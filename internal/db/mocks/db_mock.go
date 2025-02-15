// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "merch_shop/internal/models"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// BuyItemByItemID provides a mock function with given fields: ctx, userID, itemID
func (_m *DB) BuyItemByItemID(ctx context.Context, userID int, itemID int) error {
	ret := _m.Called(ctx, userID, itemID)

	if len(ret) == 0 {
		panic("no return value specified for BuyItemByItemID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, userID, itemID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, username, password
func (_m *DB) CreateUser(ctx context.Context, username string, password string) (*int, error) {
	ret := _m.Called(ctx, username, password)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*int, error)); ok {
		return rf(ctx, username, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *int); ok {
		r0 = rf(ctx, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, username
func (_m *DB) GetUser(ctx context.Context, username string) (*int, string, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 *int
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*int, string, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *int); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) string); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, username)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUserInfoByUserID provides a mock function with given fields: ctx, userID
func (_m *DB) GetUserInfoByUserID(ctx context.Context, userID int) (*int, []byte, *models.CoinTransferHistory, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserInfoByUserID")
	}

	var r0 *int
	var r1 []byte
	var r2 *models.CoinTransferHistory
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*int, []byte, *models.CoinTransferHistory, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *int); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) []byte); ok {
		r1 = rf(ctx, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, int) *models.CoinTransferHistory); ok {
		r2 = rf(ctx, userID)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*models.CoinTransferHistory)
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, int) error); ok {
		r3 = rf(ctx, userID)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// SendCoinByUsername provides a mock function with given fields: ctx, userID, destUsername, amount
func (_m *DB) SendCoinByUsername(ctx context.Context, userID int, destUsername string, amount int) error {
	ret := _m.Called(ctx, userID, destUsername, amount)

	if len(ret) == 0 {
		panic("no return value specified for SendCoinByUsername")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, int) error); ok {
		r0 = rf(ctx, userID, destUsername, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
