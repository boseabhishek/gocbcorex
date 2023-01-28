// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package core

import (
	"github.com/couchbase/stellar-nebula/contrib/cbconfig"
	"sync"
)

// Ensure, that ConfigManagerMock does implement ConfigManager.
// If this is not the case, regenerate this file with moq.
var _ ConfigManager = &ConfigManagerMock{}

// ConfigManagerMock is a mock implementation of ConfigManager.
//
//	func TestSomethingThatUsesConfigManager(t *testing.T) {
//
//		// make and configure a mocked ConfigManager
//		mockedConfigManager := &ConfigManagerMock{
//			ApplyConfigFunc: func(sourceHostname string, json *cbconfig.TerseConfigJson)  {
//				panic("mock out the ApplyConfig method")
//			},
//			RegisterCallbackFunc: func(fn RouteConfigHandler)  {
//				panic("mock out the RegisterCallback method")
//			},
//		}
//
//		// use mockedConfigManager in code that requires ConfigManager
//		// and then make assertions.
//
//	}
type ConfigManagerMock struct {
	// ApplyConfigFunc mocks the ApplyConfig method.
	ApplyConfigFunc func(sourceHostname string, json *cbconfig.TerseConfigJson)

	// RegisterCallbackFunc mocks the RegisterCallback method.
	RegisterCallbackFunc func(fn RouteConfigHandler)

	// calls tracks calls to the methods.
	calls struct {
		// ApplyConfig holds details about calls to the ApplyConfig method.
		ApplyConfig []struct {
			// SourceHostname is the sourceHostname argument value.
			SourceHostname string
			// JSON is the json argument value.
			JSON *cbconfig.TerseConfigJson
		}
		// RegisterCallback holds details about calls to the RegisterCallback method.
		RegisterCallback []struct {
			// Fn is the fn argument value.
			Fn RouteConfigHandler
		}
	}
	lockApplyConfig      sync.RWMutex
	lockRegisterCallback sync.RWMutex
}

// ApplyConfig calls ApplyConfigFunc.
func (mock *ConfigManagerMock) ApplyConfig(sourceHostname string, json *cbconfig.TerseConfigJson) {
	if mock.ApplyConfigFunc == nil {
		panic("ConfigManagerMock.ApplyConfigFunc: method is nil but ConfigManager.ApplyConfig was just called")
	}
	callInfo := struct {
		SourceHostname string
		JSON           *cbconfig.TerseConfigJson
	}{
		SourceHostname: sourceHostname,
		JSON:           json,
	}
	mock.lockApplyConfig.Lock()
	mock.calls.ApplyConfig = append(mock.calls.ApplyConfig, callInfo)
	mock.lockApplyConfig.Unlock()
	mock.ApplyConfigFunc(sourceHostname, json)
}

// ApplyConfigCalls gets all the calls that were made to ApplyConfig.
// Check the length with:
//
//	len(mockedConfigManager.ApplyConfigCalls())
func (mock *ConfigManagerMock) ApplyConfigCalls() []struct {
	SourceHostname string
	JSON           *cbconfig.TerseConfigJson
} {
	var calls []struct {
		SourceHostname string
		JSON           *cbconfig.TerseConfigJson
	}
	mock.lockApplyConfig.RLock()
	calls = mock.calls.ApplyConfig
	mock.lockApplyConfig.RUnlock()
	return calls
}

// RegisterCallback calls RegisterCallbackFunc.
func (mock *ConfigManagerMock) RegisterCallback(fn RouteConfigHandler) {
	if mock.RegisterCallbackFunc == nil {
		panic("ConfigManagerMock.RegisterCallbackFunc: method is nil but ConfigManager.RegisterCallback was just called")
	}
	callInfo := struct {
		Fn RouteConfigHandler
	}{
		Fn: fn,
	}
	mock.lockRegisterCallback.Lock()
	mock.calls.RegisterCallback = append(mock.calls.RegisterCallback, callInfo)
	mock.lockRegisterCallback.Unlock()
	mock.RegisterCallbackFunc(fn)
}

// RegisterCallbackCalls gets all the calls that were made to RegisterCallback.
// Check the length with:
//
//	len(mockedConfigManager.RegisterCallbackCalls())
func (mock *ConfigManagerMock) RegisterCallbackCalls() []struct {
	Fn RouteConfigHandler
} {
	var calls []struct {
		Fn RouteConfigHandler
	}
	mock.lockRegisterCallback.RLock()
	calls = mock.calls.RegisterCallback
	mock.lockRegisterCallback.RUnlock()
	return calls
}
