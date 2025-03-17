package model

type Seller struct {
    SellerID         string `json:"seller_id"`
    SellerChannel    string `json:"seller_channel"`
    LastEvaluation   string `json:"last_evaluation"`
    EstadoAdelantos  string `json:"estado_adelantos"`
    ResultadoReglas  string `json:"resultado_reglas"`
}