package main

import "fmt"

//=========interfaces=========//
type Mover interface {
	Move()
}

type Speaker interface {
	Speak()
}

type MoverSpeaker interface {
	Mover
	Speaker
}

//========structures========//
type Human struct {
	name string
}

type Tiger struct {
	name string
}

//======methods======//
func (h *Human) Speak() {
	fmt.Println("Hello")
}

func (h *Human) Move() {
	fmt.Println("Walk")
}

func (h *Tiger) Speak() {
	fmt.Println("Grrrr")
}

func (h *Tiger) Move() {
	fmt.Println("Run")
}

//======functions======//
func Actions(ms MoverSpeaker) {
	ms.Move()
	ms.Speak()
}

//======main======//
func main() {
	h := Human{"Human"}
	t := Tiger{"Tiger"}
	Actions(&h)
	Actions(&t)
}
