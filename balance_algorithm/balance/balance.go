package balance

// 只要实现了该接口的方法，就可以调用该接口的方法
type Balancer interface {
	DoBalance([]*Instance, ...string) (*Instance, error)
}
