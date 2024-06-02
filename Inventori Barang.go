package main

import (
	"fmt"
)

const NMAX = 1024

type barang struct {
	nama       string
	jenis      string
	jumlahStok int
	kategori   string
}
type Transaksi struct {
	jam   int
	menit int
	detik int
}
type tabBarang [NMAX]barang
type tabMasuk [NMAX]Transaksi
type tabKeluar [NMAX]Transaksi

var a, d tabBarang
var b tabMasuk
var c tabKeluar
var n, i, arrayjumlah int

func main() {
	var menuInput int
	fmt.Println("===== Masukkan Angka Untuk Menu =====")
	fmt.Println("1. Data Barang dan Transaksi")
	fmt.Println("2. Pencarian Kata Kunci Tertentu.")
	fmt.Println("3. Pencarian Kategori Tertentu.")
	fmt.Println("4. Tampilan semua data")
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&menuInput)
	switch menuInput {
	case 1:
		fmt.Println("Pencatatan Data Barang dan Transaksi...")
		PencatatanBarangdanTransaksi()
	case 2:
		fmt.Println("Pencarian Kata Kunci Barang...")
		pemasukandata(a, &d, arrayjumlah)
		pencariankatakuncibarang()
	case 3:
		fmt.Println("Pencarian Kategori Tertentu...")
		pemasukandata(a, &d, arrayjumlah)
		pencarianbarangkategori()
	case 4:
		fmt.Println("Tampilan Semua Data Barang")
		tampilarray(a, b, c)
		main()
	}
}

func pemasukandata(a tabBarang, d *tabBarang, n int) {
	for i := 0; i < n; i++ {
		d[i] = a[i]
	}
}

func PencatatanBarangdanTransaksi() {
	var menuInput int
	fmt.Println("=====Apa yang anda pilih?=====")
	fmt.Println("1. Penambahan.")
	fmt.Println("2. Perubahan.")
	fmt.Println("3. Penghapusan.")
	fmt.Println("4. Kembali Ke Menu Utama.")
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&menuInput)
	switch menuInput {
	case 1:
		fmt.Println("Penambahan Data dan Transaksi...")
		pemasukanBarangdanTransaksi()
	case 2:
		fmt.Println("Perubahan Data dan Transaksi...")
		perubahanBarangdanTransaksi()
	case 3:
		fmt.Println("Penghapusan Data dan Transaksi...")
		penghapusanBarangdanTransaksi()
	case 4:
		main()
	}
}

func pemasukanBarangdanTransaksi() {
	fmt.Print("Masukkan Jumlah Data yang ingin di Masukkan: ")
	fmt.Scan(&n)
	isiArray(&a, &b, &c, &n)
	tampilarray(a, b, c)
	PencatatanBarangdanTransaksi()
}

func isiArray(a *tabBarang, b *tabMasuk, c *tabKeluar, n *int) {
	i = 0
	arrayjumlah += *n
	for i < arrayjumlah {
		fmt.Println("Masukkan Barang")
		fmt.Println("Jenis - Nama - Jumlah Stok - Kategori")
		fmt.Scan(&a[i].nama, &a[i].jenis, &a[i].jumlahStok, &a[i].kategori)
		i++
	}
	i = 0
	for i < arrayjumlah {
		fmt.Println("Masukkan Waktu Transkasi Barang Masuk")
		fmt.Println("Jam - Menit - Detik")
		fmt.Scan(&b[i].jam, &b[i].menit, &b[i].detik)
		fmt.Println("Masukkan Waktu Transkasi Barang Keluar")
		fmt.Println("Jam - Menit - Detik")
		fmt.Scan(&c[i].jam, &c[i].menit, &c[i].detik)
		i++
	}
}

func tampilarray(a tabBarang, b tabMasuk, c tabKeluar) {
	fmt.Println("Data Masuk:")
	for i = 0; i < arrayjumlah; i++ {
		fmt.Println(i, a[i].nama, a[i].jenis, a[i].jumlahStok, a[i].kategori)
	}
	fmt.Println("Pada Pukul")
	for i = 0; i < arrayjumlah; i++ {
		fmt.Println(i, b[i].jam, ":", b[i].menit, ":", b[i].detik)
	}
	fmt.Println("Keluar Pada Pukul")
	for i = 0; i < arrayjumlah; i++ {
		fmt.Println(i, c[i].jam, ":", c[i].menit, ":", c[i].detik)
	}
}

