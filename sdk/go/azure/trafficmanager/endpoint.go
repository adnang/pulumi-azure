// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package trafficmanager

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Traffic Manager Endpoint.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/traffic_manager_endpoint_legacy.html.markdown.
type Endpoint struct {
	pulumi.CustomResourceState

	// One or more `customHeader` blocks as defined below
	CustomHeaders EndpointCustomHeadersArrayOutput `pulumi:"customHeaders"`

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
	Subnets EndpointSubnetsArrayOutput `pulumi:"subnets"`

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

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
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
		if i := args.CustomHeaders; i != nil { inputs["customHeaders"] = i.ToEndpointCustomHeadersArrayOutput() }
		if i := args.EndpointLocation; i != nil { inputs["endpointLocation"] = i.ToStringOutput() }
		if i := args.EndpointStatus; i != nil { inputs["endpointStatus"] = i.ToStringOutput() }
		if i := args.GeoMappings; i != nil { inputs["geoMappings"] = i.ToStringArrayOutput() }
		if i := args.MinChildEndpoints; i != nil { inputs["minChildEndpoints"] = i.ToIntOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Priority; i != nil { inputs["priority"] = i.ToIntOutput() }
		if i := args.ProfileName; i != nil { inputs["profileName"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Subnets; i != nil { inputs["subnets"] = i.ToEndpointSubnetsArrayOutput() }
		if i := args.Target; i != nil { inputs["target"] = i.ToStringOutput() }
		if i := args.TargetResourceId; i != nil { inputs["targetResourceId"] = i.ToStringOutput() }
		if i := args.Type; i != nil { inputs["type"] = i.ToStringOutput() }
		if i := args.Weight; i != nil { inputs["weight"] = i.ToIntOutput() }
	}
	var resource Endpoint
	err := ctx.RegisterResource("azure:trafficmanager/endpoint:Endpoint", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.CustomHeaders; i != nil { inputs["customHeaders"] = i.ToEndpointCustomHeadersArrayOutput() }
		if i := state.EndpointLocation; i != nil { inputs["endpointLocation"] = i.ToStringOutput() }
		if i := state.EndpointMonitorStatus; i != nil { inputs["endpointMonitorStatus"] = i.ToStringOutput() }
		if i := state.EndpointStatus; i != nil { inputs["endpointStatus"] = i.ToStringOutput() }
		if i := state.GeoMappings; i != nil { inputs["geoMappings"] = i.ToStringArrayOutput() }
		if i := state.MinChildEndpoints; i != nil { inputs["minChildEndpoints"] = i.ToIntOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Priority; i != nil { inputs["priority"] = i.ToIntOutput() }
		if i := state.ProfileName; i != nil { inputs["profileName"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Subnets; i != nil { inputs["subnets"] = i.ToEndpointSubnetsArrayOutput() }
		if i := state.Target; i != nil { inputs["target"] = i.ToStringOutput() }
		if i := state.TargetResourceId; i != nil { inputs["targetResourceId"] = i.ToStringOutput() }
		if i := state.Type; i != nil { inputs["type"] = i.ToStringOutput() }
		if i := state.Weight; i != nil { inputs["weight"] = i.ToIntOutput() }
	}
	var resource Endpoint
	err := ctx.ReadResource("azure:trafficmanager/endpoint:Endpoint", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type EndpointState struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders EndpointCustomHeadersArrayInput `pulumi:"customHeaders"`
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
	Subnets EndpointSubnetsArrayInput `pulumi:"subnets"`
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

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders EndpointCustomHeadersArrayInput `pulumi:"customHeaders"`
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
	Subnets EndpointSubnetsArrayInput `pulumi:"subnets"`
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
type EndpointCustomHeaders struct {
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name string `pulumi:"name"`
	Value string `pulumi:"value"`
}
var endpointCustomHeadersType = reflect.TypeOf((*EndpointCustomHeaders)(nil)).Elem()

type EndpointCustomHeadersInput interface {
	pulumi.Input

	ToEndpointCustomHeadersOutput() EndpointCustomHeadersOutput
	ToEndpointCustomHeadersOutputWithContext(ctx context.Context) EndpointCustomHeadersOutput
}

type EndpointCustomHeadersArgs struct {
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (EndpointCustomHeadersArgs) ElementType() reflect.Type {
	return endpointCustomHeadersType
}

func (a EndpointCustomHeadersArgs) ToEndpointCustomHeadersOutput() EndpointCustomHeadersOutput {
	return pulumi.ToOutput(a).(EndpointCustomHeadersOutput)
}

func (a EndpointCustomHeadersArgs) ToEndpointCustomHeadersOutputWithContext(ctx context.Context) EndpointCustomHeadersOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EndpointCustomHeadersOutput)
}

type EndpointCustomHeadersOutput struct { *pulumi.OutputState }

// The name of the Traffic Manager endpoint. Changing this forces a
// new resource to be created.
func (o EndpointCustomHeadersOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v EndpointCustomHeaders) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o EndpointCustomHeadersOutput) Value() pulumi.StringOutput {
	return o.Apply(func(v EndpointCustomHeaders) string {
		return v.Value
	}).(pulumi.StringOutput)
}

func (EndpointCustomHeadersOutput) ElementType() reflect.Type {
	return endpointCustomHeadersType
}

func (o EndpointCustomHeadersOutput) ToEndpointCustomHeadersOutput() EndpointCustomHeadersOutput {
	return o
}

func (o EndpointCustomHeadersOutput) ToEndpointCustomHeadersOutputWithContext(ctx context.Context) EndpointCustomHeadersOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EndpointCustomHeadersOutput{}) }

var endpointCustomHeadersArrayType = reflect.TypeOf((*[]EndpointCustomHeaders)(nil)).Elem()

type EndpointCustomHeadersArrayInput interface {
	pulumi.Input

	ToEndpointCustomHeadersArrayOutput() EndpointCustomHeadersArrayOutput
	ToEndpointCustomHeadersArrayOutputWithContext(ctx context.Context) EndpointCustomHeadersArrayOutput
}

type EndpointCustomHeadersArrayArgs []EndpointCustomHeadersInput

func (EndpointCustomHeadersArrayArgs) ElementType() reflect.Type {
	return endpointCustomHeadersArrayType
}

func (a EndpointCustomHeadersArrayArgs) ToEndpointCustomHeadersArrayOutput() EndpointCustomHeadersArrayOutput {
	return pulumi.ToOutput(a).(EndpointCustomHeadersArrayOutput)
}

func (a EndpointCustomHeadersArrayArgs) ToEndpointCustomHeadersArrayOutputWithContext(ctx context.Context) EndpointCustomHeadersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EndpointCustomHeadersArrayOutput)
}

type EndpointCustomHeadersArrayOutput struct { *pulumi.OutputState }

func (o EndpointCustomHeadersArrayOutput) Index(i pulumi.IntInput) EndpointCustomHeadersOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) EndpointCustomHeaders {
		return vs[0].([]EndpointCustomHeaders)[vs[1].(int)]
	}).(EndpointCustomHeadersOutput)
}

