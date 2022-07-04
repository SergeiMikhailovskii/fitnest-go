package Request

type DeleteNotificationRequest struct {
	Id int `json:"id" binding:"required"`
}
