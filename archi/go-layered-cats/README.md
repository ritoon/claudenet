# Go Layered Cats - API CRUD avec Architecture en Couches

Un exemple complet et prêt pour la production d'une API REST en Go implémentant le modèle d'architecture **Layered (N-Tier)** (architecture en couches). Ce projet démontre comment structurer une application Go avec une séparation claire des responsabilités à travers plusieurs couches architecturales.

## Aperçu

Ce projet implémente une API CRUD (Create, Read, Update, Delete) simple mais complète pour les chats avec un motif de conception de type PostgreSQL, en utilisant la bibliothèque standard Go pure (`net/http`) sans frameworks web externes. Il sert de référence éducative pour comprendre et mettre en œuvre une architecture en couches en Go.

## Qu'est-ce que l'Architecture en Couches ?

### Définition

**Layered Architecture** (aussi appelée N-Tier ou Multi-Couches) est un modèle architectural qui organise une application en couches horizontales, chacune avec une responsabilité spécifique. Chaque couche communique uniquement avec la couche directement en dessous (ou au même niveau), créant une séparation claire des préoccupations et améliorant la maintenabilité, la testabilité et la scalabilité.

### Principe Central : Séparation des Responsabilités

Chaque couche gère un aspect spécifique de l'application :

```
┌─────────────────────────────────────────────┐
│         Presentation Layer (HTTP)           │
│      (HTTP Handlers, Request/Response)      │
└─────────────────────────────────────────────┘
                      ↓ depends on
┌─────────────────────────────────────────────┐
│      Business Logic Layer (Service)         │
│  (Validation, Rules, Processing)            │
└─────────────────────────────────────────────┘
                      ↓ depends on
┌─────────────────────────────────────────────┐
│    Data Access Layer (Repository)           │
│   (Database Operations, Abstraction)        │
└─────────────────────────────────────────────┘
                      ↓ depends on
┌─────────────────────────────────────────────┐
│        Storage Layer (In-Memory/DB)         │
│     (Actual Data Persistence)               │
└─────────────────────────────────────────────┘
```

### Principes Clés

1. **Dépendance Unidirectionnelle** : Les couches ne peuvent dépendre que des couches en dessous. Les couches supérieures ne référencent jamais directement les détails internes des couches inférieures.

2. **Couches Fermées** : Chaque couche est fermée pour modification par les couches au-dessus. La communication se fait à travers des interfaces bien définies.

3. **Responsabilité Unique** : Chaque couche a une seule raison de changer. Les changements en logique métier n'affectent pas la gestion HTTP, et les changements de base de données n'affectent pas la couche de service.

4. **Abstraction** : Les couches inférieures exposent des interfaces abstraites sur lesquelles les couches supérieures dépendent, pas des implémentations concrètes.

## Architecture du Projet

Ce projet implémente 4 couches (plus la couche modèle/domaine) :

### Couche 1 : Modèle (Couche de Domaine)

**Emplacement** : `model/cat.go`

**Responsabilité** : Définir les structures de données qui représentent vos entités métier.

**Caractéristiques** :
- Contient uniquement des structures de données (structs)
- Pas de logique métier
- Annotations de marshaling/unmarshaling JSON
- Pure représentation de données

**Exemple** :
```go
type Cat struct {
    ID    string
    Name  string
    Breed string
    Age   int
    Color string
}
```

### Couche 2 : Stockage (Couche de Données)

**Emplacement** : `storage/memory.go`

**Responsabilité** : Fournir des opérations de persistance de données de bas niveau.

**Caractéristiques** :
- Opérations thread-safe (utilise des mutexes)
- Opérations de clé-valeur génériques
- Pas de logique métier ou connaissance du domaine
- Peut être remplacé par des pilotes de base de données

**Pourquoi Séparer le Stockage du Repository ?**
- Le stockage gère le "comment" de la persistance (mémoire, base de données, système de fichiers)
- Le Repository gère le "quoi" et "où" les requêtes sont exécutées
- Rend facile l'échange d'implémentations (test avec mémoire, production avec PostgreSQL)

