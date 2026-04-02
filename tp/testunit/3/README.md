# Couverture de code: fuzzing

Dans le dernier exercice nous avons vus qu'il n'était pas toujours évident d'écrire des jeux de données.
Quand des fonctions qui ont des conditions en fonctions de pluieurs états de variables passés en paramètres,
elles obligent à créer autant de cas qu'il y a de conditions.

[D'après cet article sur le fuzzing](https://go.dev/doc/security/fuzz/), essayon d'implémenter une fonction de fuzz.

```go
package complex

import (
	"testing"
)

func FuzzMyComplexFunction(f *testing.F) {
	f.Add(1, 2)
	f.Fuzz(func(t *testing.T, a int, b int) {
		MyComplexFunction(a, b)
	})
}
```

et lançons le test de fuzzing dans le terminal pendant 10 secondes:

```sh
go test -run=^$ -fuzz=Fuzz -fuzztime=10s
```

Ajouter un corpus afin de valider le maximum de cas possibles.
