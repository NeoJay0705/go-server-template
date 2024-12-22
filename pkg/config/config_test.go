package config

import "testing"

// testing

type TestData struct {
	Test string
}

func TestConfig(t *testing.T) {
	tData := new(TestData)
	err := SetConfig(tData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tData)
}
