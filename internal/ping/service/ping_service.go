package service

type PingServiceAPI interface {
	Ping() string
}

type PingServiceImpl struct{}

func (p *PingServiceImpl) Ping() string {
	return "meow"
}

func NewPingService() *PingServiceImpl {
	return &PingServiceImpl{}
}