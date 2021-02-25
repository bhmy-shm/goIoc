package Bean

import (
	"fmt"
	"reflect"
)

//该文件是IOC容器文件，专门用来存储各种类型

type BeanMapper map[reflect.Type]reflect.Value

func NewBeanMapper() BeanMapper {
	return BeanMapper(make(map[reflect.Type]reflect.Value))
}

func (this BeanMapper) add(bean interface{}) {
	beantype := reflect.TypeOf(bean)

	if beantype.Kind() != reflect.Ptr {
		return
	}
	this[beantype] = reflect.ValueOf(bean)
}

func (this BeanMapper) get(bean interface{}) reflect.Value {

	beantype := reflect.TypeOf(bean)

	//判断 bean的类型，使用接口类型断言
	if bt, ok := bean.(reflect.Type); ok { //如果传入的 bean是type类型
		beantype = bt
	} else { //如果不是，则把它变成反射的 type
		reflect.TypeOf(bean)
	}

	if v, ok := this[beantype]; ok {
		return v
	}

	//处理接口注入
	//遍历 Mapper，判断Mapper的key 是否为 接口类型
	for k, v := range this {
		if k.Implements(beantype) {
			return v //如果是接口类型，反馈value
		}
	}
	return reflect.Value{} //如果不是依旧反馈 空Value
}

func (this BeanMapper) each() {
	for i, v := range this {
		fmt.Println(i, v)
	}
}
