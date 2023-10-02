package db

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func TestOrderTx(t *testing.T) {

// 	status := "Accepted"
// 	store := NewStore(testDB)

// 	user1 := createRandomUser(t)
// 	service1 := createRandomService(t)

// 	n := 5

// 	errs := make(chan error)
// 	results := make(chan OrderTxResult)
// 	for i := 0; i < n; i++ {
// 		go func() {
// 			result, err := store.OrderTx(context.Background(), OrderTxParams{
// 				CustomerID:  int64(user1.UserID),
// 				ServiceID:   int64(service1.ServiceID),
// 				OrderStatus: status,
// 			})
// 			errs <- err
// 			results <- result
// 		}()
// 	}

// 	for i := 0; i < n; i++ {
// 		err := <-errs
// 		require.NoError(t, err)
// 		result := <-results
// 		require.NotEmpty(t, result)

// 		order := result.Order
// 		require.NotEmpty(t, order)
// 		require.Equal(t, user1.UserID, order.CustomerID)
// 		require.Equal(t, service1.ServiceID, order.ServiceID)
// 		require.Equal(t, status, order.OrderStatus)
// 		require.NotZero(t, order.OrderStarted)

// 		_, err = store.GetOrder(context.Background(), order.OrderID)
// 		require.NoError(t, err)

// 	}
// }
