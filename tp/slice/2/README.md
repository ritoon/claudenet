# Slice : CRUD (Create, Read, Update, Delete)

![maternité](img/maternite.png)

Afin de jouer avec le comportement des slices, nous partons d’un projet simple de CRUD permettant de manipuler une slice.

**Objectif :** renforcer ses connaissances sur les slices.

1. Implémenter les fonctions :

```go
Add(u *User) error
GetByID(id uint64) (*User, error)
UpdateByID(id uint64, u *User) error
DeleteByID(id uint64) error
```

2. Ajouter des appels de fonctions dans `main` et tester dans le terminal :

```sh
go run .
```
