package response

type DbResponse struct {
	Database      string 	`json:"database_name"`
	Username      string 	`json:"username"`
	Password      string 	`json:"password"`
}
