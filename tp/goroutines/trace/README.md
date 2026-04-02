# Serveur trace

Ici, nous allons voir comment utiliser le nouveau package `runtime/trace`. Il a été introduit dans la version 1.25 afin de faciliter la visualisation du runtime.

Dans votre terminal, lancez l’application.

**Objectif :** utiliser le package `runtime/trace`

```sh
go run .
```

Allez ensuite sur l’URL :
[http://localhost:8080/](http://localhost:8080/)

Puis lancez :

```sh
go tool trace trace.out
```
