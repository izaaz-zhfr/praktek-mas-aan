package main

import "fmt"

type Segitiga struct {
	Alas   float64
	Tinggi float64
}


func hitungLuasSegitiga(s Segitiga) float64 {
	return 0.5 * s.Alas * s.Tinggi
}

func main() {
	kota := "jakarta"
	fmt.Println("Kota:", kota)

	var namaAplikasi string = "Kalkulator Segitiga"
	fmt.Println("Aplikasi:", namaAplikasi)

	dataSegitiga := Segitiga{Alas: 10, Tinggi: 6}


	luas := hitungLuasSegitiga(dataSegitiga)

	fmt.Println("Alas:", dataSegitiga.Alas, "dan Tinggi:", dataSegitiga.Tinggi)
	fmt.Println("Hasil Luas Segitiga:", luas)
}
