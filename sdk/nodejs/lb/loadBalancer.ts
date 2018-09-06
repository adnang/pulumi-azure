// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Create a Load Balancer Resource.
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerState): LoadBalancer {
        return new LoadBalancer(name, <any>state, { id });
    }

    /**
     * A frontend ip configuration block as documented below.
     */
    public readonly frontendIpConfigurations: pulumi.Output<{ inboundNatRules: string[], loadBalancerRules: string[], name: string, privateIpAddress: string, privateIpAddressAllocation: string, publicIpAddressId: string, subnetId: string, zones?: string }[] | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * Specifies the name of the frontend ip configuration.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
     */
    public /*out*/ readonly privateIpAddress: pulumi.Output<string>;
    /**
     * The list of private IP address assigned to the load balancer in `frontend_ip_configuration` blocks, if any.
     */
    public /*out*/ readonly privateIpAddresses: pulumi.Output<string[]>;
    /**
     * The name of the resource group in which to create the LoadBalancer.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
     */
    public readonly sku: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: LoadBalancerState = argsOrState as LoadBalancerState | undefined;
            inputs["frontendIpConfigurations"] = state ? state.frontendIpConfigurations : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            inputs["privateIpAddresses"] = state ? state.privateIpAddresses : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as LoadBalancerArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["frontendIpConfigurations"] = args ? args.frontendIpConfigurations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["privateIpAddress"] = undefined /*out*/;
            inputs["privateIpAddresses"] = undefined /*out*/;
        }
        super("azure:lb/loadBalancer:LoadBalancer", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    /**
     * A frontend ip configuration block as documented below.
     */
    readonly frontendIpConfigurations?: pulumi.Input<pulumi.Input<{ inboundNatRules?: pulumi.Input<pulumi.Input<string>[]>, loadBalancerRules?: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string>, privateIpAddress?: pulumi.Input<string>, privateIpAddressAllocation?: pulumi.Input<string>, publicIpAddressId?: pulumi.Input<string>, subnetId?: pulumi.Input<string>, zones?: pulumi.Input<string> }>[]>;
    /**
     * Specifies the supported Azure location where the resource exists.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the frontend ip configuration.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
     */
    readonly privateIpAddress?: pulumi.Input<string>;
    /**
     * The list of private IP address assigned to the load balancer in `frontend_ip_configuration` blocks, if any.
     */
    readonly privateIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group in which to create the LoadBalancer.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * A frontend ip configuration block as documented below.
     */
    readonly frontendIpConfigurations?: pulumi.Input<pulumi.Input<{ inboundNatRules?: pulumi.Input<pulumi.Input<string>[]>, loadBalancerRules?: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string>, privateIpAddress?: pulumi.Input<string>, privateIpAddressAllocation?: pulumi.Input<string>, publicIpAddressId?: pulumi.Input<string>, subnetId?: pulumi.Input<string>, zones?: pulumi.Input<string> }>[]>;
    /**
     * Specifies the supported Azure location where the resource exists.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Specifies the name of the frontend ip configuration.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the LoadBalancer.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
