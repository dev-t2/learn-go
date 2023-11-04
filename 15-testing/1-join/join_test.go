package main

import "testing"

func TestTwoElements(t *testing.T) {
	list := []string{"Austin", "Chloe"}
	returnValue := JoinWithCommas(list)
	resultValue := "Austin and Chloe"

	if returnValue != resultValue {
		t.Errorf("JoinWithCommas(%#v) => \"%s\" != \"%s\"", list, returnValue, resultValue)
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"Austin", "Chloe", "Sally"}
	returnValue := JoinWithCommas(list)
	resultValue := "Austin, Chloe, and Sally"

	if returnValue != resultValue {
		t.Errorf("JoinWithCommas(%#v) => \"%s\" != \"%s\"", list, returnValue, resultValue)
	}
}