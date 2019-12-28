package main

import (
	"fmt"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type (
	ComplexStruct struct {
		A string `cty:"A"`
		//Note: There's an easy bug here: any typos in the Tag are not going to be caught
		Inner SimpleStruct `cty:"Inner"`
	}
)

type (
	SimpleStruct struct {
		X int `cty:"X"`
		Y int `cty:"Y"`
		S string `cty:"S"`
	}
)

func main() {

	startingStruct := ComplexStruct{
		A:     "This is value A",
		Inner: SimpleStruct{
			X: 1,
			Y: 2,
			S: "This is value S",
		},
	}

	ctyValue, err := gocty.ToCtyValue(startingStruct, cty.Object(
		map[string]cty.Type{
			"A": cty.String,
			"Inner": cty.Object(map[string]cty.Type{
				"X": cty.Number,
				"Y": cty.Number,
				"S": cty.String,
			}),
	}))
	if err != nil {
		fmt.Println("Boo!")
	}

	endingStruct := new(ComplexStruct)
	err = gocty.FromCtyValue(ctyValue, &endingStruct)

	if startingStruct == *endingStruct {
		fmt.Println("Yay!")
	}
}
