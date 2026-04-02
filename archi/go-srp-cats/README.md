# API CRUD Chat - Principe de Responsabilité Unique (SRP)

## Introduction

Ce projet démontre l'application stricte du **Principe de Responsabilité Unique (Single Responsibility Principle - SRP)**, l'un des cinq principes SOLID de Robert C. Martin pour concevoir des logiciels maintenables et extensibles.

### Citation Fondatrice

> **"Une classe devrait avoir une, et une seule, raison de changer."**
> — Robert C. Martin (Uncle Bob), Clean Architecture

Le SRP affirme que chaque module, classe ou fonction devrait être responsable d'une seule fonctionnalité. Cela signifie que le module n'a qu'une seule raison de changer : si les exigences relatives à cette seule responsabilité évoluent.

## Qu'est-ce que le Principe de Responsabilité Unique ?

Le SRP est basé sur l'idée que si une classe a plusieurs raisons de changer, elle a probablement plus qu'une responsabilité. Lorsqu'une classe a plusieurs responsabilités :

1. **Couplage accru** : Les changements dans une responsabilité affectent les autres
2. **Testabilité réduite** : Plus difficile à tester isolément
3. **Réutilisabilité limitée** : Impossible d'utiliser une partie sans les autres
4. **Maintenance complexe** : Les modifications risquent de casser d'autres fonctionnalités

En appliquant le SRP, nous créons un code qui est :
- **Modulaire** : Chaque composant est indépendant
- **Testable** : Facile à tester isolément
- **Réutilisable** : Peut être utilisé dans d'autres contextes
- **Maintenable** : Les changements sont localisés

## Diagramme d'Architecture SRP

```
┌─────────────────────────────────────────────────────────────────────┐
│                         HTTP REQUEST                               │
└────────────────────────────┬────────────────────────────────────────┘
                             │
                             ▼
                    ┌────────────────┐
                    │    ROUTER      │
                    │  Responsabilité:│
                    │ Mappage URL/HTTP│
                    └────────┬────────┘
                             │
              ┌──────────────┼──────────────┐
              ▼              ▼              ▼
         ┌────────┐  ┌──────────┐  ┌──────────┐
         │HANDLER │  │ HANDLER  │  │ HANDLER  │
         │HANDLER │  │ HANDLER  │  │ HANDLER  │
         │Créer   │  │Récupérer │  │Supprimer │
         └────┬───┘  └────┬─────┘  └────┬─────┘
              │           │             │
              └───────────┼─────────────┘
                          │
                          ▼
             ┌─────────────────────────┐
             │   SERIALIZATION         │
             │  (Request/Response)     │
             │ Responsabilité:         │
             │ Format d'entrée/sortie  │
             └────────┬────────────────┘
                      │
         ┌────────────┴────────────┐
         ▼                         ▼
    ┌─────────┐            ┌────────────┐
    │ DECODER │            │  ENCODER   │
    │ JSON→DTO│            │ Entity→JSON│
    └────┬────┘            └──────┬─────┘
         │                        │
         └────────────┬───────────┘
                      │
                      ▼
             ┌────────────────────┐
             │    CAT SERVICE      │
             │  Orchestration      │
             │ Responsabilité:     │
             │ Flux métier         │
             └──────┬──────────────┘
                    │
        ┌───────────┼───────────┐
        ▼           ▼           ▼
    ┌────────┐ ┌──────────┐ ┌──────────┐
    │VALIDATOR│ │ID GEN    │ │REPOSITORY│
    │Règles  │ │Générer ID│ │DAO       │
    │métier  │ │          │ │          │
    └────────┘ └──────────┘ └────┬─────┘
                                 │
                      ┌──────────┴────────────┐
                      ▼                       ▼
                 ┌──────────┐          ┌─────────────┐
                 │ STORAGE  │          │   ENTITY    │
                 │Moteur    │          │Pure Data    │
                 │(Memory,  │          │             │
                 │ DB, etc) │          │             │
                 └──────────┘          └─────────────┘

             ┌───────────────────────────────────────┐
             │       DOMAIN ERRORS                   │
             │  ValidationError, NotFoundError, etc  │
             └───────────────────────────────────────┘
```

