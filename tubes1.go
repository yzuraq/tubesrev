package main

import (
	"fmt"
	"strings"
)

const NMAX = 100

type userLes struct {
	nama     string
	status   string
	username string
	email    string
	password string
}

type statusLogin struct {
	statusUser    string
	apakahLogin   bool
	userYangLogin userLes
}

type mataPelajaran struct {
	nama        string
	nilaiTryOut [3]float64
	nilaiRata   float64
}

type siswa struct {
	id            int
	nama          string
	kelas         string
	username      string
	status        string
	userLes       userLes
	mataPelajaran [5]mataPelajaran
	jadwalLes     [5]jadwalLes
	telepon       string
	email         string
	tanggalLahir  string
	catatan       string
	totalNilai    float64
	ranking       int
}

type jadwalLes struct {
	day  string
	time string
}

type dataSiswa struct {
	listSiswa [NMAX]siswa
	nSiswa    int
}

type admin struct {
	id       int
	nama     string
	username string
	status   string
	userLes  userLes
}

type dataAdmin struct {
	listAdmin [NMAX]userLes
	nAdmin    int
}

func inputWithSpace(totalString *string) {
	var stringPanjang, inputKata string
	stringPanjang = ""
	fmt.Scan(&inputKata)
	for inputKata != "." {
		stringPanjang += inputKata + " "
		fmt.Scan(&inputKata)
	}
	*totalString = stringPanjang
}

func register(tabAdmin *dataAdmin, tabSiswa *dataSiswa, dataLogin *statusLogin) {
	var inputStatus string
	var userNameSama int
	fmt.Println()
	fmt.Println("============Register==============")
	fmt.Println()
	fmt.Print("Apa status anda? (Admin/Siswa): ")
	fmt.Scan(&inputStatus)
	for inputStatus != "Admin" && inputStatus != "Siswa" {
		fmt.Print("Status anda salah, status yang tersedia adalah Admin dan Siswa. Huruf pertama harus menggunakan huruf kapital")
		fmt.Scan(&inputStatus)
	}

	if inputStatus == "Admin" {
		// Menyimpan data user yang baru pada tabAdmin
		tabAdmin.listAdmin[tabAdmin.nAdmin].status = "Admin"
		fmt.Print("Masukkan nama lengkap Anda (akhiri dgn '.'): ")
		inputWithSpace(&tabAdmin.listAdmin[tabAdmin.nAdmin].nama)
		fmt.Print("Masukkan username Anda: ")
		fmt.Scan(&tabAdmin.listAdmin[tabAdmin.nAdmin].username)

		// Looping jika username sudah ada sebelumnya
		userNameSama = adaUsernameYangSama(*tabAdmin, *tabSiswa, tabAdmin.listAdmin[tabAdmin.nAdmin].username)
		for userNameSama != 0 {
			fmt.Print("Username sudah ada. Masukkan username lain.")
			fmt.Scan(&tabAdmin.listAdmin[tabAdmin.nAdmin].username)
			userNameSama = adaUsernameYangSama(*tabAdmin, *tabSiswa, tabAdmin.listAdmin[tabAdmin.nAdmin].username)
		}
		fmt.Print("Masukkan email Anda: ")
		fmt.Scan(&tabAdmin.listAdmin[tabAdmin.nAdmin].email)
		fmt.Print("Masukkan password Anda: ")
		fmt.Scan(&tabAdmin.listAdmin[tabAdmin.nAdmin].password)
		tabAdmin.nAdmin++

	} else if inputStatus == "Siswa" {
		// Menyimpan data user yang baru pada tabSiswa
		tabSiswa.listSiswa[tabSiswa.nSiswa].status = "Siswa"
		fmt.Print("Masukkan nama lengkap anda (akhiri dgn '.'): ")
		inputWithSpace(&tabSiswa.listSiswa[tabSiswa.nSiswa].nama)
		fmt.Print("Masukkan username Anda: ")
		fmt.Scan(&tabSiswa.listSiswa[tabSiswa.nSiswa].username)
		userNameSama = adaUsernameYangSama(*tabAdmin, *tabSiswa, tabSiswa.listSiswa[tabSiswa.nSiswa].username)
		for userNameSama != 0 {
			fmt.Print("Username sudah ada. Masukkan username lain.")
			fmt.Scan(&tabSiswa.listSiswa[tabSiswa.nSiswa].username)
			userNameSama = adaUsernameYangSama(*tabAdmin, *tabSiswa, tabSiswa.listSiswa[tabSiswa.nSiswa].username)
		}

		tabSiswa.listSiswa[tabSiswa.nSiswa].nama = tabSiswa.listSiswa[tabSiswa.nSiswa].nama
		fmt.Print("Masukkan email Anda: ")
		fmt.Scan(&tabSiswa.listSiswa[tabSiswa.nSiswa].userLes.email)
		fmt.Print("Masukkan password Anda: ")
		fmt.Scan(&tabSiswa.listSiswa[tabSiswa.nSiswa].userLes.password)
		tabSiswa.nSiswa++
	}

	fmt.Println("Registrasi berhasil! Silahkan login sekarang.")
	// menuUtama(tabAdmin, tabSiswa, dataLogin)
}

