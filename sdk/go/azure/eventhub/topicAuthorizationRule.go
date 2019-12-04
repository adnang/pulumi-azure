// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a ServiceBus Topic authorization Rule within a ServiceBus Topic.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/servicebus_topic_authorization_rule_legacy.html.markdown.
type TopicAuthorizationRule struct {
	pulumi.CustomResourceState

	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolOutput `pulumi:"listen"`

	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolOutput `pulumi:"manage"`

	// Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`

	// The Primary Connection String for the ServiceBus Topic authorization Rule.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`

	// The Primary Key for the ServiceBus Topic authorization Rule.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`

	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The Secondary Connection String for the ServiceBus Topic authorization Rule.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`

	// The Secondary Key for the ServiceBus Topic authorization Rule.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`

	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolOutput `pulumi:"send"`

	// Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewTopicAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewTopicAuthorizationRule(ctx *pulumi.Context,
	name string, args *TopicAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*TopicAuthorizationRule, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TopicName == nil {
		return nil, errors.New("missing required argument 'TopicName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Listen; i != nil { inputs["listen"] = i.ToBoolOutput() }
		if i := args.Manage; i != nil { inputs["manage"] = i.ToBoolOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.NamespaceName; i != nil { inputs["namespaceName"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Send; i != nil { inputs["send"] = i.ToBoolOutput() }
		if i := args.TopicName; i != nil { inputs["topicName"] = i.ToStringOutput() }
	}
	var resource TopicAuthorizationRule
	err := ctx.RegisterResource("azure:eventhub/topicAuthorizationRule:TopicAuthorizationRule", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicAuthorizationRule gets an existing TopicAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicAuthorizationRuleState, opts ...pulumi.ResourceOption) (*TopicAuthorizationRule, error) {
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
		if i := state.TopicName; i != nil { inputs["topicName"] = i.ToStringOutput() }
	}
	var resource TopicAuthorizationRule
	err := ctx.ReadResource("azure:eventhub/topicAuthorizationRule:TopicAuthorizationRule", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicAuthorizationRule resources.
type TopicAuthorizationRuleState struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolInput `pulumi:"listen"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolInput `pulumi:"manage"`
	// Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The Primary Connection String for the ServiceBus Topic authorization Rule.
	PrimaryConnectionString pulumi.StringInput `pulumi:"primaryConnectionString"`
	// The Primary Key for the ServiceBus Topic authorization Rule.
	PrimaryKey pulumi.StringInput `pulumi:"primaryKey"`
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the ServiceBus Topic authorization Rule.
	SecondaryConnectionString pulumi.StringInput `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the ServiceBus Topic authorization Rule.
	SecondaryKey pulumi.StringInput `pulumi:"secondaryKey"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolInput `pulumi:"send"`
	// Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

// The set of arguments for constructing a TopicAuthorizationRule resource.
type TopicAuthorizationRuleArgs struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolInput `pulumi:"listen"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolInput `pulumi:"manage"`
	// Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolInput `pulumi:"send"`
	// Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}
