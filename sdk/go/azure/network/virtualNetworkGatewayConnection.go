// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a connection in an existing Virtual Network Gateway.
type VirtualNetworkGatewayConnection struct {
	s *pulumi.ResourceState
}

// NewVirtualNetworkGatewayConnection registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkGatewayConnection(ctx *pulumi.Context,
	name string, args *VirtualNetworkGatewayConnectionArgs, opts ...pulumi.ResourceOpt) (*VirtualNetworkGatewayConnection, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.VirtualNetworkGatewayId == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkGatewayId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["authorizationKey"] = nil
		inputs["enableBgp"] = nil
		inputs["expressRouteCircuitId"] = nil
		inputs["ipsecPolicy"] = nil
		inputs["localNetworkGatewayId"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["peerVirtualNetworkGatewayId"] = nil
		inputs["resourceGroupName"] = nil
		inputs["routingWeight"] = nil
		inputs["sharedKey"] = nil
		inputs["tags"] = nil
		inputs["type"] = nil
		inputs["usePolicyBasedTrafficSelectors"] = nil
		inputs["virtualNetworkGatewayId"] = nil
	} else {
		inputs["authorizationKey"] = args.AuthorizationKey
		inputs["enableBgp"] = args.EnableBgp
		inputs["expressRouteCircuitId"] = args.ExpressRouteCircuitId
		inputs["ipsecPolicy"] = args.IpsecPolicy
		inputs["localNetworkGatewayId"] = args.LocalNetworkGatewayId
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["peerVirtualNetworkGatewayId"] = args.PeerVirtualNetworkGatewayId
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["routingWeight"] = args.RoutingWeight
		inputs["sharedKey"] = args.SharedKey
		inputs["tags"] = args.Tags
		inputs["type"] = args.Type
		inputs["usePolicyBasedTrafficSelectors"] = args.UsePolicyBasedTrafficSelectors
		inputs["virtualNetworkGatewayId"] = args.VirtualNetworkGatewayId
	}
	s, err := ctx.RegisterResource("azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualNetworkGatewayConnection{s: s}, nil
}

// GetVirtualNetworkGatewayConnection gets an existing VirtualNetworkGatewayConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkGatewayConnection(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VirtualNetworkGatewayConnectionState, opts ...pulumi.ResourceOpt) (*VirtualNetworkGatewayConnection, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["authorizationKey"] = state.AuthorizationKey
		inputs["enableBgp"] = state.EnableBgp
		inputs["expressRouteCircuitId"] = state.ExpressRouteCircuitId
		inputs["ipsecPolicy"] = state.IpsecPolicy
		inputs["localNetworkGatewayId"] = state.LocalNetworkGatewayId
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["peerVirtualNetworkGatewayId"] = state.PeerVirtualNetworkGatewayId
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["routingWeight"] = state.RoutingWeight
		inputs["sharedKey"] = state.SharedKey
		inputs["tags"] = state.Tags
		inputs["type"] = state.Type
		inputs["usePolicyBasedTrafficSelectors"] = state.UsePolicyBasedTrafficSelectors
		inputs["virtualNetworkGatewayId"] = state.VirtualNetworkGatewayId
	}
	s, err := ctx.ReadResource("azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualNetworkGatewayConnection{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VirtualNetworkGatewayConnection) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VirtualNetworkGatewayConnection) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The authorization key associated with the
// Express Route Circuit. This field is required only if the type is an
// ExpressRoute connection.
func (r *VirtualNetworkGatewayConnection) AuthorizationKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["authorizationKey"])
}

// If `true`, BGP (Border Gateway Protocol) is enabled
// for this connection. Defaults to `false`.
func (r *VirtualNetworkGatewayConnection) EnableBgp() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableBgp"])
}

// The ID of the Express Route Circuit
// when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
// The Express Route Circuit can be in the same or in a different subscription.
func (r *VirtualNetworkGatewayConnection) ExpressRouteCircuitId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["expressRouteCircuitId"])
}

// A `ipsec_policy` block which is documented below.
// Only a single policy can be defined for a connection. For details on
// custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
func (r *VirtualNetworkGatewayConnection) IpsecPolicy() *pulumi.Output {
	return r.s.State["ipsecPolicy"]
}

// The ID of the local network gateway
// when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
func (r *VirtualNetworkGatewayConnection) LocalNetworkGatewayId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["localNetworkGatewayId"])
}

// The location/region where the connection is
// located. Changing this forces a new resource to be created.
func (r *VirtualNetworkGatewayConnection) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The name of the connection. Changing the name forces a
// new resource to be created.
func (r *VirtualNetworkGatewayConnection) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the peer virtual
// network gateway when creating a VNet-to-VNet connection (i.e. when `type`
// is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
// in a different subscription.
func (r *VirtualNetworkGatewayConnection) PeerVirtualNetworkGatewayId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["peerVirtualNetworkGatewayId"])
}

// The name of the resource group in which to
// create the connection Changing the name forces a new resource to be created.
func (r *VirtualNetworkGatewayConnection) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The routing weight. Defaults to `10`.
func (r *VirtualNetworkGatewayConnection) RoutingWeight() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["routingWeight"])
}

// The shared IPSec key. A key must be provided if a
// Site-to-Site or VNet-to-VNet connection is created whereas ExpressRoute
// connections do not need a shared key.
func (r *VirtualNetworkGatewayConnection) SharedKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sharedKey"])
}

