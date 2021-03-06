// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"

	sqlx "github.com/jmoiron/sqlx"
)

// DBService is an autogenerated mock type for the DBService type
type DBService struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields: ctx, opts
func (_m *DBService) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	ret := _m.Called(ctx, opts)

	var r0 *sqlx.Tx
	if rf, ok := ret.Get(0).(func(context.Context, *sql.TxOptions) *sqlx.Tx); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sql.TxOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckAdmin provides a mock function with given fields: ctx, tx, password
func (_m *DBService) CheckAdmin(ctx context.Context, tx *sqlx.Tx, password string) (bool, error) {
	ret := _m.Called(ctx, tx, password)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, string) bool); ok {
		r0 = rf(ctx, tx, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.Tx, string) error); ok {
		r1 = rf(ctx, tx, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDBService interface {
	mock.TestingT
	Cleanup(func())
}

// NewDBService creates a new instance of DBService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDBService(t mockConstructorTestingTNewDBService) *DBService {
	mock := &DBService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
