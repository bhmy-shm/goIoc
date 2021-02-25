package Bean

import (
	"fmt"
	"reflect"

	"github.com/shenyisyn/goft-expr/src/expr"
)

//定义一个 BeanFactory 的全局变量
var BeanFactory *beanFactoryImpl

//初始化 BeanFactory
func init() {
	BeanFactory = NewBeanFactory()
}

type beanFactoryImpl struct {
	Mapper  BeanMapper
	ExprMap map[string]interface{}
}

func NewBeanFactory() *beanFactoryImpl {
	return &beanFactoryImpl{Mapper: make(BeanMapper), ExprMap: make(map[string]interface{})}
}

//1.IOC容器类型依赖注入
func (this *beanFactoryImpl) Set(vs ...interface{}) {

	if vs == nil || len(vs) == 0 {
		return
	}
	for _, v := range vs {
		this.Mapper.add(v)
	}
}

//2.IOC容器依赖查询
func (this *beanFactoryImpl) Get(v interface{}) interface{} {

	if v == nil {
		return nil
	}
	getvalue := this.Mapper.get(v)

	//返回getvalue拿到的值
	if getvalue.IsValid() {
		return getvalue.Interface()
	}
	return nil
}

//4.IOC - Config 表达式注入
func (this *beanFactoryImpl) Config(cfs ...interface{}) {
	//判断传入的是不是空值
	if cfs == nil || len(cfs) == 0 {
		return
	}
	//判断传入的是不是错误值
	for _, cfg := range cfs {
		t := reflect.TypeOf(cfg)
		if t.Kind() != reflect.Ptr {
			panic("cfgs must ptr")
		}
		//如果都不是则把 cfg 本身写入到类型存储当中
		this.Set(cfg)
		//再构建表达式map
		this.ExprMap[t.Elem().Name()] = cfg

		//获取BeanConfig 的方法
		v := reflect.ValueOf(cfg)
		for i := 0; i < v.NumMethod(); i++ {
			Callmethod := v.Method(i).Call(nil) //调用每一个方法

			fmt.Println("--------------------")
			//如果存在方法调用，并且方法的长度 >=1,则将方法的执行结果写入 Mapper
			if Callmethod != nil && len(Callmethod) >= 1 {
				this.Set(Callmethod[0].Interface())
			}
		}
	}
}
func (this *beanFactoryImpl) CleanConfig() {
	this.ExprMap = make(map[string]interface{})
}

//3.IOC依赖注入（重要）
func (this *beanFactoryImpl) Apply(bean interface{}) {
	if bean == nil {
		return
	}
	//1.反射判断是不是ptr，struct，如果是则提取它的 tag ，然后实现依赖关系注入
	v := reflect.ValueOf(bean)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}

	//2.遍历结构体，获取 tag 标签
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i) //拿到结构体的字段类型

		//先判断字段首字母是不是大写,再判断tag标签是不是空的
		if v.Field(i).CanSet() && field.Tag.Get("injector") != "" {

			if field.Tag.Get("injector") == "-" { //多例注入
				getv := this.Get(field.Type)
				v.Field(i).Set(reflect.ValueOf(getv))
				this.Apply(getv) //循环注入
			} else { //单例注入
				fmt.Println("使用了表达式的方法")
				ret := expr.BeanExpr(field.Tag.Get("injector"), this.ExprMap)
				if ret != nil && !ret.IsEmpty() {
					retValue := ret[0]
					if retValue != nil {
						this.Set(retValue)
						v.Field(i).Set(reflect.ValueOf(retValue))
						this.Apply(retValue) //循环注入
					}
				}
			}
		}
	}
}

func (this *beanFactoryImpl) Each() {
	this.Mapper.each()
}
