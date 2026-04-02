package postgres

import (
	"context"
	"go-layered-cats/model"

	"gorm.io/gorm"
)

type PostgresStoreDog struct {
	db *gorm.DB
}

func (p *PostgresStoreDog) Set(ctx context.Context, key string, value *model.Dog) {
	// Implémenter la logique pour stocker une valeur dans PostgreSQL
	p.db.Create(value)
}

func (p *PostgresStoreDog) Get(ctx context.Context, key string) (*model.Dog, bool) {
	// Implémenter la logique pour récupérer une valeur depuis PostgreSQL
	var cat model.Dog
	result := p.db.Where("name = ?", key).First(&cat)
	if result.Error != nil {
		return nil, false
	}
	return &cat, true
}

func (p *PostgresStoreDog) Delete(ctx context.Context, key string) {
	// Implémenter la logique pour supprimer une valeur de PostgreSQL
	p.db.Where("name = ?", key).Delete(&model.Dog{})

}

func (p *PostgresStoreDog) GetAll(ctx context.Context) []*model.Dog {
	// Implémenter la logique pour récupérer toutes les valeurs depuis PostgreSQL
	var dogs []model.Dog
	p.db.Find(&dogs)

	var result []*model.Dog
	for _, dog := range dogs {
		result = append(result, &dog)
	}
	return result
}

func (p *PostgresStoreDog) Exists(ctx context.Context, key string) bool {
	// Implémenter la logique pour vérifier l'existence d'une clé dans PostgreSQL
	var dog model.Dog
	result := p.db.Where("name = ?", key).First(&dog)
	if result.Error != nil {
		return false
	}
	return result.Error != gorm.ErrRecordNotFound
}
