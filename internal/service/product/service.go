package product

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return []Product{
		{Title: "one"},
		{Title: "two"},
		{Title: "three"},
		{Title: "four"},
	}
}
