// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an availability set for virtual machines.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/availability_set.html.markdown.
type AvailabilitySet struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `false`.
	Managed pulumi.BoolOutput `pulumi:"managed"`

	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the number of fault domains that are used. Defaults to 3.
	PlatformFaultDomainCount pulumi.IntOutput `pulumi:"platformFaultDomainCount"`

	// Specifies the number of update domains that are used. Defaults to 5.
	PlatformUpdateDomainCount pulumi.IntOutput `pulumi:"platformUpdateDomainCount"`

	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringOutput `pulumi:"proximityPlacementGroupId"`

	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewAvailabilitySet registers a new resource with the given unique name, arguments, and options.
func NewAvailabilitySet(ctx *pulumi.Context,
	name string, args *AvailabilitySetArgs, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.Managed; i != nil { inputs["managed"] = i.ToBoolOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.PlatformFaultDomainCount; i != nil { inputs["platformFaultDomainCount"] = i.ToIntOutput() }
		if i := args.PlatformUpdateDomainCount; i != nil { inputs["platformUpdateDomainCount"] = i.ToIntOutput() }
		if i := args.ProximityPlacementGroupId; i != nil { inputs["proximityPlacementGroupId"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource AvailabilitySet
	err := ctx.RegisterResource("azure:compute/availabilitySet:AvailabilitySet", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAvailabilitySet gets an existing AvailabilitySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAvailabilitySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilitySetState, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Managed; i != nil { inputs["managed"] = i.ToBoolOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.PlatformFaultDomainCount; i != nil { inputs["platformFaultDomainCount"] = i.ToIntOutput() }
		if i := state.PlatformUpdateDomainCount; i != nil { inputs["platformUpdateDomainCount"] = i.ToIntOutput() }
		if i := state.ProximityPlacementGroupId; i != nil { inputs["proximityPlacementGroupId"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource AvailabilitySet
	err := ctx.ReadResource("azure:compute/availabilitySet:AvailabilitySet", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AvailabilitySet resources.
type AvailabilitySetState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `false`.
	Managed pulumi.BoolInput `pulumi:"managed"`
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the number of fault domains that are used. Defaults to 3.
	PlatformFaultDomainCount pulumi.IntInput `pulumi:"platformFaultDomainCount"`
	// Specifies the number of update domains that are used. Defaults to 5.
	PlatformUpdateDomainCount pulumi.IntInput `pulumi:"platformUpdateDomainCount"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringInput `pulumi:"proximityPlacementGroupId"`
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a AvailabilitySet resource.
type AvailabilitySetArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `false`.
	Managed pulumi.BoolInput `pulumi:"managed"`
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the number of fault domains that are used. Defaults to 3.
	PlatformFaultDomainCount pulumi.IntInput `pulumi:"platformFaultDomainCount"`
	// Specifies the number of update domains that are used. Defaults to 5.
	PlatformUpdateDomainCount pulumi.IntInput `pulumi:"platformUpdateDomainCount"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringInput `pulumi:"proximityPlacementGroupId"`
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
