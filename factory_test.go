package prism_test

import (
	"./"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	color := prism.RGB{R: 1, G: 2, B: 3}

	r, g, b, a := color.RGBA()

	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
}
