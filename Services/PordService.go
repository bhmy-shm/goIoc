package Services

import "fmt"

type IPord interface {
	Name() string
}

const (
	version = "v1.0"
)

type PordService struct {
	Version string
	DB      *DBService `injector:"-"`
}

func NewPordService() *PordService {
	return &PordService{Version: "v1.0"}
}

func (this *PordService) PordInfo() {
	fmt.Println("this is podinfo test successful !~")
}
func (this *PordService) Name() string {
	name := "sunhaiming"
	return name
}
