package product

var allProducts = []Product{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Get(idx int) (*Product, error) {
	return &allProducts[idx], nil
}

func (s *Service) List() []Product {
	return allProducts
}
