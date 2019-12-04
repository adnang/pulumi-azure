// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a HDInsight Kafka Cluster.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/hdinsight_kafka_cluster.html.markdown.
type KafkaCluster struct {
	pulumi.CustomResourceState

	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringOutput `pulumi:"clusterVersion"`

	// A `componentVersion` block as defined below.
	ComponentVersion KafkaClusterComponentVersionOutput `pulumi:"componentVersion"`

	// A `gateway` block as defined below.
	Gateway KafkaClusterGatewayOutput `pulumi:"gateway"`

	// The HTTPS Connectivity Endpoint for this HDInsight Kafka Cluster.
	HttpsEndpoint pulumi.StringOutput `pulumi:"httpsEndpoint"`

	// Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// A `roles` block as defined below.
	Roles KafkaClusterRolesOutput `pulumi:"roles"`

	// The SSH Connectivity Endpoint for this HDInsight Kafka Cluster.
	SshEndpoint pulumi.StringOutput `pulumi:"sshEndpoint"`

	// One or more `storageAccount` block as defined below.
	StorageAccounts KafkaClusterStorageAccountsArrayOutput `pulumi:"storageAccounts"`

	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 KafkaClusterStorageAccountGen2Output `pulumi:"storageAccountGen2"`

	// A map of Tags which should be assigned to this HDInsight Kafka Cluster.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier pulumi.StringOutput `pulumi:"tier"`
}

