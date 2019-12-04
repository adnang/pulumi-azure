// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Application Rule Collection within an Azure Firewall.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/firewall_application_rule_collection.html.markdown.
type FirewallApplicationRuleCollection struct {
	pulumi.CustomResourceState

	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringOutput `pulumi:"action"`

	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringOutput `pulumi:"azureFirewallName"`

	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntOutput `pulumi:"priority"`

	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// One or more `rule` blocks as defined below.
	Rules FirewallApplicationRuleCollectionRulesArrayOutput `pulumi:"rules"`
}

// NewFirewallApplicationRuleCollection registers a new resource with the given unique name, arguments, and options.
func NewFirewallApplicationRuleCollection(ctx *pulumi.Context,
	name string, args *FirewallApplicationRuleCollectionArgs, opts ...pulumi.ResourceOption) (*FirewallApplicationRuleCollection, error) {
	if args == nil || args.Action == nil {
		return nil, errors.New("missing required argument 'Action'")
	}
	if args == nil || args.AzureFirewallName == nil {
		return nil, errors.New("missing required argument 'AzureFirewallName'")
	}
	if args == nil || args.Priority == nil {
		return nil, errors.New("missing required argument 'Priority'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Rules == nil {
		return nil, errors.New("missing required argument 'Rules'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Action; i != nil { inputs["action"] = i.ToStringOutput() }
		if i := args.AzureFirewallName; i != nil { inputs["azureFirewallName"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Priority; i != nil { inputs["priority"] = i.ToIntOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Rules; i != nil { inputs["rules"] = i.ToFirewallApplicationRuleCollectionRulesArrayOutput() }
	}
	var resource FirewallApplicationRuleCollection
	err := ctx.RegisterResource("azure:network/firewallApplicationRuleCollection:FirewallApplicationRuleCollection", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallApplicationRuleCollection gets an existing FirewallApplicationRuleCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallApplicationRuleCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallApplicationRuleCollectionState, opts ...pulumi.ResourceOption) (*FirewallApplicationRuleCollection, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Action; i != nil { inputs["action"] = i.ToStringOutput() }
		if i := state.AzureFirewallName; i != nil { inputs["azureFirewallName"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Priority; i != nil { inputs["priority"] = i.ToIntOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Rules; i != nil { inputs["rules"] = i.ToFirewallApplicationRuleCollectionRulesArrayOutput() }
	}
	var resource FirewallApplicationRuleCollection
	err := ctx.ReadResource("azure:network/firewallApplicationRuleCollection:FirewallApplicationRuleCollection", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallApplicationRuleCollection resources.
type FirewallApplicationRuleCollectionState struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringInput `pulumi:"action"`
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringInput `pulumi:"azureFirewallName"`
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntInput `pulumi:"priority"`
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// One or more `rule` blocks as defined below.
	Rules FirewallApplicationRuleCollectionRulesArrayInput `pulumi:"rules"`
}

// The set of arguments for constructing a FirewallApplicationRuleCollection resource.
type FirewallApplicationRuleCollectionArgs struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringInput `pulumi:"action"`
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringInput `pulumi:"azureFirewallName"`
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntInput `pulumi:"priority"`
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// One or more `rule` blocks as defined below.
	Rules FirewallApplicationRuleCollectionRulesArrayInput `pulumi:"rules"`
}
type FirewallApplicationRuleCollectionRules struct {
	Description *string `pulumi:"description"`
	FqdnTags *[]string `pulumi:"fqdnTags"`
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name string `pulumi:"name"`
	Protocols *[]FirewallApplicationRuleCollectionRulesProtocols `pulumi:"protocols"`
	SourceAddresses []string `pulumi:"sourceAddresses"`
	TargetFqdns *[]string `pulumi:"targetFqdns"`
}
var firewallApplicationRuleCollectionRulesType = reflect.TypeOf((*FirewallApplicationRuleCollectionRules)(nil)).Elem()

type FirewallApplicationRuleCollectionRulesInput interface {
	pulumi.Input

	ToFirewallApplicationRuleCollectionRulesOutput() FirewallApplicationRuleCollectionRulesOutput
	ToFirewallApplicationRuleCollectionRulesOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesOutput
}

type FirewallApplicationRuleCollectionRulesArgs struct {
	Description pulumi.StringInput `pulumi:"description"`
	FqdnTags pulumi.StringArrayInput `pulumi:"fqdnTags"`
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	Protocols FirewallApplicationRuleCollectionRulesProtocolsArrayInput `pulumi:"protocols"`
	SourceAddresses pulumi.StringArrayInput `pulumi:"sourceAddresses"`
	TargetFqdns pulumi.StringArrayInput `pulumi:"targetFqdns"`
}

func (FirewallApplicationRuleCollectionRulesArgs) ElementType() reflect.Type {
	return firewallApplicationRuleCollectionRulesType
}

func (a FirewallApplicationRuleCollectionRulesArgs) ToFirewallApplicationRuleCollectionRulesOutput() FirewallApplicationRuleCollectionRulesOutput {
	return pulumi.ToOutput(a).(FirewallApplicationRuleCollectionRulesOutput)
}

func (a FirewallApplicationRuleCollectionRulesArgs) ToFirewallApplicationRuleCollectionRulesOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(FirewallApplicationRuleCollectionRulesOutput)
}

type FirewallApplicationRuleCollectionRulesOutput struct { *pulumi.OutputState }

func (o FirewallApplicationRuleCollectionRulesOutput) Description() pulumi.StringOutput {
	return o.Apply(func(v FirewallApplicationRuleCollectionRules) string {
		if v.Description == nil { return *new(string) } else { return *v.Description }
	}).(pulumi.StringOutput)
}

func (o FirewallApplicationRuleCollectionRulesOutput) FqdnTags() pulumi.StringArrayOutput {
	return o.Apply(func(v FirewallApplicationRuleCollectionRules) []string {
		if v.FqdnTags == nil { return *new([]string) } else { return *v.FqdnTags }
	}).(pulumi.StringArrayOutput)
}

// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
func (o FirewallApplicationRuleCollectionRulesOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v FirewallApplicationRuleCollectionRules) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o FirewallApplicationRuleCollectionRulesOutput) Protocols() FirewallApplicationRuleCollectionRulesProtocolsArrayOutput {
	return o.Apply(func(v FirewallApplicationRuleCollectionRules) []FirewallApplicationRuleCollectionRulesProtocols {
		if v.Protocols == nil { return *new([]FirewallApplicationRuleCollectionRulesProtocols) } else { return *v.Protocols }
	}).(FirewallApplicationRuleCollectionRulesProtocolsArrayOutput)
}

func (o FirewallApplicationRuleCollectionRulesOutput) SourceAddresses() pulumi.StringArrayOutput {
	return o.Apply(func(v FirewallApplicationRuleCollectionRules) []string {
		return v.SourceAddresses
	}).(pulumi.StringArrayOutput)
}

func (o FirewallApplicationRuleCollectionRulesOutput) TargetFqdns() pulumi.StringArrayOutput {
	return o.Apply(func(v FirewallApplicationRuleCollectionRules) []string {
		if v.TargetFqdns == nil { return *new([]string) } else { return *v.TargetFqdns }
	}).(pulumi.StringArrayOutput)
}

func (FirewallApplicationRuleCollectionRulesOutput) ElementType() reflect.Type {
	return firewallApplicationRuleCollectionRulesType
}

func (o FirewallApplicationRuleCollectionRulesOutput) ToFirewallApplicationRuleCollectionRulesOutput() FirewallApplicationRuleCollectionRulesOutput {
	return o
}

func (o FirewallApplicationRuleCollectionRulesOutput) ToFirewallApplicationRuleCollectionRulesOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(FirewallApplicationRuleCollectionRulesOutput{}) }

var firewallApplicationRuleCollectionRulesArrayType = reflect.TypeOf((*[]FirewallApplicationRuleCollectionRules)(nil)).Elem()

type FirewallApplicationRuleCollectionRulesArrayInput interface {
	pulumi.Input

	ToFirewallApplicationRuleCollectionRulesArrayOutput() FirewallApplicationRuleCollectionRulesArrayOutput
	ToFirewallApplicationRuleCollectionRulesArrayOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesArrayOutput
}

type FirewallApplicationRuleCollectionRulesArrayArgs []FirewallApplicationRuleCollectionRulesInput

func (FirewallApplicationRuleCollectionRulesArrayArgs) ElementType() reflect.Type {
	return firewallApplicationRuleCollectionRulesArrayType
}

func (a FirewallApplicationRuleCollectionRulesArrayArgs) ToFirewallApplicationRuleCollectionRulesArrayOutput() FirewallApplicationRuleCollectionRulesArrayOutput {
	return pulumi.ToOutput(a).(FirewallApplicationRuleCollectionRulesArrayOutput)
}

func (a FirewallApplicationRuleCollectionRulesArrayArgs) ToFirewallApplicationRuleCollectionRulesArrayOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(FirewallApplicationRuleCollectionRulesArrayOutput)
}

type FirewallApplicationRuleCollectionRulesArrayOutput struct { *pulumi.OutputState }

func (o FirewallApplicationRuleCollectionRulesArrayOutput) Index(i pulumi.IntInput) FirewallApplicationRuleCollectionRulesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) FirewallApplicationRuleCollectionRules {
		return vs[0].([]FirewallApplicationRuleCollectionRules)[vs[1].(int)]
	}).(FirewallApplicationRuleCollectionRulesOutput)
}

func (FirewallApplicationRuleCollectionRulesArrayOutput) ElementType() reflect.Type {
	return firewallApplicationRuleCollectionRulesArrayType
}

func (o FirewallApplicationRuleCollectionRulesArrayOutput) ToFirewallApplicationRuleCollectionRulesArrayOutput() FirewallApplicationRuleCollectionRulesArrayOutput {
	return o
}

func (o FirewallApplicationRuleCollectionRulesArrayOutput) ToFirewallApplicationRuleCollectionRulesArrayOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(FirewallApplicationRuleCollectionRulesArrayOutput{}) }

