// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Traffic Manager Endpoint.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/traffic_manager_endpoint.html.markdown.
type TrafficManagerEndpoint struct {
	pulumi.CustomResourceState

	// One or more `customHeader` blocks as defined below
	CustomHeaders TrafficManagerEndpointCustomHeadersArrayOutput `pulumi:"customHeaders"`

	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation pulumi.StringOutput `pulumi:"endpointLocation"`

	EndpointMonitorStatus pulumi.StringOutput `pulumi:"endpointMonitorStatus"`

	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus pulumi.StringOutput `pulumi:"endpointStatus"`

	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings pulumi.StringArrayOutput `pulumi:"geoMappings"`

	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and defaults to `1`.
	MinChildEndpoints pulumi.IntOutput `pulumi:"minChildEndpoints"`

	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority pulumi.IntOutput `pulumi:"priority"`

	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`

	// The name of the resource group in which to
	// create the Traffic Manager endpoint.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// One or more `subnet` blocks as defined below
	Subnets TrafficManagerEndpointSubnetsArrayOutput `pulumi:"subnets"`

	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target pulumi.StringOutput `pulumi:"target"`

	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`

	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type pulumi.StringOutput `pulumi:"type"`

	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewTrafficManagerEndpoint registers a new resource with the given unique name, arguments, and options.
func NewTrafficManagerEndpoint(ctx *pulumi.Context,
	name string, args *TrafficManagerEndpointArgs, opts ...pulumi.ResourceOption) (*TrafficManagerEndpoint, error) {
	if args == nil || args.ProfileName == nil {
		return nil, errors.New("missing required argument 'ProfileName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.CustomHeaders; i != nil { inputs["customHeaders"] = i.ToTrafficManagerEndpointCustomHeadersArrayOutput() }
		if i := args.EndpointLocation; i != nil { inputs["endpointLocation"] = i.ToStringOutput() }
		if i := args.EndpointStatus; i != nil { inputs["endpointStatus"] = i.ToStringOutput() }
		if i := args.GeoMappings; i != nil { inputs["geoMappings"] = i.ToStringArrayOutput() }
		if i := args.MinChildEndpoints; i != nil { inputs["minChildEndpoints"] = i.ToIntOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Priority; i != nil { inputs["priority"] = i.ToIntOutput() }
		if i := args.ProfileName; i != nil { inputs["profileName"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Subnets; i != nil { inputs["subnets"] = i.ToTrafficManagerEndpointSubnetsArrayOutput() }
		if i := args.Target; i != nil { inputs["target"] = i.ToStringOutput() }
		if i := args.TargetResourceId; i != nil { inputs["targetResourceId"] = i.ToStringOutput() }
		if i := args.Type; i != nil { inputs["type"] = i.ToStringOutput() }
		if i := args.Weight; i != nil { inputs["weight"] = i.ToIntOutput() }
	}
	var resource TrafficManagerEndpoint
	err := ctx.RegisterResource("azure:network/trafficManagerEndpoint:TrafficManagerEndpoint", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficManagerEndpoint gets an existing TrafficManagerEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficManagerEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficManagerEndpointState, opts ...pulumi.ResourceOption) (*TrafficManagerEndpoint, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.CustomHeaders; i != nil { inputs["customHeaders"] = i.ToTrafficManagerEndpointCustomHeadersArrayOutput() }
		if i := state.EndpointLocation; i != nil { inputs["endpointLocation"] = i.ToStringOutput() }
		if i := state.EndpointMonitorStatus; i != nil { inputs["endpointMonitorStatus"] = i.ToStringOutput() }
		if i := state.EndpointStatus; i != nil { inputs["endpointStatus"] = i.ToStringOutput() }
		if i := state.GeoMappings; i != nil { inputs["geoMappings"] = i.ToStringArrayOutput() }
		if i := state.MinChildEndpoints; i != nil { inputs["minChildEndpoints"] = i.ToIntOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Priority; i != nil { inputs["priority"] = i.ToIntOutput() }
		if i := state.ProfileName; i != nil { inputs["profileName"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Subnets; i != nil { inputs["subnets"] = i.ToTrafficManagerEndpointSubnetsArrayOutput() }
		if i := state.Target; i != nil { inputs["target"] = i.ToStringOutput() }
		if i := state.TargetResourceId; i != nil { inputs["targetResourceId"] = i.ToStringOutput() }
		if i := state.Type; i != nil { inputs["type"] = i.ToStringOutput() }
		if i := state.Weight; i != nil { inputs["weight"] = i.ToIntOutput() }
	}
	var resource TrafficManagerEndpoint
	err := ctx.ReadResource("azure:network/trafficManagerEndpoint:TrafficManagerEndpoint", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficManagerEndpoint resources.
type TrafficManagerEndpointState struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders TrafficManagerEndpointCustomHeadersArrayInput `pulumi:"customHeaders"`
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation pulumi.StringInput `pulumi:"endpointLocation"`
	EndpointMonitorStatus pulumi.StringInput `pulumi:"endpointMonitorStatus"`
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus pulumi.StringInput `pulumi:"endpointStatus"`
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings pulumi.StringArrayInput `pulumi:"geoMappings"`
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and defaults to `1`.
	MinChildEndpoints pulumi.IntInput `pulumi:"minChildEndpoints"`
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority pulumi.IntInput `pulumi:"priority"`
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// The name of the resource group in which to
	// create the Traffic Manager endpoint.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// One or more `subnet` blocks as defined below
	Subnets TrafficManagerEndpointSubnetsArrayInput `pulumi:"subnets"`
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target pulumi.StringInput `pulumi:"target"`
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId pulumi.StringInput `pulumi:"targetResourceId"`
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type pulumi.StringInput `pulumi:"type"`
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight pulumi.IntInput `pulumi:"weight"`
}

// The set of arguments for constructing a TrafficManagerEndpoint resource.
type TrafficManagerEndpointArgs struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders TrafficManagerEndpointCustomHeadersArrayInput `pulumi:"customHeaders"`
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation pulumi.StringInput `pulumi:"endpointLocation"`
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus pulumi.StringInput `pulumi:"endpointStatus"`
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings pulumi.StringArrayInput `pulumi:"geoMappings"`
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and defaults to `1`.
	MinChildEndpoints pulumi.IntInput `pulumi:"minChildEndpoints"`
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority pulumi.IntInput `pulumi:"priority"`
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// The name of the resource group in which to
	// create the Traffic Manager endpoint.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// One or more `subnet` blocks as defined below
	Subnets TrafficManagerEndpointSubnetsArrayInput `pulumi:"subnets"`
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target pulumi.StringInput `pulumi:"target"`
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId pulumi.StringInput `pulumi:"targetResourceId"`
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type pulumi.StringInput `pulumi:"type"`
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight pulumi.IntInput `pulumi:"weight"`
}
type TrafficManagerEndpointCustomHeaders struct {
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name string `pulumi:"name"`
	Value string `pulumi:"value"`
}
var trafficManagerEndpointCustomHeadersType = reflect.TypeOf((*TrafficManagerEndpointCustomHeaders)(nil)).Elem()

type TrafficManagerEndpointCustomHeadersInput interface {
	pulumi.Input

	ToTrafficManagerEndpointCustomHeadersOutput() TrafficManagerEndpointCustomHeadersOutput
	ToTrafficManagerEndpointCustomHeadersOutputWithContext(ctx context.Context) TrafficManagerEndpointCustomHeadersOutput
}

type TrafficManagerEndpointCustomHeadersArgs struct {
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TrafficManagerEndpointCustomHeadersArgs) ElementType() reflect.Type {
	return trafficManagerEndpointCustomHeadersType
}

func (a TrafficManagerEndpointCustomHeadersArgs) ToTrafficManagerEndpointCustomHeadersOutput() TrafficManagerEndpointCustomHeadersOutput {
	return pulumi.ToOutput(a).(TrafficManagerEndpointCustomHeadersOutput)
}

func (a TrafficManagerEndpointCustomHeadersArgs) ToTrafficManagerEndpointCustomHeadersOutputWithContext(ctx context.Context) TrafficManagerEndpointCustomHeadersOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TrafficManagerEndpointCustomHeadersOutput)
}

type TrafficManagerEndpointCustomHeadersOutput struct { *pulumi.OutputState }

// The name of the Traffic Manager endpoint. Changing this forces a
// new resource to be created.
func (o TrafficManagerEndpointCustomHeadersOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v TrafficManagerEndpointCustomHeaders) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o TrafficManagerEndpointCustomHeadersOutput) Value() pulumi.StringOutput {
	return o.Apply(func(v TrafficManagerEndpointCustomHeaders) string {
		return v.Value
	}).(pulumi.StringOutput)
}

func (TrafficManagerEndpointCustomHeadersOutput) ElementType() reflect.Type {
	return trafficManagerEndpointCustomHeadersType
}

func (o TrafficManagerEndpointCustomHeadersOutput) ToTrafficManagerEndpointCustomHeadersOutput() TrafficManagerEndpointCustomHeadersOutput {
	return o
}

func (o TrafficManagerEndpointCustomHeadersOutput) ToTrafficManagerEndpointCustomHeadersOutputWithContext(ctx context.Context) TrafficManagerEndpointCustomHeadersOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TrafficManagerEndpointCustomHeadersOutput{}) }

var trafficManagerEndpointCustomHeadersArrayType = reflect.TypeOf((*[]TrafficManagerEndpointCustomHeaders)(nil)).Elem()

type TrafficManagerEndpointCustomHeadersArrayInput interface {
	pulumi.Input

	ToTrafficManagerEndpointCustomHeadersArrayOutput() TrafficManagerEndpointCustomHeadersArrayOutput
	ToTrafficManagerEndpointCustomHeadersArrayOutputWithContext(ctx context.Context) TrafficManagerEndpointCustomHeadersArrayOutput
}

type TrafficManagerEndpointCustomHeadersArrayArgs []TrafficManagerEndpointCustomHeadersInput

func (TrafficManagerEndpointCustomHeadersArrayArgs) ElementType() reflect.Type {
	return trafficManagerEndpointCustomHeadersArrayType
}

func (a TrafficManagerEndpointCustomHeadersArrayArgs) ToTrafficManagerEndpointCustomHeadersArrayOutput() TrafficManagerEndpointCustomHeadersArrayOutput {
	return pulumi.ToOutput(a).(TrafficManagerEndpointCustomHeadersArrayOutput)
}

func (a TrafficManagerEndpointCustomHeadersArrayArgs) ToTrafficManagerEndpointCustomHeadersArrayOutputWithContext(ctx context.Context) TrafficManagerEndpointCustomHeadersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TrafficManagerEndpointCustomHeadersArrayOutput)
}

type TrafficManagerEndpointCustomHeadersArrayOutput struct { *pulumi.OutputState }

func (o TrafficManagerEndpointCustomHeadersArrayOutput) Index(i pulumi.IntInput) TrafficManagerEndpointCustomHeadersOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) TrafficManagerEndpointCustomHeaders {
		return vs[0].([]TrafficManagerEndpointCustomHeaders)[vs[1].(int)]
	}).(TrafficManagerEndpointCustomHeadersOutput)
}