// NewKafkaCluster registers a new resource with the given unique name, arguments, and options.
func NewKafkaCluster(ctx *pulumi.Context,
	name string, args *KafkaClusterArgs, opts ...pulumi.ResourceOption) (*KafkaCluster, error) {
	if args == nil || args.ClusterVersion == nil {
		return nil, errors.New("missing required argument 'ClusterVersion'")
	}
	if args == nil || args.ComponentVersion == nil {
		return nil, errors.New("missing required argument 'ComponentVersion'")
	}
	if args == nil || args.Gateway == nil {
		return nil, errors.New("missing required argument 'Gateway'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Roles == nil {
		return nil, errors.New("missing required argument 'Roles'")
	}
	if args == nil || args.Tier == nil {
		return nil, errors.New("missing required argument 'Tier'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ClusterVersion; i != nil { inputs["clusterVersion"] = i.ToStringOutput() }
		if i := args.ComponentVersion; i != nil { inputs["componentVersion"] = i.ToKafkaClusterComponentVersionOutput() }
		if i := args.Gateway; i != nil { inputs["gateway"] = i.ToKafkaClusterGatewayOutput() }
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Roles; i != nil { inputs["roles"] = i.ToKafkaClusterRolesOutput() }
		if i := args.StorageAccounts; i != nil { inputs["storageAccounts"] = i.ToKafkaClusterStorageAccountsArrayOutput() }
		if i := args.StorageAccountGen2; i != nil { inputs["storageAccountGen2"] = i.ToKafkaClusterStorageAccountGen2Output() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.Tier; i != nil { inputs["tier"] = i.ToStringOutput() }
	}
	var resource KafkaCluster
	err := ctx.RegisterResource("azure:hdinsight/kafkaCluster:KafkaCluster", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaCluster gets an existing KafkaCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaClusterState, opts ...pulumi.ResourceOption) (*KafkaCluster, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ClusterVersion; i != nil { inputs["clusterVersion"] = i.ToStringOutput() }
		if i := state.ComponentVersion; i != nil { inputs["componentVersion"] = i.ToKafkaClusterComponentVersionOutput() }
		if i := state.Gateway; i != nil { inputs["gateway"] = i.ToKafkaClusterGatewayOutput() }
		if i := state.HttpsEndpoint; i != nil { inputs["httpsEndpoint"] = i.ToStringOutput() }
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Roles; i != nil { inputs["roles"] = i.ToKafkaClusterRolesOutput() }
		if i := state.SshEndpoint; i != nil { inputs["sshEndpoint"] = i.ToStringOutput() }
		if i := state.StorageAccounts; i != nil { inputs["storageAccounts"] = i.ToKafkaClusterStorageAccountsArrayOutput() }
		if i := state.StorageAccountGen2; i != nil { inputs["storageAccountGen2"] = i.ToKafkaClusterStorageAccountGen2Output() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.Tier; i != nil { inputs["tier"] = i.ToStringOutput() }
	}
	var resource KafkaCluster
	err := ctx.ReadResource("azure:hdinsight/kafkaCluster:KafkaCluster", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaCluster resources.
type KafkaClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringInput `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion KafkaClusterComponentVersionInput `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway KafkaClusterGatewayInput `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight Kafka Cluster.
	HttpsEndpoint pulumi.StringInput `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles KafkaClusterRolesInput `pulumi:"roles"`
	// The SSH Connectivity Endpoint for this HDInsight Kafka Cluster.
	SshEndpoint pulumi.StringInput `pulumi:"sshEndpoint"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts KafkaClusterStorageAccountsArrayInput `pulumi:"storageAccounts"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 KafkaClusterStorageAccountGen2Input `pulumi:"storageAccountGen2"`
	// A map of Tags which should be assigned to this HDInsight Kafka Cluster.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier pulumi.StringInput `pulumi:"tier"`
}

// The set of arguments for constructing a KafkaCluster resource.
type KafkaClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringInput `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion KafkaClusterComponentVersionInput `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway KafkaClusterGatewayInput `pulumi:"gateway"`
	// Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles KafkaClusterRolesInput `pulumi:"roles"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts KafkaClusterStorageAccountsArrayInput `pulumi:"storageAccounts"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 KafkaClusterStorageAccountGen2Input `pulumi:"storageAccountGen2"`
	// A map of Tags which should be assigned to this HDInsight Kafka Cluster.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier pulumi.StringInput `pulumi:"tier"`
}
type KafkaClusterComponentVersion struct {
	Kafka string `pulumi:"kafka"`
}
var kafkaClusterComponentVersionType = reflect.TypeOf((*KafkaClusterComponentVersion)(nil)).Elem()

type KafkaClusterComponentVersionInput interface {
	pulumi.Input

	ToKafkaClusterComponentVersionOutput() KafkaClusterComponentVersionOutput
	ToKafkaClusterComponentVersionOutputWithContext(ctx context.Context) KafkaClusterComponentVersionOutput
}

type KafkaClusterComponentVersionArgs struct {
	Kafka pulumi.StringInput `pulumi:"kafka"`
}

func (KafkaClusterComponentVersionArgs) ElementType() reflect.Type {
	return kafkaClusterComponentVersionType
}

func (a KafkaClusterComponentVersionArgs) ToKafkaClusterComponentVersionOutput() KafkaClusterComponentVersionOutput {
	return pulumi.ToOutput(a).(KafkaClusterComponentVersionOutput)
}

func (a KafkaClusterComponentVersionArgs) ToKafkaClusterComponentVersionOutputWithContext(ctx context.Context) KafkaClusterComponentVersionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(KafkaClusterComponentVersionOutput)
}

type KafkaClusterComponentVersionOutput struct { *pulumi.OutputState }

func (o KafkaClusterComponentVersionOutput) Kafka() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterComponentVersion) string {
		return v.Kafka
	}).(pulumi.StringOutput)
}

func (KafkaClusterComponentVersionOutput) ElementType() reflect.Type {
	return kafkaClusterComponentVersionType
}

func (o KafkaClusterComponentVersionOutput) ToKafkaClusterComponentVersionOutput() KafkaClusterComponentVersionOutput {
	return o
}

func (o KafkaClusterComponentVersionOutput) ToKafkaClusterComponentVersionOutputWithContext(ctx context.Context) KafkaClusterComponentVersionOutput {
	return o
}

func init() { pulumi.RegisterOutputType(KafkaClusterComponentVersionOutput{}) }

type KafkaClusterGateway struct {
	Enabled bool `pulumi:"enabled"`
	Password string `pulumi:"password"`
	Username string `pulumi:"username"`
}
var kafkaClusterGatewayType = reflect.TypeOf((*KafkaClusterGateway)(nil)).Elem()

type KafkaClusterGatewayInput interface {
	pulumi.Input

	ToKafkaClusterGatewayOutput() KafkaClusterGatewayOutput
	ToKafkaClusterGatewayOutputWithContext(ctx context.Context) KafkaClusterGatewayOutput
}

type KafkaClusterGatewayArgs struct {
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	Password pulumi.StringInput `pulumi:"password"`
	Username pulumi.StringInput `pulumi:"username"`
}

func (KafkaClusterGatewayArgs) ElementType() reflect.Type {
	return kafkaClusterGatewayType
}

