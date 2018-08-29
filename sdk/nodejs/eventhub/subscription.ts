// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Create a ServiceBus Subscription.
 */
export class Subscription extends pulumi.CustomResource {
    /**
     * Get an existing Subscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionState): Subscription {
        return new Subscription(name, <any>state, { id });
    }

    /**
     * The idle interval after which the
     * Subscription is automatically deleted, minimum of 5 minutes. Provided in the
     * TimeSpan format.
     */
    public readonly autoDeleteOnIdle: pulumi.Output<string>;
    public readonly deadLetteringOnFilterEvaluationExceptions: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag which controls
     * whether the Subscription has dead letter support when a message expires. Defaults
     * to false.
     */
    public readonly deadLetteringOnMessageExpiration: pulumi.Output<boolean | undefined>;
    /**
     * The TTL of messages sent to this Subscription
     * if no TTL value is set on the message itself. Provided in the TimeSpan
     * format.
     */
    public readonly defaultMessageTtl: pulumi.Output<string>;
    /**
     * Boolean flag which controls whether the
     * Subscription supports batched operations. Defaults to false.
     */
    public readonly enableBatchedOperations: pulumi.Output<boolean | undefined>;
    /**
     * The name of a Queue or Topic to automatically forward 
     * messages to.
     */
    public readonly forwardTo: pulumi.Output<string | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists.
     * Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string | undefined>;
    /**
     * The lock duration for the subscription, maximum
     * supported value is 5 minutes. Defaults to 1 minute.
     */
    public readonly lockDuration: pulumi.Output<string>;
    /**
     * The maximum number of deliveries.
     */
    public readonly maxDeliveryCount: pulumi.Output<number>;
    /**
     * Specifies the name of the ServiceBus Subscription resource.
     * Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The name of the ServiceBus Namespace to create
     * this Subscription in. Changing this forces a new resource to be created.
     */
    public readonly namespaceName: pulumi.Output<string>;
    /**
     * Boolean flag which controls whether this Subscription
     * supports the concept of a session. Defaults to false. Changing this forces a
     * new resource to be created.
     */
    public readonly requiresSession: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource group in which to
     * create the namespace. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * The name of the ServiceBus Topic to create
     * this Subscription in. Changing this forces a new resource to be created.
     */
    public readonly topicName: pulumi.Output<string>;

