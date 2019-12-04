// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a CosmosDB (formally DocumentDB) Account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/cosmosdb_account.html.markdown.
type Account struct {
	pulumi.CustomResourceState

	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Capabilities AccountCapabilitiesArrayOutput `pulumi:"capabilities"`

	// A list of connection strings available for this CosmosDB account. If the kind is `GlobalDocumentDB`, this will be empty.
	ConnectionStrings pulumi.StringArrayOutput `pulumi:"connectionStrings"`

	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy AccountConsistencyPolicyOutput `pulumi:"consistencyPolicy"`

	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover pulumi.BoolOutput `pulumi:"enableAutomaticFailover"`

	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations pulumi.BoolOutput `pulumi:"enableMultipleWriteLocations"`

	// The endpoint used to connect to the CosmosDB account.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`

	FailoverPolicies AccountFailoverPoliciesArrayOutput `pulumi:"failoverPolicies"`

	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations AccountGeoLocationsArrayOutput `pulumi:"geoLocations"`

	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter pulumi.StringOutput `pulumi:"ipRangeFilter"`

	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled pulumi.BoolOutput `pulumi:"isVirtualNetworkFilterEnabled"`

	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind pulumi.StringOutput `pulumi:"kind"`

	// The name of the Azure region to host replicated data.
	Location pulumi.StringOutput `pulumi:"location"`

	// The capability to enable - Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType pulumi.StringOutput `pulumi:"offerType"`

	// The Primary master key for the CosmosDB Account.
	PrimaryMasterKey pulumi.StringOutput `pulumi:"primaryMasterKey"`

	// The Primary read-only master Key for the CosmosDB Account.
	PrimaryReadonlyMasterKey pulumi.StringOutput `pulumi:"primaryReadonlyMasterKey"`

	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints pulumi.StringArrayOutput `pulumi:"readEndpoints"`

	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The Secondary master key for the CosmosDB Account.
	SecondaryMasterKey pulumi.StringOutput `pulumi:"secondaryMasterKey"`

	// The Secondary read-only master key for the CosmosDB Account.
	SecondaryReadonlyMasterKey pulumi.StringOutput `pulumi:"secondaryReadonlyMasterKey"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules AccountVirtualNetworkRulesArrayOutput `pulumi:"virtualNetworkRules"`

	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints pulumi.StringArrayOutput `pulumi:"writeEndpoints"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil || args.ConsistencyPolicy == nil {
		return nil, errors.New("missing required argument 'ConsistencyPolicy'")
	}
	if args == nil || args.OfferType == nil {
		return nil, errors.New("missing required argument 'OfferType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Capabilities; i != nil { inputs["capabilities"] = i.ToAccountCapabilitiesArrayOutput() }
		if i := args.ConsistencyPolicy; i != nil { inputs["consistencyPolicy"] = i.ToAccountConsistencyPolicyOutput() }
		if i := args.EnableAutomaticFailover; i != nil { inputs["enableAutomaticFailover"] = i.ToBoolOutput() }
		if i := args.EnableMultipleWriteLocations; i != nil { inputs["enableMultipleWriteLocations"] = i.ToBoolOutput() }
		if i := args.FailoverPolicies; i != nil { inputs["failoverPolicies"] = i.ToAccountFailoverPoliciesArrayOutput() }
		if i := args.GeoLocations; i != nil { inputs["geoLocations"] = i.ToAccountGeoLocationsArrayOutput() }
		if i := args.IpRangeFilter; i != nil { inputs["ipRangeFilter"] = i.ToStringOutput() }
		if i := args.IsVirtualNetworkFilterEnabled; i != nil { inputs["isVirtualNetworkFilterEnabled"] = i.ToBoolOutput() }
		if i := args.Kind; i != nil { inputs["kind"] = i.ToStringOutput() }
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.OfferType; i != nil { inputs["offerType"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.VirtualNetworkRules; i != nil { inputs["virtualNetworkRules"] = i.ToAccountVirtualNetworkRulesArrayOutput() }
	}
	var resource Account
	err := ctx.RegisterResource("azure:cosmosdb/account:Account", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Capabilities; i != nil { inputs["capabilities"] = i.ToAccountCapabilitiesArrayOutput() }
		if i := state.ConnectionStrings; i != nil { inputs["connectionStrings"] = i.ToStringArrayOutput() }
		if i := state.ConsistencyPolicy; i != nil { inputs["consistencyPolicy"] = i.ToAccountConsistencyPolicyOutput() }
		if i := state.EnableAutomaticFailover; i != nil { inputs["enableAutomaticFailover"] = i.ToBoolOutput() }
		if i := state.EnableMultipleWriteLocations; i != nil { inputs["enableMultipleWriteLocations"] = i.ToBoolOutput() }
		if i := state.Endpoint; i != nil { inputs["endpoint"] = i.ToStringOutput() }
		if i := state.FailoverPolicies; i != nil { inputs["failoverPolicies"] = i.ToAccountFailoverPoliciesArrayOutput() }
		if i := state.GeoLocations; i != nil { inputs["geoLocations"] = i.ToAccountGeoLocationsArrayOutput() }
		if i := state.IpRangeFilter; i != nil { inputs["ipRangeFilter"] = i.ToStringOutput() }
		if i := state.IsVirtualNetworkFilterEnabled; i != nil { inputs["isVirtualNetworkFilterEnabled"] = i.ToBoolOutput() }
		if i := state.Kind; i != nil { inputs["kind"] = i.ToStringOutput() }
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.OfferType; i != nil { inputs["offerType"] = i.ToStringOutput() }
		if i := state.PrimaryMasterKey; i != nil { inputs["primaryMasterKey"] = i.ToStringOutput() }
		if i := state.PrimaryReadonlyMasterKey; i != nil { inputs["primaryReadonlyMasterKey"] = i.ToStringOutput() }
		if i := state.ReadEndpoints; i != nil { inputs["readEndpoints"] = i.ToStringArrayOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.SecondaryMasterKey; i != nil { inputs["secondaryMasterKey"] = i.ToStringOutput() }
		if i := state.SecondaryReadonlyMasterKey; i != nil { inputs["secondaryReadonlyMasterKey"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.VirtualNetworkRules; i != nil { inputs["virtualNetworkRules"] = i.ToAccountVirtualNetworkRulesArrayOutput() }
		if i := state.WriteEndpoints; i != nil { inputs["writeEndpoints"] = i.ToStringArrayOutput() }
	}
	var resource Account
	err := ctx.ReadResource("azure:cosmosdb/account:Account", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type AccountState struct {
	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Capabilities AccountCapabilitiesArrayInput `pulumi:"capabilities"`
	// A list of connection strings available for this CosmosDB account. If the kind is `GlobalDocumentDB`, this will be empty.
	ConnectionStrings pulumi.StringArrayInput `pulumi:"connectionStrings"`
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy AccountConsistencyPolicyInput `pulumi:"consistencyPolicy"`
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover pulumi.BoolInput `pulumi:"enableAutomaticFailover"`
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations pulumi.BoolInput `pulumi:"enableMultipleWriteLocations"`
	// The endpoint used to connect to the CosmosDB account.
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	FailoverPolicies AccountFailoverPoliciesArrayInput `pulumi:"failoverPolicies"`
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations AccountGeoLocationsArrayInput `pulumi:"geoLocations"`
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter pulumi.StringInput `pulumi:"ipRangeFilter"`
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled pulumi.BoolInput `pulumi:"isVirtualNetworkFilterEnabled"`
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind pulumi.StringInput `pulumi:"kind"`
	// The name of the Azure region to host replicated data.
	Location pulumi.StringInput `pulumi:"location"`
	// The capability to enable - Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType pulumi.StringInput `pulumi:"offerType"`
	// The Primary master key for the CosmosDB Account.
	PrimaryMasterKey pulumi.StringInput `pulumi:"primaryMasterKey"`
	// The Primary read-only master Key for the CosmosDB Account.
	PrimaryReadonlyMasterKey pulumi.StringInput `pulumi:"primaryReadonlyMasterKey"`
	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints pulumi.StringArrayInput `pulumi:"readEndpoints"`
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Secondary master key for the CosmosDB Account.
	SecondaryMasterKey pulumi.StringInput `pulumi:"secondaryMasterKey"`
	// The Secondary read-only master key for the CosmosDB Account.
	SecondaryReadonlyMasterKey pulumi.StringInput `pulumi:"secondaryReadonlyMasterKey"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules AccountVirtualNetworkRulesArrayInput `pulumi:"virtualNetworkRules"`
	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints pulumi.StringArrayInput `pulumi:"writeEndpoints"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Capabilities AccountCapabilitiesArrayInput `pulumi:"capabilities"`
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy AccountConsistencyPolicyInput `pulumi:"consistencyPolicy"`
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover pulumi.BoolInput `pulumi:"enableAutomaticFailover"`
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations pulumi.BoolInput `pulumi:"enableMultipleWriteLocations"`
	FailoverPolicies AccountFailoverPoliciesArrayInput `pulumi:"failoverPolicies"`
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations AccountGeoLocationsArrayInput `pulumi:"geoLocations"`
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter pulumi.StringInput `pulumi:"ipRangeFilter"`
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled pulumi.BoolInput `pulumi:"isVirtualNetworkFilterEnabled"`
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind pulumi.StringInput `pulumi:"kind"`
	// The name of the Azure region to host replicated data.
	Location pulumi.StringInput `pulumi:"location"`
	// The capability to enable - Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType pulumi.StringInput `pulumi:"offerType"`
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules AccountVirtualNetworkRulesArrayInput `pulumi:"virtualNetworkRules"`
}
type AccountCapabilities struct {
	// The capability to enable - Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Name string `pulumi:"name"`
}
var accountCapabilitiesType = reflect.TypeOf((*AccountCapabilities)(nil)).Elem()

type AccountCapabilitiesInput interface {
	pulumi.Input

	ToAccountCapabilitiesOutput() AccountCapabilitiesOutput
	ToAccountCapabilitiesOutputWithContext(ctx context.Context) AccountCapabilitiesOutput
}

type AccountCapabilitiesArgs struct {
	// The capability to enable - Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Name pulumi.StringInput `pulumi:"name"`
}

func (AccountCapabilitiesArgs) ElementType() reflect.Type {
	return accountCapabilitiesType
}

func (a AccountCapabilitiesArgs) ToAccountCapabilitiesOutput() AccountCapabilitiesOutput {
	return pulumi.ToOutput(a).(AccountCapabilitiesOutput)
}

func (a AccountCapabilitiesArgs) ToAccountCapabilitiesOutputWithContext(ctx context.Context) AccountCapabilitiesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AccountCapabilitiesOutput)
}

type AccountCapabilitiesOutput struct { *pulumi.OutputState }

// The capability to enable - Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
func (o AccountCapabilitiesOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v AccountCapabilities) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (AccountCapabilitiesOutput) ElementType() reflect.Type {
	return accountCapabilitiesType
}

func (o AccountCapabilitiesOutput) ToAccountCapabilitiesOutput() AccountCapabilitiesOutput {
	return o
}

func (o AccountCapabilitiesOutput) ToAccountCapabilitiesOutputWithContext(ctx context.Context) AccountCapabilitiesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AccountCapabilitiesOutput{}) }

var accountCapabilitiesArrayType = reflect.TypeOf((*[]AccountCapabilities)(nil)).Elem()

type AccountCapabilitiesArrayInput interface {
	pulumi.Input

	ToAccountCapabilitiesArrayOutput() AccountCapabilitiesArrayOutput
	ToAccountCapabilitiesArrayOutputWithContext(ctx context.Context) AccountCapabilitiesArrayOutput
}

type AccountCapabilitiesArrayArgs []AccountCapabilitiesInput

func (AccountCapabilitiesArrayArgs) ElementType() reflect.Type {
	return accountCapabilitiesArrayType
}

func (a AccountCapabilitiesArrayArgs) ToAccountCapabilitiesArrayOutput() AccountCapabilitiesArrayOutput {
	return pulumi.ToOutput(a).(AccountCapabilitiesArrayOutput)
}

func (a AccountCapabilitiesArrayArgs) ToAccountCapabilitiesArrayOutputWithContext(ctx context.Context) AccountCapabilitiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AccountCapabilitiesArrayOutput)
}

type AccountCapabilitiesArrayOutput struct { *pulumi.OutputState }

func (o AccountCapabilitiesArrayOutput) Index(i pulumi.IntInput) AccountCapabilitiesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) AccountCapabilities {
		return vs[0].([]AccountCapabilities)[vs[1].(int)]
	}).(AccountCapabilitiesOutput)
}

