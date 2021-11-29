package main

import "log"

func main() {
	type User struct {
		Firstname string
		Lastname  string
		Email     string
		Age       int
	}

	var users []User

	// var users []User
	users = append(users, User{"Ade", "Sunhaki", "ade.sunhaki@gmail.com", 30})

	log.Println(users)

	// firstLine := "Once upon a midnight dreary"

	// for animalType, animal := range firstLine {
	// 	log.Println(animalType, ":", animal)
	// }

	for _, l := range users {
		log.Println(l.Firstname, l.Lastname)
	}

}
