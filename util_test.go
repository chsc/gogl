// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"reflect"
	"testing"
)

type testCategories struct {
	in  string
	out ParsedCategory
}

type testGoName struct {
	in  string
	out string
}

var allTestCategories = []testCategories{
	{"VERSION_1_4", ParsedCategory{CategoryVersion, Version{1, 4}, "", ""}},
	{"VERSION_2_0_DEPRECATED", ParsedCategory{CategoryDepVersion, Version{2, 0}, "", ""}},
	{"ARB_multisample", ParsedCategory{CategoryExtension, Version{0, 0}, "ARB", "multisample"}},
	{"EXT_12345678", ParsedCategory{CategoryExtension, Version{0, 0}, "EXT", "12345678"}},
	{"ABC123_a_b_", ParsedCategory{CategoryExtension, Version{0, 0}, "ABC123", "a_b_"}},
	{"A_", ParsedCategory{CategoryExtension, Version{0, 0}, "A", ""}},
}

var allTestGoName = []testGoName{
	{"", ""},
	{"_", ""},
	{"ARB_multisample", "ARBMultisample"},
	{"vertex_array_object", "VertexArrayObject"},
	{"a_b_c_", "ABC"},
}

func TestParsedCategory(t *testing.T) {
	for i := range allTestCategories {
		te := &allTestCategories[i]
		pc, err := ParseCategoryString(te.in)
		if err != nil {
			t.Fatalf(err.Error())
		}
		t.Logf("%s -> %v", te.in, te.out)
		if !reflect.DeepEqual(pc, te.out) {
			t.Errorf("not equal: out = %v", pc)
		}
	}
}

func TestGoName(t *testing.T) {
	for i := range allTestGoName {
		te := &allTestGoName[i]
		if GoName(te.in) != te.out {
			t.Errorf("String conversion failed: %v -> %v", te.in, te.out)
		}
	}
}
