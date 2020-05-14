// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a ServiceBus Queue.
//
// Deprecated: azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue
type Queue struct {
	pulumi.CustomResourceState

	// The ISO 8601 timespan duration of the idle interval after which the
	// Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringOutput `pulumi:"autoDeleteOnIdle"`
	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrOutput `pulumi:"deadLetteringOnMessageExpiration"`
	// The ISO 8601 timespan duration of the TTL of messages sent to this
	// queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTtl pulumi.StringOutput `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Default value is 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow pulumi.StringOutput `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls whether Express Entities
	// are enabled. An express queue holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST
	// be set to `false`.
	EnableExpress pulumi.BoolPtrOutput `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable
	// the queue to be partitioned across multiple message brokers. Changing this forces
	// a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST
	// be set to `true`.
	EnablePartitioning pulumi.BoolPtrOutput `pulumi:"enablePartitioning"`
	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute. (`PT1M`)
	LockDuration pulumi.StringOutput `pulumi:"lockDuration"`
	// Integer value which controls when a message is automatically deadlettered. Defaults to `10`.
	MaxDeliveryCount pulumi.IntPtrOutput `pulumi:"maxDeliveryCount"`
	// Integer value which controls the size of
	// memory allocated for the queue. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes pulumi.IntOutput `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Queue resource. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the ServiceBus Namespace to create
	// this queue in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Boolean flag which controls whether
	// the Queue requires duplicate detection. Changing this forces
	// a new resource to be created. Defaults to `false`.
	RequiresDuplicateDetection pulumi.BoolPtrOutput `pulumi:"requiresDuplicateDetection"`
	// Boolean flag which controls whether the Queue requires sessions.
	// This will allow ordered handling of unbounded sequences of related messages. With sessions enabled
	// a queue can guarantee first-in-first-out delivery of messages.
	// Changing this forces a new resource to be created. Defaults to `false`.
	RequiresSession pulumi.BoolPtrOutput `pulumi:"requiresSession"`
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &QueueArgs{}
	}
	var resource Queue
	err := ctx.RegisterResource("azure:eventhub/queue:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("azure:eventhub/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration *bool `pulumi:"deadLetteringOnMessageExpiration"`
	// The ISO 8601 timespan duration of the TTL of messages sent to this
	// queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTtl *string `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Default value is 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls whether Express Entities
	// are enabled. An express queue holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST
	// be set to `false`.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable
	// the queue to be partitioned across multiple message brokers. Changing this forces
	// a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST
	// be set to `true`.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute. (`PT1M`)
	LockDuration *string `pulumi:"lockDuration"`
	// Integer value which controls when a message is automatically deadlettered. Defaults to `10`.
	MaxDeliveryCount *int `pulumi:"maxDeliveryCount"`
	// Integer value which controls the size of
	// memory allocated for the queue. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Queue resource. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the ServiceBus Namespace to create
	// this queue in. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// Boolean flag which controls whether
	// the Queue requires duplicate detection. Changing this forces
	// a new resource to be created. Defaults to `false`.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// Boolean flag which controls whether the Queue requires sessions.
	// This will allow ordered handling of unbounded sequences of related messages. With sessions enabled
	// a queue can guarantee first-in-first-out delivery of messages.
	// Changing this forces a new resource to be created. Defaults to `false`.
	RequiresSession *bool `pulumi:"requiresSession"`
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type QueueState struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrInput
	// The ISO 8601 timespan duration of the TTL of messages sent to this
	// queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTtl pulumi.StringPtrInput
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Default value is 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Boolean flag which controls whether Express Entities
	// are enabled. An express queue holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST
	// be set to `false`.
	EnableExpress pulumi.BoolPtrInput
	// Boolean flag which controls whether to enable
	// the queue to be partitioned across multiple message brokers. Changing this forces
	// a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST
	// be set to `true`.
	EnablePartitioning pulumi.BoolPtrInput
	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute. (`PT1M`)
	LockDuration pulumi.StringPtrInput
	// Integer value which controls when a message is automatically deadlettered. Defaults to `10`.
	MaxDeliveryCount pulumi.IntPtrInput
	// Integer value which controls the size of
	// memory allocated for the queue. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes pulumi.IntPtrInput
	// Specifies the name of the ServiceBus Queue resource. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the ServiceBus Namespace to create
	// this queue in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// Boolean flag which controls whether
	// the Queue requires duplicate detection. Changing this forces
	// a new resource to be created. Defaults to `false`.
	RequiresDuplicateDetection pulumi.BoolPtrInput
	// Boolean flag which controls whether the Queue requires sessions.
	// This will allow ordered handling of unbounded sequences of related messages. With sessions enabled
	// a queue can guarantee first-in-first-out delivery of messages.
	// Changing this forces a new resource to be created. Defaults to `false`.
	RequiresSession pulumi.BoolPtrInput
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration *bool `pulumi:"deadLetteringOnMessageExpiration"`
	// The ISO 8601 timespan duration of the TTL of messages sent to this
	// queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTtl *string `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Default value is 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls whether Express Entities
	// are enabled. An express queue holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST
	// be set to `false`.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable
	// the queue to be partitioned across multiple message brokers. Changing this forces
	// a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST
	// be set to `true`.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute. (`PT1M`)
	LockDuration *string `pulumi:"lockDuration"`
	// Integer value which controls when a message is automatically deadlettered. Defaults to `10`.
	MaxDeliveryCount *int `pulumi:"maxDeliveryCount"`
	// Integer value which controls the size of
	// memory allocated for the queue. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Queue resource. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the ServiceBus Namespace to create
	// this queue in. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// Boolean flag which controls whether
	// the Queue requires duplicate detection. Changing this forces
	// a new resource to be created. Defaults to `false`.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// Boolean flag which controls whether the Queue requires sessions.
	// This will allow ordered handling of unbounded sequences of related messages. With sessions enabled
	// a queue can guarantee first-in-first-out delivery of messages.
	// Changing this forces a new resource to be created. Defaults to `false`.
	RequiresSession *bool `pulumi:"requiresSession"`
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrInput
	// The ISO 8601 timespan duration of the TTL of messages sent to this
	// queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTtl pulumi.StringPtrInput
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Default value is 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Boolean flag which controls whether Express Entities
	// are enabled. An express queue holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST
	// be set to `false`.
	EnableExpress pulumi.BoolPtrInput
	// Boolean flag which controls whether to enable
	// the queue to be partitioned across multiple message brokers. Changing this forces
	// a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST
	// be set to `true`.
	EnablePartitioning pulumi.BoolPtrInput
	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute. (`PT1M`)
	LockDuration pulumi.StringPtrInput
	// Integer value which controls when a message is automatically deadlettered. Defaults to `10`.
	MaxDeliveryCount pulumi.IntPtrInput
	// Integer value which controls the size of
	// memory allocated for the queue. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes pulumi.IntPtrInput
	// Specifies the name of the ServiceBus Queue resource. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the ServiceBus Namespace to create
	// this queue in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// Boolean flag which controls whether
	// the Queue requires duplicate detection. Changing this forces
	// a new resource to be created. Defaults to `false`.
	RequiresDuplicateDetection pulumi.BoolPtrInput
	// Boolean flag which controls whether the Queue requires sessions.
	// This will allow ordered handling of unbounded sequences of related messages. With sessions enabled
	// a queue can guarantee first-in-first-out delivery of messages.
	// Changing this forces a new resource to be created. Defaults to `false`.
	RequiresSession pulumi.BoolPtrInput
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}
