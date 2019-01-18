// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package policy

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a policy rule definition on a management group or your provider subscription. 
// 
// Policy definitions do not take effect until they are assigned to a scope using a Policy Assignment.
type Definition struct {
	s *pulumi.ResourceState
}

// NewDefinition registers a new resource with the given unique name, arguments, and options.
func NewDefinition(ctx *pulumi.Context,
	name string, args *DefinitionArgs, opts ...pulumi.ResourceOpt) (*Definition, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Mode == nil {
		return nil, errors.New("missing required argument 'Mode'")
	}
	if args == nil || args.PolicyType == nil {
		return nil, errors.New("missing required argument 'PolicyType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["displayName"] = nil
		inputs["managementGroupId"] = nil
		inputs["metadata"] = nil
		inputs["mode"] = nil
		inputs["name"] = nil
		inputs["parameters"] = nil
		inputs["policyRule"] = nil
		inputs["policyType"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["displayName"] = args.DisplayName
		inputs["managementGroupId"] = args.ManagementGroupId
		inputs["metadata"] = args.Metadata
		inputs["mode"] = args.Mode
		inputs["name"] = args.Name
		inputs["parameters"] = args.Parameters
		inputs["policyRule"] = args.PolicyRule
		inputs["policyType"] = args.PolicyType
	}
	s, err := ctx.RegisterResource("azure:policy/definition:Definition", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Definition{s: s}, nil
}

// GetDefinition gets an existing Definition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefinition(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DefinitionState, opts ...pulumi.ResourceOpt) (*Definition, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["displayName"] = state.DisplayName
		inputs["managementGroupId"] = state.ManagementGroupId
		inputs["metadata"] = state.Metadata
		inputs["mode"] = state.Mode
		inputs["name"] = state.Name
		inputs["parameters"] = state.Parameters
		inputs["policyRule"] = state.PolicyRule
		inputs["policyType"] = state.PolicyType
	}
	s, err := ctx.ReadResource("azure:policy/definition:Definition", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Definition{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Definition) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Definition) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The description of the policy definition.
func (r *Definition) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The display name of the policy definition.
func (r *Definition) DisplayName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["displayName"])
}

// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
func (r *Definition) ManagementGroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["managementGroupId"])
}

// The metadata for the policy definition. This
// is a json object representing additional metadata that should be stored
// with the policy definition.
func (r *Definition) Metadata() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["metadata"])
}

// The policy mode that allows you to specify which resource
// types will be evaluated.  The value can be "All", "Indexed" or
// "NotSpecified". Changing this resource forces a new resource to be
// created.
func (r *Definition) Mode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["mode"])
}

// The name of the policy definition. Changing this forces a
// new resource to be created.
func (r *Definition) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Parameters for the policy definition. This field
// is a json object that allows you to parameterize your policy definition.
func (r *Definition) Parameters() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["parameters"])
}

// The policy rule for the policy definition. This
// is a json object representing the rule that contains an if and
// a then block.
func (r *Definition) PolicyRule() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyRule"])
}

// The policy type.  The value can be "BuiltIn", "Custom"
// or "NotSpecified". Changing this forces a new resource to be created.
func (r *Definition) PolicyType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyType"])
}

// Input properties used for looking up and filtering Definition resources.
type DefinitionState struct {
	// The description of the policy definition.
	Description interface{}
	// The display name of the policy definition.
	DisplayName interface{}
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId interface{}
	// The metadata for the policy definition. This
	// is a json object representing additional metadata that should be stored
	// with the policy definition.
	Metadata interface{}
	// The policy mode that allows you to specify which resource
	// types will be evaluated.  The value can be "All", "Indexed" or
	// "NotSpecified". Changing this resource forces a new resource to be
	// created.
	Mode interface{}
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name interface{}
	// Parameters for the policy definition. This field
	// is a json object that allows you to parameterize your policy definition.
	Parameters interface{}
	// The policy rule for the policy definition. This
	// is a json object representing the rule that contains an if and
	// a then block.
	PolicyRule interface{}
	// The policy type.  The value can be "BuiltIn", "Custom"
	// or "NotSpecified". Changing this forces a new resource to be created.
	PolicyType interface{}
}

// The set of arguments for constructing a Definition resource.
type DefinitionArgs struct {
	// The description of the policy definition.
	Description interface{}
	// The display name of the policy definition.
	DisplayName interface{}
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId interface{}
	// The metadata for the policy definition. This
	// is a json object representing additional metadata that should be stored
	// with the policy definition.
	Metadata interface{}
	// The policy mode that allows you to specify which resource
	// types will be evaluated.  The value can be "All", "Indexed" or
	// "NotSpecified". Changing this resource forces a new resource to be
	// created.
	Mode interface{}
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name interface{}
	// Parameters for the policy definition. This field
	// is a json object that allows you to parameterize your policy definition.
	Parameters interface{}
	// The policy rule for the policy definition. This
	// is a json object representing the rule that contains an if and
	// a then block.
	PolicyRule interface{}
	// The policy type.  The value can be "BuiltIn", "Custom"
	// or "NotSpecified". Changing this forces a new resource to be created.
	PolicyType interface{}
}
