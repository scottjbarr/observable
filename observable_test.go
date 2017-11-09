package observable

import (
	"fmt"
	"testing"
)

type observer struct {
	id       string
	messages *[]string
}

func newObserver(id string) observer {
	messages := []string{}

	return observer{
		id:       id,
		messages: &messages,
	}
}

func (l observer) Update(message interface{}) {
	fmt.Printf("received %v : %v\n", l.id, message)
	m, _ := message.(string)

	*l.messages = append(*l.messages, m)
}

func (l observer) Identifier() string {
	return l.id
}

func TestObservable(t *testing.T) {
	s := NewObservable()

	observers := []observer{}

	for i := 0; i < 2; i++ {
		id := fmt.Sprintf("%v", i)
		l := newObserver(id)

		s.Register(l)

		observers = append(observers, l)
	}

	s.Notify("hello")

	s.Deregister(observers[0])

	s.Notify("hello again")

	i0 := len(*observers[0].messages)
	if i0 != 1 {
		t.Errorf("observer[%v] : got len=%v, want len=%v", 0, i0, 1)
	}

	i1 := len(*observers[1].messages)
	if i1 != 2 {
		t.Errorf("observer[%v] : got len=%v, want len=%v", 1, i1, 2)
	}
}

func TestObservable_Deregister(t *testing.T) {
	s := NewObservable()

	o := observer{}

	s.Register(o)

	if len(s.observers) != 1 {
		t.Errorf("Got %v, want %v", len(s.observers), 1)
	}

	s.Deregister(o)

	if len(s.observers) != 0 {
		t.Errorf("Got %v, want %v", len(s.observers), 0)
	}

	p := observer{
		id: "nope",
	}

	s.Deregister(p)

	if len(s.observers) != 0 {
		t.Errorf("Got %v, want %v", len(s.observers), 0)
	}
}
