package products

type ProductRepository interface {
	Create(product *Product) error
}
