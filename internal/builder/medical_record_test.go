package builder_test

import (
	"database/sql"
	"testing"

	"github.com/orvosi/api/internal/builder"
	"github.com/orvosi/api/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestBuildMedicalRecordCreator(t *testing.T) {
	t.Run("successfully build medical record creator", func(t *testing.T) {
		cfg, err := config.NewConfig("../../test/fixture/env.valid")
		assert.Nil(t, err)

		db := &sql.DB{}

		routes := builder.BuildMedicalRecordCreator(cfg, db)
		assert.NotEmpty(t, routes)
	})
}
