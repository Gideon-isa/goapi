package models

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float32 `json:"price"`
}

func CreateProduct(products *[]Product, newProduct *Product) error {
	*products = append(*products, *newProduct)

	return nil
}

// ReadProductById
func ReadProductById(products *[]Product, id int) Product {
	for _, item := range *products {
		if item.Id == id {
			return item
		}
	}

	return Product{}
}

// UpdatedProductById
func UpdatedProductById(products *[]Product, updatedProduct *Product) Product {
	for index, item := range *products {

		if item.Id == (*updatedProduct).Id {
			(*products)[index].Name = (*updatedProduct).Name
			(*products)[index].Stock = (*updatedProduct).Stock
			(*products)[index].Price = (*updatedProduct).Price

			return (*products)[index]
		}
	}

	return Product{}
}

//
func DeleteProductById(products *[]Product, id int) Product {
	var deletedProduct Product
	for index, item := range *products {
		if item.Id == id {
			n := len(*products)
			deletedProduct = item // item to be deleted
			// delete the item by:
			// assigning the last element to the index of the said deleted item
			(*products)[index] = (*products)[n-1]

			//then slicing off the last element which is
			// already assigned to the said to deleted item index
			*products = (*products)[:n-1]

			return deletedProduct
		}

	}
	// returning an empty struct
	return Product{}
}