func (EndpointCustomHeadersArrayOutput) ElementType() reflect.Type {
	return endpointCustomHeadersArrayType
}

func (o EndpointCustomHeadersArrayOutput) ToEndpointCustomHeadersArrayOutput() EndpointCustomHeadersArrayOutput {
	return o
}

func (o EndpointCustomHeadersArrayOutput) ToEndpointCustomHeadersArrayOutputWithContext(ctx context.Context) EndpointCustomHeadersArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EndpointCustomHeadersArrayOutput{}) }

type EndpointSubnets struct {
	First string `pulumi:"first"`
	Last *string `pulumi:"last"`
	Scope *int `pulumi:"scope"`
}
var endpointSubnetsType = reflect.TypeOf((*EndpointSubnets)(nil)).Elem()

type EndpointSubnetsInput interface {
	pulumi.Input

	ToEndpointSubnetsOutput() EndpointSubnetsOutput
	ToEndpointSubnetsOutputWithContext(ctx context.Context) EndpointSubnetsOutput
}

type EndpointSubnetsArgs struct {
	First pulumi.StringInput `pulumi:"first"`
	Last pulumi.StringInput `pulumi:"last"`
	Scope pulumi.IntInput `pulumi:"scope"`
}

func (EndpointSubnetsArgs) ElementType() reflect.Type {
	return endpointSubnetsType
}

