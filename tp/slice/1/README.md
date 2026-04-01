# Slices : Observation des comportements

![crevette](img/crevette.png)

Les slices sont l’un des trois types de collection permettant de manipuler des tableaux sous-jacents. On les appelle aussi des tableaux dynamiques.

Contrairement aux tableaux, leur taille n’est pas définie dans leur type, ce qui les rend plus souples d’utilisation et évite de devoir réécrire du code lorsque les spécifications changent et qu’il faut modifier la taille d’un tableau dans plusieurs signatures de fonction.

Mais attention : avec ce comportement, les slices ont quelques subtilités. Dans certains cas, ils peuvent partager le même tableau sous-jacent ; dans d’autres, ils peuvent s’en détacher.

Afin de mieux comprendre les slices, veuillez exécuter le code suivant dans votre terminal :

**Objectifs :** comprendre comment interagir avec les slices

- Création et initialisation d’un slice
- Valeur contiguë en mémoire
- Capacité et longueur des slices
- Passage d’un slice à une fonction
- Lien avec un tableau sous-jacent
- Modification, ajout et suppression d’éléments
- Le slicing

```sh
go run .
```

**Aller plus loin** : utiliser le package de la lib standard [slices](https://pkg.go.dev/slices).
