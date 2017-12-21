package registry

import (
	"github.com/suusan2go/familog-api/domain/model"
	"github.com/suusan2go/familog-api/domain/repository"
	"github.com/suusan2go/familog-api/infrastructure/gorm"
)

// Registry DI
type Registry struct {
	DB *model.DB
}

func (r Registry) DiaryEntryRepository() repository.DiaryEntryRepository {
	return gorm.DiaryEntryRepository{r.DB}
}
