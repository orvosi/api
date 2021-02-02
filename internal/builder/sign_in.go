package builder

import (
	"database/sql"

	"github.com/orvosi/api/internal/config"
	"github.com/orvosi/api/internal/http/handler"
	"github.com/orvosi/api/internal/http/router"
	"github.com/orvosi/api/internal/repository"
	"github.com/orvosi/api/usecase"
)

// BuildSigner builds sign-in workflow
// starting from handler down to repository.
func BuildSigner(cfg *config.Config, db *sql.DB) []*router.Route {
	ins := repository.NewUserInserter(db)
	uc := usecase.NewSigner(ins)
	hdr := handler.NewSigner(uc)
	return router.Signer(hdr)
}
