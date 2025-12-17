package main

import "fmt"

const (
	January = iota
	February
	March
	April
	May
	June
	Jule
	August
	September
	October
	November
	December
)

func main() {
	var months [12]int = [12]int{
		January:   31,
		February:  28,
		March:     31,
		April:     30,
		May:       31,
		June:      30,
		Jule:      31,
		August:    31,
		September: 30,
		October:   31,
		November:  30,
		December:  31,
	}

	fmt.Println(calculateDays(&months))
}

func calculateDays(months *[12]int) int {
	var days int
	for month := range months {
		days += months[month]
	}
	return days
}
