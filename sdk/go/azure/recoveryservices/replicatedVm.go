// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Azure recovery replicated vms (Azure to Azure). An replicated VM keeps a copiously updated image of the vm in another region in order to be able to start the VM in that region in case of a disaster. 
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/recovery_replicated_vm.html.markdown.
type ReplicatedVm struct {
	pulumi.CustomResourceState

	// One or more `managedDisk` block.
	ManagedDisks ReplicatedVmManagedDisksArrayOutput `pulumi:"managedDisks"`

	// The name of the network mapping.
	Name pulumi.StringOutput `pulumi:"name"`

	RecoveryReplicationPolicyId pulumi.StringOutput `pulumi:"recoveryReplicationPolicyId"`

	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`

	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName pulumi.StringOutput `pulumi:"sourceRecoveryFabricName"`

	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName pulumi.StringOutput `pulumi:"sourceRecoveryProtectionContainerName"`

	// Id of the VM to replicate
	SourceVmId pulumi.StringOutput `pulumi:"sourceVmId"`

	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId pulumi.StringOutput `pulumi:"targetAvailabilitySetId"`

	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId pulumi.StringOutput `pulumi:"targetRecoveryFabricId"`

	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId pulumi.StringOutput `pulumi:"targetRecoveryProtectionContainerId"`

	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId pulumi.StringOutput `pulumi:"targetResourceGroupId"`
}

// NewReplicatedVm registers a new resource with the given unique name, arguments, and options.
func NewReplicatedVm(ctx *pulumi.Context,
	name string, args *ReplicatedVmArgs, opts ...pulumi.ResourceOption) (*ReplicatedVm, error) {
	if args == nil || args.RecoveryReplicationPolicyId == nil {
		return nil, errors.New("missing required argument 'RecoveryReplicationPolicyId'")
	}
	if args == nil || args.RecoveryVaultName == nil {
		return nil, errors.New("missing required argument 'RecoveryVaultName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SourceRecoveryFabricName == nil {
		return nil, errors.New("missing required argument 'SourceRecoveryFabricName'")
	}
	if args == nil || args.SourceRecoveryProtectionContainerName == nil {
		return nil, errors.New("missing required argument 'SourceRecoveryProtectionContainerName'")
	}
	if args == nil || args.SourceVmId == nil {
		return nil, errors.New("missing required argument 'SourceVmId'")
	}
	if args == nil || args.TargetRecoveryFabricId == nil {
		return nil, errors.New("missing required argument 'TargetRecoveryFabricId'")
	}
	if args == nil || args.TargetRecoveryProtectionContainerId == nil {
		return nil, errors.New("missing required argument 'TargetRecoveryProtectionContainerId'")
	}
	if args == nil || args.TargetResourceGroupId == nil {
		return nil, errors.New("missing required argument 'TargetResourceGroupId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ManagedDisks; i != nil { inputs["managedDisks"] = i.ToReplicatedVmManagedDisksArrayOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.RecoveryReplicationPolicyId; i != nil { inputs["recoveryReplicationPolicyId"] = i.ToStringOutput() }
		if i := args.RecoveryVaultName; i != nil { inputs["recoveryVaultName"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.SourceRecoveryFabricName; i != nil { inputs["sourceRecoveryFabricName"] = i.ToStringOutput() }
		if i := args.SourceRecoveryProtectionContainerName; i != nil { inputs["sourceRecoveryProtectionContainerName"] = i.ToStringOutput() }
		if i := args.SourceVmId; i != nil { inputs["sourceVmId"] = i.ToStringOutput() }
		if i := args.TargetAvailabilitySetId; i != nil { inputs["targetAvailabilitySetId"] = i.ToStringOutput() }
		if i := args.TargetRecoveryFabricId; i != nil { inputs["targetRecoveryFabricId"] = i.ToStringOutput() }
		if i := args.TargetRecoveryProtectionContainerId; i != nil { inputs["targetRecoveryProtectionContainerId"] = i.ToStringOutput() }
		if i := args.TargetResourceGroupId; i != nil { inputs["targetResourceGroupId"] = i.ToStringOutput() }
	}
	var resource ReplicatedVm
	err := ctx.RegisterResource("azure:recoveryservices/replicatedVm:ReplicatedVm", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicatedVm gets an existing ReplicatedVm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicatedVm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicatedVmState, opts ...pulumi.ResourceOption) (*ReplicatedVm, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ManagedDisks; i != nil { inputs["managedDisks"] = i.ToReplicatedVmManagedDisksArrayOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.RecoveryReplicationPolicyId; i != nil { inputs["recoveryReplicationPolicyId"] = i.ToStringOutput() }
		if i := state.RecoveryVaultName; i != nil { inputs["recoveryVaultName"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.SourceRecoveryFabricName; i != nil { inputs["sourceRecoveryFabricName"] = i.ToStringOutput() }
		if i := state.SourceRecoveryProtectionContainerName; i != nil { inputs["sourceRecoveryProtectionContainerName"] = i.ToStringOutput() }
		if i := state.SourceVmId; i != nil { inputs["sourceVmId"] = i.ToStringOutput() }
		if i := state.TargetAvailabilitySetId; i != nil { inputs["targetAvailabilitySetId"] = i.ToStringOutput() }
		if i := state.TargetRecoveryFabricId; i != nil { inputs["targetRecoveryFabricId"] = i.ToStringOutput() }
		if i := state.TargetRecoveryProtectionContainerId; i != nil { inputs["targetRecoveryProtectionContainerId"] = i.ToStringOutput() }
		if i := state.TargetResourceGroupId; i != nil { inputs["targetResourceGroupId"] = i.ToStringOutput() }
	}
	var resource ReplicatedVm
	err := ctx.ReadResource("azure:recoveryservices/replicatedVm:ReplicatedVm", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicatedVm resources.
type ReplicatedVmState struct {
	// One or more `managedDisk` block.
	ManagedDisks ReplicatedVmManagedDisksArrayInput `pulumi:"managedDisks"`
	// The name of the network mapping.
	Name pulumi.StringInput `pulumi:"name"`
	RecoveryReplicationPolicyId pulumi.StringInput `pulumi:"recoveryReplicationPolicyId"`
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringInput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName pulumi.StringInput `pulumi:"sourceRecoveryFabricName"`
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName pulumi.StringInput `pulumi:"sourceRecoveryProtectionContainerName"`
	// Id of the VM to replicate
	SourceVmId pulumi.StringInput `pulumi:"sourceVmId"`
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId pulumi.StringInput `pulumi:"targetAvailabilitySetId"`
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId pulumi.StringInput `pulumi:"targetRecoveryFabricId"`
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId pulumi.StringInput `pulumi:"targetRecoveryProtectionContainerId"`
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId pulumi.StringInput `pulumi:"targetResourceGroupId"`
}

// The set of arguments for constructing a ReplicatedVm resource.
type ReplicatedVmArgs struct {
	// One or more `managedDisk` block.
	ManagedDisks ReplicatedVmManagedDisksArrayInput `pulumi:"managedDisks"`
	// The name of the network mapping.
	Name pulumi.StringInput `pulumi:"name"`
	RecoveryReplicationPolicyId pulumi.StringInput `pulumi:"recoveryReplicationPolicyId"`
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringInput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName pulumi.StringInput `pulumi:"sourceRecoveryFabricName"`
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName pulumi.StringInput `pulumi:"sourceRecoveryProtectionContainerName"`
	// Id of the VM to replicate
	SourceVmId pulumi.StringInput `pulumi:"sourceVmId"`
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId pulumi.StringInput `pulumi:"targetAvailabilitySetId"`
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId pulumi.StringInput `pulumi:"targetRecoveryFabricId"`
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId pulumi.StringInput `pulumi:"targetRecoveryProtectionContainerId"`
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId pulumi.StringInput `pulumi:"targetResourceGroupId"`
}
type ReplicatedVmManagedDisks struct {
	DiskId string `pulumi:"diskId"`
	StagingStorageAccountId string `pulumi:"stagingStorageAccountId"`
	TargetDiskType string `pulumi:"targetDiskType"`
	TargetReplicaDiskType string `pulumi:"targetReplicaDiskType"`
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId string `pulumi:"targetResourceGroupId"`
}
var replicatedVmManagedDisksType = reflect.TypeOf((*ReplicatedVmManagedDisks)(nil)).Elem()

type ReplicatedVmManagedDisksInput interface {
	pulumi.Input

	ToReplicatedVmManagedDisksOutput() ReplicatedVmManagedDisksOutput
	ToReplicatedVmManagedDisksOutputWithContext(ctx context.Context) ReplicatedVmManagedDisksOutput
}

type ReplicatedVmManagedDisksArgs struct {
	DiskId pulumi.StringInput `pulumi:"diskId"`
	StagingStorageAccountId pulumi.StringInput `pulumi:"stagingStorageAccountId"`
	TargetDiskType pulumi.StringInput `pulumi:"targetDiskType"`
	TargetReplicaDiskType pulumi.StringInput `pulumi:"targetReplicaDiskType"`
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId pulumi.StringInput `pulumi:"targetResourceGroupId"`
}

func (ReplicatedVmManagedDisksArgs) ElementType() reflect.Type {
	return replicatedVmManagedDisksType
}

func (a ReplicatedVmManagedDisksArgs) ToReplicatedVmManagedDisksOutput() ReplicatedVmManagedDisksOutput {
	return pulumi.ToOutput(a).(ReplicatedVmManagedDisksOutput)
}

func (a ReplicatedVmManagedDisksArgs) ToReplicatedVmManagedDisksOutputWithContext(ctx context.Context) ReplicatedVmManagedDisksOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ReplicatedVmManagedDisksOutput)
}

type ReplicatedVmManagedDisksOutput struct { *pulumi.OutputState }

func (o ReplicatedVmManagedDisksOutput) DiskId() pulumi.StringOutput {
	return o.Apply(func(v ReplicatedVmManagedDisks) string {
		return v.DiskId
	}).(pulumi.StringOutput)
}

func (o ReplicatedVmManagedDisksOutput) StagingStorageAccountId() pulumi.StringOutput {
	return o.Apply(func(v ReplicatedVmManagedDisks) string {
		return v.StagingStorageAccountId
	}).(pulumi.StringOutput)
}

func (o ReplicatedVmManagedDisksOutput) TargetDiskType() pulumi.StringOutput {
	return o.Apply(func(v ReplicatedVmManagedDisks) string {
		return v.TargetDiskType
	}).(pulumi.StringOutput)
}

func (o ReplicatedVmManagedDisksOutput) TargetReplicaDiskType() pulumi.StringOutput {
	return o.Apply(func(v ReplicatedVmManagedDisks) string {
		return v.TargetReplicaDiskType
	}).(pulumi.StringOutput)
}

// Id of resource group where the VM should be created when a failover is done.
func (o ReplicatedVmManagedDisksOutput) TargetResourceGroupId() pulumi.StringOutput {
	return o.Apply(func(v ReplicatedVmManagedDisks) string {
		return v.TargetResourceGroupId
	}).(pulumi.StringOutput)
}

func (ReplicatedVmManagedDisksOutput) ElementType() reflect.Type {
	return replicatedVmManagedDisksType
}

func (o ReplicatedVmManagedDisksOutput) ToReplicatedVmManagedDisksOutput() ReplicatedVmManagedDisksOutput {
	return o
}

func (o ReplicatedVmManagedDisksOutput) ToReplicatedVmManagedDisksOutputWithContext(ctx context.Context) ReplicatedVmManagedDisksOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ReplicatedVmManagedDisksOutput{}) }

var replicatedVmManagedDisksArrayType = reflect.TypeOf((*[]ReplicatedVmManagedDisks)(nil)).Elem()

type ReplicatedVmManagedDisksArrayInput interface {
	pulumi.Input

	ToReplicatedVmManagedDisksArrayOutput() ReplicatedVmManagedDisksArrayOutput
	ToReplicatedVmManagedDisksArrayOutputWithContext(ctx context.Context) ReplicatedVmManagedDisksArrayOutput
}

type ReplicatedVmManagedDisksArrayArgs []ReplicatedVmManagedDisksInput

func (ReplicatedVmManagedDisksArrayArgs) ElementType() reflect.Type {
	return replicatedVmManagedDisksArrayType
}

func (a ReplicatedVmManagedDisksArrayArgs) ToReplicatedVmManagedDisksArrayOutput() ReplicatedVmManagedDisksArrayOutput {
	return pulumi.ToOutput(a).(ReplicatedVmManagedDisksArrayOutput)
}

func (a ReplicatedVmManagedDisksArrayArgs) ToReplicatedVmManagedDisksArrayOutputWithContext(ctx context.Context) ReplicatedVmManagedDisksArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ReplicatedVmManagedDisksArrayOutput)
}

type ReplicatedVmManagedDisksArrayOutput struct { *pulumi.OutputState }

func (o ReplicatedVmManagedDisksArrayOutput) Index(i pulumi.IntInput) ReplicatedVmManagedDisksOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) ReplicatedVmManagedDisks {
		return vs[0].([]ReplicatedVmManagedDisks)[vs[1].(int)]
	}).(ReplicatedVmManagedDisksOutput)
}

func (ReplicatedVmManagedDisksArrayOutput) ElementType() reflect.Type {
	return replicatedVmManagedDisksArrayType
}

func (o ReplicatedVmManagedDisksArrayOutput) ToReplicatedVmManagedDisksArrayOutput() ReplicatedVmManagedDisksArrayOutput {
	return o
}

func (o ReplicatedVmManagedDisksArrayOutput) ToReplicatedVmManagedDisksArrayOutputWithContext(ctx context.Context) ReplicatedVmManagedDisksArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ReplicatedVmManagedDisksArrayOutput{}) }

