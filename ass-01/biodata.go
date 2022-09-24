package main

import (
	"ass-01/helpers"
	"fmt"
	"strings"
)

func main() {
	// SetBiodata(id, nama, alamat, pekerjaan, alasan)
	var perintah string

	fmt.Print(strings.Repeat("#", 20))
	fmt.Print("\n")

	// baca input dari terminal
	fmt.Print("Masukkan ID peserta *antara 1-3 : ")
	fmt.Scanln(&perintah)
	fmt.Print("\n")

	if perintah == "1" {
		helpers.SetBiodata("1", "Raihan", "jl. raya", "mahasiswa", "untuk menambah ilmu")
	} else if perintah == "2" {
		helpers.SetBiodata("2", "siput", "jl. kombo", "mahasiswa", "gabut aja")
		fmt.Print("\n")
	} else if perintah == "3" {
		helpers.SetBiodata("3", "keong", "jl. wumbo", "mahasiswa", "gk ada alasan")
		fmt.Print("\n")
	}
}
