// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Directory within an Azure Storage File Share.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_share_directory.html.markdown.
type ShareDirectory struct {
	pulumi.CustomResourceState

	// A mapping of metadata to assign to this Directory.
	Metadata pulumi.MapOutput `pulumi:"metadata"`

	// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
	ShareName pulumi.StringOutput `pulumi:"shareName"`

	// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
}

// NewShareDirectory registers a new resource with the given unique name, arguments, and options.
func NewShareDirectory(ctx *pulumi.Context,
	name string, args *ShareDirectoryArgs, opts ...pulumi.ResourceOption) (*ShareDirectory, error) {
	if args == nil || args.ShareName == nil {
		return nil, errors.New("missing required argument 'ShareName'")
	}
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Metadata; i != nil { inputs["metadata"] = i.ToMapOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ShareName; i != nil { inputs["shareName"] = i.ToStringOutput() }
		if i := args.StorageAccountName; i != nil { inputs["storageAccountName"] = i.ToStringOutput() }
	}
	var resource ShareDirectory
	err := ctx.RegisterResource("azure:storage/shareDirectory:ShareDirectory", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShareDirectory gets an existing ShareDirectory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShareDirectory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareDirectoryState, opts ...pulumi.ResourceOption) (*ShareDirectory, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Metadata; i != nil { inputs["metadata"] = i.ToMapOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ShareName; i != nil { inputs["shareName"] = i.ToStringOutput() }
		if i := state.StorageAccountName; i != nil { inputs["storageAccountName"] = i.ToStringOutput() }
	}
	var resource ShareDirectory
	err := ctx.ReadResource("azure:storage/shareDirectory:ShareDirectory", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShareDirectory resources.
type ShareDirectoryState struct {
	// A mapping of metadata to assign to this Directory.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
	ShareName pulumi.StringInput `pulumi:"shareName"`
	// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a ShareDirectory resource.
type ShareDirectoryArgs struct {
	// A mapping of metadata to assign to this Directory.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
	ShareName pulumi.StringInput `pulumi:"shareName"`
	// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}
