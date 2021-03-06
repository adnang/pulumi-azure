// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql
{
    /// <summary>
    /// Manages a MS SQL Database.
    /// </summary>
    public partial class Database : Pulumi.CustomResource
    {
        /// <summary>
        /// Time in minutes after which database is automatically paused. A value of `-1` means that automatic pause is disabled. This property is only settable for General Purpose Serverless databases.
        /// </summary>
        [Output("autoPauseDelayInMinutes")]
        public Output<int> AutoPauseDelayInMinutes { get; private set; } = null!;

        /// <summary>
        /// Specifies the collation of the database. Changing this forces a new resource to be created.
        /// </summary>
        [Output("collation")]
        public Output<string> Collation { get; private set; } = null!;

        /// <summary>
        /// The create mode of the database. Possible values are `Copy`, `Default`, `OnlineSecondary`, `PointInTimeRestore`, `Restore`, `RestoreExternalBackup`, `RestoreExternalBackupSecondary`, `RestoreLongTermRetentionBackup` and `Secondary`. 
        /// </summary>
        [Output("createMode")]
        public Output<string> CreateMode { get; private set; } = null!;

        /// <summary>
        /// The id of the source database to be referred to create the new database. This should only be used for databases with `create_mode` values that use another database as reference. Changing this forces a new resource to be created.
        /// </summary>
        [Output("creationSourceDatabaseId")]
        public Output<string> CreationSourceDatabaseId { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the elastic pool containing this database. Changing this forces a new resource to be created.
        /// </summary>
        [Output("elasticPoolId")]
        public Output<string?> ElasticPoolId { get; private set; } = null!;

        /// <summary>
        /// A `extended_auditing_policy` block as defined below.
        /// </summary>
        [Output("extendedAuditingPolicy")]
        public Output<Outputs.DatabaseExtendedAuditingPolicy?> ExtendedAuditingPolicy { get; private set; } = null!;

        /// <summary>
        /// Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
        /// </summary>
        [Output("licenseType")]
        public Output<string> LicenseType { get; private set; } = null!;

        /// <summary>
        /// The max size of the database in gigabytes. 
        /// </summary>
        [Output("maxSizeGb")]
        public Output<int> MaxSizeGb { get; private set; } = null!;

        /// <summary>
        /// Minimal capacity that database will always have allocated, if not paused. This property is only settable for General Purpose Serverless databases.
        /// </summary>
        [Output("minCapacity")]
        public Output<double> MinCapacity { get; private set; } = null!;

        /// <summary>
        /// The name of the Ms SQL Database. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
        /// </summary>
        [Output("readReplicaCount")]
        public Output<int> ReadReplicaCount { get; private set; } = null!;

        /// <summary>
        /// If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
        /// </summary>
        [Output("readScale")]
        public Output<bool> ReadScale { get; private set; } = null!;

        /// <summary>
        /// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. This property is only settable for `create_mode`= `PointInTimeRestore`  databases.
        /// </summary>
        [Output("restorePointInTime")]
        public Output<string> RestorePointInTime { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the sample schema to apply when creating this database. Possible value is `AdventureWorksLT`.
        /// </summary>
        [Output("sampleName")]
        public Output<string> SampleName { get; private set; } = null!;

        /// <summary>
        /// The id of the Ms SQL Server on which to create the database. Changing this forces a new resource to be created.
        /// </summary>
        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the sku used by the database. Changing this forces a new resource to be created. For example, `GP_S_Gen5_2`,`HS_Gen4_1`,`BC_Gen5_2`, `ElasticPool`, `Basic`,`S0`, `P2` ,`DW100c`, `DS100`.
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Threat detection policy configuration. The `threat_detection_policy` block supports fields documented below.
        /// </summary>
        [Output("threatDetectionPolicy")]
        public Output<Outputs.DatabaseThreatDetectionPolicy> ThreatDetectionPolicy { get; private set; } = null!;

        /// <summary>
        /// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones. This property is only settable for Premium and Business Critical databases.
        /// </summary>
        [Output("zoneRedundant")]
        public Output<bool> ZoneRedundant { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("azure:mssql/database:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
            : base("azure:mssql/database:Database", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new Database(name, id, state, options);
        }
    }

    public sealed class DatabaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time in minutes after which database is automatically paused. A value of `-1` means that automatic pause is disabled. This property is only settable for General Purpose Serverless databases.
        /// </summary>
        [Input("autoPauseDelayInMinutes")]
        public Input<int>? AutoPauseDelayInMinutes { get; set; }

        /// <summary>
        /// Specifies the collation of the database. Changing this forces a new resource to be created.
        /// </summary>
        [Input("collation")]
        public Input<string>? Collation { get; set; }

        /// <summary>
        /// The create mode of the database. Possible values are `Copy`, `Default`, `OnlineSecondary`, `PointInTimeRestore`, `Restore`, `RestoreExternalBackup`, `RestoreExternalBackupSecondary`, `RestoreLongTermRetentionBackup` and `Secondary`. 
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        /// <summary>
        /// The id of the source database to be referred to create the new database. This should only be used for databases with `create_mode` values that use another database as reference. Changing this forces a new resource to be created.
        /// </summary>
        [Input("creationSourceDatabaseId")]
        public Input<string>? CreationSourceDatabaseId { get; set; }

        /// <summary>
        /// Specifies the ID of the elastic pool containing this database. Changing this forces a new resource to be created.
        /// </summary>
        [Input("elasticPoolId")]
        public Input<string>? ElasticPoolId { get; set; }

        /// <summary>
        /// A `extended_auditing_policy` block as defined below.
        /// </summary>
        [Input("extendedAuditingPolicy")]
        public Input<Inputs.DatabaseExtendedAuditingPolicyArgs>? ExtendedAuditingPolicy { get; set; }

        /// <summary>
        /// Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// The max size of the database in gigabytes. 
        /// </summary>
        [Input("maxSizeGb")]
        public Input<int>? MaxSizeGb { get; set; }

        /// <summary>
        /// Minimal capacity that database will always have allocated, if not paused. This property is only settable for General Purpose Serverless databases.
        /// </summary>
        [Input("minCapacity")]
        public Input<double>? MinCapacity { get; set; }

        /// <summary>
        /// The name of the Ms SQL Database. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
        /// </summary>
        [Input("readReplicaCount")]
        public Input<int>? ReadReplicaCount { get; set; }

        /// <summary>
        /// If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
        /// </summary>
        [Input("readScale")]
        public Input<bool>? ReadScale { get; set; }

        /// <summary>
        /// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. This property is only settable for `create_mode`= `PointInTimeRestore`  databases.
        /// </summary>
        [Input("restorePointInTime")]
        public Input<string>? RestorePointInTime { get; set; }

        /// <summary>
        /// Specifies the name of the sample schema to apply when creating this database. Possible value is `AdventureWorksLT`.
        /// </summary>
        [Input("sampleName")]
        public Input<string>? SampleName { get; set; }

        /// <summary>
        /// The id of the Ms SQL Server on which to create the database. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the sku used by the database. Changing this forces a new resource to be created. For example, `GP_S_Gen5_2`,`HS_Gen4_1`,`BC_Gen5_2`, `ElasticPool`, `Basic`,`S0`, `P2` ,`DW100c`, `DS100`.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Threat detection policy configuration. The `threat_detection_policy` block supports fields documented below.
        /// </summary>
        [Input("threatDetectionPolicy")]
        public Input<Inputs.DatabaseThreatDetectionPolicyArgs>? ThreatDetectionPolicy { get; set; }

        /// <summary>
        /// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones. This property is only settable for Premium and Business Critical databases.
        /// </summary>
        [Input("zoneRedundant")]
        public Input<bool>? ZoneRedundant { get; set; }

        public DatabaseArgs()
        {
        }
    }

    public sealed class DatabaseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time in minutes after which database is automatically paused. A value of `-1` means that automatic pause is disabled. This property is only settable for General Purpose Serverless databases.
        /// </summary>
        [Input("autoPauseDelayInMinutes")]
        public Input<int>? AutoPauseDelayInMinutes { get; set; }

        /// <summary>
        /// Specifies the collation of the database. Changing this forces a new resource to be created.
        /// </summary>
        [Input("collation")]
        public Input<string>? Collation { get; set; }

        /// <summary>
        /// The create mode of the database. Possible values are `Copy`, `Default`, `OnlineSecondary`, `PointInTimeRestore`, `Restore`, `RestoreExternalBackup`, `RestoreExternalBackupSecondary`, `RestoreLongTermRetentionBackup` and `Secondary`. 
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        /// <summary>
        /// The id of the source database to be referred to create the new database. This should only be used for databases with `create_mode` values that use another database as reference. Changing this forces a new resource to be created.
        /// </summary>
        [Input("creationSourceDatabaseId")]
        public Input<string>? CreationSourceDatabaseId { get; set; }

        /// <summary>
        /// Specifies the ID of the elastic pool containing this database. Changing this forces a new resource to be created.
        /// </summary>
        [Input("elasticPoolId")]
        public Input<string>? ElasticPoolId { get; set; }

        /// <summary>
        /// A `extended_auditing_policy` block as defined below.
        /// </summary>
        [Input("extendedAuditingPolicy")]
        public Input<Inputs.DatabaseExtendedAuditingPolicyGetArgs>? ExtendedAuditingPolicy { get; set; }

        /// <summary>
        /// Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// The max size of the database in gigabytes. 
        /// </summary>
        [Input("maxSizeGb")]
        public Input<int>? MaxSizeGb { get; set; }

        /// <summary>
        /// Minimal capacity that database will always have allocated, if not paused. This property is only settable for General Purpose Serverless databases.
        /// </summary>
        [Input("minCapacity")]
        public Input<double>? MinCapacity { get; set; }

        /// <summary>
        /// The name of the Ms SQL Database. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
        /// </summary>
        [Input("readReplicaCount")]
        public Input<int>? ReadReplicaCount { get; set; }

        /// <summary>
        /// If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
        /// </summary>
        [Input("readScale")]
        public Input<bool>? ReadScale { get; set; }

        /// <summary>
        /// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. This property is only settable for `create_mode`= `PointInTimeRestore`  databases.
        /// </summary>
        [Input("restorePointInTime")]
        public Input<string>? RestorePointInTime { get; set; }

        /// <summary>
        /// Specifies the name of the sample schema to apply when creating this database. Possible value is `AdventureWorksLT`.
        /// </summary>
        [Input("sampleName")]
        public Input<string>? SampleName { get; set; }

        /// <summary>
        /// The id of the Ms SQL Server on which to create the database. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        /// <summary>
        /// Specifies the name of the sku used by the database. Changing this forces a new resource to be created. For example, `GP_S_Gen5_2`,`HS_Gen4_1`,`BC_Gen5_2`, `ElasticPool`, `Basic`,`S0`, `P2` ,`DW100c`, `DS100`.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Threat detection policy configuration. The `threat_detection_policy` block supports fields documented below.
        /// </summary>
        [Input("threatDetectionPolicy")]
        public Input<Inputs.DatabaseThreatDetectionPolicyGetArgs>? ThreatDetectionPolicy { get; set; }

        /// <summary>
        /// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones. This property is only settable for Premium and Business Critical databases.
        /// </summary>
        [Input("zoneRedundant")]
        public Input<bool>? ZoneRedundant { get; set; }

        public DatabaseState()
        {
        }
    }
}
