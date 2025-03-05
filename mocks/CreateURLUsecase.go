// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	port "github.com/matheusapostulo/url-shortener/internal/url/port"
	mock "github.com/stretchr/testify/mock"
)

// CreateURLUsecase is an autogenerated mock type for the CreateURLUsecase type
type CreateURLUsecase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: input
func (_m *CreateURLUsecase) Execute(input port.CreateURLInputDto) (port.CreateURLOutputDto, error) {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 port.CreateURLOutputDto
	var r1 error
	if rf, ok := ret.Get(0).(func(port.CreateURLInputDto) (port.CreateURLOutputDto, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(port.CreateURLInputDto) port.CreateURLOutputDto); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(port.CreateURLOutputDto)
	}

	if rf, ok := ret.Get(1).(func(port.CreateURLInputDto) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCreateURLUsecase creates a new instance of CreateURLUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreateURLUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreateURLUsecase {
	mock := &CreateURLUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
