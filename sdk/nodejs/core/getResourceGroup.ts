// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Resource Group.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleResourceGroup = azure.core.getResourceGroup({
 *     name: "dsrgTest",
 * });
 * const exampleManagedDisk = new azure.compute.ManagedDisk("example", {
 *     createOption: "Empty",
 *     diskSizeGb: 1,
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     storageAccountType: "Standard_LRS",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/resource_group.html.markdown.
 */
export function getResourceGroup(args: GetResourceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:core/getResourceGroup:getResourceGroup", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourceGroup.
 */
export interface GetResourceGroupArgs {
    /**
     * Specifies the name of the resource group.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getResourceGroup.
 */
export interface GetResourceGroupResult {
    /**
     * The location of the resource group.
     */
    readonly location: string;
    readonly name: string;
    /**
     * A mapping of tags assigned to the resource group.
     */
    readonly tags: {[key: string]: string};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
