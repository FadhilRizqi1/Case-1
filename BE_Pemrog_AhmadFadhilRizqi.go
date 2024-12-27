// by Ahmad Fadhil Rizqi

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// AlokasiMemori cek apakah n bisa dibentuk dari jumlah angka 3, 5, dan 7
// Kalau bisa, hitung jumlah angka yang dipakai, kalau tidak bisa return "TIDAK"
func AlokasiMemori(n int) string {
	// Kita coba semua kombinasi jumlah 3, 5, dan 7
	for a := 0; a <= n/3; a++ { // Loop untuk angka 3
		for b := 0; b <= n/5; b++ { // Loop untuk angka 5
			for c := 0; c <= n/7; c++ { // Loop untuk angka 7
				// Kalau totalnya pas sama n, return jumlah elemen
				if a*3+b*5+c*7 == n {
					return strconv.Itoa(a + b + c) // Konversi jumlah elemen jadi string
				}
			}
		}
	}
	// Kalau tidak ada kombinasi yang pas, return "TIDAK"
	return "TIDAK"
}

func main() {
	// Pakai bufio supaya input/output lebih cepat
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush() // Pastikan output ditulis sebelum program selesai

	// Baca jumlah test case (t)
	var t int
	fmt.Fscanln(reader, &t) // Baca input t

	// Loop untuk setiap test case
	for i := 0; i < t; i++ {
		var n int
		// Baca nilai n untuk test case ini
		fmt.Fscanln(reader, &n)
		// Panggil fungsi AlokasiMemori untuk mencari hasilnya
		result := AlokasiMemori(n)
		// Mencetak hasil
		fmt.Fprintln(writer, result)
	}
}
