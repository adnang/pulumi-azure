// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Custom Action within a Logic App Workflow
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "East US"});
 * const exampleWorkflow = new azure.logicapps.Workflow("exampleWorkflow", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleActionCustom = new azure.logicapps.ActionCustom("exampleActionCustom", {
 *     logicAppId: exampleWorkflow.id,
 *     body: `{
 *     "description": "A variable to configure the auto expiration age in days. Configured in negative number. Default is -30 (30 days old).",
 *     "inputs": {
 *         "variables": [
 *             {
 *                 "name": "ExpirationAgeInDays",
 *                 "type": "Integer",
 *                 "value": -30
 *             }
 *         ]
 *     },
 *     "runAfter": {},
 *     "type": "InitializeVariable"
 * }
 * `,
 * });
 * ```
 */
export class ActionCustom extends pulumi.CustomResource {
    /**
     * Get an existing ActionCustom resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionCustomState, opts?: pulumi.CustomResourceOptions): ActionCustom {
        return new ActionCustom(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:logicapps/actionCustom:ActionCustom';

    /**
     * Returns true if the given object is an instance of ActionCustom.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionCustom {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionCustom.__pulumiType;
    }

    /**
     * Specifies the JSON Blob defining the Body of this Custom Action.
     */
    public readonly body!: pulumi.Output<string>;
    /**
     * Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
     */
    public readonly logicAppId!: pulumi.Output<string>;
    /**
     * Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ActionCustom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionCustomArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionCustomArgs | ActionCustomState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ActionCustomState | undefined;
            inputs["body"] = state ? state.body : undefined;
            inputs["logicAppId"] = state ? state.logicAppId : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ActionCustomArgs | undefined;
            if (!args || args.body === undefined) {
                throw new Error("Missing required property 'body'");
            }
            if (!args || args.logicAppId === undefined) {
                throw new Error("Missing required property 'logicAppId'");
            }
            inputs["body"] = args ? args.body : undefined;
            inputs["logicAppId"] = args ? args.logicAppId : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ActionCustom.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionCustom resources.
 */
export interface ActionCustomState {
    /**
     * Specifies the JSON Blob defining the Body of this Custom Action.
     */
    readonly body?: pulumi.Input<string>;
    /**
     * Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
     */
    readonly logicAppId?: pulumi.Input<string>;
    /**
     * Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ActionCustom resource.
 */
export interface ActionCustomArgs {
    /**
     * Specifies the JSON Blob defining the Body of this Custom Action.
     */
    readonly body: pulumi.Input<string>;
    /**
     * Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
     */
    readonly logicAppId: pulumi.Input<string>;
    /**
     * Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
}
