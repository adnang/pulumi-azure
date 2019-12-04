// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Queue within an Azure Storage Account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_queue.html.markdown.
type Queue struct {
	pulumi.CustomResourceState

	// A mapping of MetaData which should be assigned to this Storage Queue.
	Metadata pulumi.MapOutput `pulumi:"metadata"`

	// The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the resource group in which to create the storage queue.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Metadata; i != nil { inputs["metadata"] = i.ToMapOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.StorageAccountName; i != nil { inputs["storageAccountName"] = i.ToStringOutput() }
	}
	var resource Queue
	err := ctx.RegisterResource("azure:storage/queue:Queue", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Metadata; i != nil { inputs["metadata"] = i.ToMapOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.StorageAccountName; i != nil { inputs["storageAccountName"] = i.ToStringOutput() }
	}
	var resource Queue
	err := ctx.ReadResource("azure:storage/queue:Queue", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type QueueState struct {
	// A mapping of MetaData which should be assigned to this Storage Queue.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the storage queue.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// A mapping of MetaData which should be assigned to this Storage Queue.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the storage queue.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}
