// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Uses this data source to access information about an existing Virtual Hub.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/virtual_hub.html.markdown.
func LookupVirtualHub(ctx *pulumi.Context, args *LookupVirtualHubArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubResult, error) {
	var rv LookupVirtualHubResult
	err := ctx.Invoke("azure:network/getVirtualHub:getVirtualHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualHub.
type LookupVirtualHubArgs struct {
	// The name of the Virtual Hub.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the Virtual Hub exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getVirtualHub.
type LookupVirtualHubResult struct {
	// The Address Prefix used for this Virtual Hub.
	AddressPrefix string `pulumi:"addressPrefix"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region where the Virtual Hub exists.
	Location string `pulumi:"location"`
	Name string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the Virtual Hub.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual WAN within which the Virtual Hub exists.
	VirtualWanId string `pulumi:"virtualWanId"`
}

