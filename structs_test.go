package main

import (
	"testing"
)

type versionTest struct {
	In    string
	Out   Version
	Err   bool
}

var versionTests = []versionTest {
	{"0.0", Version{0, 0}, false},
	{"0.1", Version{0, 1}, false},
	{"1.0", Version{1, 0}, false},
	{"5.6", Version{5, 6}, false},
	{"0.",  Version{0, 0}, true},
	{"",    Version{0, 0}, true},
	{".",   Version{0, 0}, true},
	{"a",   Version{0, 0}, true},
}

func TestVersion(t *testing.T) {
	for i := range versionTests {
		test := &versionTests[i]
		v, err := MakeVersionFromString(test.In)
		if err != nil != test.Err {
			t.Errorf("Converison failed %v, %v", test, err)
		}
		if v.Valid() != test.Out.Valid() {
			t.Errorf("Input and output must be valid or invalid: %v", test)
		}
		if v.Compare(test.Out) != 0 {
			t.Errorf("Input != output %v", test)
		}
	}
}

