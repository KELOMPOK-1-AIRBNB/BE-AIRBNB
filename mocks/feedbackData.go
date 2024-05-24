// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	feedback "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/feedback"
	mock "github.com/stretchr/testify/mock"
)

// feedbackData is an autogenerated mock type for the DataInterface type
type FeedbackData struct {
	mock.Mock
}

// CreateFeedback provides a mock function with given fields: input
func (_m *FeedbackData) CreateFeedback(input feedback.Core) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for CreateFeedback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(feedback.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFeedbackByHomestayId provides a mock function with given fields: homestayId
func (_m *FeedbackData) GetFeedbackByHomestayId(homestayId uint) ([]feedback.Core, error) {
	ret := _m.Called(homestayId)

	if len(ret) == 0 {
		panic("no return value specified for GetFeedbackByHomestayId")
	}

	var r0 []feedback.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]feedback.Core, error)); ok {
		return rf(homestayId)
	}
	if rf, ok := ret.Get(0).(func(uint) []feedback.Core); ok {
		r0 = rf(homestayId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]feedback.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(homestayId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newFeedbackData creates a new instance of feedbackData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newFeedbackData(t interface {
	mock.TestingT
	Cleanup(func())
}) *FeedbackData {
	mock := &FeedbackData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
