// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a integer variable in Azure Automation
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/automation_variable_int.html.markdown.
type IntVariable struct {
	pulumi.CustomResourceState

	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`

	// The description of the Automation Variable.
	Description pulumi.StringOutput `pulumi:"description"`

	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted pulumi.BoolOutput `pulumi:"encrypted"`

	// The name of the Automation Variable. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The value of the Automation Variable as a `integer`.
	Value pulumi.IntOutput `pulumi:"value"`
}

// NewIntVariable registers a new resource with the given unique name, arguments, and options.
func NewIntVariable(ctx *pulumi.Context,
	name string, args *IntVariableArgs, opts ...pulumi.ResourceOption) (*IntVariable, error) {
	if args == nil || args.AutomationAccountName == nil {
		return nil, errors.New("missing required argument 'AutomationAccountName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AutomationAccountName; i != nil { inputs["automationAccountName"] = i.ToStringOutput() }
		if i := args.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := args.Encrypted; i != nil { inputs["encrypted"] = i.ToBoolOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Value; i != nil { inputs["value"] = i.ToIntOutput() }
	}
	var resource IntVariable
	err := ctx.RegisterResource("azure:automation/intVariable:IntVariable", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntVariable gets an existing IntVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntVariableState, opts ...pulumi.ResourceOption) (*IntVariable, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AutomationAccountName; i != nil { inputs["automationAccountName"] = i.ToStringOutput() }
		if i := state.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := state.Encrypted; i != nil { inputs["encrypted"] = i.ToBoolOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Value; i != nil { inputs["value"] = i.ToIntOutput() }
	}
	var resource IntVariable
	err := ctx.ReadResource("azure:automation/intVariable:IntVariable", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntVariable resources.
type IntVariableState struct {
	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// The description of the Automation Variable.
	Description pulumi.StringInput `pulumi:"description"`
	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted pulumi.BoolInput `pulumi:"encrypted"`
	// The name of the Automation Variable. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The value of the Automation Variable as a `integer`.
	Value pulumi.IntInput `pulumi:"value"`
}

// The set of arguments for constructing a IntVariable resource.
type IntVariableArgs struct {
	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// The description of the Automation Variable.
	Description pulumi.StringInput `pulumi:"description"`
	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted pulumi.BoolInput `pulumi:"encrypted"`
	// The name of the Automation Variable. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The value of the Automation Variable as a `integer`.
	Value pulumi.IntInput `pulumi:"value"`
}
