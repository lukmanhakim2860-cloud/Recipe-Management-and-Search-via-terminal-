package main

import "fmt"

const NMAX = 1000

type resep struct {
	Judul, BahanUtama, KategoriBahan, Komposisi, Langkah string
	Durasi, JmlPencarian                                 int
}

type Tabresep [NMAX]resep

func datacontoh(T *Tabresep, n *int) {
	// I.S masukan Array T dan n sembarang
	// F.S Array T berisi 5 data resep dan n bernilai 5
	T[0] = resep{"Nasi_Goreng", "Nasi", "Karbohidrat", "Nasi,Telur", "Goreng", 15, 0}
	T[1] = resep{"Ayam_Bakar", "Ayam", "Protein", "Ayam,Madu", "Bakar", 45, 0}
	T[2] = resep{"Salad_Buah", "Apel", "Buah", "Apel,Anggur", "Campur", 10, 0}
	T[3] = resep{"Opor_Ayam", "Ayam", "Protein", "Ayam,Santan", "Rebus", 60, 0}
	T[4] = resep{"Tumis_Kangkung", "Kangkung", "Sayur", "Kangkung,Bawang", "Tumis", 5, 0}
	*n = 5
}

func tambahresep(T *Tabresep, n *int) {
	// I.S Array T dan n terdefinisi, Kondisi n bisa penuh atau belum
	// F.S jika array belum penuh resep baru ditambah di indeks ke-n dari input lalu n tambah 1

	if *n < NMAX {
		fmt.Println("\n--- Tambah resep  ---")
		fmt.Print("Judul resep      : ")
		fmt.Scan(&T[*n].Judul)
		fmt.Print("Bahan Utama      : ")
		fmt.Scan(&T[*n].BahanUtama)
		fmt.Print("Kategori Bahan   : ")
		fmt.Scan(&T[*n].KategoriBahan)
		fmt.Print("Komposisi        : ")
		fmt.Scan(&T[*n].Komposisi)
		fmt.Print("Langkah-langkah  : ")
		fmt.Scan(&T[*n].Langkah)
		fmt.Print("Durasi (menit)   : ")
		fmt.Scan(&T[*n].Durasi)

		T[*n].JmlPencarian = 0
		*n = *n + 1
		fmt.Println("Data berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas penuh!")
	}
}

func ubahresep(T *Tabresep, n int) {
	// I.S array T dan n terdefinisi, meminta no indeks yang akan diubah
	// F.S mengubah resep yang dipilih

	var idx, d int
	var j, b string

	tampilSemua(*T, n)
	if n > 0 {
		fmt.Print("Masukkan Nomor resep yang diubah: ")
		fmt.Scan(&idx)
		idx = idx - 1

		if idx >= 0 && idx < n {
			fmt.Println("\n--- Ketik '-' jika tidak ingin mengubah ---")

			fmt.Print("Judul Baru: ")
			fmt.Scan(&j)
			if j != "-" {
				T[idx].Judul = j
			}

			fmt.Print("Bahan Utama Baru: ")
			fmt.Scan(&b)
			if b != "-" {
				T[idx].BahanUtama = b
			}

			fmt.Print("Durasi Baru (Ketik 0 jika tetap): ")
			fmt.Scan(&d)
			if d != 0 {
				T[idx].Durasi = d
			}

			fmt.Println("Data berhasil diubah.")
		} else {
			fmt.Println("Nomor tidak ditemukan.")
		}
	}
}

func hapusresep(T *Tabresep, n *int) {
	// I.S Array T dan n terdefinisi
	// F.S menghapus resep no "idx" dengan cara memindahkan atau menimpa dengan indeks dikanannya

	var idx, i int
	
	tampilSemua(*T, *n)
	if *n > 0 {
		fmt.Print("Masukkan Nomor resep yang dihapus: ")
		fmt.Scan(&idx)
		idx = idx - 1

		if idx >= 0 && idx < *n {
			i = idx
			for i < *n-1 {
				T[i] = T[i+1]
				i = i + 1
			}
			*n = *n - 1
			fmt.Println("Data berhasil dihapus.")
		} else {
			fmt.Println("Nomor tidak ditemukan.")
		}
	}
}

func tampilSemua(T Tabresep, n int) {
	// I.S Array T dan n terdefinisi
	// F.S menampilkan Seluruh data dari resep di array T

	var i int

	if n == 0 {
		fmt.Println("Belum ada data.")
	} else {
		fmt.Println("\n--- Daftar resep ---")
		i = 0
		for i < n {
			fmt.Print(i + 1)
			fmt.Print(". ")
			fmt.Print(T[i].Judul)
			fmt.Print(" > Bahan: ")
			fmt.Print(T[i].BahanUtama)
			fmt.Print(" > ")
			fmt.Print(T[i].Durasi)
			fmt.Println(" menit")
			i = i + 1
		}
	}
}

