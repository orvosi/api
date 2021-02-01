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

// MedicalRecordFinder creates routes for medical record finder.
func MedicalRecordFinder(h *handler.MedicalRecordFinder) []*Route {
	var routes []*Route

	r := &Route{
		Method:  http.MethodGet,
		Path:    "/medical-records",
		Handler: h.FindByEmail,
	}

	routes = append(routes, r)
	return routes
}

// MedicalRecordUpdater creates routes for medical record updater.
func MedicalRecordUpdater(h *handler.MedicalRecordUpdater) []*Route {
	var routes []*Route

	r := &Route{
		Method:  http.MethodPut,
		Path:    "/medical-records/:id",
		Handler: h.Update,
	}

	routes = append(routes, r)
	return routes
}
