// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a MySQL Server.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/mysql_server.html.markdown.
type Server struct {
	pulumi.CustomResourceState

	// The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringOutput `pulumi:"administratorLogin"`

	// The Password associated with the `administratorLogin` for the MySQL Server.
	AdministratorLoginPassword pulumi.StringOutput `pulumi:"administratorLoginPassword"`

	// The FQDN of the MySQL Server.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// A `sku` block as defined below.
	Sku ServerSkuOutput `pulumi:"sku"`

	// Specifies if SSL should be enforced on connections. Possible values are `Enabled` and `Disabled`.
	SslEnforcement pulumi.StringOutput `pulumi:"sslEnforcement"`

	// A `storageProfile` block as defined below.
	StorageProfile ServerStorageProfileOutput `pulumi:"storageProfile"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Specifies the version of MySQL to use. Valid values are `5.6` and `5.7`. Changing this forces a new resource to be created.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil || args.AdministratorLogin == nil {
		return nil, errors.New("missing required argument 'AdministratorLogin'")
	}
	if args == nil || args.AdministratorLoginPassword == nil {
		return nil, errors.New("missing required argument 'AdministratorLoginPassword'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil || args.SslEnforcement == nil {
		return nil, errors.New("missing required argument 'SslEnforcement'")
	}
	if args == nil || args.StorageProfile == nil {
		return nil, errors.New("missing required argument 'StorageProfile'")
	}
	if args == nil || args.Version == nil {
		return nil, errors.New("missing required argument 'Version'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AdministratorLogin; i != nil { inputs["administratorLogin"] = i.ToStringOutput() }
		if i := args.AdministratorLoginPassword; i != nil { inputs["administratorLoginPassword"] = i.ToStringOutput() }
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Sku; i != nil { inputs["sku"] = i.ToServerSkuOutput() }
		if i := args.SslEnforcement; i != nil { inputs["sslEnforcement"] = i.ToStringOutput() }
		if i := args.StorageProfile; i != nil { inputs["storageProfile"] = i.ToServerStorageProfileOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.Version; i != nil { inputs["version"] = i.ToStringOutput() }
	}
	var resource Server
	err := ctx.RegisterResource("azure:mysql/server:Server", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AdministratorLogin; i != nil { inputs["administratorLogin"] = i.ToStringOutput() }
		if i := state.AdministratorLoginPassword; i != nil { inputs["administratorLoginPassword"] = i.ToStringOutput() }
		if i := state.Fqdn; i != nil { inputs["fqdn"] = i.ToStringOutput() }
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Sku; i != nil { inputs["sku"] = i.ToServerSkuOutput() }
		if i := state.SslEnforcement; i != nil { inputs["sslEnforcement"] = i.ToStringOutput() }
		if i := state.StorageProfile; i != nil { inputs["storageProfile"] = i.ToServerStorageProfileOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.Version; i != nil { inputs["version"] = i.ToStringOutput() }
	}
	var resource Server
	err := ctx.ReadResource("azure:mysql/server:Server", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type ServerState struct {
	// The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringInput `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the MySQL Server.
	AdministratorLoginPassword pulumi.StringInput `pulumi:"administratorLoginPassword"`
	// The FQDN of the MySQL Server.
	Fqdn pulumi.StringInput `pulumi:"fqdn"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku ServerSkuInput `pulumi:"sku"`
	// Specifies if SSL should be enforced on connections. Possible values are `Enabled` and `Disabled`.
	SslEnforcement pulumi.StringInput `pulumi:"sslEnforcement"`
	// A `storageProfile` block as defined below.
	StorageProfile ServerStorageProfileInput `pulumi:"storageProfile"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Specifies the version of MySQL to use. Valid values are `5.6` and `5.7`. Changing this forces a new resource to be created.
	Version pulumi.StringInput `pulumi:"version"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringInput `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the MySQL Server.
	AdministratorLoginPassword pulumi.StringInput `pulumi:"administratorLoginPassword"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku ServerSkuInput `pulumi:"sku"`
	// Specifies if SSL should be enforced on connections. Possible values are `Enabled` and `Disabled`.
	SslEnforcement pulumi.StringInput `pulumi:"sslEnforcement"`
	// A `storageProfile` block as defined below.
	StorageProfile ServerStorageProfileInput `pulumi:"storageProfile"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Specifies the version of MySQL to use. Valid values are `5.6` and `5.7`. Changing this forces a new resource to be created.
	Version pulumi.StringInput `pulumi:"version"`
}
type ServerSku struct {
	// The scale up/out capacity, representing server's compute units.
	Capacity int `pulumi:"capacity"`
	// The `family` of hardware `Gen4` or `Gen5`, before selecting your `family` check the [product documentation](https://docs.microsoft.com/en-us/azure/mysql/concepts-pricing-tiers#compute-generations-vcores-and-memory) for availability in your region.
	Family string `pulumi:"family"`
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	Name string `pulumi:"name"`
	// The tier of the particular SKU. Possible values are `Basic`, `GeneralPurpose`, and `MemoryOptimized`. For more information see the [product documentation](https://docs.microsoft.com/en-us/azure/mysql/concepts-pricing-tiers).
	Tier string `pulumi:"tier"`
}
var serverSkuType = reflect.TypeOf((*ServerSku)(nil)).Elem()

