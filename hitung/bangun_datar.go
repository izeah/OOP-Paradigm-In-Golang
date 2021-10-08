package hitung

import "math"

// OOP Class
type lingkaran struct {
	// inheritance
	entity

	// encapsulation
	diameter float64
}

// creates new Instance of Lingkaran class (object)
func NewLingkaran(diameter float64) lingkaran {
	return lingkaran{
		entity: entity{
			Nama: "Lingkaran (Circle)",
		},
		diameter: diameter,
	}
}

// Encapsulation
// Encapsulation allows for us to create methods that can only be accessed or changed by public methods.
func (l lingkaran) GetRadius() float64 {
	return l.diameter / 2
}

func (l lingkaran) Luas() float64 {
	// rumus = phi x r²
	return math.Pi * math.Pow(l.GetRadius(), 2)
}

func (l lingkaran) Keliling() float64 {
	// rumus = phi x d
	// rumus = 2 x phi x r
	return math.Pi * l.diameter
}
