package controller

import (
	"api_backend_app/utils"
	"fmt"
	"net/http"
)

// ProtectedEndpoint : Handler for "/protected" Endpoint.
func (c Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ProtectedEndpoint Handler Invoked")
		utils.ResponseJSON(w, "Yes")
	}
}
