package service

import (
    "github.com/agudelozca/verdipoc/model"
    "github.com/agudelozca/verdipoc/repository"
)

type SellerService interface {
    GetSellerByID(sellerID string) (model.Seller, error)
}

type sellerService struct {
    sellerRepo repository.SellerRepository
}

func NewSellerService(sellerRepo repository.SellerRepository) SellerService {
    return &sellerService{sellerRepo: sellerRepo}
}

func (s *sellerService) GetSellerByID(sellerID string) (model.Seller, error) {
    return s.sellerRepo.FetchSellerByID(sellerID)
}