func (a EndpointSubnetsArgs) ToEndpointSubnetsOutput() EndpointSubnetsOutput {
	return pulumi.ToOutput(a).(EndpointSubnetsOutput)
}

func (a EndpointSubnetsArgs) ToEndpointSubnetsOutputWithContext(ctx context.Context) EndpointSubnetsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EndpointSubnetsOutput)
}

type EndpointSubnetsOutput struct { *pulumi.OutputState }

func (o EndpointSubnetsOutput) First() pulumi.StringOutput {
	return o.Apply(func(v EndpointSubnets) string {
		return v.First
	}).(pulumi.StringOutput)
}

func (o EndpointSubnetsOutput) Last() pulumi.StringOutput {
	return o.Apply(func(v EndpointSubnets) string {
		if v.Last == nil { return *new(string) } else { return *v.Last }
	}).(pulumi.StringOutput)
}

func (o EndpointSubnetsOutput) Scope() pulumi.IntOutput {
	return o.Apply(func(v EndpointSubnets) int {
		if v.Scope == nil { return *new(int) } else { return *v.Scope }
	}).(pulumi.IntOutput)
}

func (EndpointSubnetsOutput) ElementType() reflect.Type {
	return endpointSubnetsType
}

func (o EndpointSubnetsOutput) ToEndpointSubnetsOutput() EndpointSubnetsOutput {
	return o
}

func (o EndpointSubnetsOutput) ToEndpointSubnetsOutputWithContext(ctx context.Context) EndpointSubnetsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EndpointSubnetsOutput{}) }

var endpointSubnetsArrayType = reflect.TypeOf((*[]EndpointSubnets)(nil)).Elem()

type EndpointSubnetsArrayInput interface {
	pulumi.Input

	ToEndpointSubnetsArrayOutput() EndpointSubnetsArrayOutput
	ToEndpointSubnetsArrayOutputWithContext(ctx context.Context) EndpointSubnetsArrayOutput
}

type EndpointSubnetsArrayArgs []EndpointSubnetsInput

func (EndpointSubnetsArrayArgs) ElementType() reflect.Type {
	return endpointSubnetsArrayType
}

func (a EndpointSubnetsArrayArgs) ToEndpointSubnetsArrayOutput() EndpointSubnetsArrayOutput {
	return pulumi.ToOutput(a).(EndpointSubnetsArrayOutput)
}

func (a EndpointSubnetsArrayArgs) ToEndpointSubnetsArrayOutputWithContext(ctx context.Context) EndpointSubnetsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EndpointSubnetsArrayOutput)
}

type EndpointSubnetsArrayOutput struct { *pulumi.OutputState }

func (o EndpointSubnetsArrayOutput) Index(i pulumi.IntInput) EndpointSubnetsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) EndpointSubnets {
		return vs[0].([]EndpointSubnets)[vs[1].(int)]
	}).(EndpointSubnetsOutput)
}

func (EndpointSubnetsArrayOutput) ElementType() reflect.Type {
	return endpointSubnetsArrayType
}

func (o EndpointSubnetsArrayOutput) ToEndpointSubnetsArrayOutput() EndpointSubnetsArrayOutput {
	return o
}

func (o EndpointSubnetsArrayOutput) ToEndpointSubnetsArrayOutputWithContext(ctx context.Context) EndpointSubnetsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EndpointSubnetsArrayOutput{}) }

