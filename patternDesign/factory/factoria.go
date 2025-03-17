package factory

type Shape interface {
    Draw() string
}

type Circle struct{}

func (c *Circle) Draw() string {
    return "Drawing Circle"
}

type Square struct{}

func (s *Square) Draw() string {
    return "Drawing Square"
}

type ShapeFactory struct{}

func (sf *ShapeFactory) GetShape(shapeType string) Shape {
    switch shapeType {
    case "circle":
        return &Circle{}
    case "square":
        return &Square{}
    default:
        return nil
    }
}
