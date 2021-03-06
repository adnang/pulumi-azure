# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SubnetRouteTableAssociation(pulumi.CustomResource):
    route_table_id: pulumi.Output[str]
    """
    The ID of the Route Table which should be associated with the Subnet. Changing this forces a new resource to be created.
    """
    subnet_id: pulumi.Output[str]
    """
    The ID of the Subnet. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, route_table_id=None, subnet_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Associates a Route Table with a Subnet within a Virtual Network.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefix="10.0.2.0/24")
        example_route_table = azure.network.RouteTable("exampleRouteTable",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            route=[{
                "name": "example",
                "addressPrefix": "10.100.0.0/14",
                "nextHopType": "VirtualAppliance",
                "nextHopInIpAddress": "10.10.1.1",
            }])
        example_subnet_route_table_association = azure.network.SubnetRouteTableAssociation("exampleSubnetRouteTableAssociation",
            subnet_id=example_subnet.id,
            route_table_id=example_route_table.id)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] route_table_id: The ID of the Route Table which should be associated with the Subnet. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet. Changing this forces a new resource to be created.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if route_table_id is None:
                raise TypeError("Missing required property 'route_table_id'")
            __props__['route_table_id'] = route_table_id
            if subnet_id is None:
                raise TypeError("Missing required property 'subnet_id'")
            __props__['subnet_id'] = subnet_id
        super(SubnetRouteTableAssociation, __self__).__init__(
            'azure:network/subnetRouteTableAssociation:SubnetRouteTableAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, route_table_id=None, subnet_id=None):
        """
        Get an existing SubnetRouteTableAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] route_table_id: The ID of the Route Table which should be associated with the Subnet. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["route_table_id"] = route_table_id
        __props__["subnet_id"] = subnet_id
        return SubnetRouteTableAssociation(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