func (AccountCapabilitiesArrayOutput) ElementType() reflect.Type {
	return accountCapabilitiesArrayType
}

func (o AccountCapabilitiesArrayOutput) ToAccountCapabilitiesArrayOutput() AccountCapabilitiesArrayOutput {
	return o
}

func (o AccountCapabilitiesArrayOutput) ToAccountCapabilitiesArrayOutputWithContext(ctx context.Context) AccountCapabilitiesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AccountCapabilitiesArrayOutput{}) }

type AccountConsistencyPolicy struct {
	// The Consistency Level to use for this CosmosDB Account - can be either `BoundedStaleness`, `Eventual`, `Session`, `Strong` or `ConsistentPrefix`.
	ConsistencyLevel string `pulumi:"consistencyLevel"`
	// When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is `5` - `86400` (1 day). Defaults to `5`. Required when `consistencyLevel` is set to `BoundedStaleness`.
	MaxIntervalInSeconds *int `pulumi:"maxIntervalInSeconds"`
	// When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is `10` – `2147483647`. Defaults to `100`. Required when `consistencyLevel` is set to `BoundedStaleness`.
	MaxStalenessPrefix *int `pulumi:"maxStalenessPrefix"`
}
var accountConsistencyPolicyType = reflect.TypeOf((*AccountConsistencyPolicy)(nil)).Elem()

