// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Group_tables is an autogenerated mock type for the Group_tables type
type Group_tables struct {
	mock.Mock
}

// Migrate_me provides a mock function with given fields:
func (_m *Group_tables) Migrate_me() {
	_m.Called()
}

// NewGroup_tables creates a new instance of Group_tables. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGroup_tables(t interface {
	mock.TestingT
	Cleanup(func())
}) *Group_tables {
	mock := &Group_tables{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
