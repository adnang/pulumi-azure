// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package scheduler

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Scheduler Job Collection.
// 
// > **NOTE:** Support for Scheduler Job Collections has been deprecated by Microsoft in favour of Logic Apps ([more information can be found at this link](https://docs.microsoft.com/en-us/azure/scheduler/migrate-from-scheduler-to-logic-apps)) - as such we plan to remove support for this data source as a part of version 2.0 of the AzureRM Provider.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/scheduler_job_collection.html.markdown.
func LookupJobCollection(ctx *pulumi.Context, args *LookupJobCollectionArgs, opts ...pulumi.InvokeOption) (*LookupJobCollectionResult, error) {
	var rv LookupJobCollectionResult
	err := ctx.Invoke("azure:scheduler/getJobCollection:getJobCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getJobCollection.
type LookupJobCollectionArgs struct {
	// Specifies the name of the Scheduler Job Collection.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group in which the Scheduler Job Collection resides.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getJobCollection.
type LookupJobCollectionResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure location where the resource exists.
	Location string `pulumi:"location"`
	Name string `pulumi:"name"`
	// The Job collection quotas as documented in the `quota` block below.
	Quotas []GetJobCollectionQuota `pulumi:"quotas"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Job Collection's pricing level's SKU.
	Sku string `pulumi:"sku"`
	// The Job Collection's state.
	State string `pulumi:"state"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

