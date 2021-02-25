package Services

type AdminService struct {
	// Pord *PordService `injector:"BeanConfig.PordService()"`
	// Pord *PordService `injector:"-"`
	Pord IPord `injector:"-"`
}

func NewAdminService() *AdminService {
	return &AdminService{}
}
