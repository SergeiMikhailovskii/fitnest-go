package Request

type DeleteActivityRequest struct {
	Id   int    `json:"id" binding:"required"`
	Type string `json:"type" binding:"required"`
}
