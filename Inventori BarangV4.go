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
var i, arrayjumlah, tambah int
var menuInput int

func main() {

	fmt.Println("===== Masukkan Angka Untuk Menu =====")
	fmt.Println("1. Data Barang dan Transaksi")
	fmt.Println("2. Pencarian Kata Kunci Tertentu.")
	fmt.Println("3. Pencarian Kategori Tertentu.")
	fmt.Println("4. Tampilan semua data")
	fmt.Println("5. Preset Data")
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&menuInput)
	switch menuInput {
	case 1:
		fmt.Println("Pencatatan Data Barang dan Transaksi...")
		PencatatanBarangdanTransaksi()
	case 2:
		fmt.Println("Pencarian Kata Kunci Barang...")
		pemasukandata(a, &d, arrayjumlah)
		pencariankatakuncibarang(d, arrayjumlah)
	case 3:
		fmt.Println("Pencarian Kategori Tertentu...")
		pemasukandata(a, &d, arrayjumlah)
		pencarianbarangkategori(d, arrayjumlah)
	case 4:
		fmt.Println("Tampilan Semua Data Barang")
		fmt.Println(arrayjumlah, tambah)
		tampilarray(a, b, c, arrayjumlah)
		main()
	case 5:
		presetData(&a, &b, &c, &arrayjumlah, &tambah)
		main()
	}
}
func presetData(a *tabBarang, b *tabMasuk, c *tabKeluar, n *int, tambah *int) {
	*n = 5
	*tambah = 5
	a[0].nama, a[0].jenis, a[0].jumlahStok, a[0].kategori = "Ikan", "Hidup", 150, "Air"
	a[1].nama, a[1].jenis, a[1].jumlahStok, a[1].kategori = "Ikan", "Tidak", 120, "Air"
	a[2].nama, a[2].jenis, a[2].jumlahStok, a[2].kategori = "Ayam", "Hidup", 200, "Darat"
	a[3].nama, a[3].jenis, a[3].jumlahStok, a[3].kategori = "Ayam", "Tidak", 220, "Darat"
	a[4].nama, a[4].jenis, a[4].jumlahStok, a[4].kategori = "Burung", "Hidup", 300, "Udara"

	b[0].jam, b[0].menit, b[0].detik = 05, 05, 05
	b[1].jam, b[1].menit, b[1].detik = 10, 10, 10
	b[2].jam, b[2].menit, b[2].detik = 15, 15, 15
	b[3].jam, b[3].menit, b[3].detik = 20, 20, 20
	b[4].jam, b[4].menit, b[4].detik = 23, 30, 00

	c[0].jam, c[0].menit, c[0].detik = 10, 10, 10
	c[1].jam, c[1].menit, c[1].detik = 15, 15, 15
	c[2].jam, c[2].menit, c[2].detik = 20, 20, 20
	c[3].jam, c[3].menit, c[3].detik = 23, 30, 00
	c[4].jam, c[4].menit, c[4].detik = 05, 05, 05
	fmt.Println("Data Masuk:")
	for i := 0; i < *n; i++ {
		fmt.Println(i, a[i].nama, a[i].jenis, a[i].jumlahStok, a[i].kategori)
	}
	fmt.Println("Pada Pukul")
	for i := 0; i < *n; i++ {
		fmt.Println(i, b[i].jam, ":", b[i].menit, ":", b[i].detik)
	}
	fmt.Println("Keluar Pada Pukul")
	for i := 0; i < *n; i++ {
		fmt.Println(i, c[i].jam, ":", c[i].menit, ":", c[i].detik)
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
		pemasukanBarangdanTransaksi(a, b, c, arrayjumlah, tambah)
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

func pemasukanBarangdanTransaksi(a tabBarang, b tabMasuk, c tabKeluar, arrayjumlah int, tambah int) {
	var menuInput int
	fmt.Print("Masukkan Jumlah Data yang ingin di Masukkan: ")
	fmt.Scan(&menuInput)
	isiArray(&a, &b, &c, &menuInput, &arrayjumlah, &tambah)
	tampilarray(a, b, c, arrayjumlah)
	PencatatanBarangdanTransaksi()
}

func isiArray(a *tabBarang, b *tabMasuk, c *tabKeluar, n *int, arrayjumlah *int, i *int) {
	*arrayjumlah += *n
	for *i < *arrayjumlah {
		fmt.Println("Masukkan Barang")
		fmt.Println("Jenis - Nama - Jumlah Stok - Kategori")
		fmt.Scan(&a[*i].nama, &a[*i].jenis, &a[*i].jumlahStok, &a[*i].kategori)
		fmt.Println("Masukkan Waktu Transkasi Barang Masuk")
		fmt.Println("Jam - Menit - Detik")
		fmt.Scan(&b[*i].jam, &b[*i].menit, &b[*i].detik)
		fmt.Println("Masukkan Waktu Transkasi Barang Keluar")
		fmt.Println("Jam - Menit - Detik")
		fmt.Scan(&c[*i].jam, &c[*i].menit, &c[*i].detik)
		*i++
	}
}

func tampilarray(a tabBarang, b tabMasuk, c tabKeluar, arrayjumlah int) {
	fmt.Println("Data Masuk:")
	for i := 0; i < arrayjumlah; i++ {
		fmt.Println(i, a[i].nama, a[i].jenis, a[i].jumlahStok, a[i].kategori)
	}
	fmt.Println("Pada Pukul")
	for i := 0; i < arrayjumlah; i++ {
		fmt.Println(i, b[i].jam, ":", b[i].menit, ":", b[i].detik)
	}
	fmt.Println("Keluar Pada Pukul")
	for i := 0; i < arrayjumlah; i++ {
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
	tampilarray(a, b, c, arrayjumlah)
	fmt.Println("Urutan Ke berapa yang ingin di ubah?")
	fmt.Print("Masukkan Angka Urutan: ")
	fmt.Scan(&menuInput)
	Perubahantrandanbar(&a, &b, &c, &menuInput, tambah)
	tampilarray(a, b, c, arrayjumlah)
	perubahanBarangdanTransaksi()
}

func Perubahantrandanbar(a *tabBarang, b *tabMasuk, c *tabKeluar, n *int, tambah int) {
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
	tampilarray(a, b, c, arrayjumlah)
	fmt.Println("Urutan Ke berapa yang ingin di ubah?")
	fmt.Print("Masukkan Angka Urutan: ")
	fmt.Scan(&menuInput)
	Perubahandat(&a, &menuInput)
	tampilarray(a, b, c, arrayjumlah)
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
	tampilarray(a, b, c, arrayjumlah)
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
		tampilarray(a, b, c, arrayjumlah)
		perubahanBarangdanTransaksi()
	case 2:
		Perubahantrankel(&c, &menuInput)
		tampilarray(a, b, c, arrayjumlah)
		perubahanBarangdanTransaksi()
	case 3:
		Perubahantranmaskel(&b, &c, &menuInput)
		tampilarray(a, b, c, arrayjumlah)
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
	tampilarray(a, b, c, arrayjumlah)
	fmt.Println("Urutan Ke berapa yang ingin di hapus?")
	fmt.Print("Masukkan Angka Urutan: ")
	fmt.Scan(&menuInput)
	Penghapusantrandanbar(&a, &b, &c, &menuInput, &arrayjumlah)
	tampilarray(a, b, c, arrayjumlah)
	penghapusanBarangdanTransaksi()
}

func Penghapusantrandanbar(a *tabBarang, b *tabMasuk, c *tabKeluar, n *int, arrayjumlah *int) {
	for x := *n; x < *arrayjumlah; x++ {
		a[x] = a[x+1]
		b[x] = b[x+1]
		c[x] = c[x+1]
	}
	*arrayjumlah--
}

func penghapusandata() {
	var menuInput int
	fmt.Println("Data Sebelum Hapus:")
	tampilarray(a, b, c, arrayjumlah)
	fmt.Println("Urutan Ke berapa yang ingin di hapus?")
	fmt.Print("Masukkan Angka Urutan: ")
	fmt.Scan(&menuInput)
	Penghapusandat(&a, &menuInput, &arrayjumlah)
	tampilarray(a, b, c, arrayjumlah)
	penghapusanBarangdanTransaksi()
}

func Penghapusandat(a *tabBarang, n *int, arrayjumlah *int) {
	for x := *n; x < *arrayjumlah; x++ {
		a[x] = a[x+1]
	}
	*arrayjumlah--
}

func penghapusantransaki() {
	var menuInput, x int
	fmt.Println("Data Sebelum Hapus:")
	tampilarray(a, b, c, arrayjumlah)
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
		Penghapusantranmas(&b, &menuInput, &arrayjumlah)
		tampilarray(a, b, c, arrayjumlah)
		penghapusanBarangdanTransaksi()
	case 2:
		Penghapusantrankel(&c, &menuInput, &arrayjumlah)
		tampilarray(a, b, c, arrayjumlah)
		penghapusanBarangdanTransaksi()
	case 3:
		Penghapusantranmaskel(&b, &c, &menuInput, &arrayjumlah)
		tampilarray(a, b, c, arrayjumlah)
		penghapusanBarangdanTransaksi()
	default:
		perubahanBarangdanTransaksi()
	}
}

func Penghapusantranmas(b *tabMasuk, n *int, arrayjumlah *int) {
	for x := *n; x < *arrayjumlah; x++ {
		b[x] = b[x+1]
	}
	*arrayjumlah--
}

func Penghapusantrankel(c *tabKeluar, n *int, arrayjumlah *int) {
	for x := *n; x < *arrayjumlah; x++ {
		c[x] = c[x+1]
	}
	*arrayjumlah--
}

func Penghapusantranmaskel(b *tabMasuk, c *tabKeluar, n *int, arrayjumlah *int) {
	for x := *n; x < *arrayjumlah; x++ {
		b[x] = b[x+1]
		c[x] = c[x+1]
	}
	*arrayjumlah--
}

func pencariankatakuncibarang(d tabBarang, arrayjumlah int) {
	var menuInput int
	fmt.Println("Masukkan Kata kunci yang ingin dicari")
	fmt.Println("1. Nama")
	fmt.Println("2. Jenis")
	fmt.Println("3. Jumlah Stok")
	fmt.Println("4. Kembali ke Menu")
	fmt.Scan(&menuInput)
	switch menuInput {
	case 1:
		pencarianNama(d, arrayjumlah)
	case 2:
		pencarianJenis(d, arrayjumlah)
	case 3:
		pencarianStok(d, arrayjumlah)
	case 4:
		main()
	}
}
func pencarianNama(d tabBarang, arrayjumlah int) {
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
func pencarianJenis(d tabBarang, arrayjumlah int) {
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
func pencarianStok(d tabBarang, arrayjumlah int) {
	var menuInput int
	fmt.Println("Masukkan jenis pencarian stok")
	fmt.Println("1. Binary Search")
	fmt.Println("2. Selection Sort (ASCENDING)")
	fmt.Println("3. Insertion Sort (DESCENDING)")
	fmt.Println("4. Sequential Search")
	fmt.Scan(&menuInput)
	switch menuInput {
	case 1:
		binarySearch(d, arrayjumlah)
	case 2:
		selectionSort(d, arrayjumlah)
	case 3:
		insertionSort(d, arrayjumlah)
	case 4:
		sequentianSearchStok(d, arrayjumlah)
	case 5:
		main()
	}
}
func binarySearch(d tabBarang, arrayjumlah int) {
	var menuInput int
	fmt.Println("Masukkan Angka Jumlah Stok yang ingin di cari")
	fmt.Scan(&menuInput)
	binSearch(d, menuInput, arrayjumlah)
	if binSearch(d, menuInput, arrayjumlah) == -1 {
		fmt.Println("Jumlah Stok Tidak Ketemu")
		main()
	} else {
		fmt.Println("Angka", menuInput, "Ketemu Pada Array", binSearch(d, menuInput, arrayjumlah))
		main()
	}
	main()
}
func binSearch(d tabBarang, n int, arrayjumlah int) int {
	var kr, kn, med int
	var found int = -1
	kr = 0
	kn = arrayjumlah - 2
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
func sequentianSearchStok(d tabBarang, arrayjumlah int) {
	var menuInput int
	fmt.Println("Masukkan Angka jumlah stok yang akan di cari")
	fmt.Scan(&menuInput)
	stockSearch(d, arrayjumlah, menuInput)
}
func stockSearch(d tabBarang, arrayjumlah int, x int) {
	var z int
	for i := 0; i < arrayjumlah; i++ {
		if x == d[i].jumlahStok {
			fmt.Println("Angka", x, "Ketemu pada array ke", i)
			fmt.Println("Nama Barang:", d[i].nama, "Jenis Barang:", d[i].jenis, "Jumlah Stok:", d[i].jumlahStok, "Kategori Barang:", d[i].kategori)
			z++
		}
	}
	if z == 0 {
		fmt.Println("Angka Stok Tidak Ketemu")
	}
}
func selectionSort(d tabBarang, arrayjumlah int) {
	secSort(&d, arrayjumlah)
	tampilArraySort(d, arrayjumlah)
	main()
}
func secSort(d *tabBarang, arrayjumlah int) {
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
func tampilArraySort(d tabBarang, arrayjumlah int) {
	for i := 0; i < arrayjumlah; i++ {
		fmt.Println(i, d[i].nama, d[i].jenis, d[i].jumlahStok, d[i].kategori)
	}
}
func insertionSort(d tabBarang, arrayjumlah int) {
	inSort(&d, arrayjumlah)
	tampilArraySort(d, arrayjumlah)
	main()
}
func inSort(d *tabBarang, arrayjumlah int) {
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

func pencarianbarangkategori(d tabBarang, arrayjumlah int) {
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
