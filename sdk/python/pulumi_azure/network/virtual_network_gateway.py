# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class VirtualNetworkGateway(pulumi.CustomResource):
    """
    Manages a Virtual Network Gateway to establish secure, cross-premises connectivity.
    
    -> **Note:** Please be aware that provisioning a Virtual Network Gateway takes a long time (between 30 minutes and 1 hour)
    """
    def __init__(__self__, __name__, __opts__=None, active_active=None, bgp_settings=None, default_local_network_gateway_id=None, enable_bgp=None, ip_configurations=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, type=None, vpn_client_configuration=None, vpn_type=None):
        """Create a VirtualNetworkGateway resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if active_active and not isinstance(active_active, bool):
            raise TypeError('Expected property active_active to be a bool')
        __self__.active_active = active_active
        """
        If `true`, an active-active Virtual Network Gateway
        will be created. An active-active gateway requires a `HighPerformance` or an
        `UltraPerformance` sku. If `false`, an active-standby gateway will be created.
        Defaults to `false`.
        """
        __props__['activeActive'] = active_active

        if bgp_settings and not isinstance(bgp_settings, dict):
            raise TypeError('Expected property bgp_settings to be a dict')
        __self__.bgp_settings = bgp_settings
        __props__['bgpSettings'] = bgp_settings

        if default_local_network_gateway_id and not isinstance(default_local_network_gateway_id, basestring):
            raise TypeError('Expected property default_local_network_gateway_id to be a basestring')
        __self__.default_local_network_gateway_id = default_local_network_gateway_id
        """
        The ID of the local network gateway
        through which outbound Internet traffic from the virtual network in which the
        gateway is created will be routed (*forced tunneling*). Refer to the
        [Azure documentation on forced tunneling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
        If not specified, forced tunneling is disabled.
        """
        __props__['defaultLocalNetworkGatewayId'] = default_local_network_gateway_id

        if enable_bgp and not isinstance(enable_bgp, bool):
            raise TypeError('Expected property enable_bgp to be a bool')
        __self__.enable_bgp = enable_bgp
        """
        If `true`, BGP (Border Gateway Protocol) will be enabled
        for this Virtual Network Gateway. Defaults to `false`.
        """
        __props__['enableBgp'] = enable_bgp

        if not ip_configurations:
            raise TypeError('Missing required property ip_configurations')
        elif not isinstance(ip_configurations, list):
            raise TypeError('Expected property ip_configurations to be a list')
        __self__.ip_configurations = ip_configurations
        """
        One or two `ip_configuration` blocks documented below.
        An active-standby gateway requires exactly one `ip_configuration` block whereas
        an active-active gateway requires exactly two `ip_configuration` blocks.
        """
        __props__['ipConfigurations'] = ip_configurations

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        The location/region where the Virtual Network Gateway is
        located. Changing the location/region forces a new resource to be created.
        """
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the Virtual Network Gateway. Changing the name
        forces a new resource to be created.
        """
        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to
        create the Virtual Network Gateway. Changing the resource group name forces
        a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if not sku:
            raise TypeError('Missing required property sku')
        elif not isinstance(sku, basestring):
            raise TypeError('Expected property sku to be a basestring')
        __self__.sku = sku
        """
        Configuration of the size and capacity of the virtual network
        gateway. Valid options are `Basic`, `Standard`, `HighPerformance`, `UltraPerformance`,
        `VpnGw1`, `VpnGw2` and `VpnGw3` and depend on the `type` and `vpn_type` arguments.
        A `PolicyBased` gateway only supports the `Basic` sku. Further, the `UltraPerformance`
        sku is only supported by an `ExpressRoute` gateway.
        """
        __props__['sku'] = sku

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
        The type of the Virtual Network Gateway. Valid options are
        `Vpn` or `ExpressRoute`. Changing the type forces a new resource to be created.
        """
        __props__['type'] = type

        if vpn_client_configuration and not isinstance(vpn_client_configuration, dict):
            raise TypeError('Expected property vpn_client_configuration to be a dict')
        __self__.vpn_client_configuration = vpn_client_configuration
        """
        A `vpn_client_configuration` block which
        is documented below. In this block the Virtual Network Gateway can be configured
        to accept IPSec point-to-site connections.
        """
        __props__['vpnClientConfiguration'] = vpn_client_configuration

        if vpn_type and not isinstance(vpn_type, basestring):
            raise TypeError('Expected property vpn_type to be a basestring')
        __self__.vpn_type = vpn_type
        """
        The routing type of the Virtual Network Gateway. Valid
        options are `RouteBased` or `PolicyBased`. Defaults to `RouteBased`.
        """
        __props__['vpnType'] = vpn_type

        super(VirtualNetworkGateway, __self__).__init__(
            'azure:network/virtualNetworkGateway:VirtualNetworkGateway',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'activeActive' in outs:
            self.active_active = outs['activeActive']
        if 'bgpSettings' in outs:
            self.bgp_settings = outs['bgpSettings']
        if 'defaultLocalNetworkGatewayId' in outs:
            self.default_local_network_gateway_id = outs['defaultLocalNetworkGatewayId']
        if 'enableBgp' in outs:
            self.enable_bgp = outs['enableBgp']
        if 'ipConfigurations' in outs:
            self.ip_configurations = outs['ipConfigurations']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'sku' in outs:
            self.sku = outs['sku']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'type' in outs:
            self.type = outs['type']
        if 'vpnClientConfiguration' in outs:
            self.vpn_client_configuration = outs['vpnClientConfiguration']
        if 'vpnType' in outs:
            self.vpn_type = outs['vpnType']
