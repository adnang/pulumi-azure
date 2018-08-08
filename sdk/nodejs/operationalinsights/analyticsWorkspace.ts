// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a Log Analytics (formally Operational Insights) Workspace.
 */
export class AnalyticsWorkspace extends pulumi.CustomResource {
    /**
     * Get an existing AnalyticsWorkspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnalyticsWorkspaceState): AnalyticsWorkspace {
        return new AnalyticsWorkspace(name, <any>state, { id });
    }

    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The Portal URL for the Log Analytics Workspace.
     */
    public /*out*/ readonly portalUrl: pulumi.Output<string>;
    /**
     * The Primary shared key for the Log Analytics Workspace.
     */
    public /*out*/ readonly primarySharedKey: pulumi.Output<string>;
    /**
     * The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * The workspace data retention in days. Possible values range between 30 and 730.
     */
    public readonly retentionInDays: pulumi.Output<number>;
    /**
     * The Secondary shared key for the Log Analytics Workspace.
     */
    public /*out*/ readonly secondarySharedKey: pulumi.Output<string>;
    /**
     * Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`).
     */
    public readonly sku: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;
    /**
     * The Workspace (or Customer) ID for the Log Analytics Workspace.
     */
    public /*out*/ readonly workspaceId: pulumi.Output<string>;

    /**
     * Create a AnalyticsWorkspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnalyticsWorkspaceArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: AnalyticsWorkspaceArgs | AnalyticsWorkspaceState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AnalyticsWorkspaceState = argsOrState as AnalyticsWorkspaceState | undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["portalUrl"] = state ? state.portalUrl : undefined;
            inputs["primarySharedKey"] = state ? state.primarySharedKey : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["retentionInDays"] = state ? state.retentionInDays : undefined;
            inputs["secondarySharedKey"] = state ? state.secondarySharedKey : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["workspaceId"] = state ? state.workspaceId : undefined;
        } else {
            const args = argsOrState as AnalyticsWorkspaceArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["retentionInDays"] = args ? args.retentionInDays : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["portalUrl"] = undefined /*out*/;
            inputs["primarySharedKey"] = undefined /*out*/;
            inputs["secondarySharedKey"] = undefined /*out*/;
            inputs["workspaceId"] = undefined /*out*/;
        }
        super("azure:operationalinsights/analyticsWorkspace:AnalyticsWorkspace", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AnalyticsWorkspace resources.
 */
export interface AnalyticsWorkspaceState {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Portal URL for the Log Analytics Workspace.
     */
    readonly portalUrl?: pulumi.Input<string>;
    /**
     * The Primary shared key for the Log Analytics Workspace.
     */
    readonly primarySharedKey?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The workspace data retention in days. Possible values range between 30 and 730.
     */
    readonly retentionInDays?: pulumi.Input<number>;
    /**
     * The Secondary shared key for the Log Analytics Workspace.
     */
    readonly secondarySharedKey?: pulumi.Input<string>;
    /**
     * Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`).
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The Workspace (or Customer) ID for the Log Analytics Workspace.
     */
    readonly workspaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AnalyticsWorkspace resource.
 */
export interface AnalyticsWorkspaceArgs {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The workspace data retention in days. Possible values range between 30 and 730.
     */
    readonly retentionInDays?: pulumi.Input<number>;
    /**
     * Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`).
     */
    readonly sku: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
