package main

import ("github.com/ecastillochin/course")

func main() {
	Go := Course{
		Name:    "Go desde cero",
		Price:   20.1,
		IsFree:  false,
		UserIDs: []uint{2, 5, 12},
		Moduls: map[uint]string{
			1: "Introduccion",
			2: "Primeros pasos",
		},
	}

	/*	CSS := Course{
			"Css desde cero",
			14.05,
			false,
			[]uint{1, 2, 3, 4},
			map[uint]string{
				1: "primeros pasos",
				2: "ya eres un master",
			},
		}
	*/
	//fmt.Println(Go)
	Go.PrintCourse()
	Go.ChangePrice(40.1)
	//	fmt.Println(Go.Price)
	//	fmt.Println(CSS)
}
