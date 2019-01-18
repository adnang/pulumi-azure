# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Account(pulumi.CustomResource):
    dsc_primary_access_key: pulumi.Output[str]
    """
    The Primary Access Key for the DSC Endpoint associated with this Automation Account.
    """
    dsc_secondary_access_key: pulumi.Output[str]
    """
    The Secondary Access Key for the DSC Endpoint associated with this Automation Account.
    """
    dsc_server_endpoint: pulumi.Output[str]
    """
    The DSC Server Endpoint associated with this Automation Account.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The SKU name of the account - only `Basic` is supported at this time. Defaults to `Basic`.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the Automation Account is created. Changing this forces a new resource to be created.
    """
    sku: pulumi.Output[dict]
    """
    A `sku` block as defined below.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, __name__, __opts__=None, location=None, name=None, resource_group_name=None, sku=None, tags=None):
        """
        Manages a Automation Account.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The SKU name of the account - only `Basic` is supported at this time. Defaults to `Basic`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Automation Account is created. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] sku: A `sku` block as defined below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resource_group_name'] = resource_group_name

        if not sku:
            raise TypeError('Missing required property sku')
        __props__['sku'] = sku

        __props__['tags'] = tags

        __props__['dsc_primary_access_key'] = None
        __props__['dsc_secondary_access_key'] = None
        __props__['dsc_server_endpoint'] = None

        super(Account, __self__).__init__(
            'azure:automation/account:Account',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

