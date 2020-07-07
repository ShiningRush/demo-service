// Code generated by mockery v1.0.0. DO NOT EDIT.

package service

import (
	entity "github.com/shiningrush/demo-service/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockAccount is an autogenerated mock type for the Account type
type MockAccount struct {
	mock.Mock
}

// ChangePhone provides a mock function with given fields: accId, newPhone
func (_m *MockAccount) ChangePhone(accId int, newPhone string) error {
	ret := _m.Called(accId, newPhone)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(accId, newPhone)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAccount provides a mock function with given fields: phone
func (_m *MockAccount) NewAccount(phone string) (*entity.Account, error) {
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

// Transfer provides a mock function with given fields: fromAccId, toAccId, amount, desc
func (_m *MockAccount) Transfer(fromAccId int, toAccId int, amount float64, desc string) error {
	ret := _m.Called(fromAccId, toAccId, amount, desc)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, float64, string) error); ok {
		r0 = rf(fromAccId, toAccId, amount, desc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
