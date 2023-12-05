package api

// User represents body of user request and response
type User struct {

	// User's name
	// Required: true
	Name string `json:"name"`
	// User's address
	Address string `json:"address"`
	// User's email
	Email string `json:"email"`
}
