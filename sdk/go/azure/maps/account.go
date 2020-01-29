// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package maps

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Azure Maps Account.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/maps_account.html.markdown.
type Account struct {
	pulumi.CustomResourceState

	// The name of the Azure Maps Account. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary key used to authenticate and authorize access to the Maps REST APIs.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The secondary key used to authenticate and authorize access to the Maps REST APIs.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// The sku of the Azure Maps Account. Possible values are `s0` and `s1`.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A mapping of tags to assign to the Azure Maps Account.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A unique identifier for the Maps Account.
	XMsClientId pulumi.StringOutput `pulumi:"xMsClientId"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SkuName == nil {
		return nil, errors.New("missing required argument 'SkuName'")
	}
	if args == nil {
		args = &AccountArgs{}
	}
	var resource Account
	err := ctx.RegisterResource("azure:maps/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure:maps/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The name of the Azure Maps Account. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The primary key used to authenticate and authorize access to the Maps REST APIs.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The secondary key used to authenticate and authorize access to the Maps REST APIs.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// The sku of the Azure Maps Account. Possible values are `s0` and `s1`.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags to assign to the Azure Maps Account.
	Tags map[string]string `pulumi:"tags"`
	// A unique identifier for the Maps Account.
	XMsClientId *string `pulumi:"xMsClientId"`
}

type AccountState struct {
	// The name of the Azure Maps Account. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The primary key used to authenticate and authorize access to the Maps REST APIs.
	PrimaryAccessKey pulumi.StringPtrInput
	// The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The secondary key used to authenticate and authorize access to the Maps REST APIs.
	SecondaryAccessKey pulumi.StringPtrInput
	// The sku of the Azure Maps Account. Possible values are `s0` and `s1`.
	SkuName pulumi.StringPtrInput
	// A mapping of tags to assign to the Azure Maps Account.
	Tags pulumi.StringMapInput
	// A unique identifier for the Maps Account.
	XMsClientId pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the Azure Maps Account. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku of the Azure Maps Account. Possible values are `s0` and `s1`.
	SkuName string `pulumi:"skuName"`
	// A mapping of tags to assign to the Azure Maps Account.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the Azure Maps Account. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The sku of the Azure Maps Account. Possible values are `s0` and `s1`.
	SkuName pulumi.StringInput
	// A mapping of tags to assign to the Azure Maps Account.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

