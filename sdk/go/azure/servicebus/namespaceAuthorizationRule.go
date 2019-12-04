// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a ServiceBus Namespace authorization Rule within a ServiceBus.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/servicebus_namespace_authorization_rule.html.markdown.
type NamespaceAuthorizationRule struct {
	pulumi.CustomResourceState

	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolOutput `pulumi:"listen"`

	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolOutput `pulumi:"manage"`

	// Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`

	// The Primary Connection String for the ServiceBus Namespace authorization Rule.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`

	// The Primary Key for the ServiceBus Namespace authorization Rule.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`

	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The Secondary Connection String for the ServiceBus Namespace authorization Rule.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`

	// The Secondary Key for the ServiceBus Namespace authorization Rule.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`

	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolOutput `pulumi:"send"`
}

// NewNamespaceAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, args *NamespaceAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Listen; i != nil { inputs["listen"] = i.ToBoolOutput() }
		if i := args.Manage; i != nil { inputs["manage"] = i.ToBoolOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.NamespaceName; i != nil { inputs["namespaceName"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Send; i != nil { inputs["send"] = i.ToBoolOutput() }
	}
	var resource NamespaceAuthorizationRule
	err := ctx.RegisterResource("azure:servicebus/namespaceAuthorizationRule:NamespaceAuthorizationRule", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceAuthorizationRule gets an existing NamespaceAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceAuthorizationRuleState, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Listen; i != nil { inputs["listen"] = i.ToBoolOutput() }
		if i := state.Manage; i != nil { inputs["manage"] = i.ToBoolOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.NamespaceName; i != nil { inputs["namespaceName"] = i.ToStringOutput() }
		if i := state.PrimaryConnectionString; i != nil { inputs["primaryConnectionString"] = i.ToStringOutput() }
		if i := state.PrimaryKey; i != nil { inputs["primaryKey"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.SecondaryConnectionString; i != nil { inputs["secondaryConnectionString"] = i.ToStringOutput() }
		if i := state.SecondaryKey; i != nil { inputs["secondaryKey"] = i.ToStringOutput() }
		if i := state.Send; i != nil { inputs["send"] = i.ToBoolOutput() }
	}
	var resource NamespaceAuthorizationRule
	err := ctx.ReadResource("azure:servicebus/namespaceAuthorizationRule:NamespaceAuthorizationRule", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceAuthorizationRule resources.
type NamespaceAuthorizationRuleState struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolInput `pulumi:"listen"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolInput `pulumi:"manage"`
	// Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The Primary Connection String for the ServiceBus Namespace authorization Rule.
	PrimaryConnectionString pulumi.StringInput `pulumi:"primaryConnectionString"`
	// The Primary Key for the ServiceBus Namespace authorization Rule.
	PrimaryKey pulumi.StringInput `pulumi:"primaryKey"`
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the ServiceBus Namespace authorization Rule.
	SecondaryConnectionString pulumi.StringInput `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the ServiceBus Namespace authorization Rule.
	SecondaryKey pulumi.StringInput `pulumi:"secondaryKey"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolInput `pulumi:"send"`
}

// The set of arguments for constructing a NamespaceAuthorizationRule resource.
type NamespaceAuthorizationRuleArgs struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolInput `pulumi:"listen"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolInput `pulumi:"manage"`
	// Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolInput `pulumi:"send"`
}
