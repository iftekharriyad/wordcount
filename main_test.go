package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	b := bytes.NewBufferString("whriewuh fskjdfhksj f \n dfhksjfh khfd fksjfh")
	exp := 44
	res := count(b, false, true)
	if res != exp {
		t.Errorf("Expected %d, got instead %d.\n", exp, res)
	}
}
