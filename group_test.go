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

	e["VERSION_1_3"] = Enums{}
	e["VERSION_1_3_DEPRECATED"] = Enums{}

	e["VERSION_2_1"] = Enums{}	
	e["VERSION_2_1_DEPRECATED"] = Enums{}

	e["VERSION_2_3"] = Enums{}	
	e["VERSION_2_3_DEPRECATED"] = Enums{}

	e["VERSION_3_1"] = Enums{}	

	e["EXT_1"] = Enums{}
	e["NV_1"] = Enums{}
	e["ATI_1"] = Enums{}

	f["EXT_1"] = []*Function{}
	f["NV_1"] = []*Function{}
	f["ATI_1"] = []*Function{}
	
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
	
	checkEnumCats(p, t, "nv", "NV_1")
	checkEnumCats(p, t, "ati", "ATI_1")
	checkEnumCats(p, t, "ext", "EXT_1")
	checkEnumCats(p, t, "gl13", "VERSION_1_3", "VERSION_1_3_DEPRECATED")
	checkEnumCats(p, t, "gl21",
		"VERSION_1_3", "VERSION_1_3_DEPRECATED", 
		"VERSION_2_1", "VERSION_2_1_DEPRECATED")
	checkEnumCats(p, t, "gl23",
		"VERSION_1_3", "VERSION_1_3_DEPRECATED", 
		"VERSION_2_1", "VERSION_2_1_DEPRECATED",
		"VERSION_2_3", "VERSION_2_3_DEPRECATED")
	checkEnumCats(p, t, "gl31",
		"VERSION_1_3", 
		"VERSION_2_1",
		"VERSION_2_3",
		"VERSION_3_1")
	checkEnumCats(p, t, "gl31c",
		"VERSION_1_3", "VERSION_1_3_DEPRECATED", 
		"VERSION_2_1", "VERSION_2_1_DEPRECATED",
		"VERSION_2_3", "VERSION_2_3_DEPRECATED",
		"VERSION_3_1")
	
	checkFuncCats(p, t, "nv", "NV_1")
	checkFuncCats(p, t, "ati", "ATI_1")
	checkFuncCats(p, t, "ext", "EXT_1")
}
