package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var PII = 3.14159265359

func main() {
	var L, m, n, result float64

	n = 1000000
	m = 0

	fmt.Print("L = ")
	fmt.Scan(&L)

	rand.Seed(time.Now().UnixNano())

	m = 0

	for i := 0.0; i <= n; i++ {

		y := rand.Float64() * L
		alpha := rand.Float64() * (math.Pi / 2.0)

		y1 := y + L*math.Sin(alpha)

		if int(y/L) != int(y1/L) {
			m++
		}

	}

	result = 2.0 * float64(n) / float64(m)

	fmt.Println(string("\033[36m"), "PII =", result)
	fmt.Println(string("\033[37m"), "PII =", PII)

}
