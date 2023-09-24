package db

import (
	"context"
	"testing"

	"github.com/nexpictora-pvt-ltd/cnx-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomService(t *testing.T) Service {
	arg := CreateServiceParams{
		ServiceName:  util.RandomUser(),
		ServicePrice: int64(util.RandomPrice()),
	}

	service, err := testQueries.CreateService(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, service)

	require.Equal(t, arg.ServiceName, service.ServiceName)
	require.Equal(t, arg.ServicePrice, service.ServicePrice)

	require.NotZero(t, service.ServiceID)

	return service
}

func TestCreateService(t *testing.T) {
	createRandomService(t)
}

func TestGetService(t *testing.T) {
	service1 := createRandomService(t)
	res, err := testQueries.GetService(context.Background(), service1.ServiceID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, service1.ServiceID, res.ServiceID)
	require.Equal(t, service1.ServiceName, res.ServiceName)
	require.Equal(t, service1.ServicePrice, res.ServicePrice)
}
