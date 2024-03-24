package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}
}

func TestSub(t *testing.T) {

	result := sub(3, 1)

	if result != 2 {
		t.Error("The result must be 5")
	}
}

func TestTimes(t *testing.T) {

	result := times(3, 3)

	if result != 9 {
		t.Error("The result must be 5")
	}
}

func TestSumX(t *testing.T) {

	result := sumX(3, 2)

	if result != 8 {
		t.Error("The result must be 5")
	}
}
