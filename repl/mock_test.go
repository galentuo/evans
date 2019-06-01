// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repl

import (
	"sync"

	"github.com/ktr0731/evans/grpc"
	"github.com/ktr0731/evans/idl"
)

var (
	lockSpecMockPackageNames   sync.RWMutex
	lockSpecMockRPC            sync.RWMutex
	lockSpecMockRPCs           sync.RWMutex
	lockSpecMockServiceNames   sync.RWMutex
	lockSpecMockTypeDescriptor sync.RWMutex
)

// Ensure, that SpecMock does implement Spec.
// If this is not the case, regenerate this file with moq.
var _ idl.Spec = &SpecMock{}

// SpecMock is a mock implementation of Spec.
//
//     func TestSomethingThatUsesSpec(t *testing.T) {
//
//         // make and configure a mocked Spec
//         mockedSpec := &SpecMock{
//             PackageNamesFunc: func() []string {
// 	               panic("mock out the PackageNames method")
//             },
//             RPCFunc: func(pkgName string, svcName string, rpcName string) (*grpc.RPC, error) {
// 	               panic("mock out the RPC method")
//             },
//             RPCsFunc: func(pkgName string, svcName string) ([]*grpc.RPC, error) {
// 	               panic("mock out the RPCs method")
//             },
//             ServiceNamesFunc: func(pkgName string) ([]string, error) {
// 	               panic("mock out the ServiceNames method")
//             },
//             TypeDescriptorFunc: func(pkgName string, msgName string) (interface{}, error) {
// 	               panic("mock out the TypeDescriptor method")
//             },
//         }
//
//         // use mockedSpec in code that requires Spec
//         // and then make assertions.
//
//     }
type SpecMock struct {
	// PackageNamesFunc mocks the PackageNames method.
	PackageNamesFunc func() []string

	// RPCFunc mocks the RPC method.
	RPCFunc func(pkgName string, svcName string, rpcName string) (*grpc.RPC, error)

	// RPCsFunc mocks the RPCs method.
	RPCsFunc func(pkgName string, svcName string) ([]*grpc.RPC, error)

	// ServiceNamesFunc mocks the ServiceNames method.
	ServiceNamesFunc func(pkgName string) ([]string, error)

	// TypeDescriptorFunc mocks the TypeDescriptor method.
	TypeDescriptorFunc func(pkgName string, msgName string) (interface{}, error)

	// calls tracks calls to the methods.
	calls struct {
		// PackageNames holds details about calls to the PackageNames method.
		PackageNames []struct {
		}
		// RPC holds details about calls to the RPC method.
		RPC []struct {
			// PkgName is the pkgName argument value.
			PkgName string
			// SvcName is the svcName argument value.
			SvcName string
			// RpcName is the rpcName argument value.
			RpcName string
		}
		// RPCs holds details about calls to the RPCs method.
		RPCs []struct {
			// PkgName is the pkgName argument value.
			PkgName string
			// SvcName is the svcName argument value.
			SvcName string
		}
		// ServiceNames holds details about calls to the ServiceNames method.
		ServiceNames []struct {
			// PkgName is the pkgName argument value.
			PkgName string
		}
		// TypeDescriptor holds details about calls to the TypeDescriptor method.
		TypeDescriptor []struct {
			// PkgName is the pkgName argument value.
			PkgName string
			// MsgName is the msgName argument value.
			MsgName string
		}
	}
}

// PackageNames calls PackageNamesFunc.
func (mock *SpecMock) PackageNames() []string {
	if mock.PackageNamesFunc == nil {
		panic("SpecMock.PackageNamesFunc: method is nil but Spec.PackageNames was just called")
	}
	callInfo := struct {
	}{}
	lockSpecMockPackageNames.Lock()
	mock.calls.PackageNames = append(mock.calls.PackageNames, callInfo)
	lockSpecMockPackageNames.Unlock()
	return mock.PackageNamesFunc()
}

// PackageNamesCalls gets all the calls that were made to PackageNames.
// Check the length with:
//     len(mockedSpec.PackageNamesCalls())
func (mock *SpecMock) PackageNamesCalls() []struct {
} {
	var calls []struct {
	}
	lockSpecMockPackageNames.RLock()
	calls = mock.calls.PackageNames
	lockSpecMockPackageNames.RUnlock()
	return calls
}

