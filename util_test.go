package main

import (
	"reflect"
	"testing"
)

type testCategories struct {
	in  string
	out ParsedCategoryString
}

var allTestCategories = []testCategories{
	{"GL_VERSION_1_4", ParsedCategoryString{CategoryVersion, Version{1, 4}, "", ""}},
	{"GL_VERSION_2_0_DEPRECATED", ParsedCategoryString{CategoryDeprecatedVersion, Version{2, 0}, "", ""}},
	{"GL_ARB_multisample", ParsedCategoryString{CategoryExtension, Version{0, 0}, "ARB", "multisample"}},
	{"GL_EXT_12345678", ParsedCategoryString{CategoryExtension, Version{0, 0}, "EXT", "12345678"}},
}

func TestParsedCategoryString(t *testing.T) {
	for i := range allTestCategories {
		te := &allTestCategories[i]
		pc, err := ParseCategoryString(te.in)
		if err != nil {
			t.Fatalf(err.String())
		}
		t.Logf("%v", te)
		if !reflect.DeepEqual(pc, te.out) {
			t.Errorf("not equal")
		}
	}
}
