// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Automation DSC Node Configuration.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/automation_dsc_nodeconfiguration.html.markdown.
type DscNodeConfiguration struct {
	pulumi.CustomResourceState

	// The name of the automation account in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`

	ConfigurationName pulumi.StringOutput `pulumi:"configurationName"`

	// The PowerShell DSC Node Configuration (mof content).
	ContentEmbedded pulumi.StringOutput `pulumi:"contentEmbedded"`

	// Specifies the name of the DSC Node Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the resource group in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewDscNodeConfiguration registers a new resource with the given unique name, arguments, and options.
func NewDscNodeConfiguration(ctx *pulumi.Context,
	name string, args *DscNodeConfigurationArgs, opts ...pulumi.ResourceOption) (*DscNodeConfiguration, error) {
	if args == nil || args.AutomationAccountName == nil {
		return nil, errors.New("missing required argument 'AutomationAccountName'")
	}
	if args == nil || args.ContentEmbedded == nil {
		return nil, errors.New("missing required argument 'ContentEmbedded'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AutomationAccountName; i != nil { inputs["automationAccountName"] = i.ToStringOutput() }
		if i := args.ContentEmbedded; i != nil { inputs["contentEmbedded"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
	}
	var resource DscNodeConfiguration
	err := ctx.RegisterResource("azure:automation/dscNodeConfiguration:DscNodeConfiguration", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDscNodeConfiguration gets an existing DscNodeConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDscNodeConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DscNodeConfigurationState, opts ...pulumi.ResourceOption) (*DscNodeConfiguration, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AutomationAccountName; i != nil { inputs["automationAccountName"] = i.ToStringOutput() }
		if i := state.ConfigurationName; i != nil { inputs["configurationName"] = i.ToStringOutput() }
		if i := state.ContentEmbedded; i != nil { inputs["contentEmbedded"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
	}
	var resource DscNodeConfiguration
	err := ctx.ReadResource("azure:automation/dscNodeConfiguration:DscNodeConfiguration", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DscNodeConfiguration resources.
type DscNodeConfigurationState struct {
	// The name of the automation account in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
	// The PowerShell DSC Node Configuration (mof content).
	ContentEmbedded pulumi.StringInput `pulumi:"contentEmbedded"`
	// Specifies the name of the DSC Node Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DscNodeConfiguration resource.
type DscNodeConfigurationArgs struct {
	// The name of the automation account in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// The PowerShell DSC Node Configuration (mof content).
	ContentEmbedded pulumi.StringInput `pulumi:"contentEmbedded"`
	// Specifies the name of the DSC Node Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}