## Tableau des Responsabilités Uniques

| Composant | Fichier | Responsabilité Unique | Raison de Changer |
|-----------|---------|----------------------|-------------------|
| **Entity** | `entity/cat.go` | Définition de la structure de données | Structure du Chat change (nouveau champ) |
| **ID Generator** | `id/generator.go` | Générer des IDs uniques | Stratégie d'ID change (UUID → Ulid → Snowflake) |
| **Validator** | `validation/cat_validator.go` | Valider les règles métier | Règles de validation du Chat changent |
| **Storage** | `storage/memory_store.go` | Persistance clé-valeur | Moteur de stockage change (Memory → PostgreSQL → Redis) |
| **Repository** | `repository/cat_repository.go` | Accès aux données du Chat | Logique de requête change |
| **Service** | `service/cat_service.go` | Orchestration du flux métier | Flux métier du Chat change |
| **Request DTO** | `serialization/request.go` | Décodage des requêtes JSON | Format d'entrée change (JSON → XML → Protobuf) |
| **Response DTO** | `serialization/response.go` | Encodage des réponses JSON | Format de sortie change |
| **Handlers** | `handler/*.go` | Traitement de chaque endpoint | Logique d'un endpoint spécifique change |
| **Router** | `routing/router.go` | Mappage des routes HTTP | Routes ou mappages HTTP changent |
| **Errors** | `errors/errors.go` | Taxonomie des erreurs métier | Types d'erreurs du domaine changent |

## Approche Naïve vs SRP

### Approche Naïve (Anti-Pattern)

```go
// ❌ MAUVAIS : Tout dans un seul fichier "main.go"
package main

import (
    "net/http"
    "encoding/json"
    "sync"
)

// Raisons de changer :
// 1. Format de données de Chat
// 2. Stratégie d'ID
// 3. Règles de validation
// 4. Moteur de stockage
// 5. Logique de requête
// 6. Format d'entrée/sortie
// 7. Logique métier
// 8. Endpoints HTTP

var cats = make(map[string]Cat)
var mu sync.RWMutex

type Cat struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Breed string `json:"breed"`
    Color string `json:"color"`
    Age   int    `json:"age"`
}

func createCat(w http.ResponseWriter, r *http.Request) {
    // Validation
    var cat Cat
    json.NewDecoder(r.Body).Decode(&cat)
    if cat.Name == "" {
        http.Error(w, "Name required", http.StatusBadRequest)
        return
    }
    // Génération d'ID
    cat.ID = generateID()
    // Stockage
    mu.Lock()
    cats[cat.ID] = cat
    mu.Unlock()
    // Réponse
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cat)
}

func updateCat(w http.ResponseWriter, r *http.Request) {
    // Même logique, dupliquée...
}

func getCat(w http.ResponseWriter, r *http.Request) {
    // Même logique, dupliquée...
}

// Problèmes :
// - Fichier énorme et complexe
// - Impossible de tester un aspect isolément
// - Changer la validation affecte potentiellement le stockage
// - Ajouter un nouvel endpoint duplique la logique
// - Passer à PostgreSQL demande de réécrire le fichier entier
```

**Problèmes majeurs :**
- Couplage extrême entre tous les aspects
- Tests difficiles et fragiles
- Changements localisés impossibles
- Réutilisabilité nulle
- Maintenabilité cauchemardesque

### Approche SRP (Ce Projet)

```
go-srp-cats/
├── internal/
│   ├── entity/cat.go              # Uniquement la structure
│   ├── id/generator.go            # Uniquement la génération d'ID
│   ├── validation/cat_validator.go  # Uniquement la validation
│   ├── storage/memory_store.go    # Uniquement le stockage
│   ├── repository/cat_repository.go # Uniquement l'accès aux données
│   ├── service/cat_service.go     # Uniquement l'orchestration
│   ├── serialization/request.go   # Uniquement le décodage des requêtes
│   ├── serialization/response.go  # Uniquement l'encodage des réponses
│   ├── handler/                   # Chaque handler = un endpoint
│   ├── routing/router.go          # Uniquement le mappage des routes
│   └── errors/errors.go           # Uniquement les types d'erreur
```

