package Config

import "goOwn-Ioc/Services"

type BeanConfig struct {
}

func NewBeanConfig() *BeanConfig {
	return &BeanConfig{}
}

func (this *BeanConfig) PordService() *Services.PordService {
	return Services.NewPordService()
}

func (this *BeanConfig) DBService() *Services.DBService {
	return Services.NewDBServices()
}
