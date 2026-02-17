package main

import "fmt"

// /
type we struct {
	are are
}

type are struct {
	the the
}

type the struct {
	best string
}

///

type hello struct {
	world string
}

///

type tech struct {
	academy string
}

type man struct {
	tech tech
}

type Level3 struct {
	man []man
}

type Level2 []Level3
type Level1 []Level2

type obj struct {
	str []Level1
}

func Slice() {

	we := we{
		are: are{
			the: the{
				best: "koda",
			},
		},
	}

	fmt.Println(we.are.the.best)

	//

	hello := hello{
		world: "hello world",
	}

	fmt.Println(hello.world)

	//

	obj := obj{
		str: []Level1{
			{}, {}, {},
			{
				{}, {
					{}, {}, {
						man: []man{
							{
								tech: tech{
									academy: "Koda Academy",
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(obj.str[3][1][2].man[0].tech.academy)

}