**Avantages majeurs :**
- Chaque fichier a UNE et UNE SEULE responsabilité
- Tests unitaires simples et ciblés
- Changements totalement localisés
- Réutilisabilité maximale
- Maintenabilité excellente

## Avantages du SRP

### 1. Maintenabilité Accrue
Changer une règle de validation ne touche que `validation/cat_validator.go`. Aucun risque d'effets de bord dans les handlers, le stockage ou le routage.

### 2. Testabilité Améliorée
Tester le validateur ne nécessite pas de configuration de base de données ou de serveur HTTP :

```go
func TestCatValidation(t *testing.T) {
    validator := NewCatValidator()
    cat := &Cat{Name: "", Breed: "Persan", Color: "Blanc", Age: 5}
    err := validator.ValidateCreate(cat)
    if err == nil {
        t.Fatal("Should fail for empty name")
    }
}
```

### 3. Réutilisabilité
Le `CatValidator` peut être utilisé dans une CLI, une API gRPC ou une queue de messages sans modification.

### 4. Flexibilité de Déploiement
Remplacer le stockage en mémoire par PostgreSQL ne nécessite de modifier que `storage/memory_store.go` et éventuellement `repository/cat_repository.go`.

### 5. Collaboration d'Équipe
Plusieurs développeurs peuvent travailler sur différents composants sans conflits :
- Dev A travaille sur les handlers
- Dev B travaille sur le validateur
- Dev C travaille sur le stockage

## Inconvénients du SRP (à Connaître)

### 1. Plus de Fichiers
Le code est éparpillé dans plus de fichiers. Nécessite une bonne organisation.

### 2. Complexité Initiale
L'architecture demande plus de temps de setup initialement.

### 3. Over-Engineering pour Petits Projets
Pour un script simple, c'est probablement excessif.

### 4. Apprentissage Requis
Nécessite de comprendre les principes SOLID et l'architecture en couches.

## Implémentation du SRP dans Ce Projet

### Exemple 1 : Changer une Règle de Validation

**Besoin :** L'âge maximum des chats passe de 50 à 40 ans.

**Impact :** Seul le fichier `validation/cat_validator.go` change :

```go
// Avant
if cat.Age < 0 || cat.Age > 50 {  // ← Changement ici
    return fmt.Errorf("cat age must be between 0 and 50")
}

// Après
if cat.Age < 0 || cat.Age > 40 {  // ← Changement uniquement ici
    return fmt.Errorf("cat age must be between 0 and 40")
}
```

Aucun autre fichier n'est affecté. Aucun redéploiement du handler, du routeur ou du service n'est nécessaire conceptuellement.

### Exemple 2 : Ajouter une Nouvelle Règle de Validation

**Besoin :** Les noms de chats ne doivent pas dépasser 50 caractères.

**Impact :** Seul `validation/cat_validator.go` change :

```go
func (v *CatValidator) validateBasicFields(cat *entity.Cat) error {
    if cat.Name == "" {
        return fmt.Errorf("cat name is required")
    }
    if len(cat.Name) > 50 {  // ← Nouvelle validation
        return fmt.Errorf("cat name must be 50 characters or less")
    }
    // ... reste du code
}
```

### Exemple 3 : Changer le Format de Sortie

**Besoin :** Passer de JSON à XML pour les réponses.

**Impact :** Seuls les fichiers de sérialisation changent :
- `serialization/response.go` - Utiliser `xml.Marshal` au lieu de `json.Marshal`
- Éventuellement `serialization/request.go` pour parser du XML

Aucun handler, aucun service, aucun validateur n'est affecté.

### Exemple 4 : Passer du Stockage Mémoire à PostgreSQL

**Besoin :** Persister les données dans PostgreSQL.

**Impact :** Principalement `storage/memory_store.go` change de façon architecturale :
- Remplacer `MemoryStore` par `PostgresStore`
- Adapter `repository/cat_repository.go` si la signature change
- Peut-être ajouter de la gestion des transactions

Les handlers, validateurs, service, routeur restent INCHANGÉS.

### Exemple 5 : Ajouter un Nouvel Endpoint

