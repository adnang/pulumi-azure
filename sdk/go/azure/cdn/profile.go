// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a CDN Profile to create a collection of CDN Endpoints.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/cdn_profile.html.markdown.
type Profile struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// Specifies the name of the CDN Profile. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the resource group in which to
	// create the CDN Profile.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The pricing related information of current CDN profile. Accepted values are `Standard_Akamai`, `Standard_ChinaCdn`, `Standard_Microsoft`, `Standard_Verizon` or `Premium_Verizon`.
	Sku pulumi.StringOutput `pulumi:"sku"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Sku; i != nil { inputs["sku"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource Profile
	err := ctx.RegisterResource("azure:cdn/profile:Profile", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Sku; i != nil { inputs["sku"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource Profile
	err := ctx.ReadResource("azure:cdn/profile:Profile", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type ProfileState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name of the CDN Profile. Changing this forces a
	// new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to
	// create the CDN Profile.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The pricing related information of current CDN profile. Accepted values are `Standard_Akamai`, `Standard_ChinaCdn`, `Standard_Microsoft`, `Standard_Verizon` or `Premium_Verizon`.
	Sku pulumi.StringInput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name of the CDN Profile. Changing this forces a
	// new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to
	// create the CDN Profile.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The pricing related information of current CDN profile. Accepted values are `Standard_Akamai`, `Standard_ChinaCdn`, `Standard_Microsoft`, `Standard_Verizon` or `Premium_Verizon`.
	Sku pulumi.StringInput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
