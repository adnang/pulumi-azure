// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package streamanalytics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Stream Analytics Job.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_job.html.markdown.
type Job struct {
	pulumi.CustomResourceState

	// Specifies the compatibility level for this job - which controls certain runtime behaviors of the streaming job. Possible values are `1.0` and 1.1`.
	CompatibilityLevel pulumi.StringOutput `pulumi:"compatibilityLevel"`
	// Specifies the Data Locale of the Job, which [should be a supported .NET Culture](https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx).
	DataLocale pulumi.StringOutput `pulumi:"dataLocale"`
	// Specifies the maximum tolerable delay in seconds where events arriving late could be included. Supported range is `-1` (indefinite) to `1814399` (20d 23h 59m 59s).  Default is `0`.
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrOutput `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// Specifies the maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order. Supported range is `0` to `599` (9m 59s). Default is `5`.
	EventsOutOfOrderMaxDelayInSeconds pulumi.IntPtrOutput `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Specifies the policy which should be applied to events which arrive out of order in the input event stream. Possible values are `Adjust` and `Drop`.  Default is `Adjust`.
	EventsOutOfOrderPolicy pulumi.StringPtrOutput `pulumi:"eventsOutOfOrderPolicy"`
	// The Job ID assigned by the Stream Analytics Job.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// The Azure Region in which the Resource Group exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size). Possible values are `Drop` and `Stop`.  Default is `Drop`.
	OutputErrorPolicy pulumi.StringPtrOutput `pulumi:"outputErrorPolicy"`
	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the number of streaming units that the streaming job uses. Supported values are `1`, `3`, `6` and multiples of `6` up to `120`.
	StreamingUnits pulumi.IntOutput `pulumi:"streamingUnits"`
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998).
	TransformationQuery pulumi.StringOutput `pulumi:"transformationQuery"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StreamingUnits == nil {
		return nil, errors.New("missing required argument 'StreamingUnits'")
	}
	if args == nil || args.TransformationQuery == nil {
		return nil, errors.New("missing required argument 'TransformationQuery'")
	}
	if args == nil {
		args = &JobArgs{}
	}
	var resource Job
	err := ctx.RegisterResource("azure:streamanalytics/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure:streamanalytics/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// Specifies the compatibility level for this job - which controls certain runtime behaviors of the streaming job. Possible values are `1.0` and 1.1`.
	CompatibilityLevel *string `pulumi:"compatibilityLevel"`
	// Specifies the Data Locale of the Job, which [should be a supported .NET Culture](https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx).
	DataLocale *string `pulumi:"dataLocale"`
	// Specifies the maximum tolerable delay in seconds where events arriving late could be included. Supported range is `-1` (indefinite) to `1814399` (20d 23h 59m 59s).  Default is `0`.
	EventsLateArrivalMaxDelayInSeconds *int `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// Specifies the maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order. Supported range is `0` to `599` (9m 59s). Default is `5`.
	EventsOutOfOrderMaxDelayInSeconds *int `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Specifies the policy which should be applied to events which arrive out of order in the input event stream. Possible values are `Adjust` and `Drop`.  Default is `Adjust`.
	EventsOutOfOrderPolicy *string `pulumi:"eventsOutOfOrderPolicy"`
	// The Job ID assigned by the Stream Analytics Job.
	JobId *string `pulumi:"jobId"`
	// The Azure Region in which the Resource Group exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size). Possible values are `Drop` and `Stop`.  Default is `Drop`.
	OutputErrorPolicy *string `pulumi:"outputErrorPolicy"`
	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the number of streaming units that the streaming job uses. Supported values are `1`, `3`, `6` and multiples of `6` up to `120`.
	StreamingUnits *int `pulumi:"streamingUnits"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998).
	TransformationQuery *string `pulumi:"transformationQuery"`
}

type JobState struct {
	// Specifies the compatibility level for this job - which controls certain runtime behaviors of the streaming job. Possible values are `1.0` and 1.1`.
	CompatibilityLevel pulumi.StringPtrInput
	// Specifies the Data Locale of the Job, which [should be a supported .NET Culture](https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx).
	DataLocale pulumi.StringPtrInput
	// Specifies the maximum tolerable delay in seconds where events arriving late could be included. Supported range is `-1` (indefinite) to `1814399` (20d 23h 59m 59s).  Default is `0`.
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrInput
	// Specifies the maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order. Supported range is `0` to `599` (9m 59s). Default is `5`.
	EventsOutOfOrderMaxDelayInSeconds pulumi.IntPtrInput
	// Specifies the policy which should be applied to events which arrive out of order in the input event stream. Possible values are `Adjust` and `Drop`.  Default is `Adjust`.
	EventsOutOfOrderPolicy pulumi.StringPtrInput
	// The Job ID assigned by the Stream Analytics Job.
	JobId pulumi.StringPtrInput
	// The Azure Region in which the Resource Group exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size). Possible values are `Drop` and `Stop`.  Default is `Drop`.
	OutputErrorPolicy pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the number of streaming units that the streaming job uses. Supported values are `1`, `3`, `6` and multiples of `6` up to `120`.
	StreamingUnits pulumi.IntPtrInput
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapInput
	// Specifies the query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998).
	TransformationQuery pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// Specifies the compatibility level for this job - which controls certain runtime behaviors of the streaming job. Possible values are `1.0` and 1.1`.
	CompatibilityLevel *string `pulumi:"compatibilityLevel"`
	// Specifies the Data Locale of the Job, which [should be a supported .NET Culture](https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx).
	DataLocale *string `pulumi:"dataLocale"`
	// Specifies the maximum tolerable delay in seconds where events arriving late could be included. Supported range is `-1` (indefinite) to `1814399` (20d 23h 59m 59s).  Default is `0`.
	EventsLateArrivalMaxDelayInSeconds *int `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// Specifies the maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order. Supported range is `0` to `599` (9m 59s). Default is `5`.
	EventsOutOfOrderMaxDelayInSeconds *int `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Specifies the policy which should be applied to events which arrive out of order in the input event stream. Possible values are `Adjust` and `Drop`.  Default is `Adjust`.
	EventsOutOfOrderPolicy *string `pulumi:"eventsOutOfOrderPolicy"`
	// The Azure Region in which the Resource Group exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size). Possible values are `Drop` and `Stop`.  Default is `Drop`.
	OutputErrorPolicy *string `pulumi:"outputErrorPolicy"`
	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the number of streaming units that the streaming job uses. Supported values are `1`, `3`, `6` and multiples of `6` up to `120`.
	StreamingUnits int `pulumi:"streamingUnits"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998).
	TransformationQuery string `pulumi:"transformationQuery"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// Specifies the compatibility level for this job - which controls certain runtime behaviors of the streaming job. Possible values are `1.0` and 1.1`.
	CompatibilityLevel pulumi.StringPtrInput
	// Specifies the Data Locale of the Job, which [should be a supported .NET Culture](https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx).
	DataLocale pulumi.StringPtrInput
	// Specifies the maximum tolerable delay in seconds where events arriving late could be included. Supported range is `-1` (indefinite) to `1814399` (20d 23h 59m 59s).  Default is `0`.
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrInput
	// Specifies the maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order. Supported range is `0` to `599` (9m 59s). Default is `5`.
	EventsOutOfOrderMaxDelayInSeconds pulumi.IntPtrInput
	// Specifies the policy which should be applied to events which arrive out of order in the input event stream. Possible values are `Adjust` and `Drop`.  Default is `Adjust`.
	EventsOutOfOrderPolicy pulumi.StringPtrInput
	// The Azure Region in which the Resource Group exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size). Possible values are `Drop` and `Stop`.  Default is `Drop`.
	OutputErrorPolicy pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the number of streaming units that the streaming job uses. Supported values are `1`, `3`, `6` and multiples of `6` up to `120`.
	StreamingUnits pulumi.IntInput
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapInput
	// Specifies the query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998).
	TransformationQuery pulumi.StringInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