type AccountConsistencyPolicyInput interface {
	pulumi.Input

	ToAccountConsistencyPolicyOutput() AccountConsistencyPolicyOutput
	ToAccountConsistencyPolicyOutputWithContext(ctx context.Context) AccountConsistencyPolicyOutput
}

type AccountConsistencyPolicyArgs struct {
	// The Consistency Level to use for this CosmosDB Account - can be either `BoundedStaleness`, `Eventual`, `Session`, `Strong` or `ConsistentPrefix`.
	ConsistencyLevel pulumi.StringInput `pulumi:"consistencyLevel"`
	// When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is `5` - `86400` (1 day). Defaults to `5`. Required when `consistencyLevel` is set to `BoundedStaleness`.
	MaxIntervalInSeconds pulumi.IntInput `pulumi:"maxIntervalInSeconds"`
	// When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is `10` – `2147483647`. Defaults to `100`. Required when `consistencyLevel` is set to `BoundedStaleness`.
	MaxStalenessPrefix pulumi.IntInput `pulumi:"maxStalenessPrefix"`
}

func (AccountConsistencyPolicyArgs) ElementType() reflect.Type {
	return accountConsistencyPolicyType
}

func (a AccountConsistencyPolicyArgs) ToAccountConsistencyPolicyOutput() AccountConsistencyPolicyOutput {
	return pulumi.ToOutput(a).(AccountConsistencyPolicyOutput)
}

