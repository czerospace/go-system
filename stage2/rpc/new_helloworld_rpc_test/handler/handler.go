package handler

// 解决名称冲突的问题
const HelloServiceName = "handler/HelloService"

type NewHelloService struct{}

func (s *NewHelloService) Hello(request string, reply *string) error {
	// 返回值事通过修改 reply 的值
	*reply = "Hello," + request
	return nil
}
