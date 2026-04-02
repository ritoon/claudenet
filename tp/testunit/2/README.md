# Couverture de code: génération d'un rapport

Afin de garantir une bonne couverture de code, il convient d'avoir des tests sur l'ensemble des fonctions critiques de son application. En d'autres termes, il n'est pas obligatoir d'avoir 100% de couverture de code.

Dans certaines fonctions qui comportent des conditions, il est possible que vous ométiez une valeur de donnée dans vos tables de tests. Pour y remédier, il est possible de générer un rapport de tests, permettant d'avoir une visibilité concrète sur l'effet des jeux de données.

1. Générez un premier rapport

```sh
go test -covermode=set -coverprofile=complex.txt
go tool cover -html=complex.txt -o complex.html
```

---

2. Ajoutez les tests

Testez des jeux de données afin d'essayer de rentrer dans les conditions et les sous conditions.

Attention, les tests doivent passer afin de pouvoir mettre à jour le document complex.html.

Que constatez-vous ?

**Il est souvent laborieux de créer des jeux de données pour des fonctions qui exisent déjà surtout quand elles sont complexes.**
