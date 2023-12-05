package docs

import "github.com/enamespace/tpl/api/swagger/api"

// swagger:route POST /users user createUserRequest
// Create a user in memory
// response:
// 	 200: createUserResponse
//   default: errResponse

// swagger:route GET /users/{name} user getUserRequest
// Get a user from memory
// response:
//   200: getUserResponse
//   default: errResponse

// This text will appear as description of createUserRequest
// swagger:parameters createUserRequest
type userParamsWrapper struct {
	// in:body
	Body api.User
}

// This text will appear as description of getUserRequest
// swagger:parameters getUserRequest
type getUserParamsWarpper struct {
	// in:path
	Name string `json:"name"`
}

// This text will appear as description of createUserResponse
// swagger:response createUserResponse
type createUserResponseWrapper struct {
	// in:body
	Body api.User
}

// This text will appear as description of getUserResponse
// swagger:response getUserResponse
type getUserResponseWrapper struct {
	// in:body
	Body api.User
}

// This text will appear as description of response body
// swagger:response errResponse
type errResponseWrapper struct {
	// Error Code
	Code int `json:"code"`
	// Error Message
	Message string `json:"message"`
}
