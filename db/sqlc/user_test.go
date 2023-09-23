package db

import (
	"context"
	"testing"
	"time"

	"github.com/nexpictora-pvt-ltd/cnx-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Name:    util.RandomUser(),
		Email:   util.RandomEmail(),
		Phone:   util.RandomPhone(),
		Address: util.RandomAddress(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Address, user.Address)
	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	res, err := testQueries.GetUser(context.Background(), user1.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, user1.UserID, res.UserID)
	require.Equal(t, user1.Name, res.Name)
	require.Equal(t, user1.Email, res.Email)
	require.Equal(t, user1.Phone, res.Phone)
	require.Equal(t, user1.Address, res.Address)
	require.WithinDuration(t, user1.CreatedAt, res.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		UserID:      user1.UserID,
		Email:       util.RandomEmail(),
		Phone:       user1.Phone,
		Address:     user1.Address,
		TotalOrders: int32(util.RandomOrder()),
	}

	res, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, user1.UserID, res.UserID)
	require.Equal(t, user1.Name, res.Name)
	require.Equal(t, arg.Email, res.Email)
	require.Equal(t, user1.Phone, res.Phone)
	require.Equal(t, user1.Address, res.Address)
	require.Equal(t, arg.TotalOrders, res.TotalOrders)
	require.WithinDuration(t, user1.CreatedAt, res.CreatedAt, time.Second)

}
