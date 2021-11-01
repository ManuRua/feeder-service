package products

type ProductRepository interface {
	Save(product *Product) error
	Count() int
}
