// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"
	helperstructs "ecommerce/web/helpers/helper_structs"

	mock "github.com/stretchr/testify/mock"

	responce "ecommerce/web/helpers/responce"
)

// AdminUsecaseInterface is an autogenerated mock type for the AdminUsecaseInterface type
type AdminUsecaseInterface struct {
	mock.Mock
}

// AdminLogin provides a mock function with given fields: ctx, adminreq
func (_m *AdminUsecaseInterface) AdminLogin(ctx context.Context, adminreq helperstructs.AdminReq) (responce.AdminData, error) {
	ret := _m.Called(ctx, adminreq)

	var r0 responce.AdminData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, helperstructs.AdminReq) (responce.AdminData, error)); ok {
		return rf(ctx, adminreq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, helperstructs.AdminReq) responce.AdminData); ok {
		r0 = rf(ctx, adminreq)
	} else {
		r0 = ret.Get(0).(responce.AdminData)
	}

	if rf, ok := ret.Get(1).(func(context.Context, helperstructs.AdminReq) error); ok {
		r1 = rf(ctx, adminreq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, email
func (_m *AdminUsecaseInterface) GetUser(ctx context.Context, email string) (responce.AdminsideUsersData, error) {
	ret := _m.Called(ctx, email)

	var r0 responce.AdminsideUsersData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (responce.AdminsideUsersData, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) responce.AdminsideUsersData); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(responce.AdminsideUsersData)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx
func (_m *AdminUsecaseInterface) GetUsers(ctx context.Context) ([]responce.AdminsideUsersData, error) {
	ret := _m.Called(ctx)

	var r0 []responce.AdminsideUsersData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]responce.AdminsideUsersData, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []responce.AdminsideUsersData); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]responce.AdminsideUsersData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reportuser provides a mock function with given fields: ctx, reportreq
func (_m *AdminUsecaseInterface) Reportuser(ctx context.Context, reportreq helperstructs.ReportReq) error {
	ret := _m.Called(ctx, reportreq)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, helperstructs.ReportReq) error); ok {
		r0 = rf(ctx, reportreq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAdminUsecaseInterface creates a new instance of AdminUsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdminUsecaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdminUsecaseInterface {
	mock := &AdminUsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}