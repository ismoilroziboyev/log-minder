package domain

import "math"

type Pagination struct {
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

func (p *Pagination) SetPaginationMetadata(limit, offset int32, total int64) {
	p.Limit = limit
	p.Offset = offset
	p.Total = total
	p.calculatePages()
}

func (p *Pagination) calculatePages() {
	if p.Limit <= 0 {
		return
	}
	p.TotalPages = int64(math.Ceil(float64(p.Total) / float64(p.Limit)))
}

type BasicFilter struct {
	Limit  int32 `json:"limit" form:"limit"`
	Offset int32 `json:"offset" form:"offset"`
}
