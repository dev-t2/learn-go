package main

import "testing"

func TestTwoElements(t *testing.T) {
	list := []string{"Austin", "Chloe"}

	if JoinWithCommas(list) != "Austin and Chloe" {
		t.Error("Test Two Elements Error")
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"Austin", "Chloe", "Sally"}

	if JoinWithCommas(list) != "Austin, Chloe, and Sally" {
		t.Error("Test Three Elements Error")
	}
}