### Couche 3 : Repository (Couche d'Accès aux Données)

**Emplacement** : `repository/cat_repository.go`

**Responsabilité** : Fournir des opérations d'accès aux données spécifiques au domaine.

**Caractéristiques** :
- Dépend de la couche Storage
- Les méthodes correspondent aux entités du domaine (Create, GetByID, Update, Delete)
- Pas de logique métier ou validation
- Abstrait les détails d'implémentation du stockage

**Exemple de Méthodes** :
- `Create(cat *Cat) error`
- `GetByID(id string) (*Cat, error)`
- `GetAll() ([]*Cat, error)`
- `Update(id string, cat *Cat) error`
- `Delete(id string) error`

**Pourquoi Cette Couche Est Importante** :
- Le motif Repository cache les requêtes de base de données derrière une interface propre
- La couche Service ne sait pas si les données proviennent de la mémoire, PostgreSQL ou MongoDB
- Rend les tests plus faciles avec des repositories fictifs

### Couche 4 : Service (Couche Logique Métier)

**Emplacement** : `service/cat_service.go`

**Responsabilité** : Implémenter les règles métier, la validation et l'orchestration.

**Caractéristiques** :
- Dépend de la couche Repository
- Contient la logique de validation
- Gère l'application des règles métier
- Génère les UUIDs pour les nouvelles entités
- Orchestre plusieurs appels de repository si nécessaire

**Exemple de Validations** :
- Le nom du chat ne peut pas être vide
- L'âge du chat doit être entre 0 et 50
- La race et la couleur du chat sont obligatoires

**Pourquoi Cette Couche Est Importante** :
- Centralise la logique métier que plusieurs handlers pourraient utiliser
- Rend les règles métier réutilisables entre HTTP, CLI ou opérations batch
- Valide les données avant la persistance
- Implémente des opérations spécifiques au domaine qui vont au-delà du simple CRUD

### Couche 5 : Handler (Couche Présentation/HTTP)

**Emplacement** : `handler/cat_handler.go`

**Responsabilité** : Gérer les requêtes et réponses HTTP.

**Caractéristiques** :
- Dépend de la couche Service
- Analyse les requêtes JSON
- Valide la méthode HTTP
- Mappe les codes de statut HTTP aux résultats métier
- Sérialise les réponses

**Points d'Accès HTTP** :
- `POST /cats` → Create
- `GET /cats` → Read All
- `GET /cats/{id}` → Read One
- `PUT /cats/{id}` → Update
- `DELETE /cats/{id}` → Delete

**Pourquoi Cette Couche Est Importante** :
- Isole les préoccupations HTTP de la logique métier
- Facile d'ajouter des handlers WebSocket ou gRPC sans toucher à la logique métier
- Testable sans simuler des serveurs HTTP

## Avantages de l'Architecture en Couches

### 1. Séparation des Responsabilités
```
Avant : Tout mélangé ensemble
Code API + Logique Métier + Requêtes BD = Difficile à tester, difficile à changer

Après : Séparation claire
API (Handler) → Règles Métier (Service) → Accès aux Données (Repository) → Stockage
```

### 2. Testabilité
Chaque couche peut être testée indépendamment :

```go
// Tester le service sans HTTP
service := service.NewCatService(mockRepository)
cat, err := service.CreateCat(&model.Cat{...})

// Tester le repository sans base de données
repo := repository.NewCatRepository(mockStore)
cat, err := repo.Create(&model.Cat{...})

// Tester le handler avec un service fictif
handler := handler.NewCatHandler(mockService)
// Faire des requêtes HTTP au handler
```

### 3. Réutilisabilité
La logique métier fonctionne sur plusieurs interfaces :

```go
// Même service utilisé pour l'API HTTP
httpHandler := handler.NewCatHandler(catService)

// Pourrait aussi être utilisé pour CLI
cliApp := cli.NewApp(catService)

// Ou traitement batch
batchProcessor := batch.NewProcessor(catService)

// Ou WebSocket si nécessaire
wsHandler := websocket.NewHandler(catService)
```