func login(tabAdmin *dataAdmin, tabSiswa *dataSiswa, dataLogin *statusLogin) {
	var inputUsername, inputPassword string
	var idxUsername int
	var percobaanLogin int = 0

	fmt.Println()
	fmt.Println("=============Log In===============")
	fmt.Println()

	fmt.Print("Masukkan username anda: ")
	fmt.Scan(&inputUsername)
	idxUsername = adaUsernameYangSama(*tabAdmin, *tabSiswa, inputUsername)

	for idxUsername == 0 && percobaanLogin < 3 {
		percobaanLogin++
		fmt.Println("Username tidak ditemukan")
		fmt.Print("Masukkan username anda: ")
		fmt.Scan(&inputUsername)
		idxUsername = adaUsernameYangSama(*tabAdmin, *tabSiswa, inputUsername)
		if percobaanLogin == 3 {
			fmt.Println("Anda telah mencoba 3 kali, silahkan registrasi terlebih dahulu")
			register(tabAdmin, tabSiswa, dataLogin)
			fmt.Println("Silahkan login kembali")
			fmt.Print("Masukkan username anda: ")
			fmt.Scan(&inputUsername)
			idxUsername = adaUsernameYangSama(*tabAdmin, *tabSiswa, inputUsername)
		}
	}

	if idxUsername >= 2000000 {
		dataLogin.statusUser = "Siswa"
		dataLogin.apakahLogin = true
		dataLogin.userYangLogin = tabSiswa.listSiswa[idxUsername-2000000].userLes
		dataLogin.userYangLogin.status = "Siswa"

		// Explicitly assign the username
		dataLogin.userYangLogin.username = tabSiswa.listSiswa[idxUsername-2000000].username

		fmt.Print("Masukkan password anda: ")
		fmt.Scan(&inputPassword)
		for inputPassword != dataLogin.userYangLogin.password {
			fmt.Println("Password salah")
			fmt.Print("Masukkan password anda: ")
			fmt.Scan(&inputPassword)
		}

		// Debugging: print out the username after login
		fmt.Println("Login successful. Username:", dataLogin.userYangLogin.username) // Added this line
	} else {
		dataLogin.statusUser = "Admin"
		dataLogin.apakahLogin = true
		dataLogin.userYangLogin = tabAdmin.listAdmin[idxUsername-1000000]

		// Explicitly assign the username for Admin too
		dataLogin.userYangLogin.username = tabAdmin.listAdmin[idxUsername-1000000].username

		fmt.Print("Masukkan password anda: ")
		fmt.Scan(&inputPassword)
		for inputPassword != dataLogin.userYangLogin.password {
			fmt.Println("Password salah")
			fmt.Print("Masukkan password anda: ")
			fmt.Scan(&inputPassword)
		}
		dataLogin.userYangLogin.status = "Admin"

		// Debugging: print out the username after login
		fmt.Println("Login successful. Username:", dataLogin.userYangLogin.username) // Added this line
	}

	fmt.Println("Login berhasil")
	fmt.Print("Selamat datang ", dataLogin.userYangLogin.nama, " pada apps Brained")
	fmt.Println()
	fmt.Println("Status anda adalah", dataLogin.userYangLogin.status)

	// After successful login, navigate to the appropriate menu
	if dataLogin.userYangLogin.status == "Admin" {
		menuAdmin(tabSiswa, tabAdmin, dataLogin)
	} else if dataLogin.userYangLogin.status == "Siswa" {
		menuSiswa(tabSiswa, dataLogin, tabAdmin)
	}
}

