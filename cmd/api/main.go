package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Julio-Norberto/api-message/internal/infra/akafka"
	"github.com/Julio-Norberto/api-message/internal/infra/repository"
	"github.com/Julio-Norberto/api-message/internal/infra/web"
	"github.com/Julio-Norberto/api-message/internal/usecases"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecases.NewCreateProductUseCase(repository)
	listProductsUseCase := usecases.NewListProductsUseCase(repository)

	productHandles := web.NewProductHandlers(createProductUseCase, listProductsUseCase)

	r := chi.NewRouter()
	r.Post("/products", productHandles.CreateProductHandler)
	r.Get("/products", productHandles.ListProductHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecases.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			//logar o erro
		}
		_, err = createProductUseCase.Execute(dto)
	}
}
