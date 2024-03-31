package main

import (
	"testing"
)

func Test_cleanCommandText(t *testing.T) {
	cases := map[string][]string{
		"hello world": {"hello", "world"},
		"hello World": {"hello", "world"},
		"HELLO world": {"hello", "world"},
	}

	for key, value := range cases {
		resulted := cleanCommandText(key)
		if len(resulted) != len(value) {
			t.Fatal("length Doesn't match")
		}

		for i := 0; i < len(resulted); i++ {
			if resulted[i] != value[i] {
				t.Fatal("words Doesn't match")
			}
		}
	}

}
