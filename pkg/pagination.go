package pkg

import (
	"database/sql"
	"math"
)

type Pagination struct {
	totalCount uint32
	page       uint32
	limit      uint32
	offset     uint32
}

func NewPagination(page uint32, limit int32, offset uint32, perPage int32, totalCounts uint32) *Pagination {
	if limit == 0 {
		limit = perPage
	}

	pagination := Pagination{
		totalCount: totalCounts,
		page:       1,
		limit:      uint32(limit),
	}

	if page != 0 {
		pagination.page = page
	}

	if limit < 0 {
		pagination.limit = 0
		pagination.offset = 0
		pagination.page = 0
	}

	pagination.offset = pagination.getOffset()
	if offset != 0 {
		pagination.offset = offset
	}

	return &pagination
}

func (p *Pagination) Offset() uint32 {
	return p.offset
}

func (p *Pagination) Limit() uint32 {
	return p.limit
}

func (p *Pagination) LimitSQL() sql.NullInt32 {
	return sql.NullInt32{Int32: int32(p.limit), Valid: p.limit > 0}
}

func (p *Pagination) PerPage() uint32 {
	if p.limit == 0 {
		return p.totalCount
	}

	return p.limit
}

func (p *Pagination) TotalItems() uint32 {
	return p.totalCount
}

func (p *Pagination) PagesCount() uint32 {
	if p.limit == 0 {
		return 1
	}

	totalCount := math.Ceil(float64(p.totalCount) / float64(p.limit))
	return uint32(totalCount)
}

func (p *Pagination) getOffset() uint32 {
	return (p.page - 1) * p.limit
}
