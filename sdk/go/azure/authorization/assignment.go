// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Assigns a given Principal (User or Application) to a given Role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/role_assignment.html.markdown.
type Assignment struct {
	pulumi.CustomResourceState

	// A unique UUID/GUID for this Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The ID of the Principal (User, Group, Service Principal, or Application) to assign the Role Definition to. Changing this forces a new resource to be created. 
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`

	// The type of the `principalId`, e.g. User, Group, Service Principal, Application, etc.
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`

	// The Scoped-ID of the Role Definition. Changing this forces a new resource to be created. Conflicts with `roleDefinitionName`.
	RoleDefinitionId pulumi.StringOutput `pulumi:"roleDefinitionId"`

	// The name of a built-in Role. Changing this forces a new resource to be created. Conflicts with `roleDefinitionId`.
	RoleDefinitionName pulumi.StringOutput `pulumi:"roleDefinitionName"`

	// The scope at which the Role Assignment applies to, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`, or `/providers/Microsoft.Management/managementGroups/myMG`. Changing this forces a new resource to be created.
	Scope pulumi.StringOutput `pulumi:"scope"`

	// If the `principalId` is a newly provisioned `Service Principal` set this value to `true` to skip the `Azure Active Directory` check which may fail due to replication lag. This argument is only valid if the `principalId` is a `Service Principal` identity. If it is not a `Service Principal` identity it will cause the role assignment to fail. Defaults to `false`.
	SkipServicePrincipalAadCheck pulumi.BoolOutput `pulumi:"skipServicePrincipalAadCheck"`
}

// NewAssignment registers a new resource with the given unique name, arguments, and options.
func NewAssignment(ctx *pulumi.Context,
	name string, args *AssignmentArgs, opts ...pulumi.ResourceOption) (*Assignment, error) {
	if args == nil || args.PrincipalId == nil {
		return nil, errors.New("missing required argument 'PrincipalId'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.PrincipalId; i != nil { inputs["principalId"] = i.ToStringOutput() }
		if i := args.RoleDefinitionId; i != nil { inputs["roleDefinitionId"] = i.ToStringOutput() }
		if i := args.RoleDefinitionName; i != nil { inputs["roleDefinitionName"] = i.ToStringOutput() }
		if i := args.Scope; i != nil { inputs["scope"] = i.ToStringOutput() }
		if i := args.SkipServicePrincipalAadCheck; i != nil { inputs["skipServicePrincipalAadCheck"] = i.ToBoolOutput() }
	}
	var resource Assignment
	err := ctx.RegisterResource("azure:authorization/assignment:Assignment", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssignment gets an existing Assignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssignmentState, opts ...pulumi.ResourceOption) (*Assignment, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.PrincipalId; i != nil { inputs["principalId"] = i.ToStringOutput() }
		if i := state.PrincipalType; i != nil { inputs["principalType"] = i.ToStringOutput() }
		if i := state.RoleDefinitionId; i != nil { inputs["roleDefinitionId"] = i.ToStringOutput() }
		if i := state.RoleDefinitionName; i != nil { inputs["roleDefinitionName"] = i.ToStringOutput() }
		if i := state.Scope; i != nil { inputs["scope"] = i.ToStringOutput() }
		if i := state.SkipServicePrincipalAadCheck; i != nil { inputs["skipServicePrincipalAadCheck"] = i.ToBoolOutput() }
	}
	var resource Assignment
	err := ctx.ReadResource("azure:authorization/assignment:Assignment", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assignment resources.
type AssignmentState struct {
	// A unique UUID/GUID for this Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the Principal (User, Group, Service Principal, or Application) to assign the Role Definition to. Changing this forces a new resource to be created. 
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	// The type of the `principalId`, e.g. User, Group, Service Principal, Application, etc.
	PrincipalType pulumi.StringInput `pulumi:"principalType"`
	// The Scoped-ID of the Role Definition. Changing this forces a new resource to be created. Conflicts with `roleDefinitionName`.
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
	// The name of a built-in Role. Changing this forces a new resource to be created. Conflicts with `roleDefinitionId`.
	RoleDefinitionName pulumi.StringInput `pulumi:"roleDefinitionName"`
	// The scope at which the Role Assignment applies to, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`, or `/providers/Microsoft.Management/managementGroups/myMG`. Changing this forces a new resource to be created.
	Scope pulumi.StringInput `pulumi:"scope"`
	// If the `principalId` is a newly provisioned `Service Principal` set this value to `true` to skip the `Azure Active Directory` check which may fail due to replication lag. This argument is only valid if the `principalId` is a `Service Principal` identity. If it is not a `Service Principal` identity it will cause the role assignment to fail. Defaults to `false`.
	SkipServicePrincipalAadCheck pulumi.BoolInput `pulumi:"skipServicePrincipalAadCheck"`
}

// The set of arguments for constructing a Assignment resource.
type AssignmentArgs struct {
	// A unique UUID/GUID for this Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the Principal (User, Group, Service Principal, or Application) to assign the Role Definition to. Changing this forces a new resource to be created. 
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	// The Scoped-ID of the Role Definition. Changing this forces a new resource to be created. Conflicts with `roleDefinitionName`.
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
	// The name of a built-in Role. Changing this forces a new resource to be created. Conflicts with `roleDefinitionId`.
	RoleDefinitionName pulumi.StringInput `pulumi:"roleDefinitionName"`
	// The scope at which the Role Assignment applies to, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`, or `/providers/Microsoft.Management/managementGroups/myMG`. Changing this forces a new resource to be created.
	Scope pulumi.StringInput `pulumi:"scope"`
	// If the `principalId` is a newly provisioned `Service Principal` set this value to `true` to skip the `Azure Active Directory` check which may fail due to replication lag. This argument is only valid if the `principalId` is a `Service Principal` identity. If it is not a `Service Principal` identity it will cause the role assignment to fail. Defaults to `false`.
	SkipServicePrincipalAadCheck pulumi.BoolInput `pulumi:"skipServicePrincipalAadCheck"`
}
