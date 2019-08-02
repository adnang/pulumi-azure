# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Domain(pulumi.CustomResource):
    endpoint: pulumi.Output[str]
    """
    The Endpoint associated with the EventGrid Domain.
    """
    input_mapping_default_values: pulumi.Output[dict]
    """
    A `input_mapping_default_values` block as defined below.
    """
    input_mapping_fields: pulumi.Output[dict]
    """
    A `input_mapping_fields` block as defined below.
    """
    input_schema: pulumi.Output[str]
    """
    Specifies the schema in which incoming events will be published to this domain. Allowed values are `cloudeventv01schema`, `customeventschema`, or `eventgridschema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, input_mapping_default_values=None, input_mapping_fields=None, input_schema=None, location=None, name=None, resource_group_name=None, tags=None, __name__=None, __opts__=None):
        """
        Manages an EventGrid Domain
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input[dict] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `cloudeventv01schema`, `customeventschema`, or `eventgridschema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventgrid_domain_legacy.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['input_mapping_default_values'] = input_mapping_default_values

        __props__['input_mapping_fields'] = input_mapping_fields

        __props__['input_schema'] = input_schema

        __props__['location'] = location

        __props__['name'] = name

        if resource_group_name is None:
            raise TypeError("Missing required property 'resource_group_name'")
        __props__['resource_group_name'] = resource_group_name

        __props__['tags'] = tags

        __props__['endpoint'] = None

        super(Domain, __self__).__init__(
            'azure:eventhub/domain:Domain',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

