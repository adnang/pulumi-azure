// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Managed Disk.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/managed_disk.html.markdown.
func LookupManagedDisk(ctx *pulumi.Context, args *LookupManagedDiskArgs, opts ...pulumi.InvokeOption) (*LookupManagedDiskResult, error) {
	var rv LookupManagedDiskResult
	err := ctx.Invoke("azure:compute/getManagedDisk:getManagedDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getManagedDisk.
type LookupManagedDiskArgs struct {
	// Specifies the name of the Managed Disk.
	Name string `pulumi:"name"`
	// Specifies the name of the Resource Group where this Managed Disk exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Tags map[string]string `pulumi:"tags"`
	Zones []string `pulumi:"zones"`
}


// A collection of values returned by getManagedDisk.
type LookupManagedDiskResult struct {
	CreateOption string `pulumi:"createOption"`
	// The ID of the Disk Encryption Set used to encrypt this Managed Disk.
	DiskEncryptionSetId string `pulumi:"diskEncryptionSetId"`
	// The number of IOPS allowed for this disk, where one operation can transfer between 4k and 256k bytes.
	DiskIopsReadWrite int `pulumi:"diskIopsReadWrite"`
	// The bandwidth allowed for this disk.
	DiskMbpsReadWrite int `pulumi:"diskMbpsReadWrite"`
	// The size of the Managed Disk in gigabytes.
	DiskSizeGb int `pulumi:"diskSizeGb"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The operating system used for this Managed Disk.
	OsType string `pulumi:"osType"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of an existing Managed Disk which this Disk was created from.
	SourceResourceId string `pulumi:"sourceResourceId"`
	// The Source URI for this Managed Disk.
	SourceUri string `pulumi:"sourceUri"`
	// The ID of the Storage Account where the `sourceUri` is located.
	StorageAccountId string `pulumi:"storageAccountId"`
	// The storage account type for the Managed Disk.
	StorageAccountType string `pulumi:"storageAccountType"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A list of Availability Zones where the Managed Disk exists.
	Zones []string `pulumi:"zones"`
}

