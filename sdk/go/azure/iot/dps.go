// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an IotHub Device Provisioning Service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/iot_dps.html.markdown.
type Dps struct {
	pulumi.CustomResourceState

	// A `linkedHub` block as defined below.
	LinkedHubs DpsLinkedHubsArrayOutput `pulumi:"linkedHubs"`

	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// A `sku` block as defined below.
	Sku DpsSkuOutput `pulumi:"sku"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewDps registers a new resource with the given unique name, arguments, and options.
func NewDps(ctx *pulumi.Context,
	name string, args *DpsArgs, opts ...pulumi.ResourceOption) (*Dps, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.LinkedHubs; i != nil { inputs["linkedHubs"] = i.ToDpsLinkedHubsArrayOutput() }
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Sku; i != nil { inputs["sku"] = i.ToDpsSkuOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource Dps
	err := ctx.RegisterResource("azure:iot/dps:Dps", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDps gets an existing Dps resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDps(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DpsState, opts ...pulumi.ResourceOption) (*Dps, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.LinkedHubs; i != nil { inputs["linkedHubs"] = i.ToDpsLinkedHubsArrayOutput() }
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Sku; i != nil { inputs["sku"] = i.ToDpsSkuOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource Dps
	err := ctx.ReadResource("azure:iot/dps:Dps", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dps resources.
type DpsState struct {
	// A `linkedHub` block as defined below.
	LinkedHubs DpsLinkedHubsArrayInput `pulumi:"linkedHubs"`
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku DpsSkuInput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a Dps resource.
type DpsArgs struct {
	// A `linkedHub` block as defined below.
	LinkedHubs DpsLinkedHubsArrayInput `pulumi:"linkedHubs"`
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku DpsSkuInput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
type DpsLinkedHubs struct {
	AllocationWeight *int `pulumi:"allocationWeight"`
	ApplyAllocationPolicy *bool `pulumi:"applyAllocationPolicy"`
	ConnectionString string `pulumi:"connectionString"`
	Hostname *string `pulumi:"hostname"`
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location string `pulumi:"location"`
}
var dpsLinkedHubsType = reflect.TypeOf((*DpsLinkedHubs)(nil)).Elem()

type DpsLinkedHubsInput interface {
	pulumi.Input

	ToDpsLinkedHubsOutput() DpsLinkedHubsOutput
	ToDpsLinkedHubsOutputWithContext(ctx context.Context) DpsLinkedHubsOutput
}

type DpsLinkedHubsArgs struct {
	AllocationWeight pulumi.IntInput `pulumi:"allocationWeight"`
	ApplyAllocationPolicy pulumi.BoolInput `pulumi:"applyAllocationPolicy"`
	ConnectionString pulumi.StringInput `pulumi:"connectionString"`
	Hostname pulumi.StringInput `pulumi:"hostname"`
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
}

func (DpsLinkedHubsArgs) ElementType() reflect.Type {
	return dpsLinkedHubsType
}

func (a DpsLinkedHubsArgs) ToDpsLinkedHubsOutput() DpsLinkedHubsOutput {
	return pulumi.ToOutput(a).(DpsLinkedHubsOutput)
}

func (a DpsLinkedHubsArgs) ToDpsLinkedHubsOutputWithContext(ctx context.Context) DpsLinkedHubsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DpsLinkedHubsOutput)
}

type DpsLinkedHubsOutput struct { *pulumi.OutputState }

func (o DpsLinkedHubsOutput) AllocationWeight() pulumi.IntOutput {
	return o.Apply(func(v DpsLinkedHubs) int {
		if v.AllocationWeight == nil { return *new(int) } else { return *v.AllocationWeight }
	}).(pulumi.IntOutput)
}

func (o DpsLinkedHubsOutput) ApplyAllocationPolicy() pulumi.BoolOutput {
	return o.Apply(func(v DpsLinkedHubs) bool {
		if v.ApplyAllocationPolicy == nil { return *new(bool) } else { return *v.ApplyAllocationPolicy }
	}).(pulumi.BoolOutput)
}

func (o DpsLinkedHubsOutput) ConnectionString() pulumi.StringOutput {
	return o.Apply(func(v DpsLinkedHubs) string {
		return v.ConnectionString
	}).(pulumi.StringOutput)
}

func (o DpsLinkedHubsOutput) Hostname() pulumi.StringOutput {
	return o.Apply(func(v DpsLinkedHubs) string {
		if v.Hostname == nil { return *new(string) } else { return *v.Hostname }
	}).(pulumi.StringOutput)
}

// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
func (o DpsLinkedHubsOutput) Location() pulumi.StringOutput {
	return o.Apply(func(v DpsLinkedHubs) string {
		return v.Location
	}).(pulumi.StringOutput)
}

func (DpsLinkedHubsOutput) ElementType() reflect.Type {
	return dpsLinkedHubsType
}

func (o DpsLinkedHubsOutput) ToDpsLinkedHubsOutput() DpsLinkedHubsOutput {
	return o
}

func (o DpsLinkedHubsOutput) ToDpsLinkedHubsOutputWithContext(ctx context.Context) DpsLinkedHubsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DpsLinkedHubsOutput{}) }

var dpsLinkedHubsArrayType = reflect.TypeOf((*[]DpsLinkedHubs)(nil)).Elem()

type DpsLinkedHubsArrayInput interface {
	pulumi.Input

	ToDpsLinkedHubsArrayOutput() DpsLinkedHubsArrayOutput
	ToDpsLinkedHubsArrayOutputWithContext(ctx context.Context) DpsLinkedHubsArrayOutput
}

type DpsLinkedHubsArrayArgs []DpsLinkedHubsInput

func (DpsLinkedHubsArrayArgs) ElementType() reflect.Type {
	return dpsLinkedHubsArrayType
}

func (a DpsLinkedHubsArrayArgs) ToDpsLinkedHubsArrayOutput() DpsLinkedHubsArrayOutput {
	return pulumi.ToOutput(a).(DpsLinkedHubsArrayOutput)
}

func (a DpsLinkedHubsArrayArgs) ToDpsLinkedHubsArrayOutputWithContext(ctx context.Context) DpsLinkedHubsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DpsLinkedHubsArrayOutput)
}

type DpsLinkedHubsArrayOutput struct { *pulumi.OutputState }

func (o DpsLinkedHubsArrayOutput) Index(i pulumi.IntInput) DpsLinkedHubsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) DpsLinkedHubs {
		return vs[0].([]DpsLinkedHubs)[vs[1].(int)]
	}).(DpsLinkedHubsOutput)
}

func (DpsLinkedHubsArrayOutput) ElementType() reflect.Type {
	return dpsLinkedHubsArrayType
}

func (o DpsLinkedHubsArrayOutput) ToDpsLinkedHubsArrayOutput() DpsLinkedHubsArrayOutput {
	return o
}

func (o DpsLinkedHubsArrayOutput) ToDpsLinkedHubsArrayOutputWithContext(ctx context.Context) DpsLinkedHubsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DpsLinkedHubsArrayOutput{}) }

type DpsSku struct {
	Capacity int `pulumi:"capacity"`
	// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}
var dpsSkuType = reflect.TypeOf((*DpsSku)(nil)).Elem()

type DpsSkuInput interface {
	pulumi.Input

	ToDpsSkuOutput() DpsSkuOutput
	ToDpsSkuOutputWithContext(ctx context.Context) DpsSkuOutput
}

type DpsSkuArgs struct {
	Capacity pulumi.IntInput `pulumi:"capacity"`
	// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (DpsSkuArgs) ElementType() reflect.Type {
	return dpsSkuType
}

func (a DpsSkuArgs) ToDpsSkuOutput() DpsSkuOutput {
	return pulumi.ToOutput(a).(DpsSkuOutput)
}

func (a DpsSkuArgs) ToDpsSkuOutputWithContext(ctx context.Context) DpsSkuOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DpsSkuOutput)
}

type DpsSkuOutput struct { *pulumi.OutputState }

func (o DpsSkuOutput) Capacity() pulumi.IntOutput {
	return o.Apply(func(v DpsSku) int {
		return v.Capacity
	}).(pulumi.IntOutput)
}

// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
func (o DpsSkuOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v DpsSku) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o DpsSkuOutput) Tier() pulumi.StringOutput {
	return o.Apply(func(v DpsSku) string {
		return v.Tier
	}).(pulumi.StringOutput)
}

func (DpsSkuOutput) ElementType() reflect.Type {
	return dpsSkuType
}

func (o DpsSkuOutput) ToDpsSkuOutput() DpsSkuOutput {
	return o
}

func (o DpsSkuOutput) ToDpsSkuOutputWithContext(ctx context.Context) DpsSkuOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DpsSkuOutput{}) }

