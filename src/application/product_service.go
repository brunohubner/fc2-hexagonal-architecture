package application

type ProductService struct {
	Persistence IProductPersistence
}

func NewProductService(persistence IProductPersistence) *ProductService {
	return &ProductService{persistence}
}

func (s *ProductService) Get(id string) (IProduct, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return product, err
}

func (s *ProductService) Create(name string, price float64) (IProduct, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	if _, err := product.IsValid(); err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (s *ProductService) Enable(product IProduct) (IProduct, error) {
	if err := product.Enable(); err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (s *ProductService) Disable(product IProduct) (IProduct, error) {
	if err := product.Disable(); err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
}
