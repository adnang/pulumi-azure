// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appinsights

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Application Insights API key.
type ApiKey struct {
	s *pulumi.ResourceState
}

// NewApiKey registers a new resource with the given unique name, arguments, and options.
func NewApiKey(ctx *pulumi.Context,
	name string, args *ApiKeyArgs, opts ...pulumi.ResourceOpt) (*ApiKey, error) {
	if args == nil || args.ApplicationInsightsId == nil {
		return nil, errors.New("missing required argument 'ApplicationInsightsId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["applicationInsightsId"] = nil
		inputs["name"] = nil
		inputs["readPermissions"] = nil
		inputs["writePermissions"] = nil
	} else {
		inputs["applicationInsightsId"] = args.ApplicationInsightsId
		inputs["name"] = args.Name
		inputs["readPermissions"] = args.ReadPermissions
		inputs["writePermissions"] = args.WritePermissions
	}
	inputs["apiKey"] = nil
	s, err := ctx.RegisterResource("azure:appinsights/apiKey:ApiKey", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiKey{s: s}, nil
}

// GetApiKey gets an existing ApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApiKeyState, opts ...pulumi.ResourceOpt) (*ApiKey, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiKey"] = state.ApiKey
		inputs["applicationInsightsId"] = state.ApplicationInsightsId
		inputs["name"] = state.Name
		inputs["readPermissions"] = state.ReadPermissions
		inputs["writePermissions"] = state.WritePermissions
	}
	s, err := ctx.ReadResource("azure:appinsights/apiKey:ApiKey", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiKey{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ApiKey) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ApiKey) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The API Key secret (Sensitive).
func (r *ApiKey) ApiKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["apiKey"])
}

// The ID of the Application Insights component on which the API key operates. Changing this forces a new resource to be created.
func (r *ApiKey) ApplicationInsightsId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["applicationInsightsId"])
}

// Specifies the name of the Application Insights API key. Changing this forces a
// new resource to be created.
func (r *ApiKey) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Specifies the list of read permissions granted to the API key. Valid values are `agentconfig`, `aggregate`, `api`, `draft`, `extendqueries`, `search`. Please note these values are case sensitive. Changing this forces a new resource to be created. 
func (r *ApiKey) ReadPermissions() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["readPermissions"])
}

// Specifies the list of write permissions granted to the API key. Valid values are `annotations`. Please note these values are case sensitive. Changing this forces a new resource to be created.
func (r *ApiKey) WritePermissions() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["writePermissions"])
}

// Input properties used for looking up and filtering ApiKey resources.
type ApiKeyState struct {
	// The API Key secret (Sensitive).
	ApiKey interface{}
	// The ID of the Application Insights component on which the API key operates. Changing this forces a new resource to be created.
	ApplicationInsightsId interface{}
	// Specifies the name of the Application Insights API key. Changing this forces a
	// new resource to be created.
	Name interface{}
	// Specifies the list of read permissions granted to the API key. Valid values are `agentconfig`, `aggregate`, `api`, `draft`, `extendqueries`, `search`. Please note these values are case sensitive. Changing this forces a new resource to be created. 
	ReadPermissions interface{}
	// Specifies the list of write permissions granted to the API key. Valid values are `annotations`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	WritePermissions interface{}
}

// The set of arguments for constructing a ApiKey resource.
type ApiKeyArgs struct {
	// The ID of the Application Insights component on which the API key operates. Changing this forces a new resource to be created.
	ApplicationInsightsId interface{}
	// Specifies the name of the Application Insights API key. Changing this forces a
	// new resource to be created.
	Name interface{}
	// Specifies the list of read permissions granted to the API key. Valid values are `agentconfig`, `aggregate`, `api`, `draft`, `extendqueries`, `search`. Please note these values are case sensitive. Changing this forces a new resource to be created. 
	ReadPermissions interface{}
	// Specifies the list of write permissions granted to the API key. Valid values are `annotations`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	WritePermissions interface{}
}
