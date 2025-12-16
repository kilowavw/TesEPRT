package main

import "fmt"

func main() {
	var soal1benar, soal2benar, soal3benar, soal4benar, soal5benar bool
	var nilai int
	for {
		fmt.Println("Tes Quiz")
		fmt.Println("Silahkan Pilih Menu: ")
		fmt.Println("1. Soal 1 ")
		fmt.Println("2. Soal 2 ")
		fmt.Println("3. Soal 3 ")
		fmt.Println("4. Soal 4 ")
		fmt.Println("5. Soal 5 ")
		fmt.Println("6. Keluar")
		var pilihan int
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Println("Anda memilih Soal 1")
			fmt.Println("___ is the best football player in the world")
			fmt.Println("1. WHO")
			fmt.Println("2. WHAT")
			fmt.Println("3. WHEN")
			fmt.Println("4. WHERE")
			var soal1 int
			fmt.Scan(&soal1)
			if soal1 == 1 {
				fmt.Println("Jawaban Benar")
				soal1benar = true
			} else {
				fmt.Println("Jawaban Salah")
			}
		} else if pilihan == 2 {
			fmt.Println("Anda memilih Soal 2")
			fmt.Println("Siapa presiden Indonesia __ (huruf kecil)")
			var soal2 string
			fmt.Scan(&soal2)
			if soal2 == "prabowo" {
				fmt.Println("Jawaban Benar")
				soal2benar = true
			} else {
				fmt.Println("Jawaban Salah")
			}
		} else if pilihan == 3 {
			fmt.Println("Anda memilih Soal 3")
			fmt.Println("Apakah ijazah jokowi asli (tulis true atau false)")
			var soal3 bool
			fmt.Scan(&soal3)
			if soal3 == false {
				fmt.Println("Jawaban Benar")
				soal3benar = true
			} else {
				fmt.Println("Jawaban Salah")
			}
		} else if pilihan == 4 {
			fmt.Println("Anda memilih Soal 4")
			fmt.Println("Berapa angka dari phi")
			var soal4 float64
			fmt.Scan(&soal4)
			if soal4 == 3.14 {
				fmt.Println("Jawaban Benar")
				soal4benar = true
			} else {
				fmt.Println("Jawaban Salah")
			}
		} else if pilihan == 5 {
			fmt.Println("Anda memilih Soal 5")
			fmt.Println("Berapa 1+1 = ")
			var soal5 int
			fmt.Scan(&soal5)
			if soal5 == 2 {
				fmt.Println("Jawaban Benar")
				soal5benar = true
			} else {
				fmt.Println("Jawaban Salah")
			}
		} else if pilihan == 6 {
			fmt.Println("Anda memilih Keluar")
			if soal1benar && soal2benar && soal3benar && soal4benar && soal5benar {
				fmt.Println("Selamat Anda Lulus dengan Nilai 100 🎉🎉🎉")
			} else if soal1benar || soal2benar || soal3benar || soal4benar || soal5benar {
				if soal1benar {
					nilai += 20
				}
				if soal2benar {
					nilai += 20
				}
				if soal3benar {
					nilai += 20
				}
				if soal4benar {
					nilai += 20
				}
				if soal5benar {
					nilai += 20
				}
				if nilai >= 60 {
					fmt.Println("Selamat Anda Lulus dengan Nilai: ", nilai)
				} else {
					fmt.Println("Maaf Anda Tidak Lulus dengan Nilai: ", nilai)
				}
			}
			break
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}
