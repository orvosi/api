package repository_test

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/orvosi/api/entity"
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

func TestUserInserter_InsertOrIgnore(t *testing.T) {
	t.Run("can't proceed due to nil user", func(t *testing.T) {
		exec := createUserInserterExecutor()

		err := exec.repo.InsertOrIgnore(context.Background(), nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyUser, err)
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
