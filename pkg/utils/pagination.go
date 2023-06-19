package utils

import (
	"math"
)

type Pagination struct {
	//Page represents the current page number, and PageSize indicates the number of records per page.
	Page     int
	PageSize int
}

type Metadata struct {
	CurrentPage  int
	PageSize     int
	FirstPage    int
	LastPage     int
	TotalRecords int64
}

// The Limit() method is defined for the Pagination struct. It returns the PageSize field, which represents the limit for the number of records to be fetched per page.
func (f *Pagination) Limit() int {
	return f.PageSize
}

func (f *Pagination) Offset() int {
	return int(math.Abs(float64((f.Page - 1) * f.PageSize)))
}
func ComputeMetadata(totalRecords *int64, CurrentPage *int, pageSize *int) Metadata {

	if *totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  *CurrentPage,
		PageSize:     *pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(*totalRecords) / float64(*pageSize))),
		TotalRecords: *totalRecords,
	}
}
