package Services

import "fmt"

type UserService struct {
	// Pord *PordService `injector:"BeanConfig.PordService()"`
	Pord *PordService `injector:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) UserInfo() {
	fmt.Println("this is Userinfo test successful !~")
}
