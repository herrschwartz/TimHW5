package main

import "fmt"

type sillystruct struct {
	key  int
	data string
}

func main() {
	var SoupRatings map[string]int
	SoupRatings = make(map[string]int)

	SoupRatings["Borscht"] = 10
	SoupRatings["Pho"] = 8
	SoupRatings["Sizzling Rice"] = 8
	SoupRatings["Tomato"] = 5
	SoupRatings["Chicken Noodle"] = 8
	SoupRatings["French Onion"] = 7
	SoupRatings["Ramen"] = 9

	StupidSlice := []float64{77.3, 65.8, 72.1, 68, 75.4}
	var average float64 = 0
	for i := 0; i < len(StupidSlice)-1; i++ {
		average += StupidSlice[i]
	}

	var silly sillystruct
	silly.data = "HEY YO!"
	silly.key = 1
	printItOut(silly)

	average = average / float64(len(StupidSlice)-1)

	var slicestring string = "Hello Person"
	newstring := slicestring[:5]

	fmt.Println(newstring)
	fmt.Println("The Temperature average is", average, "°F")
	fmt.Println("My rating for Pho is ", SoupRatings["Pho"])
}

func printItOut(s sillystruct) {
	fmt.Println(s.data)
}