func tampilDetail(r resep) {
	// I.S tipe bentukan resep terdefinisi
	// F.S semua detail data dari r ditampilkan
	fmt.Println("- Judul      :", r.Judul)
	fmt.Println("  Kategori   :", r.KategoriBahan)
	fmt.Println("  Durasi     :", r.Durasi, "menit")
	fmt.Println("  Komposisi  :", r.Komposisi)
	fmt.Println("  Langkah    :", r.Langkah)
	fmt.Println()
}

func cariSequential(T *Tabresep, n int, bahan string) {
	// I.S array T, n, dan string bahan terdefinisi
	// F.S menelusuri array dari indeks 0 hingga n-1. jika BahanUtama cocok dengan input maka JmlPencarian bertambah 1 dan detail resep ditampilkan

	var ditemukan bool
	var i int
	
	fmt.Println("\n--- Hasil Sequential Search ---")

	ditemukan = false
	i = 0
	for i < n {
		if T[i].BahanUtama == bahan {
			T[i].JmlPencarian = T[i].JmlPencarian + 1
			tampilDetail(T[i])
			ditemukan = true
		}
		i = i + 1
	}

	if !ditemukan {
		fmt.Println("resep tidak ditemukan")
	}
}

func cariBinary(T *Tabresep, n int, bahan string) {
	// I.S array T, n, dan string bahan terdefinisi
	// F.S array T diurutkan terlebih dahulu berdasarkan BahanUtama. Kemudian cari resep, Jika ditemukan JmlPencarian bertambah 1 dan detail resep ditampilkan

	var i, kiri, tengah, kanan, idxTengah, start, end int
	var ketemu, batasKiri, batasKanan bool

	urutInsertionBahan(T, n)
	kiri = 0
	kanan = n - 1
	idxTengah = -1
	ketemu = false

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2
		if T[tengah].BahanUtama == bahan {
			idxTengah = tengah
			ketemu = true
		} else if T[tengah].BahanUtama < bahan {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	fmt.Println("\n--- Hasil Binary Search ---")
	if ketemu {
		start = idxTengah
		batasKiri = false
		for start > 0 && !batasKiri {
			if T[start-1].BahanUtama == bahan {
				start = start - 1
			} else {
				batasKiri = true
			}
		}
		end = idxTengah
		batasKanan = false
		for end < n-1 && !batasKanan {
			if T[end+1].BahanUtama == bahan {
				end = end + 1
			} else {
				batasKanan = true
			}
		}

		i = start
		for i <= end {
			T[i].JmlPencarian = T[i].JmlPencarian + 1
			tampilDetail(T[i])
			i = i + 1
		}
	} else {
		fmt.Println("resep tidak ditemukan.")
	}
}

func urutInsertionBahan(T *Tabresep, n int) {
	// I.S Array T dan n terdefinisi, T  belum terurut
	// F.S.: Array T terurut secara Ascending sesuai abjad

	var i, j int
	var key resep
	var geser bool

	i = 1
	for i < n {
		key = T[i]
		j = i - 1
		geser = true

		for j >= 0 && geser {
			if T[j].BahanUtama > key.BahanUtama {
				T[j+1] = T[j]
				j = j - 1
			} else {
				geser = false
			}
		}
		T[j+1] = key
		i = i + 1
	}
}

func urutSelectionDurasi(T *Tabresep, n int) {
	// I.S array T dan n terdefinisi, T belum terurut
	// F.S array T terurut secara Ascending tergantung Durasi Masak

	var i, j, idxMin int
	var temp resep

	i = 0
	for i < n-1 {
		idxMin = i
		j = i + 1
		for j < n {
			if T[j].Durasi < T[idxMin].Durasi {
				idxMin = j
			}
			j = j + 1
		}
		temp = T[i]
		T[i] = T[idxMin]
		T[idxMin] = temp
		i = i + 1
	}
	fmt.Println("\nData diurutkan berdasarkan Durasi (Selection Sort).")
	tampilSemua(*T, n)
}

func urutInsertionJudul(T *Tabresep, n int) {
	// I.S array T dan n terdefinisi, T belum terurut
	// F.S Array T terurut secara Ascending berdasarkan abjad pada judul resep

	var i, j int
	var key resep
	var geser bool

	i = 1
	for i < n {
		key = T[i]
		j = i - 1
		geser = true

		for j >= 0 && geser {
			if T[j].Judul > key.Judul {
				T[j+1] = T[j]
				j = j - 1
			} else {
				geser = false
			}
		}
		T[j+1] = key
		i = i + 1
	}
	fmt.Println("\nData diurutkan berdasarkan Abjad Judul (Insertion Sort).")
	tampilSemua(*T, n)
}

func tampilStatistik(T Tabresep, n int) {
	// I.S array T dan n terdefinisi
	// F.S menghitung dan menampilkan total resep per Kategori bahan, lalu menampilkan daftar 5 resep  yang nilai jmlpencarianya terbanyak

	type KatCount struct {
		Nama string
		Jml  int
	}
	var arrKat [NMAX]KatCount
	var nKat, i, j, k, idxMax, limit, m int
	var found, ada bool
	var TCopy Tabresep
	var temp resep

	nKat = 0
	i = 0

	for i < n {
		j = 0
		found = false
		for j < nKat && !found {
			if arrKat[j].Nama == T[i].KategoriBahan {
				arrKat[j].Jml = arrKat[j].Jml + 1
				found = true
			}
			j = j + 1
		}
		if !found {
			arrKat[nKat].Nama = T[i].KategoriBahan
			arrKat[nKat].Jml = 1
			nKat = nKat + 1
		}
		i = i + 1
	}

	fmt.Println("\n--- Statistik Kategori ---")
	k = 0
	for k < nKat {
		fmt.Print("- ", arrKat[k].Nama, ": ", arrKat[k].Jml, " resep\n")
		k = k + 1
	}

	fmt.Println("\n--- Top 5 Menu Paling Dicari ---")
	i = 0
	for i < n {
		TCopy[i] = T[i]
		i = i + 1
	}

	i = 0
	for i < n-1 {
		idxMax = i
		j = i + 1
		for j < n {
			if TCopy[j].JmlPencarian > TCopy[idxMax].JmlPencarian {
				idxMax = j
			}
			j = j + 1
		}
		temp = TCopy[i]
		TCopy[i] = TCopy[idxMax]
		TCopy[idxMax] = temp
		i = i + 1
	}

	limit = 5
	if n < 5 {
		limit = n
	}

	ada = false
	m = 0
	for m < limit {
		if TCopy[m].JmlPencarian > 0 {
			fmt.Print(m+1, ". ", TCopy[m].Judul, " (Dicari ", TCopy[m].JmlPencarian, " kali)\n")
			ada = true
		}
		m = m + 1
	}
	if !ada {
		fmt.Println("Belum ada resep yang dicari.")
	}
}

func main() {
	var resepcnth Tabresep
	var totalresep int
	var berhenti bool = false
	var pilihan int
	var bahan string

	datacontoh(&resepcnth, &totalresep)

	for !berhenti {
		fmt.Println("\n--------------------------------")
		fmt.Println("        Aplikasi resepKu                 ")
		fmt.Println("--------------------------------")
		fmt.Println("1. Tambah resep")
		fmt.Println("2. Ubah resep")
		fmt.Println("3. Hapus resep")
		fmt.Println("4. Tampil Semua resep")
		fmt.Println("5. Cari resep (Sequential Search)")
		fmt.Println("6. Cari resep (Binary Search)")
		fmt.Println("7. Urutkan Durasi Masak (Selection Sort)")
		fmt.Println("8. Urutkan Judul Abjad (Insertion Sort)")
		fmt.Println("9. Tampilkan Statistik")
		fmt.Println("0. Keluar")
		fmt.Println("   (Mohon gunakan huruf kapital)")
		fmt.Println("--------------------------------")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahresep(&resepcnth, &totalresep)
		} else if pilihan == 2 {
			ubahresep(&resepcnth, totalresep)
		} else if pilihan == 3 {
			hapusresep(&resepcnth, &totalresep)
		} else if pilihan == 4 {
			tampilSemua(resepcnth, totalresep)
		} else if pilihan == 5 {
			fmt.Print("Masukkan Bahan Utama (Tanpa Spasi): ")
			fmt.Scan(&bahan)
			cariSequential(&resepcnth, totalresep, bahan)
		} else if pilihan == 6 {
			fmt.Print("Masukkan Bahan Utama (Tanpa Spasi): ")
			fmt.Scan(&bahan)
			cariBinary(&resepcnth, totalresep, bahan)
		} else if pilihan == 7 {
			urutSelectionDurasi(&resepcnth, totalresep)
		} else if pilihan == 8 {
			urutInsertionJudul(&resepcnth, totalresep)
		} else if pilihan == 9 {
			tampilStatistik(resepcnth, totalresep)
		} else if pilihan == 0 {
			fmt.Println("Terima kasih!")
			berhenti = true
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
