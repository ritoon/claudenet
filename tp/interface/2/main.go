package main

import "fmt"

// Implémenter l'interface fmt.Stringer pour la structure User
// ne pas retirer cette ligne, elle est utilisée pour vérifier que User implémente bien fmt.Stringer
var _ fmt.Stringer = User{}

type User struct {
	FistName string
	LastName string
}

func (u User) String() string {
	return fmt.Sprintf("User: %s %s", u.FistName, u.LastName)
}

func main() {
	var u User = User{FistName: "John", LastName: "Doe"}
	fmt.Println(u) // fmt.Stringer est implémenté pour User, donc la méthode String() est utilisée
}
