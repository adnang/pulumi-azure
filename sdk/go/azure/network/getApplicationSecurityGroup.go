// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Application Security Group.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/application_security_group.html.markdown.
func LookupApplicationSecurityGroup(ctx *pulumi.Context, args *LookupApplicationSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupApplicationSecurityGroupResult, error) {
	var rv LookupApplicationSecurityGroupResult
	err := ctx.Invoke("azure:network/getApplicationSecurityGroup:getApplicationSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationSecurityGroup.
type LookupApplicationSecurityGroupArgs struct {
	// The name of the Application Security Group.
	Name string `pulumi:"name"`
	// The name of the resource group in which the Application Security Group exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getApplicationSecurityGroup.
type LookupApplicationSecurityGroupResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The supported Azure location where the Application Security Group exists.
	Location string `pulumi:"location"`
	Name string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

