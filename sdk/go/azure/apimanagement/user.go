// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Management User.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_user.html.markdown.
type User struct {
	pulumi.CustomResourceState

	// The name of the API Management Service in which the User should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`

	// The kind of confirmation email which will be sent to this user. Possible values are `invite` and `signup`. Changing this forces a new resource to be created.
	Confirmation pulumi.StringOutput `pulumi:"confirmation"`

	// The email address associated with this user.
	Email pulumi.StringOutput `pulumi:"email"`

	// The first name for this user.
	FirstName pulumi.StringOutput `pulumi:"firstName"`

	// The last name for this user.
	LastName pulumi.StringOutput `pulumi:"lastName"`

	// A note about this user.
	Note pulumi.StringOutput `pulumi:"note"`

	// The password associated with this user.
	Password pulumi.StringOutput `pulumi:"password"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The state of this user. Possible values are `active`, `blocked` and `pending`.
	State pulumi.StringOutput `pulumi:"state"`

	// The Identifier for this User, which must be unique within the API Management Service. Changing this forces a new resource to be created.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.Email == nil {
		return nil, errors.New("missing required argument 'Email'")
	}
	if args == nil || args.FirstName == nil {
		return nil, errors.New("missing required argument 'FirstName'")
	}
	if args == nil || args.LastName == nil {
		return nil, errors.New("missing required argument 'LastName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.UserId == nil {
		return nil, errors.New("missing required argument 'UserId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ApiManagementName; i != nil { inputs["apiManagementName"] = i.ToStringOutput() }
		if i := args.Confirmation; i != nil { inputs["confirmation"] = i.ToStringOutput() }
		if i := args.Email; i != nil { inputs["email"] = i.ToStringOutput() }
		if i := args.FirstName; i != nil { inputs["firstName"] = i.ToStringOutput() }
		if i := args.LastName; i != nil { inputs["lastName"] = i.ToStringOutput() }
		if i := args.Note; i != nil { inputs["note"] = i.ToStringOutput() }
		if i := args.Password; i != nil { inputs["password"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.State; i != nil { inputs["state"] = i.ToStringOutput() }
		if i := args.UserId; i != nil { inputs["userId"] = i.ToStringOutput() }
	}
	var resource User
	err := ctx.RegisterResource("azure:apimanagement/user:User", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ApiManagementName; i != nil { inputs["apiManagementName"] = i.ToStringOutput() }
		if i := state.Confirmation; i != nil { inputs["confirmation"] = i.ToStringOutput() }
		if i := state.Email; i != nil { inputs["email"] = i.ToStringOutput() }
		if i := state.FirstName; i != nil { inputs["firstName"] = i.ToStringOutput() }
		if i := state.LastName; i != nil { inputs["lastName"] = i.ToStringOutput() }
		if i := state.Note; i != nil { inputs["note"] = i.ToStringOutput() }
		if i := state.Password; i != nil { inputs["password"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.State; i != nil { inputs["state"] = i.ToStringOutput() }
		if i := state.UserId; i != nil { inputs["userId"] = i.ToStringOutput() }
	}
	var resource User
	err := ctx.ReadResource("azure:apimanagement/user:User", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type UserState struct {
	// The name of the API Management Service in which the User should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput `pulumi:"apiManagementName"`
	// The kind of confirmation email which will be sent to this user. Possible values are `invite` and `signup`. Changing this forces a new resource to be created.
	Confirmation pulumi.StringInput `pulumi:"confirmation"`
	// The email address associated with this user.
	Email pulumi.StringInput `pulumi:"email"`
	// The first name for this user.
	FirstName pulumi.StringInput `pulumi:"firstName"`
	// The last name for this user.
	LastName pulumi.StringInput `pulumi:"lastName"`
	// A note about this user.
	Note pulumi.StringInput `pulumi:"note"`
	// The password associated with this user.
	Password pulumi.StringInput `pulumi:"password"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The state of this user. Possible values are `active`, `blocked` and `pending`.
	State pulumi.StringInput `pulumi:"state"`
	// The Identifier for this User, which must be unique within the API Management Service. Changing this forces a new resource to be created.
	UserId pulumi.StringInput `pulumi:"userId"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The name of the API Management Service in which the User should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput `pulumi:"apiManagementName"`
	// The kind of confirmation email which will be sent to this user. Possible values are `invite` and `signup`. Changing this forces a new resource to be created.
	Confirmation pulumi.StringInput `pulumi:"confirmation"`
	// The email address associated with this user.
	Email pulumi.StringInput `pulumi:"email"`
	// The first name for this user.
	FirstName pulumi.StringInput `pulumi:"firstName"`
	// The last name for this user.
	LastName pulumi.StringInput `pulumi:"lastName"`
	// A note about this user.
	Note pulumi.StringInput `pulumi:"note"`
	// The password associated with this user.
	Password pulumi.StringInput `pulumi:"password"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The state of this user. Possible values are `active`, `blocked` and `pending`.
	State pulumi.StringInput `pulumi:"state"`
	// The Identifier for this User, which must be unique within the API Management Service. Changing this forces a new resource to be created.
	UserId pulumi.StringInput `pulumi:"userId"`
}
