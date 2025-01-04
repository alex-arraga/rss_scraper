package services

import (
	"context"
	"database/sql"
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
		apiKey := uuid.New().String()
		mockUser := database.User{
			ID:        uuid.New(),
			Name:      name,
			ApiKey:    apiKey,
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

func TestGetUserByAPIKey(t *testing.T) {
	name := "John Doe"
	apiKey := uuid.New().String()

	// Create new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create db mock and an instance of userService
	mockDB := mocks_services.NewMockUserDatabase(ctrl)
	userService := &UserService{mockDB}

	t.Run("get user by api key successfully", func(t *testing.T) {
		mockUser := database.User{
			ID:        uuid.New(),
			Name:      name,
			ApiKey:    apiKey,
			CreatedAt: time.Now(),
			UpdateAt:  time.Now(),
		}

		// Expected behavior
		mockDB.EXPECT().GetUserByAPIKey(gomock.Any(), gomock.Any()).Return(mockUser, nil)

		// Really call
		result, err := userService.GetUserByAPIKey(context.Background(), apiKey)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.ID, result.ID)
	})

	t.Run("returns error if DB fails", func(t *testing.T) {
		mockDB.EXPECT().GetUserByAPIKey(gomock.Any(), gomock.Any()).Return(database.User{}, errors.New("DB error"))

		result, err := userService.GetUserByAPIKey(context.Background(), apiKey)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestGetPostsForUser(t *testing.T) {
	userID := uuid.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks_services.NewMockUserDatabase(ctrl)
	userService := &UserService{mockDB}

	t.Run("get posts successfully", func(t *testing.T) {
		mockPosts := []database.Post{
			{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       "First title",
				Description: sql.NullString{String: "First description", Valid: true},
				PublishedAt: time.Now(),
				Url:         "First URL",
				FeedID:      uuid.New(),
			},
			{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       "Second title",
				Description: sql.NullString{String: "Second description", Valid: true},
				PublishedAt: time.Now(),
				Url:         "Second URL",
				FeedID:      uuid.New(),
			},
		}

		// Behavior expected
		mockDB.EXPECT().GetPostsForUser(gomock.Any(), gomock.Any()).Return(mockPosts, nil)

		result, err := userService.GetPostsForUser(context.Background(), userID, 10)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(mockPosts), len(result))
		assert.Equal(t, mockPosts[0].ID, result[0].ID)
		assert.Equal(t, mockPosts[0].FeedID, result[0].FeedID)
		assert.Equal(t, mockPosts[0].Url, result[0].URL)
		assert.Equal(t, mockPosts[0].Title, result[0].Title)
	})

	t.Run("returns error if userID is invalid", func(t *testing.T) {
		result, err := userService.GetPostsForUser(context.Background(), uuid.Nil, 10)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
	t.Run("returns error if limit is 0", func(t *testing.T) {
		result, err := userService.GetPostsForUser(context.Background(), userID, 0)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
	t.Run("returns error if limit is a negative number", func(t *testing.T) {
		result, err := userService.GetPostsForUser(context.Background(), userID, -10)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
	t.Run("returns error if userID and limit is invalid", func(t *testing.T) {
		result, err := userService.GetPostsForUser(context.Background(), uuid.Nil, 0)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
