// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Container within an Azure Storage Account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_container.html.markdown.
type Container struct {
	pulumi.CustomResourceState

	// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
	ContainerAccessType pulumi.StringOutput `pulumi:"containerAccessType"`

	// Is there an Immutability Policy configured on this Storage Container?
	HasImmutabilityPolicy pulumi.BoolOutput `pulumi:"hasImmutabilityPolicy"`

	// Is there a Legal Hold configured on this Storage Container?
	HasLegalHold pulumi.BoolOutput `pulumi:"hasLegalHold"`

	// A mapping of MetaData for this Container.
	Metadata pulumi.MapOutput `pulumi:"metadata"`

	// The name of the Container which should be created within the Storage Account.
	Name pulumi.StringOutput `pulumi:"name"`

	// (**Deprecated**) Key-value definition of additional properties associated to the Storage Container
	Properties pulumi.StringMapOutput `pulumi:"properties"`

	// The name of the resource group in which to create the storage container. This field is no longer used and will be removed in 2.0. 
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The name of the Storage Account where the Container should be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
}

// NewContainer registers a new resource with the given unique name, arguments, and options.
func NewContainer(ctx *pulumi.Context,
	name string, args *ContainerArgs, opts ...pulumi.ResourceOption) (*Container, error) {
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ContainerAccessType; i != nil { inputs["containerAccessType"] = i.ToStringOutput() }
		if i := args.Metadata; i != nil { inputs["metadata"] = i.ToMapOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.StorageAccountName; i != nil { inputs["storageAccountName"] = i.ToStringOutput() }
	}
	var resource Container
	err := ctx.RegisterResource("azure:storage/container:Container", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainer gets an existing Container resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerState, opts ...pulumi.ResourceOption) (*Container, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ContainerAccessType; i != nil { inputs["containerAccessType"] = i.ToStringOutput() }
		if i := state.HasImmutabilityPolicy; i != nil { inputs["hasImmutabilityPolicy"] = i.ToBoolOutput() }
		if i := state.HasLegalHold; i != nil { inputs["hasLegalHold"] = i.ToBoolOutput() }
		if i := state.Metadata; i != nil { inputs["metadata"] = i.ToMapOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Properties; i != nil { inputs["properties"] = i.ToStringMapOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.StorageAccountName; i != nil { inputs["storageAccountName"] = i.ToStringOutput() }
	}
	var resource Container
	err := ctx.ReadResource("azure:storage/container:Container", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Container resources.
type ContainerState struct {
	// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
	ContainerAccessType pulumi.StringInput `pulumi:"containerAccessType"`
	// Is there an Immutability Policy configured on this Storage Container?
	HasImmutabilityPolicy pulumi.BoolInput `pulumi:"hasImmutabilityPolicy"`
	// Is there a Legal Hold configured on this Storage Container?
	HasLegalHold pulumi.BoolInput `pulumi:"hasLegalHold"`
	// A mapping of MetaData for this Container.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The name of the Container which should be created within the Storage Account.
	Name pulumi.StringInput `pulumi:"name"`
	// (**Deprecated**) Key-value definition of additional properties associated to the Storage Container
	Properties pulumi.StringMapInput `pulumi:"properties"`
	// The name of the resource group in which to create the storage container. This field is no longer used and will be removed in 2.0. 
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Storage Account where the Container should be created.
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a Container resource.
type ContainerArgs struct {
	// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
	ContainerAccessType pulumi.StringInput `pulumi:"containerAccessType"`
	// A mapping of MetaData for this Container.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The name of the Container which should be created within the Storage Account.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the storage container. This field is no longer used and will be removed in 2.0. 
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Storage Account where the Container should be created.
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}
