package main

// Version

type Version struct {
	Major int
	Minor int
}

func (v *Version) Compare(v2 *Version) int {
	if v.Major < v2.Major {
		return -1
	} else if v.Major > v2.Major {
		return 1
	} else if v.Minor < v2.Minor {
		return -1
	} else if v.Minor > v2.Minor {
		return 1
	}
	return 0
}

func (v *Version) Valid() bool {
	return v.Major == 0 && v.Minor == 0
}

// Functions

type Parameter struct {
	Name    string
	Type    string
	InArray bool
}

type Function struct {
	Name              string
	Parameters        []Parameter
	Return            string
	Version           Version
	DeprecatedVersion Version
	Category          string
}

type Functions map[string][]*Function

// Enums

type Enum struct {
	Name  string
	Value string
}

type EnumCategories map[string]Enums
type Enums map[string]Enum

// Type maps

type TypeMap map[string]string
