package main

//Lab 01
func main() {
	// 	println("Hello World!")
	// 	println("Hello academy!")
	// 	println("Hello academy22!")
	// 	println("Hello academy24!")

	//Lab 02 - Strings
	// var myString string= "Hello World!"
	// 	println(myString)
	// 	myString = "1234"
	// 		println(myString)
	// 		myString = "Welcome to Go!"
	// 			println(myString)
	// 			myString = "5678"
	// 				 println(myString)

	//Lab 03 - Slices & Loops

	names := []string{"Daniel", "Rupert", "Emma"}
	colours := []string{"Red", "Blue", "Green", "Yellow"}
	for _, name := range names {
		
		println(name)

		for _, colour := range colours {
			println(colour)
		}
	}

}