**Besoin :** Ajouter `GET /cats/breed/{breed}` pour rechercher par race.

**Impact :**
1. Créer `handler/get_cats_by_breed.go` (nouveau fichier)
2. Ajouter une méthode à `repository/cat_repository.go`
3. Ajouter une méthode à `service/cat_service.go`
4. Ajouter une route à `routing/router.go`

Les autres handlers restent INCHANGÉS. Zéro duplication de logique métier.

## Structure du Projet

```
go-srp-cats/
│
├── go.mod                          # Dépendances du module
├── README.md                       # Cette documentation
├── main.go                         # Point d'entrée - initialisation DI
│
└── internal/                       # Code privé au module
    │
    ├── entity/
    │   └── cat.go                 # ✓ Seule structure de données
    │                              # Raison de changer: structure du Chat
    │
    ├── id/
    │   └── generator.go           # ✓ Génération d'ID uniquement
    │                              # Raison de changer: stratégie d'ID
    │
    ├── validation/
    │   └── cat_validator.go       # ✓ Validation métier uniquement
    │                              # Raison de changer: règles de validation
    │
    ├── storage/
    │   └── memory_store.go        # ✓ Persistance clé-valeur uniquement
    │                              # Raison de changer: moteur de stockage
    │
    ├── repository/
    │   └── cat_repository.go      # ✓ Accès aux données uniquement
    │                              # Raison de changer: logique de requête
    │
    ├── service/
    │   └── cat_service.go         # ✓ Orchestration flux métier
    │                              # Raison de changer: flux métier
    │
    ├── serialization/
    │   ├── request.go             # ✓ Décodage requêtes JSON
    │   │                          # Raison de changer: format d'entrée
    │   └── response.go            # ✓ Encodage réponses JSON
    │                              # Raison de changer: format de sortie
    │
    ├── handler/
    │   ├── create_cat.go          # ✓ Handler POST /cats
    │   │                          # Raison de changer: logique créer
    │   ├── get_cat.go             # ✓ Handler GET /cats/{id}
    │   │                          # Raison de changer: logique récupérer
    │   ├── get_all_cats.go        # ✓ Handler GET /cats
    │   │                          # Raison de changer: logique lister
    │   ├── update_cat.go          # ✓ Handler PUT /cats/{id}
    │   │                          # Raison de changer: logique mettre à jour
    │   └── delete_cat.go          # ✓ Handler DELETE /cats/{id}
    │                              # Raison de changer: logique supprimer
    │
    ├── routing/
    │   └── router.go              # ✓ Mappage URL → handlers
    │                              # Raison de changer: mappage routes
    │
    └── errors/
        └── errors.go              # ✓ Types d'erreur domaine
                                   # Raison de changer: taxonomie erreurs
```

## Exemples d'Utilisation

### 1. Créer un Chat

```bash
curl -X POST http://localhost:8080/cats \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Minou",
    "breed": "Persan",
    "color": "Blanc",
    "age": 3
  }'
```

**Réponse (201 Created) :**
```json
{
  "id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
  "name": "Minou",
  "breed": "Persan",
  "color": "Blanc",
  "age": 3
}
```

### 2. Récupérer Tous les Chats

```bash
curl http://localhost:8080/cats
```

**Réponse (200 OK) :**
```json
[
  {
    "id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
    "name": "Minou",
    "breed": "Persan",
    "color": "Blanc",
    "age": 3
  }
]
```

### 3. Récupérer un Chat Spécifique

```bash
curl http://localhost:8080/cats/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6
```

**Réponse (200 OK) :**
```json
{
  "id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
  "name": "Minou",
  "breed": "Persan",
  "color": "Blanc",
  "age": 3
}
```

### 4. Mettre à Jour un Chat

```bash
curl -X PUT http://localhost:8080/cats/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Minou Félix",
    "breed": "Persan",
    "color": "Gris",
    "age": 4
  }'
```

**Réponse (200 OK) :**
```json
{
  "id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
  "name": "Minou Félix",
  "breed": "Persan",
  "color": "Gris",
  "age": 4
}
```

### 5. Supprimer un Chat

```bash
curl -X DELETE http://localhost:8080/cats/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6
```

**Réponse (204 No Content) :**
(Pas de body)

