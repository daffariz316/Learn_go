// // package main
// // import "fmt"
// // func main() {
// // 	fmt.Println("Hello, World!")
// // 	fmt.Println("This is a simple Go program.")
// // 	fmt.Println("It demonstrates basic syntax and structure.")
// // 	fmt.Println("Go is a statically typed, compiled programming language.")
// // }

// // package main
// // import "fmt"

// // func main() {
// // 	var name string = "John Doe"
// // 	var age int = 20 
// // 	height := 5.9

// // 	fmt.Println("Name:", name)
// // 	fmt.Println("Age:", age)
// // 	fmt.Println("Height:", height)
// // }

// // package main
// // import "fmt"
// // func main() {
// // 	var name string 
// // 	fmt.Println("Enter your name:")
// // 	fmt.Scanln(&name)
// // 	fmt.Println("Hello,", name)
// // 	fmt.Println("Welcome to the Go programming language!")
// // 	fmt.Println("This program demonstrates user input and output.")
// // 	fmt.Println("You can use fmt package for formatted I/O.")
// // 	fmt.Println("Enjoy coding in Go!")
// // }

// // package main
// // import "fmt"
// // func main() {
// // 	var nilai int
// // 	fmt.Print("Masukkan nilai: ")
// // 	fmt.Scanln(&nilai)

// // 	if nilai >= 90 {
// // 		fmt.Println("Nilai A")
// // 	}else if nilai >= 80 {
// // 		fmt.Println("Nilai B")
// // 	}else if nilai >= 70 {
// // 		fmt.Println("Nilai C")
// // 	}else if nilai >= 60 {
// // 		fmt.Println("Nilai D")
// // 	}else {
// // 		fmt.Println("Nilai E")
// // 	}
// // }

// package main
// import "fmt"
// func main() {
// 	for i := 1; i <= 5; i++{
// 		for j := 1; j <= i; j++{
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }
package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
	// Menampilkan menu awal
	for {
		fmt.Println("\nAplikasi Catatan Harian")
		fmt.Println("1. Lihat Semua Catatan")
		fmt.Println("2. Tambah Catatan")
		fmt.Println("3. Hapus Catatan")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu (1/2/3/4): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tampilkanCatatan()
		case 2:
			tambahCatatan()
		case 3:
			hapusCatatan()
		case 4:
			fmt.Println("Terima kasih! Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi untuk menampilkan semua catatan
func tampilkanCatatan() {
	file, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("\nDaftar Catatan:")
	line := 1
	for scanner.Scan() {
		fmt.Printf("%d. %s\n", line, scanner.Text())
		line++
	}
}

// Fungsi untuk menambah catatan
func tambahCatatan() {
	fmt.Print("Masukkan catatan baru: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Membaca input pengguna

	note := scanner.Text()
	file, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(note + "\n")
	if err != nil {
		fmt.Println("Error menulis ke file:", err)
		return
	}
	fmt.Println("Catatan berhasil ditambahkan!")
}

// Fungsi untuk menghapus catatan
func hapusCatatan() {
	var index int
	tampilkanCatatan()

	fmt.Print("Masukkan nomor catatan yang ingin dihapus: ")
	fmt.Scanln(&index)

	file, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	if index <= 0 || index > len(notes) {
		fmt.Println("Nomor catatan tidak valid.")
		return
	}

	// Menghapus catatan yang dipilih
	notes = append(notes[:index-1], notes[index:]...)

	// Menulis ulang file dengan catatan yang sudah dihapus
	file, err = os.Create("notes.txt")
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	for _, note := range notes {
		_, err = file.WriteString(note + "\n")
		if err != nil {
			fmt.Println("Error menulis ke file:", err)
			return
		}
	}

	fmt.Println("Catatan berhasil dihapus!")
}
