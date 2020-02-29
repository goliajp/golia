package libgorm

import (
	"github.com/jinzhu/gorm"
)

type ListFilterSet struct {
	Sort    string   `json:"order"`
	Offset  int      `json:"offset"`
	Limit   int      `json:"limit"`
	CondArr []string `json:"cond_arr"`
}

func ListFilter(db *gorm.DB, f *ListFilterSet) *gorm.DB {
	if f != nil {
		if len(f.CondArr) > 0 { // condition arr
			for _, c := range f.CondArr {
				db = db.Where(c) // add where query
			}
		}
		if f.Sort != "" {
			db = db.Order(f.Sort)
		}
		if f.Offset != 0 {
			db = db.Offset(f.Offset)
		}
		if f.Limit != 0 {
			db = db.Limit(f.Limit)
		}
	}
	return db
}
