// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// DataRetentionInterface is an autogenerated mock type for the DataRetentionInterface type
type DataRetentionInterface struct {
	mock.Mock
}

// AddChannelsToPolicy provides a mock function with given fields: policyID, channelIDs
func (_m *DataRetentionInterface) AddChannelsToPolicy(policyID string, channelIDs []string) *model.AppError {
	ret := _m.Called(policyID, channelIDs)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, []string) *model.AppError); ok {
		r0 = rf(policyID, channelIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// AddTeamsToPolicy provides a mock function with given fields: policyID, teamIDs
func (_m *DataRetentionInterface) AddTeamsToPolicy(policyID string, teamIDs []string) *model.AppError {
	ret := _m.Called(policyID, teamIDs)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, []string) *model.AppError); ok {
		r0 = rf(policyID, teamIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// CreatePolicy provides a mock function with given fields: policy
func (_m *DataRetentionInterface) CreatePolicy(policy *model.RetentionPolicyWithTeamAndChannelIDs) (*model.RetentionPolicyWithTeamAndChannelCounts, *model.AppError) {
	ret := _m.Called(policy)

	var r0 *model.RetentionPolicyWithTeamAndChannelCounts
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.RetentionPolicyWithTeamAndChannelIDs) (*model.RetentionPolicyWithTeamAndChannelCounts, *model.AppError)); ok {
		return rf(policy)
	}
	if rf, ok := ret.Get(0).(func(*model.RetentionPolicyWithTeamAndChannelIDs) *model.RetentionPolicyWithTeamAndChannelCounts); ok {
		r0 = rf(policy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RetentionPolicyWithTeamAndChannelCounts)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.RetentionPolicyWithTeamAndChannelIDs) *model.AppError); ok {
		r1 = rf(policy)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// DeletePolicy provides a mock function with given fields: policyID
func (_m *DataRetentionInterface) DeletePolicy(policyID string) *model.AppError {
	ret := _m.Called(policyID)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(policyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// GetChannelPoliciesForUser provides a mock function with given fields: userID, offset, limit
func (_m *DataRetentionInterface) GetChannelPoliciesForUser(userID string, offset int, limit int) (*model.RetentionPolicyForChannelList, *model.AppError) {
	ret := _m.Called(userID, offset, limit)

	var r0 *model.RetentionPolicyForChannelList
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string, int, int) (*model.RetentionPolicyForChannelList, *model.AppError)); ok {
		return rf(userID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) *model.RetentionPolicyForChannelList); ok {
		r0 = rf(userID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RetentionPolicyForChannelList)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) *model.AppError); ok {
		r1 = rf(userID, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetChannelsForPolicy provides a mock function with given fields: policyID, offset, limit
func (_m *DataRetentionInterface) GetChannelsForPolicy(policyID string, offset int, limit int) (*model.ChannelsWithCount, *model.AppError) {
	ret := _m.Called(policyID, offset, limit)

	var r0 *model.ChannelsWithCount
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string, int, int) (*model.ChannelsWithCount, *model.AppError)); ok {
		return rf(policyID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) *model.ChannelsWithCount); ok {
		r0 = rf(policyID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ChannelsWithCount)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) *model.AppError); ok {
		r1 = rf(policyID, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetGlobalPolicy provides a mock function with given fields:
func (_m *DataRetentionInterface) GetGlobalPolicy() (*model.GlobalRetentionPolicy, *model.AppError) {
	ret := _m.Called()

	var r0 *model.GlobalRetentionPolicy
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func() (*model.GlobalRetentionPolicy, *model.AppError)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.GlobalRetentionPolicy); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GlobalRetentionPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func() *model.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetPolicies provides a mock function with given fields: offset, limit
func (_m *DataRetentionInterface) GetPolicies(offset int, limit int) (*model.RetentionPolicyWithTeamAndChannelCountsList, *model.AppError) {
	ret := _m.Called(offset, limit)

	var r0 *model.RetentionPolicyWithTeamAndChannelCountsList
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(int, int) (*model.RetentionPolicyWithTeamAndChannelCountsList, *model.AppError)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) *model.RetentionPolicyWithTeamAndChannelCountsList); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RetentionPolicyWithTeamAndChannelCountsList)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) *model.AppError); ok {
		r1 = rf(offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetPoliciesCount provides a mock function with given fields:
func (_m *DataRetentionInterface) GetPoliciesCount() (int64, *model.AppError) {
	ret := _m.Called()

	var r0 int64
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func() (int64, *model.AppError)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() *model.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetPolicy provides a mock function with given fields: policyID
func (_m *DataRetentionInterface) GetPolicy(policyID string) (*model.RetentionPolicyWithTeamAndChannelCounts, *model.AppError) {
	ret := _m.Called(policyID)

	var r0 *model.RetentionPolicyWithTeamAndChannelCounts
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string) (*model.RetentionPolicyWithTeamAndChannelCounts, *model.AppError)); ok {
		return rf(policyID)
	}
	if rf, ok := ret.Get(0).(func(string) *model.RetentionPolicyWithTeamAndChannelCounts); ok {
		r0 = rf(policyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RetentionPolicyWithTeamAndChannelCounts)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(policyID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetTeamPoliciesForUser provides a mock function with given fields: userID, offset, limit
func (_m *DataRetentionInterface) GetTeamPoliciesForUser(userID string, offset int, limit int) (*model.RetentionPolicyForTeamList, *model.AppError) {
	ret := _m.Called(userID, offset, limit)

	var r0 *model.RetentionPolicyForTeamList
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string, int, int) (*model.RetentionPolicyForTeamList, *model.AppError)); ok {
		return rf(userID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) *model.RetentionPolicyForTeamList); ok {
		r0 = rf(userID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RetentionPolicyForTeamList)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) *model.AppError); ok {
		r1 = rf(userID, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetTeamsForPolicy provides a mock function with given fields: policyID, offset, limit
func (_m *DataRetentionInterface) GetTeamsForPolicy(policyID string, offset int, limit int) (*model.TeamsWithCount, *model.AppError) {
	ret := _m.Called(policyID, offset, limit)

	var r0 *model.TeamsWithCount
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string, int, int) (*model.TeamsWithCount, *model.AppError)); ok {
		return rf(policyID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) *model.TeamsWithCount); ok {
		r0 = rf(policyID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TeamsWithCount)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) *model.AppError); ok {
		r1 = rf(policyID, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// PatchPolicy provides a mock function with given fields: patch
func (_m *DataRetentionInterface) PatchPolicy(patch *model.RetentionPolicyWithTeamAndChannelIDs) (*model.RetentionPolicyWithTeamAndChannelCounts, *model.AppError) {
	ret := _m.Called(patch)

	var r0 *model.RetentionPolicyWithTeamAndChannelCounts
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.RetentionPolicyWithTeamAndChannelIDs) (*model.RetentionPolicyWithTeamAndChannelCounts, *model.AppError)); ok {
		return rf(patch)
	}
	if rf, ok := ret.Get(0).(func(*model.RetentionPolicyWithTeamAndChannelIDs) *model.RetentionPolicyWithTeamAndChannelCounts); ok {
		r0 = rf(patch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RetentionPolicyWithTeamAndChannelCounts)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.RetentionPolicyWithTeamAndChannelIDs) *model.AppError); ok {
		r1 = rf(patch)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// RemoveChannelsFromPolicy provides a mock function with given fields: policyID, channelIDs
func (_m *DataRetentionInterface) RemoveChannelsFromPolicy(policyID string, channelIDs []string) *model.AppError {
	ret := _m.Called(policyID, channelIDs)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, []string) *model.AppError); ok {
		r0 = rf(policyID, channelIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// RemoveTeamsFromPolicy provides a mock function with given fields: policyID, teamIDs
func (_m *DataRetentionInterface) RemoveTeamsFromPolicy(policyID string, teamIDs []string) *model.AppError {
	ret := _m.Called(policyID, teamIDs)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, []string) *model.AppError); ok {
		r0 = rf(policyID, teamIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

type mockConstructorTestingTNewDataRetentionInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataRetentionInterface creates a new instance of DataRetentionInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataRetentionInterface(t mockConstructorTestingTNewDataRetentionInterface) *DataRetentionInterface {
	mock := &DataRetentionInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
