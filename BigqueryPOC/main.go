package main

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"

	"cloud.google.com/go/bigquery"
	"github.com/agudelozca/verdipoc/controller"
	"github.com/agudelozca/verdipoc/repository"
	"github.com/agudelozca/verdipoc/service"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func main() {
	const projectID = "meli-bi-data"
	ctx := context.Background()
	base64Credentials := "ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAibWVsaS1iaS1kYXRhIiwKICAicHJpdmF0ZV9yZXlfaWQiOiAiM2FhMDQwZWZjZmFmMGMyNGUyZDQ2ODMzZmZjYWI2N2U5MDBhZTQ2OCIsCiAgInByaXZhdGVfa2V5IjogIi0=="
	credentials, err := base64.StdEncoding.DecodeString(base64Credentials)
	if err != nil {
		log.Fatalf("Error decoding base64 credentials: %v", err)
	}
	client, err := bigquery.NewClient(ctx, projectID, option.WithCredentialsJSON(credentials))
	if err != nil {
		log.Fatalf("Error creating BigQuery client: %v", err)
	}
	defer client.Close()

	sellerRepo := repository.NewSellerRepository(client)
	sellerService := service.NewSellerService(sellerRepo)
	sellerController := controller.NewSellerController(sellerService)

	r := mux.NewRouter()
	r.HandleFunc("/sellers/{sellerID}", sellerController.GetSeller).Methods("GET")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
