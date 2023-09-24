package db

import (
	"context"
	"database/sql"
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

func TestListAllServices(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomService(t)
	}
	res, err := testQueries.ListAllServices(context.Background())
	require.NoError(t, err)
	// Here we are checking whether the t object, the res i.e response have a length of 5

	//Here we are checking that each account is empty or not, this will ensure that we have got proper data back.
	for _, service := range res {
		require.NotEmpty(t, service)
	}
}

func TestListLimitedServices(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomService(t)
	}

	arg := ListLimitedServicesParams{
		Limit:  5,
		Offset: 5,
	}

	res, err := testQueries.ListLimitedServices(context.Background(), arg)
	require.NoError(t, err)
	// Here we are checking whether the t object, the res i.e response have a length of 5
	require.Len(t, res, 5)

	//Here we are checking that each account is empty or not, this will ensure that we have got proper data back.
	for _, service := range res {
		require.NotEmpty(t, service)
	}

}

func TestUpdateService(t *testing.T) {
	service1 := createRandomService(t)

	arg := UpdateServiceParams{
		ServiceID:    service1.ServiceID,
		ServiceName:  util.RandomUser(),
		ServicePrice: 683,
	}

	res, err := testQueries.UpdateService(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, service1.ServiceID, res.ServiceID)
	require.Equal(t, arg.ServiceName, res.ServiceName)
	require.Equal(t, arg.ServicePrice, res.ServicePrice)

}

func TestDeleteService(t *testing.T) {

	service1 := createRandomService(t)

	err := testQueries.DeleteService(context.Background(), service1.ServiceID)

	require.NoError(t, err)

	res, err := testQueries.GetService(context.Background(), service1.ServiceID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
}
