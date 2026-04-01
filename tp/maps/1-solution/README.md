## Les Maps

![dico](img/dictionnaires.png)

Ce type de collection permet d'interagir rapidement entre clés et valeurs. Toutefois, ce n’est pas le meilleur choix si des contraintes strictes de temps de réponse existent, et l’espace mémoire préempté n’est pas contigu.

Quels sont les cas d’usage alors ?

- Permettre un usage simple et surtout une récupération des données par clé sans avoir à parcourir tout un tableau.
- Dédoublonner des paires clé/valeur.
- Communiquer avec du JSON sur des payloads comportant des champs non définis dans une structure.

**Objectif :** prise en main des maps avec un CRUD.

1. Récupérer le package pour la génération d’UUID :

```sh
go get -u github.com/google/uuid
```

2. Implémenter les méthodes CRUD.

Attention : il faut forcer l’allocation lorsqu’une map n’est pas encore initialisée.

```go
m := make(map[Tk]Tv)
```

3. Tester dans le terminal :

```sh
go run .
```
