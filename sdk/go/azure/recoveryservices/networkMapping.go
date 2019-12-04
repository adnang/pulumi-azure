// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a site recovery network mapping on Azure. A network mapping decides how to translate connected netwroks when a VM is migrated from one region to another.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/recovery_network_mapping.html.markdown.
type NetworkMapping struct {
	pulumi.CustomResourceState

	// The name of the network mapping.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`

	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The id of the primary network.
	SourceNetworkId pulumi.StringOutput `pulumi:"sourceNetworkId"`

	// Specifies the ASR fabric where mapping should be created.
	SourceRecoveryFabricName pulumi.StringOutput `pulumi:"sourceRecoveryFabricName"`

	// The id of the recovery network.
	TargetNetworkId pulumi.StringOutput `pulumi:"targetNetworkId"`

	// The Azure Site Recovery fabric object corresponding to the recovery Azure region.
	TargetRecoveryFabricName pulumi.StringOutput `pulumi:"targetRecoveryFabricName"`
}

// NewNetworkMapping registers a new resource with the given unique name, arguments, and options.
func NewNetworkMapping(ctx *pulumi.Context,
	name string, args *NetworkMappingArgs, opts ...pulumi.ResourceOption) (*NetworkMapping, error) {
	if args == nil || args.RecoveryVaultName == nil {
		return nil, errors.New("missing required argument 'RecoveryVaultName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SourceNetworkId == nil {
		return nil, errors.New("missing required argument 'SourceNetworkId'")
	}
	if args == nil || args.SourceRecoveryFabricName == nil {
		return nil, errors.New("missing required argument 'SourceRecoveryFabricName'")
	}
	if args == nil || args.TargetNetworkId == nil {
		return nil, errors.New("missing required argument 'TargetNetworkId'")
	}
	if args == nil || args.TargetRecoveryFabricName == nil {
		return nil, errors.New("missing required argument 'TargetRecoveryFabricName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.RecoveryVaultName; i != nil { inputs["recoveryVaultName"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.SourceNetworkId; i != nil { inputs["sourceNetworkId"] = i.ToStringOutput() }
		if i := args.SourceRecoveryFabricName; i != nil { inputs["sourceRecoveryFabricName"] = i.ToStringOutput() }
		if i := args.TargetNetworkId; i != nil { inputs["targetNetworkId"] = i.ToStringOutput() }
		if i := args.TargetRecoveryFabricName; i != nil { inputs["targetRecoveryFabricName"] = i.ToStringOutput() }
	}
	var resource NetworkMapping
	err := ctx.RegisterResource("azure:recoveryservices/networkMapping:NetworkMapping", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkMapping gets an existing NetworkMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkMappingState, opts ...pulumi.ResourceOption) (*NetworkMapping, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.RecoveryVaultName; i != nil { inputs["recoveryVaultName"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.SourceNetworkId; i != nil { inputs["sourceNetworkId"] = i.ToStringOutput() }
		if i := state.SourceRecoveryFabricName; i != nil { inputs["sourceRecoveryFabricName"] = i.ToStringOutput() }
		if i := state.TargetNetworkId; i != nil { inputs["targetNetworkId"] = i.ToStringOutput() }
		if i := state.TargetRecoveryFabricName; i != nil { inputs["targetRecoveryFabricName"] = i.ToStringOutput() }
	}
	var resource NetworkMapping
	err := ctx.ReadResource("azure:recoveryservices/networkMapping:NetworkMapping", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkMapping resources.
type NetworkMappingState struct {
	// The name of the network mapping.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringInput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The id of the primary network.
	SourceNetworkId pulumi.StringInput `pulumi:"sourceNetworkId"`
	// Specifies the ASR fabric where mapping should be created.
	SourceRecoveryFabricName pulumi.StringInput `pulumi:"sourceRecoveryFabricName"`
	// The id of the recovery network.
	TargetNetworkId pulumi.StringInput `pulumi:"targetNetworkId"`
	// The Azure Site Recovery fabric object corresponding to the recovery Azure region.
	TargetRecoveryFabricName pulumi.StringInput `pulumi:"targetRecoveryFabricName"`
}

// The set of arguments for constructing a NetworkMapping resource.
type NetworkMappingArgs struct {
	// The name of the network mapping.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringInput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The id of the primary network.
	SourceNetworkId pulumi.StringInput `pulumi:"sourceNetworkId"`
	// Specifies the ASR fabric where mapping should be created.
	SourceRecoveryFabricName pulumi.StringInput `pulumi:"sourceRecoveryFabricName"`
	// The id of the recovery network.
	TargetNetworkId pulumi.StringInput `pulumi:"targetNetworkId"`
	// The Azure Site Recovery fabric object corresponding to the recovery Azure region.
	TargetRecoveryFabricName pulumi.StringInput `pulumi:"targetRecoveryFabricName"`
}
