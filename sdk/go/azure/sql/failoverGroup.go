// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Create a failover group of databases on a collection of Azure SQL servers.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/sql_failover_group.html.markdown.
type FailoverGroup struct {
	pulumi.CustomResourceState

	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayOutput `pulumi:"databases"`

	// the location of the failover group.
	Location pulumi.StringOutput `pulumi:"location"`

	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServersArrayOutput `pulumi:"partnerServers"`

	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyOutput `pulumi:"readWriteEndpointFailoverPolicy"`

	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyOutput `pulumi:"readonlyEndpointFailoverPolicy"`

	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// local replication role of the failover group instance.
	Role pulumi.StringOutput `pulumi:"role"`

	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewFailoverGroup registers a new resource with the given unique name, arguments, and options.
func NewFailoverGroup(ctx *pulumi.Context,
	name string, args *FailoverGroupArgs, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	if args == nil || args.PartnerServers == nil {
		return nil, errors.New("missing required argument 'PartnerServers'")
	}
	if args == nil || args.ReadWriteEndpointFailoverPolicy == nil {
		return nil, errors.New("missing required argument 'ReadWriteEndpointFailoverPolicy'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Databases; i != nil { inputs["databases"] = i.ToStringArrayOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.PartnerServers; i != nil { inputs["partnerServers"] = i.ToFailoverGroupPartnerServersArrayOutput() }
		if i := args.ReadWriteEndpointFailoverPolicy; i != nil { inputs["readWriteEndpointFailoverPolicy"] = i.ToFailoverGroupReadWriteEndpointFailoverPolicyOutput() }
		if i := args.ReadonlyEndpointFailoverPolicy; i != nil { inputs["readonlyEndpointFailoverPolicy"] = i.ToFailoverGroupReadonlyEndpointFailoverPolicyOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.ServerName; i != nil { inputs["serverName"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource FailoverGroup
	err := ctx.RegisterResource("azure:sql/failoverGroup:FailoverGroup", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFailoverGroup gets an existing FailoverGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFailoverGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FailoverGroupState, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Databases; i != nil { inputs["databases"] = i.ToStringArrayOutput() }
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.PartnerServers; i != nil { inputs["partnerServers"] = i.ToFailoverGroupPartnerServersArrayOutput() }
		if i := state.ReadWriteEndpointFailoverPolicy; i != nil { inputs["readWriteEndpointFailoverPolicy"] = i.ToFailoverGroupReadWriteEndpointFailoverPolicyOutput() }
		if i := state.ReadonlyEndpointFailoverPolicy; i != nil { inputs["readonlyEndpointFailoverPolicy"] = i.ToFailoverGroupReadonlyEndpointFailoverPolicyOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Role; i != nil { inputs["role"] = i.ToStringOutput() }
		if i := state.ServerName; i != nil { inputs["serverName"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource FailoverGroup
	err := ctx.ReadResource("azure:sql/failoverGroup:FailoverGroup", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FailoverGroup resources.
type FailoverGroupState struct {
	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayInput `pulumi:"databases"`
	// the location of the failover group.
	Location pulumi.StringInput `pulumi:"location"`
	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServersArrayInput `pulumi:"partnerServers"`
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyInput `pulumi:"readWriteEndpointFailoverPolicy"`
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyInput `pulumi:"readonlyEndpointFailoverPolicy"`
	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// local replication role of the failover group instance.
	Role pulumi.StringInput `pulumi:"role"`
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a FailoverGroup resource.
type FailoverGroupArgs struct {
	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayInput `pulumi:"databases"`
	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServersArrayInput `pulumi:"partnerServers"`
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyInput `pulumi:"readWriteEndpointFailoverPolicy"`
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyInput `pulumi:"readonlyEndpointFailoverPolicy"`
	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
type FailoverGroupPartnerServers struct {
	// the SQL server ID
	Id string `pulumi:"id"`
	// the location of the failover group.
	Location *string `pulumi:"location"`
	// local replication role of the failover group instance.
	Role *string `pulumi:"role"`
}
var failoverGroupPartnerServersType = reflect.TypeOf((*FailoverGroupPartnerServers)(nil)).Elem()

type FailoverGroupPartnerServersInput interface {
	pulumi.Input

	ToFailoverGroupPartnerServersOutput() FailoverGroupPartnerServersOutput
	ToFailoverGroupPartnerServersOutputWithContext(ctx context.Context) FailoverGroupPartnerServersOutput
}

type FailoverGroupPartnerServersArgs struct {
	// the SQL server ID
	Id pulumi.StringInput `pulumi:"id"`
	// the location of the failover group.
	Location pulumi.StringInput `pulumi:"location"`
	// local replication role of the failover group instance.
	Role pulumi.StringInput `pulumi:"role"`
}

func (FailoverGroupPartnerServersArgs) ElementType() reflect.Type {
	return failoverGroupPartnerServersType
}

func (a FailoverGroupPartnerServersArgs) ToFailoverGroupPartnerServersOutput() FailoverGroupPartnerServersOutput {
	return pulumi.ToOutput(a).(FailoverGroupPartnerServersOutput)
}

func (a FailoverGroupPartnerServersArgs) ToFailoverGroupPartnerServersOutputWithContext(ctx context.Context) FailoverGroupPartnerServersOutput {
	return pulumi.ToOutputWithContext(ctx, a).(FailoverGroupPartnerServersOutput)
}

type FailoverGroupPartnerServersOutput struct { *pulumi.OutputState }

// the SQL server ID
func (o FailoverGroupPartnerServersOutput) Id() pulumi.StringOutput {
	return o.Apply(func(v FailoverGroupPartnerServers) string {
		return v.Id
	}).(pulumi.StringOutput)
}

// the location of the failover group.
func (o FailoverGroupPartnerServersOutput) Location() pulumi.StringOutput {
	return o.Apply(func(v FailoverGroupPartnerServers) string {
		if v.Location == nil { return *new(string) } else { return *v.Location }
	}).(pulumi.StringOutput)
}

// local replication role of the failover group instance.
func (o FailoverGroupPartnerServersOutput) Role() pulumi.StringOutput {
	return o.Apply(func(v FailoverGroupPartnerServers) string {
		if v.Role == nil { return *new(string) } else { return *v.Role }
	}).(pulumi.StringOutput)
}

func (FailoverGroupPartnerServersOutput) ElementType() reflect.Type {
	return failoverGroupPartnerServersType
}

func (o FailoverGroupPartnerServersOutput) ToFailoverGroupPartnerServersOutput() FailoverGroupPartnerServersOutput {
	return o
}

func (o FailoverGroupPartnerServersOutput) ToFailoverGroupPartnerServersOutputWithContext(ctx context.Context) FailoverGroupPartnerServersOutput {
	return o
}

func init() { pulumi.RegisterOutputType(FailoverGroupPartnerServersOutput{}) }

var failoverGroupPartnerServersArrayType = reflect.TypeOf((*[]FailoverGroupPartnerServers)(nil)).Elem()

type FailoverGroupPartnerServersArrayInput interface {
	pulumi.Input

	ToFailoverGroupPartnerServersArrayOutput() FailoverGroupPartnerServersArrayOutput
	ToFailoverGroupPartnerServersArrayOutputWithContext(ctx context.Context) FailoverGroupPartnerServersArrayOutput
}

type FailoverGroupPartnerServersArrayArgs []FailoverGroupPartnerServersInput

func (FailoverGroupPartnerServersArrayArgs) ElementType() reflect.Type {
	return failoverGroupPartnerServersArrayType
}

func (a FailoverGroupPartnerServersArrayArgs) ToFailoverGroupPartnerServersArrayOutput() FailoverGroupPartnerServersArrayOutput {
	return pulumi.ToOutput(a).(FailoverGroupPartnerServersArrayOutput)
}

func (a FailoverGroupPartnerServersArrayArgs) ToFailoverGroupPartnerServersArrayOutputWithContext(ctx context.Context) FailoverGroupPartnerServersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(FailoverGroupPartnerServersArrayOutput)
}

type FailoverGroupPartnerServersArrayOutput struct { *pulumi.OutputState }

func (o FailoverGroupPartnerServersArrayOutput) Index(i pulumi.IntInput) FailoverGroupPartnerServersOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) FailoverGroupPartnerServers {
		return vs[0].([]FailoverGroupPartnerServers)[vs[1].(int)]
	}).(FailoverGroupPartnerServersOutput)
}

func (FailoverGroupPartnerServersArrayOutput) ElementType() reflect.Type {
	return failoverGroupPartnerServersArrayType
}

func (o FailoverGroupPartnerServersArrayOutput) ToFailoverGroupPartnerServersArrayOutput() FailoverGroupPartnerServersArrayOutput {
	return o
}

func (o FailoverGroupPartnerServersArrayOutput) ToFailoverGroupPartnerServersArrayOutputWithContext(ctx context.Context) FailoverGroupPartnerServersArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(FailoverGroupPartnerServersArrayOutput{}) }

type FailoverGroupReadWriteEndpointFailoverPolicy struct {
	// Applies only if `mode` is `Automatic`. The grace period in minutes before failover with data loss is attempted
	GraceMinutes *int `pulumi:"graceMinutes"`
	// Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
	Mode string `pulumi:"mode"`
}
var failoverGroupReadWriteEndpointFailoverPolicyType = reflect.TypeOf((*FailoverGroupReadWriteEndpointFailoverPolicy)(nil)).Elem()

type FailoverGroupReadWriteEndpointFailoverPolicyInput interface {
	pulumi.Input

	ToFailoverGroupReadWriteEndpointFailoverPolicyOutput() FailoverGroupReadWriteEndpointFailoverPolicyOutput
	ToFailoverGroupReadWriteEndpointFailoverPolicyOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointFailoverPolicyOutput
}

type FailoverGroupReadWriteEndpointFailoverPolicyArgs struct {
	// Applies only if `mode` is `Automatic`. The grace period in minutes before failover with data loss is attempted
	GraceMinutes pulumi.IntInput `pulumi:"graceMinutes"`
	// Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
	Mode pulumi.StringInput `pulumi:"mode"`
}

func (FailoverGroupReadWriteEndpointFailoverPolicyArgs) ElementType() reflect.Type {
	return failoverGroupReadWriteEndpointFailoverPolicyType
}

func (a FailoverGroupReadWriteEndpointFailoverPolicyArgs) ToFailoverGroupReadWriteEndpointFailoverPolicyOutput() FailoverGroupReadWriteEndpointFailoverPolicyOutput {
	return pulumi.ToOutput(a).(FailoverGroupReadWriteEndpointFailoverPolicyOutput)
}

func (a FailoverGroupReadWriteEndpointFailoverPolicyArgs) ToFailoverGroupReadWriteEndpointFailoverPolicyOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointFailoverPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(FailoverGroupReadWriteEndpointFailoverPolicyOutput)
}

type FailoverGroupReadWriteEndpointFailoverPolicyOutput struct { *pulumi.OutputState }

// Applies only if `mode` is `Automatic`. The grace period in minutes before failover with data loss is attempted
func (o FailoverGroupReadWriteEndpointFailoverPolicyOutput) GraceMinutes() pulumi.IntOutput {
	return o.Apply(func(v FailoverGroupReadWriteEndpointFailoverPolicy) int {
		if v.GraceMinutes == nil { return *new(int) } else { return *v.GraceMinutes }
	}).(pulumi.IntOutput)
}

// Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
func (o FailoverGroupReadWriteEndpointFailoverPolicyOutput) Mode() pulumi.StringOutput {
	return o.Apply(func(v FailoverGroupReadWriteEndpointFailoverPolicy) string {
		return v.Mode
	}).(pulumi.StringOutput)
}

func (FailoverGroupReadWriteEndpointFailoverPolicyOutput) ElementType() reflect.Type {
	return failoverGroupReadWriteEndpointFailoverPolicyType
}

func (o FailoverGroupReadWriteEndpointFailoverPolicyOutput) ToFailoverGroupReadWriteEndpointFailoverPolicyOutput() FailoverGroupReadWriteEndpointFailoverPolicyOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointFailoverPolicyOutput) ToFailoverGroupReadWriteEndpointFailoverPolicyOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointFailoverPolicyOutput {
	return o
}

func init() { pulumi.RegisterOutputType(FailoverGroupReadWriteEndpointFailoverPolicyOutput{}) }

type FailoverGroupReadonlyEndpointFailoverPolicy struct {
	// Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
	Mode string `pulumi:"mode"`
}
var failoverGroupReadonlyEndpointFailoverPolicyType = reflect.TypeOf((*FailoverGroupReadonlyEndpointFailoverPolicy)(nil)).Elem()

type FailoverGroupReadonlyEndpointFailoverPolicyInput interface {
	pulumi.Input

	ToFailoverGroupReadonlyEndpointFailoverPolicyOutput() FailoverGroupReadonlyEndpointFailoverPolicyOutput
	ToFailoverGroupReadonlyEndpointFailoverPolicyOutputWithContext(ctx context.Context) FailoverGroupReadonlyEndpointFailoverPolicyOutput
}

type FailoverGroupReadonlyEndpointFailoverPolicyArgs struct {
	// Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
	Mode pulumi.StringInput `pulumi:"mode"`
}

func (FailoverGroupReadonlyEndpointFailoverPolicyArgs) ElementType() reflect.Type {
	return failoverGroupReadonlyEndpointFailoverPolicyType
}

func (a FailoverGroupReadonlyEndpointFailoverPolicyArgs) ToFailoverGroupReadonlyEndpointFailoverPolicyOutput() FailoverGroupReadonlyEndpointFailoverPolicyOutput {
	return pulumi.ToOutput(a).(FailoverGroupReadonlyEndpointFailoverPolicyOutput)
}

func (a FailoverGroupReadonlyEndpointFailoverPolicyArgs) ToFailoverGroupReadonlyEndpointFailoverPolicyOutputWithContext(ctx context.Context) FailoverGroupReadonlyEndpointFailoverPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(FailoverGroupReadonlyEndpointFailoverPolicyOutput)
}

type FailoverGroupReadonlyEndpointFailoverPolicyOutput struct { *pulumi.OutputState }

// Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
func (o FailoverGroupReadonlyEndpointFailoverPolicyOutput) Mode() pulumi.StringOutput {
	return o.Apply(func(v FailoverGroupReadonlyEndpointFailoverPolicy) string {
		return v.Mode
	}).(pulumi.StringOutput)
}

func (FailoverGroupReadonlyEndpointFailoverPolicyOutput) ElementType() reflect.Type {
	return failoverGroupReadonlyEndpointFailoverPolicyType
}

func (o FailoverGroupReadonlyEndpointFailoverPolicyOutput) ToFailoverGroupReadonlyEndpointFailoverPolicyOutput() FailoverGroupReadonlyEndpointFailoverPolicyOutput {
	return o
}

func (o FailoverGroupReadonlyEndpointFailoverPolicyOutput) ToFailoverGroupReadonlyEndpointFailoverPolicyOutputWithContext(ctx context.Context) FailoverGroupReadonlyEndpointFailoverPolicyOutput {
	return o
}

func init() { pulumi.RegisterOutputType(FailoverGroupReadonlyEndpointFailoverPolicyOutput{}) }

