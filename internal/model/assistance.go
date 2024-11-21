package model

type AssistanceResponse struct {
    Id          int     `json:"id"` 
    CompanyName string  `json:"company_name"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       int     `json:"price"`
    Deadline    string  `json:"deadline"`
}

type AssistanceDetailedResponse struct {
    Id          int     `json:"id"` 
    User        User    `json:"user"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       int     `json:"price"`
    Deadline    string  `json:"deadline"`
}

type AssistanceRequest struct {
    UserId      int     `json:"user_id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       int     `json:"price"`
    Deadline    string  `json:"deadline"`
}
