// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Route Table.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/route_table.html.markdown.
func LookupRouteTable(ctx *pulumi.Context, args *GetRouteTableArgs, opts ...pulumi.InvokeOption) (*GetRouteTableResult, error) {
	var rv GetRouteTableResult
	err := ctx.Invoke("azure:network/getRouteTable:getRouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteTable.
type GetRouteTableArgs struct {
	// The name of the Route Table.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the Route Table exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getRouteTable.
type GetRouteTableResult struct {
	// The Azure Region in which the Route Table exists.
	Location string `pulumi:"location"`
	// The name of the Route.
	Name string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `route` blocks as documented below.
	Routes []GetRouteTableRoutesResult `pulumi:"routes"`
	// The collection of Subnets associated with this route table.
	Subnets []string `pulumi:"subnets"`
	// A mapping of tags assigned to the Route Table.
	Tags map[string]string `pulumi:"tags"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
type GetRouteTableRoutesResult struct {
	// The destination CIDR to which the route applies.
	AddressPrefix string `pulumi:"addressPrefix"`
	// The name of the Route Table.
	Name string `pulumi:"name"`
	// Contains the IP address packets should be forwarded to.
	NextHopInIpAddress string `pulumi:"nextHopInIpAddress"`
	// The type of Azure hop the packet should be sent to.
	NextHopType string `pulumi:"nextHopType"`
}
