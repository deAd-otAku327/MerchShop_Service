// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Cryptor is an autogenerated mock type for the Cryptor type
type Cryptor struct {
	mock.Mock
}

// CompareHashAndPassword provides a mock function with given fields: hash, password
func (_m *Cryptor) CompareHashAndPassword(hash string, password string) error {
	ret := _m.Called(hash, password)

	if len(ret) == 0 {
		panic("no return value specified for CompareHashAndPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(hash, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EncryptKeyword provides a mock function with given fields: pass
func (_m *Cryptor) EncryptKeyword(pass string) (string, error) {
	ret := _m.Called(pass)

	if len(ret) == 0 {
		panic("no return value specified for EncryptKeyword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(pass)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(pass)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCryptor creates a new instance of Cryptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCryptor(t interface {
	mock.TestingT
	Cleanup(func())
}) *Cryptor {
	mock := &Cryptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
