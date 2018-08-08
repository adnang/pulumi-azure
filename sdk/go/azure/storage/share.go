// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Create an Azure Storage File Share.
type Share struct {
	s *pulumi.ResourceState
}

// NewShare registers a new resource with the given unique name, arguments, and options.
func NewShare(ctx *pulumi.Context,
	name string, args *ShareArgs, opts ...pulumi.ResourceOpt) (*Share, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["quota"] = nil
		inputs["resourceGroupName"] = nil
		inputs["storageAccountName"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["quota"] = args.Quota
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["storageAccountName"] = args.StorageAccountName
	}
	inputs["url"] = nil
	s, err := ctx.RegisterResource("azure:storage/share:Share", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Share{s: s}, nil
}

// GetShare gets an existing Share resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShare(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ShareState, opts ...pulumi.ResourceOpt) (*Share, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["quota"] = state.Quota
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["storageAccountName"] = state.StorageAccountName
		inputs["url"] = state.Url
	}
	s, err := ctx.ReadResource("azure:storage/share:Share", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Share{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Share) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Share) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The name of the share. Must be unique within the storage account where the share is located.
func (r *Share) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5 TB (5120 GB). Default is 5120.
func (r *Share) Quota() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["quota"])
}

// The name of the resource group in which to
// create the share. Changing this forces a new resource to be created.
func (r *Share) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Specifies the storage account in which to create the share.
// Changing this forces a new resource to be created.
func (r *Share) StorageAccountName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["storageAccountName"])
}

// The URL of the share
func (r *Share) Url() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["url"])
}

// Input properties used for looking up and filtering Share resources.
type ShareState struct {
	// The name of the share. Must be unique within the storage account where the share is located.
	Name interface{}
	// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5 TB (5120 GB). Default is 5120.
	Quota interface{}
	// The name of the resource group in which to
	// create the share. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Specifies the storage account in which to create the share.
	// Changing this forces a new resource to be created.
	StorageAccountName interface{}
	// The URL of the share
	Url interface{}
}

// The set of arguments for constructing a Share resource.
type ShareArgs struct {
	// The name of the share. Must be unique within the storage account where the share is located.
	Name interface{}
	// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5 TB (5120 GB). Default is 5120.
	Quota interface{}
	// The name of the resource group in which to
	// create the share. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Specifies the storage account in which to create the share.
	// Changing this forces a new resource to be created.
	StorageAccountName interface{}
}
