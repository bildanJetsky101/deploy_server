package cartdto

type CartResponse struct {
	ID        int    `json:"id"`
	QTY       int    `json:"qty"`
	ProductID int    `json:"product_id"`
	SubTotal  int    `json:"subtotal"`
	Status    string `json:"status"`
	UserID    int    `json:"user_id"`
}