func (a KafkaClusterGatewayArgs) ToKafkaClusterGatewayOutput() KafkaClusterGatewayOutput {
	return pulumi.ToOutput(a).(KafkaClusterGatewayOutput)
}

func (a KafkaClusterGatewayArgs) ToKafkaClusterGatewayOutputWithContext(ctx context.Context) KafkaClusterGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(KafkaClusterGatewayOutput)
}

type KafkaClusterGatewayOutput struct { *pulumi.OutputState }

func (o KafkaClusterGatewayOutput) Enabled() pulumi.BoolOutput {
	return o.Apply(func(v KafkaClusterGateway) bool {
		return v.Enabled
	}).(pulumi.BoolOutput)
}

func (o KafkaClusterGatewayOutput) Password() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterGateway) string {
		return v.Password
	}).(pulumi.StringOutput)
}

func (o KafkaClusterGatewayOutput) Username() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterGateway) string {
		return v.Username
	}).(pulumi.StringOutput)
}

func (KafkaClusterGatewayOutput) ElementType() reflect.Type {
	return kafkaClusterGatewayType
}

func (o KafkaClusterGatewayOutput) ToKafkaClusterGatewayOutput() KafkaClusterGatewayOutput {
	return o
}

func (o KafkaClusterGatewayOutput) ToKafkaClusterGatewayOutputWithContext(ctx context.Context) KafkaClusterGatewayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(KafkaClusterGatewayOutput{}) }

type KafkaClusterRoles struct {
	HeadNode KafkaClusterRolesHeadNode `pulumi:"headNode"`
	WorkerNode KafkaClusterRolesWorkerNode `pulumi:"workerNode"`
	ZookeeperNode KafkaClusterRolesZookeeperNode `pulumi:"zookeeperNode"`
}
var kafkaClusterRolesType = reflect.TypeOf((*KafkaClusterRoles)(nil)).Elem()

type KafkaClusterRolesInput interface {
	pulumi.Input

	ToKafkaClusterRolesOutput() KafkaClusterRolesOutput
	ToKafkaClusterRolesOutputWithContext(ctx context.Context) KafkaClusterRolesOutput
}

type KafkaClusterRolesArgs struct {
	HeadNode KafkaClusterRolesHeadNodeInput `pulumi:"headNode"`
	WorkerNode KafkaClusterRolesWorkerNodeInput `pulumi:"workerNode"`
	ZookeeperNode KafkaClusterRolesZookeeperNodeInput `pulumi:"zookeeperNode"`
}

func (KafkaClusterRolesArgs) ElementType() reflect.Type {
	return kafkaClusterRolesType
}

func (a KafkaClusterRolesArgs) ToKafkaClusterRolesOutput() KafkaClusterRolesOutput {
	return pulumi.ToOutput(a).(KafkaClusterRolesOutput)
}

func (a KafkaClusterRolesArgs) ToKafkaClusterRolesOutputWithContext(ctx context.Context) KafkaClusterRolesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(KafkaClusterRolesOutput)
}

type KafkaClusterRolesOutput struct { *pulumi.OutputState }

func (o KafkaClusterRolesOutput) HeadNode() KafkaClusterRolesHeadNodeOutput {
	return o.Apply(func(v KafkaClusterRoles) KafkaClusterRolesHeadNode {
		return v.HeadNode
	}).(KafkaClusterRolesHeadNodeOutput)
}

func (o KafkaClusterRolesOutput) WorkerNode() KafkaClusterRolesWorkerNodeOutput {
	return o.Apply(func(v KafkaClusterRoles) KafkaClusterRolesWorkerNode {
		return v.WorkerNode
	}).(KafkaClusterRolesWorkerNodeOutput)
}

func (o KafkaClusterRolesOutput) ZookeeperNode() KafkaClusterRolesZookeeperNodeOutput {
	return o.Apply(func(v KafkaClusterRoles) KafkaClusterRolesZookeeperNode {
		return v.ZookeeperNode
	}).(KafkaClusterRolesZookeeperNodeOutput)
}

func (KafkaClusterRolesOutput) ElementType() reflect.Type {
	return kafkaClusterRolesType
}

func (o KafkaClusterRolesOutput) ToKafkaClusterRolesOutput() KafkaClusterRolesOutput {
	return o
}

func (o KafkaClusterRolesOutput) ToKafkaClusterRolesOutputWithContext(ctx context.Context) KafkaClusterRolesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(KafkaClusterRolesOutput{}) }

