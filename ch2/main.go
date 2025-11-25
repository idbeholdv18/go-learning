package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"go-learning/ch2/tempconv"
)

var unit = flag.String("u", "c", "\"c\" for fahrenheit to celsius or \"f\" for celsius to fahrenheit")

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		var f = tempconv.Fahrenheit(t)
		var c = tempconv.Celsius(t)
		var k = tempconv.Kelvin(t)

		fmt.Printf("%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n",
			c, tempconv.CToF(c), c, tempconv.CToK(c),
			f, tempconv.FToK(f), f, tempconv.FToC(f),
			k, tempconv.KToC(k), k, tempconv.KToF(k))
	}
}
