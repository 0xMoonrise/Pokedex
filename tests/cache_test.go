package test

import (
    "testing"
    "github.com/0xMoonrise/pokedexcli/src"
)

func TestCacheReturnSomething(t *testing.T) {
    result   := src.NewCache("its okay!")
	expected := "its okay!"
    if result != expected {
        t.Errorf("Expected %v, but got %v", expected, result)
    }
}

func TestCacheReturnNothing(t *testing.T) {
    result   := src.NewCache("")
    expected := ""
    if result != expected {
        t.Errorf("\nExpected '%v', but got '%v'", expected, result)
    }
}
