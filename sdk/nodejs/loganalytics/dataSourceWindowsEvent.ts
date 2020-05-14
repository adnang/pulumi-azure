// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Log Analytics Windows Event DataSource.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAnalyticsWorkspace = new azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "PerGB2018",
 * });
 * const exampleDataSourceWindowsEvent = new azure.loganalytics.DataSourceWindowsEvent("exampleDataSourceWindowsEvent", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     workspaceName: exampleAnalyticsWorkspace.name,
 *     eventLogName: "Application",
 *     eventTypes: ["error"],
 * });
 * ```
 */
export class DataSourceWindowsEvent extends pulumi.CustomResource {
    /**
     * Get an existing DataSourceWindowsEvent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataSourceWindowsEventState, opts?: pulumi.CustomResourceOptions): DataSourceWindowsEvent {
        return new DataSourceWindowsEvent(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:loganalytics/dataSourceWindowsEvent:DataSourceWindowsEvent';

    /**
     * Returns true if the given object is an instance of DataSourceWindowsEvent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSourceWindowsEvent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSourceWindowsEvent.__pulumiType;
    }

    /**
     * Specifies the name of the Windows Event Log to collect events from.
     */
    public readonly eventLogName!: pulumi.Output<string>;
    /**
     * Specifies an array of event types applied to the specified event log. Possible values include `error`, `warning` and `information`.
     */
    public readonly eventTypes!: pulumi.Output<string[]>;
    /**
     * The name which should be used for this Log Analytics Windows Event DataSource. Changing this forces a new Log Analytics Windows Event DataSource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
     */
    public readonly workspaceName!: pulumi.Output<string>;

    /**
     * Create a DataSourceWindowsEvent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataSourceWindowsEventArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataSourceWindowsEventArgs | DataSourceWindowsEventState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DataSourceWindowsEventState | undefined;
            inputs["eventLogName"] = state ? state.eventLogName : undefined;
            inputs["eventTypes"] = state ? state.eventTypes : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["workspaceName"] = state ? state.workspaceName : undefined;
        } else {
            const args = argsOrState as DataSourceWindowsEventArgs | undefined;
            if (!args || args.eventLogName === undefined) {
                throw new Error("Missing required property 'eventLogName'");
            }
            if (!args || args.eventTypes === undefined) {
                throw new Error("Missing required property 'eventTypes'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["eventLogName"] = args ? args.eventLogName : undefined;
            inputs["eventTypes"] = args ? args.eventTypes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DataSourceWindowsEvent.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataSourceWindowsEvent resources.
 */
export interface DataSourceWindowsEventState {
    /**
     * Specifies the name of the Windows Event Log to collect events from.
     */
    readonly eventLogName?: pulumi.Input<string>;
    /**
     * Specifies an array of event types applied to the specified event log. Possible values include `error`, `warning` and `information`.
     */
    readonly eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name which should be used for this Log Analytics Windows Event DataSource. Changing this forces a new Log Analytics Windows Event DataSource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
     */
    readonly workspaceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataSourceWindowsEvent resource.
 */
export interface DataSourceWindowsEventArgs {
    /**
     * Specifies the name of the Windows Event Log to collect events from.
     */
    readonly eventLogName: pulumi.Input<string>;
    /**
     * Specifies an array of event types applied to the specified event log. Possible values include `error`, `warning` and `information`.
     */
    readonly eventTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name which should be used for this Log Analytics Windows Event DataSource. Changing this forces a new Log Analytics Windows Event DataSource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
     */
    readonly workspaceName: pulumi.Input<string>;
}
