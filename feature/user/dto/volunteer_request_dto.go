package dto

type VoluteerRequestCreatingDTO struct {
	UserID int    `json:"user_id" binding:"required"`
	Type   string `json:"type" binding:"required"`
	Status int    `json:"status" binding:"required"`
}
