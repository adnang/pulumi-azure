// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management Service Diagnostic.
type Diagnostic struct {
	pulumi.CustomResourceState

	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	ApiManagementLoggerId pulumi.StringOutput `pulumi:"apiManagementLoggerId"`
	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput  `pulumi:"apiManagementName"`
	Enabled           pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewDiagnostic registers a new resource with the given unique name, arguments, and options.
func NewDiagnostic(ctx *pulumi.Context,
	name string, args *DiagnosticArgs, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	if args == nil || args.ApiManagementLoggerId == nil {
		return nil, errors.New("missing required argument 'ApiManagementLoggerId'")
	}
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.Identifier == nil {
		return nil, errors.New("missing required argument 'Identifier'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DiagnosticArgs{}
	}
	var resource Diagnostic
	err := ctx.RegisterResource("azure:apimanagement/diagnostic:Diagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiagnostic gets an existing Diagnostic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticState, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	var resource Diagnostic
	err := ctx.ReadResource("azure:apimanagement/diagnostic:Diagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Diagnostic resources.
type diagnosticState struct {
	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	ApiManagementLoggerId *string `pulumi:"apiManagementLoggerId"`
	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	Enabled           *bool   `pulumi:"enabled"`
	// The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
	Identifier *string `pulumi:"identifier"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type DiagnosticState struct {
	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	ApiManagementLoggerId pulumi.StringPtrInput
	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	Enabled           pulumi.BoolPtrInput
	// The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
	Identifier pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (DiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticState)(nil)).Elem()
}

type diagnosticArgs struct {
	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	ApiManagementLoggerId string `pulumi:"apiManagementLoggerId"`
	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	Enabled           *bool  `pulumi:"enabled"`
	// The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
	Identifier string `pulumi:"identifier"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Diagnostic resource.
type DiagnosticArgs struct {
	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	ApiManagementLoggerId pulumi.StringInput
	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	Enabled           pulumi.BoolPtrInput
	// The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
	Identifier pulumi.StringInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (DiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticArgs)(nil)).Elem()
}
