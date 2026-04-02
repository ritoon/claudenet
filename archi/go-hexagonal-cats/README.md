# Architecture Hexagonale - API Chat

Un projet Go complet qui implémente l'**Architecture Hexagonale** (Ports & Adapters), le pattern architectural proposé par Alistair Cockburn. Ce projet démontre les principes d'une architecture maintenable, testable et flexible grâce à une séparation nette entre la logique métier et les détails d'implémentation.

## Table des matières

1. [Vue d'ensemble de l'Architecture Hexagonale](#vue-densemble-de-larchitecture-hexagonale)
2. [Concepts fondamentaux](#concepts-fondamentaux)
3. [Structure du projet](#structure-du-projet)
4. [Installation et exécution](#installation-et-exécution)
5. [API Endpoints](#api-endpoints)
6. [Exemples curl](#exemples-curl)
7. [Architecture du projet](#architecture-du-projet)
8. [Ajouter un nouvel adaptateur](#ajouter-un-nouvel-adaptateur)
9. [Avantages et inconvénients](#avantages-et-inconvénients)

---

## Vue d'ensemble de l'Architecture Hexagonale

### Qu'est-ce que c'est ?

L'**Architecture Hexagonale**, aussi appelée **Ports & Adapters**, est un pattern architectural créé par Alistair Cockburn qui résout le problème fondamental du développement logiciel : comment isoler la logique métier de l'infrastructure technique ?

Le concept clé est que votre **application est un hexagone isolé** au centre. Tous les échanges avec le monde extérieur (bases de données, HTTP, interfaces utilisateur, etc.) se font via des **Ports et Adapters** situés sur les faces de l'hexagone.

### Pourquoi une forme hexagonale ?

La forme hexagonale est purement symbolique. Elle représente :
- **L'application au centre** : la logique métier pure
- **Les six faces** : les différents ports d'entrée et de sortie
- **Les adaptateurs** : implémentations concrètes des ports

```
                    ┌─────────────────────────────────────┐
                    │                                     │
                    │         ADAPTERS INBOUND           │
                    │    (Driving / Client-side)         │
                    │                                     │
         ┌──────────┴─────────────────────────────────┬──┴────────┐
         │                                              │           │
         │     ┌──────────────────────────────┐        │           │
         │     │                              │        │           │
         │     │  HTTP Adapter  │ gRPC Adapter        │           │
         │     │  (REST API)    │ (RPC)               │           │
         │     │                │                      │           │
         │     │   PORT: CatService (Inbound)         │           │
         │     │     ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓         │           │
         │     │  ╔═══════════════════════════╗       │           │
         │     │  ║   APPLICATION CORE        ║       │           │
         │     │  ║  (Business Logic Layer)   ║       │           │
         │     │  ║                           ║       │           │
         │     │  ║  - Domain Entities        ║       │           │
         │     │  ║  - Application Service    ║       │           │
         │     │  ║  - Ports (Interfaces)     ║       │           │
         │     │  ╚═══════════════════════════╝       │           │
         │     │     ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓         │           │
         │     │   PORT: CatRepository (Outbound)     │           │
         │     │                │                      │           │
         │     │  Memory Repo   │ PostgreSQL │ Redis   │           │
         │     │  (In-Memory)   │ (SQL)      │ (Cache) │           │
         │     │                │                      │           │
         │     └──────────────────────────────────┬───┘           │
         │                                              │           │
         └──────────────────────────────────────────┬──┘           │
                    │                              │               │
                    │         ADAPTERS OUTBOUND   │               │
                    │    (Driven / Server-side)   │               │
                    │                              │               │
                    └──────────────────────────────┘               │
                                                                   │
                    MONDE EXTÉRIEUR                                │
           (HTTP, Bases de données, Services tiers)  ─────────────┘
```

---

## Concepts fondamentaux

### 1. **Ports (Interfaces)**

Un **Port** est une interface qui définit un contrat. Il y a deux types de ports :

#### **Port Inbound (Driving Port)**
- Définit ce que **l'application peut faire** (ses use cases)
- Exposé par la logique métier
- Implémenté par l'application
- Appelé par les adaptateurs inbound
- Exemple : `CatService` interface

```go
type CatService interface {
    CreateCat(name, breed, color string, age int) (*domain.Cat, error)
    GetCatByID(id string) (*domain.Cat, error)
    GetAllCats() ([]*domain.Cat, error)
    UpdateCat(id string, request domain.UpdateCatRequest) (*domain.Cat, error)
    DeleteCat(id string) error
}
```

#### **Port Outbound (Driven Port)**
- Définit ce que **l'application a besoin** (ses dépendances)
- Requis par la logique métier
- Implémenté par les adaptateurs outbound
- Utilisé par l'application
- Exemple : `CatRepository` interface

```go
type CatRepository interface {
    Save(cat *domain.Cat) error
    FindByID(id string) (*domain.Cat, error)
    FindAll() ([]*domain.Cat, error)
    Delete(id string) error
    Exists(id string) bool
}
```

### 2. **Adapters (Implémentations)**

Un **Adapter** est une implémentation concrète d'un Port. Il adapte le monde extérieur aux exigences de l'application.

#### **Adapter Inbound (Driving Adapter)**
- Reçoit les requêtes du monde extérieur (HTTP, CLI, gRPC)
- Appelle les inbound ports (CatService)
- Convertit les données externes en domaine
- Exemple : `HTTPHandler`

```go
type Handler struct {
    catService ports.CatService
}

func (h *Handler) CreateCat(w http.ResponseWriter, r *http.Request) {
    // Décoder la requête HTTP
    // Appeler h.catService.CreateCat()
    // Envoyer la réponse HTTP
}
```

#### **Adapter Outbound (Driven Adapter)**
- Implémente les outbound ports (CatRepository)
- Persiste les données dans le monde extérieur (DB, fichiers, cache)
- Convertit les entités de domaine en format de stockage
- Exemple : `MemoryRepository`, `PostgreSQLRepository`

```go
type CatRepository struct {
    mu   sync.RWMutex
    cats map[string]*domain.Cat
}

func (r *CatRepository) Save(cat *domain.Cat) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.cats[cat.ID] = cat
    return nil
}
```

### 3. **Flux de requête à travers l'hexagone**

```
[Client HTTP]
       ↓
   HTTP Request
       ↓
[HTTP Adapter - Handler.CreateCat()]
       ↓
   Décoder JSON
       ↓
[Inbound Port - CatService.CreateCat()]
       ↓
[Application Core - CatApplicationService.CreateCat()]
       ↓
   Valider métier (nom non vide, âge >= 0)
       ↓
   Générer un ID unique
       ↓
[Outbound Port - CatRepository.Save()]
       ↓
[Memory Adapter - CatRepository.Save()]
       ↓
   Stocker en mémoire
       ↓
[Retour à CatApplicationService]
       ↓
[Retour au Handler]
       ↓
   Encoder en JSON
       ↓
HTTP Response 201 Created
       ↓
[Client HTTP reçoit la réponse]
```

---

## Structure du projet

```
go-hexagonal-cats/
├── README.md                           # Cette documentation
├── go.mod                              # Module Go
├── main.go                             # Point d'entrée - Dependency Injection
│
└── internal/                           # Code interne (non exposé)
    │
    ├── core/                           # Le cœur de l'application
    │   │
    │   ├── domain/                     # Entités métier pures
    │   │   └── cat.go                  # Entité Cat et DTOs
    │   │
    │   ├── ports/                      # Contrats (Interfaces)
    │   │   ├── inbound.go              # Ports inbound (CatService)
    │   │   └── outbound.go             # Ports outbound (CatRepository)
    │   │
    │   └── service/                    # Logique métier
    │       └── cat_service.go          # Implémentation de CatService
    │                                   # (utilise CatRepository)
    │
    └── adapters/                       # Adaptateurs (Implémentations)
        │
        ├── inbound/                    # Côté driving (entrée)
        │   └── http/
        │       └── handler.go          # Adaptateur HTTP
        │                               # (implémente HTTP, utilise CatService)
        │
        └── outbound/                   # Côté driven (sortie)
            └── memory/
                └── cat_repository.go   # Adaptateur mémoire
                                        # (implémente CatRepository)
```

### Explication de la séparation des répertoires

| Répertoire | Responsabilité | Dépendances |
|------------|----------------|------------|
| `internal/core/domain/` | Entités métier pures, sans logique | Aucune |
| `internal/core/ports/` | Contrats abstraits | Seulement `domain/` |
| `internal/core/service/` | Logique métier, use cases | `domain/` et `ports/` |
| `internal/adapters/inbound/` | Interfaçage externe (HTTP, gRPC) | `ports/` |
| `internal/adapters/outbound/` | Implémentations de persistance | `domain/` et `ports/` |

### Règle d'or des dépendances

```
Domain
  ↑
Ports
  ↑
Service (utilise les ports)
  ↑
Adapters (implémentent ou utilisent les ports)
```

**Les dépendances vont TOUJOURS vers le haut.** La logique métier ne connaît jamais les adaptateurs.

---

## Installation et exécution

### Prérequis
- Go 1.21 ou supérieur

### Démarrer le serveur

```bash
# Se placer dans le répertoire du projet
cd go-hexagonal-cats

# Lancer le serveur
go run main.go
```

La sortie :
```
Cat API server starting on http://localhost:8080
API endpoints:
  POST   /api/cats          - Create a new cat
  GET    /api/cats          - Get all cats
  GET    /api/cats/:id      - Get a specific cat
  PUT    /api/cats/:id      - Update a cat
  DELETE /api/cats/:id      - Delete a cat
```

### Construire un binaire

```bash
go build -o cat-api main.go
./cat-api
```

---

## API Endpoints

### Modèle Chat (Cat)

```json
{
  "id": "a1b2c3d4",
  "name": "Minou",
  "breed": "Persan",
  "age": 3,
  "color": "Orange"
}
```

### Endpoints

| Méthode | URL | Description | Status |
|---------|-----|-------------|--------|
| `POST` | `/api/cats` | Créer un nouveau chat | 201 Created |
| `GET` | `/api/cats` | Récupérer tous les chats | 200 OK |
| `GET` | `/api/cats/:id` | Récupérer un chat par ID | 200 OK ou 404 |
| `PUT` | `/api/cats/:id` | Mettre à jour un chat | 200 OK ou 404 |
| `DELETE` | `/api/cats/:id` | Supprimer un chat | 204 No Content ou 404 |

---

## Exemples curl

### 1. Créer un nouveau chat

```bash
curl -X POST http://localhost:8080/api/cats \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Minou",
    "breed": "Persan",
    "age": 3,
    "color": "Orange"
  }'
```

Réponse (201 Created):
```json
{
  "id": "a1b2c3d4e5f6g7h8",
  "name": "Minou",
  "breed": "Persan",
  "age": 3,
  "color": "Orange"
}
```

### 2. Récupérer tous les chats

```bash
curl http://localhost:8080/api/cats
```

Réponse (200 OK):
```json
[
  {
    "id": "a1b2c3d4e5f6g7h8",
    "name": "Minou",
    "breed": "Persan",
    "age": 3,
    "color": "Orange"
  },
  {
    "id": "b2c3d4e5f6g7h8i9",
    "name": "Felix",
    "breed": "Siamois",
    "age": 5,
    "color": "Chocolat"
  }
]
```

### 3. Récupérer un chat spécifique

```bash
curl http://localhost:8080/api/cats/a1b2c3d4e5f6g7h8
```

Réponse (200 OK):
```json
{
  "id": "a1b2c3d4e5f6g7h8",
  "name": "Minou",
  "breed": "Persan",
  "age": 3,
  "color": "Orange"
}
```

Réponse si le chat n'existe pas (404 Not Found):
```json
{
  "error": "not found"
}
```

### 4. Mettre à jour un chat

```bash
curl -X PUT http://localhost:8080/api/cats/a1b2c3d4e5f6g7h8 \
  -H "Content-Type: application/json" \
  -d '{
    "age": 4,
    "color": "Gris"
  }'
```

Réponse (200 OK):
```json
{
  "id": "a1b2c3d4e5f6g7h8",
  "name": "Minou",
  "breed": "Persan",
  "age": 4,
  "color": "Gris"
}
```

### 5. Supprimer un chat

```bash
curl -X DELETE http://localhost:8080/api/cats/a1b2c3d4e5f6g7h8
```

Réponse (204 No Content - pas de body)

---

## Architecture du projet

### Le cœur non-contaminé (Core)

La partie `internal/core/` est complètement indépendante de tout framework, base de données ou technologie spécifique. Elle contient :

- **Domain** : Les entités métier (`Cat`)
- **Ports** : Les contrats abstraits
- **Service** : La logique métier qui dépend des ports

### Les adaptateurs remplaçables

Les adaptateurs permettent de **changer les détails techniques sans modifier le cœur** :

```go
// Exemple : Changer de repository
// De : In-Memory
catRepository := memoryAdapter.NewCatRepository()

// À : PostgreSQL (nouvel adapter outbound)
catRepository := postgresAdapter.NewCatRepository(dbConnection)

// Le service n'a besoin d'aucun changement !
catService := service.NewCatApplicationService(catRepository)
```

### Dependency Injection en main.go

```go
func main() {
    // 1. Créer l'adaptateur outbound (persistance)
    catRepository := memoryAdapter.NewCatRepository()

    // 2. Créer le service (logique métier)
    catService := service.NewCatApplicationService(catRepository)

    // 3. Créer l'adaptateur inbound (HTTP)
    httpHandler := httpAdapter.NewHandler(catService)

    // 4. Configurer les routes
    mux := http.NewServeMux()
    httpHandler.RegisterRoutes(mux)

    // 5. Démarrer le serveur
    http.ListenAndServe(":8080", mux)
}
```

---

## Ajouter un nouvel adaptateur

### Cas 1 : Ajouter une nouvelle API inbound (gRPC)

**Objectif** : Exposer la même logique métier via gRPC au lieu d'HTTP.

**Étapes** :

1. Créer le répertoire `internal/adapters/inbound/grpc/`

2. Créer un adaptateur gRPC qui implémente le port inbound :

```go
// internal/adapters/inbound/grpc/server.go
package grpc

import (
    "go-hexagonal-cats/internal/core/ports"
)

type GrpcServer struct {
    catService ports.CatService
}

func NewGrpcServer(catService ports.CatService) *GrpcServer {
    return &GrpcServer{catService: catService}
}

func (s *GrpcServer) CreateCat(ctx context.Context, req *pb.CreateCatRequest) (*pb.Cat, error) {
    // Appeler s.catService.CreateCat()
    // Convertir le résultat en pb.Cat
    // Retourner
}
```

3. Modifier `main.go` pour enregistrer le serveur gRPC :

```go
func main() {
    catRepository := memoryAdapter.NewCatRepository()
    catService := service.NewCatApplicationService(catRepository)

    // Adaptateur HTTP existant
    httpHandler := httpAdapter.NewHandler(catService)

    // Nouvel adaptateur gRPC
    grpcServer := grpcAdapter.NewGrpcServer(catService)

    // Démarrer les deux serveurs...
}
```

**Avantage** : Le service métier n'a pas besoin de changer du tout !

### Cas 2 : Ajouter une nouvelle persistance outbound (PostgreSQL)

**Objectif** : Remplacer le stockage en mémoire par une vraie base de données.

**Étapes** :

1. Créer le répertoire `internal/adapters/outbound/postgres/`

2. Créer un adaptateur qui implémente le port outbound :

```go
// internal/adapters/outbound/postgres/repository.go
package postgres

import (
    "database/sql"
    "go-hexagonal-cats/internal/core/domain"
)

type CatRepository struct {
    db *sql.DB
}

func NewCatRepository(db *sql.DB) *CatRepository {
    return &CatRepository{db: db}
}

func (r *CatRepository) Save(cat *domain.Cat) error {
    query := `INSERT INTO cats (id, name, breed, age, color)
             VALUES ($1, $2, $3, $4, $5)
             ON CONFLICT (id) DO UPDATE SET ...`
    _, err := r.db.Exec(query, cat.ID, cat.Name, cat.Breed, cat.Age, cat.Color)
    return err
}

func (r *CatRepository) FindByID(id string) (*domain.Cat, error) {
    // Exécuter une requête SELECT
}

// ... implémenter les autres méthodes
```

3. Modifier `main.go` pour utiliser la nouvelle implémentation :

```go
func main() {
    // À la place de :
    // catRepository := memoryAdapter.NewCatRepository()

    // Utiliser :
    db := connectToPostgres() // Ouvrir la connexion DB
    catRepository := postgresAdapter.NewCatRepository(db)

    // Le reste ne change pas !
    catService := service.NewCatApplicationService(catRepository)
    httpHandler := httpAdapter.NewHandler(catService)
    // ...
}
```

**Avantage** : Le service métier et le handler HTTP ne changent pas du tout !

---

## Avantages et inconvénients

### Avantages ✅

1. **Logique métier indépendante**
   - Le cœur de l'application ne dépend d'aucun framework ou technologie
   - Facile à tester sans mock complexe

2. **Adaptateurs interchangeables**
   - Changer de base de données sans refactoriser le service
   - Ajouter une nouvelle interface (HTTP, gRPC, WebSocket) sans modifier le métier

3. **Testabilité**
   - Tests unitaires purs du service (pas de dépendances externes)
   - Tests d'intégration seulement sur les adaptateurs

4. **Clarté architecturale**
   - Les dépendances sont explicites et visibles
   - Comprendre le flux de données est facile

5. **Maintenabilité**
   - Chaque couche a une responsabilité claire
   - Changements futurs localisés et contrôlés

### Inconvénients ❌

1. **Complexité initiale**
   - Plus de fichiers et d'interfaces au démarrage
   - Courbe d'apprentissage pour les débutants

2. **Boilerplate**
   - Code structuré peut paraître verbeux
   - Beaucoup de petites classes/interfaces

3. **Over-engineering possible**
   - Pour des projets très simples, peut être excessif
   - Faut pas créer des ports inutiles

4. **Dépendances circulaires à éviter**
   - Faut faire attention à la direction des dépendances
   - Erreur courante : adapter qui dépend du service

### Quand utiliser l'Architecture Hexagonale ?

**✅ Utilisez-la si** :
- Projet long terme avec évolution probable
- Équipe multiple
- Plusieurs interfaces prévues (HTTP + gRPC + CLI)
- Besoin de testabilité haute
- Changements de technologie possibles

**❌ N'utilisez-la pas si** :
- Prototype rapidement jeté
- Projet une seule personne et figé
- Logique métier très simple
- Deadline courte très strict

---

## Explication des fichiers clés

### `internal/core/domain/cat.go`

Contient les entités métier pures - sans dépendances externes.

```go
type Cat struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Breed string `json:"breed"`
    Age   int    `json:"age"`
    Color string `json:"color"`
}
```

### `internal/core/ports/inbound.go`

Définit ce que l'application **peut faire** - les contrats exposés.

```go
type CatService interface {
    CreateCat(name, breed, color string, age int) (*domain.Cat, error)
    GetCatByID(id string) (*domain.Cat, error)
    // ...
}
```

### `internal/core/ports/outbound.go`

Définit ce que l'application **a besoin** - les dépendances.

```go
type CatRepository interface {
    Save(cat *domain.Cat) error
    FindByID(id string) (*domain.Cat, error)
    // ...
}
```

### `internal/core/service/cat_service.go`

Implémente la logique métier - utilise les outbound ports, implémente les inbound ports.

```go
type CatApplicationService struct {
    repository ports.CatRepository
}

func (s *CatApplicationService) CreateCat(...) (*domain.Cat, error) {
    // Validation métier
    // Appel à s.repository.Save()
}
```

### `internal/adapters/inbound/http/handler.go`

Reçoit les requêtes HTTP, appelle les inbound ports.

```go
type Handler struct {
    catService ports.CatService
}

func (h *Handler) CreateCat(w http.ResponseWriter, r *http.Request) {
    // Décoder JSON
    // h.catService.CreateCat()
    // Encoder JSON et répondre
}
```

### `internal/adapters/outbound/memory/cat_repository.go`

Implémente le port outbound - stockage en mémoire.

```go
type CatRepository struct {
    mu   sync.RWMutex
    cats map[string]*domain.Cat
}

func (r *CatRepository) Save(cat *domain.Cat) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.cats[cat.ID] = cat
    return nil
}
```

### `main.go`

Wire tous les composants ensemble - injection de dépendances.

```go
func main() {
    // Créer les adaptateurs outbound
    catRepository := memoryAdapter.NewCatRepository()

    // Créer le service
    catService := service.NewCatApplicationService(catRepository)

    // Créer les adaptateurs inbound
    httpHandler := httpAdapter.NewHandler(catService)

    // Démarrer le serveur
}
```

---

## Conclusion

Ce projet démontre comment l'Architecture Hexagonale crée une **séparation claire entre la logique métier et les détails techniques**.

Le résultat : une application flexible, testable et maintenable où changer une technologie (DB, API, etc.) ne demande que d'ajouter ou modifier un adaptateur, sans toucher au cœur métier.

---

## Ressources

- **Alistair Cockburn** - Créateur du pattern Hexagonal Architecture
  - Article original : http://alistair.cockburn.us/Hexagonal+architecture

- **Clean Architecture** - Robert C. Martin (Uncle Bob)
  - Livre recommandé sur les principes architecturaux similaires

---

**Auteur** : Projet de démonstration
**Date** : 2026
**License** : MIT
