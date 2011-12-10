// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"testing"
)

func checkEnumCats(paks Packages, t *testing.T, pak string, cats... string) {
	if p, ok := paks[pak]; ok {
		if len(cats) != len(p.Enums) {
			t.Errorf("len(cats) != len(p.Enums)")
			return
		}
		for _, c := range cats {
			if _, ok := p.Enums[c]; !ok {
				t.Errorf("Enum not found: %v::%v", pak, c)
				return
			}
		}
		return
	}
	t.Errorf("Package not found: %v", pak)
}

func checkFuncCats(paks Packages, t *testing.T, pak string, cats... string) {
	if p, ok := paks[pak]; ok {
		if len(cats) != len(p.Functions) {
			t.Errorf("len(cats) != len(p.Functions)")
			return
		}
		for _, c := range cats {
			if _, ok := p.Functions[c]; !ok {
				t.Errorf("Function not found: %v::%v", pak, c)
				return
			}
		}
		return
	}
	t.Errorf("Package not found: %v", pak)
}

func TestPackageGrouping(t *testing.T) {
	e := make(EnumCategories)
	f := make(FunctionCategories)

	e["GL_VERSION_1_3"] = Enums{}
	e["GL_VERSION_1_3_DEPRECATED"] = Enums{}

	e["GL_VERSION_2_1"] = Enums{}	
	e["GL_VERSION_2_1_DEPRECATED"] = Enums{}

	e["GL_VERSION_2_3"] = Enums{}	
	e["GL_VERSION_2_3_DEPRECATED"] = Enums{}

	e["GL_VERSION_3_1"] = Enums{}	

	e["GL_EXT_1"] = Enums{}
	e["GL_NV_1"] = Enums{}
	e["GL_ATI_1"] = Enums{}

	f["GL_EXT_1"] = []*Function{}
	f["GL_NV_1"] = []*Function{}
	f["GL_ATI_1"] = []*Function{}
	
	suppV := []Version{{1,3}, {2,1}, {2,3}, {3,1}}
	deprV := []Version{{3,1}}
	
	p := GroupEnumsAndFunctions(e, f, 
		func (category string) (packageNames []string) { 
			return GroupPackagesByVendorFunc(category, suppV, deprV)
		})
	
	for n, pa := range p {
		t.Logf("%s:\n %s\n", n, pa)
	}

	if len(p) != 8 {
		t.Errorf("Wrong number of categories.")
	}
	
	checkEnumCats(p, t, "nv", "GL_NV_1")
	checkEnumCats(p, t, "ati", "GL_ATI_1")
	checkEnumCats(p, t, "ext", "GL_EXT_1")
	checkEnumCats(p, t, "gl13", "GL_VERSION_1_3", "GL_VERSION_1_3_DEPRECATED")
	checkEnumCats(p, t, "gl21",
		"GL_VERSION_1_3", "GL_VERSION_1_3_DEPRECATED", 
		"GL_VERSION_2_1", "GL_VERSION_2_1_DEPRECATED")
	checkEnumCats(p, t, "gl23",
		"GL_VERSION_1_3", "GL_VERSION_1_3_DEPRECATED", 
		"GL_VERSION_2_1", "GL_VERSION_2_1_DEPRECATED",
		"GL_VERSION_2_3", "GL_VERSION_2_3_DEPRECATED")
	checkEnumCats(p, t, "gl31",
		"GL_VERSION_1_3", 
		"GL_VERSION_2_1",
		"GL_VERSION_2_3",
		"GL_VERSION_3_1")
	checkEnumCats(p, t, "gl31c",
		"GL_VERSION_1_3", "GL_VERSION_1_3_DEPRECATED", 
		"GL_VERSION_2_1", "GL_VERSION_2_1_DEPRECATED",
		"GL_VERSION_2_3", "GL_VERSION_2_3_DEPRECATED",
		"GL_VERSION_3_1")
	
	checkFuncCats(p, t, "nv", "GL_NV_1")
	checkFuncCats(p, t, "ati", "GL_ATI_1")
	checkFuncCats(p, t, "ext", "GL_EXT_1")
}