func perubahanBarangdanTransaksi() {
	var menuInput int
	fmt.Println("=====Apa Yang akan di Ubah?=====")
	fmt.Println("1. Data dan Transaksi")
	fmt.Println("2. Data")
	fmt.Println("3. Transaksi")
	fmt.Println("4.Kembali ke Menu Utama")
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&menuInput)
	switch menuInput {
	case 1:
		perubahantransaksidanbarang()
	case 2:
		perubahandata()
	case 3:
		perubahantransaki()
	case 4:
		PencatatanBarangdanTransaksi()
	}
}

func perubahantransaksidanbarang() {
	var menuInput int
	fmt.Println("Data Sebelum Ubah:")
	tampilarray(a, b, c)
	fmt.Println("Urutan Ke berapa yang ingin di ubah?")
	fmt.Print("Masukkan Angka Urutan: ")
	fmt.Scan(&menuInput)
	Perubahantrandanbar(&a, &b, &c, &menuInput)
	tampilarray(a, b, c)
	perubahanBarangdanTransaksi()
}

func Perubahantrandanbar(a *tabBarang, b *tabMasuk, c *tabKeluar, n *int) {
	fmt.Println("Masukkan Barang")
	fmt.Println("Jenis - Nama - Jumlah Stok - Kategori")
	fmt.Scan(&a[*n].nama, &a[*n].jenis, &a[*n].jumlahStok, &a[*n].kategori)
	fmt.Println("Masukkan Waktu Transkasi Barang Masuk")
	fmt.Println("Jam - Menit - Detik")
	fmt.Scan(&b[*n].jam, &b[*n].menit, &b[*n].detik)
	fmt.Println("Masukkan Waktu Transkasi Barang Keluar")
	fmt.Println("Jam - Menit - Detik")
	fmt.Scan(&c[*n].jam, &c[*n].menit, &c[*n].detik)
}

func perubahandata() {
	var menuInput int
	fmt.Println("Data Sebelum Ubah:")
	tampilarray(a, b, c)
	fmt.Println("Urutan Ke berapa yang ingin di ubah?")
	fmt.Print("Masukkan Angka Urutan: ")
	fmt.Scan(&menuInput)
	Perubahandat(&a, &menuInput)
	tampilarray(a, b, c)
	perubahanBarangdanTransaksi()
}

func Perubahandat(a *tabBarang, n *int) {
	fmt.Println("Masukkan Barang")
	fmt.Println("Jenis - Nama - Jumlah Stok - Kategori")
	fmt.Scan(&a[*n].nama, &a[*n].jenis, &a[*n].jumlahStok, &a[*n].kategori)
}

func perubahantransaki() {
	var menuInput, x int
	fmt.Println("Data Sebelum Ubah:")
	tampilarray(a, b, c)
	fmt.Println("Urutan Ke berapa yang ingin di ubah?")
	fmt.Print("Masukkan Angka Urutan: ")
	fmt.Scan(&menuInput)
	fmt.Println("Transaksi apakah yang ingin di hapus?")
	fmt.Println("1. Transaksi Masuk")
	fmt.Println("2. Transaksi Keluar")
	fmt.Println("3. Transaksi Masuk dan Keluar")
	fmt.Println("4. Keluar Menu Perubahan Transaksi")
	fmt.Scan(&x)
	switch x {
	case 1:
		Perubahantranmas(&b, &menuInput)
		tampilarray(a, b, c)
		perubahanBarangdanTransaksi()
	case 2:
		Perubahantrankel(&c, &menuInput)
		tampilarray(a, b, c)
		perubahanBarangdanTransaksi()
	case 3:
		Perubahantranmaskel(&b, &c, &menuInput)
		tampilarray(a, b, c)
		perubahanBarangdanTransaksi()
	default:
		perubahanBarangdanTransaksi()
	}
}

func Perubahantranmas(b *tabMasuk, n *int) {
	fmt.Println("Masukkan Waktu Transkasi Barang Masuk")
	fmt.Println("Jam - Menit - Detik")
	fmt.Scan(&b[*n].jam, &b[*n].menit, &b[*n].detik)
}

func Perubahantrankel(c *tabKeluar, n *int) {
	fmt.Println("Masukkan Waktu Transkasi Barang Keluar")
	fmt.Println("Jam - Menit - Detik")
	fmt.Scan(&c[*n].jam, &c[*n].menit, &c[*n].detik)
}