type FirewallApplicationRuleCollectionRulesProtocols struct {
	Port *int `pulumi:"port"`
	Type string `pulumi:"type"`
}
var firewallApplicationRuleCollectionRulesProtocolsType = reflect.TypeOf((*FirewallApplicationRuleCollectionRulesProtocols)(nil)).Elem()

type FirewallApplicationRuleCollectionRulesProtocolsInput interface {
	pulumi.Input

	ToFirewallApplicationRuleCollectionRulesProtocolsOutput() FirewallApplicationRuleCollectionRulesProtocolsOutput
	ToFirewallApplicationRuleCollectionRulesProtocolsOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesProtocolsOutput
}

type FirewallApplicationRuleCollectionRulesProtocolsArgs struct {
	Port pulumi.IntInput `pulumi:"port"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (FirewallApplicationRuleCollectionRulesProtocolsArgs) ElementType() reflect.Type {
	return firewallApplicationRuleCollectionRulesProtocolsType
}

func (a FirewallApplicationRuleCollectionRulesProtocolsArgs) ToFirewallApplicationRuleCollectionRulesProtocolsOutput() FirewallApplicationRuleCollectionRulesProtocolsOutput {
	return pulumi.ToOutput(a).(FirewallApplicationRuleCollectionRulesProtocolsOutput)
}

func (a FirewallApplicationRuleCollectionRulesProtocolsArgs) ToFirewallApplicationRuleCollectionRulesProtocolsOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesProtocolsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(FirewallApplicationRuleCollectionRulesProtocolsOutput)
}

type FirewallApplicationRuleCollectionRulesProtocolsOutput struct { *pulumi.OutputState }

func (o FirewallApplicationRuleCollectionRulesProtocolsOutput) Port() pulumi.IntOutput {
	return o.Apply(func(v FirewallApplicationRuleCollectionRulesProtocols) int {
		if v.Port == nil { return *new(int) } else { return *v.Port }
	}).(pulumi.IntOutput)
}

func (o FirewallApplicationRuleCollectionRulesProtocolsOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v FirewallApplicationRuleCollectionRulesProtocols) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (FirewallApplicationRuleCollectionRulesProtocolsOutput) ElementType() reflect.Type {
	return firewallApplicationRuleCollectionRulesProtocolsType
}

func (o FirewallApplicationRuleCollectionRulesProtocolsOutput) ToFirewallApplicationRuleCollectionRulesProtocolsOutput() FirewallApplicationRuleCollectionRulesProtocolsOutput {
	return o
}

func (o FirewallApplicationRuleCollectionRulesProtocolsOutput) ToFirewallApplicationRuleCollectionRulesProtocolsOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesProtocolsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(FirewallApplicationRuleCollectionRulesProtocolsOutput{}) }

var firewallApplicationRuleCollectionRulesProtocolsArrayType = reflect.TypeOf((*[]FirewallApplicationRuleCollectionRulesProtocols)(nil)).Elem()

type FirewallApplicationRuleCollectionRulesProtocolsArrayInput interface {
	pulumi.Input

	ToFirewallApplicationRuleCollectionRulesProtocolsArrayOutput() FirewallApplicationRuleCollectionRulesProtocolsArrayOutput
	ToFirewallApplicationRuleCollectionRulesProtocolsArrayOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesProtocolsArrayOutput
}

type FirewallApplicationRuleCollectionRulesProtocolsArrayArgs []FirewallApplicationRuleCollectionRulesProtocolsInput

func (FirewallApplicationRuleCollectionRulesProtocolsArrayArgs) ElementType() reflect.Type {
	return firewallApplicationRuleCollectionRulesProtocolsArrayType
}

func (a FirewallApplicationRuleCollectionRulesProtocolsArrayArgs) ToFirewallApplicationRuleCollectionRulesProtocolsArrayOutput() FirewallApplicationRuleCollectionRulesProtocolsArrayOutput {
	return pulumi.ToOutput(a).(FirewallApplicationRuleCollectionRulesProtocolsArrayOutput)
}

func (a FirewallApplicationRuleCollectionRulesProtocolsArrayArgs) ToFirewallApplicationRuleCollectionRulesProtocolsArrayOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesProtocolsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(FirewallApplicationRuleCollectionRulesProtocolsArrayOutput)
}

type FirewallApplicationRuleCollectionRulesProtocolsArrayOutput struct { *pulumi.OutputState }

func (o FirewallApplicationRuleCollectionRulesProtocolsArrayOutput) Index(i pulumi.IntInput) FirewallApplicationRuleCollectionRulesProtocolsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) FirewallApplicationRuleCollectionRulesProtocols {
		return vs[0].([]FirewallApplicationRuleCollectionRulesProtocols)[vs[1].(int)]
	}).(FirewallApplicationRuleCollectionRulesProtocolsOutput)
}

func (FirewallApplicationRuleCollectionRulesProtocolsArrayOutput) ElementType() reflect.Type {
	return firewallApplicationRuleCollectionRulesProtocolsArrayType
}

func (o FirewallApplicationRuleCollectionRulesProtocolsArrayOutput) ToFirewallApplicationRuleCollectionRulesProtocolsArrayOutput() FirewallApplicationRuleCollectionRulesProtocolsArrayOutput {
	return o
}

func (o FirewallApplicationRuleCollectionRulesProtocolsArrayOutput) ToFirewallApplicationRuleCollectionRulesProtocolsArrayOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionRulesProtocolsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(FirewallApplicationRuleCollectionRulesProtocolsArrayOutput{}) }

