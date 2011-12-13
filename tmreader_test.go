// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"strings"
	"testing"
)

var testTypeMapStr = "Test1,*,*,	    GLType1,*,*\n" +
	"Test2,*,*,	    GLType2,*,*\n" +
	"Test3,*,*,	    GLType3,*,*\n"

func checkType(tn, gltype string, tm TypeMap, t *testing.T) {
	if ty, ok := tm[tn]; ok {
		if ty == gltype {
			t.Logf("Type found: %v::%v", tn, gltype)
			return
		}
		t.Errorf("Type not found: %v::%v", tn, gltype)
		return
	}
	t.Errorf("Type not found: %v", tn)
}

func TestTypeMapReader(t *testing.T) {
	r := strings.NewReader(testTypeMapStr)
	tm, err := ReadTypeMap(r)
	if err != nil {
		t.Fatalf("Read type map failed: %v", err)
	}
	t.Logf("%v", tm)

	if len(tm) != 3 {
		t.Errorf("Wrong number of types.")
	}

	checkType("Test1", "GLType1", tm, t)
	checkType("Test2", "GLType2", tm, t)
	checkType("Test3", "GLType3", tm, t)
}
