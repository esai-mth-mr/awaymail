// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package google

import (
	context "github.com/cloudsrc/api.awaymail.v1.go/src/shared/utils/context"
	gmail "google.golang.org/api/gmail/v1"

	google "github.com/cloudsrc/api.awaymail.v1.go/src/infrastructure/google"

	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"
)

// Wrapper is an autogenerated mock type for the Wrapper type
type Wrapper struct {
	mock.Mock
}

// DeleteMessage provides a mock function with given fields: ctxSess, messageId
func (_m *Wrapper) DeleteMessage(ctxSess *context.Context, messageId string) error {
	ret := _m.Called(ctxSess, messageId)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, string) error); ok {
		r0 = rf(ctxSess, messageId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMessages provides a mock function with given fields: ctxSess, messageId
func (_m *Wrapper) GetMessages(ctxSess *context.Context, messageId string) (*gmail.Message, error) {
	ret := _m.Called(ctxSess, messageId)

	var r0 *gmail.Message
	if rf, ok := ret.Get(0).(func(*context.Context, string) *gmail.Message); ok {
		r0 = rf(ctxSess, messageId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gmail.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctxSess, messageId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessagesList provides a mock function with given fields: ctxSess, nextPageToken, label
func (_m *Wrapper) GetMessagesList(ctxSess *context.Context, nextPageToken string, label string) ([]*gmail.Message, string, error) {
	ret := _m.Called(ctxSess, nextPageToken, label)

	var r0 []*gmail.Message
	if rf, ok := ret.Get(0).(func(*context.Context, string, string) []*gmail.Message); ok {
		r0 = rf(ctxSess, nextPageToken, label)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gmail.Message)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(*context.Context, string, string) string); ok {
		r1 = rf(ctxSess, nextPageToken, label)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*context.Context, string, string) error); ok {
		r2 = rf(ctxSess, nextPageToken, label)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetProfile provides a mock function with given fields: ctxSess, agent, refreshToken
func (_m *Wrapper) GetProfile(ctxSess *context.Context, agent string, refreshToken string) (google.UserProfile, error) {
	ret := _m.Called(ctxSess, agent, refreshToken)

	var r0 google.UserProfile
	if rf, ok := ret.Get(0).(func(*context.Context, string, string) google.UserProfile); ok {
		r0 = rf(ctxSess, agent, refreshToken)
	} else {
		r0 = ret.Get(0).(google.UserProfile)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string, string) error); ok {
		r1 = rf(ctxSess, agent, refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: ctxSess, req
func (_m *Wrapper) SendMessage(ctxSess *context.Context, req google.Messsage) error {
	ret := _m.Called(ctxSess, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, google.Messsage) error); ok {
		r0 = rf(ctxSess, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMessageWithAttachment provides a mock function with given fields: ctxSess, req
func (_m *Wrapper) SendMessageWithAttachment(ctxSess *context.Context, req google.Messsage) error {
	ret := _m.Called(ctxSess, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, google.Messsage) error); ok {
		r0 = rf(ctxSess, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMessage provides a mock function with given fields: ctxSess, id, isRead
func (_m *Wrapper) UpdateMessage(ctxSess *context.Context, id string, isRead *bool) (*gmail.Message, error) {
	ret := _m.Called(ctxSess, id, isRead)

	var r0 *gmail.Message
	if rf, ok := ret.Get(0).(func(*context.Context, string, *bool) *gmail.Message); ok {
		r0 = rf(ctxSess, id, isRead)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gmail.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string, *bool) error); ok {
		r1 = rf(ctxSess, id, isRead)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: ctxSess, agent, refreshToken
func (_m *Wrapper) ValidateToken(ctxSess *context.Context, agent string, refreshToken string) (*oauth2.Token, error) {
	ret := _m.Called(ctxSess, agent, refreshToken)

	var r0 *oauth2.Token
	if rf, ok := ret.Get(0).(func(*context.Context, string, string) *oauth2.Token); ok {
		r0 = rf(ctxSess, agent, refreshToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*context.Context, string, string) error); ok {
		r1 = rf(ctxSess, agent, refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
