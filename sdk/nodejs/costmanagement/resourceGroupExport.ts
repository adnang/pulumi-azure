// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Azure Cost Management Export for a Resource Group.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/cost_management_export_resource_group.html.markdown.
 */
export class ResourceGroupExport extends pulumi.CustomResource {
    /**
     * Get an existing ResourceGroupExport resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceGroupExportState, opts?: pulumi.CustomResourceOptions): ResourceGroupExport {
        return new ResourceGroupExport(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:costmanagement/resourceGroupExport:ResourceGroupExport';

    /**
     * Returns true if the given object is an instance of ResourceGroupExport.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceGroupExport {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceGroupExport.__pulumiType;
    }

    /**
     * Is the cost management export active? Default is `true`.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * A `deliveryInfo` block as defined below.
     */
    public readonly deliveryInfo!: pulumi.Output<outputs.costmanagement.ResourceGroupExportDeliveryInfo>;
    /**
     * Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `query` block as defined below.
     */
    public readonly query!: pulumi.Output<outputs.costmanagement.ResourceGroupExportQuery>;
    /**
     * The date the export will stop capturing information. 
     */
    public readonly recurrencePeriodEnd!: pulumi.Output<string>;
    /**
     * The date the export will start capturing information.
     */
    public readonly recurrencePeriodStart!: pulumi.Output<string>;
    /**
     * How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
     */
    public readonly recurrenceType!: pulumi.Output<string>;
    /**
     * The id of the resource group in which to export information.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;

    /**
     * Create a ResourceGroupExport resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceGroupExportArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceGroupExportArgs | ResourceGroupExportState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResourceGroupExportState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["deliveryInfo"] = state ? state.deliveryInfo : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["query"] = state ? state.query : undefined;
            inputs["recurrencePeriodEnd"] = state ? state.recurrencePeriodEnd : undefined;
            inputs["recurrencePeriodStart"] = state ? state.recurrencePeriodStart : undefined;
            inputs["recurrenceType"] = state ? state.recurrenceType : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
        } else {
            const args = argsOrState as ResourceGroupExportArgs | undefined;
            if (!args || args.deliveryInfo === undefined) {
                throw new Error("Missing required property 'deliveryInfo'");
            }
            if (!args || args.query === undefined) {
                throw new Error("Missing required property 'query'");
            }
            if (!args || args.recurrencePeriodEnd === undefined) {
                throw new Error("Missing required property 'recurrencePeriodEnd'");
            }
            if (!args || args.recurrencePeriodStart === undefined) {
                throw new Error("Missing required property 'recurrencePeriodStart'");
            }
            if (!args || args.recurrenceType === undefined) {
                throw new Error("Missing required property 'recurrenceType'");
            }
            if (!args || args.resourceGroupId === undefined) {
                throw new Error("Missing required property 'resourceGroupId'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["deliveryInfo"] = args ? args.deliveryInfo : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["query"] = args ? args.query : undefined;
            inputs["recurrencePeriodEnd"] = args ? args.recurrencePeriodEnd : undefined;
            inputs["recurrencePeriodStart"] = args ? args.recurrencePeriodStart : undefined;
            inputs["recurrenceType"] = args ? args.recurrenceType : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ResourceGroupExport.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceGroupExport resources.
 */
export interface ResourceGroupExportState {
    /**
     * Is the cost management export active? Default is `true`.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * A `deliveryInfo` block as defined below.
     */
    readonly deliveryInfo?: pulumi.Input<inputs.costmanagement.ResourceGroupExportDeliveryInfo>;
    /**
     * Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `query` block as defined below.
     */
    readonly query?: pulumi.Input<inputs.costmanagement.ResourceGroupExportQuery>;
    /**
     * The date the export will stop capturing information. 
     */
    readonly recurrencePeriodEnd?: pulumi.Input<string>;
    /**
     * The date the export will start capturing information.
     */
    readonly recurrencePeriodStart?: pulumi.Input<string>;
    /**
     * How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
     */
    readonly recurrenceType?: pulumi.Input<string>;
    /**
     * The id of the resource group in which to export information.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourceGroupExport resource.
 */
export interface ResourceGroupExportArgs {
    /**
     * Is the cost management export active? Default is `true`.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * A `deliveryInfo` block as defined below.
     */
    readonly deliveryInfo: pulumi.Input<inputs.costmanagement.ResourceGroupExportDeliveryInfo>;
    /**
     * Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `query` block as defined below.
     */
    readonly query: pulumi.Input<inputs.costmanagement.ResourceGroupExportQuery>;
    /**
     * The date the export will stop capturing information. 
     */
    readonly recurrencePeriodEnd: pulumi.Input<string>;
    /**
     * The date the export will start capturing information.
     */
    readonly recurrencePeriodStart: pulumi.Input<string>;
    /**
     * How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
     */
    readonly recurrenceType: pulumi.Input<string>;
    /**
     * The id of the resource group in which to export information.
     */
    readonly resourceGroupId: pulumi.Input<string>;
}