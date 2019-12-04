// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an AzureNetwork DDoS Protection Plan.
// 
// > **NOTE** Azure only allows `one` DDoS Protection Plan per region.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_ddos_protection_plan.html.markdown.
type DdosProtectionPlan struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// Specifies the name of the Network DDoS Protection Plan. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The Resource ID list of the Virtual Networks associated with DDoS Protection Plan.
	VirtualNetworkIds pulumi.StringArrayOutput `pulumi:"virtualNetworkIds"`
}

// NewDdosProtectionPlan registers a new resource with the given unique name, arguments, and options.
func NewDdosProtectionPlan(ctx *pulumi.Context,
	name string, args *DdosProtectionPlanArgs, opts ...pulumi.ResourceOption) (*DdosProtectionPlan, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource DdosProtectionPlan
	err := ctx.RegisterResource("azure:network/ddosProtectionPlan:DdosProtectionPlan", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdosProtectionPlan gets an existing DdosProtectionPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdosProtectionPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosProtectionPlanState, opts ...pulumi.ResourceOption) (*DdosProtectionPlan, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.VirtualNetworkIds; i != nil { inputs["virtualNetworkIds"] = i.ToStringArrayOutput() }
	}
	var resource DdosProtectionPlan
	err := ctx.ReadResource("azure:network/ddosProtectionPlan:DdosProtectionPlan", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdosProtectionPlan resources.
type DdosProtectionPlanState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name of the Network DDoS Protection Plan. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The Resource ID list of the Virtual Networks associated with DDoS Protection Plan.
	VirtualNetworkIds pulumi.StringArrayInput `pulumi:"virtualNetworkIds"`
}

// The set of arguments for constructing a DdosProtectionPlan resource.
type DdosProtectionPlanArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name of the Network DDoS Protection Plan. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
