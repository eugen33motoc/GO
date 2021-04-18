package main

import (
	"fmt"
	"math/rand"
	"time"
)

func tinta(laturaPatratului float64)float64 {
	var n int
	m:=0
	fmt.Print("numarul de trageri la tinta= ")
	fmt.Scan(&n)

	for i:=0; i<n; i++ {
		x := rand.Float64() * laturaPatratului
		y := rand.Float64() * laturaPatratului

		if (x - laturaPatratului/2)*(x - laturaPatratului/2) +
			(y - laturaPatratului/2)*(y - laturaPatratului/2) <= (laturaPatratului/2)*(laturaPatratului/2){
			m++
		}
	}

	return float64(m*4) / float64(n)
}

func main(){
	rand.Seed(time.Now().UnixNano())

	piTinta:=tinta(1000)
	fmt.Println(piTinta)


}