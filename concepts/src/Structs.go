package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func structs() {
	Devi := User{"Devi", "Devi@gmail.com", 25}
	fmt.Printf("%+v\n", Devi)
	fmt.Printf("%+v\n", Devi.Email)

	var name = new(User)
	name.Name = "Devi"
	name.Email = "Devi@gmail.com"
	name.Age = 20
	fmt.Println(name)

	var deva = &User{"Deva", "deva@gmail.com", 22}
	fmt.Println(deva)
}
