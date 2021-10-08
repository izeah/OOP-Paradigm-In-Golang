package hitung

import "math"

// OOP Class
type kubus struct {
	// inheritance
	entity
	// encapsulation
	sisi float64
}

// creates new Instance of Kubus class (object)
func NewKubus(sisi float64) kubus {
	return kubus{
		entity: entity{
			Nama: "Kubus (Cube)",
		},
		sisi: sisi,
	}
}

// Encapsulation
// Encapsulation allows for us to create methods that can only be accessed or changed by public methods.
func (k kubus) GetSisi() float64 {
	return k.sisi
}

func (k kubus) Luas() float64 {
	// rumus = 6 x s²
	return 6 * math.Pow(k.sisi, 2)
}

func (k kubus) Keliling() float64 {
	// rumus = 12 x panjang sisi
	return 12 * k.sisi
}

func (k kubus) Volume() float64 {
	// rumus = s³
	return math.Pow(k.sisi, 3)
}
