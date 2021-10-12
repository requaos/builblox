package builder

import (
	"github.com/requaos/builblox/internal/blox"
	"github.com/requaos/builblox/internal/class"
)

type Container interface {
	Open() blox.Blox
	Name() string
	Level() int
}

type container struct {
	name  string
	level int
	b     blox.Blox
}

func (e *container) Open() blox.Blox {
	return e.b
}

func (e *container) Name() string {
	var s string
	if v, ok := lookup[e.name]; ok {
		s = v
	}
	return s
}

func (e *container) Level() int {
	return e.level
}

func Build(blox blox.Blox) Container {
	c := new(container)
	c.b = blox
	contents := c.b.Contents()
	for _, t := range class.Order {
		if v, ok := contents[t]; ok {
			c.level += v
			c.name += t.String()
		}
	}
	return c
}

var lookup = map[string]string{
	"earthwindfirewaterspirit": "orb",
	"earthwindfirewater":       "cross",
	"earthwindfirespirit":      "lightning",
	"earthwindwaterspirit":     "wave",
	"earthfirewaterspirit":     "steam",
	"windfirewaterspirit":      "cyclone",
	"earthwindfire":            "storm",
	"earthwindwater":           "hail",
	"earthwindspirit":          "ghost",
	"earthfirewater":           "sacrifice",
	"earthfirespirit":          "tomb",
	"earthwaterspirit":         "mud",
	"windfirewater":            "energy",
	"windfirespirit":           "chakra",
	"windwaterspirit":          "eel",
	"firewaterspirit":          "booze",
	"earthwind":                "dust",
	"earthfire":                "clay",
	"earthwater":               "slip",
	"earthspirit":              "gaia",
	"windfire":                 "raze",
	"windwater":                "splash",
	"windspirit":               "aero",
	"firewater":                "crystal",
	"firespirit":               "lava",
	"waterspirit":              "mermaid",
}
