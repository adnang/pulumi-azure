# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class VirtualMachine(pulumi.CustomResource):
    auto_patching: pulumi.Output[dict]
    """
    An `auto_patching` block as defined below.

      * `dayOfWeek` (`str`) - The day of week to apply the patch on.
      * `maintenanceWindowDurationInMinutes` (`float`) - The size of the Maintenance Window in minutes.
      * `maintenanceWindowStartingHour` (`float`) - The Hour, in the Virtual Machine Time-Zone when the patching maintenance window should begin.
    """
    key_vault_credential: pulumi.Output[dict]
    """
    (Optional) An `key_vault_credential` block as defined below.

      * `keyVaultUrl` (`str`) - The azure Key Vault url. Changing this forces a new resource to be created.
      * `name` (`str`) - The credential name.
      * `servicePrincipalName` (`str`) - The service principal name to access key vault. Changing this forces a new resource to be created.
      * `servicePrincipalSecret` (`str`) - The service principal name secret to access key vault. Changing this forces a new resource to be created.
    """
    r_services_enabled: pulumi.Output[bool]
    """
    Should R Services be enabled?
    """
    sql_connectivity_port: pulumi.Output[float]
    """
    The SQL Server port. Defaults to `1433`.
    """
    sql_connectivity_type: pulumi.Output[str]
    """
    The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
    """
    sql_connectivity_update_password: pulumi.Output[str]
    """
    The SQL Server sysadmin login password.
    """
    sql_connectivity_update_username: pulumi.Output[str]
    """
    The SQL Server sysadmin login to create.
    """
    sql_license_type: pulumi.Output[str]
    """
    The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    virtual_machine_id: pulumi.Output[str]
    """
    The ID of the Virtual Machine. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, auto_patching=None, key_vault_credential=None, r_services_enabled=None, sql_connectivity_port=None, sql_connectivity_type=None, sql_connectivity_update_password=None, sql_connectivity_update_username=None, sql_license_type=None, tags=None, virtual_machine_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Microsoft SQL Virtual Machine



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auto_patching: An `auto_patching` block as defined below.
        :param pulumi.Input[dict] key_vault_credential: (Optional) An `key_vault_credential` block as defined below.
        :param pulumi.Input[bool] r_services_enabled: Should R Services be enabled?
        :param pulumi.Input[float] sql_connectivity_port: The SQL Server port. Defaults to `1433`.
        :param pulumi.Input[str] sql_connectivity_type: The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        :param pulumi.Input[str] sql_connectivity_update_password: The SQL Server sysadmin login password.
        :param pulumi.Input[str] sql_connectivity_update_username: The SQL Server sysadmin login to create.
        :param pulumi.Input[str] sql_license_type: The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] virtual_machine_id: The ID of the Virtual Machine. Changing this forces a new resource to be created.

        The **auto_patching** object supports the following:

          * `dayOfWeek` (`pulumi.Input[str]`) - The day of week to apply the patch on.
          * `maintenanceWindowDurationInMinutes` (`pulumi.Input[float]`) - The size of the Maintenance Window in minutes.
          * `maintenanceWindowStartingHour` (`pulumi.Input[float]`) - The Hour, in the Virtual Machine Time-Zone when the patching maintenance window should begin.

        The **key_vault_credential** object supports the following:

          * `keyVaultUrl` (`pulumi.Input[str]`) - The azure Key Vault url. Changing this forces a new resource to be created.
          * `name` (`pulumi.Input[str]`) - The credential name.
          * `servicePrincipalName` (`pulumi.Input[str]`) - The service principal name to access key vault. Changing this forces a new resource to be created.
          * `servicePrincipalSecret` (`pulumi.Input[str]`) - The service principal name secret to access key vault. Changing this forces a new resource to be created.
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

            __props__['auto_patching'] = auto_patching
            __props__['key_vault_credential'] = key_vault_credential
            __props__['r_services_enabled'] = r_services_enabled
            __props__['sql_connectivity_port'] = sql_connectivity_port
            __props__['sql_connectivity_type'] = sql_connectivity_type
            __props__['sql_connectivity_update_password'] = sql_connectivity_update_password
            __props__['sql_connectivity_update_username'] = sql_connectivity_update_username
            if sql_license_type is None:
                raise TypeError("Missing required property 'sql_license_type'")
            __props__['sql_license_type'] = sql_license_type
            __props__['tags'] = tags
            if virtual_machine_id is None:
                raise TypeError("Missing required property 'virtual_machine_id'")
            __props__['virtual_machine_id'] = virtual_machine_id
        super(VirtualMachine, __self__).__init__(
            'azure:mssql/virtualMachine:VirtualMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auto_patching=None, key_vault_credential=None, r_services_enabled=None, sql_connectivity_port=None, sql_connectivity_type=None, sql_connectivity_update_password=None, sql_connectivity_update_username=None, sql_license_type=None, tags=None, virtual_machine_id=None):
        """
        Get an existing VirtualMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auto_patching: An `auto_patching` block as defined below.
        :param pulumi.Input[dict] key_vault_credential: (Optional) An `key_vault_credential` block as defined below.
        :param pulumi.Input[bool] r_services_enabled: Should R Services be enabled?
        :param pulumi.Input[float] sql_connectivity_port: The SQL Server port. Defaults to `1433`.
        :param pulumi.Input[str] sql_connectivity_type: The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        :param pulumi.Input[str] sql_connectivity_update_password: The SQL Server sysadmin login password.
        :param pulumi.Input[str] sql_connectivity_update_username: The SQL Server sysadmin login to create.
        :param pulumi.Input[str] sql_license_type: The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] virtual_machine_id: The ID of the Virtual Machine. Changing this forces a new resource to be created.

        The **auto_patching** object supports the following:

          * `dayOfWeek` (`pulumi.Input[str]`) - The day of week to apply the patch on.
          * `maintenanceWindowDurationInMinutes` (`pulumi.Input[float]`) - The size of the Maintenance Window in minutes.
          * `maintenanceWindowStartingHour` (`pulumi.Input[float]`) - The Hour, in the Virtual Machine Time-Zone when the patching maintenance window should begin.

        The **key_vault_credential** object supports the following:

          * `keyVaultUrl` (`pulumi.Input[str]`) - The azure Key Vault url. Changing this forces a new resource to be created.
          * `name` (`pulumi.Input[str]`) - The credential name.
          * `servicePrincipalName` (`pulumi.Input[str]`) - The service principal name to access key vault. Changing this forces a new resource to be created.
          * `servicePrincipalSecret` (`pulumi.Input[str]`) - The service principal name secret to access key vault. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_patching"] = auto_patching
        __props__["key_vault_credential"] = key_vault_credential
        __props__["r_services_enabled"] = r_services_enabled
        __props__["sql_connectivity_port"] = sql_connectivity_port
        __props__["sql_connectivity_type"] = sql_connectivity_type
        __props__["sql_connectivity_update_password"] = sql_connectivity_update_password
        __props__["sql_connectivity_update_username"] = sql_connectivity_update_username
        __props__["sql_license_type"] = sql_license_type
        __props__["tags"] = tags
        __props__["virtual_machine_id"] = virtual_machine_id
        return VirtualMachine(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
