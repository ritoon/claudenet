# Array dans le monde réel - partie 2

![machine](img/machine.png)

L’array, ou tableau, est un type de collection bas niveau permettant d’interagir avec des éléments finis. Cela permet notamment de l’utiliser pour créer un buffer afin de lire des portions d’un document ou d’un flux de données transmis sur le réseau.

Afin de lire un document, plusieurs options s’offrent à vous :

- Si ce document est léger, rien n’empêche de le lire d’une traite et de le transmettre à un writer.
- Imaginons maintenant que ce document soit plus lourd : dans ce cas, il convient de le lire en sous-parties afin d’éviter de saturer la RAM.

**Objectif :** jouer avec la taille du buffer afin d’optimiser le nombre d’opérations par nanoseconde et ainsi permettre un maximum de lectures.

**Étape 1 :** exécuter les tests de benchmark dans votre terminal

```sh
go test -bench=.
```

**Étape 2 :** dans le buffer, remplacer la valeur 16 par d’autres valeurs plus grandes.

```go
// ligne 22
var buf [16 << 10]byte // tampon fixe sur la pile, 16 kio
```

**Étape 3 :** recommencer les étapes 1 et 2 jusqu’à obtenir une valeur plafond.

## Que constatez-vous ?

Le nombre d’opérations par seconde est plus important avec un buffer bien adapté. À l’inverse, il peut aussi se révéler néfaste s’il est trop petit ou trop grand.

## Aller plus loin

- [Bufio de la lib standard](https://pkg.go.dev/bufio@go1.25.4)
- [Streaming without a buffer](https://medium.com/stupid-gopher-tricks/streaming-data-in-go-without-buffering-3285ddd2a1e5)