type KafkaClusterRolesHeadNode struct {
	Password *string `pulumi:"password"`
	SshKeys *[]string `pulumi:"sshKeys"`
	SubnetId *string `pulumi:"subnetId"`
	Username string `pulumi:"username"`
	VirtualNetworkId *string `pulumi:"virtualNetworkId"`
	VmSize string `pulumi:"vmSize"`
}
var kafkaClusterRolesHeadNodeType = reflect.TypeOf((*KafkaClusterRolesHeadNode)(nil)).Elem()

type KafkaClusterRolesHeadNodeInput interface {
	pulumi.Input

	ToKafkaClusterRolesHeadNodeOutput() KafkaClusterRolesHeadNodeOutput
	ToKafkaClusterRolesHeadNodeOutputWithContext(ctx context.Context) KafkaClusterRolesHeadNodeOutput
}

type KafkaClusterRolesHeadNodeArgs struct {
	Password pulumi.StringInput `pulumi:"password"`
	SshKeys pulumi.StringArrayInput `pulumi:"sshKeys"`
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
	Username pulumi.StringInput `pulumi:"username"`
	VirtualNetworkId pulumi.StringInput `pulumi:"virtualNetworkId"`
	VmSize pulumi.StringInput `pulumi:"vmSize"`
}

func (KafkaClusterRolesHeadNodeArgs) ElementType() reflect.Type {
	return kafkaClusterRolesHeadNodeType
}

func (a KafkaClusterRolesHeadNodeArgs) ToKafkaClusterRolesHeadNodeOutput() KafkaClusterRolesHeadNodeOutput {
	return pulumi.ToOutput(a).(KafkaClusterRolesHeadNodeOutput)
}

func (a KafkaClusterRolesHeadNodeArgs) ToKafkaClusterRolesHeadNodeOutputWithContext(ctx context.Context) KafkaClusterRolesHeadNodeOutput {
	return pulumi.ToOutputWithContext(ctx, a).(KafkaClusterRolesHeadNodeOutput)
}

type KafkaClusterRolesHeadNodeOutput struct { *pulumi.OutputState }

func (o KafkaClusterRolesHeadNodeOutput) Password() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesHeadNode) string {
		if v.Password == nil { return *new(string) } else { return *v.Password }
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesHeadNodeOutput) SshKeys() pulumi.StringArrayOutput {
	return o.Apply(func(v KafkaClusterRolesHeadNode) []string {
		if v.SshKeys == nil { return *new([]string) } else { return *v.SshKeys }
	}).(pulumi.StringArrayOutput)
}

func (o KafkaClusterRolesHeadNodeOutput) SubnetId() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesHeadNode) string {
		if v.SubnetId == nil { return *new(string) } else { return *v.SubnetId }
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesHeadNodeOutput) Username() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesHeadNode) string {
		return v.Username
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesHeadNodeOutput) VirtualNetworkId() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesHeadNode) string {
		if v.VirtualNetworkId == nil { return *new(string) } else { return *v.VirtualNetworkId }
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesHeadNodeOutput) VmSize() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesHeadNode) string {
		return v.VmSize
	}).(pulumi.StringOutput)
}

func (KafkaClusterRolesHeadNodeOutput) ElementType() reflect.Type {
	return kafkaClusterRolesHeadNodeType
}

func (o KafkaClusterRolesHeadNodeOutput) ToKafkaClusterRolesHeadNodeOutput() KafkaClusterRolesHeadNodeOutput {
	return o
}

func (o KafkaClusterRolesHeadNodeOutput) ToKafkaClusterRolesHeadNodeOutputWithContext(ctx context.Context) KafkaClusterRolesHeadNodeOutput {
	return o
}

func init() { pulumi.RegisterOutputType(KafkaClusterRolesHeadNodeOutput{}) }

type KafkaClusterRolesWorkerNode struct {
	MinInstanceCount *int `pulumi:"minInstanceCount"`
	NumberOfDisksPerNode int `pulumi:"numberOfDisksPerNode"`
	Password *string `pulumi:"password"`
	SshKeys *[]string `pulumi:"sshKeys"`
	SubnetId *string `pulumi:"subnetId"`
	TargetInstanceCount int `pulumi:"targetInstanceCount"`
	Username string `pulumi:"username"`
	VirtualNetworkId *string `pulumi:"virtualNetworkId"`
	VmSize string `pulumi:"vmSize"`
}
var kafkaClusterRolesWorkerNodeType = reflect.TypeOf((*KafkaClusterRolesWorkerNode)(nil)).Elem()

