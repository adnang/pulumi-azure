// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notificationhub

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Notification Hub within a Notification Hub Namespace.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/notification_hub.html.markdown.
type Hub struct {
	pulumi.CustomResourceState

	// A `apnsCredential` block as defined below.
	ApnsCredential HubApnsCredentialOutput `pulumi:"apnsCredential"`

	// A `gcmCredential` block as defined below.
	GcmCredential HubGcmCredentialOutput `pulumi:"gcmCredential"`

	// The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// The name to use for this Notification Hub. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the Notification Hub Namespace in which to create this Notification Hub. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`

	// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewHub registers a new resource with the given unique name, arguments, and options.
func NewHub(ctx *pulumi.Context,
	name string, args *HubArgs, opts ...pulumi.ResourceOption) (*Hub, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ApnsCredential; i != nil { inputs["apnsCredential"] = i.ToHubApnsCredentialOutput() }
		if i := args.GcmCredential; i != nil { inputs["gcmCredential"] = i.ToHubGcmCredentialOutput() }
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.NamespaceName; i != nil { inputs["namespaceName"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
	}
	var resource Hub
	err := ctx.RegisterResource("azure:notificationhub/hub:Hub", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHub gets an existing Hub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubState, opts ...pulumi.ResourceOption) (*Hub, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ApnsCredential; i != nil { inputs["apnsCredential"] = i.ToHubApnsCredentialOutput() }
		if i := state.GcmCredential; i != nil { inputs["gcmCredential"] = i.ToHubGcmCredentialOutput() }
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.NamespaceName; i != nil { inputs["namespaceName"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
	}
	var resource Hub
	err := ctx.ReadResource("azure:notificationhub/hub:Hub", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hub resources.
type HubState struct {
	// A `apnsCredential` block as defined below.
	ApnsCredential HubApnsCredentialInput `pulumi:"apnsCredential"`
	// A `gcmCredential` block as defined below.
	GcmCredential HubGcmCredentialInput `pulumi:"gcmCredential"`
	// The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// The name to use for this Notification Hub. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Notification Hub Namespace in which to create this Notification Hub. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Hub resource.
type HubArgs struct {
	// A `apnsCredential` block as defined below.
	ApnsCredential HubApnsCredentialInput `pulumi:"apnsCredential"`
	// A `gcmCredential` block as defined below.
	GcmCredential HubGcmCredentialInput `pulumi:"gcmCredential"`
	// The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// The name to use for this Notification Hub. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Notification Hub Namespace in which to create this Notification Hub. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}
type HubApnsCredential struct {
	ApplicationMode string `pulumi:"applicationMode"`
	BundleId string `pulumi:"bundleId"`
	KeyId string `pulumi:"keyId"`
	TeamId string `pulumi:"teamId"`
	Token string `pulumi:"token"`
}
var hubApnsCredentialType = reflect.TypeOf((*HubApnsCredential)(nil)).Elem()

type HubApnsCredentialInput interface {
	pulumi.Input

	ToHubApnsCredentialOutput() HubApnsCredentialOutput
	ToHubApnsCredentialOutputWithContext(ctx context.Context) HubApnsCredentialOutput
}

type HubApnsCredentialArgs struct {
	ApplicationMode pulumi.StringInput `pulumi:"applicationMode"`
	BundleId pulumi.StringInput `pulumi:"bundleId"`
	KeyId pulumi.StringInput `pulumi:"keyId"`
	TeamId pulumi.StringInput `pulumi:"teamId"`
	Token pulumi.StringInput `pulumi:"token"`
}

func (HubApnsCredentialArgs) ElementType() reflect.Type {
	return hubApnsCredentialType
}

func (a HubApnsCredentialArgs) ToHubApnsCredentialOutput() HubApnsCredentialOutput {
	return pulumi.ToOutput(a).(HubApnsCredentialOutput)
}

func (a HubApnsCredentialArgs) ToHubApnsCredentialOutputWithContext(ctx context.Context) HubApnsCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, a).(HubApnsCredentialOutput)
}

type HubApnsCredentialOutput struct { *pulumi.OutputState }

func (o HubApnsCredentialOutput) ApplicationMode() pulumi.StringOutput {
	return o.Apply(func(v HubApnsCredential) string {
		return v.ApplicationMode
	}).(pulumi.StringOutput)
}

func (o HubApnsCredentialOutput) BundleId() pulumi.StringOutput {
	return o.Apply(func(v HubApnsCredential) string {
		return v.BundleId
	}).(pulumi.StringOutput)
}

func (o HubApnsCredentialOutput) KeyId() pulumi.StringOutput {
	return o.Apply(func(v HubApnsCredential) string {
		return v.KeyId
	}).(pulumi.StringOutput)
}

func (o HubApnsCredentialOutput) TeamId() pulumi.StringOutput {
	return o.Apply(func(v HubApnsCredential) string {
		return v.TeamId
	}).(pulumi.StringOutput)
}

func (o HubApnsCredentialOutput) Token() pulumi.StringOutput {
	return o.Apply(func(v HubApnsCredential) string {
		return v.Token
	}).(pulumi.StringOutput)
}

func (HubApnsCredentialOutput) ElementType() reflect.Type {
	return hubApnsCredentialType
}

func (o HubApnsCredentialOutput) ToHubApnsCredentialOutput() HubApnsCredentialOutput {
	return o
}

func (o HubApnsCredentialOutput) ToHubApnsCredentialOutputWithContext(ctx context.Context) HubApnsCredentialOutput {
	return o
}

func init() { pulumi.RegisterOutputType(HubApnsCredentialOutput{}) }

type HubGcmCredential struct {
	ApiKey string `pulumi:"apiKey"`
}
var hubGcmCredentialType = reflect.TypeOf((*HubGcmCredential)(nil)).Elem()

type HubGcmCredentialInput interface {
	pulumi.Input

	ToHubGcmCredentialOutput() HubGcmCredentialOutput
	ToHubGcmCredentialOutputWithContext(ctx context.Context) HubGcmCredentialOutput
}

type HubGcmCredentialArgs struct {
	ApiKey pulumi.StringInput `pulumi:"apiKey"`
}

func (HubGcmCredentialArgs) ElementType() reflect.Type {
	return hubGcmCredentialType
}

func (a HubGcmCredentialArgs) ToHubGcmCredentialOutput() HubGcmCredentialOutput {
	return pulumi.ToOutput(a).(HubGcmCredentialOutput)
}

func (a HubGcmCredentialArgs) ToHubGcmCredentialOutputWithContext(ctx context.Context) HubGcmCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, a).(HubGcmCredentialOutput)
}

type HubGcmCredentialOutput struct { *pulumi.OutputState }

func (o HubGcmCredentialOutput) ApiKey() pulumi.StringOutput {
	return o.Apply(func(v HubGcmCredential) string {
		return v.ApiKey
	}).(pulumi.StringOutput)
}

func (HubGcmCredentialOutput) ElementType() reflect.Type {
	return hubGcmCredentialType
}

func (o HubGcmCredentialOutput) ToHubGcmCredentialOutput() HubGcmCredentialOutput {
	return o
}

func (o HubGcmCredentialOutput) ToHubGcmCredentialOutputWithContext(ctx context.Context) HubGcmCredentialOutput {
	return o
}

func init() { pulumi.RegisterOutputType(HubGcmCredentialOutput{}) }

