package class

import (
	"errors"
)

var (
	Earth         = element{name: "earth"}
	Wind          = element{name: "wind"}
	Fire          = element{name: "fire"}
	Water         = element{name: "water"}
	Spirit        = element{name: "spirit"}
	Order = []Type{Earth,Wind,Fire,Water,Spirit}
	stringToTypes = map[string]Type{
		"fire":   Fire,
		"water":  Water,
		"earth":  Earth,
		"wind":   Wind,
		"spirit": Spirit,
	}
)

type Type interface {
	String() string
	a()
}

type element struct {
	name string
}

func (e element) String() string {
	e.a()
	return e.name
}
func (element) a() {}

var ErrInvalidType = errors.New("invalid type")

func New(name string) (Type, error) {
	t, ok := stringToTypes[name]
	if !ok {
		return nil, ErrInvalidType
	}
	return t, nil
}