type KafkaClusterRolesWorkerNodeInput interface {
	pulumi.Input

	ToKafkaClusterRolesWorkerNodeOutput() KafkaClusterRolesWorkerNodeOutput
	ToKafkaClusterRolesWorkerNodeOutputWithContext(ctx context.Context) KafkaClusterRolesWorkerNodeOutput
}

type KafkaClusterRolesWorkerNodeArgs struct {
	MinInstanceCount pulumi.IntInput `pulumi:"minInstanceCount"`
	NumberOfDisksPerNode pulumi.IntInput `pulumi:"numberOfDisksPerNode"`
	Password pulumi.StringInput `pulumi:"password"`
	SshKeys pulumi.StringArrayInput `pulumi:"sshKeys"`
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
	TargetInstanceCount pulumi.IntInput `pulumi:"targetInstanceCount"`
	Username pulumi.StringInput `pulumi:"username"`
	VirtualNetworkId pulumi.StringInput `pulumi:"virtualNetworkId"`
	VmSize pulumi.StringInput `pulumi:"vmSize"`
}

func (KafkaClusterRolesWorkerNodeArgs) ElementType() reflect.Type {
	return kafkaClusterRolesWorkerNodeType
}

func (a KafkaClusterRolesWorkerNodeArgs) ToKafkaClusterRolesWorkerNodeOutput() KafkaClusterRolesWorkerNodeOutput {
	return pulumi.ToOutput(a).(KafkaClusterRolesWorkerNodeOutput)
}

func (a KafkaClusterRolesWorkerNodeArgs) ToKafkaClusterRolesWorkerNodeOutputWithContext(ctx context.Context) KafkaClusterRolesWorkerNodeOutput {
	return pulumi.ToOutputWithContext(ctx, a).(KafkaClusterRolesWorkerNodeOutput)
}

type KafkaClusterRolesWorkerNodeOutput struct { *pulumi.OutputState }

func (o KafkaClusterRolesWorkerNodeOutput) MinInstanceCount() pulumi.IntOutput {
	return o.Apply(func(v KafkaClusterRolesWorkerNode) int {
		if v.MinInstanceCount == nil { return *new(int) } else { return *v.MinInstanceCount }
	}).(pulumi.IntOutput)
}

func (o KafkaClusterRolesWorkerNodeOutput) NumberOfDisksPerNode() pulumi.IntOutput {
	return o.Apply(func(v KafkaClusterRolesWorkerNode) int {
		return v.NumberOfDisksPerNode
	}).(pulumi.IntOutput)
}

func (o KafkaClusterRolesWorkerNodeOutput) Password() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesWorkerNode) string {
		if v.Password == nil { return *new(string) } else { return *v.Password }
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesWorkerNodeOutput) SshKeys() pulumi.StringArrayOutput {
	return o.Apply(func(v KafkaClusterRolesWorkerNode) []string {
		if v.SshKeys == nil { return *new([]string) } else { return *v.SshKeys }
	}).(pulumi.StringArrayOutput)
}

func (o KafkaClusterRolesWorkerNodeOutput) SubnetId() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesWorkerNode) string {
		if v.SubnetId == nil { return *new(string) } else { return *v.SubnetId }
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesWorkerNodeOutput) TargetInstanceCount() pulumi.IntOutput {
	return o.Apply(func(v KafkaClusterRolesWorkerNode) int {
		return v.TargetInstanceCount
	}).(pulumi.IntOutput)
}

func (o KafkaClusterRolesWorkerNodeOutput) Username() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesWorkerNode) string {
		return v.Username
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesWorkerNodeOutput) VirtualNetworkId() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesWorkerNode) string {
		if v.VirtualNetworkId == nil { return *new(string) } else { return *v.VirtualNetworkId }
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesWorkerNodeOutput) VmSize() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesWorkerNode) string {
		return v.VmSize
	}).(pulumi.StringOutput)
}

func (KafkaClusterRolesWorkerNodeOutput) ElementType() reflect.Type {
	return kafkaClusterRolesWorkerNodeType
}

