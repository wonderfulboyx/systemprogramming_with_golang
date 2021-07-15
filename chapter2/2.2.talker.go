package main

import (
	"fmt"
)

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("hello, my name is %s\n", g.name)
}

func GreeterMain() {
	talker := &Greeter{name: "wonderfulboyx"}
	talker.Talk()
}
