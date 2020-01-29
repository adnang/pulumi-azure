// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cosmosdb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing CosmosDB (formally DocumentDB) Account.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/cosmosdb_account.html.markdown.
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure:cosmosdb/getAccount:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccount.
type LookupAccountArgs struct {
	// Specifies the name of the CosmosDB Account.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group in which the CosmosDB Account resides.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getAccount.
type LookupAccountResult struct {
	// Capabilities enabled on this Cosmos DB account.
	Capabilities []GetAccountCapability `pulumi:"capabilities"`
	ConsistencyPolicies []GetAccountConsistencyPolicy `pulumi:"consistencyPolicies"`
	// If automatic failover is enabled for this CosmosDB Account.
	EnableAutomaticFailover bool `pulumi:"enableAutomaticFailover"`
	// If multi-master is enabled for this Cosmos DB account.
	EnableMultipleWriteLocations bool `pulumi:"enableMultipleWriteLocations"`
	// The endpoint used to connect to the CosmosDB account.
	Endpoint string `pulumi:"endpoint"`
	GeoLocations []GetAccountGeoLocation `pulumi:"geoLocations"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current IP Filter for this CosmosDB account
	IpRangeFilter string `pulumi:"ipRangeFilter"`
	// If virtual network filtering is enabled for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled bool `pulumi:"isVirtualNetworkFilterEnabled"`
	// The Kind of the CosmosDB account.
	Kind string `pulumi:"kind"`
	// The name of the Azure region hosting replicated data.
	Location string `pulumi:"location"`
	Name string `pulumi:"name"`
	// The Offer Type to used by this CosmosDB Account.
	OfferType string `pulumi:"offerType"`
	// The Primary master key for the CosmosDB Account.
	PrimaryMasterKey string `pulumi:"primaryMasterKey"`
	// The Primary read-only master Key for the CosmosDB Account.
	PrimaryReadonlyMasterKey string `pulumi:"primaryReadonlyMasterKey"`
	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints []string `pulumi:"readEndpoints"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Secondary master key for the CosmosDB Account.
	SecondaryMasterKey string `pulumi:"secondaryMasterKey"`
	// The Secondary read-only master key for the CosmosDB Account.
	SecondaryReadonlyMasterKey string `pulumi:"secondaryReadonlyMasterKey"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Subnets that are allowed to access this CosmosDB account.
	VirtualNetworkRules []GetAccountVirtualNetworkRule `pulumi:"virtualNetworkRules"`
	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints []string `pulumi:"writeEndpoints"`
}