// RPC calls RPCFunc.
func (mock *SpecMock) RPC(pkgName string, svcName string, rpcName string) (*grpc.RPC, error) {
	if mock.RPCFunc == nil {
		panic("SpecMock.RPCFunc: method is nil but Spec.RPC was just called")
	}
	callInfo := struct {
		PkgName string
		SvcName string
		RpcName string
	}{
		PkgName: pkgName,
		SvcName: svcName,
		RpcName: rpcName,
	}
	lockSpecMockRPC.Lock()
	mock.calls.RPC = append(mock.calls.RPC, callInfo)
	lockSpecMockRPC.Unlock()
	return mock.RPCFunc(pkgName, svcName, rpcName)
}

// RPCCalls gets all the calls that were made to RPC.
// Check the length with:
//     len(mockedSpec.RPCCalls())
func (mock *SpecMock) RPCCalls() []struct {
	PkgName string
	SvcName string
	RpcName string
} {
	var calls []struct {
		PkgName string
		SvcName string
		RpcName string
	}
	lockSpecMockRPC.RLock()
	calls = mock.calls.RPC
	lockSpecMockRPC.RUnlock()
	return calls
}

// RPCs calls RPCsFunc.
func (mock *SpecMock) RPCs(pkgName string, svcName string) ([]*grpc.RPC, error) {
	if mock.RPCsFunc == nil {
		panic("SpecMock.RPCsFunc: method is nil but Spec.RPCs was just called")
	}
	callInfo := struct {
		PkgName string
		SvcName string
	}{
		PkgName: pkgName,
		SvcName: svcName,
	}
	lockSpecMockRPCs.Lock()
	mock.calls.RPCs = append(mock.calls.RPCs, callInfo)
	lockSpecMockRPCs.Unlock()
	return mock.RPCsFunc(pkgName, svcName)
}

// RPCsCalls gets all the calls that were made to RPCs.
// Check the length with:
//     len(mockedSpec.RPCsCalls())
func (mock *SpecMock) RPCsCalls() []struct {
	PkgName string
	SvcName string
} {
	var calls []struct {
		PkgName string
		SvcName string
	}
	lockSpecMockRPCs.RLock()
	calls = mock.calls.RPCs
	lockSpecMockRPCs.RUnlock()
	return calls
}

// ServiceNames calls ServiceNamesFunc.
func (mock *SpecMock) ServiceNames(pkgName string) ([]string, error) {
	if mock.ServiceNamesFunc == nil {
		panic("SpecMock.ServiceNamesFunc: method is nil but Spec.ServiceNames was just called")
	}
	callInfo := struct {
		PkgName string
	}{
		PkgName: pkgName,
	}
	lockSpecMockServiceNames.Lock()
	mock.calls.ServiceNames = append(mock.calls.ServiceNames, callInfo)
	lockSpecMockServiceNames.Unlock()
	return mock.ServiceNamesFunc(pkgName)
}

// ServiceNamesCalls gets all the calls that were made to ServiceNames.
// Check the length with:
//     len(mockedSpec.ServiceNamesCalls())
func (mock *SpecMock) ServiceNamesCalls() []struct {
	PkgName string
} {
	var calls []struct {
		PkgName string
	}
	lockSpecMockServiceNames.RLock()
	calls = mock.calls.ServiceNames
	lockSpecMockServiceNames.RUnlock()
	return calls
}

// TypeDescriptor calls TypeDescriptorFunc.
func (mock *SpecMock) TypeDescriptor(pkgName string, msgName string) (interface{}, error) {
	if mock.TypeDescriptorFunc == nil {
		panic("SpecMock.TypeDescriptorFunc: method is nil but Spec.TypeDescriptor was just called")
	}
	callInfo := struct {
		PkgName string
		MsgName string
	}{
		PkgName: pkgName,
		MsgName: msgName,
	}
	lockSpecMockTypeDescriptor.Lock()
	mock.calls.TypeDescriptor = append(mock.calls.TypeDescriptor, callInfo)
	lockSpecMockTypeDescriptor.Unlock()
	return mock.TypeDescriptorFunc(pkgName, msgName)
}

// TypeDescriptorCalls gets all the calls that were made to TypeDescriptor.
// Check the length with:
//     len(mockedSpec.TypeDescriptorCalls())
func (mock *SpecMock) TypeDescriptorCalls() []struct {
	PkgName string
	MsgName string
} {
	var calls []struct {
		PkgName string
		MsgName string
	}
	lockSpecMockTypeDescriptor.RLock()
	calls = mock.calls.TypeDescriptor
	lockSpecMockTypeDescriptor.RUnlock()
	return calls
}
