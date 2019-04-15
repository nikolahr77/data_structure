package main

import (
	"strconv"
)

// Sumator contains two integers that we will use to sum
type Sumator struct {
	a int
	b int
}

// IntSum sums two integers and returns them
func (s Sumator) IntSum() int {
	return s.a + s.b
}

// FloatSum sums two integers and returns a float64 number
func (s Sumator) FloatSum() float64 {
	f := float64(s.a + s.b)
	return f
}

// StringSum sums two integers and returns a string
func (s Sumator) StringSum() string {
	str := strconv.Itoa(s.a + s.b)
	return str
}

// StringSum converts two strings to int64 and sums them
func StringSum(a, b string) int64 {
	i, _ := strconv.ParseInt(a, 10, 64)
	v, _ := strconv.ParseInt(b, 10, 64)
	return i + v
}