func (TrafficManagerEndpointCustomHeadersArrayOutput) ElementType() reflect.Type {
	return trafficManagerEndpointCustomHeadersArrayType
}

func (o TrafficManagerEndpointCustomHeadersArrayOutput) ToTrafficManagerEndpointCustomHeadersArrayOutput() TrafficManagerEndpointCustomHeadersArrayOutput {
	return o
}

func (o TrafficManagerEndpointCustomHeadersArrayOutput) ToTrafficManagerEndpointCustomHeadersArrayOutputWithContext(ctx context.Context) TrafficManagerEndpointCustomHeadersArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TrafficManagerEndpointCustomHeadersArrayOutput{}) }

type TrafficManagerEndpointSubnets struct {
	First string `pulumi:"first"`
	Last *string `pulumi:"last"`
	Scope *int `pulumi:"scope"`
}
var trafficManagerEndpointSubnetsType = reflect.TypeOf((*TrafficManagerEndpointSubnets)(nil)).Elem()

type TrafficManagerEndpointSubnetsInput interface {
	pulumi.Input

	ToTrafficManagerEndpointSubnetsOutput() TrafficManagerEndpointSubnetsOutput
	ToTrafficManagerEndpointSubnetsOutputWithContext(ctx context.Context) TrafficManagerEndpointSubnetsOutput
}

