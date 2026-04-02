## Les channels : la communication inter-goroutines

![Hello animation](img/chan.png)

Afin de trouver une solution plus élégante, revenons au point de départ et allons découvrir l'utilisation des channels.

Pour crécap voici comment utiliser un channel

```go
// création d'un channel d'int
ch := make(chan int)
// Écrire dans un channel
go func (ch chan int){
	ch <- 1
}(ch)
// Lire la valeur
fmt.Println(<-ch)
// Le fermer
close(ch)
```

Remaquez qu'ici dans notre exemple l'utilisation de la fonction anonyme permet d'écrir dans le channel non bufferisé. Il est necessaire de se décrocher du fil d'execution sur ce type de channel.

Il est possible d'avoir des channel bufferisé avec à la création l'utilisation du paramètre optionnel de la fonction `make` tel que :

```go
// création d'un channel d'int bufferisé
ch := make(chan int, 2)
```

---

## A faire :

1. Dans la fonction `main`, sous la création du _WaitGroup_, ajoutez la création d’un canal de communication non bufferisé :

```go
ch := make(chan int)
```

2. Toujours dans la fonction `main ` à l'interieur de la boucle `for`, sous l’appel concurrent de la fonction `HeavyWork`, ajoutez :

```go
ch <- i
```

3. Modifiez la signature de la fonction `HeavyWork` afin d’avoir :

```go
func HeavyWork(workID chan int, wg *sync.WaitGroup) {
```

---

## Qu’est-ce que vous constatez ?

L’orchestrateur est bien impacté avec un ordre défini, mais l’application n’est plus exécutée en concurrence, elle s’exécute désormais de manière séquentielle.
