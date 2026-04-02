# Concurrence : sortir du fil d’exécution avec l'orchestrateur

![Hello animation](img/octopus.png)

Ce premier TP a pour objectif de vous familiariser avec la concurrence en Go et l’utilisation du mot-clé `go`. Vous allez découvrir comment exécuter des fonctions de manière concurrente, gérer les routines et comprendre les avantages de la programmation concurrente dans Go.

---

## A faire :

1. Dans votre terminal, lancez une première fois le programme et constatez le résultat :

```sh
go run .
```

2. Modifiez à la ligne 12 l’appel de la fonction en ajoutant le mot clef `go`.

3. Relancez la commande `go run .` dans le terminal afin de voir la différence produite.

---

## Que constatez-vous ?

**Utiliser la concurrence permet de sortir du fil d’exécution principal de l’application et d’exécuter plusieurs processus simultanément.**

![Hello animation](img/hello.gif)

---

## Aller plus loin

- [Concurrency is not Parallelism by Rob Pike](https://www.youtube.com/watch?v=oV9rvDllKEg)
- [Visualizing Concurrency in Go](https://divan.dev/posts/go_concurrency_visualize/)
