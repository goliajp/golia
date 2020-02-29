package libgorm

import (
	"github.com/jinzhu/gorm"
)

func Preload(m *gorm.DB, preloadArr []string) *gorm.DB {
	if len(preloadArr) > 0 {
		for _, p := range preloadArr {
			m = m.Preload(p)
		}
	}
	return m
}
