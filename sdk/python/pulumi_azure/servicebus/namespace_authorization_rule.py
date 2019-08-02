# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NamespaceAuthorizationRule(pulumi.CustomResource):
    listen: pulumi.Output[bool]
    """
    Grants listen access to this this Authorization Rule. Defaults to `false`.
    """
    manage: pulumi.Output[bool]
    """
    Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
    """
    namespace_name: pulumi.Output[str]
    """
    Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
    """
    primary_connection_string: pulumi.Output[str]
    """
    The Primary Connection String for the ServiceBus Namespace authorization Rule.
    """
    primary_key: pulumi.Output[str]
    """
    The Primary Key for the ServiceBus Namespace authorization Rule.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
    """
    secondary_connection_string: pulumi.Output[str]
    """
    The Secondary Connection String for the ServiceBus Namespace authorization Rule.
    """
    secondary_key: pulumi.Output[str]
    """
    The Secondary Key for the ServiceBus Namespace authorization Rule.
    """
    send: pulumi.Output[bool]
    """
    Grants send access to this this Authorization Rule. Defaults to `false`.
    """
    def __init__(__self__, resource_name, opts=None, listen=None, manage=None, name=None, namespace_name=None, resource_group_name=None, send=None, __name__=None, __opts__=None):
        """
        Manages a ServiceBus Namespace authorization Rule within a ServiceBus.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] listen: Grants listen access to this this Authorization Rule. Defaults to `false`.
        :param pulumi.Input[bool] manage: Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] send: Grants send access to this this Authorization Rule. Defaults to `false`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/servicebus_namespace_authorization_rule.html.markdown.
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

        __props__['listen'] = listen

        __props__['manage'] = manage

        __props__['name'] = name

        if namespace_name is None:
            raise TypeError("Missing required property 'namespace_name'")
        __props__['namespace_name'] = namespace_name

        if resource_group_name is None:
            raise TypeError("Missing required property 'resource_group_name'")
        __props__['resource_group_name'] = resource_group_name

        __props__['send'] = send

        __props__['primary_connection_string'] = None
        __props__['primary_key'] = None
        __props__['secondary_connection_string'] = None
        __props__['secondary_key'] = None

        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:eventhub/namespaceAuthorizationRule:NamespaceAuthorizationRule")])
        opts = alias_opts if opts is None else opts.merge(alias_opts)
        super(NamespaceAuthorizationRule, __self__).__init__(
            'azure:servicebus/namespaceAuthorizationRule:NamespaceAuthorizationRule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

