package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository}
}

// GetProduct returns a product by id
// This product was not supposed to be returned. We should return DTO instead.
// However, we will return it for now to keep the example simple.
func (u *ProductUseCase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}
