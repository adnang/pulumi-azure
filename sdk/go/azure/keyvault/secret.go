// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Key Vault Secret.
// 
// > **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
type Secret struct {
	s *pulumi.ResourceState
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOpt) (*Secret, error) {
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil || args.VaultUri == nil {
		return nil, errors.New("missing required argument 'VaultUri'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["contentType"] = nil
		inputs["name"] = nil
		inputs["tags"] = nil
		inputs["value"] = nil
		inputs["vaultUri"] = nil
	} else {
		inputs["contentType"] = args.ContentType
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
		inputs["value"] = args.Value
		inputs["vaultUri"] = args.VaultUri
	}
	inputs["version"] = nil
	s, err := ctx.RegisterResource("azure:keyvault/secret:Secret", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Secret{s: s}, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecretState, opts ...pulumi.ResourceOpt) (*Secret, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["contentType"] = state.ContentType
		inputs["name"] = state.Name
		inputs["tags"] = state.Tags
		inputs["value"] = state.Value
		inputs["vaultUri"] = state.VaultUri
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("azure:keyvault/secret:Secret", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Secret{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Secret) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Secret) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the content type for the Key Vault Secret.
func (r *Secret) ContentType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["contentType"])
}

// Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
func (r *Secret) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A mapping of tags to assign to the resource.
func (r *Secret) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Specifies the value of the Key Vault Secret.
func (r *Secret) Value() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["value"])
}

// Specifies the URI used to access the Key Vault instance, available on the `azurerm_key_vault` resource.
func (r *Secret) VaultUri() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vaultUri"])
}

// The current version of the Key Vault Secret.
func (r *Secret) Version() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering Secret resources.
type SecretState struct {
	// Specifies the content type for the Key Vault Secret.
	ContentType interface{}
	// Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
	Name interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the value of the Key Vault Secret.
	Value interface{}
	// Specifies the URI used to access the Key Vault instance, available on the `azurerm_key_vault` resource.
	VaultUri interface{}
	// The current version of the Key Vault Secret.
	Version interface{}
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	// Specifies the content type for the Key Vault Secret.
	ContentType interface{}
	// Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
	Name interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the value of the Key Vault Secret.
	Value interface{}
	// Specifies the URI used to access the Key Vault instance, available on the `azurerm_key_vault` resource.
	VaultUri interface{}
}
