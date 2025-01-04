package services

import (
	"context"
	"errors"
	"testing"
	"time"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	mocks_services "github.com/alex-arraga/rss_project/internal/mocks/services"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks_services.NewMockUserDatabase(ctrl)

	userService := &UserService{mockDB}

	t.Run("successfully creates user", func(t *testing.T) {
		name := "John Doe"
		mockUser := database.User{
			ID:        uuid.New(),
			Name:      name,
			CreatedAt: time.Now(),
			UpdateAt:  time.Now(),
		}

		mockDB.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(mockUser, nil)

		result, err := userService.CreateUser(context.Background(), name)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.ID, result.ID)
		assert.Equal(t, name, result.Name)
	})

	t.Run("returns error if name is empty", func(t *testing.T) {
		result, err := userService.CreateUser(context.Background(), "")

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("returns error if DB fails", func(t *testing.T) {
		mockDB.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(database.User{}, errors.New("DB error"))

		result, err := userService.CreateUser(context.Background(), "Jane Doe")

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
