package main

import (
	"container/ring"
	"fmt"
	"time"
)

func main() {
	data := ring.New(5)
	aktivitas := []string{
		"Membuka aplikasi",
		"Melihat profil pengguna",
		"Mengirim pesan",
		"Menerima notifikasi",
		"Melakukan pembayaran",
		"Mencari produk baru", // Ini akan menggantikan aktivitas pertama
	}

	fmt.Println("-- Menambahkan atktifiatas ke dalam ring ---")
	for _, akt := range aktivitas {
		now := time.Now().Format("2006-01-02 15:04:05")
		data.Value = fmt.Sprintf("[%s] %s", now, akt)

		fmt.Println(data.Value)

		data = data.Next()
	}

	fmt.Println("-- Riwayat 5 aktifitas terakhir ---")
	data.Do(func(value any) {
		fmt.Println(value)
	})
}
