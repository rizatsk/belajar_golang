package main

import "fmt"

func learnerSlice() {
	fmt.Println("Running learner slice")

	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println("daySlice1: ", daySlice1)
	fmt.Println("days: ", days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println("daySlice1: ", daySlice1)
	fmt.Println("daySlice2: ", daySlice2)
	fmt.Println("days: ", days)

	// Membuat slice dari awal
	newSlice := make([]string, 2, 3) // (typeData, panjang, kapasitas)
	newSlice[0] = "Rizat"
	newSlice[1] = "Sakmir"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Joko")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Jaka"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fmt.Println("Done learner slice")
}
