// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PublisherConnection is an autogenerated mock type for the PublisherConnection type
type PublisherConnection struct {
	mock.Mock
}

// Close provides a mock function with no fields
func (_m *PublisherConnection) Close() {
	_m.Called()
}

// PublishMsg provides a mock function with given fields: msg
func (_m *PublisherConnection) PublishMsg(msg []byte) error {
	ret := _m.Called(msg)

	if len(ret) == 0 {
		panic("no return value specified for PublishMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublisherConfig provides a mock function with given fields: name
func (_m *PublisherConnection) PublisherConfig(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for PublisherConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPublisherConnection creates a new instance of PublisherConnection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPublisherConnection(t interface {
	mock.TestingT
	Cleanup(func())
}) *PublisherConnection {
	mock := &PublisherConnection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
