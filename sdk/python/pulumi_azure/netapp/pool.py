# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Pool(pulumi.CustomResource):
    account_name: pulumi.Output[str]
    """
    The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the NetApp Pool. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
    """
    service_level: pulumi.Output[str]
    """
    The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
    """
    size_in_tb: pulumi.Output[float]
    """
    Provisioned size of the pool in TB. Value must be between `4` and `500`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, location=None, name=None, resource_group_name=None, service_level=None, size_in_tb=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Pool within a NetApp Account.

        ## NetApp Pool Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.netapp.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_pool = azure.netapp.Pool("examplePool",
            account_name=example_account.name,
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            service_level="Premium",
            size_in_tb=4)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Pool. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_level: The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        :param pulumi.Input[float] size_in_tb: Provisioned size of the pool in TB. Value must be between `4` and `500`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_level is None:
                raise TypeError("Missing required property 'service_level'")
            __props__['service_level'] = service_level
            if size_in_tb is None:
                raise TypeError("Missing required property 'size_in_tb'")
            __props__['size_in_tb'] = size_in_tb
            __props__['tags'] = tags
        super(Pool, __self__).__init__(
            'azure:netapp/pool:Pool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_name=None, location=None, name=None, resource_group_name=None, service_level=None, size_in_tb=None, tags=None):
        """
        Get an existing Pool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Pool. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_level: The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        :param pulumi.Input[float] size_in_tb: Provisioned size of the pool in TB. Value must be between `4` and `500`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["service_level"] = service_level
        __props__["size_in_tb"] = size_in_tb
        __props__["tags"] = tags
        return Pool(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