func (a AccountConsistencyPolicyArgs) ToAccountConsistencyPolicyOutputWithContext(ctx context.Context) AccountConsistencyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AccountConsistencyPolicyOutput)
}

type AccountConsistencyPolicyOutput struct { *pulumi.OutputState }

// The Consistency Level to use for this CosmosDB Account - can be either `BoundedStaleness`, `Eventual`, `Session`, `Strong` or `ConsistentPrefix`.
func (o AccountConsistencyPolicyOutput) ConsistencyLevel() pulumi.StringOutput {
	return o.Apply(func(v AccountConsistencyPolicy) string {
		return v.ConsistencyLevel
	}).(pulumi.StringOutput)
}

// When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is `5` - `86400` (1 day). Defaults to `5`. Required when `consistencyLevel` is set to `BoundedStaleness`.
func (o AccountConsistencyPolicyOutput) MaxIntervalInSeconds() pulumi.IntOutput {
	return o.Apply(func(v AccountConsistencyPolicy) int {
		if v.MaxIntervalInSeconds == nil { return *new(int) } else { return *v.MaxIntervalInSeconds }
	}).(pulumi.IntOutput)
}

// When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is `10` – `2147483647`. Defaults to `100`. Required when `consistencyLevel` is set to `BoundedStaleness`.
func (o AccountConsistencyPolicyOutput) MaxStalenessPrefix() pulumi.IntOutput {
	return o.Apply(func(v AccountConsistencyPolicy) int {
		if v.MaxStalenessPrefix == nil { return *new(int) } else { return *v.MaxStalenessPrefix }
	}).(pulumi.IntOutput)
}

