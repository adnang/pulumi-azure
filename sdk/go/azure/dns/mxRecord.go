// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Enables you to manage DNS MX Records within Azure DNS.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/dns_mx_record.html.markdown.
type MxRecord struct {
	pulumi.CustomResourceState

	// The name of the DNS MX Record.
	Name pulumi.StringOutput `pulumi:"name"`

	// A list of values that make up the MX record. Each `record` block supports fields documented below.
	Records MxRecordRecordsArrayOutput `pulumi:"records"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntOutput `pulumi:"ttl"`

	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewMxRecord registers a new resource with the given unique name, arguments, and options.
func NewMxRecord(ctx *pulumi.Context,
	name string, args *MxRecordArgs, opts ...pulumi.ResourceOption) (*MxRecord, error) {
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
		if i := args.Records; i != nil { inputs["records"] = i.ToMxRecordRecordsArrayOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.Ttl; i != nil { inputs["ttl"] = i.ToIntOutput() }
		if i := args.ZoneName; i != nil { inputs["zoneName"] = i.ToStringOutput() }
	}
	var resource MxRecord
	err := ctx.RegisterResource("azure:dns/mxRecord:MxRecord", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMxRecord gets an existing MxRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMxRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MxRecordState, opts ...pulumi.ResourceOption) (*MxRecord, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Records; i != nil { inputs["records"] = i.ToMxRecordRecordsArrayOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.Ttl; i != nil { inputs["ttl"] = i.ToIntOutput() }
		if i := state.ZoneName; i != nil { inputs["zoneName"] = i.ToStringOutput() }
	}
	var resource MxRecord
	err := ctx.ReadResource("azure:dns/mxRecord:MxRecord", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MxRecord resources.
type MxRecordState struct {
	// The name of the DNS MX Record.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of values that make up the MX record. Each `record` block supports fields documented below.
	Records MxRecordRecordsArrayInput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntInput `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput `pulumi:"zoneName"`
}

// The set of arguments for constructing a MxRecord resource.
type MxRecordArgs struct {
	// The name of the DNS MX Record.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of values that make up the MX record. Each `record` block supports fields documented below.
	Records MxRecordRecordsArrayInput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntInput `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput `pulumi:"zoneName"`
}
type MxRecordRecords struct {
	Exchange string `pulumi:"exchange"`
	Preference string `pulumi:"preference"`
}
var mxRecordRecordsType = reflect.TypeOf((*MxRecordRecords)(nil)).Elem()

type MxRecordRecordsInput interface {
	pulumi.Input

	ToMxRecordRecordsOutput() MxRecordRecordsOutput
	ToMxRecordRecordsOutputWithContext(ctx context.Context) MxRecordRecordsOutput
}

type MxRecordRecordsArgs struct {
	Exchange pulumi.StringInput `pulumi:"exchange"`
	Preference pulumi.StringInput `pulumi:"preference"`
}

func (MxRecordRecordsArgs) ElementType() reflect.Type {
	return mxRecordRecordsType
}

func (a MxRecordRecordsArgs) ToMxRecordRecordsOutput() MxRecordRecordsOutput {
	return pulumi.ToOutput(a).(MxRecordRecordsOutput)
}

func (a MxRecordRecordsArgs) ToMxRecordRecordsOutputWithContext(ctx context.Context) MxRecordRecordsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(MxRecordRecordsOutput)
}

type MxRecordRecordsOutput struct { *pulumi.OutputState }

func (o MxRecordRecordsOutput) Exchange() pulumi.StringOutput {
	return o.Apply(func(v MxRecordRecords) string {
		return v.Exchange
	}).(pulumi.StringOutput)
}

func (o MxRecordRecordsOutput) Preference() pulumi.StringOutput {
	return o.Apply(func(v MxRecordRecords) string {
		return v.Preference
	}).(pulumi.StringOutput)
}

func (MxRecordRecordsOutput) ElementType() reflect.Type {
	return mxRecordRecordsType
}

func (o MxRecordRecordsOutput) ToMxRecordRecordsOutput() MxRecordRecordsOutput {
	return o
}

func (o MxRecordRecordsOutput) ToMxRecordRecordsOutputWithContext(ctx context.Context) MxRecordRecordsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(MxRecordRecordsOutput{}) }

var mxRecordRecordsArrayType = reflect.TypeOf((*[]MxRecordRecords)(nil)).Elem()

type MxRecordRecordsArrayInput interface {
	pulumi.Input

	ToMxRecordRecordsArrayOutput() MxRecordRecordsArrayOutput
	ToMxRecordRecordsArrayOutputWithContext(ctx context.Context) MxRecordRecordsArrayOutput
}

type MxRecordRecordsArrayArgs []MxRecordRecordsInput

func (MxRecordRecordsArrayArgs) ElementType() reflect.Type {
	return mxRecordRecordsArrayType
}

func (a MxRecordRecordsArrayArgs) ToMxRecordRecordsArrayOutput() MxRecordRecordsArrayOutput {
	return pulumi.ToOutput(a).(MxRecordRecordsArrayOutput)
}

func (a MxRecordRecordsArrayArgs) ToMxRecordRecordsArrayOutputWithContext(ctx context.Context) MxRecordRecordsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(MxRecordRecordsArrayOutput)
}

type MxRecordRecordsArrayOutput struct { *pulumi.OutputState }

func (o MxRecordRecordsArrayOutput) Index(i pulumi.IntInput) MxRecordRecordsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) MxRecordRecords {
		return vs[0].([]MxRecordRecords)[vs[1].(int)]
	}).(MxRecordRecordsOutput)
}

func (MxRecordRecordsArrayOutput) ElementType() reflect.Type {
	return mxRecordRecordsArrayType
}

func (o MxRecordRecordsArrayOutput) ToMxRecordRecordsArrayOutput() MxRecordRecordsArrayOutput {
	return o
}

func (o MxRecordRecordsArrayOutput) ToMxRecordRecordsArrayOutputWithContext(ctx context.Context) MxRecordRecordsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(MxRecordRecordsArrayOutput{}) }

