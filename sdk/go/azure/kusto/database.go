// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) Database
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/kusto_database.html.markdown.
type Database struct {
	pulumi.CustomResourceState

	// Specifies the name of the Kusto Cluster this database will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`

	// The time the data that should be kept in cache for fast queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	HotCachePeriod pulumi.StringOutput `pulumi:"hotCachePeriod"`

	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// The name of the Kusto Database to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The size of the database in bytes.
	Size pulumi.Float64Output `pulumi:"size"`

	// The time the data should be kept before it stops being accessible to queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	SoftDeletePeriod pulumi.StringOutput `pulumi:"softDeletePeriod"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ClusterName; i != nil { inputs["clusterName"] = i.ToStringOutput() }
		if i := args.HotCachePeriod; i != nil { inputs["hotCachePeriod"] = i.ToStringOutput() }
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.SoftDeletePeriod; i != nil { inputs["softDeletePeriod"] = i.ToStringOutput() }
	}
	var resource Database
	err := ctx.RegisterResource("azure:kusto/database:Database", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ClusterName; i != nil { inputs["clusterName"] = i.ToStringOutput() }
		if i := state.HotCachePeriod; i != nil { inputs["hotCachePeriod"] = i.ToStringOutput() }
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Size; i != nil { inputs["size"] = i.ToFloat64Output() }
		if i := state.SoftDeletePeriod; i != nil { inputs["softDeletePeriod"] = i.ToStringOutput() }
	}
	var resource Database
	err := ctx.ReadResource("azure:kusto/database:Database", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type DatabaseState struct {
	// Specifies the name of the Kusto Cluster this database will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The time the data that should be kept in cache for fast queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	HotCachePeriod pulumi.StringInput `pulumi:"hotCachePeriod"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// The name of the Kusto Database to create. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The size of the database in bytes.
	Size pulumi.Float64Input `pulumi:"size"`
	// The time the data should be kept before it stops being accessible to queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	SoftDeletePeriod pulumi.StringInput `pulumi:"softDeletePeriod"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Specifies the name of the Kusto Cluster this database will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The time the data that should be kept in cache for fast queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	HotCachePeriod pulumi.StringInput `pulumi:"hotCachePeriod"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// The name of the Kusto Database to create. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The time the data should be kept before it stops being accessible to queries as ISO 8601 timespan. Default is unlimited. For more information see: [ISO 8601 Timespan](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	SoftDeletePeriod pulumi.StringInput `pulumi:"softDeletePeriod"`
}
