package test

import (
	"context"
	"testing"
	"time"

	"github.com/OzkrOssa/idktd/internal/users/core/domain"
	"github.com/OzkrOssa/idktd/internal/users/core/port/mock"
	"github.com/OzkrOssa/idktd/internal/users/core/service"
	"github.com/OzkrOssa/idktd/internal/users/core/util"
	mock_cache "github.com/OzkrOssa/idktd/pkg/storage/cache/mock"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

type registerInput struct {
	user *domain.User
}

type expectedOutput struct {
	user *domain.User
	err  error
}

func TestUserService_Register(t *testing.T) {
	ctx := context.Background()
	email := gofakeit.Email()
	name := gofakeit.Name()
	password := gofakeit.Password(true, true, true, true, true, 10)
	hashedPassword, _ := util.HashPassword(password)

	userInput := &domain.User{
		Email:    email,
		Name:     name,
		Password: password,
	}
	userOutput := &domain.User{
		ID:        gofakeit.Uint64(),
		Email:     email,
		Name:      name,
		Password:  hashedPassword,
		Role:      domain.RoleReader,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	serializedUser, _ := util.Serialize(userOutput)
	cacheKey := util.GenerateCacheKey("user", userOutput.ID)
	ttl := time.Duration(0)

	testCases := []struct {
		desc     string
		mocks    func(repo *mock.UserRepository, cache *mock_cache.CacheRepository)
		input    registerInput
		expected expectedOutput
	}{
		{
			desc: "Success",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().CreateUser(ctx, userInput).Return(userOutput, nil)
				cache.EXPECT().Set(ctx, cacheKey, serializedUser, ttl).Return(nil)
				cache.EXPECT().DeleteByPrefix(ctx, "users:*").Return(nil)
			},
			input: registerInput{user: userInput},
			expected: expectedOutput{
				user: userOutput,
				err:  nil,
			},
		},
		{
			desc: "Fail_DuplicateData",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().CreateUser(ctx, userInput).Return(nil, domain.ErrConflictingData)
			},
			input: registerInput{user: userInput},
			expected: expectedOutput{
				user: nil,
				err:  domain.ErrConflictingData,
			},
		},
		{
			desc: "Fail_ErrInternal",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().CreateUser(ctx, userInput).Return(nil, domain.ErrInternal)
			},
			input: registerInput{user: userInput},
			expected: expectedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
		},
		{
			desc: "Fail_SetCache",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().CreateUser(ctx, userInput).Return(userOutput, nil)
				cache.EXPECT().Set(ctx, cacheKey, serializedUser, ttl).Return(domain.ErrInternal)
			},
			input: registerInput{user: userInput},
			expected: expectedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
		},
		{
			desc: "Fail_DeleteCacheByPrefix",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().CreateUser(ctx, userInput).Return(userOutput, nil)
				cache.EXPECT().Set(ctx, cacheKey, serializedUser, ttl).Return(nil)
				cache.EXPECT().DeleteByPrefix(ctx, "users:*").Return(domain.ErrInternal)
			},
			input: registerInput{user: userInput},
			expected: expectedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			repo := mock.NewUserRepository(t)
			cache := mock_cache.NewCacheRepository(t)
			tc.mocks(repo, cache)

			userService := service.NewUserService(repo, cache)

			user, err := userService.Register(ctx, tc.input.user)
			assert.Equal(t, tc.expected.err, err, "Error mismatch")
			assert.Equal(t, tc.expected.user, user, "User mismatch")
		})
	}

}

type getUserTestedInput struct {
	ID uint64
}

type getUserExpectedOutput struct {
	user *domain.User
	err  error
}

