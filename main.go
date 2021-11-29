package main

import "log"

func main() {
	myVar := "asga"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to cat")
	default:
		log.Println("Not found")

	}
}
