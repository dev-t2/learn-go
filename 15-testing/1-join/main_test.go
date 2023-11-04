package main

import (
	"fmt"
	"testing"
)

type TestData struct {
	list 				[]string
	resultValue string
}

func errorString(list []string, returnValue, resultValue string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) => \"%s\" != \"%s\"", list, returnValue, resultValue) 
}

func TestOneElement(t *testing.T) {
	list := []string{"Austin"}
	returnValue := JoinWithCommas(list)
	resultValue := "Austin"

	if returnValue != resultValue {
		t.Errorf(errorString(list, returnValue, resultValue))
	}
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

func TestJoinWithCommas(t * testing.T) {
	tests := []TestData{
		{list: []string{"Austin"}, resultValue: "Austin"},
		{list: []string{"Austin", "Chloe"}, resultValue: "Austin and Chloe"},
		{list: []string{"Austin", "Chloe", "Sally"}, resultValue: "Austin, Chloe, and Sally"},
	}

	for _, test := range tests {
		returnValue := JoinWithCommas(test.list)

		if returnValue != test.resultValue {
			t.Errorf(errorString(test.list, returnValue, test.resultValue))
		}
	}
}