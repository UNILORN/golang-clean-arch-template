package shop

type shopResponseModel struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type postShopResponse struct {
	Shop shopResponseModel `json:"shop"`
}

type getShopsResponse struct {
	Shop []shopResponseModel `json:"shop"`
}
