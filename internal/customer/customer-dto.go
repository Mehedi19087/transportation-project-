package customer 


type CreateCustomerRequest struct {
	 CustomerName string `json:"customer_name" binding:"required"`
	 Mobile string `json:"mobile,omitempty"`
	 Email          string `json:"email"`                  
     Address        string `json:"address"`                 
     OpeningBalance string `json:"opening_balance"`  
     Status         string `json:"status"`   
}

type UpdateCustomerRequest struct {
	 CustomerName string `json:"customer_name"`
	 Mobile string `json:"mobile,omitempty"`
	 Email          string `json:"email"`                  
     Address        string `json:"address"`                 
     OpeningBalance string `json:"opening_balance"`  
     Status         string `json:"status"`   
}