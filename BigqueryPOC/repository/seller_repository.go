package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/agudelozca/verdipoc/model"
)

type SellerRepository interface {
	FetchSellerByID(sellerID string) (model.Seller, error)
}

type sellerRepository struct {
	client *bigquery.Client
}

func NewSellerRepository(client *bigquery.Client) SellerRepository {
	return &sellerRepository{client: client}
}

func (r *sellerRepository) FetchSellerByID(sellerID string) (model.Seller, error) {
	ctx := context.Background()

	queryString := fmt.Sprintf(`
        SELECT 
            CUS_CUST_ID AS seller_id,
            SELLER_CHANNEL as seller_channel,
            MAX(EXECUTION_DATE) AS last_evaluation,
            CASE
                WHEN PRICING_GROUP_RESULT IS NOT NULL 
                AND ADVANCE_ACCESS_GROUP_RESULT IS NOT NULL
                AND DISALLOWED_RULES_RESULT IS NULL
                THEN 'Puede hacer adelantos'
                ELSE 'No puede hacer adelantos'
            END AS estado_adelantos,
            CASE
                WHEN DISALLOWED_RULES_RESULT IS NULL THEN 'No tiene situaciones de apagado'
                ELSE DISALLOWED_RULES_RESULT
            END AS resultado_reglas
        FROM ddme000138-53for0npjlh-furyid.TBL.BT_MIA_ELEGIBILITY_AUDIENCE_HISTORICAL
        WHERE CUS_CUST_ID = @sellerID
        GROUP BY
            CUS_CUST_ID,
            SELLER_CHANNEL,
            PRICING_GROUP_RESULT,
            ADVANCE_ACCESS_GROUP_RESULT,
            DISALLOWED_RULES_RESULT
    `)

	query := r.client.Query(queryString)
	query.Parameters = []bigquery.QueryParameter{
		{
			Name:  "sellerID",
			Value: sellerID,
		},
	}

	it, err := query.Read(ctx)
	if err != nil {
		return model.Seller{}, err
	}

	var seller model.Seller
	err = it.Next(&seller)
	if err != nil {
		return model.Seller{}, err
	}

	return seller, nil
}
