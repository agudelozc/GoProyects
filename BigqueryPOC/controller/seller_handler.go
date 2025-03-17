package controller

import (
    "encoding/json"
    "net/http"
    "github.com/agudelozca/verdipoc/service"
    "github.com/gorilla/mux"
)

type SellerController struct {
    sellerService service.SellerService
}

func NewSellerController(sellerService service.SellerService) *SellerController {
    return &SellerController{sellerService: sellerService}
}

func (c *SellerController) GetSeller(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    sellerID := vars["sellerID"]

    seller, err := c.sellerService.GetSellerByID(sellerID)
    if err != nil {
        http.Error(w, "Error fetching seller", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(seller)
}