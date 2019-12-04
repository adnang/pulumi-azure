// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bot

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Bot Channels Registration.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/bot_channels_registration.html.markdown.
type ChannelsRegistration struct {
	pulumi.CustomResourceState

	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey pulumi.StringOutput `pulumi:"developerAppInsightsApiKey"`

	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId pulumi.StringOutput `pulumi:"developerAppInsightsApplicationId"`

	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey pulumi.StringOutput `pulumi:"developerAppInsightsKey"`

	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified. 
	DisplayName pulumi.StringOutput `pulumi:"displayName"`

	// The Bot Channels Registration endpoint.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringOutput `pulumi:"microsoftAppId"`

	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringOutput `pulumi:"sku"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewChannelsRegistration registers a new resource with the given unique name, arguments, and options.
func NewChannelsRegistration(ctx *pulumi.Context,
	name string, args *ChannelsRegistrationArgs, opts ...pulumi.ResourceOption) (*ChannelsRegistration, error) {
	if args == nil || args.MicrosoftAppId == nil {
		return nil, errors.New("missing required argument 'MicrosoftAppId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.DeveloperAppInsightsApiKey; i != nil { inputs["developerAppInsightsApiKey"] = i.ToStringOutput() }
		if i := args.DeveloperAppInsightsApplicationId; i != nil { inputs["developerAppInsightsApplicationId"] = i.ToStringOutput() }
		if i := args.DeveloperAppInsightsKey; i != nil { inputs["developerAppInsightsKey"] = i.ToStringOutput() }
		if i := args.DisplayName; i != nil { inputs["displayName"] = i.ToStringOutput() }
		if i := args.Endpoint; i != nil { inputs["endpoint"] = i.ToStringOutput() }
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.MicrosoftAppId; i != nil { inputs["microsoftAppId"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Sku; i != nil { inputs["sku"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource ChannelsRegistration
	err := ctx.RegisterResource("azure:bot/channelsRegistration:ChannelsRegistration", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannelsRegistration gets an existing ChannelsRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannelsRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelsRegistrationState, opts ...pulumi.ResourceOption) (*ChannelsRegistration, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.DeveloperAppInsightsApiKey; i != nil { inputs["developerAppInsightsApiKey"] = i.ToStringOutput() }
		if i := state.DeveloperAppInsightsApplicationId; i != nil { inputs["developerAppInsightsApplicationId"] = i.ToStringOutput() }
		if i := state.DeveloperAppInsightsKey; i != nil { inputs["developerAppInsightsKey"] = i.ToStringOutput() }
		if i := state.DisplayName; i != nil { inputs["displayName"] = i.ToStringOutput() }
		if i := state.Endpoint; i != nil { inputs["endpoint"] = i.ToStringOutput() }
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.MicrosoftAppId; i != nil { inputs["microsoftAppId"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Sku; i != nil { inputs["sku"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource ChannelsRegistration
	err := ctx.ReadResource("azure:bot/channelsRegistration:ChannelsRegistration", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChannelsRegistration resources.
type ChannelsRegistrationState struct {
	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey pulumi.StringInput `pulumi:"developerAppInsightsApiKey"`
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId pulumi.StringInput `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey pulumi.StringInput `pulumi:"developerAppInsightsKey"`
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified. 
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// The Bot Channels Registration endpoint.
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringInput `pulumi:"microsoftAppId"`
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringInput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a ChannelsRegistration resource.
type ChannelsRegistrationArgs struct {
	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey pulumi.StringInput `pulumi:"developerAppInsightsApiKey"`
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId pulumi.StringInput `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey pulumi.StringInput `pulumi:"developerAppInsightsKey"`
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified. 
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// The Bot Channels Registration endpoint.
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringInput `pulumi:"microsoftAppId"`
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringInput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