func TestUserService_GetUser(t *testing.T) {
	ctx := context.Background()
	id := gofakeit.Uint64()

	userOutput := &domain.User{
		ID:       id,
		Email:    gofakeit.Email(),
		Name:     gofakeit.Name(),
		Role:     domain.RoleReader,
		Password: gofakeit.Password(true, true, true, true, true, 10),
	}

	cacheKey := util.GenerateCacheKey("user", id)
	userSerialized, _ := util.Serialize(userOutput)
	ttl := time.Duration(0)

	testCases := []struct {
		desc     string
		mocks    func(repo *mock.UserRepository, cache *mock_cache.CacheRepository)
		input    getUserTestedInput
		expected getUserExpectedOutput
	}{
		{
			desc: "Success_FromCache",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				cache.EXPECT().Get(ctx, cacheKey).Return(userSerialized, nil)
			},
			input: getUserTestedInput{ID: id},
			expected: getUserExpectedOutput{
				user: userOutput,
				err:  nil,
			},
		},
		{
			desc: "Success_FromDB",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				cache.EXPECT().Get(ctx, cacheKey).Return(nil, domain.ErrInternal)
				repo.EXPECT().GetUserByID(ctx, id).Return(userOutput, nil)
				cache.EXPECT().Set(ctx, cacheKey, userSerialized, ttl).Return(nil)
			},
			input: getUserTestedInput{ID: id},
			expected: getUserExpectedOutput{
				user: userOutput,
				err:  nil,
			},
		},
		{
			desc: "Fail_NotFound",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				cache.EXPECT().Get(ctx, cacheKey).Return(nil, domain.ErrDataNotFound)
				repo.EXPECT().GetUserByID(ctx, id).Return(nil, domain.ErrDataNotFound)
			},
			input: getUserTestedInput{ID: id},
			expected: getUserExpectedOutput{
				user: nil,
				err:  domain.ErrDataNotFound,
			},
		},
		{
			desc: "Fail_ErrInternal",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				cache.EXPECT().Get(ctx, cacheKey).Return(nil, domain.ErrInternal)
				repo.EXPECT().GetUserByID(ctx, id).Return(nil, domain.ErrInternal)
			},
			input: getUserTestedInput{ID: id},
			expected: getUserExpectedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
		},
		{
			desc: "Fail_SetCache",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				cache.EXPECT().Get(ctx, cacheKey).Return(nil, domain.ErrInternal)
				repo.EXPECT().GetUserByID(ctx, id).Return(userOutput, nil)
				cache.EXPECT().Set(ctx, cacheKey, userSerialized, ttl).Return(domain.ErrInternal)
			},
			input: getUserTestedInput{ID: id},
			expected: getUserExpectedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
		},
		{
			desc: "Fail_Deserialize",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				cache.EXPECT().Get(ctx, cacheKey).Return([]byte("invalid"), nil)
			},
			input: getUserTestedInput{ID: id},
			expected: getUserExpectedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			repo := mock.NewUserRepository(t)
			cache := mock_cache.NewCacheRepository(t)
			tc.mocks(repo, cache)

			userService := service.NewUserService(repo, cache)

			user, err := userService.GetUser(ctx, id)
			assert.Equal(t, tc.expected.err, err, "Error mismatch")
			assert.Equal(t, tc.expected.user, user, "User mismatch")
		})
	}
}

type listUsersTestedInput struct {
	skip  uint64
	limit uint64
}

type listUsersExpectedOutput struct {
	users []domain.User
	err   error
}

func TestUserService_ListUsers(t *testing.T) {
	ctx := context.Background()
	skip, limit := gofakeit.Uint64(), gofakeit.Uint64()
	params := util.GenerateCacheKeyParams(skip, limit)
	cacheKey := util.GenerateCacheKey("users", params)

	var users []domain.User

	for i := 0; i < 10; i++ {
		userPassword := gofakeit.Password(true, true, true, true, false, 8)
		hashedPassword, _ := util.HashPassword(userPassword)

		users = append(users, domain.User{
			ID:       gofakeit.Uint64(),
			Name:     gofakeit.Name(),
			Email:    gofakeit.Email(),
			Role:     domain.RoleReader,
			Password: hashedPassword,
		})
	}

	usersSerialized, _ := util.Serialize(users)
	ttl := time.Duration(0)

	testCases := []struct {
		desc     string
		mocks    func(repo *mock.UserRepository, cache *mock_cache.CacheRepository)
		input    listUsersTestedInput
		expected listUsersExpectedOutput
	}{
		{
			desc: "Success_FromCache",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				cache.EXPECT().Get(ctx, cacheKey).Return(usersSerialized, nil)
			},
			input: listUsersTestedInput{
				skip:  skip,
				limit: limit,
			},
			expected: listUsersExpectedOutput{
				users: users,
				err:   nil,
			},
		},
		{
			desc: "Success_FromDB",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				cache.EXPECT().Get(ctx, cacheKey).Return(nil, domain.ErrDataNotFound)
				repo.EXPECT().ListUsers(ctx, skip, limit).Return(users, nil)
				cache.EXPECT().Set(ctx, cacheKey, usersSerialized, ttl).Return(nil)
			},
			input: listUsersTestedInput{
				skip:  skip,
				limit: limit,
			},
			expected: listUsersExpectedOutput{
				users: users,
				err:   nil,
			},
		},
		{
			desc: "Fail_ErrInternal",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				cache.EXPECT().Get(ctx, cacheKey).Return(nil, domain.ErrDataNotFound)
				repo.EXPECT().ListUsers(ctx, skip, limit).Return(nil, domain.ErrInternal)
			},
			input: listUsersTestedInput{
				skip:  skip,
				limit: limit,
			},
			expected: listUsersExpectedOutput{
				users: nil,
				err:   domain.ErrInternal,
			},
		},
		{
			desc: "Fail_Deserialize",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				cache.EXPECT().Get(ctx, cacheKey).Return([]byte("invalid"), nil)
			},
			input: listUsersTestedInput{
				skip:  skip,
				limit: limit,
			},
			expected: listUsersExpectedOutput{
				users: nil,
				err:   domain.ErrInternal,
			},
		},
		{
			desc: "Fail_SetCache",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				cache.EXPECT().Get(ctx, cacheKey).Return(nil, domain.ErrDataNotFound)
				repo.EXPECT().ListUsers(ctx, skip, limit).Return(users, nil)
				cache.EXPECT().Set(ctx, cacheKey, usersSerialized, ttl).Return(domain.ErrInternal)
			},
			input: listUsersTestedInput{
				skip:  skip,
				limit: limit,
			},
			expected: listUsersExpectedOutput{
				users: nil,
				err:   domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			repo := mock.NewUserRepository(t)
			cache := mock_cache.NewCacheRepository(t)
			tc.mocks(repo, cache)

			userService := service.NewUserService(repo, cache)

			users, err := userService.ListUsers(ctx, tc.input.skip, tc.input.limit)
			assert.Equal(t, tc.expected.err, err, "Error mismatch")
			assert.Equal(t, tc.expected.users, users, "Users mismatch")
		})
	}
}

