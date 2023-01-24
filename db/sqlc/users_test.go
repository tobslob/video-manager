package db

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tobslob/video-manager/utils"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(utils.RandomString(8))
	require.NoError(t, err)

	userId, _ := uuid.NewRandom()
	userArg := CreateUserParams{
		ID:             userId,
		UserName:       utils.RandomUsername(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomUsername(),
		Email:          utils.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), userArg)
	if err != nil {
		log.Fatal("Error while creating user:", err)
	}

	require.NotEmpty(t, user)

	require.Equal(t, userArg.FullName, user.FullName)
	require.Equal(t, userArg.Email, user.Email)
	require.Equal(t, userArg.UserName, user.UserName)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.UserName)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.UserName, user2.UserName)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
}