func Perubahantranmaskel(b *tabMasuk, c *tabKeluar, n *int) {
	fmt.Println("Masukkan Waktu Transkasi Barang Masuk")
	fmt.Println("Jam - Menit - Detik")
	fmt.Scan(&b[*n].jam, &b[*n].menit, &b[*n].detik)
	fmt.Println("Masukkan Waktu Transkasi Barang Keluar")
	fmt.Println("Jam - Menit - Detik")
	fmt.Scan(&c[*n].jam, &c[*n].menit, &c[*n].detik)
}

func penghapusanBarangdanTransaksi() {
	var menuInput int
	fmt.Println("=====Apa Yang akan di hapus?=====")
	fmt.Println("1. Data dan Transaksi")
	fmt.Println("2. Data")
	fmt.Println("3. Transaksi")
	fmt.Println("4. Kembali ke Menu Utama")
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&menuInput)
	switch menuInput {
	case 1:
		penghapusantransaksidanbarang()
	case 2:
		penghapusandata()
	case 3:
		penghapusantransaki()
	case 4:
		PencatatanBarangdanTransaksi()
	}
}

func penghapusantransaksidanbarang() {
	var menuInput int
	fmt.Println("Data Sebelum Hapus:")
	tampilarray(a, b, c)
	fmt.Println("Urutan Ke berapa yang ingin di hapus?")
	fmt.Print("Masukkan Angka Urutan: ")
	fmt.Scan(&menuInput)
	Penghapusantrandanbar(&a, &b, &c, &menuInput)
	tampilarray(a, b, c)
	penghapusanBarangdanTransaksi()
}

func Penghapusantrandanbar(a *tabBarang, b *tabMasuk, c *tabKeluar, n *int) {
	for x := *n; x < arrayjumlah; x++ {
		a[x] = a[x+1]
		b[x] = b[x+1]
		c[x] = c[x+1]
	}
	arrayjumlah--
}

func penghapusandata() {
	var menuInput int
	fmt.Println("Data Sebelum Hapus:")
	tampilarray(a, b, c)
	fmt.Println("Urutan Ke berapa yang ingin di hapus?")
	fmt.Print("Masukkan Angka Urutan: ")
	fmt.Scan(&menuInput)
	Penghapusandat(&a, &menuInput)
	tampilarray(a, b, c)
	penghapusanBarangdanTransaksi()
}

func Penghapusandat(a *tabBarang, n *int) {
	for x := *n; x < arrayjumlah; x++ {
		a[x] = a[x+1]
	}
	arrayjumlah--
}

func penghapusantransaki() {
	var menuInput, x int
	fmt.Println("Data Sebelum Hapus:")
	tampilarray(a, b, c)
	fmt.Println("Urutan Ke berapa yang ingin di hapus?")
	fmt.Print("Masukkan Angka Urutan: ")
	fmt.Scan(&menuInput)
	fmt.Println("Transaksi apakah yang ingin di hapus?")
	fmt.Println("1. Transaksi Masuk")
	fmt.Println("2. Transaksi Keluar")
	fmt.Println("3. Transaksi Masuk dan Keluar")
	fmt.Println("4. Keluar Menu Perubahan Transaksi")
	fmt.Scan(&x)
	switch x {
	case 1:
		Penghapusantranmas(&b, &menuInput)
		tampilarray(a, b, c)
		penghapusanBarangdanTransaksi()
	case 2:
		Penghapusantrankel(&c, &menuInput)
		tampilarray(a, b, c)
		penghapusanBarangdanTransaksi()
	case 3:
		Penghapusantranmaskel(&b, &c, &menuInput)
		tampilarray(a, b, c)
		penghapusanBarangdanTransaksi()
	default:
		perubahanBarangdanTransaksi()
	}
}

func Penghapusantranmas(b *tabMasuk, n *int) {
	for x := *n; x < arrayjumlah; x++ {
		b[x] = b[x+1]
	}
	arrayjumlah--
}

func Penghapusantrankel(c *tabKeluar, n *int) {
	for x := *n; x < arrayjumlah; x++ {
		c[x] = c[x+1]
	}
	arrayjumlah--
}

func Penghapusantranmaskel(b *tabMasuk, c *tabKeluar, n *int) {
	for x := *n; x < arrayjumlah; x++ {
		b[x] = b[x+1]
		c[x] = c[x+1]
	}
	arrayjumlah--
}

