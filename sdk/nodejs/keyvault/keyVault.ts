// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a Key Vault.
 * 
 * ~> **NOTE:** It's possible to define Key Vault Access Policies both within the `azurerm_key_vault` resource via the `access_policy` block and by using the `azurerm_key_vault_access_policy` resource. However it's not possible to use both methods to manage Access Policies within a KeyVault, since there'll be conflicts.
 */
export class KeyVault extends pulumi.CustomResource {
    /**
     * Get an existing KeyVault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyVaultState): KeyVault {
        return new KeyVault(name, <any>state, { id });
    }

    /**
     * An access policy block as described below. A maximum of 16
     * may be declared.
     */
    public readonly accessPolicies: pulumi.Output<{ applicationId?: string, certificatePermissions?: string[], keyPermissions: string[], objectId: string, secretPermissions: string[], tenantId: string }[]>;
    /**
     * Boolean flag to specify whether Azure Virtual
     * Machines are permitted to retrieve certificates stored as secrets from the key
     * vault. Defaults to false.
     */
    public readonly enabledForDeployment: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag to specify whether Azure
     * Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
     * Defaults to false.
     */
    public readonly enabledForDiskEncryption: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag to specify whether
     * Azure Resource Manager is permitted to retrieve secrets from the key vault.
     * Defaults to false.
     */
    public readonly enabledForTemplateDeployment: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists.
     * Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * SKU name to specify whether the key vault is a `standard`
     * or `premium` vault.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The name of the resource group in which to
     * create the Key Vault. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * An SKU block as described below.
     */
    public readonly sku: pulumi.Output<{ name: string }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;
    /**
     * The Azure Active Directory tenant ID that should be used
     * for authenticating requests to the key vault. Must match the `tenant_id` used
     * above.
     */
    public readonly tenantId: pulumi.Output<string>;
    /**
     * The URI of the vault for performing operations on keys and secrets.
     */
    public /*out*/ readonly vaultUri: pulumi.Output<string>;

    /**
     * Create a KeyVault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyVaultArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyVaultArgs | KeyVaultState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: KeyVaultState = argsOrState as KeyVaultState | undefined;
            inputs["accessPolicies"] = state ? state.accessPolicies : undefined;
            inputs["enabledForDeployment"] = state ? state.enabledForDeployment : undefined;
            inputs["enabledForDiskEncryption"] = state ? state.enabledForDiskEncryption : undefined;
            inputs["enabledForTemplateDeployment"] = state ? state.enabledForTemplateDeployment : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["vaultUri"] = state ? state.vaultUri : undefined;
        } else {
            const args = argsOrState as KeyVaultArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            if (!args || args.tenantId === undefined) {
                throw new Error("Missing required property 'tenantId'");
            }
            inputs["accessPolicies"] = args ? args.accessPolicies : undefined;
            inputs["enabledForDeployment"] = args ? args.enabledForDeployment : undefined;
            inputs["enabledForDiskEncryption"] = args ? args.enabledForDiskEncryption : undefined;
            inputs["enabledForTemplateDeployment"] = args ? args.enabledForTemplateDeployment : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["vaultUri"] = undefined /*out*/;
        }
        super("azure:keyvault/keyVault:KeyVault", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyVault resources.
 */
export interface KeyVaultState {
    /**
     * An access policy block as described below. A maximum of 16
     * may be declared.
     */
    readonly accessPolicies?: pulumi.Input<pulumi.Input<{ applicationId?: pulumi.Input<string>, certificatePermissions?: pulumi.Input<pulumi.Input<string>[]>, keyPermissions: pulumi.Input<pulumi.Input<string>[]>, objectId: pulumi.Input<string>, secretPermissions: pulumi.Input<pulumi.Input<string>[]>, tenantId: pulumi.Input<string> }>[]>;
    /**
     * Boolean flag to specify whether Azure Virtual
     * Machines are permitted to retrieve certificates stored as secrets from the key
     * vault. Defaults to false.
     */
    readonly enabledForDeployment?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether Azure
     * Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
     * Defaults to false.
     */
    readonly enabledForDiskEncryption?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether
     * Azure Resource Manager is permitted to retrieve secrets from the key vault.
     * Defaults to false.
     */
    readonly enabledForTemplateDeployment?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists.
     * Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * SKU name to specify whether the key vault is a `standard`
     * or `premium` vault.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the Key Vault. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * An SKU block as described below.
     */
    readonly sku?: pulumi.Input<{ name: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The Azure Active Directory tenant ID that should be used
     * for authenticating requests to the key vault. Must match the `tenant_id` used
     * above.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * The URI of the vault for performing operations on keys and secrets.
     */
    readonly vaultUri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyVault resource.
 */
export interface KeyVaultArgs {
    /**
     * An access policy block as described below. A maximum of 16
     * may be declared.
     */
    readonly accessPolicies?: pulumi.Input<pulumi.Input<{ applicationId?: pulumi.Input<string>, certificatePermissions?: pulumi.Input<pulumi.Input<string>[]>, keyPermissions: pulumi.Input<pulumi.Input<string>[]>, objectId: pulumi.Input<string>, secretPermissions: pulumi.Input<pulumi.Input<string>[]>, tenantId: pulumi.Input<string> }>[]>;
    /**
     * Boolean flag to specify whether Azure Virtual
     * Machines are permitted to retrieve certificates stored as secrets from the key
     * vault. Defaults to false.
     */
    readonly enabledForDeployment?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether Azure
     * Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
     * Defaults to false.
     */
    readonly enabledForDiskEncryption?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether
     * Azure Resource Manager is permitted to retrieve secrets from the key vault.
     * Defaults to false.
     */
    readonly enabledForTemplateDeployment?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists.
     * Changing this forces a new resource to be created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * SKU name to specify whether the key vault is a `standard`
     * or `premium` vault.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the Key Vault. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * An SKU block as described below.
     */
    readonly sku: pulumi.Input<{ name: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The Azure Active Directory tenant ID that should be used
     * for authenticating requests to the key vault. Must match the `tenant_id` used
     * above.
     */
    readonly tenantId: pulumi.Input<string>;
}
