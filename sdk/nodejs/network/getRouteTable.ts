// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Route Table.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const example = azure.network.getRouteTable({
 *     name: "myroutetable",
 *     resourceGroupName: "some-resource-group",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/route_table.html.markdown.
 */
export function getRouteTable(args: GetRouteTableArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteTableResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getRouteTable:getRouteTable", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteTable.
 */
export interface GetRouteTableArgs {
    /**
     * The name of the Route Table.
     */
    readonly name: string;
    /**
     * The name of the Resource Group in which the Route Table exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getRouteTable.
 */
export interface GetRouteTableResult {
    /**
     * The Azure Region in which the Route Table exists.
     */
    readonly location: string;
    /**
     * The name of the Route.
     */
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * One or more `route` blocks as documented below.
     */
    readonly routes: outputs.network.GetRouteTableRoute[];
    /**
     * The collection of Subnets associated with this route table.
     */
    readonly subnets: string[];
    /**
     * A mapping of tags assigned to the Route Table.
     */
    readonly tags: {[key: string]: string};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
