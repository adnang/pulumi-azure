// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a certificate in an Azure Batch account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/batch_certificate.html.markdown.
type Certificate struct {
	pulumi.CustomResourceState

	// Specifies the name of the Batch account. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`

	// The base64-encoded contents of the certificate.
	Certificate pulumi.StringOutput `pulumi:"certificate"`

	// The format of the certificate. Possible values are `Cer` or `Pfx`.
	Format pulumi.StringOutput `pulumi:"format"`

	// The generated name of the certificate.
	Name pulumi.StringOutput `pulumi:"name"`

	// The password to access the certificate's private key. This must and can only be specified when `format` is `Pfx`.
	Password pulumi.StringOutput `pulumi:"password"`

	// The public key of the certificate.
	PublicData pulumi.StringOutput `pulumi:"publicData"`

	// The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The thumbprint of the certificate. At this time the only supported value is 'SHA1'.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`

	ThumbprintAlgorithm pulumi.StringOutput `pulumi:"thumbprintAlgorithm"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Certificate == nil {
		return nil, errors.New("missing required argument 'Certificate'")
	}
	if args == nil || args.Format == nil {
		return nil, errors.New("missing required argument 'Format'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Thumbprint == nil {
		return nil, errors.New("missing required argument 'Thumbprint'")
	}
	if args == nil || args.ThumbprintAlgorithm == nil {
		return nil, errors.New("missing required argument 'ThumbprintAlgorithm'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AccountName; i != nil { inputs["accountName"] = i.ToStringOutput() }
		if i := args.Certificate; i != nil { inputs["certificate"] = i.ToStringOutput() }
		if i := args.Format; i != nil { inputs["format"] = i.ToStringOutput() }
		if i := args.Password; i != nil { inputs["password"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Thumbprint; i != nil { inputs["thumbprint"] = i.ToStringOutput() }
		if i := args.ThumbprintAlgorithm; i != nil { inputs["thumbprintAlgorithm"] = i.ToStringOutput() }
	}
	var resource Certificate
	err := ctx.RegisterResource("azure:batch/certificate:Certificate", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AccountName; i != nil { inputs["accountName"] = i.ToStringOutput() }
		if i := state.Certificate; i != nil { inputs["certificate"] = i.ToStringOutput() }
		if i := state.Format; i != nil { inputs["format"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Password; i != nil { inputs["password"] = i.ToStringOutput() }
		if i := state.PublicData; i != nil { inputs["publicData"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Thumbprint; i != nil { inputs["thumbprint"] = i.ToStringOutput() }
		if i := state.ThumbprintAlgorithm; i != nil { inputs["thumbprintAlgorithm"] = i.ToStringOutput() }
	}
	var resource Certificate
	err := ctx.ReadResource("azure:batch/certificate:Certificate", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type CertificateState struct {
	// Specifies the name of the Batch account. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The base64-encoded contents of the certificate.
	Certificate pulumi.StringInput `pulumi:"certificate"`
	// The format of the certificate. Possible values are `Cer` or `Pfx`.
	Format pulumi.StringInput `pulumi:"format"`
	// The generated name of the certificate.
	Name pulumi.StringInput `pulumi:"name"`
	// The password to access the certificate's private key. This must and can only be specified when `format` is `Pfx`.
	Password pulumi.StringInput `pulumi:"password"`
	// The public key of the certificate.
	PublicData pulumi.StringInput `pulumi:"publicData"`
	// The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The thumbprint of the certificate. At this time the only supported value is 'SHA1'.
	Thumbprint pulumi.StringInput `pulumi:"thumbprint"`
	ThumbprintAlgorithm pulumi.StringInput `pulumi:"thumbprintAlgorithm"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Specifies the name of the Batch account. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The base64-encoded contents of the certificate.
	Certificate pulumi.StringInput `pulumi:"certificate"`
	// The format of the certificate. Possible values are `Cer` or `Pfx`.
	Format pulumi.StringInput `pulumi:"format"`
	// The password to access the certificate's private key. This must and can only be specified when `format` is `Pfx`.
	Password pulumi.StringInput `pulumi:"password"`
	// The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The thumbprint of the certificate. At this time the only supported value is 'SHA1'.
	Thumbprint pulumi.StringInput `pulumi:"thumbprint"`
	ThumbprintAlgorithm pulumi.StringInput `pulumi:"thumbprintAlgorithm"`
}
