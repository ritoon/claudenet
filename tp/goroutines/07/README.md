## Les channels : pourquoi les fermer et comment ?

Le principe d’un **channel** est de fournir un moyen efficace de communication entre les goroutines. Précédemment, nous avons vu qu’il était possible d’y **envoyer des données** et de **synchroniser** des goroutines.

Voyez le channel non pas comme un simple **bâton de relais** entre deux coureurs, mais comme un **objet avec un état** :
s’il est **ouvert**, il travaille et laisse passer les messages ; à l’**inverse**, s’il est **fermé**, **plus aucun envoi n’est possible**.

Pour fermer un canal, on utilise la fonction interne de Go : `close(myChan)`.

> **Important :** après fermeture, **les envois (`chan <- x`) provoquent un panic**.

> Les **réceptions (`x := <-chan`) restent possibles** : on peut encore lire les valeurs **déjà présentes** (dans un canal bufferisé), puis les lectures renverront la **valeur zéro** du type et `ok = false`.

```go
ch := make(chan string)
close(ch)
ch <- "La porte entrouverte — un souffle frais du matin, le monde s’invite" // panic: send on closed channel
```

[play](https://go.dev/play/p/TLvcmYpccl4)

À la lecture, on peut savoir si un channel est encore “actif” grâce au second résultat :

```go
value, ok := <-ch
if !ok {
    // le channel est fermé et vidé
}
```

**Objectif :** en vérifiant `ok`, corrigez le code pour **éviter les fuites de goroutines** et prévenir tout **panic** dû à un envoi sur un channel fermé.
