// Code generated by mockery v1.0.0. DO NOT EDIT.

package account

import (
	entity "github.com/shiningrush/demo-service/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockRepo is an autogenerated mock type for the Repo type
type MockRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: account
func (_m *MockRepo) Create(account entity.Account) (*entity.Account, error) {
	ret := _m.Called(account)

	var r0 *entity.Account
	if rf, ok := ret.Get(0).(func(entity.Account) *entity.Account); ok {
		r0 = rf(account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Account) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *MockRepo) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *MockRepo) Get(id int) (*entity.Account, error) {
	ret := _m.Called(id)

	var r0 *entity.Account
	if rf, ok := ret.Get(0).(func(int) *entity.Account); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWithPhone provides a mock function with given fields: phone
func (_m *MockRepo) GetWithPhone(phone string) (*entity.Account, error) {
	ret := _m.Called(phone)

	var r0 *entity.Account
	if rf, ok := ret.Get(0).(func(string) *entity.Account); ok {
		r0 = rf(phone)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(phone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: account
func (_m *MockRepo) Update(account entity.Account) (*entity.Account, error) {
	ret := _m.Called(account)

	var r0 *entity.Account
	if rf, ok := ret.Get(0).(func(entity.Account) *entity.Account); ok {
		r0 = rf(account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Account) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
