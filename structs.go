package main

import (
	"errors"
	"fmt"

	"strconv"
	"strings"
)

// Version

type Version struct {
	Major int
	Minor int
}

func MakeVersionFromString(version string) (Version, error) {
	split := strings.Split(version, ".")
	if len(split) != 2 {
		return Version{0, 0}, errors.New("Invalid version string.")
	}
	return MakeVersionFromMinorMajorString(split[0], split[1])
}

func MakeVersionFromMinorMajorString(minor, major string) (Version, error) {
	majorNumber, err := strconv.Atoi(minor)
	if err != nil {
		return Version{0, 0}, errors.New("Invalid major version number.")
	}
	minorNumber, err := strconv.Atoi(major)
	if err != nil {
		return Version{0, 0}, errors.New("Invalid minor version number.")
	}
	return Version{majorNumber, minorNumber}, nil
}

func (v Version) Compare(v2 Version) int {
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

func (v Version) Valid() bool {
	return v.Major != 0 || v.Minor != 0
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
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
	Offset            int
	GlxRopCode        int
}

type Functions []*Function

func (fs Functions) Find(name string) *Function {
	for _, f := range fs {
		if f.Name == name {
			return f
		}
	}
	return nil
}

type FunctionCategories map[string]Functions

// Enums

type Enum struct {
	Name  string
	Value string
}

type Enums map[string]Enum
type EnumCategories map[string]Enums

// Type maps

type TypeMap map[string]string

// Packages

type Package struct {
	Enums     EnumCategories
	Functions FunctionCategories
}

type Packages map[string]*Package
