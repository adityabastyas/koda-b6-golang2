package main

import "fmt"

type we struct {
	are  are
}

type are struct {
	the the
}

type the struct {
	 best string
}




func main(){

	we := we{
		are: are{
			the: the{
				best: "koda",
			},
		},
	}

	fmt.Println(we.are.the.best)

}
