# Go Clean Cats - API CRUD Clean Architecture

Un exemple prêt pour la production d'une API CRUD de gestion de chats construite selon les principes de Clean Architecture (architecture d'Uncle Bob) en Go en utilisant uniquement la bibliothèque standard.

## Table des matières

1. [Aperçu](#aperçu)
2. [Clean Architecture expliquée](#clean-architecture-expliquée)
3. [Structure du projet](#structure-du-projet)
4. [Pour commencer](#pour-commencer)
5. [Points d'accès de l'API](#points-daccès-de-lapi)
6. [Exemples](#exemples)
7. [Avantages et compromis de l'architecture](#avantages-et-compromis-de-larchitecture)

---

## Aperçu

Ce projet démontre comment construire une application Go évolutive, testable et maintenable en utilisant les principes de Clean Architecture. L'API gère une collection de chats avec des opérations CRUD complètes, les stockant en mémoire.

**Caractéristiques principales :**
- Clean Architecture avec séparation stricte des couches
- Implémentation du Dependency Inversion Principle (DIP)
- Zéro dépendance externe (bibliothèque standard uniquement)
- Persistance en mémoire avec opérations thread-safe
- Codes d'état HTTP appropriés et gestion des erreurs
- Code entièrement commenté expliquant les décisions architecturales

---

## Clean Architecture expliquée

### Qu'est-ce que Clean Architecture ?

Clean Architecture est un motif architectural introduit par Uncle Bob (Robert C. Martin) qui met l'accent sur la séparation des responsabilités en couches distinctes. L'objectif principal est de créer des systèmes qui sont :

- **Indépendants des frameworks** - La logique métier ne dépend pas de HTTP, des bases de données ou d'aucun framework
- **Testables** - Les règles métier principales peuvent être testées sans dépendances externes
- **Indépendants de l'interface utilisateur** - La logique métier fonctionne indépendamment de sa présentation
- **Indépendants de la base de données** - Vous pouvez changer de base de données sans modifier la logique métier
- **Indépendants des agents externes** - Les services externes peuvent être remplacés

### Les cercles concentriques

Clean Architecture est souvent visualisée comme des cercles concentriques, où chaque cercle représente une couche d'abstraction :

```
┌─────────────────────────────────────────────────────────────────┐
│                    Frameworks & Tools                            │
│                  (HTTP, Bases de données)                        │
├─────────────────────────────────────────────────────────────────┤
│                  Couche Infrastructure                           │
│         (Routeurs, Adaptateurs de BD, Présentateurs)           │
├─────────────────────────────────────────────────────────────────┤
│                     Couche Adapter                               │
│           (Contrôleurs, Passerelles, Présentateurs)            │
├─────────────────────────────────────────────────────────────────┤
│                  Couche Application                              │
│                  (Use Cases / Interacteurs)                      │
├─────────────────────────────────────────────────────────────────┤
│                     Couche Entity                                │
│                (Règles métier, Entities)                         │
│                    (Pas de dépendances)                          │
└─────────────────────────────────────────────────────────────────┘
```

### La Dependency Rule

La règle la plus importante de Clean Architecture est :

**« Les dépendances du code source doivent pointer vers l'intérieur, jamais vers l'extérieur. »**

Cela signifie :
- Les couches externes dépendent des couches internes
- Les couches internes ne dépendent jamais des couches externes
- La communication s'effectue vers l'intérieur à travers les interfaces/abstractions

```
                    Couche Entity
                         ↑
                         │ dépend de
                         │
                    Couche Use Case
                         ↑
                         │ dépend de
                         │
                    Couche Adapter
                         ↑
                         │ dépend de
                         │
                   Couche Infrastructure
```

### Les quatre couches de ce projet

#### 1. Couche Domain (Entities) - Cercle interne

**Localisation :** `domain/`

**Caractéristiques :**
- Contient les entities métier et les règles
- Aucune dépendance sur quoi que ce soit d'autre
- Ne connaît rien des bases de données, HTTP ou frameworks
- Logique métier pure

**Fichiers :**
- `domain/entity/cat.go` - L'entity Cat avec les attributs métier essentiels
- `domain/repository/cat_repository.go` - Interface Repository que les use cases utilisent

**Exemple - L'entity Cat :**
```go
type Cat struct {
    ID    string
    Name  string
    Breed string
    Age   int
    Color string
}
```

L'entity contient uniquement ce qui importe pour le métier. Elle ne sait rien des balises JSON, des colonnes de base de données ou des réponses API.

#### 2. Couche Use Case (Règles métier de l'application)

**Localisation :** `usecase/`

**Caractéristiques :**
- Contient les règles métier spécifiques à l'application
- Orchestre les entities et Repository
- Définit les interfaces pour les dépendances externes (ne dépend pas d'elles)
- Indépendant de tout framework ou base de données

**Fichiers :**
- `usecase/cat_usecase.go` - Orchestration de la logique métier

**Motif de conception clé :**
```go
type CatUseCase struct {
    catRepository repository.CatRepository  // Dépend de l'interface
}
```

Le use case reçoit l'interface Repository via l'injection de dépendances. Il ne crée pas le Repository et ne sait pas comment il est implémenté.

**Exemple - Logique métier :**
```go
func (uc *CatUseCase) CreateCat(id, name, breed string, age int, color string) (*entity.Cat, error) {
    if name == "" || breed == "" || color == "" {
        return nil, errors.New("name, breed, and color are required")
    }
    if age < 0 || age > 120 {
        return nil, errors.New("age must be between 0 and 120")
    }
    // ... crée le chat via le Repository
}
```

La logique de validation est indépendante de whether you're using HTTP, gRPC, or CLI interfaces.

#### 3. Couche Adapter (Adaptateurs d'interfaces)

**Localisation :** `adapter/`

**Caractéristiques :**
- Convertit les données entre les use cases et les formats externes
- Les contrôleurs gèrent les requêtes HTTP
- Les présentateurs formatent les réponses
- Ne contient pas de logique métier
- Crée un pont entre les couches externes et internes

**Fichiers :**
- `adapter/controller/cat_controller.go` - Gestionnaires de requêtes HTTP
- `adapter/presenter/cat_presenter.go` - Formateurs de réponses

**Exemple - Contrôleur (adaptateur d'entrée) :**
```go
func (cc *CatController) HandleCreateCat(w http.ResponseWriter, r *http.Request) {
    var req CreateCatRequest
    json.NewDecoder(r.Body).Decode(&req)  // Décode la requête HTTP

    cat, err := cc.catUseCase.CreateCat(...)  // Appelle le use case
    // ...
}
```

**Exemple - Présentateur (adaptateur de sortie) :**
```go
type CatResponse struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    // ...
}

func (p *CatPresenter) PresentCat(cat *entity.Cat) CatResponse {
    return CatResponse{
        ID:    cat.ID,
        Name:  cat.Name,
        // ...
    }
}
```

Remarquez que l'entity n'a pas de balises JSON - c'est le travail du présentateur !

#### 4. Couche Infrastructure (Frameworks & Drivers) - Cercle externe

**Localisation :** `infrastructure/`

**Caractéristiques :**
- Implémente les interfaces définies dans les couches internes
- Contient du code spécifique au framework
- Routage HTTP, pilotes de base de données, I/O de fichier
- Connaît tous les détails du fonctionnement

**Fichiers :**
- `infrastructure/router/router.go` - Configuration du routage HTTP
- `infrastructure/persistence/memory_cat_repository.go` - Implémentation du Repository

**Exemple - Implémentation du Repository :**
```go
type MemoryCatRepository struct {
    cats map[string]*entity.Cat
    mu   sync.RWMutex
}

func (r *MemoryCatRepository) Create(cat *entity.Cat) (*entity.Cat, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    if _, exists := r.cats[cat.ID]; exists {
        return nil, errors.New("cat with this ID already exists")
    }

    r.cats[cat.ID] = cat
    return cat, nil
}
```

Ceci implémente l'interface `CatRepository` définie dans la couche domain. Vous pourriez échanger ceci avec une implémentation PostgreSQL sans rien changer dans les use cases !

---

## Structure du projet

```
go-clean-cats/
├── README.md                                    # Ce fichier
├── go.mod                                       # Définition du module Go
├── main.go                                      # Point d'entrée de l'application
│
├── domain/                                      # Couche 1 : Règles métier (Entities)
│   ├── entity/
│   │   └── cat.go                              # Entity Cat - objet métier essentiel
│   └── repository/
│       └── cat_repository.go                   # Interface pour la persistance des données
│
├── usecase/                                     # Couche 2 : Règles métier de l'application
│   └── cat_usecase.go                          # Orchestre les entities et Repository
│
├── adapter/                                     # Couche 3 : Adaptateurs d'interfaces
│   ├── controller/
│   │   └── cat_controller.go                   # Gestionnaires de requêtes HTTP
│   └── presenter/
│       └── cat_presenter.go                    # Formateurs de réponses
│
└── infrastructure/                              # Couche 4 : Frameworks & Drivers
    ├── router/
    │   └── router.go                           # Configuration du routage HTTP
    └── persistence/
        └── memory_cat_repository.go            # Implémentation du Repository en mémoire
```

### Diagramme du flux de dépendances

```
main.go (Composition Root)
    ↓
Crée les instances dans l'ordre :
    1. MemoryCatRepository (infrastructure/persistence)
    2. CatUseCase (usecase) - reçoit le Repository
    3. CatPresenter (adapter/presenter)
    4. CatController (adapter/controller) - reçoit le use case et le présentateur
    5. Router (infrastructure/router) - reçoit le contrôleur

Le flux montre que les dépendances s'écoulent vers l'intérieur :
├── Infrastructure crée les implémentations
├── Adapters utilisent ces implémentations
├── Use Cases dépendent des interfaces du domain
└── Domain n'a pas de dépendances
```

---

## Pour commencer

### Prérequis
- Go 1.21 ou supérieur

### Exécuter le serveur

1. Accédez au répertoire du projet :
```bash
cd go-clean-cats
```

2. Exécutez le serveur :
```bash
go run main.go
```

Vous devriez voir :
```
Starting Cat CRUD API server on :8080
Available endpoints:
  POST   /cats              - Create a new cat
  GET    /cats              - Get all cats
  GET    /cats/{id}         - Get a specific cat
  PUT    /cats/{id}         - Update a cat
  DELETE /cats/{id}         - Delete a cat
```

3. Le serveur écoute maintenant sur `http://localhost:8080`

---

## Points d'accès de l'API

### 1. Créer un chat
**Point d'accès :** `POST /cats`

**Corps de la requête :**
```json
{
  "id": "cat-001",
  "name": "Whiskers",
  "breed": "Persian",
  "age": 3,
  "color": "Orange"
}
```

**Réponse en cas de succès :** `201 Created`
```json
{
  "id": "cat-001",
  "name": "Whiskers",
  "breed": "Persian",
  "age": 3,
  "color": "Orange"
}
```

**Réponse en cas d'erreur :** `400 Bad Request`
```
Invalid request body
name, breed, and color are required
age must be between 0 and 120
```

---

### 2. Obtenir tous les chats
**Point d'accès :** `GET /cats`

**Réponse en cas de succès :** `200 OK`
```json
[
  {
    "id": "cat-001",
    "name": "Whiskers",
    "breed": "Persian",
    "age": 3,
    "color": "Orange"
  },
  {
    "id": "cat-002",
    "name": "Mittens",
    "breed": "Siamese",
    "age": 5,
    "color": "Cream"
  }
]
```

---

### 3. Obtenir un chat spécifique
**Point d'accès :** `GET /cats/{id}`

**Exemple :** `GET /cats/cat-001`

**Réponse en cas de succès :** `200 OK`
```json
{
  "id": "cat-001",
  "name": "Whiskers",
  "breed": "Persian",
  "age": 3,
  "color": "Orange"
}
```

**Réponse en cas d'erreur :** `404 Not Found`
```
Cat not found
```

---

### 4. Mettre à jour un chat
**Point d'accès :** `PUT /cats/{id}`

**Exemple :** `PUT /cats/cat-001`

**Corps de la requête :**
```json
{
  "name": "Whiskers Jr.",
  "breed": "Persian",
  "age": 4,
  "color": "Orange"
}
```

**Réponse en cas de succès :** `200 OK`
```json
{
  "id": "cat-001",
  "name": "Whiskers Jr.",
  "breed": "Persian",
  "age": 4,
  "color": "Orange"
}
```

---

### 5. Supprimer un chat
**Point d'accès :** `DELETE /cats/{id}`

**Exemple :** `DELETE /cats/cat-001`

**Réponse en cas de succès :** `204 No Content`

**Réponse en cas d'erreur :** `400 Bad Request`
```
cat not found
```

---

## Exemples

### Utilisation de cURL

#### Créer un chat :
```bash
curl -X POST http://localhost:8080/cats \
  -H "Content-Type: application/json" \
  -d '{
    "id": "cat-001",
    "name": "Whiskers",
    "breed": "Persian",
    "age": 3,
    "color": "Orange"
  }'
```

#### Obtenir tous les chats :
```bash
curl http://localhost:8080/cats
```

#### Obtenir un chat spécifique :
```bash
curl http://localhost:8080/cats/cat-001
```

#### Mettre à jour un chat :
```bash
curl -X PUT http://localhost:8080/cats/cat-001 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Whiskers Jr.",
    "breed": "Persian",
    "age": 4,
    "color": "Orange"
  }'
```

#### Supprimer un chat :
```bash
curl -X DELETE http://localhost:8080/cats/cat-001
```

### Utilisation du client HTTP de Go

```go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func main() {
    // Create a cat
    cat := map[string]interface{}{
        "id":    "cat-001",
        "name":  "Whiskers",
        "breed": "Persian",
        "age":   3,
        "color": "Orange",
    }

    body, _ := json.Marshal(cat)
    resp, _ := http.Post("http://localhost:8080/cats",
        "application/json",
        bytes.NewBuffer(body))
    defer resp.Body.Close()

    // Response is in resp.Body
}
```

---

## Avantages et compromis de l'architecture

### Avantages de Clean Architecture

#### 1. **Testabilité**
La logique métier est complètement séparée des frameworks. Vous pouvez tester les use cases sans HTTP, bases de données ou dépendances externes.

```go
// Facile à tester - pas besoin de HTTP ou de base de données
func TestCreateCat(t *testing.T) {
    repo := &MockCatRepository{}
    useCase := usecase.NewCatUseCase(repo)

    cat, err := useCase.CreateCat("id1", "Whiskers", "Persian", 3, "Orange")
    assert.NoError(t, err)
    assert.Equal(t, "Whiskers", cat.Name)
}
```

#### 2. **Indépendance du framework**
Passez de HTTP à gRPC, de la mémoire à PostgreSQL, sans toucher à la logique métier.

```go
// Même use case, framework différent
catUseCase := usecase.NewCatUseCase(postgresRepository)
catUseCase = usecase.NewCatUseCase(mongoRepository)
catUseCase = usecase.NewCatUseCase(memoryRepository)

// Ou interface différente
grpcService := grpc.NewCatService(catUseCase)
restAPI := rest.NewCatAPI(catUseCase)
```

#### 3. **Maintenabilité**
La séparation claire des responsabilités rend le code plus facile à comprendre, modifier et étendre.

#### 4. **Scalabilité**
Facile d'ajouter de nouvelles fonctionnalités sans affecter le code existant. De nouveaux Repository, contrôleurs ou présentateurs peuvent être ajoutés indépendamment.

#### 5. **Inversion des dépendances**
Les modules de haut niveau ne dépendent pas des modules de bas niveau. Les deux dépendent des abstractions.

```
✓ Bon :   UseCase dépend de l'interface Repository
✗ Mauvais : UseCase dépend de MemoryCatRepository
```

### Compromis et inconvénients

#### 1. **Complexité**
Plus de couches signifie plus de code et plus de fichiers à comprendre initialement.

```
CRUD simple sans Clean Architecture :
├── main.go
└── db.go

Avec Clean Architecture :
├── domain/entity/
├── domain/repository/
├── usecase/
├── adapter/controller/
├── adapter/presenter/
├── infrastructure/router/
└── infrastructure/persistence/
```

#### 2. **Sur-ingénierie pour les petits projets**
Pour un script simple ou un petit projet, Clean Architecture pourrait être trop complexe.

**Quand Clean Architecture vaut le coût :**
- Taille d'équipe > 1 personne
- Longévité du projet > 3 mois
- Changements de requirements attendus
- Plusieurs interfaces nécessaires (HTTP, CLI, gRPC)

**Quand les approches plus simples suffisent :**
- Script monofichier
- Outil interne aux requirements fixes
- Prototype ou preuve de concept

#### 3. **Surcharge de performance**
Plus d'appels de fonction et d'indirections, mais négligeable dans la plupart des cas.

#### 4. **Courbe d'apprentissage plus raide**
Les développeurs nouveaux à Clean Architecture ont besoin de temps pour comprendre les couches et le flux des dépendances.

### Quand utiliser Clean Architecture

**Bons cas d'usage :**
- API web en production
- Applications longue durée avec changing requirements
- Grandes équipes avec différentes personnes travaillant sur différentes couches
- Applications qui ont besoin de plusieurs interfaces (web, CLI, mobile API)
- Applications qui ont besoin de swapper les implémentations (base de données, processeur de paiement)

**Trop complexe pour :**
- Scripts rapides
- Outils ponctuels
- Prototypes
- Applications stables avec une interface unique et fixe

---

## Comment ce projet implémente Clean Architecture

### 1. La couche Domain est isolée
- L'entity `Cat` ne sait rien de HTTP ou des bases de données
- L'interface `CatRepository` est définie ici, pas implémentée

### 2. Inversion des dépendances en action
- Le use case dépend de l'interface `CatRepository`
- Infrastructure implémente cette interface
- L'une ou l'autre couche peut être changée indépendamment

### 3. Frontières claires entre les couches
```
domain/       ← Pas de dépendances
    ↓
usecase/      ← Dépend uniquement des interfaces du domain
    ↓
adapter/      ← Dépend du usecase (implémente les adaptateurs)
    ↓
infrastructure/ ← Dépend de tout ce qui précède (implémente tout)
```

### 4. Composition Root
Le fichier `main.go` est la "composition root" où toutes les dépendances sont connectées :

```go
// Couche par couche, les dépendances s'écoulent vers l'intérieur
catRepository := persistence.NewMemoryCatRepository()        // Infrastructure
catUseCase := usecase.NewCatUseCase(catRepository)          // UseCase
catPresenter := presenter.NewCatPresenter()                  // Adapter
catController := controller.NewCatController(catUseCase, catPresenter)  // Adapter
handler := router.SetupRouter(catController)                 // Infrastructure
```

C'est le seul endroit où les implémentations concrètes sont instanciées et connectées.

### 5. Facile à tester
Test du use case sans HTTP ou base de données :

```go
// Mock l'interface Repository
type MockRepository struct{}
func (m *MockRepository) Create(cat *entity.Cat) (*entity.Cat, error) {
    return cat, nil
}

// Test le use case avec le mock
repo := &MockRepository{}
uc := usecase.NewCatUseCase(repo)
result, _ := uc.CreateCat("id1", "Whiskers", "Persian", 3, "Orange")
```

### 6. Facile à étendre
Ajouter un nouveau use case ou Repository ne nécessite pas de changer le code existant :

```go
// Ajouter une nouvelle implémentation de Repository
type PostgresCatRepository struct { /* ... */ }

// Implémente l'interface CatRepository
func (p *PostgresCatRepository) Create(cat *entity.Cat) (*entity.Cat, error) {
    // Implémentation spécifique à PostgreSQL
}

// Connectez-le dans main.go
catRepository := postgres.NewPostgresCatRepository()
catUseCase := usecase.NewCatUseCase(catRepository)
```

---

## Principes clés démontrés

### 1. Dependency Inversion Principle (DIP)
Les modules de haut niveau (use cases) définissent les interfaces. Les modules de bas niveau (persistence) les implémentent.

```go
// Domain définit l'interface
type CatRepository interface { /* ... */ }

// Infrastructure l'implémente
type MemoryCatRepository struct { /* ... */ }

// UseCase dépend de l'interface, pas de l'implémentation
func NewCatUseCase(catRepository repository.CatRepository) *CatUseCase { /* ... */ }
```

### 2. Single Responsibility Principle (SRP)
Chaque couche a une responsabilité unique :
- Domain : Règles métier
- UseCase : Orchestration
- Adapter : Conversion de format
- Infrastructure : Détails du framework

### 3. Open/Closed Principle (OCP)
Ouvert à l'extension (ajouter de nouveaux Repository), fermé à la modification (code existant inchangé).

### 4. Interface Segregation Principle (ISP)
Chaque couche expose uniquement ce dont elle a besoin. L'interface `CatRepository` expose uniquement les opérations CRUD.

---

## Lectures complémentaires

- **Articles d'Uncle Bob :** https://blog.cleancoder.com/
- **Livre Clean Architecture :** « Clean Architecture » par Robert C. Martin
- **Golang Clean Architecture :** https://github.com/golang-standards/project-layout

---

## Licence

Ce projet est fourni comme un exemple éducatif des principes de Clean Architecture en Go.
