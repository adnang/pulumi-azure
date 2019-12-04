// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ad

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/azuread_application.html.markdown.
type Application struct {
	pulumi.CustomResourceState

	// The Application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`

	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolOutput `pulumi:"availableToOtherTenants"`

	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringOutput `pulumi:"homepage"`

	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayOutput `pulumi:"identifierUris"`

	// The display name for the application.
	Name pulumi.StringOutput `pulumi:"name"`

	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolOutput `pulumi:"oauth2AllowImplicitFlow"`

	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayOutput `pulumi:"replyUrls"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AvailableToOtherTenants; i != nil { inputs["availableToOtherTenants"] = i.ToBoolOutput() }
		if i := args.Homepage; i != nil { inputs["homepage"] = i.ToStringOutput() }
		if i := args.IdentifierUris; i != nil { inputs["identifierUris"] = i.ToStringArrayOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Oauth2AllowImplicitFlow; i != nil { inputs["oauth2AllowImplicitFlow"] = i.ToBoolOutput() }
		if i := args.ReplyUrls; i != nil { inputs["replyUrls"] = i.ToStringArrayOutput() }
	}
	var resource Application
	err := ctx.RegisterResource("azure:ad/application:Application", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ApplicationId; i != nil { inputs["applicationId"] = i.ToStringOutput() }
		if i := state.AvailableToOtherTenants; i != nil { inputs["availableToOtherTenants"] = i.ToBoolOutput() }
		if i := state.Homepage; i != nil { inputs["homepage"] = i.ToStringOutput() }
		if i := state.IdentifierUris; i != nil { inputs["identifierUris"] = i.ToStringArrayOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Oauth2AllowImplicitFlow; i != nil { inputs["oauth2AllowImplicitFlow"] = i.ToBoolOutput() }
		if i := state.ReplyUrls; i != nil { inputs["replyUrls"] = i.ToStringArrayOutput() }
	}
	var resource Application
	err := ctx.ReadResource("azure:ad/application:Application", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type ApplicationState struct {
	// The Application ID.
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolInput `pulumi:"availableToOtherTenants"`
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringInput `pulumi:"homepage"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayInput `pulumi:"identifierUris"`
	// The display name for the application.
	Name pulumi.StringInput `pulumi:"name"`
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolInput `pulumi:"oauth2AllowImplicitFlow"`
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayInput `pulumi:"replyUrls"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolInput `pulumi:"availableToOtherTenants"`
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringInput `pulumi:"homepage"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayInput `pulumi:"identifierUris"`
	// The display name for the application.
	Name pulumi.StringInput `pulumi:"name"`
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolInput `pulumi:"oauth2AllowImplicitFlow"`
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayInput `pulumi:"replyUrls"`
}