func (AccountConsistencyPolicyOutput) ElementType() reflect.Type {
	return accountConsistencyPolicyType
}

func (o AccountConsistencyPolicyOutput) ToAccountConsistencyPolicyOutput() AccountConsistencyPolicyOutput {
	return o
}

func (o AccountConsistencyPolicyOutput) ToAccountConsistencyPolicyOutputWithContext(ctx context.Context) AccountConsistencyPolicyOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AccountConsistencyPolicyOutput{}) }

type AccountFailoverPolicies struct {
	// The ID of the virtual network subnet.
	Id *string `pulumi:"id"`
	// The name of the Azure region to host replicated data.
	Location string `pulumi:"location"`
	Priority int `pulumi:"priority"`
}
var accountFailoverPoliciesType = reflect.TypeOf((*AccountFailoverPolicies)(nil)).Elem()

type AccountFailoverPoliciesInput interface {
	pulumi.Input

	ToAccountFailoverPoliciesOutput() AccountFailoverPoliciesOutput
	ToAccountFailoverPoliciesOutputWithContext(ctx context.Context) AccountFailoverPoliciesOutput
}

type AccountFailoverPoliciesArgs struct {
	// The ID of the virtual network subnet.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the Azure region to host replicated data.
	Location pulumi.StringInput `pulumi:"location"`
	Priority pulumi.IntInput `pulumi:"priority"`
}

func (AccountFailoverPoliciesArgs) ElementType() reflect.Type {
	return accountFailoverPoliciesType
}

func (a AccountFailoverPoliciesArgs) ToAccountFailoverPoliciesOutput() AccountFailoverPoliciesOutput {
	return pulumi.ToOutput(a).(AccountFailoverPoliciesOutput)
}

func (a AccountFailoverPoliciesArgs) ToAccountFailoverPoliciesOutputWithContext(ctx context.Context) AccountFailoverPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AccountFailoverPoliciesOutput)
}

type AccountFailoverPoliciesOutput struct { *pulumi.OutputState }

// The ID of the virtual network subnet.
func (o AccountFailoverPoliciesOutput) Id() pulumi.StringOutput {
	return o.Apply(func(v AccountFailoverPolicies) string {
		if v.Id == nil { return *new(string) } else { return *v.Id }
	}).(pulumi.StringOutput)
}

// The name of the Azure region to host replicated data.
func (o AccountFailoverPoliciesOutput) Location() pulumi.StringOutput {
	return o.Apply(func(v AccountFailoverPolicies) string {
		return v.Location
	}).(pulumi.StringOutput)
}

func (o AccountFailoverPoliciesOutput) Priority() pulumi.IntOutput {
	return o.Apply(func(v AccountFailoverPolicies) int {
		return v.Priority
	}).(pulumi.IntOutput)
}

func (AccountFailoverPoliciesOutput) ElementType() reflect.Type {
	return accountFailoverPoliciesType
}

func (o AccountFailoverPoliciesOutput) ToAccountFailoverPoliciesOutput() AccountFailoverPoliciesOutput {
	return o
}

func (o AccountFailoverPoliciesOutput) ToAccountFailoverPoliciesOutputWithContext(ctx context.Context) AccountFailoverPoliciesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AccountFailoverPoliciesOutput{}) }

var accountFailoverPoliciesArrayType = reflect.TypeOf((*[]AccountFailoverPolicies)(nil)).Elem()

