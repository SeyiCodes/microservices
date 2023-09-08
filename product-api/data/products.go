package data

import (
   t "time"
   "io"
   "encoding/json"
)

type Product struct {
	ID           	int			`json:"id"`
	Name         	string		`json:"name"`
	Description 	string		`json:"description"`
	Price     		float32		`json:"price"`
	SKU 			string		`json:"sku"`
	CreatedOn 		string		`json:"-"`
	UpdatedOn 		string		`json:"-"`
	DeletedOn 		string		`json:"-"`
}

type Products []*Product

func (p*Products) ToJSON(w io.Writer) error {
	en := json.NewEncoder(w)
	return en.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
			ID:            	1,
			Name:          	"Latte",
			Description:  	"Frotty milky coffee",
			Price:        	2.45,
			SKU:           	"lat323",
			CreatedOn:     	t.Now().UTC().String(),
			UpdatedOn:     	t.Now().UTC().String(),
    },
	&Product{
			ID:          	2,
			Name:        	"Cappuccino",
			Description: 	"Espresso mixed with steamed milk foam",
			Price:       	3.25,
			SKU:         	"Cap456",
			CreatedOn:     	t.Now().UTC().String(),
        	UpdatedOn:    	t.Now().UTC().String(),
	},
	&Product{
			ID:          	3,
			Name:        	"Mocha",
			Description: 	"Espresso with hot milk, chocolate, and whipped cream",
			Price:       	3.75,
			SKU:         	"moc789",
			CreatedOn:   	t.Now().UTC().String(),
			UpdatedOn:   	t.Now().UTC().String(),
	},
	&Product{
			ID:          	4,
			Name:        	"Espresso",
			Description: 	"Strong black coffee made by forcing steam through finely-ground coffee beans",
			Price:       	2.00,
			SKU:         	"esp101",
			CreatedOn:   	t.Now().UTC().String(),
			UpdatedOn:   	t.Now().UTC().String(),
	},
	&Product{
			ID:          	5,
			Name:        	"Chai Latte",
			Description: 	"A spiced tea with milk and various spices",
			Price:       	4.50,
			SKU:         	"chai456",
			CreatedOn:   	t.Now().UTC().String(),
			UpdatedOn:   	t.Now().UTC().String(),
	},
	&Product{
			ID:          	6,
			Name:        	"Iced Coffee",
			Description: 	"Chilled coffee beverage",
			Price:       	2.95,
			SKU:         	"ice789",
			CreatedOn:   	t.Now().UTC().String(),
			UpdatedOn:   	t.Now().UTC().String(),
	},
	&Product{
			ID:          	7,
			Name:        	"Green Tea",
			Description: 	"Healthy green tea",
			Price:       	2.25,
			SKU:         	"grn123",
			CreatedOn:   	t.Now().UTC().String(),
			UpdatedOn:   	t.Now().UTC().String(),
	},
}