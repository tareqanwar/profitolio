// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AddAssetToPortfolioInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	Name             string  `json:"name"`
	PortfolioID      string  `json:"portfolioId"`
	Symbol           string  `json:"symbol"`
}

type AddAssetToPortfolioPayload struct {
	Asset            *PortfolioAsset `json:"asset"`
	ClientMutationID *string         `json:"clientMutationId"`
	Errors           []string        `json:"errors"`
}

type AddPortfolioTransactionInput struct {
	ClientMutationID *string             `json:"clientMutationId"`
	Fee              string              `json:"fee"`
	Notes            *string             `json:"notes"`
	PortfolioAssetID string              `json:"portfolioAssetId"`
	Price            string              `json:"price"`
	Quantity         string              `json:"quantity"`
	TransactionType  TransactionTypeEnum `json:"transactionType"`
}

type AddPortfolioTransactionPayload struct {
	ClientMutationID *string               `json:"clientMutationId"`
	Errors           []string              `json:"errors"`
	Transaction      *PortfolioTransaction `json:"transaction"`
}

type CreatePortfolioInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	Name             string  `json:"name"`
}

type CreatePortfolioPayload struct {
	ClientMutationID *string    `json:"clientMutationId"`
	Errors           []string   `json:"errors"`
	Portfolio        *Portfolio `json:"portfolio"`
}

type Portfolio struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	PortfolioAssets []*PortfolioAsset `json:"portfolioAssets"`
}

type PortfolioAsset struct {
	ID                    string                  `json:"id"`
	Name                  string                  `json:"name"`
	Portfolio             *Portfolio              `json:"portfolio"`
	PortfolioID           string                  `json:"portfolioId"`
	PortfolioTransactions []*PortfolioTransaction `json:"portfolioTransactions"`
	Symbol                string                  `json:"symbol"`
}

type PortfolioTransaction struct {
	Fee              float64             `json:"fee"`
	ID               string              `json:"id"`
	Notes            *string             `json:"notes"`
	PortfolioAsset   *PortfolioAsset     `json:"portfolioAsset"`
	PortfolioAssetID string              `json:"portfolioAssetId"`
	Price            float64             `json:"price"`
	Quantity         float64             `json:"quantity"`
	TransactionType  TransactionTypeEnum `json:"transactionType"`
}

type TransactionTypeEnum string

const (
	TransactionTypeEnumBuy  TransactionTypeEnum = "BUY"
	TransactionTypeEnumSell TransactionTypeEnum = "SELL"
)

var AllTransactionTypeEnum = []TransactionTypeEnum{
	TransactionTypeEnumBuy,
	TransactionTypeEnumSell,
}

func (e TransactionTypeEnum) IsValid() bool {
	switch e {
	case TransactionTypeEnumBuy, TransactionTypeEnumSell:
		return true
	}
	return false
}

func (e TransactionTypeEnum) String() string {
	return string(e)
}

func (e *TransactionTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransactionTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransactionTypeEnum", str)
	}
	return nil
}

func (e TransactionTypeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}