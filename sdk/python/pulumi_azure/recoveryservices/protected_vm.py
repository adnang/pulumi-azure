# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ProtectedVM(pulumi.CustomResource):
    backup_policy_id: pulumi.Output[str]
    """
    Specifies the id of the backup policy to use. Changing this forces a new resource to be created.
    """
    recovery_vault_name: pulumi.Output[str]
    """
    Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
    """
    source_vm_id: pulumi.Output[str]
    """
    Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, __name__, __opts__=None, backup_policy_id=None, recovery_vault_name=None, resource_group_name=None, source_vm_id=None, tags=None):
        """
        Manages an Recovery Protected VM.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] backup_policy_id: Specifies the id of the backup policy to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_vm_id: Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not backup_policy_id:
            raise TypeError('Missing required property backup_policy_id')
        __props__['backup_policy_id'] = backup_policy_id

        if not recovery_vault_name:
            raise TypeError('Missing required property recovery_vault_name')
        __props__['recovery_vault_name'] = recovery_vault_name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resource_group_name'] = resource_group_name

        if not source_vm_id:
            raise TypeError('Missing required property source_vm_id')
        __props__['source_vm_id'] = source_vm_id

        __props__['tags'] = tags

        super(ProtectedVM, __self__).__init__(
            'azure:recoveryservices/protectedVM:ProtectedVM',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

