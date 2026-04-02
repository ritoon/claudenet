# Interface : gérer l'inconnu

![Rick](img/strange.png)

Il arrive qu’il ne soit pas possible d’utiliser une structure définie avec des champs pour récupérer un payload. Par exemple, dans le cas d’un `PATCH`, le client n’est pas obligé d’envoyer tous les champs à mettre à jour. Dans d’autres situations, le payload peut contenir une partie dite _ouverte_, où certains éléments sont dynamiques.

Dans tous ces cas, il est possible d’utiliser une interface de valeur afin de binder le payload.

**Objectif :** utiliser une interface de valeur pour binder un payload non connu en implémentant la fonction `handlerCaracters`.

1. Récupérer Gin Gonic :

```sh
go get -u github.com/gin-gonic/gin
```

2. Dans la méthode `handlerCaracters`, appeler `getCaracter()`.

3. Initialiser une variable `payload` de type `any`.

4. Réaliser un bind du payload avec le package `encoding/json`.

5. Retourner le payload et accéder à l’adresse suivante dans votre navigateur :
   [http://localhost:8080/caracters](http://localhost:8080/caracters)
