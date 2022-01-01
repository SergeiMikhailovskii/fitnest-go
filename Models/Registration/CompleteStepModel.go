package Registration

type CompleteStepModel struct {
	Sex         string  `json:"sex" binding:"required"`
	DateOfBirth string  `json:"date_of_birth" binding:"required"`
	Weight      float64 `json:"weight" binding:"required"`
	Height      float64 `json:"height" binding:"required"`
}
