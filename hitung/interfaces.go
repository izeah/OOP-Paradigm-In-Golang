package hitung

// reusable struct
type entity struct {
	Nama string
}

// abstraction
// providing essenstial information to the outside world
// and hiding their background details

// apapun Type-nya, jika sudah menggunakan kontrak/interface ini
// harus memiliki method yang diminta dalam isi kontrak/interface-nya
type HitungBangunDatar interface {
	Luas() float64
	Keliling() float64
}

type HitungBangunRuang interface {
	Volume() float64
}

// inheritance of abstraction
type HitungMatematika interface {
	HitungBangunDatar
	HitungBangunRuang
}
