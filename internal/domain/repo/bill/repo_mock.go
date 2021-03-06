// Code generated by mockery v1.0.0. DO NOT EDIT.

package bill

import (
	entity "github.com/shiningrush/demo-service/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockRepo is an autogenerated mock type for the Repo type
type MockRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: account
func (_m *MockRepo) Create(account entity.Bill) (*entity.Bill, error) {
	ret := _m.Called(account)

	var r0 *entity.Bill
	if rf, ok := ret.Get(0).(func(entity.Bill) *entity.Bill); ok {
		r0 = rf(account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Bill)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Bill) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
