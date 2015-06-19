package named_test

import (
	"."
	"../../clouseau/reckon"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	factory := named.NewFactory()

	color := factory.GetColor("Hello World")
	reckon.That(color).Is.Nil()

	color = factory.GetColor(named.Red)

	r, g, b, a := color.RGBA()
	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
	reckon.That(r).Is.EqualTo(uint32(255))
	reckon.That(g).Is.EqualTo(uint32(0))
	reckon.That(b).Is.EqualTo(uint32(0))
	reckon.That(a).Is.EqualTo(uint32(255))

	color = factory.GetColor(named.Green)

	r, g, b, a = color.RGBA()
	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
	reckon.That(r).Is.EqualTo(uint32(0))
	reckon.That(g).Is.EqualTo(uint32(128))
	reckon.That(b).Is.EqualTo(uint32(0))
	reckon.That(a).Is.EqualTo(uint32(255))

	color = factory.GetColor(named.Lime)

	r, g, b, a = color.RGBA()
	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
	reckon.That(r).Is.EqualTo(uint32(0))
	reckon.That(g).Is.EqualTo(uint32(255))
	reckon.That(b).Is.EqualTo(uint32(0))
	reckon.That(a).Is.EqualTo(uint32(255))

	color = factory.GetColor(named.Blue)

	r, g, b, a = color.RGBA()
	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
	reckon.That(r).Is.EqualTo(uint32(0))
	reckon.That(g).Is.EqualTo(uint32(0))
	reckon.That(b).Is.EqualTo(uint32(255))
	reckon.That(a).Is.EqualTo(uint32(255))

	color = factory.GetColor(named.White)

	r, g, b, a = color.RGBA()
	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
	reckon.That(r).Is.EqualTo(uint32(255))
	reckon.That(g).Is.EqualTo(uint32(255))
	reckon.That(b).Is.EqualTo(uint32(255))
	reckon.That(a).Is.EqualTo(uint32(255))
}