### 4. Maintenabilité
Les changements restent localisés :

- **Changer le format HTTP (JSON → XML) ?** Modifier uniquement la couche Handler
- **Ajouter une règle de validation ?** Modifier uniquement la couche Service
- **Changer de base de données ?** Modifier uniquement les couches Repository/Storage
- **Mettre à jour le modèle Chat ?** Toucher à toutes les couches, mais chaque modification est minimale

### 5. Scalabilité
Facile d'améliorer chaque couche :

- Ajouter la mise en cache dans la couche Repository
- Ajouter la limitation de débit dans la couche Handler
- Ajouter la journalisation d'audit dans la couche Service
- Ajouter le pooling de connexion dans la couche Storage

### 6. Injection de Dépendances
Les dépendances claires facilitent l'injection de mocks :

```go
// Utilisation réelle
realHandler := handler.NewCatHandler(realService)

// Test
mockService := new(MockCatService)
testHandler := handler.NewCatHandler(mockService)
```

### 7. Architecture Claire
Les nouveaux développeurs peuvent :
- Trouver la logique HTTP dans Handler
- Trouver la validation dans Service
- Trouver les requêtes de base de données dans Repository
- Savoir exactement ce que fait chaque fichier

## Inconvénients de l'Architecture en Couches

### 1. Sur-ingénierie pour les Cas Simples
Une fonction utilitaire simple n'a pas besoin de 4 couches. L'architecture en couches ajoute une surcharge.

**Quand elle devient excessive** :
- Scripts simples (< 100 lignes)
- Utilitaires ponctuels
- Preuves de concept

### 2. Complexité Accrue
Plus de couches = plus de fichiers à parcourir, plus de code à maintenir.

**Comment atténuer** :
- Ne pas créer de couches inutilisées
- Garder les couches minces
- Utiliser une nomenclature claire

### 3. Surcharge de Performance Potentielle
Chaque appel de fonction traverse plusieurs couches.

**Impact réel** : Négligeable dans la plupart des applications (nanosecondes par appel)
**Quand c'est important** : Les systèmes de trading à très haute fréquence

### 4. Peut Masquer une Mauvaise Conception
Vous pouvez toujours écrire du mauvais code dans chaque couche. La structuration en couches ne corrige pas les problèmes de conception fondamentaux.

**Atténuation** : Revues de code, tests, documentation

### 5. Base de Données Partagée Entre Services
Plusieurs services accédant à la même base de données peuvent créer un couplage étroit.

**Solution** : Utiliser des bases de données séparées par service (motif microservices)

## Comment Ce Projet Implémente l'Architecture en Couches

### Flux de Requête (Chemin de Dépendance)

```
Requête HTTP
    ↓
[Handler Layer] Reçoit la requête, valide la méthode HTTP
    ↓
[Handler] Analyse JSON → Crée une instance de modèle
    ↓
[Handler] Appelle la méthode de la couche Service
    ↓
[Service Layer] Valide les règles métier (âge, nom, race)
    ↓
[Service] Appelle la méthode de la couche Repository
    ↓
[Repository Layer] Traduit en opérations de stockage
    ↓
[Repository] Appelle la couche Storage
    ↓
[Storage Layer] Effectue l'opération de données réelle (map en mémoire)
    ↓
La réponse remonte à travers les couches
    ↓
[Handler] Retourne la réponse HTTP avec le code de statut approprié
    ↓
Réponse HTTP au Client
```

### Exemple : Créer un Chat

```go
// Handler reçoit : POST /cats { "name": "Whiskers", ... }
handler.CreateCat(w, r) {
    // 1. Analyser JSON
    var cat Cat = r.Body

    // 2. Appeler le service
    createdCat := h.service.CreateCat(&cat)
}

service.CreateCat(cat) {
    // 3. Valider les règles métier
    if cat.Age < 0 { return error }

    // 4. Générer ID
    cat.ID = uuid.New()

    // 5. Appeler le repository
    s.repo.Create(cat)
}

repo.Create(cat) {
    // 6. Appeler le stockage
    r.store.Set(cat.ID, cat)
}

store.Set(key, value) {
    // 7. Effectuer l'opération de stockage réelle
    m.data[key] = value
}
```

