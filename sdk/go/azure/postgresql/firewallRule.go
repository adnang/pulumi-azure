// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Firewall Rule for a PostgreSQL Server
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/postgresql_firewall_rule.html.markdown.
type FirewallRule struct {
	pulumi.CustomResourceState

	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`

	// Specifies the name of the PostgreSQL Firewall Rule. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`

	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil || args.EndIpAddress == nil {
		return nil, errors.New("missing required argument 'EndIpAddress'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.StartIpAddress == nil {
		return nil, errors.New("missing required argument 'StartIpAddress'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.EndIpAddress; i != nil { inputs["endIpAddress"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.ServerName; i != nil { inputs["serverName"] = i.ToStringOutput() }
		if i := args.StartIpAddress; i != nil { inputs["startIpAddress"] = i.ToStringOutput() }
	}
	var resource FirewallRule
	err := ctx.RegisterResource("azure:postgresql/firewallRule:FirewallRule", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.EndIpAddress; i != nil { inputs["endIpAddress"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.ServerName; i != nil { inputs["serverName"] = i.ToStringOutput() }
		if i := state.StartIpAddress; i != nil { inputs["startIpAddress"] = i.ToStringOutput() }
	}
	var resource FirewallRule
	err := ctx.ReadResource("azure:postgresql/firewallRule:FirewallRule", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type FirewallRuleState struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress pulumi.StringInput `pulumi:"endIpAddress"`
	// Specifies the name of the PostgreSQL Firewall Rule. Changing this forces a
	// new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress pulumi.StringInput `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress pulumi.StringInput `pulumi:"endIpAddress"`
	// Specifies the name of the PostgreSQL Firewall Rule. Changing this forces a
	// new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress pulumi.StringInput `pulumi:"startIpAddress"`
}
