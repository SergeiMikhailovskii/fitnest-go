package Request

type AddActivityRequest struct {
	Amount *int   `json:"amount" binding:"required"`
	Type   string `json:"type" binding:"required"`
}
