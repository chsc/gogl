// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"testing"
	_ "reflect"
)

func TestPackageGrouping(t *testing.T) {
	e := make(EnumCategories)
	f := make(FunctionCategories)

	e["GL_VERSION_1_3"] = Enums{}
	e["GL_VERSION_1_3_DEPRECATED"] = Enums{}

	e["GL_VERSION_2_1"] = Enums{}	
	e["GL_VERSION_2_1_DEPRECATED"] = Enums{}

	e["GL_VERSION_3_1"] = Enums{}	

	e["GL_EXT_1"] = Enums{}
	e["GL_NV_1"] = Enums{}
	e["GL_ATI_1"] = Enums{}

	f["GL_EXT_1"] = []*Function{}
	f["GL_NV_1"] = []*Function{}
	f["GL_ATI_1"] = []*Function{}
	
	suppV := []Version{{1,0}, {1,3}, {2, 3}, {3,1}}
	deprV := []Version{{3,1}}
	
	p := GroupEnumsAndFunctions(e, f, 
		func (category string) (packageNames []string) { 
			return GroupPackagesByVendorFunc(category, suppV, deprV)
		})
	
	for n, pa := range p {
		t.Logf("%s\n %s\n", n, pa)
	}
	
	t.Fatal("Unit test not complete")
}
