// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RequesterElided is an autogenerated mock type for the RequesterElided type
type RequesterElided struct {
	mock.Mock
}

// Get provides a mock function with given fields: path, url
func (_m *RequesterElided) Get(path string, url string) error {
	ret := _m.Called(path, url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(path, url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
