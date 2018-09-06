# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Subnet(pulumi.CustomResource):
    """
    Manages a subnet. Subnets represent network segments within the IP space defined by the virtual network.
    
    ~> **NOTE on Virtual Networks and Subnet's:** Terraform currently
    provides both a standalone Subnet resource, and allows for Subnets to be defined in-line within the Virtual Network resource.
    At this time you cannot use a Virtual Network with in-line Subnets in conjunction with any Subnet resources. Doing so will cause a conflict of Subnet configurations and will overwrite Subnet's.
    """
    def __init__(__self__, __name__, __opts__=None, address_prefix=None, ip_configurations=None, name=None, network_security_group_id=None, resource_group_name=None, route_table_id=None, service_endpoints=None, virtual_network_name=None):
        """Create a Subnet resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not address_prefix:
            raise TypeError('Missing required property address_prefix')
        elif not isinstance(address_prefix, basestring):
            raise TypeError('Expected property address_prefix to be a basestring')
        __self__.address_prefix = address_prefix
        """
        The address prefix to use for the subnet.
        """
        __props__['addressPrefix'] = address_prefix

        if ip_configurations and not isinstance(ip_configurations, list):
            raise TypeError('Expected property ip_configurations to be a list')
        __self__.ip_configurations = ip_configurations
        """
        The collection of IP Configurations with IPs within this subnet.
        """
        __props__['ipConfigurations'] = ip_configurations

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the subnet. Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if network_security_group_id and not isinstance(network_security_group_id, basestring):
            raise TypeError('Expected property network_security_group_id to be a basestring')
        __self__.network_security_group_id = network_security_group_id
        """
        The ID of the Network Security Group to associate with the subnet.
        """
        __props__['networkSecurityGroupId'] = network_security_group_id

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if route_table_id and not isinstance(route_table_id, basestring):
            raise TypeError('Expected property route_table_id to be a basestring')
        __self__.route_table_id = route_table_id
        """
        The ID of the Route Table to associate with the subnet.
        """
        __props__['routeTableId'] = route_table_id

        if service_endpoints and not isinstance(service_endpoints, list):
            raise TypeError('Expected property service_endpoints to be a list')
        __self__.service_endpoints = service_endpoints
        """
        The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.Storage`, `Microsoft.Sql`.
        """
        __props__['serviceEndpoints'] = service_endpoints

        if not virtual_network_name:
            raise TypeError('Missing required property virtual_network_name')
        elif not isinstance(virtual_network_name, basestring):
            raise TypeError('Expected property virtual_network_name to be a basestring')
        __self__.virtual_network_name = virtual_network_name
        """
        The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
        """
        __props__['virtualNetworkName'] = virtual_network_name

        super(Subnet, __self__).__init__(
            'azure:network/subnet:Subnet',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'addressPrefix' in outs:
            self.address_prefix = outs['addressPrefix']
        if 'ipConfigurations' in outs:
            self.ip_configurations = outs['ipConfigurations']
        if 'name' in outs:
            self.name = outs['name']
        if 'networkSecurityGroupId' in outs:
            self.network_security_group_id = outs['networkSecurityGroupId']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'routeTableId' in outs:
            self.route_table_id = outs['routeTableId']
        if 'serviceEndpoints' in outs:
            self.service_endpoints = outs['serviceEndpoints']
        if 'virtualNetworkName' in outs:
            self.virtual_network_name = outs['virtualNetworkName']
