// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Management Product Assignment to a Group.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_product_group.html.markdown.
type ProductGroup struct {
	pulumi.CustomResourceState

	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`

	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName pulumi.StringOutput `pulumi:"groupName"`

	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringOutput `pulumi:"productId"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewProductGroup registers a new resource with the given unique name, arguments, and options.
func NewProductGroup(ctx *pulumi.Context,
	name string, args *ProductGroupArgs, opts ...pulumi.ResourceOption) (*ProductGroup, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.GroupName == nil {
		return nil, errors.New("missing required argument 'GroupName'")
	}
	if args == nil || args.ProductId == nil {
		return nil, errors.New("missing required argument 'ProductId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ApiManagementName; i != nil { inputs["apiManagementName"] = i.ToStringOutput() }
		if i := args.GroupName; i != nil { inputs["groupName"] = i.ToStringOutput() }
		if i := args.ProductId; i != nil { inputs["productId"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
	}
	var resource ProductGroup
	err := ctx.RegisterResource("azure:apimanagement/productGroup:ProductGroup", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductGroup gets an existing ProductGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductGroupState, opts ...pulumi.ResourceOption) (*ProductGroup, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ApiManagementName; i != nil { inputs["apiManagementName"] = i.ToStringOutput() }
		if i := state.GroupName; i != nil { inputs["groupName"] = i.ToStringOutput() }
		if i := state.ProductId; i != nil { inputs["productId"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
	}
	var resource ProductGroup
	err := ctx.ReadResource("azure:apimanagement/productGroup:ProductGroup", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductGroup resources.
type ProductGroupState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput `pulumi:"apiManagementName"`
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName pulumi.StringInput `pulumi:"groupName"`
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringInput `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ProductGroup resource.
type ProductGroupArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput `pulumi:"apiManagementName"`
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName pulumi.StringInput `pulumi:"groupName"`
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringInput `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}
