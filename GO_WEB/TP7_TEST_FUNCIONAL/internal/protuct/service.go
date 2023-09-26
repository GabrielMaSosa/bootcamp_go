package product

import "gthub.com/GabrielMaSosa/test-funcional/internal/domain"

type ProductService interface {
	ServiceGetAll() (dt []*domain.Product, err error)
	UpdateItem(id int, data domain.Product) (dt *domain.Product, err error)
	Delete(id int) (dt *domain.Product, err error)
	UpdatePartial(id int, data map[string]interface{}) (dt *domain.Product, err error)
}