### Exemple : Gestion des Erreurs À Travers les Couches

```
Handler (HTTP 400) ← Service (erreur de validation) ← Erreur de données
Handler (HTTP 404) ← Service ("non trouvé") ← Repository (null)
Handler (HTTP 500) ← Service (erreur de base de données) ← Erreur de Repository
```

## Explication de la Structure des Fichiers

```
go-layered-cats/
├── README.md                    # Ce fichier - documentation de l'architecture
├── go.mod                       # Définition du module Go et dépendances
├── main.go                      # Point d'entrée de l'application, initialisation des couches
│
├── model/                       # Couche de domaine - structures de données
│   └── cat.go                   # Définition de l'entité Chat
│
├── storage/                     # Couche de stockage - persistance de bas niveau
│   └── memory.go                # Implémentation du stockage en mémoire
│
├── repository/                  # Couche Repository - accès aux données
│   └── cat_repository.go        # Repository spécifique aux chats (requêtes du domaine)
│
├── service/                     # Couche de service - logique métier
│   └── cat_service.go           # Logique métier spécifique aux chats
│
└── handler/                     # Couche Handler - présentation HTTP
    └── cat_handler.go           # Gestion des requêtes/réponses HTTP pour chats
```

## Installation et Utilisation

### Prérequis

- Go 1.21 ou version ultérieure
- Commande `git`

### Configuration

```bash
# Cloner ou naviguer jusqu'au répertoire du projet
cd go-layered-cats

# Télécharger les dépendances
go mod download

# Exécuter le serveur
go run main.go
```

Le serveur démarrera sur `http://localhost:8080`

### Points d'Accès API

#### Créer un Chat
```bash
curl -X POST http://localhost:8080/cats \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Whiskers",
    "breed": "Persian",
    "age": 3,
    "color": "Orange"
  }'
```

Réponse (201 Created) :
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Whiskers",
  "breed": "Persian",
  "age": 3,
  "color": "Orange"
}
```

#### Obtenir Tous les Chats
```bash
curl http://localhost:8080/cats
```

Réponse (200 OK) :
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Whiskers",
    "breed": "Persian",
    "age": 3,
    "color": "Orange"
  }
]
```

#### Obtenir un Chat Spécifique
```bash
curl http://localhost:8080/cats/550e8400-e29b-41d4-a716-446655440000
```

Réponse (200 OK) :
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Whiskers",
  "breed": "Persian",
  "age": 3,
  "color": "Orange"
}
```

#### Mettre à Jour un Chat
```bash
curl -X PUT http://localhost:8080/cats/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Whiskers Updated",
    "breed": "Persian",
    "age": 4,
    "color": "Orange"
  }'
```

Réponse (200 OK) :
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Whiskers Updated",
  "breed": "Persian",
  "age": 4,
  "color": "Orange"
}
```

#### Supprimer un Chat
```bash
curl -X DELETE http://localhost:8080/cats/550e8400-e29b-41d4-a716-446655440000
```

Réponse (204 No Content)

#### Vérification de Santé
```bash
curl http://localhost:8080/health
```

Réponse (200 OK) :
```json
{"status":"healthy"}
```

## Exemples de Code pour Apprendre

### Comprendre les Dépendances

Toutes les couches dépendent de la couche directement en dessous :

```go
// main.go - Montre clairement la chaîne de dépendances
store := storage.NewMemoryStore()           // Level 0
catRepo := repository.NewCatRepository(store)  // Dépend du stockage
catService := service.NewCatService(catRepo)   // Dépend du repository
catHandler := handler.NewCatHandler(catService) // Dépend du service
```

### Ajouter une Nouvelle Règle Métier

Si vous devez ajouter une validation que les chats ne peuvent pas être nommés "Felix" :

