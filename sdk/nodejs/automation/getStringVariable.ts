// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Automation String Variable.
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const example = azure.automation.getStringVariable({
 *     automationAccountName: "tfex-example-account",
 *     name: "tfex-example-var",
 *     resourceGroupName: "tfex-example-rg",
 * });
 * 
 * export const variableId = example.id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/automation_variable_string.html.markdown.
 */
export function getStringVariable(args: GetStringVariableArgs, opts?: pulumi.InvokeOptions): Promise<GetStringVariableResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:automation/getStringVariable:getStringVariable", {
        "automationAccountName": args.automationAccountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getStringVariable.
 */
export interface GetStringVariableArgs {
    /**
     * The name of the automation account in which the Automation Variable exists.
     */
    readonly automationAccountName: string;
    /**
     * The name of the Automation Variable.
     */
    readonly name: string;
    /**
     * The Name of the Resource Group where the automation account exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getStringVariable.
 */
export interface GetStringVariableResult {
    readonly automationAccountName: string;
    /**
     * The description of the Automation Variable.
     */
    readonly description: string;
    /**
     * Specifies if the Automation Variable is encrypted. Defaults to `false`.
     */
    readonly encrypted: boolean;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The value of the Automation Variable as a `string`.
     */
    readonly value: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
