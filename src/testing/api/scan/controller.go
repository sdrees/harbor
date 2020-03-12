// Code generated by mockery v1.0.0. DO NOT EDIT.

package scan

import (
	artifact "github.com/goharbor/harbor/src/api/artifact"
	all "github.com/goharbor/harbor/src/pkg/scan/all"

	context "context"

	daoscan "github.com/goharbor/harbor/src/pkg/scan/dao/scan"

	job "github.com/goharbor/harbor/src/jobservice/job"

	mock "github.com/stretchr/testify/mock"

	report "github.com/goharbor/harbor/src/pkg/scan/report"

	scan "github.com/goharbor/harbor/src/api/scan"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// DeleteReports provides a mock function with given fields: digests
func (_m *Controller) DeleteReports(digests ...string) error {
	_va := make([]interface{}, len(digests))
	for _i := range digests {
		_va[_i] = digests[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(digests...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetReport provides a mock function with given fields: ctx, _a1, mimeTypes
func (_m *Controller) GetReport(ctx context.Context, _a1 *artifact.Artifact, mimeTypes []string) ([]*daoscan.Report, error) {
	ret := _m.Called(ctx, _a1, mimeTypes)

	var r0 []*daoscan.Report
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact, []string) []*daoscan.Report); ok {
		r0 = rf(ctx, _a1, mimeTypes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*daoscan.Report)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifact.Artifact, []string) error); ok {
		r1 = rf(ctx, _a1, mimeTypes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScanLog provides a mock function with given fields: uuid
func (_m *Controller) GetScanLog(uuid string) ([]byte, error) {
	ret := _m.Called(uuid)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStats provides a mock function with given fields: requester
func (_m *Controller) GetStats(requester string) (*all.Stats, error) {
	ret := _m.Called(requester)

	var r0 *all.Stats
	if rf, ok := ret.Get(0).(func(string) *all.Stats); ok {
		r0 = rf(requester)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*all.Stats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(requester)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSummary provides a mock function with given fields: ctx, _a1, mimeTypes, options
func (_m *Controller) GetSummary(ctx context.Context, _a1 *artifact.Artifact, mimeTypes []string, options ...report.Option) (map[string]interface{}, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, _a1, mimeTypes)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact, []string, ...report.Option) map[string]interface{}); ok {
		r0 = rf(ctx, _a1, mimeTypes, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifact.Artifact, []string, ...report.Option) error); ok {
		r1 = rf(ctx, _a1, mimeTypes, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleJobHooks provides a mock function with given fields: trackID, change
func (_m *Controller) HandleJobHooks(trackID string, change *job.StatusChange) error {
	ret := _m.Called(trackID, change)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *job.StatusChange) error); ok {
		r0 = rf(trackID, change)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scan provides a mock function with given fields: ctx, _a1, options
func (_m *Controller) Scan(ctx context.Context, _a1 *artifact.Artifact, options ...scan.Option) error {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact, ...scan.Option) error); ok {
		r0 = rf(ctx, _a1, options...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
