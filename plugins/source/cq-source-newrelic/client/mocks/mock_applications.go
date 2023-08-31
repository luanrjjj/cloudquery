package mocks

import (
	context "context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/newrelic/newrelic-client-go/v2/pkg/apm"
	"net/http"
	reflect "reflect"
)

// MockApplicationService is a mock of ApplicationsService interface
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockActionsServiceMockRecorder is the mock recorder for MockActionsService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	fmt.Printf("M VAI: %v+\n", mock)
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// ListApplications mocks base method
func (m *MockApplicationService) ListApplications(arg0 context.Context) (apm.ListApplicationsParams, *http.Response, error) {
	fmt.Printf("PROBLEMA NO MOCKAUSHDUASHDUA: %v+\n", arg0)
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplications", arg0)
	//ListApplications(*apm.ListApplicationsParams) ([]*apm.Application, error)
	ret0, _ := ret[0].(apm.ListApplicationsParams)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[1].(error)

	return ret0, ret1, ret2
}

// ListApplications indicates an expected call of ListApplications
func (mr *MockApplicationServiceMockRecorder) ListApplications(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	fmt.Printf("PROBLEMA NO MOCK1: %v+\n", arg0)
	value := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockApplicationService)(nil).ListApplications), arg0)
	fmt.Printf("PROBLEMA NO MOCK1: %v+\n", value)
	return value
}
