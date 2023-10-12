package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/nexpictora-pvt-ltd/cnx-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T) Order {
	arg := CreateOrderParams{
		UserID:      int32(util.RandomOrder()),
		ServiceIds:  int32(util.RandomOrder()),
		OrderStatus: util.RandomOrderStatus(),
	}
	// Generate a random delivery time
	deliveryTime := util.RandomDeliveryTime()

	// Check if the deliveryTime is not zero and assign it as a valid sql.NullTime
	var orderDelivered sql.NullTime
	if !deliveryTime.IsZero() {
		orderDelivered.Time = deliveryTime
		orderDelivered.Valid = true
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.UserID, order.UserID)
	require.Equal(t, arg.ServiceIds, order.ServiceIds)
	require.Equal(t, arg.OrderStatus, order.OrderStatus)
	// require.Equal(t, arg.OrderDelivered, order.OrderDelivered.Time)
	require.NotZero(t, order.OrderStarted)
	require.NotZero(t, order.OrderID)

	return order
}

func TestCreateOrder(t *testing.T) {
	createRandomOrder(t)
}

func TestGetOrder(t *testing.T) {
	order1 := createRandomOrder(t)
	res, err := testQueries.GetOrder(context.Background(), order1.OrderID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	// require.Equal(t, order1.OrderID, res.OrderID)
	// require.Equal(t, order1.UserID, res.UserID)
	// require.Equal(t, order1.OrderStatus, res.OrderStatus)
	// require.Equal(t, order1.ServiceIds, res.ServiceIds)
	// require.NotEmpty(t, res.OrderStarted)
	// require.NotEmpty(t, res.OrderDelivered)
}

func TestListOrders(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrder(t)
	}

	arg := ListOrdersParams{
		Limit:  5,
		Offset: 5,
	}

	res, err := testQueries.ListOrders(context.Background(), arg)
	require.NoError(t, err)
	// Here we are checking whether the t object, the res i.e response have a length of 5
	require.Len(t, res, 5)

	//Here we are checking that each account is empty or not, this will ensure that we have got proper data back.
	for _, order := range res {
		require.NotEmpty(t, order)
	}
}

func TestUpdateOrder(t *testing.T) {
	order1 := createRandomOrder(t)

	arg := UpdateOrderParams{
		OrderID:     order1.OrderID,
		OrderStatus: util.RandomOrderStatus(),
	}
	res, err := testQueries.UpdateOrder(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, order1.OrderID, res.OrderID)
	require.Equal(t, arg.OrderStatus, res.OrderStatus)
	require.Equal(t, order1.UserID, res.UserID)
	require.Equal(t, order1.ServiceIds, res.ServiceIds)
	require.WithinDuration(t, order1.OrderStarted, res.OrderStarted, time.Second)
	require.NotEmpty(t, res.OrderDelivered)

}

func TestDeleteOrder(t *testing.T) {
	order1 := createRandomOrder(t)

	err := testQueries.DeleteOrder(context.Background(), order1.OrderID)

	require.NoError(t, err)

	res, err := testQueries.GetOrder(context.Background(), order1.OrderID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
}
