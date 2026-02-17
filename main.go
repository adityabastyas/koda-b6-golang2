package main

import "fmt"

func main() {
	Slice()

	scores := []int{50, 75, 66, 20, 32, 90}

	posisi := 0

	for x := range scores {
		if scores[x] == 66 {
			posisi = x
		}
	}

	kiri := scores[:posisi+1]
	kanan := scores[posisi+1:]
	// fmt.Println(kanan)

	// scores = append(kiri, append(88) kanan...)
	scores = append(kiri, append([]int{88}, kanan...)...)

	fmt.Println(scores)
}
