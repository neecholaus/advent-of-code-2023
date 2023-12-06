package main

import "strconv"

type Map struct {
	RawName     string
	Source      string
	Destination string
	MapIndexes  []*MapIndex
}

type MapIndex struct {
	DestinationStart int
	SourceStart      int
	Range            int
}

func (mi *MapIndex) intOrPanic(intAsString string) int {
	num, err := strconv.Atoi(intAsString)
	if err != nil {
		panic("cannot make string into int")
	}
	return num
}