type ServerSkuInput interface {
	pulumi.Input

	ToServerSkuOutput() ServerSkuOutput
	ToServerSkuOutputWithContext(ctx context.Context) ServerSkuOutput
}

type ServerSkuArgs struct {
	// The scale up/out capacity, representing server's compute units.
	Capacity pulumi.IntInput `pulumi:"capacity"`
	// The `family` of hardware `Gen4` or `Gen5`, before selecting your `family` check the [product documentation](https://docs.microsoft.com/en-us/azure/mysql/concepts-pricing-tiers#compute-generations-vcores-and-memory) for availability in your region.
	Family pulumi.StringInput `pulumi:"family"`
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	Name pulumi.StringInput `pulumi:"name"`
	// The tier of the particular SKU. Possible values are `Basic`, `GeneralPurpose`, and `MemoryOptimized`. For more information see the [product documentation](https://docs.microsoft.com/en-us/azure/mysql/concepts-pricing-tiers).
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (ServerSkuArgs) ElementType() reflect.Type {
	return serverSkuType
}

func (a ServerSkuArgs) ToServerSkuOutput() ServerSkuOutput {
	return pulumi.ToOutput(a).(ServerSkuOutput)
}

func (a ServerSkuArgs) ToServerSkuOutputWithContext(ctx context.Context) ServerSkuOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServerSkuOutput)
}

type ServerSkuOutput struct { *pulumi.OutputState }

// The scale up/out capacity, representing server's compute units.
func (o ServerSkuOutput) Capacity() pulumi.IntOutput {
	return o.Apply(func(v ServerSku) int {
		return v.Capacity
	}).(pulumi.IntOutput)
}

// The `family` of hardware `Gen4` or `Gen5`, before selecting your `family` check the [product documentation](https://docs.microsoft.com/en-us/azure/mysql/concepts-pricing-tiers#compute-generations-vcores-and-memory) for availability in your region.
func (o ServerSkuOutput) Family() pulumi.StringOutput {
	return o.Apply(func(v ServerSku) string {
		return v.Family
	}).(pulumi.StringOutput)
}

// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
func (o ServerSkuOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v ServerSku) string {
		return v.Name
	}).(pulumi.StringOutput)
}

// The tier of the particular SKU. Possible values are `Basic`, `GeneralPurpose`, and `MemoryOptimized`. For more information see the [product documentation](https://docs.microsoft.com/en-us/azure/mysql/concepts-pricing-tiers).
func (o ServerSkuOutput) Tier() pulumi.StringOutput {
	return o.Apply(func(v ServerSku) string {
		return v.Tier
	}).(pulumi.StringOutput)
}

func (ServerSkuOutput) ElementType() reflect.Type {
	return serverSkuType
}

func (o ServerSkuOutput) ToServerSkuOutput() ServerSkuOutput {
	return o
}

