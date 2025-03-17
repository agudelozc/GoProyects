package decorator

type Component interface {
    Operation() string
}

type ConcreteComponent struct{}

func (cc *ConcreteComponent) Operation() string {
    return "ConcreteComponent operation"
}

type Decorator interface {
    Component
}

type ConcreteDecoratorA struct {
    component Component
}

func NewConcreteDecoratorA(component Component) *ConcreteDecoratorA {
    return &ConcreteDecoratorA{component}
}

func (cda *ConcreteDecoratorA) Operation() string {
    return cda.component.Operation() + " with added behavior A"
}

type ConcreteDecoratorB struct {
    component Component
}

func NewConcreteDecoratorB(component Component) *ConcreteDecoratorB {
    return &ConcreteDecoratorB{component}
}

func (cdb *ConcreteDecoratorB) Operation() string {
    return cdb.component.Operation() + " with added behavior B"
}