**Modifier uniquement la couche service** :
```go
// Dans service/cat_service.go
func validateCat(cat *model.Cat) error {
    // ... validations existantes ...

    if strings.ToLower(cat.Name) == "felix" {
        return errors.New("cats cannot be named Felix (copyright protection)")
    }

    return nil
}
```

**Aucune modification nécessaire** dans les couches Handler, Repository ou Storage.

### Changer le Backend de Stockage

Pour passer du stockage en mémoire à une base de données :

**Modifier uniquement la couche repository** (et ajouter un nouvel adaptateur de stockage) :

```go
// repository/cat_repository.go - reste inchangé en termes d'interface
type CatRepository struct {
    store CatStore  // Interface, pas un type concret
}

// Créer un nouvel adaptateur de stockage
// storage/postgres.go
type PostgresStore struct {
    db *sql.DB
}

// Implémente la même interface que MemoryStore
func (p *PostgresStore) Set(key string, value interface{}) { ... }
func (p *PostgresStore) Get(key string) (interface{}, bool) { ... }

// Utilisation dans main.go
postgresStore := storage.NewPostgresStore(dbConnection)
catRepo := repository.NewCatRepository(postgresStore)  // Même interface !
```

## Stratégie de Test

### Test Unitaire de la Couche Service

```go
// service_test.go
func TestCreateCatValidation(t *testing.T) {
    mockRepo := new(MockRepository)
    service := service.NewCatService(mockRepo)

    // Tester l'âge invalide
    cat := &model.Cat{Name: "Test", Breed: "Test", Age: 60, Color: "Test"}
    _, err := service.CreateCat(cat)

    if err == nil {
        t.Fatal("expected error for age > 50")
    }
}
```

### Test Unitaire de la Couche Repository

```go
// repository_test.go
func TestCreateAndRetrieve(t *testing.T) {
    store := storage.NewMemoryStore()
    repo := repository.NewCatRepository(store)

    cat := &model.Cat{ID: "123", Name: "Test"}
    repo.Create(cat)

    retrieved, _ := repo.GetByID("123")
    if retrieved.Name != "Test" {
        t.Fatal("cat not stored correctly")
    }
}
```

### Test d'Intégration de la Couche Handler

```go
// handler_test.go
func TestCreateCatEndpoint(t *testing.T) {
    mockService := new(MockCatService)
    handler := handler.NewCatHandler(mockService)

    req := httptest.NewRequest("POST", "/cats", body)
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    if w.Code != 201 {
        t.Fatalf("expected 201, got %d", w.Code)
    }
}
```

## Étendre Ce Projet

### Ajouter la Journalisation
```go
// service/cat_service.go
func (s *CatService) CreateCat(cat *model.Cat) (*model.Cat, error) {
    log.Printf("Creating cat: %s", cat.Name)
    // ... reste de la logique
}
```

### Ajouter la Limitation de Débit
```go
// handler/cat_handler.go
func (h *CatHandler) CreateCat(w http.ResponseWriter, r *http.Request) {
    if !h.rateLimiter.Allow() {
        writeJSONError(w, "Rate limit exceeded", http.StatusTooManyRequests)
        return
    }
    // ... reste de la logique
}
```

### Ajouter une Base de Données
```go
// storage/postgres.go - nouveau fichier
type PostgresStore struct {
    db *sql.DB
}

// Implémenter la même interface que MemoryStore
// Ensuite dans main.go, il suffit de changer l'implémentation du store
```

### Ajouter un Middleware
```go
// main.go
handler := loggingMiddleware(authMiddleware(catHandler))
```

## Conclusion

L'architecture en couches fournit un moyen éprouvé de structurer les applications avec une séparation claire des responsabilités. Bien qu'elle ajoute une certaine complexité, elle paie en dividendes en termes de maintenabilité, testabilité et scalabilité.

Ce projet sert d'implémentation de référence montrant :

- Des définitions de couches claires et des responsabilités
- Un flux de dépendances approprié (toujours vers le bas)
- Une conception testable
- Facile à étendre et modifier
- Gestion d'erreurs prête pour la production

Utilisez ceci comme modèle pour vos propres applications Go !

## Ressources Supplémentaires

- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
