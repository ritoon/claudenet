# Atomic function

![Hello animation](img/atomic.png)

Au TP précédent, vous avez constaté que les goroutines lancés n'affichaient pas le résultat linéaire en fonction des appels et que l'orchestrateur priorisait lui même l'execution des fonctions appelés avec des goroutines, le résultat d'affichage n'est pas consistent dans les suite de nombres.

---

## A faire :

Ici, l'objectif est de prendre en main les fonctions atomics afin de réaliser un rendu linéaire dans l'iteration a l'affichage.

D'après la [documentation officiel](https://pkg.go.dev/sync/atomic#Int32.Add), modifiez le code afin que le terminal puisse afficher les work ids dans un ordre définit. Enfin, relancer dans le terminal votre programme.

1. Déclarez dans main une variable `iter` tel que :

```go
var iter atomic.Int32
```

2. Modifiez la signature de fonction HeavyWork afin de lui passer en paramètre l' atomic.Int32 déclaré dans main à la place de l'itérateur de la boucle for.

```go
func HeavyWork(workID *atomic.Int32, wg *sync.WaitGroup) {
```

3. Dans la fonction `HeavyWork` utilisez uniquement la fonction atomic suivante :

```go
(*Int32).Add()
```

4. Relanez votre programme dans le terminal.

---

## Qu'est ce que vous remarquez ?

Même si il semble y avoir un ordre plus logic, les IDs ne se suivent toujours pas, c'est encore l'orchestrateur de Go qui décide de prioriser les goroutines de l'une à l'autre.

Ce qui est mis en attente par l'orchestrateur n'est pas la fonction elle même mais les lignes de codes qu'elle contient.

Une fonction atomic permet seulement de garantir qu’une opération unique sur une variable (lecture, écriture, addition, comparaison-et-échange…) est indivisible et cohérente vis-à-vis des autres goroutines —
elle ne garantit pas l’ordre d’exécution global ni la cohérence de plusieurs champs ou étapes successives.

### Voici des cas d’usage concrets où sync/atomic est utile (et le type conseillé) :

- Compteur de requêtes/erreurs haute fréquence (metrics) — atomic.Int64

- Taille courante d’une file (traités au chargement) — atomic.Int64

- Générateur d’ID/sequence (index de ring-buffer, ticket) — atomic.Uint64

- Drapeau start/stop pour arrêter proprement des workers — atomic.Bool

- Feature flag (activer/désactiver une option à chaud) — atomic.Bool

- Circuit breaker : état open/half-open/closed codé en entier — atomic.Int32

- Limiter de débit : compteur de “tokens” consommés — atomic.Int64

- Max/Min observé (latence, taille, etc.) via boucle CAS — atomic.Int64

- Compteur de références (lifetime d’un objet partagé) — atomic.Int64

- Version/génération d’un cache pour invalider sans lock — atomic.Uint64

- Hot-swap de configuration “read-mostly” — atomic.Pointer[T] ou atomic.Value

- Pointer vers structure immuable (table de routage, règles) — atomic.Pointer[T]

- État d’un pool (nb connexions/worker actifs) — atomic.Int64

- Work stealing deques : indices tête/queue — atomic.Uint64

- Compteur sharded (N compteurs atomiques pour réduire la contention, agrégés périodiquement) — atomic.Int64 par shard

- Instrumentation légère (nb goroutines en cours sur une section) — atomic.Int64

- One-time init maison quand sync.Once est trop lourd/flexible — atomic.Bool + CAS

- Déduplication rapide (marquer “déjà vu” par bit/état entier) — atomic.Uint32

---

## Aller plus loin

- [Benchmarking sync/atomic](https://www.sambaiz.net/en/article/498/)
- [Atomic Operations and Mutexes in Go for Concurrency Control](https://medium.com/pickme-engineering-blog/atomic-operations-and-mutexes-in-go-for-concurrency-control-13d7ece91c7f)
- [Go Counter Benchmark](https://go-benchmarks.com/counter)
