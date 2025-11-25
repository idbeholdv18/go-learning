package main

import "fmt"

func main() {
	var i int;

	i = increment(&i);
	increment(&i);
	v := increment(&i);

	fmt.Println(i);
	fmt.Println(v);
	fmt.Println(&v == &i);
}

func increment(v *int) int {
	*v++

	return *v;
}
