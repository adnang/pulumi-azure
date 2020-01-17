// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Recovery Services VM Protection Policy.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const policy = azure.recoveryservices.getVMProtectionPolicy({
 *     name: "policy",
 *     recoveryVaultName: "recoveryVault",
 *     resourceGroupName: "resourceGroup",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/recovery_services_protection_policy_vm.html.markdown.
 */
export function getVMProtectionPolicy(args: GetVMProtectionPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetVMProtectionPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:recoveryservices/getVMProtectionPolicy:getVMProtectionPolicy", {
        "name": args.name,
        "recoveryVaultName": args.recoveryVaultName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getVMProtectionPolicy.
 */
export interface GetVMProtectionPolicyArgs {
    /**
     * Specifies the name of the Recovery Services VM Protection Policy.
     */
    readonly name: string;
    /**
     * Specifies the name of the Recovery Services Vault.
     */
    readonly recoveryVaultName: string;
    /**
     * The name of the resource group in which the Recovery Services VM Protection Policy resides.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getVMProtectionPolicy.
 */
export interface GetVMProtectionPolicyResult {
    readonly name: string;
    readonly recoveryVaultName: string;
    readonly resourceGroupName: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