type AccountFailoverPoliciesArrayInput interface {
	pulumi.Input

	ToAccountFailoverPoliciesArrayOutput() AccountFailoverPoliciesArrayOutput
	ToAccountFailoverPoliciesArrayOutputWithContext(ctx context.Context) AccountFailoverPoliciesArrayOutput
}

type AccountFailoverPoliciesArrayArgs []AccountFailoverPoliciesInput

func (AccountFailoverPoliciesArrayArgs) ElementType() reflect.Type {
	return accountFailoverPoliciesArrayType
}

func (a AccountFailoverPoliciesArrayArgs) ToAccountFailoverPoliciesArrayOutput() AccountFailoverPoliciesArrayOutput {
	return pulumi.ToOutput(a).(AccountFailoverPoliciesArrayOutput)
}

func (a AccountFailoverPoliciesArrayArgs) ToAccountFailoverPoliciesArrayOutputWithContext(ctx context.Context) AccountFailoverPoliciesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AccountFailoverPoliciesArrayOutput)
}

type AccountFailoverPoliciesArrayOutput struct { *pulumi.OutputState }

func (o AccountFailoverPoliciesArrayOutput) Index(i pulumi.IntInput) AccountFailoverPoliciesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) AccountFailoverPolicies {
		return vs[0].([]AccountFailoverPolicies)[vs[1].(int)]
	}).(AccountFailoverPoliciesOutput)
}

func (AccountFailoverPoliciesArrayOutput) ElementType() reflect.Type {
	return accountFailoverPoliciesArrayType
}

func (o AccountFailoverPoliciesArrayOutput) ToAccountFailoverPoliciesArrayOutput() AccountFailoverPoliciesArrayOutput {
	return o
}

func (o AccountFailoverPoliciesArrayOutput) ToAccountFailoverPoliciesArrayOutputWithContext(ctx context.Context) AccountFailoverPoliciesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AccountFailoverPoliciesArrayOutput{}) }

type AccountGeoLocations struct {
	// The failover priority of the region. A failover priority of `0` indicates a write region. The maximum value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the regions in which the database account exists. Changing this causes the location to be re-provisioned and cannot be changed for the location with failover priority `0`.
	FailoverPriority int `pulumi:"failoverPriority"`
	// The ID of the virtual network subnet.
	Id *string `pulumi:"id"`
	// The name of the Azure region to host replicated data.
	Location string `pulumi:"location"`
	// The string used to generate the document endpoints for this region. If not specified it defaults to `${cosmosdb_account.name}-${location}`. Changing this causes the location to be deleted and re-provisioned and cannot be changed for the location with failover priority `0`.
	Prefix *string `pulumi:"prefix"`
}
var accountGeoLocationsType = reflect.TypeOf((*AccountGeoLocations)(nil)).Elem()

type AccountGeoLocationsInput interface {
	pulumi.Input

	ToAccountGeoLocationsOutput() AccountGeoLocationsOutput
	ToAccountGeoLocationsOutputWithContext(ctx context.Context) AccountGeoLocationsOutput
}

type AccountGeoLocationsArgs struct {
	// The failover priority of the region. A failover priority of `0` indicates a write region. The maximum value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the regions in which the database account exists. Changing this causes the location to be re-provisioned and cannot be changed for the location with failover priority `0`.
	FailoverPriority pulumi.IntInput `pulumi:"failoverPriority"`
	// The ID of the virtual network subnet.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the Azure region to host replicated data.
	Location pulumi.StringInput `pulumi:"location"`
	// The string used to generate the document endpoints for this region. If not specified it defaults to `${cosmosdb_account.name}-${location}`. Changing this causes the location to be deleted and re-provisioned and cannot be changed for the location with failover priority `0`.
	Prefix pulumi.StringInput `pulumi:"prefix"`
}

func (AccountGeoLocationsArgs) ElementType() reflect.Type {
	return accountGeoLocationsType
}

func (a AccountGeoLocationsArgs) ToAccountGeoLocationsOutput() AccountGeoLocationsOutput {
	return pulumi.ToOutput(a).(AccountGeoLocationsOutput)
}

func (a AccountGeoLocationsArgs) ToAccountGeoLocationsOutputWithContext(ctx context.Context) AccountGeoLocationsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AccountGeoLocationsOutput)
}

