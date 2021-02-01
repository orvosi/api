package builder

import (
	"database/sql"

	"github.com/orvosi/api/internal/config"
	"github.com/orvosi/api/internal/http/handler"
	"github.com/orvosi/api/internal/http/router"
	"github.com/orvosi/api/internal/repository"
	"github.com/orvosi/api/usecase"
)

// BuildMedicalRecordCreator builds medical record creation workflow
// starting from handler down to repository.
func BuildMedicalRecordCreator(cfg *config.Config, db *sql.DB) []*router.Route {
	ins := repository.NewMedicalRecordInserter(db)
	uc := usecase.NewMedicalRecordCreator(ins)
	hdr := handler.NewMedicalRecordCreator(uc)
	return router.MedicalRecordCreator(hdr)
}

// BuildMedicalRecordFinder builds medical record find workflow
// starting from handler down to repository.
func BuildMedicalRecordFinder(cfg *config.Config, db *sql.DB) []*router.Route {
	sel := repository.NewMedicalRecordSelector(db)
	uc := usecase.NewMedicalRecordFinder(sel)
	hdr := handler.NewMedicalRecordFinder(uc)
	return router.MedicalRecordFinder(hdr)
}

// BuildMedicalRecordUpdater builds medical record update workflow
// starting from handler down to repository.
func BuildMedicalRecordUpdater(cfg *config.Config, db *sql.DB) []*router.Route {
	sel := repository.NewMedicalRecordUpdater(db)
	uc := usecase.NewMedicalRecordUpdater(sel)
	hdr := handler.NewMedicalRecordUpdater(uc)
	return router.MedicalRecordUpdater(hdr)
}
