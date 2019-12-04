// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Custom Action within a Logic App Workflow
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/logic_app_action_custom.html.markdown.
type ActionCustom struct {
	pulumi.CustomResourceState

	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body pulumi.StringOutput `pulumi:"body"`

	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringOutput `pulumi:"logicAppId"`

	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewActionCustom registers a new resource with the given unique name, arguments, and options.
func NewActionCustom(ctx *pulumi.Context,
	name string, args *ActionCustomArgs, opts ...pulumi.ResourceOption) (*ActionCustom, error) {
	if args == nil || args.Body == nil {
		return nil, errors.New("missing required argument 'Body'")
	}
	if args == nil || args.LogicAppId == nil {
		return nil, errors.New("missing required argument 'LogicAppId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Body; i != nil { inputs["body"] = i.ToStringOutput() }
		if i := args.LogicAppId; i != nil { inputs["logicAppId"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
	}
	var resource ActionCustom
	err := ctx.RegisterResource("azure:logicapps/actionCustom:ActionCustom", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionCustom gets an existing ActionCustom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionCustom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionCustomState, opts ...pulumi.ResourceOption) (*ActionCustom, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Body; i != nil { inputs["body"] = i.ToStringOutput() }
		if i := state.LogicAppId; i != nil { inputs["logicAppId"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
	}
	var resource ActionCustom
	err := ctx.ReadResource("azure:logicapps/actionCustom:ActionCustom", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionCustom resources.
type ActionCustomState struct {
	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body pulumi.StringInput `pulumi:"body"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringInput `pulumi:"logicAppId"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
}

// The set of arguments for constructing a ActionCustom resource.
type ActionCustomArgs struct {
	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body pulumi.StringInput `pulumi:"body"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringInput `pulumi:"logicAppId"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
}
