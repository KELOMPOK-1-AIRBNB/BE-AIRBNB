// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	homestay "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays"
	mock "github.com/stretchr/testify/mock"
)

// HomestayData is an autogenerated mock type for the DataInterface type
type HomestayData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id, idUser
func (_m *HomestayData) Delete(id uint, idUser uint) error {
	ret := _m.Called(id, idUser)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(id, idUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetHomestayById provides a mock function with given fields: id
func (_m *HomestayData) GetHomestayById(id uint) (homestay.Core, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetHomestayById")
	}

	var r0 homestay.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (homestay.Core, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) homestay.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(homestay.Core)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHomestayByUserId provides a mock function with given fields: id
func (_m *HomestayData) GetHomestayByUserId(id uint) ([]homestay.Core, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetHomestayByUserId")
	}

	var r0 []homestay.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]homestay.Core, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) []homestay.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]homestay.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMyHomestay provides a mock function with given fields: id
func (_m *HomestayData) GetMyHomestay(id uint) ([]homestay.Core, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetMyHomestay")
	}

	var r0 []homestay.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]homestay.Core, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) []homestay.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]homestay.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByHomestayId provides a mock function with given fields: id
func (_m *HomestayData) GetUserByHomestayId(id uint) (homestay.Core, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByHomestayId")
	}

	var r0 homestay.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (homestay.Core, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) homestay.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(homestay.Core)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: input
func (_m *HomestayData) Insert(input homestay.Core) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(homestay.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MakeHost provides a mock function with given fields: id, input
func (_m *HomestayData) MakeHost(id uint, input homestay.Core) error {
	ret := _m.Called(id, input)

	if len(ret) == 0 {
		panic("no return value specified for MakeHost")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, homestay.Core) error); ok {
		r0 = rf(id, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields: id
func (_m *HomestayData) SelectAll(id uint) ([]homestay.Core, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for SelectAll")
	}

	var r0 []homestay.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]homestay.Core, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) []homestay.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]homestay.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllForUser provides a mock function with given fields:
func (_m *HomestayData) SelectAllForUser() ([]homestay.Core, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SelectAllForUser")
	}

	var r0 []homestay.Core
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]homestay.Core, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []homestay.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]homestay.Core)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, input
func (_m *HomestayData) Update(id uint, input homestay.Core) error {
	ret := _m.Called(id, input)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, homestay.Core) error); ok {
		r0 = rf(id, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewHomestayData creates a new instance of HomestayData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHomestayData(t interface {
	mock.TestingT
	Cleanup(func())
}) *HomestayData {
	mock := &HomestayData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
