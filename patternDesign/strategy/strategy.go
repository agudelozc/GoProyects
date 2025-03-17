package strategy


//interfaz llamada Strategy que tiene un método Execute que devuelve una cadena.
type Strategy interface {
    Execute() string
}
//Estas estructuras representan estrategias específicas.
type ConcreteStrategyA struct{}

//Estos son los métodos Execute para las estructuras ConcreteStrategyA. Estos métodos implementan la interfaz Strategy.
func (sa *ConcreteStrategyA) Execute() string {
    return "Strategy A"
}
//Estas estructuras representan estrategias específicas.
type ConcreteStrategyB struct{}

//Estos son los métodos Execute para las estructuras ConcreteStrategyB. Estos métodos implementan la interfaz Strategy.
func (sb *ConcreteStrategyB) Execute() string {
    return "Strategy B"
}

type Context struct {
    strategy Strategy
}

func NewContext(strategy Strategy) *Context {
    return &Context{strategy}
}
//Este método permite cambiar la estrategia del contexto.
func (c *Context) SetStrategy(strategy Strategy) {
    c.strategy = strategy
}
//Este método ejecuta la estrategia actual del contexto.
func (c *Context) ExecuteStrategy() string {
    return c.strategy.Execute()
}