## Codes d'État HTTP

| Code | Cas d'Usage | Exemple |
|------|------------|---------|
| **201 Created** | Création réussie | POST /cats crée un Chat |
| **200 OK** | Opération réussie | GET /cats, PUT /cats/{id} |
| **204 No Content** | Suppression réussie | DELETE /cats/{id} |
| **400 Bad Request** | Données invalides | Validation échouée, format JSON invalide |
| **404 Not Found** | Ressource inexistante | GET /cats/{id} avec ID invalide |
| **405 Method Not Allowed** | Méthode HTTP incorrecte | GET sur un endpoint POST |

## Exécution du Projet

### Prérequis
- Go 1.21 ou supérieur

### Démarrer le serveur

```bash
go run main.go
```

Vous verrez :
```
Starting Cat CRUD API server on http://localhost:8080
Endpoints:
  POST   /cats       - Create a new cat
  GET    /cats       - Get all cats
  GET    /cats/{id}  - Get a specific cat
  PUT    /cats/{id}  - Update a cat
  DELETE /cats/{id}  - Delete a cat
```

### Exécuter les tests

```bash
go test ./internal/...
```

## Violations de SRP Évitées

Ce projet démontre comment ÉVITER ces pièges communs :

### ❌ Validation dans le Handler
```go
// MAUVAIS : Handler fait aussi de la validation
func createCat(w http.ResponseWriter, r *http.Request) {
    var cat Cat
    json.NewDecoder(r.Body).Decode(&cat)
    if cat.Name == "" {  // ← Validation ici
        http.Error(w, "Name required", http.StatusBadRequest)
        return
    }
    // Puis sauvegarde...
}
```

### ✓ Validation Séparée
```go
// BON : Validation isolée dans le validateur
type CatValidator struct{}
func (v *CatValidator) ValidateCreate(cat *Cat) error {
    if cat.Name == "" {
        return fmt.Errorf("cat name is required")
    }
    return nil
}

// Handler délègue
func (h *CreateCatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // ...
    if err := h.service.CreateCat(req.Name, req.Breed, req.Color, req.Age); err != nil {
        // ...
    }
}
```

### ❌ Logique Métier dans le Stockage
```go
// MAUVAIS : Repository fait de la validation
func (r *Repository) Save(cat *Cat) error {
    if cat.Name == "" {  // ← Validation ici
        return fmt.Errorf("invalid cat")
    }
    // Puis sauvegarde...
}
```

### ✓ Séparation Logique/Stockage
```go
// BON : Repository ne fait que du stockage
func (r *CatRepository) Save(cat *Cat) error {
    r.store.Set(cat.ID, cat)
    return nil
}

// Validation faite avant
service.CreateCat() → validator.ValidateCreate() → repository.Save()
```

## Principes SOLID Connexes

Bien que ce projet se concentre sur le SRP, il applique aussi :

### O - Open/Closed Principle
Les composants sont ouverts à l'extension (nouveaux types de validation, handlers) mais fermés à la modification.

### D - Dependency Inversion
Le `main.go` injecte les dépendances plutôt que de les créer à l'intérieur.

```go
svc := service.NewCatService(repo, validator, idGen)
handler := handler.NewCreateCatHandler(svc)
```

### L - Liskov Substitution Principle
Chaque composant peut être remplacé par une implémentation alternative (MongoDB au lieu de la mémoire) sans casser le contrat.

## Conclusion

Le Principe de Responsabilité Unique n'est pas une règle absolue mais une **guideline puissante** pour créer du code :
- **Maintenable** : Changements localisés et prévisibles
- **Testable** : Chaque composant testé isolément
- **Réutilisable** : Composants utilisables dans d'autres contextes
- **Scalable** : Structure supporte la croissance

Ce projet démontre que l'application rigoureuse du SRP crée une architecture flexible et professionnelle, capable d'évoluer sans régression.

### Ressources

- **Clean Code** par Robert C. Martin
- **Clean Architecture** par Robert C. Martin
- **SOLID Principles** : https://en.wikipedia.org/wiki/SOLID

---

**Créé comme démo du Principe de Responsabilité Unique (SRP)**
