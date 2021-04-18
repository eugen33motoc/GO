package main

import (
	"fmt"
	"strconv"
)

type Date struct {
	day, month, year int
}

func createStringDate(d Date) string {

	var date string
	date = ""

	if d.day < 10 {
		date += "0"
	}
	date += strconv.Itoa(d.day) + "/"

	if d.month < 10 {
		date += "0"
	}
	date += strconv.Itoa(d.month) + "/"

	date += strconv.Itoa(d.year)

	return date
}

func anBisect(year int) bool {
	if year%100 == 0 && year%400 == 0 {
		return true
	}
	if year%100 != 0 && year%4 == 0 {
		return true
	}
	return false
}

var days = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func AddDay(d Date) Date {
	if d.day < days[d.month] {
		d.day++
	} else {
		if d.day == days[d.month] {
			if d.month == 2 && anBisect(d.year) {
				d.day++
			} else {
				d.day = 1
				if d.month == 12 {
					d.month = 1
					d.year++
				} else {
					d.month++
				}
			}

		} else {
			d.month++
			d.day = 1
			//situatia in care suntem pe 29 feb
		}
	}

	return d
}

func AddMonth(d Date) Date {
	if d.month < 12 {
		d.month++
	} else {
		d.month = 1
		d.year++
	}
	return d
}

func AddYear(d Date) Date {
	d.year++
	return d
}

func Compare(d1, d2 Date) int {
	if d1.year > d2.year {
		return 1
	}
	if d1.year < d2.year {
		return -1
	}
	if d1.month > d2.month {
		return 1
	}
	if d1.month < d2.month {
		return -1
	}
	if d1.day > d2.day {
		return 1
	}
	if d1.day < d2.day {
		return -1
	}
	return 0
}

func SubDate(d1, d2 Date) int {
	if Compare(d1, d2) == 1 {
		d1, d2 = d2, d1
	}
	nr := 0
	for Compare(d1, d2) == -1 {
		d1 = AddDay(d1)
		nr++
	}
	return nr
}

func SubDate2(d1, d2 Date) {

	var noOfYears, noOfMonths, noOfWeeks, noOfDays int

	noOfDays = 0
	noOfMonths = 0
	noOfWeeks = 0
	noOfYears = 0

	if Compare(d1, d2) == 1 {
		d1, d2 = d2, d1
	}

	for d1.day != d2.day {
		d1 = AddDay(d1)
		noOfDays++
	}

	for d1.month != d2.month {
		d1 = AddMonth(d1)
		noOfMonths++
	}

	for d1.year != d2.year {
		d1 = AddYear(d1)
		noOfYears++
	}

	noOfWeeks = noOfDays / 7
	noOfDays -= 7 * noOfWeeks

	fmt.Print(noOfYears, " years, ", noOfMonths, " months, ", noOfWeeks, " weeks, ", noOfDays, " days")
	fmt.Print("   ~   ", noOfYears*365+noOfMonths*30+noOfWeeks*7+noOfDays, " days")
}

func addNDays(d Date, n int) {

	fmt.Print("\n\n", createStringDate(d), " + ", n, " = ")

	for n > 0 {
		d = AddDay(d)
		n--
	}

	fmt.Print(createStringDate(d))
}

func main() {
	var d1 = Date{11, 3, 2021}
	var d2 = Date{3, 3, 1999}
	//fmt.Println(SubDate(d1, d2))

	fmt.Print("Time between ", createStringDate(d1), " and ", createStringDate(d2), " : ", SubDate(d1, d2), " days")
	fmt.Print("\nTime between ", createStringDate(d1), " and ", createStringDate(d2), " : ")
	SubDate2(d1, d2)



	addNDays(d2, 1000)
}
