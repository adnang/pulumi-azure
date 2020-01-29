// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package core

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing User Assigned Identity.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/user_assigned_identity_legacy.html.markdown.
func GetUserAssignedIdentity(ctx *pulumi.Context, args *GetUserAssignedIdentityArgs, opts ...pulumi.InvokeOption) (*GetUserAssignedIdentityResult, error) {
	var rv GetUserAssignedIdentityResult
	err := ctx.Invoke("azure:core/getUserAssignedIdentity:getUserAssignedIdentity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserAssignedIdentity.
type GetUserAssignedIdentityArgs struct {
	// The name of the User Assigned Identity.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the User Assigned Identity exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getUserAssignedIdentity.
type GetUserAssignedIdentityResult struct {
	// The Client ID of the User Assigned Identity.
	ClientId string `pulumi:"clientId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure location where the User Assigned Identity exists.
	Location string `pulumi:"location"`
	Name string `pulumi:"name"`
	// The Service Principal ID of the User Assigned Identity.
	PrincipalId string `pulumi:"principalId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the User Assigned Identity.
	Tags map[string]string `pulumi:"tags"`
}

