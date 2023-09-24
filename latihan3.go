package main

import (
	"fmt"
)

// jadwal=variable	struct=tipe data
type jadwal struct {
	hari   string //hari, matkul, kelas = isinya
	matkul string
	kelas  string
}
type dosen struct {
	nama   string
	matkul string
}

func main() {
	fmt.Println("halo")

	dataku := map[string]string{
		"nama":          "Regita",
		"jenis kelamin": "Perempuan",
		"umur":          "17",
		"kelas":         "A",
	}
	fmt.Println(dataku)
	datamu := map[string]string{
		"nama":  "Maulia",
		"kelas": "B",
	}
	fmt.Println(len(datamu))

	//jika ingin hapus data  di map => delete(map, key)
	//map[key]=value
	dataku["umur"] = "19" //ganti elemen / ubah data di map dg key
	fmt.Println("Namaku", dataku["nama"])
	fmt.Println("Umurku", dataku["umur"])

	kelasa := jadwal{
		hari:   "Senin",
		matkul: "SO",
	}
	kelasb := jadwal{
		hari:   "Jumat",
		matkul: "Jarkom",
	}

	dosen1 := dosen{
		nama:   "Nur Chasanah, S.Kom.,M.Kom.",
		matkul: "SO",
	}
	dosen2 := dosen{
		nama:   "Dadang Iskandar, S.T.,M.Eng.",
		matkul: "Jarkom",
	}

	fmt.Println("Hari", kelasa.hari, ", ada matkul", kelasa.matkul)
	fmt.Println("Hari", kelasb.hari, datamu["nama"], "ada matkul", kelasb.matkul)

	//pointer
	var hari2 *string = &kelasa.hari
	*hari2 = "Selasa"

	fmt.Println("kuliah pengganti", kelasa.matkul, "di hari", kelasa.hari)
	fmt.Println("Dosen", dosen1.matkul, "adalah", dosen1.nama)
	fmt.Println("Dosen", dosen2.matkul, "adalah", dosen2.nama)

}