type updateUserTestedInput struct {
	user *domain.User
}

type updateUserExpectedOutput struct {
	user *domain.User
	err  error
}

func TestUserService_UpdateUser(t *testing.T) {
	ctx := context.Background()
	id := gofakeit.Uint64()

	userInput := &domain.User{
		ID:    id,
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
		Role:  domain.RoleReader,
	}

	userOutput := &domain.User{
		ID:    id,
		Name:  userInput.Name,
		Email: userInput.Email,
		Role:  domain.RoleReader,
	}

	existingUser := &domain.User{
		ID:    id,
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
		Role:  domain.RoleReader,
	}

	cacheKey := util.GenerateCacheKey("user", id)
	userSerialized, _ := util.Serialize(userOutput)
	ttl := time.Duration(0)

	testCases := []struct {
		desc     string
		mocks    func(repo *mock.UserRepository, cache *mock_cache.CacheRepository)
		input    updateUserTestedInput
		expected updateUserExpectedOutput
	}{
		{
			desc: "Success",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(existingUser, nil)
				repo.EXPECT().UpdateUser(ctx, userInput).Return(userOutput, nil)
				cache.EXPECT().Delete(ctx, cacheKey).Return(nil)
				cache.EXPECT().Set(ctx, cacheKey, userSerialized, ttl).Return(nil)
				cache.EXPECT().DeleteByPrefix(ctx, "users:*").Return(nil)

			},
			input: updateUserTestedInput{
				user: userInput,
			},
			expected: updateUserExpectedOutput{
				user: userOutput,
				err:  nil,
			},
		},
		{
			desc: "Fail_NotFound",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(nil, domain.ErrDataNotFound)

			},
			input: updateUserTestedInput{
				user: userInput,
			},
			expected: updateUserExpectedOutput{
				user: nil,
				err:  domain.ErrDataNotFound,
			},
		},
		{
			desc: "Fail_ErrInternalGetById",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(nil, domain.ErrInternal)

			},
			input: updateUserTestedInput{
				user: userInput,
			},
			expected: updateUserExpectedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
		},
		{
			desc: "Fail_EmptyData",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(existingUser, nil)
			},
			input: updateUserTestedInput{
				user: &domain.User{
					ID: id,
				},
			},
			expected: updateUserExpectedOutput{
				user: nil,
				err:  domain.ErrNoUpdatedData,
			},
		},
		{
			desc: "Fail_SameData",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(existingUser, nil)
			},
			input: updateUserTestedInput{
				user: existingUser,
			},
			expected: updateUserExpectedOutput{
				user: nil,
				err:  domain.ErrNoUpdatedData,
			},
		},
		{
			desc: "Fail_DuplicateData",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(existingUser, nil)
				repo.EXPECT().UpdateUser(ctx, userInput).Return(nil, domain.ErrConflictingData)
			},
			input: updateUserTestedInput{
				user: userInput,
			},
			expected: updateUserExpectedOutput{
				user: nil,
				err:  domain.ErrConflictingData,
			},
		},
		{
			desc: "Fail_ErrInternalUpdate",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(existingUser, nil)
				repo.EXPECT().UpdateUser(ctx, userInput).Return(nil, domain.ErrInternal)
			},
			input: updateUserTestedInput{
				user: userInput,
			},
			expected: updateUserExpectedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
		},
		{
			desc: "Fail_DeleteCache",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(existingUser, nil)
				repo.EXPECT().UpdateUser(ctx, userInput).Return(userOutput, nil)
				cache.EXPECT().Delete(ctx, cacheKey).Return(domain.ErrInternal)
			},
			input: updateUserTestedInput{
				user: userInput,
			},
			expected: updateUserExpectedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
		},
		{
			desc: "Fail_SetCache",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(existingUser, nil)
				repo.EXPECT().UpdateUser(ctx, userInput).Return(userOutput, nil)
				cache.EXPECT().Delete(ctx, cacheKey).Return(nil)
				cache.EXPECT().Set(ctx, cacheKey, userSerialized, ttl).Return(domain.ErrInternal)
			},
			input: updateUserTestedInput{
				user: userInput,
			},
			expected: updateUserExpectedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
		},
		{
			desc: "Fail_DeleteByPrefix",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(existingUser, nil)
				repo.EXPECT().UpdateUser(ctx, userInput).Return(userOutput, nil)
				cache.EXPECT().Delete(ctx, cacheKey).Return(nil)
				cache.EXPECT().Set(ctx, cacheKey, userSerialized, ttl).Return(nil)
				cache.EXPECT().DeleteByPrefix(ctx, "users:*").Return(domain.ErrInternal)
			},
			input: updateUserTestedInput{
				user: userInput,
			},
			expected: updateUserExpectedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			repo := mock.NewUserRepository(t)
			cache := mock_cache.NewCacheRepository(t)
			tc.mocks(repo, cache)
			userService := service.NewUserService(repo, cache)

			user, err := userService.UpdateUser(ctx, tc.input.user)

			assert.Equal(t, tc.expected.err, err, "Error mismatch")
			assert.Equal(t, tc.expected.user, user, "Users mismatch")
		})
	}
}