type AccountGeoLocationsOutput struct { *pulumi.OutputState }

// The failover priority of the region. A failover priority of `0` indicates a write region. The maximum value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the regions in which the database account exists. Changing this causes the location to be re-provisioned and cannot be changed for the location with failover priority `0`.
func (o AccountGeoLocationsOutput) FailoverPriority() pulumi.IntOutput {
	return o.Apply(func(v AccountGeoLocations) int {
		return v.FailoverPriority
	}).(pulumi.IntOutput)
}

// The ID of the virtual network subnet.
func (o AccountGeoLocationsOutput) Id() pulumi.StringOutput {
	return o.Apply(func(v AccountGeoLocations) string {
		if v.Id == nil { return *new(string) } else { return *v.Id }
	}).(pulumi.StringOutput)
}

// The name of the Azure region to host replicated data.
func (o AccountGeoLocationsOutput) Location() pulumi.StringOutput {
	return o.Apply(func(v AccountGeoLocations) string {
		return v.Location
	}).(pulumi.StringOutput)
}

// The string used to generate the document endpoints for this region. If not specified it defaults to `${cosmosdb_account.name}-${location}`. Changing this causes the location to be deleted and re-provisioned and cannot be changed for the location with failover priority `0`.
func (o AccountGeoLocationsOutput) Prefix() pulumi.StringOutput {
	return o.Apply(func(v AccountGeoLocations) string {
		if v.Prefix == nil { return *new(string) } else { return *v.Prefix }
	}).(pulumi.StringOutput)
}

func (AccountGeoLocationsOutput) ElementType() reflect.Type {
	return accountGeoLocationsType
}

func (o AccountGeoLocationsOutput) ToAccountGeoLocationsOutput() AccountGeoLocationsOutput {
	return o
}

func (o AccountGeoLocationsOutput) ToAccountGeoLocationsOutputWithContext(ctx context.Context) AccountGeoLocationsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AccountGeoLocationsOutput{}) }

var accountGeoLocationsArrayType = reflect.TypeOf((*[]AccountGeoLocations)(nil)).Elem()

type AccountGeoLocationsArrayInput interface {
	pulumi.Input

	ToAccountGeoLocationsArrayOutput() AccountGeoLocationsArrayOutput
	ToAccountGeoLocationsArrayOutputWithContext(ctx context.Context) AccountGeoLocationsArrayOutput
}

type AccountGeoLocationsArrayArgs []AccountGeoLocationsInput

func (AccountGeoLocationsArrayArgs) ElementType() reflect.Type {
	return accountGeoLocationsArrayType
}

func (a AccountGeoLocationsArrayArgs) ToAccountGeoLocationsArrayOutput() AccountGeoLocationsArrayOutput {
	return pulumi.ToOutput(a).(AccountGeoLocationsArrayOutput)
}

func (a AccountGeoLocationsArrayArgs) ToAccountGeoLocationsArrayOutputWithContext(ctx context.Context) AccountGeoLocationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AccountGeoLocationsArrayOutput)
}

type AccountGeoLocationsArrayOutput struct { *pulumi.OutputState }

func (o AccountGeoLocationsArrayOutput) Index(i pulumi.IntInput) AccountGeoLocationsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) AccountGeoLocations {
		return vs[0].([]AccountGeoLocations)[vs[1].(int)]
	}).(AccountGeoLocationsOutput)
}

func (AccountGeoLocationsArrayOutput) ElementType() reflect.Type {
	return accountGeoLocationsArrayType
}

func (o AccountGeoLocationsArrayOutput) ToAccountGeoLocationsArrayOutput() AccountGeoLocationsArrayOutput {
	return o
}

func (o AccountGeoLocationsArrayOutput) ToAccountGeoLocationsArrayOutputWithContext(ctx context.Context) AccountGeoLocationsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AccountGeoLocationsArrayOutput{}) }

type AccountVirtualNetworkRules struct {
	// The ID of the virtual network subnet.
	Id string `pulumi:"id"`
}
var accountVirtualNetworkRulesType = reflect.TypeOf((*AccountVirtualNetworkRules)(nil)).Elem()

