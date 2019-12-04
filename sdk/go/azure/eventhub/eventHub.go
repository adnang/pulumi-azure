// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Event Hubs as a nested resource within a Event Hubs namespace.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventhub.html.markdown.
type EventHub struct {
	pulumi.CustomResourceState

	// A `captureDescription` block as defined below.
	CaptureDescription EventHubCaptureDescriptionOutput `pulumi:"captureDescription"`

	Location pulumi.StringOutput `pulumi:"location"`

	// Specifies the number of days to retain the events for this Event Hub. Needs to be between 1 and 7 days; or 1 day when using a Basic SKU for the parent EventHub Namespace.
	MessageRetention pulumi.IntOutput `pulumi:"messageRetention"`

	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`

	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount pulumi.IntOutput `pulumi:"partitionCount"`

	// The identifiers for partitions created for Event Hubs.
	PartitionIds pulumi.StringArrayOutput `pulumi:"partitionIds"`

	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewEventHub registers a new resource with the given unique name, arguments, and options.
func NewEventHub(ctx *pulumi.Context,
	name string, args *EventHubArgs, opts ...pulumi.ResourceOption) (*EventHub, error) {
	if args == nil || args.MessageRetention == nil {
		return nil, errors.New("missing required argument 'MessageRetention'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.PartitionCount == nil {
		return nil, errors.New("missing required argument 'PartitionCount'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.CaptureDescription; i != nil { inputs["captureDescription"] = i.ToEventHubCaptureDescriptionOutput() }
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.MessageRetention; i != nil { inputs["messageRetention"] = i.ToIntOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.NamespaceName; i != nil { inputs["namespaceName"] = i.ToStringOutput() }
		if i := args.PartitionCount; i != nil { inputs["partitionCount"] = i.ToIntOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
	}
	var resource EventHub
	err := ctx.RegisterResource("azure:eventhub/eventHub:EventHub", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHub gets an existing EventHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubState, opts ...pulumi.ResourceOption) (*EventHub, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.CaptureDescription; i != nil { inputs["captureDescription"] = i.ToEventHubCaptureDescriptionOutput() }
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.MessageRetention; i != nil { inputs["messageRetention"] = i.ToIntOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.NamespaceName; i != nil { inputs["namespaceName"] = i.ToStringOutput() }
		if i := state.PartitionCount; i != nil { inputs["partitionCount"] = i.ToIntOutput() }
		if i := state.PartitionIds; i != nil { inputs["partitionIds"] = i.ToStringArrayOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
	}
	var resource EventHub
	err := ctx.ReadResource("azure:eventhub/eventHub:EventHub", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHub resources.
type EventHubState struct {
	// A `captureDescription` block as defined below.
	CaptureDescription EventHubCaptureDescriptionInput `pulumi:"captureDescription"`
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the number of days to retain the events for this Event Hub. Needs to be between 1 and 7 days; or 1 day when using a Basic SKU for the parent EventHub Namespace.
	MessageRetention pulumi.IntInput `pulumi:"messageRetention"`
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount pulumi.IntInput `pulumi:"partitionCount"`
	// The identifiers for partitions created for Event Hubs.
	PartitionIds pulumi.StringArrayInput `pulumi:"partitionIds"`
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EventHub resource.
type EventHubArgs struct {
	// A `captureDescription` block as defined below.
	CaptureDescription EventHubCaptureDescriptionInput `pulumi:"captureDescription"`
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the number of days to retain the events for this Event Hub. Needs to be between 1 and 7 days; or 1 day when using a Basic SKU for the parent EventHub Namespace.
	MessageRetention pulumi.IntInput `pulumi:"messageRetention"`
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount pulumi.IntInput `pulumi:"partitionCount"`
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}
type EventHubCaptureDescription struct {
	Destination EventHubCaptureDescriptionDestination `pulumi:"destination"`
	Enabled bool `pulumi:"enabled"`
	Encoding string `pulumi:"encoding"`
	IntervalInSeconds *int `pulumi:"intervalInSeconds"`
	SizeLimitInBytes *int `pulumi:"sizeLimitInBytes"`
	SkipEmptyArchives *bool `pulumi:"skipEmptyArchives"`
}
var eventHubCaptureDescriptionType = reflect.TypeOf((*EventHubCaptureDescription)(nil)).Elem()

type EventHubCaptureDescriptionInput interface {
	pulumi.Input

	ToEventHubCaptureDescriptionOutput() EventHubCaptureDescriptionOutput
	ToEventHubCaptureDescriptionOutputWithContext(ctx context.Context) EventHubCaptureDescriptionOutput
}

type EventHubCaptureDescriptionArgs struct {
	Destination EventHubCaptureDescriptionDestinationInput `pulumi:"destination"`
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	Encoding pulumi.StringInput `pulumi:"encoding"`
	IntervalInSeconds pulumi.IntInput `pulumi:"intervalInSeconds"`
	SizeLimitInBytes pulumi.IntInput `pulumi:"sizeLimitInBytes"`
	SkipEmptyArchives pulumi.BoolInput `pulumi:"skipEmptyArchives"`
}

func (EventHubCaptureDescriptionArgs) ElementType() reflect.Type {
	return eventHubCaptureDescriptionType
}

func (a EventHubCaptureDescriptionArgs) ToEventHubCaptureDescriptionOutput() EventHubCaptureDescriptionOutput {
	return pulumi.ToOutput(a).(EventHubCaptureDescriptionOutput)
}

func (a EventHubCaptureDescriptionArgs) ToEventHubCaptureDescriptionOutputWithContext(ctx context.Context) EventHubCaptureDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EventHubCaptureDescriptionOutput)
}

type EventHubCaptureDescriptionOutput struct { *pulumi.OutputState }

func (o EventHubCaptureDescriptionOutput) Destination() EventHubCaptureDescriptionDestinationOutput {
	return o.Apply(func(v EventHubCaptureDescription) EventHubCaptureDescriptionDestination {
		return v.Destination
	}).(EventHubCaptureDescriptionDestinationOutput)
}

func (o EventHubCaptureDescriptionOutput) Enabled() pulumi.BoolOutput {
	return o.Apply(func(v EventHubCaptureDescription) bool {
		return v.Enabled
	}).(pulumi.BoolOutput)
}

func (o EventHubCaptureDescriptionOutput) Encoding() pulumi.StringOutput {
	return o.Apply(func(v EventHubCaptureDescription) string {
		return v.Encoding
	}).(pulumi.StringOutput)
}

func (o EventHubCaptureDescriptionOutput) IntervalInSeconds() pulumi.IntOutput {
	return o.Apply(func(v EventHubCaptureDescription) int {
		if v.IntervalInSeconds == nil { return *new(int) } else { return *v.IntervalInSeconds }
	}).(pulumi.IntOutput)
}

func (o EventHubCaptureDescriptionOutput) SizeLimitInBytes() pulumi.IntOutput {
	return o.Apply(func(v EventHubCaptureDescription) int {
		if v.SizeLimitInBytes == nil { return *new(int) } else { return *v.SizeLimitInBytes }
	}).(pulumi.IntOutput)
}

func (o EventHubCaptureDescriptionOutput) SkipEmptyArchives() pulumi.BoolOutput {
	return o.Apply(func(v EventHubCaptureDescription) bool {
		if v.SkipEmptyArchives == nil { return *new(bool) } else { return *v.SkipEmptyArchives }
	}).(pulumi.BoolOutput)
}

func (EventHubCaptureDescriptionOutput) ElementType() reflect.Type {
	return eventHubCaptureDescriptionType
}

func (o EventHubCaptureDescriptionOutput) ToEventHubCaptureDescriptionOutput() EventHubCaptureDescriptionOutput {
	return o
}

func (o EventHubCaptureDescriptionOutput) ToEventHubCaptureDescriptionOutputWithContext(ctx context.Context) EventHubCaptureDescriptionOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EventHubCaptureDescriptionOutput{}) }

type EventHubCaptureDescriptionDestination struct {
	ArchiveNameFormat string `pulumi:"archiveNameFormat"`
	BlobContainerName string `pulumi:"blobContainerName"`
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name string `pulumi:"name"`
	StorageAccountId string `pulumi:"storageAccountId"`
}
var eventHubCaptureDescriptionDestinationType = reflect.TypeOf((*EventHubCaptureDescriptionDestination)(nil)).Elem()

type EventHubCaptureDescriptionDestinationInput interface {
	pulumi.Input

	ToEventHubCaptureDescriptionDestinationOutput() EventHubCaptureDescriptionDestinationOutput
	ToEventHubCaptureDescriptionDestinationOutputWithContext(ctx context.Context) EventHubCaptureDescriptionDestinationOutput
}

type EventHubCaptureDescriptionDestinationArgs struct {
	ArchiveNameFormat pulumi.StringInput `pulumi:"archiveNameFormat"`
	BlobContainerName pulumi.StringInput `pulumi:"blobContainerName"`
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
}

func (EventHubCaptureDescriptionDestinationArgs) ElementType() reflect.Type {
	return eventHubCaptureDescriptionDestinationType
}

func (a EventHubCaptureDescriptionDestinationArgs) ToEventHubCaptureDescriptionDestinationOutput() EventHubCaptureDescriptionDestinationOutput {
	return pulumi.ToOutput(a).(EventHubCaptureDescriptionDestinationOutput)
}

func (a EventHubCaptureDescriptionDestinationArgs) ToEventHubCaptureDescriptionDestinationOutputWithContext(ctx context.Context) EventHubCaptureDescriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EventHubCaptureDescriptionDestinationOutput)
}

type EventHubCaptureDescriptionDestinationOutput struct { *pulumi.OutputState }

func (o EventHubCaptureDescriptionDestinationOutput) ArchiveNameFormat() pulumi.StringOutput {
	return o.Apply(func(v EventHubCaptureDescriptionDestination) string {
		return v.ArchiveNameFormat
	}).(pulumi.StringOutput)
}

func (o EventHubCaptureDescriptionDestinationOutput) BlobContainerName() pulumi.StringOutput {
	return o.Apply(func(v EventHubCaptureDescriptionDestination) string {
		return v.BlobContainerName
	}).(pulumi.StringOutput)
}

// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
func (o EventHubCaptureDescriptionDestinationOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v EventHubCaptureDescriptionDestination) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o EventHubCaptureDescriptionDestinationOutput) StorageAccountId() pulumi.StringOutput {
	return o.Apply(func(v EventHubCaptureDescriptionDestination) string {
		return v.StorageAccountId
	}).(pulumi.StringOutput)
}

func (EventHubCaptureDescriptionDestinationOutput) ElementType() reflect.Type {
	return eventHubCaptureDescriptionDestinationType
}

func (o EventHubCaptureDescriptionDestinationOutput) ToEventHubCaptureDescriptionDestinationOutput() EventHubCaptureDescriptionDestinationOutput {
	return o
}

func (o EventHubCaptureDescriptionDestinationOutput) ToEventHubCaptureDescriptionDestinationOutputWithContext(ctx context.Context) EventHubCaptureDescriptionDestinationOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EventHubCaptureDescriptionDestinationOutput{}) }

