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

//

type fruit struct {
	is string
}

type item struct {
	fruit fruit
}

type favourite struct {
	favourite []item
}

type my []favourite

//

type num struct {
	first  [2]int
	second [3]int
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
									academy: "Tech Academy",
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(obj.str[3][1][2].man[0].tech.academy)

	//

	my := my{
		{
			favourite: []item{
				{},
				{},
				{},
				{
					fruit: fruit{
						is: "Apple",
					},
				},
			},
		},
	}

	fmt.Println(my[0].favourite[3].fruit.is)

	//

	num := num{
		first: [2]int{
			2, 15,
		},
		second: [3]int{
			8, 90, 17,
		},
	}

	fmt.Println(num.first[1] + num.second[2])

}
