// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appinsights

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Application Insights Analytics Item component.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/application_insights_analytics_item.html.markdown.
type AnalyticsItem struct {
	pulumi.CustomResourceState

	// The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringOutput `pulumi:"applicationInsightsId"`

	// The content for the Analytics Item, for example the query text if `type` is `query`.
	Content pulumi.StringOutput `pulumi:"content"`

	// The alias to use for the function. Required when `type` is `function`.
	FunctionAlias pulumi.StringOutput `pulumi:"functionAlias"`

	// Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
	Scope pulumi.StringOutput `pulumi:"scope"`

	// A string containing the time the Analytics Item was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`

	// A string containing the time the Analytics Item was last modified.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`

	// The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
	Type pulumi.StringOutput `pulumi:"type"`

	// A string indicating the version of the query format
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewAnalyticsItem registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsItem(ctx *pulumi.Context,
	name string, args *AnalyticsItemArgs, opts ...pulumi.ResourceOption) (*AnalyticsItem, error) {
	if args == nil || args.ApplicationInsightsId == nil {
		return nil, errors.New("missing required argument 'ApplicationInsightsId'")
	}
	if args == nil || args.Content == nil {
		return nil, errors.New("missing required argument 'Content'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ApplicationInsightsId; i != nil { inputs["applicationInsightsId"] = i.ToStringOutput() }
		if i := args.Content; i != nil { inputs["content"] = i.ToStringOutput() }
		if i := args.FunctionAlias; i != nil { inputs["functionAlias"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Scope; i != nil { inputs["scope"] = i.ToStringOutput() }
		if i := args.Type; i != nil { inputs["type"] = i.ToStringOutput() }
	}
	var resource AnalyticsItem
	err := ctx.RegisterResource("azure:appinsights/analyticsItem:AnalyticsItem", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyticsItem gets an existing AnalyticsItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsItemState, opts ...pulumi.ResourceOption) (*AnalyticsItem, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ApplicationInsightsId; i != nil { inputs["applicationInsightsId"] = i.ToStringOutput() }
		if i := state.Content; i != nil { inputs["content"] = i.ToStringOutput() }
		if i := state.FunctionAlias; i != nil { inputs["functionAlias"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Scope; i != nil { inputs["scope"] = i.ToStringOutput() }
		if i := state.TimeCreated; i != nil { inputs["timeCreated"] = i.ToStringOutput() }
		if i := state.TimeModified; i != nil { inputs["timeModified"] = i.ToStringOutput() }
		if i := state.Type; i != nil { inputs["type"] = i.ToStringOutput() }
		if i := state.Version; i != nil { inputs["version"] = i.ToStringOutput() }
	}
	var resource AnalyticsItem
	err := ctx.ReadResource("azure:appinsights/analyticsItem:AnalyticsItem", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalyticsItem resources.
type AnalyticsItemState struct {
	// The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringInput `pulumi:"applicationInsightsId"`
	// The content for the Analytics Item, for example the query text if `type` is `query`.
	Content pulumi.StringInput `pulumi:"content"`
	// The alias to use for the function. Required when `type` is `function`.
	FunctionAlias pulumi.StringInput `pulumi:"functionAlias"`
	// Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
	Scope pulumi.StringInput `pulumi:"scope"`
	// A string containing the time the Analytics Item was created.
	TimeCreated pulumi.StringInput `pulumi:"timeCreated"`
	// A string containing the time the Analytics Item was last modified.
	TimeModified pulumi.StringInput `pulumi:"timeModified"`
	// The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
	Type pulumi.StringInput `pulumi:"type"`
	// A string indicating the version of the query format
	Version pulumi.StringInput `pulumi:"version"`
}

// The set of arguments for constructing a AnalyticsItem resource.
type AnalyticsItemArgs struct {
	// The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringInput `pulumi:"applicationInsightsId"`
	// The content for the Analytics Item, for example the query text if `type` is `query`.
	Content pulumi.StringInput `pulumi:"content"`
	// The alias to use for the function. Required when `type` is `function`.
	FunctionAlias pulumi.StringInput `pulumi:"functionAlias"`
	// Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
	Scope pulumi.StringInput `pulumi:"scope"`
	// The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
	Type pulumi.StringInput `pulumi:"type"`
}
