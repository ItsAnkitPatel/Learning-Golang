package main

// func add(a, b int) int {
// 	return a + b
// }

// for returning multiple values
func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {

	// result := add(4, 5)
	// fmt.Println("result", result)

	// fmt.Println(getLanguages())

	// lang1, lang2, lang3 := getLanguages()
	// lang1, lang2, _ := getLanguages() // if want to not use any value and to suppress golang error use '_'
	// fmt.Println(lang1)
	// fmt.Println(lang2)
	// fmt.Println(lang3)

	// fn := func(a int) int {
	// 	return 2
	// }
	// processIt(fn)

	fn := processIt()
	fn(6)

}
