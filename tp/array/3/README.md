# Array dans le monde réel - partie 3

![pizza](img/pizza.png)

L’array, ou tableau, est un type de collection bas niveau permettant d’interagir avec des éléments finis et de récupérer une sous-partie d’un ensemble plus large.

Une fois que vous avez chargé un ensemble de données, il est possible d’en extraire une portion afin de l’analyser ou de la modifier.

Un cas d’usage serait, par exemple, de récupérer une image pour en transformer le contraste, la luminosité ou les couleurs. Mais prenons un exemple plus simple : l’identification du mimetype d’un document afin de le détecter.

**Objectif :** permettre de lire une petite partie d’un fichier et d’en retourner son mimetype.

**Étape 1 :** ajouter dans le `switch` deux cas permettant de récupérer le mimetype d’un PDF et d’un JPG.

```go
switch {
    case n >= 8 && [8]byte(buf[:8]) == pngSig:
        return "image/png", nil
    // Ajouter ici le test pour savoir si c'est un PDF ou une image JPG
}
return "application/octet-stream", nil
```

Les noms des mimetypes sont les suivants :

- `application/pdf`
- `image/jpeg`
