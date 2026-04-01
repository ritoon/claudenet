package main

import "fmt"

func main() {

	printTitle("Création")

	// création d'une variable de type slice pour contenir des petits entier.
	var tbl1 []uint8
	fmt.Println("création de la variable tbl1:", tbl1)
	fmt.Printf("type du tbl1: %T\n", tbl1)

	printTitle("Création : Initialisation")

	// initialisation du Slice dans la variable tbl1 avec 10 éléments.
	tbl1 = make([]uint8, 10)
	fmt.Println("initialisation du tableau sous-jacent dans tbl1:", tbl1)

	// insersion de valeurs dans le slice tbl1
	for i := 0; i < len(tbl1); i++ {
		tbl1[i] = uint8(i)
	}
	fmt.Println("initialisation des valeurs contenus dans tbl1:", tbl1)

	printTitle("Création: déclaration litteral")
	tbl5 := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("déclaration litteral tbl5:", tbl5)

	printTitle("Capacité et longeur")

	// affiche la capacité et la longeur du slice.
	displayCapAndLen("tbl1", tbl1)

	// si on ajoute un élément il change de slice sous jacent.
	tbl1 = append(tbl1, 255)
	fmt.Println("ajout d'un élément dans le slice:", tbl1)

	// affiche la capacité et la longeur du slice après un ajout.
	displayCapAndLen("tbl1", tbl1)

	// évolution de la capacité d'une Slice sur n éléments.
	growingSliceWithAppend(5_000)

	printTitle("Valeurs pointeur: Espace contigue de mémoire")
	// boucle pour afficher la valeur pointeur des élément que contient le table.
	for i := 0; i < len(tbl1); i++ {
		// les valeurs des pointeurs se suivent -> l'espace alloué est contigue.
		fmt.Printf("valeur pointeur des éléments contenus dans tbl1[%d] %p\n", i, &tbl1[i])
	}

	printTitle("Valeurs pointeur: tableau sous-jacent et passage à une autre fonction")

	// Si j'affiche la valeur pointeur du Slice et que je le passe à une fonction,
	// il a le même slice sous-jacent.
	passSliceToAFunction("tbl1", tbl1)
	fmt.Printf("valeur pointeur du slice tbl1 :%p\n", &tbl1)

	printTitle("Suppression d'une valeur")

	// suppression d'un élément dans le Slice.
	fmt.Println("valeurs actuels du slice: \t\t\t", tbl1)
	displayCapAndLen("tbl1", tbl1)
	tbl1 = append(tbl1[:2], tbl1[3:]...)
	fmt.Println("suppression d'un élément dans le slice tbl1:\t", tbl1)
	displayCapAndLen("tbl1", tbl1)

	printTitle("Copier une slice vers une autre")

	// copier les valeurs d'une Slice à l'autre
	tbl2 := make([]uint8, len(tbl1))
	copy(tbl2, tbl1)
	fmt.Println("valeurs actuels du nouveau tbl1: \t\t", tbl1)
	fmt.Println("valeurs actuels du nouveau tbl2: \t\t", tbl2)

	// les deux slices ont des tableaus sous-jacent différent
	printTitle("Modification à partir d'une copie")
	tbl2[0] = 9
	fmt.Println("valeurs actuels du nouveau tbl1: \t\t", tbl1)
	fmt.Println("valeurs actuels du nouveau tbl2: \t\t", tbl2)

	printTitle("Slicing")

	// récupérer une tranche d'une Slice vers un autre
	tbl3 := tbl1[3:6]
	fmt.Println("valeurs actuels du nouveau tbl3: \t\t", tbl3)
	displayCapAndLen("tbl1", tbl1)
	displayCapAndLen("tbl3", tbl3)

	printTitle("Slicing: nouvelle capacité")
	// récupérer une tranche d'une Slice vers un autre en modifiant la longueur.
	tbl4 := tbl1[3:6:6]
	fmt.Println("valeurs actuels du nouveau slice 4: \t\t", tbl4)
	displayCapAndLen("tbl1", tbl1)
	displayCapAndLen("tbl4", tbl4)

	printTitle("Slicing: partage du tableau sous-jacent")

	// le Slice 1, 3, 4 restent connecté si une valeur est modifiée dans l'un des Slices.
	// ces slices font référence au même tableau sous-jacent.
	// le slice 2 lui est déconnecté et n'a pas le même tableau sous-jacent:
	tbl1[3] = 200
	fmt.Printf("valeur du slice 1 :%v\n", tbl1[3])
	fmt.Printf("valeur du slice 3 :%v\n", tbl3[0])
	fmt.Printf("valeur du slice 4 :%v\n", tbl4[0])

	printTitle("Slicing: ajout de valeurs => déconnexion du tableau sous-jacent")

	// Si on ajoute une valeur à un slice provenant d'un autre slice
	// ils ne font plus référence au même tableau sous-jacent.
	displayCapAndLen("tbl1", tbl1)
	displayCapAndLen("tbl4", tbl4)
	fmt.Println("ajout d'une valeur dans le tbl4")
	tbl4 = append(tbl4, 44)
	displayCapAndLen("tbl1", tbl1)
	displayCapAndLen("tbl4", tbl4)
	fmt.Printf("valeur du slice 1 :%v\n", tbl1[3:6])
	fmt.Printf("valeur du slice 4 :%v\n", tbl4)
	tbl4[0] = 9
	fmt.Println("modification d'une valeur dans tbl4")
	fmt.Printf("valeur du slice 1 :%v\n", tbl1[3:6])
	fmt.Printf("valeur du slice 4 :%v\n", tbl4)
}

func passSliceToAFunction(tblName string, tbl []uint8) {
	for i := 0; i < len(tbl); i++ {
		fmt.Printf("valeur pointeur des éléments contenus dans %v[%v] : %p (dans la fonction)\n", tblName, i, &tbl[i])
	}
	fmt.Printf("valeur pointeur du slice %v : %p (dans une fonction)\n", tblName, &tbl)
}

func displayCapAndLen(tblName string, tbl []uint8) {
	fmt.Printf("le %v a une capacité %d et une longueur %d\n", tblName, cap(tbl), len(tbl))
}

func growingSliceWithAppend(n int) {
	printTitle("Capacité et longeur: Calcul de nouvelles capacités")
	tbl := make([]int, 0)
	capacity := cap(tbl)
	fmt.Println("capacité initial du slice:", capacity)
	for i := 0; i < n; i++ {
		tbl = append(tbl, i)
		newCapacity := cap(tbl)
		if newCapacity != capacity {
			if capacity != 0 {
				fmt.Println("nouvelle capacité du slice:", newCapacity, "\t +", (newCapacity*100)/capacity, "%")
			} else {
				fmt.Println("nouvelle capacité du slice:", newCapacity)
			}
			capacity = newCapacity
		}
	}
}

func printTitle(msg string) {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("=>", msg)
}
