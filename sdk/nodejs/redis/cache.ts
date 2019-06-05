// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Redis Cache.
 * 
 * ## Example Usage
 * 
 * This example provisions a Standard Redis Cache. Other examples of the `azurerm_redis_cache` resource can be found in [the `./examples/redis-cache` directory within the Github Repository](https://github.com/terraform-providers/terraform-provider-azurerm/tree/master/examples/redis-cache)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "West Europe",
 *     name: "example-resources",
 * });
 * // NOTE: the Name used for Redis needs to be globally unique
 * const exampleCache = new azure.redis.Cache("example", {
 *     capacity: 2,
 *     enableNonSslPort: false,
 *     family: "C",
 *     location: exampleResourceGroup.location,
 *     minimumTlsVersion: "1.2",
 *     name: "example-cache",
 *     redisConfiguration: {},
 *     resourceGroupName: exampleResourceGroup.name,
 *     skuName: "Standard",
 * });
 * ```
 * 
 * ## Default Redis Configuration Values
 * 
 * | Redis Value                     | Basic        | Standard     | Premium      |
 * | ------------------------------- | ------------ | ------------ | ------------ |
 * | enable_authentication           | true         | true         | true         |
 * | maxmemory_reserved              | 2            | 50           | 200          |
 * | maxfragmentationmemory_reserved | 2            | 50           | 200          |
 * | maxmemory_delta                 | 2            | 50           | 200          |
 * | maxmemory_policy                | volatile-lru | volatile-lru | volatile-lru |
 * 
 * > **NOTE:** The `maxmemory_reserved`, `maxmemory_delta` and `maxfragmentationmemory-reserved` settings are only available for Standard and Premium caches. More details are available in the Relevant Links section below._
 * 
 * ---
 * 
 * A `patch_schedule` block supports the following:
 * 
 * * `day_of_week` (Required) the Weekday name - possible values include `Monday`, `Tuesday`, `Wednesday` etc.
 * 
 * * `start_hour_utc` - (Optional) the Start Hour for maintenance in UTC - possible values range from `0 - 23`.
 * 
 * > **Note:** The Patch Window lasts for `5` hours from the `start_hour_utc`.
 * 
 * ## Relevant Links
 * 
 *  - [Azure Redis Cache: SKU specific configuration limitations](https://azure.microsoft.com/en-us/documentation/articles/cache-configure/#advanced-settings)
 *  - [Redis: Available Configuration Settings](http://redis.io/topics/config)
 */
export class Cache extends pulumi.CustomResource {
    /**
     * Get an existing Cache resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CacheState, opts?: pulumi.CustomResourceOptions): Cache {
        return new Cache(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:redis/cache:Cache';

    /**
     * Returns true if the given object is an instance of Cache.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cache {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cache.__pulumiType;
    }

    /**
     * The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
     */
    public readonly capacity!: pulumi.Output<number>;
    /**
     * Enable the non-SSL port (6789) - disabled by default.
     */
    public readonly enableNonSslPort!: pulumi.Output<boolean | undefined>;
    /**
     * The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
     */
    public readonly family!: pulumi.Output<string>;
    /**
     * The Hostname of the Redis Instance
     */
    public /*out*/ readonly hostname!: pulumi.Output<string>;
    /**
     * The location of the resource group.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The minimum TLS version.  Defaults to `1.0`.
     */
    public readonly minimumTlsVersion!: pulumi.Output<string | undefined>;
    /**
     * The name of the Redis instance. Changing this forces a
     * new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of `patch_schedule` blocks as defined below - only available for Premium SKU's.
     */
    public readonly patchSchedules!: pulumi.Output<{ dayOfWeek: string, startHourUtc?: number }[] | undefined>;
    /**
     * The non-SSL Port of the Redis Instance
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * The Primary Access Key for the Redis Instance
     */
    public /*out*/ readonly primaryAccessKey!: pulumi.Output<string>;
    /**
     * The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
     */
    public readonly privateStaticIpAddress!: pulumi.Output<string>;
    /**
     * A `redis_configuration` as defined below - with some limitations by SKU - defaults/details are shown below.
     */
    public readonly redisConfiguration!: pulumi.Output<{ aofBackupEnabled?: boolean, aofStorageConnectionString0?: string, aofStorageConnectionString1?: string, enableAuthentication?: boolean, maxclients: number, maxfragmentationmemoryReserved: number, maxmemoryDelta: number, maxmemoryPolicy?: string, maxmemoryReserved: number, notifyKeyspaceEvents?: string, rdbBackupEnabled?: boolean, rdbBackupFrequency?: number, rdbBackupMaxSnapshotCount?: number, rdbStorageConnectionString?: string }>;
    /**
     * The name of the resource group in which to
     * create the Redis instance.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Secondary Access Key for the Redis Instance
     */
    public /*out*/ readonly secondaryAccessKey!: pulumi.Output<string>;
    /**
     * *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
     */
    public readonly shardCount!: pulumi.Output<number | undefined>;
    /**
     * The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * The SSL Port of the Redis Instance
     */
    public /*out*/ readonly sslPort!: pulumi.Output<number>;
    /**
     * The ID of the Subnet within which the Redis Cache should be deployed. Changing this forces a new resource to be created.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any}>;
    /**
     * A list of a single item of the Availability Zone which the Redis Cache should be allocated in.
     */
    public readonly zones!: pulumi.Output<string | undefined>;

    /**
     * Create a Cache resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CacheArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CacheArgs | CacheState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CacheState | undefined;
            inputs["capacity"] = state ? state.capacity : undefined;
            inputs["enableNonSslPort"] = state ? state.enableNonSslPort : undefined;
            inputs["family"] = state ? state.family : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["minimumTlsVersion"] = state ? state.minimumTlsVersion : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["patchSchedules"] = state ? state.patchSchedules : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["primaryAccessKey"] = state ? state.primaryAccessKey : undefined;
            inputs["privateStaticIpAddress"] = state ? state.privateStaticIpAddress : undefined;
            inputs["redisConfiguration"] = state ? state.redisConfiguration : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryAccessKey"] = state ? state.secondaryAccessKey : undefined;
            inputs["shardCount"] = state ? state.shardCount : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["sslPort"] = state ? state.sslPort : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as CacheArgs | undefined;
            if (!args || args.capacity === undefined) {
                throw new Error("Missing required property 'capacity'");
            }
            if (!args || args.family === undefined) {
                throw new Error("Missing required property 'family'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.skuName === undefined) {
                throw new Error("Missing required property 'skuName'");
            }
            inputs["capacity"] = args ? args.capacity : undefined;
            inputs["enableNonSslPort"] = args ? args.enableNonSslPort : undefined;
            inputs["family"] = args ? args.family : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["minimumTlsVersion"] = args ? args.minimumTlsVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["patchSchedules"] = args ? args.patchSchedules : undefined;
            inputs["privateStaticIpAddress"] = args ? args.privateStaticIpAddress : undefined;
            inputs["redisConfiguration"] = args ? args.redisConfiguration : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["shardCount"] = args ? args.shardCount : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["hostname"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["primaryAccessKey"] = undefined /*out*/;
            inputs["secondaryAccessKey"] = undefined /*out*/;
            inputs["sslPort"] = undefined /*out*/;
        }
        super(Cache.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cache resources.
 */
export interface CacheState {
    /**
     * The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
     */
    readonly capacity?: pulumi.Input<number>;
    /**
     * Enable the non-SSL port (6789) - disabled by default.
     */
    readonly enableNonSslPort?: pulumi.Input<boolean>;
    /**
     * The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
     */
    readonly family?: pulumi.Input<string>;
    /**
     * The Hostname of the Redis Instance
     */
    readonly hostname?: pulumi.Input<string>;
    /**
     * The location of the resource group.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The minimum TLS version.  Defaults to `1.0`.
     */
    readonly minimumTlsVersion?: pulumi.Input<string>;
    /**
     * The name of the Redis instance. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of `patch_schedule` blocks as defined below - only available for Premium SKU's.
     */
    readonly patchSchedules?: pulumi.Input<pulumi.Input<{ dayOfWeek: pulumi.Input<string>, startHourUtc?: pulumi.Input<number> }>[]>;
    /**
     * The non-SSL Port of the Redis Instance
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The Primary Access Key for the Redis Instance
     */
    readonly primaryAccessKey?: pulumi.Input<string>;
    /**
     * The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
     */
    readonly privateStaticIpAddress?: pulumi.Input<string>;
    /**
     * A `redis_configuration` as defined below - with some limitations by SKU - defaults/details are shown below.
     */
    readonly redisConfiguration?: pulumi.Input<{ aofBackupEnabled?: pulumi.Input<boolean>, aofStorageConnectionString0?: pulumi.Input<string>, aofStorageConnectionString1?: pulumi.Input<string>, enableAuthentication?: pulumi.Input<boolean>, maxclients?: pulumi.Input<number>, maxfragmentationmemoryReserved?: pulumi.Input<number>, maxmemoryDelta?: pulumi.Input<number>, maxmemoryPolicy?: pulumi.Input<string>, maxmemoryReserved?: pulumi.Input<number>, notifyKeyspaceEvents?: pulumi.Input<string>, rdbBackupEnabled?: pulumi.Input<boolean>, rdbBackupFrequency?: pulumi.Input<number>, rdbBackupMaxSnapshotCount?: pulumi.Input<number>, rdbStorageConnectionString?: pulumi.Input<string> }>;
    /**
     * The name of the resource group in which to
     * create the Redis instance.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The Secondary Access Key for the Redis Instance
     */
    readonly secondaryAccessKey?: pulumi.Input<string>;
    /**
     * *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
     */
    readonly shardCount?: pulumi.Input<number>;
    /**
     * The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
     */
    readonly skuName?: pulumi.Input<string>;
    /**
     * The SSL Port of the Redis Instance
     */
    readonly sslPort?: pulumi.Input<number>;
    /**
     * The ID of the Subnet within which the Redis Cache should be deployed. Changing this forces a new resource to be created.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A list of a single item of the Availability Zone which the Redis Cache should be allocated in.
     */
    readonly zones?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cache resource.
 */
export interface CacheArgs {
    /**
     * The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
     */
    readonly capacity: pulumi.Input<number>;
    /**
     * Enable the non-SSL port (6789) - disabled by default.
     */
    readonly enableNonSslPort?: pulumi.Input<boolean>;
    /**
     * The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
     */
    readonly family: pulumi.Input<string>;
    /**
     * The location of the resource group.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The minimum TLS version.  Defaults to `1.0`.
     */
    readonly minimumTlsVersion?: pulumi.Input<string>;
    /**
     * The name of the Redis instance. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of `patch_schedule` blocks as defined below - only available for Premium SKU's.
     */
    readonly patchSchedules?: pulumi.Input<pulumi.Input<{ dayOfWeek: pulumi.Input<string>, startHourUtc?: pulumi.Input<number> }>[]>;
    /**
     * The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
     */
    readonly privateStaticIpAddress?: pulumi.Input<string>;
    /**
     * A `redis_configuration` as defined below - with some limitations by SKU - defaults/details are shown below.
     */
    readonly redisConfiguration?: pulumi.Input<{ aofBackupEnabled?: pulumi.Input<boolean>, aofStorageConnectionString0?: pulumi.Input<string>, aofStorageConnectionString1?: pulumi.Input<string>, enableAuthentication?: pulumi.Input<boolean>, maxclients?: pulumi.Input<number>, maxfragmentationmemoryReserved?: pulumi.Input<number>, maxmemoryDelta?: pulumi.Input<number>, maxmemoryPolicy?: pulumi.Input<string>, maxmemoryReserved?: pulumi.Input<number>, notifyKeyspaceEvents?: pulumi.Input<string>, rdbBackupEnabled?: pulumi.Input<boolean>, rdbBackupFrequency?: pulumi.Input<number>, rdbBackupMaxSnapshotCount?: pulumi.Input<number>, rdbStorageConnectionString?: pulumi.Input<string> }>;
    /**
     * The name of the resource group in which to
     * create the Redis instance.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
     */
    readonly shardCount?: pulumi.Input<number>;
    /**
     * The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
     */
    readonly skuName: pulumi.Input<string>;
    /**
     * The ID of the Subnet within which the Redis Cache should be deployed. Changing this forces a new resource to be created.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A list of a single item of the Availability Zone which the Redis Cache should be allocated in.
     */
    readonly zones?: pulumi.Input<string>;
}
