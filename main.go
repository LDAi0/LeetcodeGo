package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)
	var freeCoups int
	freeCoups = d / 7
	d -= freeCoups
	freeCoups += freeCoups / n
	fmt.Println(freeCoups)
}

//CALCULATOR!!!!
