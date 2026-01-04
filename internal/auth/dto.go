package auth


type AuthReq struct {
	 Name  string  `json:"name" binding:"required"`
	 Password string `json:"password" binding:"required"`
     ProductID *uint `json:"product_id"`
}

type UpdateReq struct {
	 Name  *string `json:"name"`
	 Password *string `json:"password"`
     Status   *string  `json:"status"`
}