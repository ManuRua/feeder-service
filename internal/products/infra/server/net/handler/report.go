package handler

import (
	"feeder-service/internal/products/application/count"
	"fmt"
)

type ReportHandler struct {
	countProductsUseCase count.CountProductsUseCase
}

// NewReportHandler creates a new handler for reporting about products
func NewReportHandler(
	countProductsUseCase count.CountProductsUseCase,
) ReportHandler {
	handler := ReportHandler{
		countProductsUseCase,
	}

	return handler
}

// Handle executes use case and print the result
func (h ReportHandler) Handle() {
	counts := h.countProductsUseCase.CountProducts()

	fmt.Printf("Received %d unique product skus, %d duplicates and %d discarded values.\n", counts.Uniques, counts.Duplicated, counts.Invalids)
}
