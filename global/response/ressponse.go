package response

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"messgae"`
	Data    any    `json:"data,omitempty"`
}
