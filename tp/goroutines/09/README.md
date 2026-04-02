## Le pouvoir du `select`

Nous avons vu dans le service précédent que tous les sous-services peuvent provoquer une fuite de mémoire lorsqu’ils sont utilisés en concurrence. Dans un projet, on ne choisit pas toujours l’**écosystème** des packages que l’on utilise, et il peut arriver qu’ils n’aient pas été conçus dès le départ pour fonctionner en mode concurrent.

Bonne nouvelle : on peut mettre en place des **garde-fous** grâce à l’utilisation de `select`. Ce **mot-clé** permet d’attendre plusieurs channels et de réagir à **celui qui est prêt**.

Exemple minimal :

```go
msg := make(chan string)
timeout := time.After(3 * time.Second)

go DoSomeThingGreat(msg)

select {
case s := <-msg:
    fmt.Println(s)
case <-timeout:
    fmt.Println("timeout")
    return
}
```

Ici, nous utilisons deux channels :

- `msg` transporte un message,
- `timeout` (via `time.After`) déclenche une temporisation unique de 3 secondes.

> Remarque : utilisez `time.NewTicker`/`ticker.C` pour des événements répétés, et `time.After` pour un **timeout** ponctuel (évite les fuites liées à `time.Tick`).

Dans ce TP, vous avez eu le temps de **refactoriser** le projet pour le rendre plus propre. Il contient deux packages :

- l’ancien service **API**,
- le **worker** qui lance les tâches d’appel.

**Objectif :** intégrez un `select` dans la fonction `Run` du package `worker` afin qu’elle **se termine après 5 secondes** (timeout propre).

[Pour aller plus loin](https://medium.com/@j.d.livni/write-a-go-worker-pool-in-15-minutes-c9b42f640923)
