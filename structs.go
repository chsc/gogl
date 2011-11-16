package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Version

type Version struct {
	Major int
	Minor int
}

func MakeVersionFromString(version string) (Version, os.Error) {
	split := strings.Split(version, ".")
	if len(split) != 2 {
		return Version{0, 0}, os.NewError("Invalid version string.")
	}
	return MakeVersionFromMinorMajorString(split[0], split[1])
}

func MakeVersionFromMinorMajorString(minor, major string) (Version, os.Error) {
	majorNumber, err := strconv.Atoi(minor)
	if err != nil {
		return Version{0, 0}, os.NewError("Invalid major version number.")
	}
	minorNumber, err := strconv.Atoi(major)
	if err != nil {
		return Version{0, 0}, os.NewError("Invalid minor version number.")
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
