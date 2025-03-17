package builder

type Car struct {
    wheels  int
    color   string
    brand   string
    airbags bool
}

type CarBuilder struct {
    car *Car
}

func NewCarBuilder() *CarBuilder {
    return &CarBuilder{&Car{}}
}

func (cb *CarBuilder) SetWheels(wheels int) *CarBuilder {
    cb.car.wheels = wheels
    return cb
}

func (cb *CarBuilder) SetColor(color string) *CarBuilder {
    cb.car.color = color
    return cb
}

func (cb *CarBuilder) SetBrand(brand string) *CarBuilder {
    cb.car.brand = brand
    return cb
}

func (cb *CarBuilder) SetAirbags(airbags bool) *CarBuilder {
    cb.car.airbags = airbags
    return cb
}

func (cb *CarBuilder) Build() *Car {
    return cb.car
}
