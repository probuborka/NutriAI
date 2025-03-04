// Code generated by MockGen. DO NOT EDIT.
// Source: /home/user/dev/golang/github/probuborka/NutriAI/internal/usecase/recommendation/recommendation_usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/probuborka/NutriAI/internal/entity"
)

// Mockai is a mock of ai interface.
type Mockai struct {
	ctrl     *gomock.Controller
	recorder *MockaiMockRecorder
}

// MockaiMockRecorder is the mock recorder for Mockai.
type MockaiMockRecorder struct {
	mock *Mockai
}

// NewMockai creates a new mock instance.
func NewMockai(ctrl *gomock.Controller) *Mockai {
	mock := &Mockai{ctrl: ctrl}
	mock.recorder = &MockaiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockai) EXPECT() *MockaiMockRecorder {
	return m.recorder
}

// RecommendationNew mocks base method.
func (m *Mockai) RecommendationNew(userRecommendation entity.UserRecommendationRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecommendationNew", userRecommendation)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecommendationNew indicates an expected call of RecommendationNew.
func (mr *MockaiMockRecorder) RecommendationNew(userRecommendation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecommendationNew", reflect.TypeOf((*Mockai)(nil).RecommendationNew), userRecommendation)
}

// Mockcache is a mock of cache interface.
type Mockcache struct {
	ctrl     *gomock.Controller
	recorder *MockcacheMockRecorder
}

// MockcacheMockRecorder is the mock recorder for Mockcache.
type MockcacheMockRecorder struct {
	mock *Mockcache
}

// NewMockcache creates a new mock instance.
func NewMockcache(ctrl *gomock.Controller) *Mockcache {
	mock := &Mockcache{ctrl: ctrl}
	mock.recorder = &MockcacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockcache) EXPECT() *MockcacheMockRecorder {
	return m.recorder
}

// FindByIDNew mocks base method.
func (m *Mockcache) FindByIDNew(ctx context.Context, id string) (entity.UserRecommendationRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDNew", ctx, id)
	ret0, _ := ret[0].(entity.UserRecommendationRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDNew indicates an expected call of FindByIDNew.
func (mr *MockcacheMockRecorder) FindByIDNew(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDNew", reflect.TypeOf((*Mockcache)(nil).FindByIDNew), ctx, id)
}

// SaveNew mocks base method.
func (m *Mockcache) SaveNew(ctx context.Context, recommendation entity.UserRecommendationRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNew", ctx, recommendation)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNew indicates an expected call of SaveNew.
func (mr *MockcacheMockRecorder) SaveNew(ctx, recommendation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNew", reflect.TypeOf((*Mockcache)(nil).SaveNew), ctx, recommendation)
}
