package pagination

import "gorm.io/gorm"

// PaginateResult represents the pagination metadata and data
type PaginateResult struct {
	Data        interface{} `json:"data"`
	CurrentPage int         `json:"current_page"`
	From        int         `json:"from"`
	To          int         `json:"to"`
	LastPage    int         `json:"last_page"`
	PerPage     int         `json:"per_page"`
	Total       int64       `json:"total"`
}

// Paginate returns a function that performs pagination on GORM queries
func Paginate(db *gorm.DB, page, limit int, rawFunc func(*gorm.DB) *gorm.DB, output interface{}) (PaginateResult, error) {
	offset := (page - 1) * limit

	query := db
	if rawFunc != nil {
		query = rawFunc(query)
	}

	var total int64
	query.Model(output).Count(&total)

	err := query.Offset(offset).Limit(limit).Find(output).Error
	if err != nil {
		return PaginateResult{}, nil
	}

	to := offset + limit
	if to > int(total) {
		to = int(total)
	}

	return PaginateResult{
		Data:        output,
		CurrentPage: page,
		From:        offset + 1,
		To:          to,
		LastPage:    (int(total) + limit - 1) / limit,
		PerPage:     limit,
		Total:       total,
	}, nil
}

// RawPaginate returns a function that performs pagination on GORM joins, select or raw queries
/*func RawPaginate(db *gorm.DB, page, limit int, rawFunc func(*gorm.DB) *gorm.DB, output interface{}) (PaginateResult, error) {
	offset := (page - 1) * limit

	query := rawFunc(db)

	var total int64
	query.Model(output).Count(&total)

	err := query.Offset(offset).Limit(limit).Find(output).Error
	if err != nil {
		return PaginateResult{}, nil
	}

	to := offset + limit
	if to > int(total) {
		to = int(total)
	}

	return PaginateResult{
		Data:        output,
		CurrentPage: page,
		From:        offset + 1,
		To:          to,
		LastPage:    (int(total) + limit - 1) / limit,
		PerPage:     limit,
		Total:       total,
	}, nil
}
*/
