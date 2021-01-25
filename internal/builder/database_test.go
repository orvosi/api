package builder_test

import (
	"testing"

	"github.com/orvosi/api/internal/builder"
	"github.com/orvosi/api/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestBuildSQLDatabase(t *testing.T) {
	t.Run("fail to build sql.DB due to unknown driver", func(t *testing.T) {
		cfg, err := config.NewConfig("../../test/fixture/env.valid")
		assert.Nil(t, err)

		db, err := builder.BuildSQLDatabase("unknown", cfg)

		assert.NotNil(t, err)
		assert.Nil(t, db)
	})
}
