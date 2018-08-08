// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a virtual network including any configured subnets. Each subnet can
 * optionally be configured with a security group to be associated with the subnet.
 * 
 * ~> **NOTE on Virtual Networks and Subnet's:** Terraform currently
 * provides both a standalone [Subnet resource](subnet.html), and allows for Subnets to be defined in-line within the [Virtual Network resource](virtual_network.html).
 * At this time you cannot use a Virtual Network with in-line Subnets in conjunction with any Subnet resources. Doing so will cause a conflict of Subnet configurations and will overwrite Subnet's.
 */
export class VirtualNetwork extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualNetworkState): VirtualNetwork {
        return new VirtualNetwork(name, <any>state, { id });
    }

    /**
     * The address space that is used the virtual
     * network. You can supply more than one address space. Changing this forces
     * a new resource to be created.
     */
    public readonly addressSpaces: pulumi.Output<string[]>;
    /**
     * List of IP addresses of DNS servers
     */
    public readonly dnsServers: pulumi.Output<string[] | undefined>;
    /**
     * The location/region where the virtual network is
     * created. Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * The name of the subnet.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The name of the resource group in which to
     * create the virtual network.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * Can be specified multiple times to define multiple
     * subnets. Each `subnet` block supports fields documented below.
     */
    public readonly subnets: pulumi.Output<{ addressPrefix: string, name: string, securityGroup?: string }[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a VirtualNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: VirtualNetworkArgs | VirtualNetworkState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: VirtualNetworkState = argsOrState as VirtualNetworkState | undefined;
            inputs["addressSpaces"] = state ? state.addressSpaces : undefined;
            inputs["dnsServers"] = state ? state.dnsServers : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["subnets"] = state ? state.subnets : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as VirtualNetworkArgs | undefined;
            if (!args || args.addressSpaces === undefined) {
                throw new Error("Missing required property 'addressSpaces'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["addressSpaces"] = args ? args.addressSpaces : undefined;
            inputs["dnsServers"] = args ? args.dnsServers : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subnets"] = args ? args.subnets : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        super("azure:network/virtualNetwork:VirtualNetwork", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualNetwork resources.
 */
export interface VirtualNetworkState {
    /**
     * The address space that is used the virtual
     * network. You can supply more than one address space. Changing this forces
     * a new resource to be created.
     */
    readonly addressSpaces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of IP addresses of DNS servers
     */
    readonly dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The location/region where the virtual network is
     * created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the subnet.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the virtual network.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Can be specified multiple times to define multiple
     * subnets. Each `subnet` block supports fields documented below.
     */
    readonly subnets?: pulumi.Input<pulumi.Input<{ addressPrefix: pulumi.Input<string>, name: pulumi.Input<string>, securityGroup?: pulumi.Input<string> }>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a VirtualNetwork resource.
 */
export interface VirtualNetworkArgs {
    /**
     * The address space that is used the virtual
     * network. You can supply more than one address space. Changing this forces
     * a new resource to be created.
     */
    readonly addressSpaces: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of IP addresses of DNS servers
     */
    readonly dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The location/region where the virtual network is
     * created. Changing this forces a new resource to be created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the subnet.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the virtual network.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Can be specified multiple times to define multiple
     * subnets. Each `subnet` block supports fields documented below.
     */
    readonly subnets?: pulumi.Input<pulumi.Input<{ addressPrefix: pulumi.Input<string>, name: pulumi.Input<string>, securityGroup?: pulumi.Input<string> }>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
