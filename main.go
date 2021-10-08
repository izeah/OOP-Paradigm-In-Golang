package main

import (
	"fmt"
	"oop-paradigm-in-golang/hitung"
	"oop-paradigm-in-golang/person"
)

/**
 * PSEUDO-CODE OF OOP AND PROCEDURE PROGRAMMING PARADIGM
 *
 * *   $query = "SELECT first_name, last_name FROM users LIMIT 10";
 *
 *   ! OOP Style
 * *   $mysqli = new mysqli($host, $user, $pass, $db);
 * *   $result = $mysqli->query($query);
 *
 *   ! Procedural Style
 * *   $link = mysqli_connect($host, $user, $pass, $db);
 * *   $result = mysqli_query($link, $query);
 */

func main() {
	// polymorphism #1
	var hitungBangunDatar hitung.HitungBangunDatar
	var hitungMatematika hitung.HitungMatematika

	// Menghitung Lingkaran
	// type HitungBangunDatar interface {
	// 	Luas() float64
	// 	Keliling() float64
	// }
	lingkaran := hitung.NewLingkaran(10)
	hitungBangunDatar = lingkaran
	fmt.Println(lingkaran.Nama)
	fmt.Println("Jari-jari \t\t:", lingkaran.GetRadius())               // Jari-jari : 5
	fmt.Println("Luas Lingkaran \t\t:", hitungBangunDatar.Luas())       // Luas Lingkaran : 78.53981633974483
	fmt.Println("Keliling Lingkaran \t:", hitungBangunDatar.Keliling()) // Keliling Lingkaran : 31.41592653589793

	fmt.Printf("\n\n")

	// Menghitung Kubus
	// type HitungMatematika interface {
	// 	Luas() float64
	// 	Keliling() float64
	//  Volume() float64
	// }
	kubus := hitung.NewKubus(7)
	hitungMatematika = kubus
	fmt.Println(kubus.Nama)
	fmt.Println("Sisi \t\t:", kubus.GetSisi())                     // Sisi : 7
	fmt.Println("Luas Kubus \t:", hitungMatematika.Luas())         // Luas Kubus : 294
	fmt.Println("Keliling Kubus \t:", hitungMatematika.Keliling()) // Keliling Kubus : 84
	fmt.Println("Volume Kubus \t:", hitungMatematika.Volume())     // Volume Kubus : 343

	// pake polymorphism #2
	fmt.Println("Volume Kubus \t:", SeberapaVolume(kubus)) // Volume Kubus : 343

	// pake polymorphism #3
	kubusInterface := NewKubusByInterface(4)
	fmt.Println("Volume Kubus \t:", kubusInterface.Volume()) // Volume Kubus : 343

	fmt.Printf("\n\n")
	person := person.NewPerson("Faiz", "Mohammad", "Hafidza", 22)
	fmt.Println("My name is \t\t:", person.GetFullName()) // Faiz Mohammad Hafidza
	fmt.Println("Usia \t\t\t:", person.GetAge())          // 22
	person.UlangTahun()
	fmt.Println("Setelah Ulang Tahun \t:", person.GetAge()) // expected = 23, actual = 22
	person.UpdateUsia(15)
	fmt.Println("Update Usia \t\t:", person.GetAge()) // expected = 15, actual = 15
}

// polymorphism #2
// type HitungBangunRuang interface {
//     Volume() float64
// }
func SeberapaVolume(bangunRuang hitung.HitungBangunRuang) float64 {
	return bangunRuang.Volume()
}

// polymorphism #3
func NewKubusByInterface(sisi float64) hitung.HitungBangunRuang {
	return hitung.NewKubus(sisi)
}
