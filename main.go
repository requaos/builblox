package main

import (
	"fmt"
	"github.com/requaos/builblox/internal/blox"
	"github.com/requaos/builblox/internal/builder"
	"github.com/requaos/builblox/internal/class"
	"math/rand"
)

func main() {
	rand.Seed(42)
	for cycle := 0; cycle < 15; cycle++ {
		lim := 7 * rand.Intn(50)
		b := make([]blox.Blox, 0, lim)
		for i := 0; i < lim; i++ {
			b = append(b, blox.New(class.Order[rand.Intn(len(class.Order)-cycle%(len(class.Order)-1))]))
		}
		bs, err := b[0].Merge(b[1:]...)
		if err != nil {
			fmt.Println(err)
			return
		}

		c := builder.Build(bs)
		fmt.Printf("\nType: %4s\nLevel: %-5d\nRaw:\n%#v", c.Name(), c.Level(), c.Open())
	}
}
