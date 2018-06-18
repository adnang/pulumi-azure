// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scheduler

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access the properties of an Azure scheduler job collection.
func LookupJobCollection(ctx *pulumi.Context, args *GetJobCollectionArgs) (*GetJobCollectionResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	outputs, err := ctx.Invoke("azure:scheduler/getJobCollection:getJobCollection", inputs)
	if err != nil {
		return nil, err
	}
	return &GetJobCollectionResult{
		Location: outputs["location"],
		Quotas: outputs["quotas"],
		Sku: outputs["sku"],
		State: outputs["state"],
		Tags: outputs["tags"],
	}, nil
}

// A collection of arguments for invoking getJobCollection.
type GetJobCollectionArgs struct {
	// Specifies the name of the Scheduler Job Collection. 
	Name interface{}
	// Specifies the name of the resource group in which the Scheduler Job Collection resides. 
	ResourceGroupName interface{}
}

// A collection of values returned by getJobCollection.
type GetJobCollectionResult struct {
	// The Azure location where the resource exists. 
	Location interface{}
	// The Job collection quotas as documented in the `quota` block below. 
	Quotas interface{}
	// The Job Collection's pricing level's SKU. 
	Sku interface{}
	// The Job Collection's state. 
	State interface{}
	// A mapping of tags assigned to the resource.
	Tags interface{}
}