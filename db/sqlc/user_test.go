package db

import (
	"context"
	"testing"

	"github.com/nexpictora-pvt-ltd/cnx-backend/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
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
}
