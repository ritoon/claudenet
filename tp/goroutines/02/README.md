# WaitGroup : apprendre à s'arrêter au bon moment

![Hello animation](img/bipbip-coyote.png)

Au TP précédent, vous avez constaté que les goroutines lancées n’affichaient pas le résultat. Le fil d’exécution étant détaché, l’application a quitté avant même qu’elles n’aient eu le temps de commencer à exécuter les méthodes du package `fmt`.

Afin d’y remédier, nous allons utiliser les _WaitGroups_, intégrés dans le package `sync`, qui permettent d’attendre que les actions soient terminées avant de quitter le programme.

D’après la [documentation officielle](https://pkg.go.dev/sync#example-WaitGroup), modifiez le code afin que le terminal puisse afficher les _work IDs_. Enfin, relancez votre programme dans le terminal.

Utilisez uniquement les fonctions suivantes :

```go
(*WaitGroup).Add()
(*WaitGroup).Done()
(*WaitGroup).Wait()
```

et en option

```go
(*WaitGroup).Go()
```

---

## A faire :

1. Déclarez dans `main` une variable `wg` telle que :

```go
var wg sync.WaitGroup
```

2. Ajoutez les fonctions `wg.Add()` et `wg.Done()` dans la fonction `main` en suivant la [documentation](https://pkg.go.dev/sync#example-WaitGroup-AddAndDone). Remarquez que vous pouvez aussi mettre en dehors juste avant la boucle `for` la méthode `wg.Add()` avec le `NBWork` en paramètre.

3. Modifiez la signature de la fonction `HeavyWork` afin de lui passer en paramètre le _WaitGroup_ déclaré dans `main` :

```go
func HeavyWork(workID int, wg *sync.WaitGroup) {
```

4. Ajoutez à la fin de la fonction `HeavyWork` l'appel à la méthode `(*WaitGroup).Done()` afin de signifier que la tâche est bien terminée. C'est ainsi que le compteur des waitgroupes peut se décrémenter.

5. Relancez votre programme dans le terminal.

### En bonus :

1. Vous pouvez jouer sur la valeur insérée dans `NBWork` et augmenter ou diminuer cette valeur. Vous constaterez qu’il y a toujours un ordre aléatoire dans l’affichage des valeurs de `i` dans le terminal.

2. Vous pouvez utiliser la fonction `(*WaitGroup).Go()` à la place du `(*WaitGroup).Add()` et du mot clée `go`.

3. Vou pouvez modifier l'appel du `wg.Done()` et le mettant au début de la fonction avec un `defer wg.Done()` afin qu'il soit toujours executé.

---

## Qu’est-ce que vous remarquez ?

Les IDs ne se suivent pas : c’est l’orchestrateur de Go qui décide de prioriser l’exécution des goroutines les unes par rapport aux autres.

---

## Aller plus loin

- [Gist of Go: Wait groups](https://antonz.org/go-concurrency/wait-groups/)
