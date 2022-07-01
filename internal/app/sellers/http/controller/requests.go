package controller

type postSellerRequest struct {
	CID         string `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
}

type updateSellerRequest struct {
	CID         string `json:"cid"`
	CompanyName string `json:"company_name""`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type postLocalityRequest struct {
	ID           string `json:"id" binding:"required"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name" binding:"required"`
	CountryName  string `json:"country_name" binding:"required"`
}