func adaUsernameYangSama(tabAdmin dataAdmin, tabSiswa dataSiswa, inputUsername string) int {
	for i := 0; i < tabAdmin.nAdmin; i++ {
		if tabAdmin.listAdmin[i].username == inputUsername {
			return 1000000 + i
		}
	}
	for i := 0; i < tabSiswa.nSiswa; i++ {
		if tabSiswa.listSiswa[i].username == inputUsername {
			return 2000000 + i
		}
	}
	return 0
}

func hitungNilaiRataMataPelajaran(mp *mataPelajaran) {
	var total float64
	for i := 0; i < 3; i++ {
		total += mp.nilaiTryOut[i]
	}
	mp.nilaiRata = total / 3
}

func hitungTotalNilaiSiswa(siswa *siswa) {
	var total float64
	for i := 0; i < 5; i++ {
		total += siswa.mataPelajaran[i].nilaiRata
	}
	siswa.totalNilai = total / 5
}

func insertionSortRanking(tabSiswa *dataSiswa) {
	for i := 1; i < tabSiswa.nSiswa; i++ {
		siswaTemp := tabSiswa.listSiswa[i]
		j := i - 1
		for j >= 0 && tabSiswa.listSiswa[j].totalNilai < siswaTemp.totalNilai {
			tabSiswa.listSiswa[j+1] = tabSiswa.listSiswa[j]
			j--
		}
		tabSiswa.listSiswa[j+1] = siswaTemp
	}

	for i := 0; i < tabSiswa.nSiswa; i++ {
		tabSiswa.listSiswa[i].ranking = i + 1
	}
}

