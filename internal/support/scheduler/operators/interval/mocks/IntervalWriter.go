// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

// IntervalWriter is an autogenerated mock type for the IntervalWriter type
type IntervalWriter struct {
	mock.Mock
}

// AddInterval provides a mock function with given fields: _a0
func (_m *IntervalWriter) AddInterval(_a0 models.Interval) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Interval) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Interval) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IntervalById provides a mock function with given fields: id
func (_m *IntervalWriter) IntervalById(id string) (models.Interval, error) {
	ret := _m.Called(id)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(string) models.Interval); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IntervalByName provides a mock function with given fields: name
func (_m *IntervalWriter) IntervalByName(name string) (models.Interval, error) {
	ret := _m.Called(name)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(string) models.Interval); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Intervals provides a mock function with given fields:
func (_m *IntervalWriter) Intervals() ([]models.Interval, error) {
	ret := _m.Called()

	var r0 []models.Interval
	if rf, ok := ret.Get(0).(func() []models.Interval); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Interval)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IntervalsWithLimit provides a mock function with given fields: limit
func (_m *IntervalWriter) IntervalsWithLimit(limit int) ([]models.Interval, error) {
	ret := _m.Called(limit)

	var r0 []models.Interval
	if rf, ok := ret.Get(0).(func(int) []models.Interval); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Interval)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