func (o KafkaClusterRolesWorkerNodeOutput) ToKafkaClusterRolesWorkerNodeOutput() KafkaClusterRolesWorkerNodeOutput {
	return o
}

func (o KafkaClusterRolesWorkerNodeOutput) ToKafkaClusterRolesWorkerNodeOutputWithContext(ctx context.Context) KafkaClusterRolesWorkerNodeOutput {
	return o
}

func init() { pulumi.RegisterOutputType(KafkaClusterRolesWorkerNodeOutput{}) }

type KafkaClusterRolesZookeeperNode struct {
	Password *string `pulumi:"password"`
	SshKeys *[]string `pulumi:"sshKeys"`
	SubnetId *string `pulumi:"subnetId"`
	Username string `pulumi:"username"`
	VirtualNetworkId *string `pulumi:"virtualNetworkId"`
	VmSize string `pulumi:"vmSize"`
}
var kafkaClusterRolesZookeeperNodeType = reflect.TypeOf((*KafkaClusterRolesZookeeperNode)(nil)).Elem()

type KafkaClusterRolesZookeeperNodeInput interface {
	pulumi.Input

	ToKafkaClusterRolesZookeeperNodeOutput() KafkaClusterRolesZookeeperNodeOutput
	ToKafkaClusterRolesZookeeperNodeOutputWithContext(ctx context.Context) KafkaClusterRolesZookeeperNodeOutput
}

type KafkaClusterRolesZookeeperNodeArgs struct {
	Password pulumi.StringInput `pulumi:"password"`
	SshKeys pulumi.StringArrayInput `pulumi:"sshKeys"`
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
	Username pulumi.StringInput `pulumi:"username"`
	VirtualNetworkId pulumi.StringInput `pulumi:"virtualNetworkId"`
	VmSize pulumi.StringInput `pulumi:"vmSize"`
}

func (KafkaClusterRolesZookeeperNodeArgs) ElementType() reflect.Type {
	return kafkaClusterRolesZookeeperNodeType
}

func (a KafkaClusterRolesZookeeperNodeArgs) ToKafkaClusterRolesZookeeperNodeOutput() KafkaClusterRolesZookeeperNodeOutput {
	return pulumi.ToOutput(a).(KafkaClusterRolesZookeeperNodeOutput)
}

func (a KafkaClusterRolesZookeeperNodeArgs) ToKafkaClusterRolesZookeeperNodeOutputWithContext(ctx context.Context) KafkaClusterRolesZookeeperNodeOutput {
	return pulumi.ToOutputWithContext(ctx, a).(KafkaClusterRolesZookeeperNodeOutput)
}

type KafkaClusterRolesZookeeperNodeOutput struct { *pulumi.OutputState }

func (o KafkaClusterRolesZookeeperNodeOutput) Password() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesZookeeperNode) string {
		if v.Password == nil { return *new(string) } else { return *v.Password }
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesZookeeperNodeOutput) SshKeys() pulumi.StringArrayOutput {
	return o.Apply(func(v KafkaClusterRolesZookeeperNode) []string {
		if v.SshKeys == nil { return *new([]string) } else { return *v.SshKeys }
	}).(pulumi.StringArrayOutput)
}

func (o KafkaClusterRolesZookeeperNodeOutput) SubnetId() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesZookeeperNode) string {
		if v.SubnetId == nil { return *new(string) } else { return *v.SubnetId }
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesZookeeperNodeOutput) Username() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesZookeeperNode) string {
		return v.Username
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesZookeeperNodeOutput) VirtualNetworkId() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesZookeeperNode) string {
		if v.VirtualNetworkId == nil { return *new(string) } else { return *v.VirtualNetworkId }
	}).(pulumi.StringOutput)
}

func (o KafkaClusterRolesZookeeperNodeOutput) VmSize() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterRolesZookeeperNode) string {
		return v.VmSize
	}).(pulumi.StringOutput)
}

func (KafkaClusterRolesZookeeperNodeOutput) ElementType() reflect.Type {
	return kafkaClusterRolesZookeeperNodeType
}

func (o KafkaClusterRolesZookeeperNodeOutput) ToKafkaClusterRolesZookeeperNodeOutput() KafkaClusterRolesZookeeperNodeOutput {
	return o
}

func (o KafkaClusterRolesZookeeperNodeOutput) ToKafkaClusterRolesZookeeperNodeOutputWithContext(ctx context.Context) KafkaClusterRolesZookeeperNodeOutput {
	return o
}

