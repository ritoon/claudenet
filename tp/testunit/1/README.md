# Premiers pas

**Objectif :** se familiariser avec les tests unitaires et comprendre leurs limites.

Un test unitaire permet de valider le comportement du code dans différentes conditions d’utilisation, concurrentes ou non.

Dans ce premier TP, nous allons mettre en place une **table de tests** pour **chaque fonction exportée** du package `calculate`.

---

1. Création du fichier de test

Créez un fichier nommé `calculate_test.go` à côté de `calculate.go`.
La nomenclature est stricte : le suffixe `_test` exclut ces fichiers de la compilation mais sont exécutés dans le terminal avec la commande `go test .`

Dans le fichier de test, le nom du package est généralement le même que celui testé. Vous pouvez aussi utiliser le nom du package **suivi de** `_test` pour tester le package **comme un client externe**.

---

2. Écriture des fonctions de test

Créez **une fonction de test par fonction à tester**, par exemple :

```go
func TestAdd(t *testing.T) {
```

qui correspond à la fonction :

```go
func Add(a, b int) int {

```

---

3. Ajout d’une table de tests

Pour tester plusieurs jeux de paramètres, créez une **table de tests** à l'interieur de chaques fonction de test :

```go
data := []struct {
    testTitle string
    valueInA  int
    valueInB  int
    expected  int
}{
    {"a", 1, 1, 2},
    {"b", 1, 9, 10},
}
```

---

4. Boucler sur la table de tests

```go
for _, d := range data {
    got := Add(d.valueInA, d.valueInB)
    if got != d.expected {
        t.Errorf("test %q: got %d, expected %d", d.testTitle, got, d.expected)
    }
}
```

---

5. Pousser les limites avec des jeux de données

Ajoutez des **cas limites** pour tenter de casser les résultats attendus.
Vous aurez probablement besoin du package [math](https://pkg.go.dev/math#pkg-constants).

---

6. Lancer la commande de test

Allez dans votre terminal, déplacez-vous dans le dossier de package à tester et lancez la commande `go test .`, afin d'avoir plus de détails sur les opérations utilisez la commande `go test -v -cover .`, où le paramètres`-v` qui donne de la visibilité sur les opérations et le paramètres `-cover` correspond à l'affichage de la couverture de code.

Il est par ailleurs possible de cibler une fonction à tester avec la commande `go test -run TestAdd`.

---

7. Renforcer les fonctions

Pour rendre les fonctions plus robustes, ajoutez des **préconditions** et, si nécessaire, **modifiez les signatures** afin de retourner une erreur quand l’opération est dangereuse ou impossible.
