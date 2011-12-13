// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"reflect"
	"strings"
	"testing"
)

var testFuntionsStr = "###########\n" +
	"\n" +
	"# comment 1\n" +
	" ### comment 2\n" +
	"	#comment 3\n" +
	"version:	1.0 1.1 4.2\n" +
	"# comment\n" +
	"Foo1(p1, p2)\n" +
	"	return		void\n" +
	"	param		p1		type1 in value\n" +
	"	param		p2		type2 in value\n" +
	"	category	cat_1		   # comment\n" +
	"	version		1.0\n" +
	"	glxropcode	85\n" +
	"	offset		158\n" +
	"\n" +
	"Foo2(p1)\n" +
	"	return		void\n" +
	"	param		p1		type3 in value\n" +
	"	category	cat_2		   # comment\n" +
	"	version		2.1\n" +
	"	glxropcode	95\n" +
	"	offset		168\n"

func checkFunc(f *Function, fcats FunctionCategories, t *testing.T) {
	if functions, ok := fcats[f.Category]; ok {
		if foo := functions.Find(f.Name); foo != nil {
			if reflect.DeepEqual(foo, f) {
				t.Logf("Enum found: %v::%v = %v", f.Category, f.Name, f)
				return
			}
			t.Errorf("Functions not equal: %v: %v != %v", f.Category, foo, f)
			return
		}
		t.Errorf("Function not found: %v::%v", f.Category, f.Name)
		return
	}
	t.Errorf("Category not found: %v", f.Category)
}

func TestReadFunctions(t *testing.T) {
	r := strings.NewReader(testFuntionsStr)
	f, v, err := ReadFunctions(r)
	if err != nil {
		t.Fatalf("Read functions failed: %v", err)
	}
	t.Logf("%v", f)
	t.Logf("%v", v)

	if len(f) != 2 {
		t.Errorf("Wrong number of categories.")
	}

	checkFunc(&Function{"Foo1", []Parameter{{"p1", "type1", false}, {"p2", "type2", false}}, "void", Version{1, 0}, Version{0, 0}, "cat_1", 158, 85}, f, t)
	checkFunc(&Function{"Foo2", []Parameter{{"p1", "type3", false}}, "void", Version{2, 1}, Version{0, 0}, "cat_2", 168, 95}, f, t)
	// TODO: add more tests: in array, out array, glxflags, offset, ...	
}
