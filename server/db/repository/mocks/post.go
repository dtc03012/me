// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/dtc03012/me/db/entity"
	mock "github.com/stretchr/testify/mock"

	option "github.com/dtc03012/me/db/option"

	sqlx "github.com/jmoiron/sqlx"
)

// Post is an autogenerated mock type for the Post type
type Post struct {
	mock.Mock
}

// GetBulkPost provides a mock function with given fields: ctx, tx, opt
func (_m *Post) GetBulkPost(ctx context.Context, tx *sqlx.Tx, opt *option.PostOption) ([]*entity.Post, error) {
	ret := _m.Called(ctx, tx, opt)

	var r0 []*entity.Post
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, *option.PostOption) []*entity.Post); ok {
		r0 = rf(ctx, tx, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.Tx, *option.PostOption) error); ok {
		r1 = rf(ctx, tx, opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPost provides a mock function with given fields: ctx, tx, pid
func (_m *Post) GetPost(ctx context.Context, tx *sqlx.Tx, pid int32) (*entity.Post, error) {
	ret := _m.Called(ctx, tx, pid)

	var r0 *entity.Post
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, int32) *entity.Post); ok {
		r0 = rf(ctx, tx, pid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.Tx, int32) error); ok {
		r1 = rf(ctx, tx, pid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertPost provides a mock function with given fields: ctx, tx, post, tags
func (_m *Post) InsertPost(ctx context.Context, tx *sqlx.Tx, post *entity.Post, tags []string) error {
	ret := _m.Called(ctx, tx, post, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, *entity.Post, []string) error); ok {
		r0 = rf(ctx, tx, post, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPost interface {
	mock.TestingT
	Cleanup(func())
}

// NewPost creates a new instance of Post. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPost(t mockConstructorTestingTNewPost) *Post {
	mock := &Post{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}