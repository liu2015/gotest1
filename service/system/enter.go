package system

// 结构体组合多个结构体，达到继承的效果
type ServiceGroup struct {
	JwtService
	ApiService
	MenuService
}
