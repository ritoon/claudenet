package main

import "fmt"

type User struct {
	FistName string
	LastName string
}

type Animal struct {
	Name string
}

type LuggagesNumb int

type Alien struct {
	Name string
}

func main() {

	// Utiliser les structures pour initialiser des éléments
	var u User = User{FistName: "John", LastName: "Doe"}
	var a Animal = Animal{Name: "Dog"}
	var l LuggagesNumb = 3
	var al Alien = Alien{Name: "Zorg"}

	car := []interface{}{u, a, l, al} // insérer les éléments créé dans une variable car

	// boucler dans les éléments de car
	for _, v := range car {
		// avec le switch type résolvez le type d'origine
		switch obj := v.(type) {
		case User:
			fmt.Printf("User: %s %s\n", obj.FistName, obj.LastName)
		case Animal:
			fmt.Printf("Animal: %s\n", obj.Name)
		case LuggagesNumb:
			fmt.Printf("LuggagesNumb: %d\n", obj)
		default:
			fmt.Printf("Type inconnu %T , valeur: %v\n", obj, obj)
		}
	}
}
