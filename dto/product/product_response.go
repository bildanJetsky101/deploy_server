package productdto

type ProductResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" form:"title"`
	Price int    `json:"price" form:"price"`
	Stock int    `json:"stock" form:"stock"`
	Desc  string `json:"desc"  form:"desc"`
	Image string `json:"image" form:"image"`
}
