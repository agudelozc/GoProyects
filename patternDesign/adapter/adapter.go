package adapter

type Target interface {
    Request() string
}

type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
    return "Adaptee specific request"
    
}

type Adapter struct {
    adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) *Adapter {
    return &Adapter{adaptee}
}

func (a *Adapter) Request() string {
    return a.adaptee.SpecificRequest()
}
