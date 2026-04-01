## Slice : tableau sous-jacent

![coffre](img/coffre.png)

Afin de bien comprendre les effets d’un slice lorsqu’il change de tableau sous-jacent, nous allons afficher son pointeur grâce au package `unsafe`.

**Objectif :** assimiler le concept du tableau sous-jacent dans un slice.

1. Ajouter des éléments dans le slice `tbl1` afin qu’il modifie le pointeur de son tableau sous-jacent.

```go
tbl1 := append(tbl1, 1, 2, 3, 4)
displayInfos(tbl1, &tbl)
```

2. Créer une copie de `tbl1` vers un autre slice `tbl2`. La création de `tbl2` doit lui donner une capacité plus grande que `tbl1`, lui permettant de contenir 1 000 éléments.

```go
tbl2 := make([]string, len(tbl1), 1_000)
copy(tbl2, tbl1)
```

3. Tester dans votre terminal afin de vérifier que `tbl2` ne change pas de tableau sous-jacent lorsqu’on lui ajoute quelques éléments.

```sh
go run .
```
