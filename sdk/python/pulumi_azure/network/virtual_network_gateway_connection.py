# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class VirtualNetworkGatewayConnection(pulumi.CustomResource):
    """
    Manages a connection in an existing Virtual Network Gateway.
    """
    def __init__(__self__, __name__, __opts__=None, authorization_key=None, enable_bgp=None, express_route_circuit_id=None, ipsec_policy=None, local_network_gateway_id=None, location=None, name=None, peer_virtual_network_gateway_id=None, resource_group_name=None, routing_weight=None, shared_key=None, tags=None, type=None, use_policy_based_traffic_selectors=None, virtual_network_gateway_id=None):
        """Create a VirtualNetworkGatewayConnection resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if authorization_key and not isinstance(authorization_key, basestring):
            raise TypeError('Expected property authorization_key to be a basestring')
        __self__.authorization_key = authorization_key
        """
        The authorization key associated with the
        Express Route Circuit. This field is required only if the type is an
        ExpressRoute connection.
        """
        __props__['authorizationKey'] = authorization_key

        if enable_bgp and not isinstance(enable_bgp, bool):
            raise TypeError('Expected property enable_bgp to be a bool')
        __self__.enable_bgp = enable_bgp
        """
        If `true`, BGP (Border Gateway Protocol) is enabled
        for this connection. Defaults to `false`.
        """
        __props__['enableBgp'] = enable_bgp

        if express_route_circuit_id and not isinstance(express_route_circuit_id, basestring):
            raise TypeError('Expected property express_route_circuit_id to be a basestring')
        __self__.express_route_circuit_id = express_route_circuit_id
        """
        The ID of the Express Route Circuit
        when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
        The Express Route Circuit can be in the same or in a different subscription.
        """
        __props__['expressRouteCircuitId'] = express_route_circuit_id

        if ipsec_policy and not isinstance(ipsec_policy, dict):
            raise TypeError('Expected property ipsec_policy to be a dict')
        __self__.ipsec_policy = ipsec_policy
        """
        A `ipsec_policy` block which is documented below.
        Only a single policy can be defined for a connection. For details on
        custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
        """
        __props__['ipsecPolicy'] = ipsec_policy

        if local_network_gateway_id and not isinstance(local_network_gateway_id, basestring):
            raise TypeError('Expected property local_network_gateway_id to be a basestring')
        __self__.local_network_gateway_id = local_network_gateway_id
        """
        The ID of the local network gateway
        when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
        """
        __props__['localNetworkGatewayId'] = local_network_gateway_id

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        The location/region where the connection is
        located. Changing this forces a new resource to be created.
        """
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the connection. Changing the name forces a
        new resource to be created.
        """
        __props__['name'] = name

        if peer_virtual_network_gateway_id and not isinstance(peer_virtual_network_gateway_id, basestring):
            raise TypeError('Expected property peer_virtual_network_gateway_id to be a basestring')
        __self__.peer_virtual_network_gateway_id = peer_virtual_network_gateway_id
        """
        The ID of the peer virtual
        network gateway when creating a VNet-to-VNet connection (i.e. when `type`
        is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
        in a different subscription.
        """
        __props__['peerVirtualNetworkGatewayId'] = peer_virtual_network_gateway_id

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to
        create the connection Changing the name forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if routing_weight and not isinstance(routing_weight, int):
            raise TypeError('Expected property routing_weight to be a int')
        __self__.routing_weight = routing_weight
        """
        The routing weight. Defaults to `10`.
        """
        __props__['routingWeight'] = routing_weight

        if shared_key and not isinstance(shared_key, basestring):
            raise TypeError('Expected property shared_key to be a basestring')
        __self__.shared_key = shared_key
        """
        The shared IPSec key. A key must be provided if a
        Site-to-Site or VNet-to-VNet connection is created whereas ExpressRoute
        connections do not need a shared key.
        """
        __props__['sharedKey'] = shared_key

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        __props__['tags'] = tags

        if not type:
            raise TypeError('Missing required property type')
        elif not isinstance(type, basestring):
            raise TypeError('Expected property type to be a basestring')
        __self__.type = type
        """
        The type of connection. Valid options are `IPsec`
        (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
        Each connection type requires different mandatory arguments (refer to the
        examples above). Changing the connection type will force a new connection
        to be created.
        """
        __props__['type'] = type

        if use_policy_based_traffic_selectors and not isinstance(use_policy_based_traffic_selectors, bool):
            raise TypeError('Expected property use_policy_based_traffic_selectors to be a bool')
        __self__.use_policy_based_traffic_selectors = use_policy_based_traffic_selectors
        """
        If `true`, policy-based traffic
        selectors are enabled for this connection. Enabling policy-based traffic
        selectors requires an `ipsec_policy` block. Defaults to `false`.
        """
        __props__['usePolicyBasedTrafficSelectors'] = use_policy_based_traffic_selectors

        if not virtual_network_gateway_id:
            raise TypeError('Missing required property virtual_network_gateway_id')
        elif not isinstance(virtual_network_gateway_id, basestring):
            raise TypeError('Expected property virtual_network_gateway_id to be a basestring')
        __self__.virtual_network_gateway_id = virtual_network_gateway_id
        """
        The ID of the Virtual Network Gateway
        in which the connection will be created. Changing the gateway forces a new
        resource to be created.
        """
        __props__['virtualNetworkGatewayId'] = virtual_network_gateway_id

        super(VirtualNetworkGatewayConnection, __self__).__init__(
            'azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'authorizationKey' in outs:
            self.authorization_key = outs['authorizationKey']
        if 'enableBgp' in outs:
            self.enable_bgp = outs['enableBgp']
        if 'expressRouteCircuitId' in outs:
            self.express_route_circuit_id = outs['expressRouteCircuitId']
        if 'ipsecPolicy' in outs:
            self.ipsec_policy = outs['ipsecPolicy']
        if 'localNetworkGatewayId' in outs:
            self.local_network_gateway_id = outs['localNetworkGatewayId']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'peerVirtualNetworkGatewayId' in outs:
            self.peer_virtual_network_gateway_id = outs['peerVirtualNetworkGatewayId']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'routingWeight' in outs:
            self.routing_weight = outs['routingWeight']
        if 'sharedKey' in outs:
            self.shared_key = outs['sharedKey']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'type' in outs:
            self.type = outs['type']
        if 'usePolicyBasedTrafficSelectors' in outs:
            self.use_policy_based_traffic_selectors = outs['usePolicyBasedTrafficSelectors']
        if 'virtualNetworkGatewayId' in outs:
            self.virtual_network_gateway_id = outs['virtualNetworkGatewayId']
