// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access the properties of an Azure CosmosDB (formally DocumentDB) Account.
func LookupAccount(ctx *pulumi.Context, args *GetAccountArgs) (*GetAccountResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	outputs, err := ctx.Invoke("azure:cosmosdb/getAccount:getAccount", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAccountResult{
		Capabilities: outputs["capabilities"],
		ConsistencyPolicies: outputs["consistencyPolicies"],
		EnableAutomaticFailover: outputs["enableAutomaticFailover"],
		Endpoint: outputs["endpoint"],
		GeoLocations: outputs["geoLocations"],
		IpRangeFilter: outputs["ipRangeFilter"],
		Kind: outputs["kind"],
		Location: outputs["location"],
		OfferType: outputs["offerType"],
		PrimaryMasterKey: outputs["primaryMasterKey"],
		PrimaryReadonlyMasterKey: outputs["primaryReadonlyMasterKey"],
		ReadEndpoints: outputs["readEndpoints"],
		SecondaryMasterKey: outputs["secondaryMasterKey"],
		SecondaryReadonlyMasterKey: outputs["secondaryReadonlyMasterKey"],
		Tags: outputs["tags"],
		WriteEndpoints: outputs["writeEndpoints"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAccount.
type GetAccountArgs struct {
	// Specifies the name of the CosmosDB Account. 
	Name interface{}
	// Specifies the name of the resource group in which the CosmosDB Account resides. 
	ResourceGroupName interface{}
}

// A collection of values returned by getAccount.
type GetAccountResult struct {
	// Capabilities enabled on this Cosmos DB account.
	Capabilities interface{}
	ConsistencyPolicies interface{}
	// If automatic failover is enabled for this CosmosDB Account.
	EnableAutomaticFailover interface{}
	// The endpoint used to connect to the CosmosDB account.
	Endpoint interface{}
	GeoLocations interface{}
	// The current IP Filter for this CosmosDB account
	IpRangeFilter interface{}
	// The Kind of the CosmosDB account.
	Kind interface{}
	// The name of the Azure region hosting replicated data.
	Location interface{}
	// The Offer Type to used by this CosmosDB Account.
	OfferType interface{}
	// The Primary master key for the CosmosDB Account.
	PrimaryMasterKey interface{}
	// The Primary read-only master Key for the CosmosDB Account.
	PrimaryReadonlyMasterKey interface{}
	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints interface{}
	// The Secondary master key for the CosmosDB Account.
	SecondaryMasterKey interface{}
	// The Secondary read-only master key for the CosmosDB Account.
	SecondaryReadonlyMasterKey interface{}
	// A mapping of tags assigned to the resource.
	Tags interface{}
	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
