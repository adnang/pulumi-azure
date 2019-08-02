// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an EventGrid Domain
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventgrid_domain_legacy.html.markdown.
type Domain struct {
	s *pulumi.ResourceState
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOpt) (*Domain, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["inputMappingDefaultValues"] = nil
		inputs["inputMappingFields"] = nil
		inputs["inputSchema"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
	} else {
		inputs["inputMappingDefaultValues"] = args.InputMappingDefaultValues
		inputs["inputMappingFields"] = args.InputMappingFields
		inputs["inputSchema"] = args.InputSchema
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
	}
	inputs["endpoint"] = nil
	s, err := ctx.RegisterResource("azure:eventhub/domain:Domain", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Domain{s: s}, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DomainState, opts ...pulumi.ResourceOpt) (*Domain, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["endpoint"] = state.Endpoint
		inputs["inputMappingDefaultValues"] = state.InputMappingDefaultValues
		inputs["inputMappingFields"] = state.InputMappingFields
		inputs["inputSchema"] = state.InputSchema
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:eventhub/domain:Domain", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Domain{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Domain) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Domain) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The Endpoint associated with the EventGrid Domain.
func (r *Domain) Endpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endpoint"])
}

// A `input_mapping_default_values` block as defined below.
func (r *Domain) InputMappingDefaultValues() *pulumi.Output {
	return r.s.State["inputMappingDefaultValues"]
}

// A `input_mapping_fields` block as defined below.
func (r *Domain) InputMappingFields() *pulumi.Output {
	return r.s.State["inputMappingFields"]
}

// Specifies the schema in which incoming events will be published to this domain. Allowed values are `cloudeventv01schema`, `customeventschema`, or `eventgridschema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
func (r *Domain) InputSchema() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["inputSchema"])
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *Domain) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
func (r *Domain) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
func (r *Domain) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *Domain) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Domain resources.
type DomainState struct {
	// The Endpoint associated with the EventGrid Domain.
	Endpoint interface{}
	// A `input_mapping_default_values` block as defined below.
	InputMappingDefaultValues interface{}
	// A `input_mapping_fields` block as defined below.
	InputMappingFields interface{}
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `cloudeventv01schema`, `customeventschema`, or `eventgridschema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
	InputSchema interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// A `input_mapping_default_values` block as defined below.
	InputMappingDefaultValues interface{}
	// A `input_mapping_fields` block as defined below.
	InputMappingFields interface{}
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `cloudeventv01schema`, `customeventschema`, or `eventgridschema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
	InputSchema interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
