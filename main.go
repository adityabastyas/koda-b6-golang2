package main

import "fmt"

func main() {
	Slice()

	scores := []int{50, 75, 20, 32, 66, 90}

	posisi := -1

	for x := range scores {
		if scores[x] == 66 {
			posisi = x
		}
	}

	if posisi != -1 {
		scores = append(scores[:posisi], append([]int{88}, scores[posisi:]...)...)
	}

	// kiri := scores[:posisi+1]
	// kanan := scores[posisi+1:]
	// fmt.Println(kanan)

	// scores = append(kiri, append(88) kanan...)
	// scores = append(kiri, append([]int{88}, kanan...)...)

	fmt.Println(scores)
}
