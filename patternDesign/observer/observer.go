package observer

type Subject struct {
    observers []Observer
    state     string
}

//Metodo que permite adjuntar un nuevo observador al sujeto
func (s *Subject) Attach(o Observer) {
    s.observers = append(s.observers, o)
}
//Este método notifica a todos los observadores adjuntos cuando hay un cambio en el estado del sujeto.
func (s *Subject) Notify() {
    for _, observer := range s.observers {
        observer.Update()
    }
}
// Este método permite cambiar el estado del sujeto y notificar a todos los observadores sobre este cambio.
func (s *Subject) SetState(state string) {
    s.state = state
    s.Notify()
}

type Observer interface {
    Update()
}

type ConcreteObserver struct {
    Subject *Subject
}

func (co *ConcreteObserver) Update() {
    // Actualizar basado en el nuevo estado del sujeto
	co.Subject.state = "new state"
}
