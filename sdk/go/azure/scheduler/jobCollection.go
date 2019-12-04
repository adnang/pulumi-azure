// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scheduler

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Scheduler Job Collection.
// 
// > **NOTE:** Support for Scheduler Job Collections has been deprecated by Microsoft in favour of Logic Apps ([more information can be found at this link](https://docs.microsoft.com/en-us/azure/scheduler/migrate-from-scheduler-to-logic-apps)) - as such we plan to remove support for this resource as a part of version 2.0 of the AzureRM Provider.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/scheduler_job_collection.html.markdown.
type JobCollection struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// Specifies the name of the Scheduler Job Collection. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// Configures the Job collection quotas as documented in the `quota` block below. 
	Quota JobCollectionQuotaOutput `pulumi:"quota"`

	// The name of the resource group in which to create the Scheduler Job Collection. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// Sets the Job Collection's pricing level's SKU. Possible values include: `Standard`, `Free`, `P10Premium`, `P20Premium`.
	Sku pulumi.StringOutput `pulumi:"sku"`

	// Sets Job Collection's state. Possible values include: `Enabled`, `Disabled`, `Suspended`.
	State pulumi.StringOutput `pulumi:"state"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewJobCollection registers a new resource with the given unique name, arguments, and options.
func NewJobCollection(ctx *pulumi.Context,
	name string, args *JobCollectionArgs, opts ...pulumi.ResourceOption) (*JobCollection, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Quota; i != nil { inputs["quota"] = i.ToJobCollectionQuotaOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Sku; i != nil { inputs["sku"] = i.ToStringOutput() }
		if i := args.State; i != nil { inputs["state"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource JobCollection
	err := ctx.RegisterResource("azure:scheduler/jobCollection:JobCollection", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobCollection gets an existing JobCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobCollectionState, opts ...pulumi.ResourceOption) (*JobCollection, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Quota; i != nil { inputs["quota"] = i.ToJobCollectionQuotaOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Sku; i != nil { inputs["sku"] = i.ToStringOutput() }
		if i := state.State; i != nil { inputs["state"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource JobCollection
	err := ctx.ReadResource("azure:scheduler/jobCollection:JobCollection", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobCollection resources.
type JobCollectionState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name of the Scheduler Job Collection. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Configures the Job collection quotas as documented in the `quota` block below. 
	Quota JobCollectionQuotaInput `pulumi:"quota"`
	// The name of the resource group in which to create the Scheduler Job Collection. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Sets the Job Collection's pricing level's SKU. Possible values include: `Standard`, `Free`, `P10Premium`, `P20Premium`.
	Sku pulumi.StringInput `pulumi:"sku"`
	// Sets Job Collection's state. Possible values include: `Enabled`, `Disabled`, `Suspended`.
	State pulumi.StringInput `pulumi:"state"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a JobCollection resource.
type JobCollectionArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name of the Scheduler Job Collection. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Configures the Job collection quotas as documented in the `quota` block below. 
	Quota JobCollectionQuotaInput `pulumi:"quota"`
	// The name of the resource group in which to create the Scheduler Job Collection. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Sets the Job Collection's pricing level's SKU. Possible values include: `Standard`, `Free`, `P10Premium`, `P20Premium`.
	Sku pulumi.StringInput `pulumi:"sku"`
	// Sets Job Collection's state. Possible values include: `Enabled`, `Disabled`, `Suspended`.
	State pulumi.StringInput `pulumi:"state"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
type JobCollectionQuota struct {
	MaxJobCount *int `pulumi:"maxJobCount"`
	MaxRecurrenceFrequency string `pulumi:"maxRecurrenceFrequency"`
	MaxRecurrenceInterval *int `pulumi:"maxRecurrenceInterval"`
	MaxRetryInterval *int `pulumi:"maxRetryInterval"`
}
var jobCollectionQuotaType = reflect.TypeOf((*JobCollectionQuota)(nil)).Elem()

type JobCollectionQuotaInput interface {
	pulumi.Input

	ToJobCollectionQuotaOutput() JobCollectionQuotaOutput
	ToJobCollectionQuotaOutputWithContext(ctx context.Context) JobCollectionQuotaOutput
}

type JobCollectionQuotaArgs struct {
	MaxJobCount pulumi.IntInput `pulumi:"maxJobCount"`
	MaxRecurrenceFrequency pulumi.StringInput `pulumi:"maxRecurrenceFrequency"`
	MaxRecurrenceInterval pulumi.IntInput `pulumi:"maxRecurrenceInterval"`
	MaxRetryInterval pulumi.IntInput `pulumi:"maxRetryInterval"`
}

func (JobCollectionQuotaArgs) ElementType() reflect.Type {
	return jobCollectionQuotaType
}

func (a JobCollectionQuotaArgs) ToJobCollectionQuotaOutput() JobCollectionQuotaOutput {
	return pulumi.ToOutput(a).(JobCollectionQuotaOutput)
}

func (a JobCollectionQuotaArgs) ToJobCollectionQuotaOutputWithContext(ctx context.Context) JobCollectionQuotaOutput {
	return pulumi.ToOutputWithContext(ctx, a).(JobCollectionQuotaOutput)
}

type JobCollectionQuotaOutput struct { *pulumi.OutputState }

func (o JobCollectionQuotaOutput) MaxJobCount() pulumi.IntOutput {
	return o.Apply(func(v JobCollectionQuota) int {
		if v.MaxJobCount == nil { return *new(int) } else { return *v.MaxJobCount }
	}).(pulumi.IntOutput)
}

func (o JobCollectionQuotaOutput) MaxRecurrenceFrequency() pulumi.StringOutput {
	return o.Apply(func(v JobCollectionQuota) string {
		return v.MaxRecurrenceFrequency
	}).(pulumi.StringOutput)
}

func (o JobCollectionQuotaOutput) MaxRecurrenceInterval() pulumi.IntOutput {
	return o.Apply(func(v JobCollectionQuota) int {
		if v.MaxRecurrenceInterval == nil { return *new(int) } else { return *v.MaxRecurrenceInterval }
	}).(pulumi.IntOutput)
}

func (o JobCollectionQuotaOutput) MaxRetryInterval() pulumi.IntOutput {
	return o.Apply(func(v JobCollectionQuota) int {
		if v.MaxRetryInterval == nil { return *new(int) } else { return *v.MaxRetryInterval }
	}).(pulumi.IntOutput)
}

func (JobCollectionQuotaOutput) ElementType() reflect.Type {
	return jobCollectionQuotaType
}

func (o JobCollectionQuotaOutput) ToJobCollectionQuotaOutput() JobCollectionQuotaOutput {
	return o
}

func (o JobCollectionQuotaOutput) ToJobCollectionQuotaOutputWithContext(ctx context.Context) JobCollectionQuotaOutput {
	return o
}

func init() { pulumi.RegisterOutputType(JobCollectionQuotaOutput{}) }

