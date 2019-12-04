// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Version of a Shared Image within a Shared Image Gallery.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/shared_image_version.html.markdown.
func LookupSharedImageVersion(ctx *pulumi.Context, args *GetSharedImageVersionArgs, opts ...pulumi.InvokeOption) (*GetSharedImageVersionResult, error) {
	var rv GetSharedImageVersionResult
	err := ctx.Invoke("azure:compute/getSharedImageVersion:getSharedImageVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSharedImageVersion.
type GetSharedImageVersionArgs struct {
	// The name of the Shared Image in which the Shared Image exists.
	GalleryName string `pulumi:"galleryName"`
	// The name of the Shared Image in which this Version exists.
	ImageName string `pulumi:"imageName"`
	// The name of the Image Version.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the Shared Image Gallery exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getSharedImageVersion.
type GetSharedImageVersionResult struct {
	// Is this Image Version excluded from the `latest` filter?
	ExcludeFromLatest bool `pulumi:"excludeFromLatest"`
	GalleryName string `pulumi:"galleryName"`
	ImageName string `pulumi:"imageName"`
	// The supported Azure location where the Shared Image Gallery exists.
	Location string `pulumi:"location"`
	// The ID of the Managed Image which was the source of this Shared Image Version.
	ManagedImageId string `pulumi:"managedImageId"`
	// The Azure Region in which this Image Version exists.
	Name string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the Shared Image.
	Tags map[string]string `pulumi:"tags"`
	// One or more `targetRegion` blocks as documented below.
	TargetRegions []GetSharedImageVersionTargetRegionsResult `pulumi:"targetRegions"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
type GetSharedImageVersionTargetRegionsResult struct {
	// The name of the Image Version.
	Name string `pulumi:"name"`
	// The number of replicas of the Image Version to be created per region.
	RegionalReplicaCount int `pulumi:"regionalReplicaCount"`
}