func init() { pulumi.RegisterOutputType(KafkaClusterRolesZookeeperNodeOutput{}) }

type KafkaClusterStorageAccountGen2 struct {
	FilesystemId string `pulumi:"filesystemId"`
	IsDefault bool `pulumi:"isDefault"`
	ManagedIdentityResourceId string `pulumi:"managedIdentityResourceId"`
	StorageResourceId string `pulumi:"storageResourceId"`
}
var kafkaClusterStorageAccountGen2Type = reflect.TypeOf((*KafkaClusterStorageAccountGen2)(nil)).Elem()

type KafkaClusterStorageAccountGen2Input interface {
	pulumi.Input

	ToKafkaClusterStorageAccountGen2Output() KafkaClusterStorageAccountGen2Output
	ToKafkaClusterStorageAccountGen2OutputWithContext(ctx context.Context) KafkaClusterStorageAccountGen2Output
}

type KafkaClusterStorageAccountGen2Args struct {
	FilesystemId pulumi.StringInput `pulumi:"filesystemId"`
	IsDefault pulumi.BoolInput `pulumi:"isDefault"`
	ManagedIdentityResourceId pulumi.StringInput `pulumi:"managedIdentityResourceId"`
	StorageResourceId pulumi.StringInput `pulumi:"storageResourceId"`
}

func (KafkaClusterStorageAccountGen2Args) ElementType() reflect.Type {
	return kafkaClusterStorageAccountGen2Type
}

func (a KafkaClusterStorageAccountGen2Args) ToKafkaClusterStorageAccountGen2Output() KafkaClusterStorageAccountGen2Output {
	return pulumi.ToOutput(a).(KafkaClusterStorageAccountGen2Output)
}

func (a KafkaClusterStorageAccountGen2Args) ToKafkaClusterStorageAccountGen2OutputWithContext(ctx context.Context) KafkaClusterStorageAccountGen2Output {
	return pulumi.ToOutputWithContext(ctx, a).(KafkaClusterStorageAccountGen2Output)
}

type KafkaClusterStorageAccountGen2Output struct { *pulumi.OutputState }

func (o KafkaClusterStorageAccountGen2Output) FilesystemId() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterStorageAccountGen2) string {
		return v.FilesystemId
	}).(pulumi.StringOutput)
}

func (o KafkaClusterStorageAccountGen2Output) IsDefault() pulumi.BoolOutput {
	return o.Apply(func(v KafkaClusterStorageAccountGen2) bool {
		return v.IsDefault
	}).(pulumi.BoolOutput)
}

func (o KafkaClusterStorageAccountGen2Output) ManagedIdentityResourceId() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterStorageAccountGen2) string {
		return v.ManagedIdentityResourceId
	}).(pulumi.StringOutput)
}

func (o KafkaClusterStorageAccountGen2Output) StorageResourceId() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterStorageAccountGen2) string {
		return v.StorageResourceId
	}).(pulumi.StringOutput)
}

func (KafkaClusterStorageAccountGen2Output) ElementType() reflect.Type {
	return kafkaClusterStorageAccountGen2Type
}

func (o KafkaClusterStorageAccountGen2Output) ToKafkaClusterStorageAccountGen2Output() KafkaClusterStorageAccountGen2Output {
	return o
}

func (o KafkaClusterStorageAccountGen2Output) ToKafkaClusterStorageAccountGen2OutputWithContext(ctx context.Context) KafkaClusterStorageAccountGen2Output {
	return o
}

func init() { pulumi.RegisterOutputType(KafkaClusterStorageAccountGen2Output{}) }

type KafkaClusterStorageAccounts struct {
	IsDefault bool `pulumi:"isDefault"`
	StorageAccountKey string `pulumi:"storageAccountKey"`
	StorageContainerId string `pulumi:"storageContainerId"`
}
var kafkaClusterStorageAccountsType = reflect.TypeOf((*KafkaClusterStorageAccounts)(nil)).Elem()

type KafkaClusterStorageAccountsInput interface {
	pulumi.Input

	ToKafkaClusterStorageAccountsOutput() KafkaClusterStorageAccountsOutput
	ToKafkaClusterStorageAccountsOutputWithContext(ctx context.Context) KafkaClusterStorageAccountsOutput
}

