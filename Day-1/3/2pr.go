package main

import (
	"fmt"
	"math"
)

func main() {
	var r, C float64 = 25.0, 0.0
	C = 2 * r * math.Pi
	C = math.Round(C*100) / 100
	fmt.Println("Длинна окружности с радиусом:", r, "равна:", C)
}
