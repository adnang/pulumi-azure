// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Enables you to manage DNS CAA Records within Azure DNS.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/dns_caa_record.html.markdown.
type CaaRecord struct {
	pulumi.CustomResourceState

	// The name of the DNS CAA Record.
	Name pulumi.StringOutput `pulumi:"name"`

	// A list of values that make up the CAA record. Each `record` block supports fields documented below.
	Records CaaRecordRecordsArrayOutput `pulumi:"records"`

	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntOutput `pulumi:"ttl"`

	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewCaaRecord registers a new resource with the given unique name, arguments, and options.
func NewCaaRecord(ctx *pulumi.Context,
	name string, args *CaaRecordArgs, opts ...pulumi.ResourceOption) (*CaaRecord, error) {
	if args == nil || args.Records == nil {
		return nil, errors.New("missing required argument 'Records'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Ttl == nil {
		return nil, errors.New("missing required argument 'Ttl'")
	}
	if args == nil || args.ZoneName == nil {
		return nil, errors.New("missing required argument 'ZoneName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Records; i != nil { inputs["records"] = i.ToCaaRecordRecordsArrayOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.Ttl; i != nil { inputs["ttl"] = i.ToIntOutput() }
		if i := args.ZoneName; i != nil { inputs["zoneName"] = i.ToStringOutput() }
	}
	var resource CaaRecord
	err := ctx.RegisterResource("azure:dns/caaRecord:CaaRecord", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaaRecord gets an existing CaaRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaaRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaaRecordState, opts ...pulumi.ResourceOption) (*CaaRecord, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Records; i != nil { inputs["records"] = i.ToCaaRecordRecordsArrayOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.Ttl; i != nil { inputs["ttl"] = i.ToIntOutput() }
		if i := state.ZoneName; i != nil { inputs["zoneName"] = i.ToStringOutput() }
	}
	var resource CaaRecord
	err := ctx.ReadResource("azure:dns/caaRecord:CaaRecord", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaaRecord resources.
type CaaRecordState struct {
	// The name of the DNS CAA Record.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of values that make up the CAA record. Each `record` block supports fields documented below.
	Records CaaRecordRecordsArrayInput `pulumi:"records"`
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntInput `pulumi:"ttl"`
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput `pulumi:"zoneName"`
}

// The set of arguments for constructing a CaaRecord resource.
type CaaRecordArgs struct {
	// The name of the DNS CAA Record.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of values that make up the CAA record. Each `record` block supports fields documented below.
	Records CaaRecordRecordsArrayInput `pulumi:"records"`
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntInput `pulumi:"ttl"`
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput `pulumi:"zoneName"`
}
type CaaRecordRecords struct {
	Flags int `pulumi:"flags"`
	Tag string `pulumi:"tag"`
	Value string `pulumi:"value"`
}
var caaRecordRecordsType = reflect.TypeOf((*CaaRecordRecords)(nil)).Elem()

type CaaRecordRecordsInput interface {
	pulumi.Input

	ToCaaRecordRecordsOutput() CaaRecordRecordsOutput
	ToCaaRecordRecordsOutputWithContext(ctx context.Context) CaaRecordRecordsOutput
}

type CaaRecordRecordsArgs struct {
	Flags pulumi.IntInput `pulumi:"flags"`
	Tag pulumi.StringInput `pulumi:"tag"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (CaaRecordRecordsArgs) ElementType() reflect.Type {
	return caaRecordRecordsType
}

func (a CaaRecordRecordsArgs) ToCaaRecordRecordsOutput() CaaRecordRecordsOutput {
	return pulumi.ToOutput(a).(CaaRecordRecordsOutput)
}

func (a CaaRecordRecordsArgs) ToCaaRecordRecordsOutputWithContext(ctx context.Context) CaaRecordRecordsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(CaaRecordRecordsOutput)
}

type CaaRecordRecordsOutput struct { *pulumi.OutputState }

func (o CaaRecordRecordsOutput) Flags() pulumi.IntOutput {
	return o.Apply(func(v CaaRecordRecords) int {
		return v.Flags
	}).(pulumi.IntOutput)
}

func (o CaaRecordRecordsOutput) Tag() pulumi.StringOutput {
	return o.Apply(func(v CaaRecordRecords) string {
		return v.Tag
	}).(pulumi.StringOutput)
}

func (o CaaRecordRecordsOutput) Value() pulumi.StringOutput {
	return o.Apply(func(v CaaRecordRecords) string {
		return v.Value
	}).(pulumi.StringOutput)
}

func (CaaRecordRecordsOutput) ElementType() reflect.Type {
	return caaRecordRecordsType
}

func (o CaaRecordRecordsOutput) ToCaaRecordRecordsOutput() CaaRecordRecordsOutput {
	return o
}

func (o CaaRecordRecordsOutput) ToCaaRecordRecordsOutputWithContext(ctx context.Context) CaaRecordRecordsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(CaaRecordRecordsOutput{}) }

var caaRecordRecordsArrayType = reflect.TypeOf((*[]CaaRecordRecords)(nil)).Elem()

type CaaRecordRecordsArrayInput interface {
	pulumi.Input

	ToCaaRecordRecordsArrayOutput() CaaRecordRecordsArrayOutput
	ToCaaRecordRecordsArrayOutputWithContext(ctx context.Context) CaaRecordRecordsArrayOutput
}

type CaaRecordRecordsArrayArgs []CaaRecordRecordsInput

func (CaaRecordRecordsArrayArgs) ElementType() reflect.Type {
	return caaRecordRecordsArrayType
}

func (a CaaRecordRecordsArrayArgs) ToCaaRecordRecordsArrayOutput() CaaRecordRecordsArrayOutput {
	return pulumi.ToOutput(a).(CaaRecordRecordsArrayOutput)
}

func (a CaaRecordRecordsArrayArgs) ToCaaRecordRecordsArrayOutputWithContext(ctx context.Context) CaaRecordRecordsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(CaaRecordRecordsArrayOutput)
}

type CaaRecordRecordsArrayOutput struct { *pulumi.OutputState }

func (o CaaRecordRecordsArrayOutput) Index(i pulumi.IntInput) CaaRecordRecordsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) CaaRecordRecords {
		return vs[0].([]CaaRecordRecords)[vs[1].(int)]
	}).(CaaRecordRecordsOutput)
}

func (CaaRecordRecordsArrayOutput) ElementType() reflect.Type {
	return caaRecordRecordsArrayType
}

func (o CaaRecordRecordsArrayOutput) ToCaaRecordRecordsArrayOutput() CaaRecordRecordsArrayOutput {
	return o
}

func (o CaaRecordRecordsArrayOutput) ToCaaRecordRecordsArrayOutputWithContext(ctx context.Context) CaaRecordRecordsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(CaaRecordRecordsArrayOutput{}) }

