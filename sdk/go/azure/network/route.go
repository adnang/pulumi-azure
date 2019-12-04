// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Route within a Route Table.
// 
// > **NOTE on Route Tables and Routes:** This provider currently
// provides both a standalone Route resource, and allows for Routes to be defined in-line within the Route Table resource.
// At this time you cannot use a Route Table with in-line Routes in conjunction with any Route resources. Doing so will cause a conflict of Route configurations and will overwrite Routes.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/route.html.markdown.
type Route struct {
	pulumi.CustomResourceState

	// The destination CIDR to which the route applies, such as `10.1.0.0/16`
	AddressPrefix pulumi.StringOutput `pulumi:"addressPrefix"`

	// The name of the route. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
	NextHopInIpAddress pulumi.StringOutput `pulumi:"nextHopInIpAddress"`

	// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
	NextHopType pulumi.StringOutput `pulumi:"nextHopType"`

	// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The name of the route table within which create the route. Changing this forces a new resource to be created.
	RouteTableName pulumi.StringOutput `pulumi:"routeTableName"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil || args.AddressPrefix == nil {
		return nil, errors.New("missing required argument 'AddressPrefix'")
	}
	if args == nil || args.NextHopType == nil {
		return nil, errors.New("missing required argument 'NextHopType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RouteTableName == nil {
		return nil, errors.New("missing required argument 'RouteTableName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AddressPrefix; i != nil { inputs["addressPrefix"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.NextHopInIpAddress; i != nil { inputs["nextHopInIpAddress"] = i.ToStringOutput() }
		if i := args.NextHopType; i != nil { inputs["nextHopType"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.RouteTableName; i != nil { inputs["routeTableName"] = i.ToStringOutput() }
	}
	var resource Route
	err := ctx.RegisterResource("azure:network/route:Route", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AddressPrefix; i != nil { inputs["addressPrefix"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.NextHopInIpAddress; i != nil { inputs["nextHopInIpAddress"] = i.ToStringOutput() }
		if i := state.NextHopType; i != nil { inputs["nextHopType"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.RouteTableName; i != nil { inputs["routeTableName"] = i.ToStringOutput() }
	}
	var resource Route
	err := ctx.ReadResource("azure:network/route:Route", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type RouteState struct {
	// The destination CIDR to which the route applies, such as `10.1.0.0/16`
	AddressPrefix pulumi.StringInput `pulumi:"addressPrefix"`
	// The name of the route. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
	NextHopInIpAddress pulumi.StringInput `pulumi:"nextHopInIpAddress"`
	// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
	NextHopType pulumi.StringInput `pulumi:"nextHopType"`
	// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the route table within which create the route. Changing this forces a new resource to be created.
	RouteTableName pulumi.StringInput `pulumi:"routeTableName"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The destination CIDR to which the route applies, such as `10.1.0.0/16`
	AddressPrefix pulumi.StringInput `pulumi:"addressPrefix"`
	// The name of the route. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
	NextHopInIpAddress pulumi.StringInput `pulumi:"nextHopInIpAddress"`
	// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
	NextHopType pulumi.StringInput `pulumi:"nextHopType"`
	// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the route table within which create the route. Changing this forces a new resource to be created.
	RouteTableName pulumi.StringInput `pulumi:"routeTableName"`
}
