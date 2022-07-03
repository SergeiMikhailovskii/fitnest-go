package Request

type PinNotificationRequest struct {
	Id  int   `json:"id" binding:"required"`
	Pin *bool `json:"pin" binding:"required"`
}
