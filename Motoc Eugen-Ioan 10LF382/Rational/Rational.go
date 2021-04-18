package main

import (
	"fmt"
	"math"
	"strconv"
)

type Rational struct {
	Numa, Numi int
}

func (r Rational) print(str string) {
	fmt.Printf("%s%d%s%d%s", str, r.Numa, "/", r.Numi, "\n")
}

func (r Rational) getNuma() int {
	return r.Numa
}

func (r Rational) getNumi() int {
	return r.Numi
}

func (r Rational) newInstance(x, y int) Rational {
	t := Rational{x, y}
	return t
}

func (r Rational) inverse() Rational {
	r.Numa, r.Numi = r.Numi, r.Numa
	return r
}

func (r Rational) add(t Rational) Rational {
	r.Numa = (r.Numa * t.Numi) + (r.Numi * t.Numa)
	r.Numi = r.Numi * t.Numi
	return r.simplify()
}

func (r Rational) substract(t Rational) Rational {
	r.Numa = (r.Numa * t.Numi) - (r.Numi * t.Numa)
	r.Numi = r.Numi * t.Numi
	return r.simplify()
}

func (r Rational) multiply(t Rational) Rational {
	r.Numa = r.Numa * t.Numa
	r.Numi = r.Numi * t.Numi
	return r.simplify()
}

func (r Rational) divide(t Rational) Rational {
	r = r.multiply(t.inverse())
	return r.simplify()
}

func (r Rational) addInt(t int) Rational {
	r.Numa = r.Numa + t*r.Numi
	return r.simplify()
}

func (r Rational) substractInt(t int) Rational {
	r.Numa = r.Numa - t*r.Numi
	return r.simplify()
}

func (r Rational) multiplyInt(t int) Rational {
	r.Numa *= t
	return r.simplify()
}

func (r Rational) divideInt(t int) Rational {
	r.Numi *= t
	return r.simplify()
}

func (r Rational) addNuma(t int) Rational {
	r.Numa += t
	return r
}

func (r Rational) addNumi(t int) Rational {
	r.Numi += t
	return r
}

func (r Rational) addNumaAndNumi(t int) Rational {
	r.Numa += t
	r.Numi += t
	return r
}

func (r Rational) substractNuma(t int) Rational {
	r.Numa -= t
	return r
}

func (r Rational) substractNumi(t int) Rational {
	r.Numi = r.Numi - t
	if r.Numi < 0 {
		r.Numa *= (-1)
	}
	return r
}

func (r Rational) substractNumaAndNumi(t int) Rational {
	r.Numa = r.Numa - t
	r.Numi = r.Numi - t
	if r.Numi < 0 {
		r.Numa *= (-1)
	}
	return r
}

func (r Rational) isNull() bool {
	return r.Numa == 0
}

func (r Rational) getRealValue() float32 {
	var realValue = float32(r.Numa) / float32(r.Numi)
	return realValue
}

func (r Rational) getAbsValue() Rational {
	if r.Numa < 0 {
		r.Numa *= (-1)
	}
	if r.Numi < 0 {
		r.Numi *= (-1)
	}
	return r
}

func (r Rational) simplify() Rational {
	c := cmmdc(r.Numa, r.Numi)
	r.Numa /= c
	r.Numi /= c

	if r.Numi < 0 {
		r.Numa *= -1
		r.Numi *= -1
	}

	return r
}

func (r Rational) pow(n int) Rational {
	r.Numa = int(math.Pow(float64(r.Numa), float64(n)))
	r.Numi = int(math.Pow(float64(r.Numi), float64(n)))
	return r
}

func (r Rational) biggerThan(t Rational) bool {
	if r.Numi == t.Numi {
		return r.Numa > t.Numa
	} else {
		if r.Numa == t.Numa {
			return r.Numi < t.Numi
		} else {
			var x = (r.Numi * t.Numi) / cmmdc(r.Numi, t.Numi)
			var u, v Rational

			u = u.newInstance(r.Numa*(x/r.Numi), r.Numi*(x/r.Numi))
			v = v.newInstance(t.Numa*(x/t.Numi), t.Numi*(x/t.Numi))

			return u.biggerThan(v)
		}
	}

	return false
}

func (r Rational) smallerThan(t Rational) bool {
	if r.equals(t) {
		return false
	} else {
		if r.biggerThan(t) {
			return false
		} else {
			return true
		}
	}
	return false
}

func (r Rational) equals(t Rational) bool {
	if r.Numa == t.Numa && r.Numi == t.Numi {
		return true
	}
	return false
}

func (r Rational) isNatural() bool {
	return r.Numi == 1
}


func getFromFloat32(x float32) Rational {
	//number := fmt.Sprintf("%f", x)
	number := strconv.FormatFloat(float64(x), 'f', -1, 32)
	fmt.Println("String(", x, ") = \"", number, "\"")
	var fr = false
	var numarator = 0
	var numitor = 1
	for i := 0; i < len(number); i++ {
		if fr == true {
			numitor = numitor * 10
		}
		if number[i] == '.' || number[i] == ',' {
			fr = true
		} else {
			var digit = number[i] - '0'
			numarator = numarator*10 + int(digit)
		}
	}
	return Rational{numarator, numitor}
}

func cmmdc(a, b int) int {
	for a%b != 0 {
		r := a % b
		a = b
		b = r
	}
	return b
}

func main() {
	var num1, num2, suma, diferenta, produs, divide Rational
	var n_int int
	var x_float float32

	fmt.Print("Primul numarator: ")
	fmt.Scan(&num1.Numa)
	fmt.Print("Primul numitor:  ")
	fmt.Scan(&num1.Numi)
	fmt.Print("Al doilea numarator: ")
	fmt.Scan(&num2.Numa)
	fmt.Print("Al doilea numitor:   ")
	fmt.Scan(&num2.Numi)

	fmt.Print("n_int = ")
	fmt.Scan(&n_int)

	fmt.Print("x_float = ")
	fmt.Scan(&x_float)

	fmt.Println()

	num1.print("num1 = ")
	num2.print("num2 = ")

	suma = num1.add(num2)
	suma.print("num1+num2 = ")

	diferenta = num1.substract(num2)
	diferenta.print("num1-num2 = ")

	produs = num1.multiply(num2)
	produs.print("num1*num2 = ")

	divide = num1.divide(num2)
	divide.print("num1/num2 = ")

	fmt.Println()

	num1 = num1.addInt(n_int)
	num1.print("num1+n_int = ")

	num1 = num1.substractInt(n_int)
	num1.print("num1-n_int = ")

	num1 = num1.multiplyInt(n_int)
	num1.print("num1*n_int = ")

	num1 = num1.divideInt(n_int)
	num1.print("num1/n_int = ")

	num1 = num1.pow(n_int)
	num1.print("num1^n_int = ")

	fmt.Println()

	getFromFloat32(x_float).print("x_float = ")
	getFromFloat32(x_float).simplify().print("x_float = ")

}
