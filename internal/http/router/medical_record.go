package router

import (
	"net/http"

	"github.com/orvosi/api/internal/http/handler"
)

// MedicalRecordCreator creates routes for medical record creator.
func MedicalRecordCreator(h *handler.MedicalRecordCreator) []*Route {
	var routes []*Route

	r := &Route{
		Method:  http.MethodPost,
		Path:    "/medical-records",
		Handler: h.Create,
	}

	routes = append(routes, r)
	return routes
}
