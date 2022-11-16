//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// pkgcrossplaneio
package pkgcrossplaneio

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateProviderRevision_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateProviderRevision_ManifestParameters(props *ProviderRevisionProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateProviderRevision_OfParameters(c constructs.IConstruct) error {
	if c == nil {
		return fmt.Errorf("parameter c is required, but nil was provided")
	}

	return nil
}

func validateNewProviderRevisionParameters(scope constructs.Construct, id *string, props *ProviderRevisionProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

