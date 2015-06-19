package hex_test

import (
	"."
	"../../clouseau/reckon"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	factory := hex.NewFactory()

	color := factory.GetColor("#0")

	reckon.That(color).Is.Nil()

	color = factory.GetColor("#010203")

	r, g, b, a := color.RGBA()

	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
	reckon.That(r).Is.EqualTo(uint32(1))
	reckon.That(g).Is.EqualTo(uint32(2))
	reckon.That(b).Is.EqualTo(uint32(3))
	reckon.That(a).Is.EqualTo(uint32(255))

	color = factory.GetColor("010203")

	r, g, b, a = color.RGBA()

	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
	reckon.That(r).Is.EqualTo(uint32(1))
	reckon.That(g).Is.EqualTo(uint32(2))
	reckon.That(b).Is.EqualTo(uint32(3))
	reckon.That(a).Is.EqualTo(uint32(255))

	color = factory.GetColor("102030")

	r, g, b, a = color.RGBA()

	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
	reckon.That(r).Is.EqualTo(uint32(16))
	reckon.That(g).Is.EqualTo(uint32(32))
	reckon.That(b).Is.EqualTo(uint32(48))
	reckon.That(a).Is.EqualTo(uint32(255))

	color = factory.GetColor("ffffff")

	r, g, b, a = color.RGBA()

	fmt.Printf("red: %v, green: %v, blue: %v, alpha %v\n", r, g, b, a)
	reckon.That(r).Is.EqualTo(uint32(255))
	reckon.That(g).Is.EqualTo(uint32(255))
	reckon.That(b).Is.EqualTo(uint32(255))
	reckon.That(a).Is.EqualTo(uint32(255))
}