type AccountVirtualNetworkRulesInput interface {
	pulumi.Input

	ToAccountVirtualNetworkRulesOutput() AccountVirtualNetworkRulesOutput
	ToAccountVirtualNetworkRulesOutputWithContext(ctx context.Context) AccountVirtualNetworkRulesOutput
}

type AccountVirtualNetworkRulesArgs struct {
	// The ID of the virtual network subnet.
	Id pulumi.StringInput `pulumi:"id"`
}

func (AccountVirtualNetworkRulesArgs) ElementType() reflect.Type {
	return accountVirtualNetworkRulesType
}

func (a AccountVirtualNetworkRulesArgs) ToAccountVirtualNetworkRulesOutput() AccountVirtualNetworkRulesOutput {
	return pulumi.ToOutput(a).(AccountVirtualNetworkRulesOutput)
}

func (a AccountVirtualNetworkRulesArgs) ToAccountVirtualNetworkRulesOutputWithContext(ctx context.Context) AccountVirtualNetworkRulesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AccountVirtualNetworkRulesOutput)
}

type AccountVirtualNetworkRulesOutput struct { *pulumi.OutputState }

// The ID of the virtual network subnet.
func (o AccountVirtualNetworkRulesOutput) Id() pulumi.StringOutput {
	return o.Apply(func(v AccountVirtualNetworkRules) string {
		return v.Id
	}).(pulumi.StringOutput)
}

func (AccountVirtualNetworkRulesOutput) ElementType() reflect.Type {
	return accountVirtualNetworkRulesType
}

func (o AccountVirtualNetworkRulesOutput) ToAccountVirtualNetworkRulesOutput() AccountVirtualNetworkRulesOutput {
	return o
}

func (o AccountVirtualNetworkRulesOutput) ToAccountVirtualNetworkRulesOutputWithContext(ctx context.Context) AccountVirtualNetworkRulesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AccountVirtualNetworkRulesOutput{}) }

var accountVirtualNetworkRulesArrayType = reflect.TypeOf((*[]AccountVirtualNetworkRules)(nil)).Elem()

type AccountVirtualNetworkRulesArrayInput interface {
	pulumi.Input

	ToAccountVirtualNetworkRulesArrayOutput() AccountVirtualNetworkRulesArrayOutput
	ToAccountVirtualNetworkRulesArrayOutputWithContext(ctx context.Context) AccountVirtualNetworkRulesArrayOutput
}

type AccountVirtualNetworkRulesArrayArgs []AccountVirtualNetworkRulesInput

func (AccountVirtualNetworkRulesArrayArgs) ElementType() reflect.Type {
	return accountVirtualNetworkRulesArrayType
}

func (a AccountVirtualNetworkRulesArrayArgs) ToAccountVirtualNetworkRulesArrayOutput() AccountVirtualNetworkRulesArrayOutput {
	return pulumi.ToOutput(a).(AccountVirtualNetworkRulesArrayOutput)
}

func (a AccountVirtualNetworkRulesArrayArgs) ToAccountVirtualNetworkRulesArrayOutputWithContext(ctx context.Context) AccountVirtualNetworkRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AccountVirtualNetworkRulesArrayOutput)
}

type AccountVirtualNetworkRulesArrayOutput struct { *pulumi.OutputState }

func (o AccountVirtualNetworkRulesArrayOutput) Index(i pulumi.IntInput) AccountVirtualNetworkRulesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) AccountVirtualNetworkRules {
		return vs[0].([]AccountVirtualNetworkRules)[vs[1].(int)]
	}).(AccountVirtualNetworkRulesOutput)
}

func (AccountVirtualNetworkRulesArrayOutput) ElementType() reflect.Type {
	return accountVirtualNetworkRulesArrayType
}

func (o AccountVirtualNetworkRulesArrayOutput) ToAccountVirtualNetworkRulesArrayOutput() AccountVirtualNetworkRulesArrayOutput {
	return o
}

func (o AccountVirtualNetworkRulesArrayOutput) ToAccountVirtualNetworkRulesArrayOutputWithContext(ctx context.Context) AccountVirtualNetworkRulesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AccountVirtualNetworkRulesArrayOutput{}) }

