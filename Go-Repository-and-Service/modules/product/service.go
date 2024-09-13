package product

type Service interface {
	CreateProduct(product *Product) error
	GetAllProducts() ([]Product, error)
	GetProductById(id int) (Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id int) error
}

type productService struct {
	repo Repository
}

func NewProductService(repo Repository) Service {
	return &productService{repo}
}

func (c *productService) CreateProduct(product *Product) error {
	return c.repo.Create(product)
}

func (c *productService) GetAllProducts() ([]Product, error) {
	return c.repo.FindAll()
}

func (c *productService) GetProductById(id int) (Product, error) {
	return c.repo.FindOne(id)
}

func (c *productService) UpdateProduct(product *Product) error {
	return c.repo.Update(product)
}

func (c *productService) DeleteProduct(id int) error {
	return c.repo.Delete(id)
}
