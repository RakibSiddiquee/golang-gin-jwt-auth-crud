package helpers

import (
	"gorm.io/gorm"
)

func IsUniqueValue(db *gorm.DB, tableName, fieldName, value string) bool {
	var count int64

	db.Table(tableName).Where(fieldName+" = ?", value).Count(&count)

	return count == 0
}