func pencariankatakuncibarang() {
	var menuInput int
	fmt.Println("Masukkan Kata kunci yang ingin dicari")
	fmt.Println("1. Nama")
	fmt.Println("2. Jenis")
	fmt.Println("3. Jumlah Stok")
	fmt.Println("4. Kembali ke Menu")
	fmt.Scan(&menuInput)
	switch menuInput {
	case 1:
		pencarianNama()
	case 2:
		pencarianJenis()
	case 3:
		pencarianStok()
	case 4:
		main()
	}
}
func pencarianNama() {
	var cari string
	fmt.Print("Masukkan kata kunci yang dicari: ")
	fmt.Scan(&cari)
	i = 0
	for i < arrayjumlah {
		if d[i].nama == cari {
			fmt.Println("Nama Barang:", d[i].nama, "Jenis Barang:", d[i].jenis, "Jumlah Stok:", d[i].jumlahStok, "Kategori Barang:", d[i].kategori)
		}
		i++
	}
	main()
}
func pencarianJenis() {
	var cari string
	fmt.Print("Masukkan kata kunci yang dicari: ")
	fmt.Scan(&cari)
	i = 0
	for i < arrayjumlah {
		if d[i].jenis == cari {
			fmt.Println("Nama Barang:", d[i].nama, "Jenis Barang:", d[i].jenis, "Jumlah Stok:", d[i].jumlahStok, "Kategori Barang:", d[i].kategori)
		}
		i++
	}
	main()
}
func pencarianStok() {
	var menuInput int
	fmt.Println("Masukkan jenis pencarian stok")
	fmt.Println("1. Binary Search")
	fmt.Println("2. Selection Sort (ASCENDING)")
	fmt.Println("3. Insertion Sort (DESCENDING)")
	fmt.Scan(&menuInput)
	switch menuInput {
	case 1:
		binarySearch()
	case 2:
		selectionSort()
	case 3:
		insertionSort()
	case 4:
		main()
	}
}
func binarySearch() {
	var menuInput int
	fmt.Println("Masukkan Angka Jumlah Stok yang ingin di cari")
	fmt.Scan(&menuInput)
	binSearch(d, menuInput)
	if binSearch(d, menuInput) == -1 {
		fmt.Println("Jumlah Stok Tidak Ketemu")
	} else {
		fmt.Println("Angka", menuInput, "Ketemu Pada Array", binSearch(d, menuInput))
	}
	main()
}
func binSearch(d tabBarang, n int) int {
	var kr, kn, med int
	var found int = -1
	kr = 0
	kn = arrayjumlah - 1
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if n < d[med].jumlahStok {
			kn = med - 1
			kr = kr - 1
		} else if n > d[med].jumlahStok {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}
func selectionSort() {
	secSort(&d)
	tampilArraySort(d)
	main()
}
func secSort(d *tabBarang) {
	var i, j, idx_min int
	var t barang
	i = 1
	for i <= arrayjumlah-1 {
		idx_min = i - 1
		j = i
		for j < arrayjumlah {
			if d[idx_min].jumlahStok > d[j].jumlahStok {
				idx_min = j
			}
			j = j + 1
		}
		t = d[idx_min]
		d[idx_min] = d[i-1]
		d[i-1] = t
		i = i + 1
	}
}
func tampilArraySort(d tabBarang) {
	for i = 0; i < arrayjumlah; i++ {
		fmt.Println(i, d[i].nama, d[i].jenis, d[i].jumlahStok, d[i].kategori)
	}
}
func insertionSort() {
	inSort(&d)
	tampilArraySort(d)
	main()
}
func inSort(d *tabBarang) {
	var i, j int
	var temp barang
	i = 1
	for i <= arrayjumlah-1 {
		j = i
		temp = d[j]
		for j > 0 && temp.jumlahStok > d[j-1].jumlahStok {
			d[j] = d[j-1]
			j = j - 1
		}
		d[j] = temp
		i = i + 1
	}
}

func pencarianbarangkategori() {
	var cari string
	fmt.Print("Masukkan kategori yang dicari: ")
	fmt.Scan(&cari)
	i = 0
	for i < arrayjumlah {
		if d[i].kategori == cari {
			fmt.Println("Nama Barang:", d[i].nama, "Jenis Barang:", d[i].jenis, "Jumlah Stok:", d[i].jumlahStok, "Kategori Barang:", d[i].kategori)
		}
		i++
	}
	main()
}
