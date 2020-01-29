// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package storage

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Storage Container.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/storage_container.html.markdown.
func GetStorageContainer(ctx *pulumi.Context, args *GetStorageContainerArgs, opts ...pulumi.InvokeOption) (*GetStorageContainerResult, error) {
	var rv GetStorageContainerResult
	err := ctx.Invoke("azure:storage/getStorageContainer:getStorageContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStorageContainer.
type GetStorageContainerArgs struct {
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the Container.
	Name string `pulumi:"name"`
	// The name of the Storage Account where the Container was created.
	StorageAccountName string `pulumi:"storageAccountName"`
}


// A collection of values returned by getStorageContainer.
type GetStorageContainerResult struct {
	// The Access Level configured for this Container.
	ContainerAccessType string `pulumi:"containerAccessType"`
	// Is there an Immutability Policy configured on this Storage Container?
	HasImmutabilityPolicy bool `pulumi:"hasImmutabilityPolicy"`
	// Is there a Legal Hold configured on this Storage Container?
	HasLegalHold bool `pulumi:"hasLegalHold"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A mapping of MetaData for this Container.
	Metadata map[string]string `pulumi:"metadata"`
	Name string `pulumi:"name"`
	StorageAccountName string `pulumi:"storageAccountName"`
}