func binarySearchSiswa(tabSiswa dataSiswa, targetID int) int {
	left, right := 0, tabSiswa.nSiswa-1
	for left <= right {
		mid := (left + right) / 2
		if tabSiswa.listSiswa[mid].id == targetID {
			return mid
		}
		if tabSiswa.listSiswa[mid].id < targetID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// MASUK MENU ADMIN SAMA FITURNYA
func menuAdmin(tabSiswa *dataSiswa, tabAdmin *dataAdmin, dataLogin *statusLogin) {
	var pilihan int
	continueMenu := true
	for continueMenu {
		fmt.Println("Menu Admin")
		fmt.Println("1. Input Data Siswa")
		fmt.Println("2. Input Nilai Try Out Siswa")
		fmt.Println("3. Tampilkan Data Siswa dan Urutkan Berdasarkan ID")
		fmt.Println("4. Edit Data Siswa")
		fmt.Println("5. Hapus Data Siswa")
		fmt.Println("6. Keluar Menu Admin")
		fmt.Print("Masukkan pilihan Anda: ")

		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			inputDataSiswa(tabSiswa)
		case 2:
			inputNilaiTryOut(tabSiswa)
		case 3:
			tampilkanDataSiswa(tabSiswa)
		case 4:
			editDataSiswa(tabSiswa)
		case 5:
			hapusSiswa(tabSiswa)
		case 6:
			fmt.Println("Keluar dari menu admin.")
			continueMenu = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanUsernameSiswa(tabSiswa *dataSiswa) {
	fmt.Println("Daftar Username Siswa:")
	for i := 0; i < tabSiswa.nSiswa; i++ {
		fmt.Printf("%d. Username: %s, Nama: %s\n", i+1, tabSiswa.listSiswa[i].username, tabSiswa.listSiswa[i].nama)
	}
}

func cariSiswaBerdasarkanUsername(tabSiswa dataSiswa, username string) int {
	for i := 0; i < tabSiswa.nSiswa; i++ {
		if tabSiswa.listSiswa[i].username == username {
			return i
		}
	}
	return -1
}

// MENU 1
func inputDataSiswa(tabSiswa *dataSiswa) {
	var username string
	var siswaIdx int
	var idSiswa int
	var kelas, mataPelajaranInput, telepon, tanggalLahir, jadwalHari, jadwalJam string

	tampilkanUsernameSiswa(tabSiswa)

	fmt.Print("Masukkan username siswa yang akan diinput data dirinya: ")
	fmt.Scan(&username)

	siswaIdx = cariSiswaBerdasarkanUsername(*tabSiswa, username)

	if siswaIdx == -1 {
		siswaIdx = tabSiswa.nSiswa
		tabSiswa.nSiswa++ // Increment nSiswa to add new student
	}

	fmt.Println("Siswa ditemukan:", tabSiswa.listSiswa[siswaIdx].nama)

	fmt.Print("Masukkan ID siswa: ")
	fmt.Scan(&idSiswa)
	fmt.Print("Masukkan kelas les siswa: ")
	fmt.Scan(&kelas)
	fmt.Print("Masukkan nomor telepon siswa: ")
	fmt.Scan(&telepon)
	fmt.Print("Masukkan tanggal lahir siswa (dd/mm/yyyy): ")
	fmt.Scan(&tanggalLahir)

	for i := 0; i < 3; i++ {
		fmt.Printf("Masukkan mata pelajaran ke-%d: ", i+1)
		fmt.Scan(&mataPelajaranInput)
		tabSiswa.listSiswa[siswaIdx].mataPelajaran[i].nama = mataPelajaranInput
	}

	// Separate day and time input for jadwal les
	for i := 0; i < 3; i++ {
		fmt.Printf("Masukkan jadwal les untuk mata pelajaran %s\n", tabSiswa.listSiswa[siswaIdx].mataPelajaran[i].nama)

		fmt.Print("Hari: ")
		fmt.Scan(&jadwalHari)

		fmt.Print("Jam: ")
		fmt.Scan(&jadwalJam)

		tabSiswa.listSiswa[siswaIdx].jadwalLes[i].day = jadwalHari
		tabSiswa.listSiswa[siswaIdx].jadwalLes[i].time = jadwalJam
	}

	tabSiswa.listSiswa[siswaIdx].username = username
	tabSiswa.listSiswa[siswaIdx].id = idSiswa
	tabSiswa.listSiswa[siswaIdx].kelas = kelas
	tabSiswa.listSiswa[siswaIdx].telepon = telepon
	tabSiswa.listSiswa[siswaIdx].tanggalLahir = tanggalLahir

	fmt.Println("Data siswa telah diperbarui.")
}

// MENU 2
func selectionSortSiswaByID(tabSiswa *dataSiswa) {
	for i := 0; i < tabSiswa.nSiswa-1; i++ {
		minIdx := i
		for j := i + 1; j < tabSiswa.nSiswa; j++ {
			if tabSiswa.listSiswa[j].id < tabSiswa.listSiswa[minIdx].id {
				minIdx = j
			}
		}
		if minIdx != i {
			tabSiswa.listSiswa[i], tabSiswa.listSiswa[minIdx] = tabSiswa.listSiswa[minIdx], tabSiswa.listSiswa[i]
		}
	}
}

func binarySearchSiswaByID(tabSiswa *dataSiswa, idSiswa int) int {
	left, right := 0, tabSiswa.nSiswa-1
	for left <= right {
		mid := (left + right) / 2
		if tabSiswa.listSiswa[mid].id == idSiswa {
			return mid
		}
		if tabSiswa.listSiswa[mid].id < idSiswa {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func inputNilaiTryOut(tabSiswa *dataSiswa) {
	var idSiswa, mataPelajaranIndex int
	var nilai float64

	fmt.Print("Masukkan ID siswa yang akan mengikuti try out:")
	fmt.Scan(&idSiswa)

	selectionSortSiswaByID(tabSiswa)

	siswaIdx := binarySearchSiswaByID(tabSiswa, idSiswa)

	if siswaIdx == -1 {
		fmt.Print("Siswa tidak ditemukan!")
		return
	}

	fmt.Println("Siswa ditemukan:", tabSiswa.listSiswa[siswaIdx].nama)
	fmt.Println("Mata Pelajaran yang dipilih (maksimal 5 mata pelajaran):")
	for i := 0; i < 3; i++ {
		if tabSiswa.listSiswa[siswaIdx].mataPelajaran[i].nama != "" {
			fmt.Printf("%d. %s\n", i+1, tabSiswa.listSiswa[siswaIdx].mataPelajaran[i].nama)
		}
	}

	for i := 0; i < 3; i++ {
		if tabSiswa.listSiswa[siswaIdx].mataPelajaran[i].nama != "" {
			mataPelajaranIndex = i
			for j := 0; j < 3; j++ {
				fmt.Printf("Masukkan nilai try out %d untuk mata pelajaran %s: ", j+1, tabSiswa.listSiswa[siswaIdx].mataPelajaran[i].nama)
				fmt.Scan(&nilai)
				tabSiswa.listSiswa[siswaIdx].mataPelajaran[i].nilaiTryOut[j] = nilai
			}
			hitungNilaiRataMataPelajaran(&tabSiswa.listSiswa[siswaIdx].mataPelajaran[mataPelajaranIndex])
		}
	}

	hitungTotalNilaiSiswa(&tabSiswa.listSiswa[siswaIdx])
	insertionSortRanking(tabSiswa)

}

// MENU 3
func tampilkanDataSiswa(tabSiswa *dataSiswa) {
	selectionSortSiswaByID(tabSiswa)

	fmt.Println("= Data Siswa yang Diinputkan =")
	fmt.Println("=========================================================================")
	fmt.Println("| ID  | Nama Siswa       | Kelas  | Nomor Telepon       | Tanggal Lahir |")
	fmt.Println("=========================================================================")

	for i := 0; i < tabSiswa.nSiswa; i++ {
		fmt.Printf("| %-3d | %-16s | %-6s | %-18s | %-12s |\n",
			tabSiswa.listSiswa[i].id,
			tabSiswa.listSiswa[i].nama,
			tabSiswa.listSiswa[i].kelas,
			tabSiswa.listSiswa[i].telepon,
			tabSiswa.listSiswa[i].tanggalLahir)
	}

	fmt.Println("========================================================================")
}

// MENU 4
func editDataSiswa(tabSiswa *dataSiswa) {
	var idSiswa int
	var siswaIdx int
	var kelas, telepon, tanggalLahir, mataPelajaranInput, jadwalInput string

	fmt.Print("Masukkan ID siswa yang ingin diedit:")
	fmt.Scan(&idSiswa)

	selectionSortSiswaByID(tabSiswa)

	siswaIdx = binarySearchSiswaByID(tabSiswa, idSiswa)

	if siswaIdx == -1 {
		fmt.Println("Siswa tidak ditemukan!")
		return
	}

	fmt.Println("Siswa ditemukan:", tabSiswa.listSiswa[siswaIdx].nama)

	fmt.Print("Masukkan kelas les baru siswa (Kosongkan untuk tidak mengubah):")
	fmt.Scan(&kelas)
	if kelas != "" {
		tabSiswa.listSiswa[siswaIdx].kelas = kelas
	}

	fmt.Print("Masukkan nomor telepon baru siswa (Kosongkan untuk tidak mengubah):")
	fmt.Scan(&telepon)
	if telepon != "" {
		tabSiswa.listSiswa[siswaIdx].telepon = telepon
	}

	fmt.Print("Masukkan tanggal lahir baru siswa (dd/mm/yyyy) (Kosongkan untuk tidak mengubah):")
	fmt.Scan(&tanggalLahir)
	if tanggalLahir != "" {
		tabSiswa.listSiswa[siswaIdx].tanggalLahir = tanggalLahir
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Masukkan mata pelajaran ke-%d (Kosongkan jika tidak ingin mengedit): ", i+1)
		fmt.Scan(&mataPelajaranInput)
		if mataPelajaranInput != "" {
			tabSiswa.listSiswa[siswaIdx].mataPelajaran[i].nama = mataPelajaranInput
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Masukkan jadwal les untuk mata pelajaran %s (Kosongkan jika tidak ingin mengedit):\n", tabSiswa.listSiswa[siswaIdx].mataPelajaran[i].nama)
		fmt.Scan(&jadwalInput)
		if jadwalInput != "" {
			tabSiswa.listSiswa[siswaIdx].jadwalLes[i].day = jadwalInput
			tabSiswa.listSiswa[siswaIdx].jadwalLes[i].time = jadwalInput
		}
	}

	fmt.Println("Data siswa telah diperbarui.")
}

// MENU 5
func hapusSiswa(tabSiswa *dataSiswa) {
	var idSiswa int

	fmt.Print("Masukkan ID siswa yang ingin dihapus:")
	fmt.Scan(&idSiswa)

	selectionSortSiswaByID(tabSiswa)

	siswaIdx := binarySearchSiswaByID(tabSiswa, idSiswa)

	if siswaIdx == -1 {
		fmt.Println("Siswa tidak ditemukan!")
		return
	}

	fmt.Println("Siswa ditemukan:", tabSiswa.listSiswa[siswaIdx].nama)

	for i := siswaIdx; i < tabSiswa.nSiswa-1; i++ {
		tabSiswa.listSiswa[i] = tabSiswa.listSiswa[i+1]
	}

	tabSiswa.nSiswa--

	fmt.Println("Siswa telah dihapus.")
}

// Menu Siswa
func menuSiswa(tabSiswa *dataSiswa, dataLogin *statusLogin, tabAdmin *dataAdmin) {
	var pilihan int
	continueMenu := true

	for continueMenu {
		fmt.Println("Menu Siswa")
		fmt.Println("1. Lihat Data Diri")
		fmt.Println("2. Lihat Nilai Try Out Mata Pelajaran")
		fmt.Println("3. Lihat Ranking Siswa")
		fmt.Println("4. Cek Jadwal Les")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Masukkan pilihan Anda: ")

		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			lihatDataDiri(tabSiswa, dataLogin)
		case 2:
			lihatNilaiTryOut(tabSiswa, dataLogin, tabAdmin)
		case 3:
			tampilkanRankingSiswa(tabSiswa)
		case 4:
			cekJadwalLes(tabSiswa, dataLogin)
		case 5:
			// menuUtama(tabAdmin, tabSiswa, dataLogin) // Go back to the main menu
			continueMenu = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// MENU 1
func lihatDataDiri(tabSiswa *dataSiswa, dataLogin *statusLogin) {
	username := dataLogin.userYangLogin.username

	var idxSiswa int
	for i := 0; i < tabSiswa.nSiswa; i++ {
		if tabSiswa.listSiswa[i].username == username {
			idxSiswa = i

		}
	}

	if idxSiswa == -1 {
		fmt.Println("Siswa tidak ditemukan!")
		return
	}

	fmt.Println("=== Data Diri Siswa ===")
	fmt.Println("Nama Siswa:", tabSiswa.listSiswa[idxSiswa].nama)
	fmt.Println("Kelas:", tabSiswa.listSiswa[idxSiswa].kelas)
	fmt.Println("Nomor Telepon:", tabSiswa.listSiswa[idxSiswa].telepon)
	fmt.Println("Tanggal Lahir:", tabSiswa.listSiswa[idxSiswa].tanggalLahir)

	fmt.Println("Mata Pelajaran yang Diambil:")
	for i := 0; i < 5; i++ {
		if tabSiswa.listSiswa[idxSiswa].mataPelajaran[i].nama != "" {
			fmt.Printf("%d. %s\n", i+1, tabSiswa.listSiswa[idxSiswa].mataPelajaran[i].nama)
		}
	}
}

// MENU 2
func lihatNilaiTryOut(tabSiswa *dataSiswa, dataLogin *statusLogin, tabAdmin *dataAdmin) {
	var mataPelajaranInput string
	found := false

	// Ensure username from login is correctly assigned
	username := strings.ToLower(strings.TrimSpace(dataLogin.userYangLogin.username))
	fmt.Println("Searching for student with username:", username) // Debugging line

	// Call cariSiswaBerdasarkanUsername function to get the student index
	idxSiswa := cariSiswaBerdasarkanUsername(*tabSiswa, username)

	// If student not found, return to main menu
	if idxSiswa == -1 {
		fmt.Println("Siswa tidak ditemukan!")
		menuUtama(tabAdmin, tabSiswa, dataLogin) // Return to main menu
		return
	}

	// Output student's data and try-out scores
	fmt.Println("=== Nilai Try Out Siswa ===")
	fmt.Println("Nama Siswa:", tabSiswa.listSiswa[idxSiswa].nama)
	fmt.Println("Status Siswa:", tabSiswa.listSiswa[idxSiswa].status)
	fmt.Println("==================================")

	// Input mata pelajaran to search for
	fmt.Print("Masukkan nama mata pelajaran yang ingin dilihat nilainya: ")
	fmt.Scan(&mataPelajaranInput)

	// Convert input to lowercase for case-insensitive comparison
	mataPelajaranInput = strings.ToLower(mataPelajaranInput)

	// Search for the subject and display the try-out scores
	for i := 0; i < len(tabSiswa.listSiswa[idxSiswa].mataPelajaran); i++ {
		// Convert stored mata pelajaran to lowercase for comparison
		subjectName := strings.ToLower(tabSiswa.listSiswa[idxSiswa].mataPelajaran[i].nama)

		if subjectName == mataPelajaranInput {
			fmt.Println("==================================")
			fmt.Printf("Mata Pelajaran: %s\n", mataPelajaranInput)
			fmt.Printf("Nilai Try Out: %v\n", tabSiswa.listSiswa[idxSiswa].mataPelajaran[i].nilaiTryOut)
			fmt.Println("==================================")
			found = true
		}
	}

	if !found {
		fmt.Println("Mata pelajaran yang Anda pilih tidak ada dalam jadwal Anda.")
	}

}

// MENU 3
func tampilkanRankingSiswa(tabSiswa *dataSiswa) {
	fmt.Println("============================================")
	fmt.Println("=== Ranking Siswa Berdasarkan Total Nilai ===")
	fmt.Println("============================================")
	fmt.Println("| No |       Nama Siswa        | Total Nilai |")
	fmt.Println("============================================")

	for i := 0; i < tabSiswa.nSiswa; i++ {
		fmt.Printf("| %2d | %-22s | %-12.2f |\n", tabSiswa.listSiswa[i].ranking, tabSiswa.listSiswa[i].nama, tabSiswa.listSiswa[i].totalNilai)
	}

	fmt.Println("============================================")
}

// MENU 4
func cekJadwalLes(tabSiswa *dataSiswa, dataLogin *statusLogin) {
	var mataPelajaranInput string
	var found bool

	// Search for the student's username
	username := dataLogin.userYangLogin.username
	fmt.Println("Searching for student with username:", username) // Add this line for debugging

	idxSiswa := -1
	for i := 0; i < tabSiswa.nSiswa; i++ {
		if tabSiswa.listSiswa[i].username == username {
			idxSiswa = i
			break
		}
	}

	if idxSiswa == -1 {
		fmt.Println("Siswa tidak ditemukan!")
		return
	}

	fmt.Println("=== Jadwal Les Siswa ===")
	fmt.Println("Nama Siswa:", tabSiswa.listSiswa[idxSiswa].nama)
	fmt.Println("Status Siswa:", tabSiswa.listSiswa[idxSiswa].status)
	fmt.Println("==================================")

	fmt.Print("Masukkan nama mata pelajaran: ")
	fmt.Scan(&mataPelajaranInput)

	for i := 0; i < len(tabSiswa.listSiswa[idxSiswa].mataPelajaran); i++ {
		if tabSiswa.listSiswa[idxSiswa].mataPelajaran[i].nama == mataPelajaranInput {
			fmt.Println("==================================")
			fmt.Printf("Mata Pelajaran: %s\n", mataPelajaranInput)
			fmt.Printf("     Jadwal: Hari: %s, Jam: %s\n",
				tabSiswa.listSiswa[idxSiswa].jadwalLes[i].day,
				tabSiswa.listSiswa[idxSiswa].jadwalLes[i].time)
			fmt.Println("==================================")
			found = true
		}
	}

	if !found {
		fmt.Println("Mata pelajaran tidak ditemukan di jadwal Anda.")
	}
}

// PEMANGGILAN MENU
func menuUtama(tabAdmin *dataAdmin, tabSiswa *dataSiswa, dataLogin *statusLogin) {
	var choice int
	var continueMenu = true

	for continueMenu {
		fmt.Println("Selamat datang di sistem administrasi Brained!")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Print("Masukan Pilihan Anda : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			register(tabAdmin, tabSiswa, dataLogin)
		case 2:
			login(tabAdmin, tabSiswa, dataLogin)
		case 3:
			fmt.Println("Terima kasih telah menggunakan sistem administrasi Brained!")
			continueMenu = false
		default:
			fmt.Println("Pilihan tidak valid.")

			// fmt.Println("Apakah Anda ingin kembali ke menu utama? (Y/N)")
			// var kembali string
			// fmt.Scan(&kembali)

			// if kembali != "Y" && kembali != "y" {
			// 	if dataLogin.userYangLogin.status == "Admin" {
			// 		menuAdmin(tabSiswa, tabAdmin, dataLogin)
			// 	} else if dataLogin.userYangLogin.status == "Siswa" {
			// 		menuSiswa(tabSiswa, dataLogin, tabAdmin)
			// 	}
			// 	continueMenu = false
			// }
		}
	}
}

func tampilkanWelcomePage() {
	letters := map[rune][]string{
		'B': {
			"██████  ",
			"██   ██ ",
			"██████  ",
			"██   ██ ",
			"██████  ",
		},
		'R': {
			"██████  ",
			"██   ██ ",
			"██████  ",
			"██   ██ ",
			"██   ██ ",
		},
		'A': {
			"   ██   ",
			"  ████  ",
			" ██  ██ ",
			"████████",
			"██    ██",
		},
		'I': {
			"██████",
			"  ██  ",
			"  ██  ",
			"  ██  ",
			"██████",
		},
		'N': {
			"██    ██",
			"███   ██",
			"████  ██",
			"██ ██ ██",
			"██   ███",
		},
		'E': {
			"████████",
			"██      ",
			"██████  ",
			"██      ",
			"████████",
		},
		'D': {
			"██████  ",
			"██   ██ ",
			"██    ██",
			"██   ██ ",
			"██████  ",
		},
	}

	word := "BRAINED"

	fmt.Println()

	// Menampilkan setiap baris dari ASCII art
	for row := 0; row < 5; row++ {
		line := ""
		for _, char := range word {
			if asciiArt, exists := letters[char]; exists {
				line += asciiArt[row] + " "
			}
		}
		fmt.Println(line)
	}

	fmt.Println()
}

func main() {
	var tabAdmin dataAdmin
	var tabSiswa dataSiswa
	var dataLogin statusLogin

	tampilkanWelcomePage()

	menuUtama(&tabAdmin, &tabSiswa, &dataLogin)
}