type TrafficManagerEndpointSubnetsArgs struct {
	First pulumi.StringInput `pulumi:"first"`
	Last pulumi.StringInput `pulumi:"last"`
	Scope pulumi.IntInput `pulumi:"scope"`
}

func (TrafficManagerEndpointSubnetsArgs) ElementType() reflect.Type {
	return trafficManagerEndpointSubnetsType
}

func (a TrafficManagerEndpointSubnetsArgs) ToTrafficManagerEndpointSubnetsOutput() TrafficManagerEndpointSubnetsOutput {
	return pulumi.ToOutput(a).(TrafficManagerEndpointSubnetsOutput)
}

func (a TrafficManagerEndpointSubnetsArgs) ToTrafficManagerEndpointSubnetsOutputWithContext(ctx context.Context) TrafficManagerEndpointSubnetsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TrafficManagerEndpointSubnetsOutput)
}

type TrafficManagerEndpointSubnetsOutput struct { *pulumi.OutputState }

func (o TrafficManagerEndpointSubnetsOutput) First() pulumi.StringOutput {
	return o.Apply(func(v TrafficManagerEndpointSubnets) string {
		return v.First
	}).(pulumi.StringOutput)
}

func (o TrafficManagerEndpointSubnetsOutput) Last() pulumi.StringOutput {
	return o.Apply(func(v TrafficManagerEndpointSubnets) string {
		if v.Last == nil { return *new(string) } else { return *v.Last }
	}).(pulumi.StringOutput)
}

