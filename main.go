package main

import (
	"fmt"
	"goOwn-Ioc/Bean"
	"goOwn-Ioc/Config"
	"goOwn-Ioc/Services"
)

func main() {
	test4()
}

func test4() {

	//1.实现表达式map的自动注入依赖
	beanConfig := Config.NewBeanConfig()
	bean := Bean.NewBeanFactory()
	bean.Config(beanConfig)

	{
		user := Services.NewUserService() //初始化user
		bean.Apply(user)                  //初始化user的同时将其依赖注入
		fmt.Println(user.Pord.DB)
	}
	// {
	// 	admin := Services.NewAdminService() //初始化user
	// 	bean.Apply(admin)                   //初始化user的同时将其依赖注入
	// 	str := admin.Pord.Name()            //查看结果
	// 	println(str)
	// }
}

func test3() {

	//1.实现表达式map的自动注入依赖
	beanConfig := Config.NewBeanConfig()
	bean := Bean.NewBeanFactory()
	bean.ExprMap = map[string]interface{}{
		"BenConfig": beanConfig,
	}

	//2.同时兼容 "-"
	pd := &Services.PordService{
		Version: "v2.0",
	}
	//3.Config方法可以同时兼容"-",以及表达式
	bean.Config(beanConfig, pd)

	//2.调用user
	{
		user := Services.NewUserService() //初始化user
		bean.Apply(user)                  //初始化user的同时将其依赖注入
		fmt.Println(user.Pord)            //查看结果
	}
}

func test2() {
	bean := Bean.NewBeanFactory()
	bean.Set(Services.NewPordService()) //提前将Pord注入IOC容器

	{
		user := Services.NewUserService() //初始化user
		bean.Apply(user)                  //初始化user的同时将其依赖注入
		fmt.Println(user.Pord)            //查看结果
	}
	bean.Each() //遍历 Mapper查看详细的 key - value
}

// func test1() {
// 	pord := Services.NewPordService()
// 	// user := Services.NewUserService(pord)

// 	fmt.Println("1", user.Pord, user.Pord.Version)
// 	user.Pord.PordInfo()
// }
