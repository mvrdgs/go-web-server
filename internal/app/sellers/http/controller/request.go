package controller

type request struct {
	CID         string `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
}

type updateReq struct {
	CID         string `json:"cid"`
	CompanyName string `json:"company_name""`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}
