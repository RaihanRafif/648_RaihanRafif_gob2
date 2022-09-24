package helpers

import "fmt"

func SetBiodata(id, nama, alamat, pekerjaan, alasan string) {
	var bio Biodata
	bio.ID = id
	bio.Nama = nama
	bio.Alamat = alamat
	bio.Pekerjaan = pekerjaan
	bio.Alasan = alasan

	fmt.Printf("ID : %s,\nNama : %s,\nAlamat : %s,\nPekerjaan : %s,\nAlasan : %s", bio.ID, bio.Nama, bio.Alamat, bio.Pekerjaan, bio.Alasan)
}
