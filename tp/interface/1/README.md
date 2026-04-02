# Interface de valeur : En voiture Simone

![Hello animation](img/voiture.png)

L’interface de valeur permet de contenir tout type de valeur. Par exemple, la signature de la fonction `fmt.Println(a ...any)`.

Le type `any` est un alias du type d’interface de valeur.

```go
type any = interface{}
```

**Objectif :** utiliser les interfaces de valeur dans toutes leurs formes.

1. Dans la fonction `main`, utiliser les structures définies pour créer des utilisateurs et des animaux.
2. Ajouter les éléments créés dans la variable `car`.
3. Créer une boucle pour parcourir la slice `car` et utiliser un _type switch_ pour afficher le type d’origine.

```go
switch elem.(type) {
case string:
    v, ok := elem.(string)
    if ok {
        fmt.Println("type détecté : string ->", v)
    }
default:
    fmt.Printf("type non répertorié %T - valeur : %v\n", elem, elem)
}
```
