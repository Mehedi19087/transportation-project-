package dealer 


type DealerReq struct {
	 DealerName string `json:"dealer_name" binding:"required"`
	 Destination string `json:"destination"`
	 Status string `json:"status"`
}

type DealerUpdateReq struct {
	 DealerName string `json:"dealer_name"`
	 Destination string `json:"destination"`
	 Status string `json:"status"`
}