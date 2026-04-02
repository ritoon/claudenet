## Les mutexes: la gestion des droits de passage

![Hello animation](img/custom.png)

Un autre outil très utilisé avec les goroutines est le mutex du package `sync`. Il permet d'executer une portion de code dans une seule goroutine à la fois; ainsi l'écriture lecture est sécurisée.

---

## A faire :

**Objectifs:**

- Donner l'information à l’orchestrateur d'executer qu'une goroutine à la fois sur une portion de code .
- Utiliser des [mutex](https://pkg.go.dev/sync#Mutex).

1. Dans votre fonction `main`, après la création du waitgroup, ajoutez la création d’un mutex :

```go
var mx sync.Mutex
```

2. passez la variable `mx` en paramètre de la fonction `Write` et `Read` par valeur pointeur.

3. Les nouvelles signatures des fonctions `Write` et `Read` devrait être :

```go
func Write(myUnSafeVariable *int, mx *sync.Mutex, wg *sync.WaitGroup) {
func Read(myUnSafeVariable *int, mx *sync.Mutex, wg *sync.WaitGroup) {
```

4. Dans les fonctions `Write` et `Read`, utilisez uniquement les fonctions de mutex suivantes au niveau de la zone à isoler :

```go
(m *Mutex) Lock()
... // your safe code
(m *Mutex) Unlock()
```

### En bonus :

1. Création d'un test unitaire permettant d'appeler le read et le write

2. Utiliser pour faire tourner le test la race condition tel que:

```sh
go test -v -race .
```

3. Supprimez les mutexes et relancez le test.

---

## Qu’est-ce que vous constatez ?

En conclusion, nous avons réussi à exécuter du code concurrent en toute sécurité grâce à l’utilisation des mutex, des WaitGroups et des mutexes.

Selon vous, existe-t-il un moyen plus simple d’obtenir le même résultat ?

---

## Aller plus loin

- [Go Synchronization Methods Benchmark](https://go-benchmarks.com/synchronization-methods)
- [Atomic vs Mutex](https://bbengfort.github.io/2022/11/atomic-vs-mutex/)