type userDeleteExpectedOutput struct {
	err error
}

func TestUserService_DeleteUser(t *testing.T) {
	ctx := context.Background()
	id := gofakeit.Uint64()

	cacheKey := util.GenerateCacheKey("user", id)

	testCases := []struct {
		desc     string
		mocks    func(repo *mock.UserRepository, cache *mock_cache.CacheRepository)
		input    uint64
		expected userDeleteExpectedOutput
	}{
		{
			desc: "Success",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(&domain.User{}, nil)
				cache.EXPECT().Delete(ctx, cacheKey).Return(nil)
				cache.EXPECT().DeleteByPrefix(ctx, "users:*").Return(nil)
				repo.EXPECT().DeleteUser(ctx, id).Return(nil)
			},
			input: id,
			expected: userDeleteExpectedOutput{
				err: nil,
			},
		},
		{
			desc: "Fail_NotFound",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(nil, domain.ErrDataNotFound)
			},
			input: id,
			expected: userDeleteExpectedOutput{
				err: domain.ErrDataNotFound,
			},
		},
		{
			desc: "Fail_ErrInternalGetByID",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(nil, domain.ErrInternal)
			},
			input: id,
			expected: userDeleteExpectedOutput{
				err: domain.ErrInternal,
			},
		},
		{
			desc: "Fail_DeleteCache",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(&domain.User{}, nil)
				cache.EXPECT().Delete(ctx, cacheKey).Return(domain.ErrInternal)
			},
			input: id,
			expected: userDeleteExpectedOutput{
				err: domain.ErrInternal,
			},
		},
		{
			desc: "Fail_DeleteByPrefix",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(&domain.User{}, nil)
				cache.EXPECT().Delete(ctx, cacheKey).Return(nil)
				cache.EXPECT().DeleteByPrefix(ctx, "users:*").Return(domain.ErrInternal)
			},
			input: id,
			expected: userDeleteExpectedOutput{
				err: domain.ErrInternal,
			},
		},
		{
			desc: "Fail_ErrInternalDelete",
			mocks: func(repo *mock.UserRepository, cache *mock_cache.CacheRepository) {
				repo.EXPECT().GetUserByID(ctx, id).Return(&domain.User{}, nil)
				cache.EXPECT().Delete(ctx, cacheKey).Return(nil)
				cache.EXPECT().DeleteByPrefix(ctx, "users:*").Return(nil)
				repo.EXPECT().DeleteUser(ctx, id).Return(domain.ErrInternal)
			},
			input: id,
			expected: userDeleteExpectedOutput{
				err: domain.ErrInternal,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			repo := mock.NewUserRepository(t)
			cache := mock_cache.NewCacheRepository(t)
			tc.mocks(repo, cache)
			userService := service.NewUserService(repo, cache)

			err := userService.DeleteUser(ctx, tc.input)

			assert.Equal(t, tc.expected.err, err, "Error mismatch")
		})
	}

}
