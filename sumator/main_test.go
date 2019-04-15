package main

import "testing"

func TestSumator_IntSum(t *testing.T) {
	v := Sumator{3, -9}
	if v.IntSum() != -6 {
		t.Errorf("The sum is not correct, want %d, got %d", -6, v.IntSum())
	}
}

func TestSumator_FloatSum(t *testing.T) {
	v := Sumator{3, -9}
	if v.FloatSum() != -6 {
		t.Errorf("The sum is not correct, want %d, got %f", -6, v.FloatSum())
	}
}

func TestSumator_StringSum(t *testing.T) {
	v := Sumator{3, -9}
	if v.StringSum() != "-6" {
		t.Errorf("The sum is not correct, want %s, got %s", "-6", v.StringSum())
	}
}

func TestStringSum(t *testing.T) {
	if StringSum("5", "6") != 11 {
		t.Errorf("The sum is not correct, want %d, got %d", 11, StringSum("5", "6"))
	}
}
