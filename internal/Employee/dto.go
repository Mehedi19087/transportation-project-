package employee

type CreateEmployeeReq struct {
    FullName       string   `json:"full_name" binding:"required"`
    Email          *string  `json:"email"`
    Mobile         *string  `json:"mobile"`
    BirthDate      *string  `json:"birth_date"`
    Gender         *string  `json:"gender"`
    BloodGroup     *string  `json:"blood_group"`
    Designation    *string  `json:"designation"`
    EmploymentType *string  `json:"employment_type"`
    BranchName     *string  `json:"branch_name"`
    JoinDate       *string  `json:"join_date"`
    Salary         *float64 `json:"salary"`
    PaymentMethod  *string  `json:"payment_method"`
    Nid            *string  `json:"nid"`
    Address        *string  `json:"address"`
    Status         *string  `json:"status"`
    Image          *string  `json:"image"`
    Role           *string  `json:"role"`
}

type UpdateEmployeeReq struct {
    FullName       *string  `json:"full_name"`
    Email          *string  `json:"email"`
    Mobile         *string  `json:"mobile"`
    BirthDate      *string  `json:"birth_date"`
    Gender         *string  `json:"gender"`
    BloodGroup     *string  `json:"blood_group"`
    Designation    *string  `json:"designation"`
    EmploymentType *string  `json:"employment_type"`
    BranchName     *string  `json:"branch_name"`
    JoinDate       *string  `json:"join_date"`
    Salary         *float64 `json:"salary"`
    PaymentMethod  *string  `json:"payment_method"`
    Nid            *string  `json:"nid"`
    Address        *string  `json:"address"`
    Status         *string  `json:"status"`
    Image          *string  `json:"image"`
    Role           *string  `json:"role"`
}