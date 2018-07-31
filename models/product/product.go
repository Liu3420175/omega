package product

import "omega/models"

type Product struct {
	models.Base
	Structure            string
	Upc                  string
	Name                 string
	Title                string
	Description          string
	Resume               string
	SaledNumber          int64
	StockNumber          int64
	IsActive             bool

}




type ProductDescription struct {
	Product              *Product
	Title                string
	Name                 string
}