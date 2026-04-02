## Les channels : lecteur / rédacteur

Le programme présente bien une séquentialité : l’ordre d’exécution et d’affichage est respecté, mais le code **n’est plus concurrent**.

**Objectif :** trouver un moyen simple de corriger le code pour qu’il **s’exécute de manière concurrente**.

**Astuce :** une seconde boucle `for` peut être utilisée.

**Pour les experts :** vous pouvez affiner les signatures de fonctions utilisant des channels en les **typant selon leur usage** (lecture seule ou écriture seule).

Par exemple :

```go
func HeavyWork(workID chan int, wg *sync.WaitGroup) {
```

peut aussi s’écrire :

```go
func HeavyWork(workID <-chan int, wg *sync.WaitGroup) {
```

puisque l’on ne fait que lire depuis ce canal.
Cette notation n’est pas obligatoire, mais elle **améliore la lisibilité** et explicite l’intention (lecture vs écriture).

À l’inverse, pour une fonction qui **écrit** dans le channel :

```go
func WriteMessage(workID chan<- int, wg *sync.WaitGroup) {
```

**Conclusion :** les channels sont un outil puissant de **communication et de synchronisation** entre goroutines, facilitant l’échange de données et la conception de flux concurrents.
