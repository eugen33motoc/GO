package main;

import (
	"bufio"
	"fmt"
	"os"

	r "./spline"
)

func main() {
	var n int
	var x, y float64
	var vx, vy []float64

	file, _ := os.Open("input.txt")
	file2, _ := os.Create("output.obj")
	w := bufio.NewWriter(file2)

	fmt.Fscan(file, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(file, &x)
		fmt.Fscan(file, &y)
		vx = append(vx, x)
		vy = append(vy, y)
	}

	p := 20
	h := (vx[n-1] - vx[0]) / float64(p)

	s := NewNaturalCubicSpline(vx, vy, 0, 0)

	/*for i := 0; i < n; i++ {
		fmt.Fprintf(w, "v %f %f %f \n", vx[i], vy[i], 50.0)
	}*/

	for i := 0; i <= p; i++ {
		xt := vx[0] + float64(i)*h
		yt := s.At(xt)
		fmt.Fprintf(w, "v %f %f %f \n", xt, yt, 50.0)
	}

	w.Flush()

}