type KafkaClusterStorageAccountsArgs struct {
	IsDefault pulumi.BoolInput `pulumi:"isDefault"`
	StorageAccountKey pulumi.StringInput `pulumi:"storageAccountKey"`
	StorageContainerId pulumi.StringInput `pulumi:"storageContainerId"`
}

func (KafkaClusterStorageAccountsArgs) ElementType() reflect.Type {
	return kafkaClusterStorageAccountsType
}

func (a KafkaClusterStorageAccountsArgs) ToKafkaClusterStorageAccountsOutput() KafkaClusterStorageAccountsOutput {
	return pulumi.ToOutput(a).(KafkaClusterStorageAccountsOutput)
}

func (a KafkaClusterStorageAccountsArgs) ToKafkaClusterStorageAccountsOutputWithContext(ctx context.Context) KafkaClusterStorageAccountsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(KafkaClusterStorageAccountsOutput)
}

type KafkaClusterStorageAccountsOutput struct { *pulumi.OutputState }

func (o KafkaClusterStorageAccountsOutput) IsDefault() pulumi.BoolOutput {
	return o.Apply(func(v KafkaClusterStorageAccounts) bool {
		return v.IsDefault
	}).(pulumi.BoolOutput)
}

func (o KafkaClusterStorageAccountsOutput) StorageAccountKey() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterStorageAccounts) string {
		return v.StorageAccountKey
	}).(pulumi.StringOutput)
}

func (o KafkaClusterStorageAccountsOutput) StorageContainerId() pulumi.StringOutput {
	return o.Apply(func(v KafkaClusterStorageAccounts) string {
		return v.StorageContainerId
	}).(pulumi.StringOutput)
}

func (KafkaClusterStorageAccountsOutput) ElementType() reflect.Type {
	return kafkaClusterStorageAccountsType
}

func (o KafkaClusterStorageAccountsOutput) ToKafkaClusterStorageAccountsOutput() KafkaClusterStorageAccountsOutput {
	return o
}

func (o KafkaClusterStorageAccountsOutput) ToKafkaClusterStorageAccountsOutputWithContext(ctx context.Context) KafkaClusterStorageAccountsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(KafkaClusterStorageAccountsOutput{}) }

var kafkaClusterStorageAccountsArrayType = reflect.TypeOf((*[]KafkaClusterStorageAccounts)(nil)).Elem()

type KafkaClusterStorageAccountsArrayInput interface {
	pulumi.Input

	ToKafkaClusterStorageAccountsArrayOutput() KafkaClusterStorageAccountsArrayOutput
	ToKafkaClusterStorageAccountsArrayOutputWithContext(ctx context.Context) KafkaClusterStorageAccountsArrayOutput
}

type KafkaClusterStorageAccountsArrayArgs []KafkaClusterStorageAccountsInput

func (KafkaClusterStorageAccountsArrayArgs) ElementType() reflect.Type {
	return kafkaClusterStorageAccountsArrayType
}

func (a KafkaClusterStorageAccountsArrayArgs) ToKafkaClusterStorageAccountsArrayOutput() KafkaClusterStorageAccountsArrayOutput {
	return pulumi.ToOutput(a).(KafkaClusterStorageAccountsArrayOutput)
}

func (a KafkaClusterStorageAccountsArrayArgs) ToKafkaClusterStorageAccountsArrayOutputWithContext(ctx context.Context) KafkaClusterStorageAccountsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(KafkaClusterStorageAccountsArrayOutput)
}

type KafkaClusterStorageAccountsArrayOutput struct { *pulumi.OutputState }

func (o KafkaClusterStorageAccountsArrayOutput) Index(i pulumi.IntInput) KafkaClusterStorageAccountsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) KafkaClusterStorageAccounts {
		return vs[0].([]KafkaClusterStorageAccounts)[vs[1].(int)]
	}).(KafkaClusterStorageAccountsOutput)
}

func (KafkaClusterStorageAccountsArrayOutput) ElementType() reflect.Type {
	return kafkaClusterStorageAccountsArrayType
}

func (o KafkaClusterStorageAccountsArrayOutput) ToKafkaClusterStorageAccountsArrayOutput() KafkaClusterStorageAccountsArrayOutput {
	return o
}

func (o KafkaClusterStorageAccountsArrayOutput) ToKafkaClusterStorageAccountsArrayOutputWithContext(ctx context.Context) KafkaClusterStorageAccountsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(KafkaClusterStorageAccountsArrayOutput{}) }

