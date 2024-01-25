// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	models "github.com/FrancoLiberali/uala_challenge/app/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// AddFollower provides a mock function with given fields: userID, newFollowerID
func (_m *IRepository) AddFollower(userID uint, newFollowerID uint) error {
	ret := _m.Called(userID, newFollowerID)

	if len(ret) == 0 {
		panic("no return value specified for AddFollower")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, newFollowerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTweet provides a mock function with given fields: tweet
func (_m *IRepository) CreateTweet(tweet models.Tweet) (uuid.UUID, error) {
	ret := _m.Called(tweet)

	if len(ret) == 0 {
		panic("no return value specified for CreateTweet")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Tweet) (uuid.UUID, error)); ok {
		return rf(tweet)
	}
	if rf, ok := ret.Get(0).(func(models.Tweet) uuid.UUID); ok {
		r0 = rf(tweet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Tweet) error); ok {
		r1 = rf(tweet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRepository {
	mock := &IRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