func (o ServerSkuOutput) ToServerSkuOutputWithContext(ctx context.Context) ServerSkuOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ServerSkuOutput{}) }

type ServerStorageProfile struct {
	// Defines whether autogrow is enabled or disabled for the storage. Valid values are `Enabled` or `Disabled`.
	AutoGrow *string `pulumi:"autoGrow"`
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays *int `pulumi:"backupRetentionDays"`
	// Enable Geo-redundant or not for server backup. Valid values for this property are `Enabled` or `Disabled`, not supported for the `basic` tier.
	GeoRedundantBackup *string `pulumi:"geoRedundantBackup"`
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
	StorageMb int `pulumi:"storageMb"`
}
var serverStorageProfileType = reflect.TypeOf((*ServerStorageProfile)(nil)).Elem()

type ServerStorageProfileInput interface {
	pulumi.Input

	ToServerStorageProfileOutput() ServerStorageProfileOutput
	ToServerStorageProfileOutputWithContext(ctx context.Context) ServerStorageProfileOutput
}

type ServerStorageProfileArgs struct {
	// Defines whether autogrow is enabled or disabled for the storage. Valid values are `Enabled` or `Disabled`.
	AutoGrow pulumi.StringInput `pulumi:"autoGrow"`
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays pulumi.IntInput `pulumi:"backupRetentionDays"`
	// Enable Geo-redundant or not for server backup. Valid values for this property are `Enabled` or `Disabled`, not supported for the `basic` tier.
	GeoRedundantBackup pulumi.StringInput `pulumi:"geoRedundantBackup"`
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
	StorageMb pulumi.IntInput `pulumi:"storageMb"`
}

func (ServerStorageProfileArgs) ElementType() reflect.Type {
	return serverStorageProfileType
}

func (a ServerStorageProfileArgs) ToServerStorageProfileOutput() ServerStorageProfileOutput {
	return pulumi.ToOutput(a).(ServerStorageProfileOutput)
}

func (a ServerStorageProfileArgs) ToServerStorageProfileOutputWithContext(ctx context.Context) ServerStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServerStorageProfileOutput)
}

type ServerStorageProfileOutput struct { *pulumi.OutputState }

// Defines whether autogrow is enabled or disabled for the storage. Valid values are `Enabled` or `Disabled`.
func (o ServerStorageProfileOutput) AutoGrow() pulumi.StringOutput {
	return o.Apply(func(v ServerStorageProfile) string {
		if v.AutoGrow == nil { return *new(string) } else { return *v.AutoGrow }
	}).(pulumi.StringOutput)
}

// Backup retention days for the server, supported values are between `7` and `35` days.
func (o ServerStorageProfileOutput) BackupRetentionDays() pulumi.IntOutput {
	return o.Apply(func(v ServerStorageProfile) int {
		if v.BackupRetentionDays == nil { return *new(int) } else { return *v.BackupRetentionDays }
	}).(pulumi.IntOutput)
}

// Enable Geo-redundant or not for server backup. Valid values for this property are `Enabled` or `Disabled`, not supported for the `basic` tier.
func (o ServerStorageProfileOutput) GeoRedundantBackup() pulumi.StringOutput {
	return o.Apply(func(v ServerStorageProfile) string {
		if v.GeoRedundantBackup == nil { return *new(string) } else { return *v.GeoRedundantBackup }
	}).(pulumi.StringOutput)
}

// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
func (o ServerStorageProfileOutput) StorageMb() pulumi.IntOutput {
	return o.Apply(func(v ServerStorageProfile) int {
		return v.StorageMb
	}).(pulumi.IntOutput)
}

func (ServerStorageProfileOutput) ElementType() reflect.Type {
	return serverStorageProfileType
}

func (o ServerStorageProfileOutput) ToServerStorageProfileOutput() ServerStorageProfileOutput {
	return o
}

func (o ServerStorageProfileOutput) ToServerStorageProfileOutputWithContext(ctx context.Context) ServerStorageProfileOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ServerStorageProfileOutput{}) }