// A mapping of tags to assign to the resource.
func (r *VirtualNetworkGatewayConnection) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The type of connection. Valid options are `IPsec`
// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
// Each connection type requires different mandatory arguments (refer to the
// examples above). Changing the connection type will force a new connection
// to be created.
func (r *VirtualNetworkGatewayConnection) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// If `true`, policy-based traffic
// selectors are enabled for this connection. Enabling policy-based traffic
// selectors requires an `ipsec_policy` block. Defaults to `false`.
func (r *VirtualNetworkGatewayConnection) UsePolicyBasedTrafficSelectors() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["usePolicyBasedTrafficSelectors"])
}

// The ID of the Virtual Network Gateway
// in which the connection will be created. Changing the gateway forces a new
// resource to be created.
func (r *VirtualNetworkGatewayConnection) VirtualNetworkGatewayId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["virtualNetworkGatewayId"])
}

// Input properties used for looking up and filtering VirtualNetworkGatewayConnection resources.
type VirtualNetworkGatewayConnectionState struct {
	// The authorization key associated with the
	// Express Route Circuit. This field is required only if the type is an
	// ExpressRoute connection.
	AuthorizationKey interface{}
	// If `true`, BGP (Border Gateway Protocol) is enabled
	// for this connection. Defaults to `false`.
	EnableBgp interface{}
	// The ID of the Express Route Circuit
	// when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
	// The Express Route Circuit can be in the same or in a different subscription.
	ExpressRouteCircuitId interface{}
	// A `ipsec_policy` block which is documented below.
	// Only a single policy can be defined for a connection. For details on
	// custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
	IpsecPolicy interface{}
	// The ID of the local network gateway
	// when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
	LocalNetworkGatewayId interface{}
	// The location/region where the connection is
	// located. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the connection. Changing the name forces a
	// new resource to be created.
	Name interface{}
	// The ID of the peer virtual
	// network gateway when creating a VNet-to-VNet connection (i.e. when `type`
	// is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
	// in a different subscription.
	PeerVirtualNetworkGatewayId interface{}
	// The name of the resource group in which to
	// create the connection Changing the name forces a new resource to be created.
	ResourceGroupName interface{}
	// The routing weight. Defaults to `10`.
	RoutingWeight interface{}
	// The shared IPSec key. A key must be provided if a
	// Site-to-Site or VNet-to-VNet connection is created whereas ExpressRoute
	// connections do not need a shared key.
	SharedKey interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The type of connection. Valid options are `IPsec`
	// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
	// Each connection type requires different mandatory arguments (refer to the
	// examples above). Changing the connection type will force a new connection
	// to be created.
	Type interface{}
	// If `true`, policy-based traffic
	// selectors are enabled for this connection. Enabling policy-based traffic
	// selectors requires an `ipsec_policy` block. Defaults to `false`.
	UsePolicyBasedTrafficSelectors interface{}
	// The ID of the Virtual Network Gateway
	// in which the connection will be created. Changing the gateway forces a new
	// resource to be created.
	VirtualNetworkGatewayId interface{}
}

// The set of arguments for constructing a VirtualNetworkGatewayConnection resource.
type VirtualNetworkGatewayConnectionArgs struct {
	// The authorization key associated with the
	// Express Route Circuit. This field is required only if the type is an
	// ExpressRoute connection.
	AuthorizationKey interface{}
	// If `true`, BGP (Border Gateway Protocol) is enabled
	// for this connection. Defaults to `false`.
	EnableBgp interface{}
	// The ID of the Express Route Circuit
	// when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
	// The Express Route Circuit can be in the same or in a different subscription.
	ExpressRouteCircuitId interface{}
	// A `ipsec_policy` block which is documented below.
	// Only a single policy can be defined for a connection. For details on
	// custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
	IpsecPolicy interface{}
	// The ID of the local network gateway
	// when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
	LocalNetworkGatewayId interface{}
	// The location/region where the connection is
	// located. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the connection. Changing the name forces a
	// new resource to be created.
	Name interface{}
	// The ID of the peer virtual
	// network gateway when creating a VNet-to-VNet connection (i.e. when `type`
	// is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
	// in a different subscription.
	PeerVirtualNetworkGatewayId interface{}
	// The name of the resource group in which to
	// create the connection Changing the name forces a new resource to be created.
	ResourceGroupName interface{}
	// The routing weight. Defaults to `10`.
	RoutingWeight interface{}
	// The shared IPSec key. A key must be provided if a
	// Site-to-Site or VNet-to-VNet connection is created whereas ExpressRoute
	// connections do not need a shared key.
	SharedKey interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The type of connection. Valid options are `IPsec`
	// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
	// Each connection type requires different mandatory arguments (refer to the
	// examples above). Changing the connection type will force a new connection
	// to be created.
	Type interface{}
	// If `true`, policy-based traffic
	// selectors are enabled for this connection. Enabling policy-based traffic
	// selectors requires an `ipsec_policy` block. Defaults to `false`.
	UsePolicyBasedTrafficSelectors interface{}
	// The ID of the Virtual Network Gateway
	// in which the connection will be created. Changing the gateway forces a new
	// resource to be created.
	VirtualNetworkGatewayId interface{}
}
