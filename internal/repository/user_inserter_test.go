package repository_test

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/orvosi/api/internal/repository"
	"github.com/stretchr/testify/assert"
)

type UserInserter_Executor struct {
	repo *repository.UserInserter
	sql  sqlmock.Sqlmock
}

func TestNewUserInserter(t *testing.T) {
	t.Run("successfully create an instance of UserInserter", func(t *testing.T) {
		exec := createUserInserterExecutor()
		assert.NotNil(t, exec.repo)
	})
}

func createUserInserterExecutor() *UserInserter_Executor {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Panicf("error opening a stub database connection: %v\n", err)
	}

	repo := repository.NewUserInserter(db)
	return &UserInserter_Executor{
		repo: repo,
		sql:  mock,
	}
}
