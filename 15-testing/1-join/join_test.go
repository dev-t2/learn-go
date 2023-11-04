package main

import (
	"fmt"
	"testing"
)

func errorString(list []string, returnValue, resultValue string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) => \"%s\" != \"%s\"", list, returnValue, resultValue) 
}

func TestTwoElements(t *testing.T) {
	list := []string{"Austin", "Chloe"}
	returnValue := JoinWithCommas(list)
	resultValue := "Austin and Chloe"

	if returnValue != resultValue {
		t.Errorf(errorString(list, returnValue, resultValue))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"Austin", "Chloe", "Sally"}
	returnValue := JoinWithCommas(list)
	resultValue := "Austin, Chloe, and Sally"

	if returnValue != resultValue {
		t.Errorf(errorString(list, returnValue, resultValue))
	}
}