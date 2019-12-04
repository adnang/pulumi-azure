// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Stream Analytics Output to a ServiceBus Topic.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_servicebus_topic.html.markdown.
type OutputServicebusTopic struct {
	pulumi.CustomResourceState

	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// A `serialization` block as defined below.
	Serialization OutputServicebusTopicSerializationOutput `pulumi:"serialization"`

	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringOutput `pulumi:"servicebusNamespace"`

	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringOutput `pulumi:"sharedAccessPolicyKey"`

	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringOutput `pulumi:"sharedAccessPolicyName"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`

	// The name of the Service Bus Topic.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewOutputServicebusTopic registers a new resource with the given unique name, arguments, and options.
func NewOutputServicebusTopic(ctx *pulumi.Context,
	name string, args *OutputServicebusTopicArgs, opts ...pulumi.ResourceOption) (*OutputServicebusTopic, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Serialization == nil {
		return nil, errors.New("missing required argument 'Serialization'")
	}
	if args == nil || args.ServicebusNamespace == nil {
		return nil, errors.New("missing required argument 'ServicebusNamespace'")
	}
	if args == nil || args.SharedAccessPolicyKey == nil {
		return nil, errors.New("missing required argument 'SharedAccessPolicyKey'")
	}
	if args == nil || args.SharedAccessPolicyName == nil {
		return nil, errors.New("missing required argument 'SharedAccessPolicyName'")
	}
	if args == nil || args.StreamAnalyticsJobName == nil {
		return nil, errors.New("missing required argument 'StreamAnalyticsJobName'")
	}
	if args == nil || args.TopicName == nil {
		return nil, errors.New("missing required argument 'TopicName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Serialization; i != nil { inputs["serialization"] = i.ToOutputServicebusTopicSerializationOutput() }
		if i := args.ServicebusNamespace; i != nil { inputs["servicebusNamespace"] = i.ToStringOutput() }
		if i := args.SharedAccessPolicyKey; i != nil { inputs["sharedAccessPolicyKey"] = i.ToStringOutput() }
		if i := args.SharedAccessPolicyName; i != nil { inputs["sharedAccessPolicyName"] = i.ToStringOutput() }
		if i := args.StreamAnalyticsJobName; i != nil { inputs["streamAnalyticsJobName"] = i.ToStringOutput() }
		if i := args.TopicName; i != nil { inputs["topicName"] = i.ToStringOutput() }
	}
	var resource OutputServicebusTopic
	err := ctx.RegisterResource("azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputServicebusTopic gets an existing OutputServicebusTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputServicebusTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputServicebusTopicState, opts ...pulumi.ResourceOption) (*OutputServicebusTopic, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Serialization; i != nil { inputs["serialization"] = i.ToOutputServicebusTopicSerializationOutput() }
		if i := state.ServicebusNamespace; i != nil { inputs["servicebusNamespace"] = i.ToStringOutput() }
		if i := state.SharedAccessPolicyKey; i != nil { inputs["sharedAccessPolicyKey"] = i.ToStringOutput() }
		if i := state.SharedAccessPolicyName; i != nil { inputs["sharedAccessPolicyName"] = i.ToStringOutput() }
		if i := state.StreamAnalyticsJobName; i != nil { inputs["streamAnalyticsJobName"] = i.ToStringOutput() }
		if i := state.TopicName; i != nil { inputs["topicName"] = i.ToStringOutput() }
	}
	var resource OutputServicebusTopic
	err := ctx.ReadResource("azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputServicebusTopic resources.
type OutputServicebusTopicState struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputServicebusTopicSerializationInput `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringInput `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringInput `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringInput `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput `pulumi:"streamAnalyticsJobName"`
	// The name of the Service Bus Topic.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

// The set of arguments for constructing a OutputServicebusTopic resource.
type OutputServicebusTopicArgs struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputServicebusTopicSerializationInput `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringInput `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringInput `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringInput `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput `pulumi:"streamAnalyticsJobName"`
	// The name of the Service Bus Topic.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}
type OutputServicebusTopicSerialization struct {
	Encoding *string `pulumi:"encoding"`
	FieldDelimiter *string `pulumi:"fieldDelimiter"`
	Format *string `pulumi:"format"`
	Type string `pulumi:"type"`
}
var outputServicebusTopicSerializationType = reflect.TypeOf((*OutputServicebusTopicSerialization)(nil)).Elem()

type OutputServicebusTopicSerializationInput interface {
	pulumi.Input

	ToOutputServicebusTopicSerializationOutput() OutputServicebusTopicSerializationOutput
	ToOutputServicebusTopicSerializationOutputWithContext(ctx context.Context) OutputServicebusTopicSerializationOutput
}

type OutputServicebusTopicSerializationArgs struct {
	Encoding pulumi.StringInput `pulumi:"encoding"`
	FieldDelimiter pulumi.StringInput `pulumi:"fieldDelimiter"`
	Format pulumi.StringInput `pulumi:"format"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (OutputServicebusTopicSerializationArgs) ElementType() reflect.Type {
	return outputServicebusTopicSerializationType
}

func (a OutputServicebusTopicSerializationArgs) ToOutputServicebusTopicSerializationOutput() OutputServicebusTopicSerializationOutput {
	return pulumi.ToOutput(a).(OutputServicebusTopicSerializationOutput)
}

func (a OutputServicebusTopicSerializationArgs) ToOutputServicebusTopicSerializationOutputWithContext(ctx context.Context) OutputServicebusTopicSerializationOutput {
	return pulumi.ToOutputWithContext(ctx, a).(OutputServicebusTopicSerializationOutput)
}

type OutputServicebusTopicSerializationOutput struct { *pulumi.OutputState }

func (o OutputServicebusTopicSerializationOutput) Encoding() pulumi.StringOutput {
	return o.Apply(func(v OutputServicebusTopicSerialization) string {
		if v.Encoding == nil { return *new(string) } else { return *v.Encoding }
	}).(pulumi.StringOutput)
}

func (o OutputServicebusTopicSerializationOutput) FieldDelimiter() pulumi.StringOutput {
	return o.Apply(func(v OutputServicebusTopicSerialization) string {
		if v.FieldDelimiter == nil { return *new(string) } else { return *v.FieldDelimiter }
	}).(pulumi.StringOutput)
}

func (o OutputServicebusTopicSerializationOutput) Format() pulumi.StringOutput {
	return o.Apply(func(v OutputServicebusTopicSerialization) string {
		if v.Format == nil { return *new(string) } else { return *v.Format }
	}).(pulumi.StringOutput)
}

func (o OutputServicebusTopicSerializationOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v OutputServicebusTopicSerialization) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (OutputServicebusTopicSerializationOutput) ElementType() reflect.Type {
	return outputServicebusTopicSerializationType
}

func (o OutputServicebusTopicSerializationOutput) ToOutputServicebusTopicSerializationOutput() OutputServicebusTopicSerializationOutput {
	return o
}

func (o OutputServicebusTopicSerializationOutput) ToOutputServicebusTopicSerializationOutputWithContext(ctx context.Context) OutputServicebusTopicSerializationOutput {
	return o
}

func init() { pulumi.RegisterOutputType(OutputServicebusTopicSerializationOutput{}) }

