package web

import (
	"encoding/json"
	"net/http"

	"github.com/Julio-Norberto/api-message/internal/usecases"
)

type ProductHandlers struct {
	CreateProductUseCase *usecases.CreateProductUseCase
	ListProductsUseCase  *usecases.ListProductsUseCase
}

func NewProductHandlers(createProductUseCase *usecases.CreateProductUseCase, listProductsUseCase *usecases.ListProductsUseCase) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase: createProductUseCase,
		ListProductsUseCase:  listProductsUseCase,
	}
}

func (p *ProductHandlers) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var input usecases.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.CreateProductUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
