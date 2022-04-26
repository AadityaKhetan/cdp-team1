// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "qwik.in/customers-admin/domain/model"
)

// CustomerServiceInterface is an autogenerated mock type for the CustomerServiceInterface type
type CustomerServiceInterface struct {
	mock.Mock
}

// CreateCustomer provides a mock function with given fields: customer
func (_m *CustomerServiceInterface) CreateCustomer(customer model.Customer) (*model.Customer, error) {
	ret := _m.Called(customer)

	var r0 *model.Customer
	if rf, ok := ret.Get(0).(func(model.Customer) *model.Customer); ok {
		r0 = rf(customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Customer) error); ok {
		r1 = rf(customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCustomer provides a mock function with given fields: customerId
func (_m *CustomerServiceInterface) DeleteCustomer(customerId string) (*string, error) {
	ret := _m.Called(customerId)

	var r0 *string
	if rf, ok := ret.Get(0).(func(string) *string); ok {
		r0 = rf(customerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(customerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerByEmail provides a mock function with given fields: customerEmail
func (_m *CustomerServiceInterface) GetCustomerByEmail(customerEmail string) (*model.Customer, error) {
	ret := _m.Called(customerEmail)

	var r0 *model.Customer
	if rf, ok := ret.Get(0).(func(string) *model.Customer); ok {
		r0 = rf(customerEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(customerEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerById provides a mock function with given fields: customerId
func (_m *CustomerServiceInterface) GetCustomerById(customerId string) (*model.Customer, error) {
	ret := _m.Called(customerId)

	var r0 *model.Customer
	if rf, ok := ret.Get(0).(func(string) *model.Customer); ok {
		r0 = rf(customerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(customerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCustomer provides a mock function with given fields: customerId, customer
func (_m *CustomerServiceInterface) UpdateCustomer(customerId string, customer model.Customer) (*model.Customer, error) {
	ret := _m.Called(customerId, customer)

	var r0 *model.Customer
	if rf, ok := ret.Get(0).(func(string, model.Customer) *model.Customer); ok {
		r0 = rf(customerId, customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.Customer) error); ok {
		r1 = rf(customerId, customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
