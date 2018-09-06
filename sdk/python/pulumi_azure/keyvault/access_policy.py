# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class AccessPolicy(pulumi.CustomResource):
    """
    Manages a Key Vault Access Policy.
    
    ~> **NOTE:** It's possible to define Key Vault Access Policies both within the `azurerm_key_vault` resource via the `access_policy` block and by using the `azurerm_key_vault_access_policy` resource. However it's not possible to use both methods to manage Access Policies within a KeyVault, since there'll be conflicts.
    
    -> **NOTE:** Azure permits a maximum of 16 Access Policies per Key Vault - [more information can be found in this document](https://docs.microsoft.com/en-us/azure/key-vault/key-vault-secure-your-key-vault#data-plane-access-control).
    """
    def __init__(__self__, __name__, __opts__=None, application_id=None, certificate_permissions=None, key_permissions=None, object_id=None, resource_group_name=None, secret_permissions=None, tenant_id=None, vault_name=None):
        """Create a AccessPolicy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if application_id and not isinstance(application_id, basestring):
            raise TypeError('Expected property application_id to be a basestring')
        __self__.application_id = application_id
        """
        The object ID of an Application in Azure Active Directory.
        """
        __props__['applicationId'] = application_id

        if certificate_permissions and not isinstance(certificate_permissions, list):
            raise TypeError('Expected property certificate_permissions to be a list')
        __self__.certificate_permissions = certificate_permissions
        """
        List of certificate permissions, must be one or more from
        the following: `create`, `delete`, `deleteissuers`, `get`, `getissuers`, `import`, `list`, `listissuers`,
        `managecontacts`, `manageissuers`, `purge`, `recover`, `setissuers` and `update`.
        """
        __props__['certificatePermissions'] = certificate_permissions

        if not key_permissions:
            raise TypeError('Missing required property key_permissions')
        elif not isinstance(key_permissions, list):
            raise TypeError('Expected property key_permissions to be a list')
        __self__.key_permissions = key_permissions
        """
        List of key permissions, must be one or more from
        the following: `backup`, `create`, `decrypt`, `delete`, `encrypt`, `get`, `import`, `list`, `purge`,
        `recover`, `restore`, `sign`, `unwrapKey`, `update`, `verify` and `wrapKey`.
        """
        __props__['keyPermissions'] = key_permissions

        if not object_id:
            raise TypeError('Missing required property object_id')
        elif not isinstance(object_id, basestring):
            raise TypeError('Expected property object_id to be a basestring')
        __self__.object_id = object_id
        """
        The object ID of a user, service principal or security
        group in the Azure Active Directory tenant for the vault. The object ID must
        be unique for the list of access policies. Changing this forces a new resource
        to be created.
        """
        __props__['objectId'] = object_id

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to
        create the namespace. Changing this forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if not secret_permissions:
            raise TypeError('Missing required property secret_permissions')
        elif not isinstance(secret_permissions, list):
            raise TypeError('Expected property secret_permissions to be a list')
        __self__.secret_permissions = secret_permissions
        """
        List of secret permissions, must be one or more
        from the following: `backup`, `delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
        """
        __props__['secretPermissions'] = secret_permissions

        if not tenant_id:
            raise TypeError('Missing required property tenant_id')
        elif not isinstance(tenant_id, basestring):
            raise TypeError('Expected property tenant_id to be a basestring')
        __self__.tenant_id = tenant_id
        """
        The Azure Active Directory tenant ID that should be used
        for authenticating requests to the key vault. Changing this forces a new resource
        to be created.
        """
        __props__['tenantId'] = tenant_id

        if not vault_name:
            raise TypeError('Missing required property vault_name')
        elif not isinstance(vault_name, basestring):
            raise TypeError('Expected property vault_name to be a basestring')
        __self__.vault_name = vault_name
        """
        Specifies the name of the Key Vault resource. Changing this
        forces a new resource to be created.
        """
        __props__['vaultName'] = vault_name

        super(AccessPolicy, __self__).__init__(
            'azure:keyvault/accessPolicy:AccessPolicy',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'applicationId' in outs:
            self.application_id = outs['applicationId']
        if 'certificatePermissions' in outs:
            self.certificate_permissions = outs['certificatePermissions']
        if 'keyPermissions' in outs:
            self.key_permissions = outs['keyPermissions']
        if 'objectId' in outs:
            self.object_id = outs['objectId']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'secretPermissions' in outs:
            self.secret_permissions = outs['secretPermissions']
        if 'tenantId' in outs:
            self.tenant_id = outs['tenantId']
        if 'vaultName' in outs:
            self.vault_name = outs['vaultName']
