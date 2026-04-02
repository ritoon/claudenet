package sqlite

import (
	"context"
	"go-layered-cats/model"

	"gorm.io/gorm"
)

type SqliteStoreCat struct {
	db *gorm.DB
}

func (p *SqliteStoreCat) Set(ctx context.Context, key string, value *model.Cat) {
	// Implémenter la logique pour stocker une valeur dans PostgreSQL
	p.db.Create(value)
}

func (p *SqliteStoreCat) Get(ctx context.Context, key string) (*model.Cat, bool) {
	// Implémenter la logique pour récupérer une valeur depuis PostgreSQL
	var cat model.Cat
	result := p.db.Where("name = ?", key).First(&cat)
	if result.Error != nil {
		return nil, false
	}
	return &cat, true
}

func (p *SqliteStoreCat) Delete(ctx context.Context, key string) {
	// Implémenter la logique pour supprimer une valeur de PostgreSQL
	p.db.Where("name = ?", key).Delete(&model.Cat{})

}

func (p *SqliteStoreCat) GetAll(ctx context.Context) []*model.Cat {
	// Implémenter la logique pour récupérer toutes les valeurs depuis PostgreSQL
	var cats []model.Cat
	p.db.Find(&cats)

	var result []*model.Cat
	for _, cat := range cats {
		result = append(result, &cat)
	}
	return result
}

func (p *SqliteStoreCat) Exists(ctx context.Context, key string) bool {
	// Implémenter la logique pour vérifier l'existence d'une clé dans PostgreSQL
	var cat model.Cat
	result := p.db.Where("name = ?", key).First(&cat)
	if result.Error != nil {
		return false
	}
	return result.Error != gorm.ErrRecordNotFound
}
