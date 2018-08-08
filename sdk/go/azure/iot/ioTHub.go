// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a IotHub
type IoTHub struct {
	s *pulumi.ResourceState
}

// NewIoTHub registers a new resource with the given unique name, arguments, and options.
func NewIoTHub(ctx *pulumi.Context,
	name string, args *IoTHubArgs, opts ...pulumi.ResourceOpt) (*IoTHub, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["sku"] = nil
		inputs["tags"] = nil
	} else {
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["sku"] = args.Sku
		inputs["tags"] = args.Tags
	}
	inputs["hostname"] = nil
	inputs["sharedAccessPolicies"] = nil
	inputs["type"] = nil
	s, err := ctx.RegisterResource("azure:iot/ioTHub:IoTHub", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IoTHub{s: s}, nil
}

// GetIoTHub gets an existing IoTHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIoTHub(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IoTHubState, opts ...pulumi.ResourceOpt) (*IoTHub, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["hostname"] = state.Hostname
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["sharedAccessPolicies"] = state.SharedAccessPolicies
		inputs["sku"] = state.Sku
		inputs["tags"] = state.Tags
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("azure:iot/ioTHub:IoTHub", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IoTHub{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IoTHub) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IoTHub) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The hostname of the IotHub Resource.
func (r *IoTHub) Hostname() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["hostname"])
}

// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
func (r *IoTHub) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The name of the sku. Possible values are `F1`, `S1`, `S2`, and `S3`.
func (r *IoTHub) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
func (r *IoTHub) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A list of `shared_access_policy` blocks as defined below.
func (r *IoTHub) SharedAccessPolicies() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["sharedAccessPolicies"])
}

// A `sku` block as defined below. 
func (r *IoTHub) Sku() *pulumi.Output {
	return r.s.State["sku"]
}

// A mapping of tags to assign to the resource.
func (r *IoTHub) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

func (r *IoTHub) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering IoTHub resources.
type IoTHubState struct {
	// The hostname of the IotHub Resource.
	Hostname interface{}
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the sku. Possible values are `F1`, `S1`, `S2`, and `S3`.
	Name interface{}
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A list of `shared_access_policy` blocks as defined below.
	SharedAccessPolicies interface{}
	// A `sku` block as defined below. 
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	Type interface{}
}

// The set of arguments for constructing a IoTHub resource.
type IoTHubArgs struct {
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the sku. Possible values are `F1`, `S1`, `S2`, and `S3`.
	Name interface{}
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `sku` block as defined below. 
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
