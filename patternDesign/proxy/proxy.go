package proxy

type Subject interface {
    Request() string
}

type RealSubject struct{}

func (rs *RealSubject) Request() string {
    return "RealSubject request"
}

type Proxy struct {
    realSubject *RealSubject
}

func NewProxy() *Proxy {
    return &Proxy{}
}

func (p *Proxy) Request() string {
    if p.realSubject == nil {
        p.realSubject = &RealSubject{}
    }
    return "Proxy: " + p.realSubject.Request()
}