func (o TrafficManagerEndpointSubnetsOutput) Scope() pulumi.IntOutput {
	return o.Apply(func(v TrafficManagerEndpointSubnets) int {
		if v.Scope == nil { return *new(int) } else { return *v.Scope }
	}).(pulumi.IntOutput)
}

func (TrafficManagerEndpointSubnetsOutput) ElementType() reflect.Type {
	return trafficManagerEndpointSubnetsType
}

func (o TrafficManagerEndpointSubnetsOutput) ToTrafficManagerEndpointSubnetsOutput() TrafficManagerEndpointSubnetsOutput {
	return o
}

func (o TrafficManagerEndpointSubnetsOutput) ToTrafficManagerEndpointSubnetsOutputWithContext(ctx context.Context) TrafficManagerEndpointSubnetsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TrafficManagerEndpointSubnetsOutput{}) }

var trafficManagerEndpointSubnetsArrayType = reflect.TypeOf((*[]TrafficManagerEndpointSubnets)(nil)).Elem()

type TrafficManagerEndpointSubnetsArrayInput interface {
	pulumi.Input

	ToTrafficManagerEndpointSubnetsArrayOutput() TrafficManagerEndpointSubnetsArrayOutput
	ToTrafficManagerEndpointSubnetsArrayOutputWithContext(ctx context.Context) TrafficManagerEndpointSubnetsArrayOutput
}

type TrafficManagerEndpointSubnetsArrayArgs []TrafficManagerEndpointSubnetsInput

func (TrafficManagerEndpointSubnetsArrayArgs) ElementType() reflect.Type {
	return trafficManagerEndpointSubnetsArrayType
}

func (a TrafficManagerEndpointSubnetsArrayArgs) ToTrafficManagerEndpointSubnetsArrayOutput() TrafficManagerEndpointSubnetsArrayOutput {
	return pulumi.ToOutput(a).(TrafficManagerEndpointSubnetsArrayOutput)
}

func (a TrafficManagerEndpointSubnetsArrayArgs) ToTrafficManagerEndpointSubnetsArrayOutputWithContext(ctx context.Context) TrafficManagerEndpointSubnetsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TrafficManagerEndpointSubnetsArrayOutput)
}

type TrafficManagerEndpointSubnetsArrayOutput struct { *pulumi.OutputState }

func (o TrafficManagerEndpointSubnetsArrayOutput) Index(i pulumi.IntInput) TrafficManagerEndpointSubnetsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) TrafficManagerEndpointSubnets {
		return vs[0].([]TrafficManagerEndpointSubnets)[vs[1].(int)]
	}).(TrafficManagerEndpointSubnetsOutput)
}

func (TrafficManagerEndpointSubnetsArrayOutput) ElementType() reflect.Type {
	return trafficManagerEndpointSubnetsArrayType
}

func (o TrafficManagerEndpointSubnetsArrayOutput) ToTrafficManagerEndpointSubnetsArrayOutput() TrafficManagerEndpointSubnetsArrayOutput {
	return o
}

func (o TrafficManagerEndpointSubnetsArrayOutput) ToTrafficManagerEndpointSubnetsArrayOutputWithContext(ctx context.Context) TrafficManagerEndpointSubnetsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TrafficManagerEndpointSubnetsArrayOutput{}) }

