package prism_test

import (
	"./"
	"encoding/xml"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	color := prism.RGB{R: 1, G: 2, B: 3}

	r, g, b, a := color.RGBA()

	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
}

func Test_xml(t *testing.T) {
	color := prism.RGB{}
	bytes := []byte("<rgb red=\"1\" green=\"2\" blue=\"3\"/>")

	err := xml.Unmarshal(bytes, &color)

	if err != nil {
		t.Error(err)
	}

	r, g, b, a := color.RGBA()

	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
}
