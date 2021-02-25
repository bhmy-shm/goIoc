package Services

type DBService struct {
	Address string
}

func NewDBServices() *DBService {
	return &DBService{Address: "127.0.0.1:3306"}
}
