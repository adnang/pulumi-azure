// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Azure SignalR service.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleService = new azure.signalr.Service("exampleService", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: {
 *         name: "Free_F1",
 *         capacity: 1,
 *     },
 *     cors: [{
 *         allowedOrigins: ["http://www.example.com"],
 *     }],
 *     features: [{
 *         flag: "ServiceMode",
 *         value: "Default",
 *     }],
 * });
 * ```
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:signalr/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * A `cors` block as documented below.
     */
    public readonly cors!: pulumi.Output<outputs.signalr.ServiceCor[]>;
    /**
     * A `features` block as documented below.
     */
    public readonly features!: pulumi.Output<outputs.signalr.ServiceFeature[]>;
    /**
     * The FQDN of the SignalR service.
     */
    public /*out*/ readonly hostname!: pulumi.Output<string>;
    /**
     * The publicly accessible IP of the SignalR service.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the SignalR service. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The primary access key for the SignalR service.
     */
    public /*out*/ readonly primaryAccessKey!: pulumi.Output<string>;
    /**
     * The primary connection string for the SignalR service.
     */
    public /*out*/ readonly primaryConnectionString!: pulumi.Output<string>;
    /**
     * The publicly accessible port of the SignalR service which is designed for browser/client use.
     */
    public /*out*/ readonly publicPort!: pulumi.Output<number>;
    /**
     * The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The secondary access key for the SignalR service.
     */
    public /*out*/ readonly secondaryAccessKey!: pulumi.Output<string>;
    /**
     * The secondary connection string for the SignalR service.
     */
    public /*out*/ readonly secondaryConnectionString!: pulumi.Output<string>;
    /**
     * The publicly accessible port of the SignalR service which is designed for customer server side use.
     */
    public /*out*/ readonly serverPort!: pulumi.Output<number>;
    /**
     * A `sku` block as documented below.
     */
    public readonly sku!: pulumi.Output<outputs.signalr.ServiceSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceState | undefined;
            inputs["cors"] = state ? state.cors : undefined;
            inputs["features"] = state ? state.features : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["primaryAccessKey"] = state ? state.primaryAccessKey : undefined;
            inputs["primaryConnectionString"] = state ? state.primaryConnectionString : undefined;
            inputs["publicPort"] = state ? state.publicPort : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryAccessKey"] = state ? state.secondaryAccessKey : undefined;
            inputs["secondaryConnectionString"] = state ? state.secondaryConnectionString : undefined;
            inputs["serverPort"] = state ? state.serverPort : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["cors"] = args ? args.cors : undefined;
            inputs["features"] = args ? args.features : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["hostname"] = undefined /*out*/;
            inputs["ipAddress"] = undefined /*out*/;
            inputs["primaryAccessKey"] = undefined /*out*/;
            inputs["primaryConnectionString"] = undefined /*out*/;
            inputs["publicPort"] = undefined /*out*/;
            inputs["secondaryAccessKey"] = undefined /*out*/;
            inputs["secondaryConnectionString"] = undefined /*out*/;
            inputs["serverPort"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Service.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * A `cors` block as documented below.
     */
    readonly cors?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceCor>[]>;
    /**
     * A `features` block as documented below.
     */
    readonly features?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceFeature>[]>;
    /**
     * The FQDN of the SignalR service.
     */
    readonly hostname?: pulumi.Input<string>;
    /**
     * The publicly accessible IP of the SignalR service.
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the SignalR service. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The primary access key for the SignalR service.
     */
    readonly primaryAccessKey?: pulumi.Input<string>;
    /**
     * The primary connection string for the SignalR service.
     */
    readonly primaryConnectionString?: pulumi.Input<string>;
    /**
     * The publicly accessible port of the SignalR service which is designed for browser/client use.
     */
    readonly publicPort?: pulumi.Input<number>;
    /**
     * The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The secondary access key for the SignalR service.
     */
    readonly secondaryAccessKey?: pulumi.Input<string>;
    /**
     * The secondary connection string for the SignalR service.
     */
    readonly secondaryConnectionString?: pulumi.Input<string>;
    /**
     * The publicly accessible port of the SignalR service which is designed for customer server side use.
     */
    readonly serverPort?: pulumi.Input<number>;
    /**
     * A `sku` block as documented below.
     */
    readonly sku?: pulumi.Input<inputs.signalr.ServiceSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * A `cors` block as documented below.
     */
    readonly cors?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceCor>[]>;
    /**
     * A `features` block as documented below.
     */
    readonly features?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceFeature>[]>;
    /**
     * Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the SignalR service. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `sku` block as documented below.
     */
    readonly sku: pulumi.Input<inputs.signalr.ServiceSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