    /**
     * Create a Subscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionArgs | SubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SubscriptionState = argsOrState as SubscriptionState | undefined;
            inputs["autoDeleteOnIdle"] = state ? state.autoDeleteOnIdle : undefined;
            inputs["deadLetteringOnFilterEvaluationExceptions"] = state ? state.deadLetteringOnFilterEvaluationExceptions : undefined;
            inputs["deadLetteringOnMessageExpiration"] = state ? state.deadLetteringOnMessageExpiration : undefined;
            inputs["defaultMessageTtl"] = state ? state.defaultMessageTtl : undefined;
            inputs["enableBatchedOperations"] = state ? state.enableBatchedOperations : undefined;
            inputs["forwardTo"] = state ? state.forwardTo : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["lockDuration"] = state ? state.lockDuration : undefined;
            inputs["maxDeliveryCount"] = state ? state.maxDeliveryCount : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceName"] = state ? state.namespaceName : undefined;
            inputs["requiresSession"] = state ? state.requiresSession : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["topicName"] = state ? state.topicName : undefined;
        } else {
            const args = argsOrState as SubscriptionArgs | undefined;
            if (!args || args.maxDeliveryCount === undefined) {
                throw new Error("Missing required property 'maxDeliveryCount'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.topicName === undefined) {
                throw new Error("Missing required property 'topicName'");
            }
            inputs["autoDeleteOnIdle"] = args ? args.autoDeleteOnIdle : undefined;
            inputs["deadLetteringOnFilterEvaluationExceptions"] = args ? args.deadLetteringOnFilterEvaluationExceptions : undefined;
            inputs["deadLetteringOnMessageExpiration"] = args ? args.deadLetteringOnMessageExpiration : undefined;
            inputs["defaultMessageTtl"] = args ? args.defaultMessageTtl : undefined;
            inputs["enableBatchedOperations"] = args ? args.enableBatchedOperations : undefined;
            inputs["forwardTo"] = args ? args.forwardTo : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["lockDuration"] = args ? args.lockDuration : undefined;
            inputs["maxDeliveryCount"] = args ? args.maxDeliveryCount : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["requiresSession"] = args ? args.requiresSession : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["topicName"] = args ? args.topicName : undefined;
        }
        super("azure:eventhub/subscription:Subscription", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subscription resources.
 */
export interface SubscriptionState {
    /**
     * The idle interval after which the
     * Subscription is automatically deleted, minimum of 5 minutes. Provided in the
     * TimeSpan format.
     */
    readonly autoDeleteOnIdle?: pulumi.Input<string>;
    readonly deadLetteringOnFilterEvaluationExceptions?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls
     * whether the Subscription has dead letter support when a message expires. Defaults
     * to false.
     */
    readonly deadLetteringOnMessageExpiration?: pulumi.Input<boolean>;
    /**
     * The TTL of messages sent to this Subscription
     * if no TTL value is set on the message itself. Provided in the TimeSpan
     * format.
     */
    readonly defaultMessageTtl?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether the
     * Subscription supports batched operations. Defaults to false.
     */
    readonly enableBatchedOperations?: pulumi.Input<boolean>;
    /**
     * The name of a Queue or Topic to automatically forward 
     * messages to.
     */
    readonly forwardTo?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists.
     * Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The lock duration for the subscription, maximum
     * supported value is 5 minutes. Defaults to 1 minute.
     */
    readonly lockDuration?: pulumi.Input<string>;
    /**
     * The maximum number of deliveries.
     */
    readonly maxDeliveryCount?: pulumi.Input<number>;
    /**
     * Specifies the name of the ServiceBus Subscription resource.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Namespace to create
     * this Subscription in. Changing this forces a new resource to be created.
     */
    readonly namespaceName?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether this Subscription
     * supports the concept of a session. Defaults to false. Changing this forces a
     * new resource to be created.
     */
    readonly requiresSession?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to
     * create the namespace. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Topic to create
     * this Subscription in. Changing this forces a new resource to be created.
     */
    readonly topicName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subscription resource.
 */
export interface SubscriptionArgs {
    /**
     * The idle interval after which the
     * Subscription is automatically deleted, minimum of 5 minutes. Provided in the
     * TimeSpan format.
     */
    readonly autoDeleteOnIdle?: pulumi.Input<string>;
    readonly deadLetteringOnFilterEvaluationExceptions?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls
     * whether the Subscription has dead letter support when a message expires. Defaults
     * to false.
     */
    readonly deadLetteringOnMessageExpiration?: pulumi.Input<boolean>;
    /**
     * The TTL of messages sent to this Subscription
     * if no TTL value is set on the message itself. Provided in the TimeSpan
     * format.
     */
    readonly defaultMessageTtl?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether the
     * Subscription supports batched operations. Defaults to false.
     */
    readonly enableBatchedOperations?: pulumi.Input<boolean>;
    /**
     * The name of a Queue or Topic to automatically forward 
     * messages to.
     */
    readonly forwardTo?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists.
     * Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The lock duration for the subscription, maximum
     * supported value is 5 minutes. Defaults to 1 minute.
     */
    readonly lockDuration?: pulumi.Input<string>;
    /**
     * The maximum number of deliveries.
     */
    readonly maxDeliveryCount: pulumi.Input<number>;
    /**
     * Specifies the name of the ServiceBus Subscription resource.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Namespace to create
     * this Subscription in. Changing this forces a new resource to be created.
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether this Subscription
     * supports the concept of a session. Defaults to false. Changing this forces a
     * new resource to be created.
     */
    readonly requiresSession?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to
     * create the namespace. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Topic to create
     * this Subscription in. Changing this forces a new resource to be created.
     */
    readonly topicName: pulumi.Input<string>;
}
