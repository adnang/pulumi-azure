// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package keyvault

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Key Vault.
// 
// > **NOTE:** It's possible to define Key Vault Access Policies both within the `keyvault.KeyVault` resource via the `accessPolicy` block and by using the `keyvault.AccessPolicy` resource. However it's not possible to use both methods to manage Access Policies within a KeyVault, since there'll be conflicts.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/key_vault.html.markdown.
type KeyVault struct {
	pulumi.CustomResourceState

	// [A list](https://www.terraform.io/docs/configuration/attr-as-blocks.html) of up to 16 objects describing access policies, as described below.
	AccessPolicies KeyVaultAccessPolicyArrayOutput `pulumi:"accessPolicies"`
	// Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
	EnabledForDeployment pulumi.BoolPtrOutput `pulumi:"enabledForDeployment"`
	// Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
	EnabledForDiskEncryption pulumi.BoolPtrOutput `pulumi:"enabledForDiskEncryption"`
	// Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
	EnabledForTemplateDeployment pulumi.BoolPtrOutput `pulumi:"enabledForTemplateDeployment"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Key Vault. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkAcls` block as defined below.
	NetworkAcls KeyVaultNetworkAclsOutput `pulumi:"networkAcls"`
	// The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// ) A `sku` block as described below.
	Sku KeyVaultSkuOutput `pulumi:"sku"`
	// The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The URI of the Key Vault, used for performing operations on keys and secrets.
	VaultUri pulumi.StringOutput `pulumi:"vaultUri"`
}

// NewKeyVault registers a new resource with the given unique name, arguments, and options.
func NewKeyVault(ctx *pulumi.Context,
	name string, args *KeyVaultArgs, opts ...pulumi.ResourceOption) (*KeyVault, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TenantId == nil {
		return nil, errors.New("missing required argument 'TenantId'")
	}
	if args == nil {
		args = &KeyVaultArgs{}
	}
	var resource KeyVault
	err := ctx.RegisterResource("azure:keyvault/keyVault:KeyVault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyVault gets an existing KeyVault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyVaultState, opts ...pulumi.ResourceOption) (*KeyVault, error) {
	var resource KeyVault
	err := ctx.ReadResource("azure:keyvault/keyVault:KeyVault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyVault resources.
type keyVaultState struct {
	// [A list](https://www.terraform.io/docs/configuration/attr-as-blocks.html) of up to 16 objects describing access policies, as described below.
	AccessPolicies []KeyVaultAccessPolicy `pulumi:"accessPolicies"`
	// Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
	EnabledForDeployment *bool `pulumi:"enabledForDeployment"`
	// Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
	EnabledForDiskEncryption *bool `pulumi:"enabledForDiskEncryption"`
	// Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
	EnabledForTemplateDeployment *bool `pulumi:"enabledForTemplateDeployment"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Key Vault. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkAcls` block as defined below.
	NetworkAcls *KeyVaultNetworkAcls `pulumi:"networkAcls"`
	// The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// ) A `sku` block as described below.
	Sku *KeyVaultSku `pulumi:"sku"`
	// The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId *string `pulumi:"tenantId"`
	// The URI of the Key Vault, used for performing operations on keys and secrets.
	VaultUri *string `pulumi:"vaultUri"`
}

type KeyVaultState struct {
	// [A list](https://www.terraform.io/docs/configuration/attr-as-blocks.html) of up to 16 objects describing access policies, as described below.
	AccessPolicies KeyVaultAccessPolicyArrayInput
	// Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
	EnabledForDeployment pulumi.BoolPtrInput
	// Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
	EnabledForDiskEncryption pulumi.BoolPtrInput
	// Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
	EnabledForTemplateDeployment pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Key Vault. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkAcls` block as defined below.
	NetworkAcls KeyVaultNetworkAclsPtrInput
	// The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// ) A `sku` block as described below.
	Sku KeyVaultSkuPtrInput
	// The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
	SkuName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId pulumi.StringPtrInput
	// The URI of the Key Vault, used for performing operations on keys and secrets.
	VaultUri pulumi.StringPtrInput
}

func (KeyVaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyVaultState)(nil)).Elem()
}

type keyVaultArgs struct {
	// [A list](https://www.terraform.io/docs/configuration/attr-as-blocks.html) of up to 16 objects describing access policies, as described below.
	AccessPolicies []KeyVaultAccessPolicy `pulumi:"accessPolicies"`
	// Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
	EnabledForDeployment *bool `pulumi:"enabledForDeployment"`
	// Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
	EnabledForDiskEncryption *bool `pulumi:"enabledForDiskEncryption"`
	// Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
	EnabledForTemplateDeployment *bool `pulumi:"enabledForTemplateDeployment"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Key Vault. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkAcls` block as defined below.
	NetworkAcls *KeyVaultNetworkAcls `pulumi:"networkAcls"`
	// The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// ) A `sku` block as described below.
	Sku *KeyVaultSku `pulumi:"sku"`
	// The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a KeyVault resource.
type KeyVaultArgs struct {
	// [A list](https://www.terraform.io/docs/configuration/attr-as-blocks.html) of up to 16 objects describing access policies, as described below.
	AccessPolicies KeyVaultAccessPolicyArrayInput
	// Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
	EnabledForDeployment pulumi.BoolPtrInput
	// Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
	EnabledForDiskEncryption pulumi.BoolPtrInput
	// Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
	EnabledForTemplateDeployment pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Key Vault. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkAcls` block as defined below.
	NetworkAcls KeyVaultNetworkAclsPtrInput
	// The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// ) A `sku` block as described below.
	Sku KeyVaultSkuPtrInput
	// The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
	SkuName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId pulumi.StringInput
}

func (KeyVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyVaultArgs)(nil)).Elem()
}

