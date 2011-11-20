package main

import (
	"testing"
	"strings"
)

var testStr = 
	"###########\n" +
	"\n" +
	"# comment 1\n" +
	" ### comment 2\n" +
	"	#comment 3\n" +
	"cat_1 enum:\n" +
	"# comment\n" +
	"enum1	 = 0x00000100 	\n" +
	"enum2	 = 0x00000200	# comment 2\n" +
	"passthru: /* passthru comment */\n" + 
	"enum3	 = 6	# comment 3\n" +
	"cat_2 enum:\n" +
	"# comment\n" +
	"enum1	 = 0x00000600	# comment 1\n" +
	"enum2		  = 0x00000800\n" +
	"passthru: /* passthru comment */\n" + 
	"enum3	 	= 2	# comment 2\n" +
	"cat_3 enum:\n" +
	"# comment\n" +
	"use cat_2	    enum2 # comment 1\n"

func checkEnum(cat, en, value string, ecats EnumCategories, t *testing.T) {
	if enums, ok := ecats[cat]; ok {
		if e, ok := enums[en]; ok {
			if e.Value == value {
				t.Logf("Enum found: %v::%v = %v", cat, en, value)
				return
			}
			t.Errorf("Enum not found: %v::%v", cat, en)
			return
		}
		t.Errorf("Enum not found: %v::%v", cat, en)
		return
	}
	t.Errorf("Category not found: %v", cat)
}

func TestReadEnums(t *testing.T) {
	r := strings.NewReader(testStr)
	e, err := ReadEnums(r)
	if err != nil {
		t.Errorf("Read enums failed: %v", err)
	}
	t.Logf("%v", e)

	if len(e) != 3 {
		t.Errorf("wrong length")
	}

	checkEnum("cat_1", "enum1", "0x00000100", e, t)
	checkEnum("cat_1", "enum2", "0x00000200", e, t)
	checkEnum("cat_1", "enum3", "6", e, t)
	checkEnum("cat_2", "enum1", "0x00000600", e, t)
	checkEnum("cat_2", "enum2", "0x00000800", e, t)
	checkEnum("cat_2", "enum3", "2", e, t)
	checkEnum("cat_3", "enum2", "0x00000800", e, t)		
}

