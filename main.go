package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	user := User {
		FirstName: "Faris",
		LastName: "Adlin",
		PhoneNumber: "09851249124",
	}

	log.Println(user.FirstName, user.LastName, user.PhoneNumber, "BirthDate:", user.BirthDate)
}

// Uppercase letter means that is Global meanwhile lowecase letter that is Private meanwhile