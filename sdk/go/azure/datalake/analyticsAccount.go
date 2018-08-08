// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datalake

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manage an Azure Data Lake Analytics Account.
type AnalyticsAccount struct {
	s *pulumi.ResourceState
}

// NewAnalyticsAccount registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsAccount(ctx *pulumi.Context,
	name string, args *AnalyticsAccountArgs, opts ...pulumi.ResourceOpt) (*AnalyticsAccount, error) {
	if args == nil || args.DefaultStoreAccountName == nil {
		return nil, errors.New("missing required argument 'DefaultStoreAccountName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["defaultStoreAccountName"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
		inputs["tier"] = nil
	} else {
		inputs["defaultStoreAccountName"] = args.DefaultStoreAccountName
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
		inputs["tier"] = args.Tier
	}
	s, err := ctx.RegisterResource("azure:datalake/analyticsAccount:AnalyticsAccount", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AnalyticsAccount{s: s}, nil
}

// GetAnalyticsAccount gets an existing AnalyticsAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsAccount(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AnalyticsAccountState, opts ...pulumi.ResourceOpt) (*AnalyticsAccount, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["defaultStoreAccountName"] = state.DefaultStoreAccountName
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
		inputs["tier"] = state.Tier
	}
	s, err := ctx.ReadResource("azure:datalake/analyticsAccount:AnalyticsAccount", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AnalyticsAccount{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AnalyticsAccount) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AnalyticsAccount) ID() *pulumi.IDOutput {
	return r.s.ID
}

// Specifies the data lake store to use by default. Changing this forces a new resource to be created.
func (r *AnalyticsAccount) DefaultStoreAccountName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultStoreAccountName"])
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *AnalyticsAccount) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the Data Lake Analytics Account. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
func (r *AnalyticsAccount) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which to create the Data Lake Analytics Account.
func (r *AnalyticsAccount) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *AnalyticsAccount) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The monthly commitment tier for Data Lake Analytics Account. Accepted values are `Consumption`, `Commitment_100000AUHours`, `Commitment_10000AUHours`, `Commitment_1000AUHours`, `Commitment_100AUHours`, `Commitment_500000AUHours`, `Commitment_50000AUHours`, `Commitment_5000AUHours`, or `Commitment_500AUHours`.
func (r *AnalyticsAccount) Tier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tier"])
}

// Input properties used for looking up and filtering AnalyticsAccount resources.
type AnalyticsAccountState struct {
	// Specifies the data lake store to use by default. Changing this forces a new resource to be created.
	DefaultStoreAccountName interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Data Lake Analytics Account. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name interface{}
	// The name of the resource group in which to create the Data Lake Analytics Account.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The monthly commitment tier for Data Lake Analytics Account. Accepted values are `Consumption`, `Commitment_100000AUHours`, `Commitment_10000AUHours`, `Commitment_1000AUHours`, `Commitment_100AUHours`, `Commitment_500000AUHours`, `Commitment_50000AUHours`, `Commitment_5000AUHours`, or `Commitment_500AUHours`.
	Tier interface{}
}

// The set of arguments for constructing a AnalyticsAccount resource.
type AnalyticsAccountArgs struct {
	// Specifies the data lake store to use by default. Changing this forces a new resource to be created.
	DefaultStoreAccountName interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Data Lake Analytics Account. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name interface{}
	// The name of the resource group in which to create the Data Lake Analytics Account.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The monthly commitment tier for Data Lake Analytics Account. Accepted values are `Consumption`, `Commitment_100000AUHours`, `Commitment_10000AUHours`, `Commitment_1000AUHours`, `Commitment_100AUHours`, `Commitment_500000AUHours`, `Commitment_50000AUHours`, `Commitment_5000AUHours`, or `Commitment_500AUHours`.
	Tier interface{}
}
