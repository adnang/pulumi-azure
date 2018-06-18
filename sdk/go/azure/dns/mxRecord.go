// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Enables you to manage DNS MX Records within Azure DNS.
type MxRecord struct {
	s *pulumi.ResourceState
}

// NewMxRecord registers a new resource with the given unique name, arguments, and options.
func NewMxRecord(ctx *pulumi.Context,
	name string, args *MxRecordArgs, opts ...pulumi.ResourceOpt) (*MxRecord, error) {
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
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["records"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
		inputs["ttl"] = nil
		inputs["zoneName"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["records"] = args.Records
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
		inputs["ttl"] = args.Ttl
		inputs["zoneName"] = args.ZoneName
	}
	s, err := ctx.RegisterResource("azure:dns/mxRecord:MxRecord", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MxRecord{s: s}, nil
}

// GetMxRecord gets an existing MxRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMxRecord(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MxRecordState, opts ...pulumi.ResourceOpt) (*MxRecord, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["records"] = state.Records
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
		inputs["ttl"] = state.Ttl
		inputs["zoneName"] = state.ZoneName
	}
	s, err := ctx.ReadResource("azure:dns/mxRecord:MxRecord", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MxRecord{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *MxRecord) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *MxRecord) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The name of the DNS MX Record.
func (r *MxRecord) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A list of values that make up the SRV record. Each `record` block supports fields documented below.
func (r *MxRecord) Records() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["records"])
}

// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
func (r *MxRecord) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *MxRecord) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The Time To Live (TTL) of the DNS record.
func (r *MxRecord) Ttl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ttl"])
}

// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
func (r *MxRecord) ZoneName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zoneName"])
}

// Input properties used for looking up and filtering MxRecord resources.
type MxRecordState struct {
	// The name of the DNS MX Record.
	Name interface{}
	// A list of values that make up the SRV record. Each `record` block supports fields documented below.
	Records interface{}
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The Time To Live (TTL) of the DNS record.
	Ttl interface{}
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName interface{}
}

// The set of arguments for constructing a MxRecord resource.
type MxRecordArgs struct {
	// The name of the DNS MX Record.
	Name interface{}
	// A list of values that make up the SRV record. Each `record` block supports fields documented below.
	Records interface{}
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The Time To Live (TTL) of the DNS record.
	Ttl interface{}
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName interface{}
}