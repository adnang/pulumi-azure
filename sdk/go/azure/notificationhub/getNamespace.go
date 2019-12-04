// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notificationhub

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Notification Hub Namespace.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/notification_hub_namespace.html.markdown.
func LookupNamespace(ctx *pulumi.Context, args *GetNamespaceArgs, opts ...pulumi.InvokeOption) (*GetNamespaceResult, error) {
	var rv GetNamespaceResult
	err := ctx.Invoke("azure:notificationhub/getNamespace:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNamespace.
type GetNamespaceArgs struct {
	// Specifies the Name of the Notification Hub Namespace.
	Name string `pulumi:"name"`
	// Specifies the Name of the Resource Group within which the Notification Hub exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getNamespace.
type GetNamespaceResult struct {
	// Is this Notification Hub Namespace enabled?
	Enabled bool `pulumi:"enabled"`
	// The Azure Region in which this Notification Hub Namespace exists.
	Location string `pulumi:"location"`
	// (Required) The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard.`
	Name string `pulumi:"name"`
	// The Type of Namespace, such as `Messaging` or `NotificationHub`.
	NamespaceType string `pulumi:"namespaceType"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServicebusEndpoint string `pulumi:"servicebusEndpoint"`
	// A `sku` block as defined below.
	Sku GetNamespaceSkuResult `pulumi:"sku"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
type GetNamespaceSkuResult struct {
	// Specifies the Name of the Notification Hub Namespace.
	Name string `pulumi:"name"`
}
