package blox

import (
	"github.com/requaos/builblox/internal/class"
)

type Blox interface {
	Class() string
	String() string
	Merge(...Blox) (Blox, error)
	Split() ([]Blox, error)
	Contents() map[class.Type]int
	Destroy()
}

type exampleBlox struct {
	class string
	data  map[class.Type]int
}

func New(name class.Type) *exampleBlox {
	return &exampleBlox{
		class: name.String(),
		data: map[class.Type]int{
			name: 1,
		},
	}
}

func (e *exampleBlox) Class() string {
	return e.class
}

func (e *exampleBlox) String() string {
	return e.Class()
}

func (e *exampleBlox) Contents() map[class.Type]int {
	return e.data
}

func (e *exampleBlox) Destroy() {
	e.data = nil
}

func (e *exampleBlox) Merge(with ...Blox) (Blox, error) {
	for _, v := range with {
		for name, count := range v.Contents() {
			e.data[name] += count
		}
		v.Destroy()
	}
	return e, nil
}

func (e *exampleBlox) Split() ([]Blox, error) {
	ret := make([]Blox, 0, 2*len(e.data))
	for name, count := range e.data {
		for i := 0; i < count; i++ {
			ret = append(ret, New(name))
		}
	}
	return ret, nil
}
