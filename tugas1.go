package main

import "fmt"

const phi = 3.14

func Luaspersegi(sisi float32) float32 {
	return sisi * sisi
}

func Luassegitiga(alas, tinggi float32) float32 {
	return float32(alas*tinggi) / 2
}

func Luaslingkaran(jari float32) float32 {
	return phi * float32(jari*jari)
}

func Energipotensial(massa, ketinggian float32) float32 {
	gravitasi := 10.0
	return massa * float32(gravitasi) * ketinggian
}

func Energikinetik(massa, kecepatan float32) float32 {
	return 0.5 * massa * (kecepatan * kecepatan)
}

func frekuensi(n int, t float32) float32 {
	return float32(n) / t
}

func main() {
	var n int
	var sisi, alas, tinggi, jari, massa, kecepatan, ketinggian, t, C, R, F, K float32
	var pilihan string

	fmt.Println("Menghitung Luas Persegi")
	fmt.Print("sisi: ")
	fmt.Scan(&sisi)
	luaspersegi := Luaspersegi(sisi)
	fmt.Println("Luas = ", luaspersegi)

	fmt.Println("Menghitung Luas Segitiga")
	fmt.Print("alas: ")
	fmt.Scan(&alas)
	fmt.Print("tinggi ")
	fmt.Scan(&tinggi)
	luassegitiga := Luassegitiga(alas, tinggi)
	fmt.Println("Luas = ", luassegitiga)

	fmt.Println("Menghitung Luas Lingkaran")
	fmt.Print("jari-jari: ")
	fmt.Scan(&jari)
	luaslingkaran := Luaslingkaran(jari)
	fmt.Println("Luas = ", luaslingkaran)

	fmt.Println("Menghitung Energi Potensial")
	fmt.Print("massa(kg): ")
	fmt.Scan(&massa)
	fmt.Print("ketinggian(m): ")
	fmt.Scan(&ketinggian)
	energipotensial := Energipotensial(massa, ketinggian)
	fmt.Println("Ep = ", energipotensial, " Joule")

	fmt.Println("Menghitung Energi Kinetik")
	fmt.Print("massa: ")
	fmt.Scan(&massa)
	fmt.Print("kecepatan: ")
	fmt.Scan(&kecepatan)
	energiKinetik := Energikinetik(massa, kecepatan)
	fmt.Println("Ek = ", energiKinetik)

	fmt.Println("Menghitung Frekuensi")
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Print("t: ")
	fmt.Scan(&t)
	frek := frekuensi(n, t)
	fmt.Println("= ", frek)

	fmt.Println("Konversi Suhu")
	fmt.Println("1. C ke (R-F-K)")
	fmt.Println("2. R ke (C-F-K)")
	fmt.Println("3. F ke (C-R-K)")
	fmt.Println("4. K ke (C-R-F)")
	fmt.Print("Masukkan pilihan Anda (misal, 1): ")
	fmt.Scan(&pilihan)

	if pilihan == "1" {
		fmt.Print("C = ")
		fmt.Scan(&C)
		R := 4.0 / 5.0 * C
		F := (9.0 / 5.0 * C) + 32
		K := C + 273
		fmt.Println(C, "°C =", R, "°R")
		fmt.Println(C, "°C =", F, "°F")
		fmt.Println(C, "°C =", K, "°K")
	} else if pilihan == "2" {
		fmt.Print("R = ")
		fmt.Scan(&R)
		C := 5.0 / 4.0 * R
		F := (9.0 / 4.0 * R) + 32
		K := (5.0 / 4.0 * R) + 273
		fmt.Println(R, "°R =", C, "°C")
		fmt.Println(R, "°R =", F, "°F")
		fmt.Println(R, "°R =", K, "°K")
	} else if pilihan == "3" {
		fmt.Print("F = ")
		fmt.Scan(&F)
		C := (5.0 / 9.0) * (F - 32)
		R := (4.0 / 9.0) * (F - 32)
		K := (5.0/9.0)*(F-32) + 273
		fmt.Println(F, "°F =", C, "°C")
		fmt.Println(F, "°F =", R, "°R")
		fmt.Println(F, "°F =", K, "°K")
	} else if pilihan == "4" {
		fmt.Print("K = ")
		fmt.Scan(&K)
		C := K - 273
		R := (4.0 / 5.0) * (K - 273)
		F := (9.0/5.0)*(K-273) + 32
		fmt.Println(R, "°R =", C, "°C")
		fmt.Println(R, "°R =", F, "°F")
		fmt.Println(R, "°R =", K, "°K")
	} else {
		fmt.Println("Pilihan tidak tersedia")
	}
	return
}
