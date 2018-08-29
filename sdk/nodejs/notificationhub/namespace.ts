// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a Notification Hub Namespace.
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceState): Namespace {
        return new Namespace(name, <any>state, { id });
    }

    /**
     * Is this Notification Hub Namespace enabled? Defaults to `true`.
     */
    public readonly enabled: pulumi.Output<boolean | undefined>;
    /**
     * The Azure Region in which this Notification Hub Namespace should be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
     */
    public readonly namespaceType: pulumi.Output<string>;
    /**
     * The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * The ServiceBus Endpoint for this Notification Hub Namespace.
     */
    public /*out*/ readonly servicebusEndpoint: pulumi.Output<string>;
    /**
     * A `sku` block as defined below.
     */
    public readonly sku: pulumi.Output<{ name: string }>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceArgs | NamespaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: NamespaceState = argsOrState as NamespaceState | undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceType"] = state ? state.namespaceType : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["servicebusEndpoint"] = state ? state.servicebusEndpoint : undefined;
            inputs["sku"] = state ? state.sku : undefined;
        } else {
            const args = argsOrState as NamespaceArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.namespaceType === undefined) {
                throw new Error("Missing required property 'namespaceType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceType"] = args ? args.namespaceType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["servicebusEndpoint"] = undefined /*out*/;
        }
        super("azure:notificationhub/namespace:Namespace", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Namespace resources.
 */
export interface NamespaceState {
    /**
     * Is this Notification Hub Namespace enabled? Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The Azure Region in which this Notification Hub Namespace should be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
     */
    readonly namespaceType?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The ServiceBus Endpoint for this Notification Hub Namespace.
     */
    readonly servicebusEndpoint?: pulumi.Input<string>;
    /**
     * A `sku` block as defined below.
     */
    readonly sku?: pulumi.Input<{ name: pulumi.Input<string> }>;
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * Is this Notification Hub Namespace enabled? Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The Azure Region in which this Notification Hub Namespace should be created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
     */
    readonly namespaceType: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `sku` block as defined below.
     */
    readonly sku: pulumi.Input<{ name: pulumi.Input<string> }>;
}
