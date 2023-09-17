package main

import "fmt"

func main() {
	var bil = []int{7, 4, 1, 9}

	n := len(bil)
	for i := 0; i < n-i; i++ {
		for j := 0; j < n-i-1; j++ {
			if bil[j] < bil[j+1] {
				bil[j], bil[j+1] = bil[j+1], bil[j]
			}
		}
	}
	fmt.Println("Bilangan terbesar-terkecil: ", bil)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Data awal: ", slice)
	sliceawal := slice[:5]

	tambahdata := []int{9, 8, 7}
	sliceawal = append(sliceawal, tambahdata...) //variadic parameters = parameter yg dapat menerima sejumlah argumen variabel dg tipe data sama
	fmt.Println("Data akhir: ", sliceawal)
}
