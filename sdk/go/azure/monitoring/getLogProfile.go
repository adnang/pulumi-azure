// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access the properties of a Log Profile.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/monitor_log_profile.html.markdown.
func LookupLogProfile(ctx *pulumi.Context, args *GetLogProfileArgs, opts ...pulumi.InvokeOption) (*GetLogProfileResult, error) {
	var rv GetLogProfileResult
	err := ctx.Invoke("azure:monitoring/getLogProfile:getLogProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogProfile.
type GetLogProfileArgs struct {
	// Specifies the Name of the Log Profile.
	Name string `pulumi:"name"`
}

// A collection of values returned by getLogProfile.
type GetLogProfileResult struct {
	// List of categories of the logs.
	Categories []string `pulumi:"categories"`
	// List of regions for which Activity Log events are stored or streamed.
	Locations []string `pulumi:"locations"`
	Name string `pulumi:"name"`
	RetentionPolicy GetLogProfileRetentionPolicyResult `pulumi:"retentionPolicy"`
	// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to.
	ServicebusRuleId string `pulumi:"servicebusRuleId"`
	// The resource id of the storage account in which the Activity Log is stored.
	StorageAccountId string `pulumi:"storageAccountId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
type GetLogProfileRetentionPolicyResult struct {
	// The number of days for the retention policy.
	Days int `pulumi:"days"`
	// A boolean value indicating whether the retention policy is enabled.
	Enabled bool `pulumi:"enabled"`
}
