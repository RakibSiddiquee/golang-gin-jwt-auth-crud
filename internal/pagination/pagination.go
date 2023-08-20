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
func Paginate(db *gorm.DB, page, limit int, preloadFunc func(*gorm.DB) *gorm.DB, output interface{}) (PaginateResult, error) {
	offset := (page - 1) * limit

	query := db
	if preloadFunc != nil {
		query = preloadFunc(query)
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
func RawPaginate(db *gorm.DB, page, limit int, rawFunc func(*gorm.DB) *gorm.DB, output interface{}) (PaginateResult, error) {
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

// Paginate returns a function that performs pagination on GORM queries
/*func Paginate(page, limit int, preloadFunc func(*gorm.DB) *gorm.DB) func(db *gorm.DB, out interface{}) (result PaginateResult, err error) {
	return func(db *gorm.DB, out interface{}) (result PaginateResult, err error) {
		offset := (page - 1) * limit
		query := db

		if preloadFunc != nil {
			query = preloadFunc(query)
		}

		// Perform the query for paginated data
		err = query.Offset(offset).Limit(limit).Find(out).Error
		if err != nil {
			return result, err
		}

		// Get the total count of records
		query.Model(out).Count(&result.Total)

		result.Data = out
		result.CurrentPage = page
		result.From = offset + 1

		to := offset + limit
		if to > int(result.Total) {
			to = int(result.Total)
		}
		result.To = to
		result.LastPage = (int(result.Total) + limit - 1) / limit
		result.PerPage = limit

		return result, nil
	}
}
*/
