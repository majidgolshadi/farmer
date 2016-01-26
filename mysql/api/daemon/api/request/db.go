package request

type DbRequest struct {
	Database      string 	`json:"database" binding:"required"`
}
