// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API within an API Management Service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_api.html.markdown.
type Api struct {
	pulumi.CustomResourceState

	// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`

	// A description of the API Management API, which may include HTML formatting tags.
	Description pulumi.StringOutput `pulumi:"description"`

	// The display name of the API.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`

	// A `import` block as documented below.
	Import ApiImportOutput `pulumi:"import"`

	// Is this the current API Revision?
	IsCurrent pulumi.BoolOutput `pulumi:"isCurrent"`

	// Is this API Revision online/accessible via the Gateway?
	IsOnline pulumi.BoolOutput `pulumi:"isOnline"`

	// The name of the API Management API. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of it's resource paths within the API Management Service.
	Path pulumi.StringOutput `pulumi:"path"`

	// A list of protocols the operations in this API can be invoked. Possible values are `http` and `https`.
	Protocols pulumi.StringArrayOutput `pulumi:"protocols"`

	// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The Revision which used for this API.
	Revision pulumi.StringOutput `pulumi:"revision"`

	// Absolute URL of the backend service implementing this API.
	ServiceUrl pulumi.StringOutput `pulumi:"serviceUrl"`

	// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to `false`.
	SoapPassThrough pulumi.BoolOutput `pulumi:"soapPassThrough"`

	// A `subscriptionKeyParameterNames` block as documented below.
	SubscriptionKeyParameterNames ApiSubscriptionKeyParameterNamesOutput `pulumi:"subscriptionKeyParameterNames"`

	// The Version number of this API, if this API is versioned.
	Version pulumi.StringOutput `pulumi:"version"`

	// The ID of the Version Set which this API is associated with.
	VersionSetId pulumi.StringOutput `pulumi:"versionSetId"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil || args.Protocols == nil {
		return nil, errors.New("missing required argument 'Protocols'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Revision == nil {
		return nil, errors.New("missing required argument 'Revision'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ApiManagementName; i != nil { inputs["apiManagementName"] = i.ToStringOutput() }
		if i := args.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := args.DisplayName; i != nil { inputs["displayName"] = i.ToStringOutput() }
		if i := args.Import; i != nil { inputs["import"] = i.ToApiImportOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Path; i != nil { inputs["path"] = i.ToStringOutput() }
		if i := args.Protocols; i != nil { inputs["protocols"] = i.ToStringArrayOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.Revision; i != nil { inputs["revision"] = i.ToStringOutput() }
		if i := args.ServiceUrl; i != nil { inputs["serviceUrl"] = i.ToStringOutput() }
		if i := args.SoapPassThrough; i != nil { inputs["soapPassThrough"] = i.ToBoolOutput() }
		if i := args.SubscriptionKeyParameterNames; i != nil { inputs["subscriptionKeyParameterNames"] = i.ToApiSubscriptionKeyParameterNamesOutput() }
		if i := args.Version; i != nil { inputs["version"] = i.ToStringOutput() }
		if i := args.VersionSetId; i != nil { inputs["versionSetId"] = i.ToStringOutput() }
	}
	var resource Api
	err := ctx.RegisterResource("azure:apimanagement/api:Api", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ApiManagementName; i != nil { inputs["apiManagementName"] = i.ToStringOutput() }
		if i := state.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := state.DisplayName; i != nil { inputs["displayName"] = i.ToStringOutput() }
		if i := state.Import; i != nil { inputs["import"] = i.ToApiImportOutput() }
		if i := state.IsCurrent; i != nil { inputs["isCurrent"] = i.ToBoolOutput() }
		if i := state.IsOnline; i != nil { inputs["isOnline"] = i.ToBoolOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Path; i != nil { inputs["path"] = i.ToStringOutput() }
		if i := state.Protocols; i != nil { inputs["protocols"] = i.ToStringArrayOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.Revision; i != nil { inputs["revision"] = i.ToStringOutput() }
		if i := state.ServiceUrl; i != nil { inputs["serviceUrl"] = i.ToStringOutput() }
		if i := state.SoapPassThrough; i != nil { inputs["soapPassThrough"] = i.ToBoolOutput() }
		if i := state.SubscriptionKeyParameterNames; i != nil { inputs["subscriptionKeyParameterNames"] = i.ToApiSubscriptionKeyParameterNamesOutput() }
		if i := state.Version; i != nil { inputs["version"] = i.ToStringOutput() }
		if i := state.VersionSetId; i != nil { inputs["versionSetId"] = i.ToStringOutput() }
	}
	var resource Api
	err := ctx.ReadResource("azure:apimanagement/api:Api", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type ApiState struct {
	// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput `pulumi:"apiManagementName"`
	// A description of the API Management API, which may include HTML formatting tags.
	Description pulumi.StringInput `pulumi:"description"`
	// The display name of the API.
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// A `import` block as documented below.
	Import ApiImportInput `pulumi:"import"`
	// Is this the current API Revision?
	IsCurrent pulumi.BoolInput `pulumi:"isCurrent"`
	// Is this API Revision online/accessible via the Gateway?
	IsOnline pulumi.BoolInput `pulumi:"isOnline"`
	// The name of the API Management API. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of it's resource paths within the API Management Service.
	Path pulumi.StringInput `pulumi:"path"`
	// A list of protocols the operations in this API can be invoked. Possible values are `http` and `https`.
	Protocols pulumi.StringArrayInput `pulumi:"protocols"`
	// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Revision which used for this API.
	Revision pulumi.StringInput `pulumi:"revision"`
	// Absolute URL of the backend service implementing this API.
	ServiceUrl pulumi.StringInput `pulumi:"serviceUrl"`
	// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to `false`.
	SoapPassThrough pulumi.BoolInput `pulumi:"soapPassThrough"`
	// A `subscriptionKeyParameterNames` block as documented below.
	SubscriptionKeyParameterNames ApiSubscriptionKeyParameterNamesInput `pulumi:"subscriptionKeyParameterNames"`
	// The Version number of this API, if this API is versioned.
	Version pulumi.StringInput `pulumi:"version"`
	// The ID of the Version Set which this API is associated with.
	VersionSetId pulumi.StringInput `pulumi:"versionSetId"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput `pulumi:"apiManagementName"`
	// A description of the API Management API, which may include HTML formatting tags.
	Description pulumi.StringInput `pulumi:"description"`
	// The display name of the API.
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// A `import` block as documented below.
	Import ApiImportInput `pulumi:"import"`
	// The name of the API Management API. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of it's resource paths within the API Management Service.
	Path pulumi.StringInput `pulumi:"path"`
	// A list of protocols the operations in this API can be invoked. Possible values are `http` and `https`.
	Protocols pulumi.StringArrayInput `pulumi:"protocols"`
	// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Revision which used for this API.
	Revision pulumi.StringInput `pulumi:"revision"`
	// Absolute URL of the backend service implementing this API.
	ServiceUrl pulumi.StringInput `pulumi:"serviceUrl"`
	// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to `false`.
	SoapPassThrough pulumi.BoolInput `pulumi:"soapPassThrough"`
	// A `subscriptionKeyParameterNames` block as documented below.
	SubscriptionKeyParameterNames ApiSubscriptionKeyParameterNamesInput `pulumi:"subscriptionKeyParameterNames"`
	// The Version number of this API, if this API is versioned.
	Version pulumi.StringInput `pulumi:"version"`
	// The ID of the Version Set which this API is associated with.
	VersionSetId pulumi.StringInput `pulumi:"versionSetId"`
}
type ApiImport struct {
	ContentFormat string `pulumi:"contentFormat"`
	ContentValue string `pulumi:"contentValue"`
	WsdlSelector *ApiImportWsdlSelector `pulumi:"wsdlSelector"`
}
var apiImportType = reflect.TypeOf((*ApiImport)(nil)).Elem()

type ApiImportInput interface {
	pulumi.Input

	ToApiImportOutput() ApiImportOutput
	ToApiImportOutputWithContext(ctx context.Context) ApiImportOutput
}

type ApiImportArgs struct {
	ContentFormat pulumi.StringInput `pulumi:"contentFormat"`
	ContentValue pulumi.StringInput `pulumi:"contentValue"`
	WsdlSelector ApiImportWsdlSelectorInput `pulumi:"wsdlSelector"`
}

func (ApiImportArgs) ElementType() reflect.Type {
	return apiImportType
}

func (a ApiImportArgs) ToApiImportOutput() ApiImportOutput {
	return pulumi.ToOutput(a).(ApiImportOutput)
}

func (a ApiImportArgs) ToApiImportOutputWithContext(ctx context.Context) ApiImportOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ApiImportOutput)
}

type ApiImportOutput struct { *pulumi.OutputState }

func (o ApiImportOutput) ContentFormat() pulumi.StringOutput {
	return o.Apply(func(v ApiImport) string {
		return v.ContentFormat
	}).(pulumi.StringOutput)
}

func (o ApiImportOutput) ContentValue() pulumi.StringOutput {
	return o.Apply(func(v ApiImport) string {
		return v.ContentValue
	}).(pulumi.StringOutput)
}

func (o ApiImportOutput) WsdlSelector() ApiImportWsdlSelectorOutput {
	return o.Apply(func(v ApiImport) ApiImportWsdlSelector {
		if v.WsdlSelector == nil { return *new(ApiImportWsdlSelector) } else { return *v.WsdlSelector }
	}).(ApiImportWsdlSelectorOutput)
}

func (ApiImportOutput) ElementType() reflect.Type {
	return apiImportType
}

func (o ApiImportOutput) ToApiImportOutput() ApiImportOutput {
	return o
}

func (o ApiImportOutput) ToApiImportOutputWithContext(ctx context.Context) ApiImportOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ApiImportOutput{}) }

type ApiImportWsdlSelector struct {
	EndpointName string `pulumi:"endpointName"`
	ServiceName string `pulumi:"serviceName"`
}
var apiImportWsdlSelectorType = reflect.TypeOf((*ApiImportWsdlSelector)(nil)).Elem()

type ApiImportWsdlSelectorInput interface {
	pulumi.Input

	ToApiImportWsdlSelectorOutput() ApiImportWsdlSelectorOutput
	ToApiImportWsdlSelectorOutputWithContext(ctx context.Context) ApiImportWsdlSelectorOutput
}

type ApiImportWsdlSelectorArgs struct {
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ApiImportWsdlSelectorArgs) ElementType() reflect.Type {
	return apiImportWsdlSelectorType
}

func (a ApiImportWsdlSelectorArgs) ToApiImportWsdlSelectorOutput() ApiImportWsdlSelectorOutput {
	return pulumi.ToOutput(a).(ApiImportWsdlSelectorOutput)
}

func (a ApiImportWsdlSelectorArgs) ToApiImportWsdlSelectorOutputWithContext(ctx context.Context) ApiImportWsdlSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ApiImportWsdlSelectorOutput)
}

type ApiImportWsdlSelectorOutput struct { *pulumi.OutputState }

func (o ApiImportWsdlSelectorOutput) EndpointName() pulumi.StringOutput {
	return o.Apply(func(v ApiImportWsdlSelector) string {
		return v.EndpointName
	}).(pulumi.StringOutput)
}

func (o ApiImportWsdlSelectorOutput) ServiceName() pulumi.StringOutput {
	return o.Apply(func(v ApiImportWsdlSelector) string {
		return v.ServiceName
	}).(pulumi.StringOutput)
}

func (ApiImportWsdlSelectorOutput) ElementType() reflect.Type {
	return apiImportWsdlSelectorType
}

func (o ApiImportWsdlSelectorOutput) ToApiImportWsdlSelectorOutput() ApiImportWsdlSelectorOutput {
	return o
}

func (o ApiImportWsdlSelectorOutput) ToApiImportWsdlSelectorOutputWithContext(ctx context.Context) ApiImportWsdlSelectorOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ApiImportWsdlSelectorOutput{}) }

type ApiSubscriptionKeyParameterNames struct {
	Header string `pulumi:"header"`
	Query string `pulumi:"query"`
}
var apiSubscriptionKeyParameterNamesType = reflect.TypeOf((*ApiSubscriptionKeyParameterNames)(nil)).Elem()

type ApiSubscriptionKeyParameterNamesInput interface {
	pulumi.Input

	ToApiSubscriptionKeyParameterNamesOutput() ApiSubscriptionKeyParameterNamesOutput
	ToApiSubscriptionKeyParameterNamesOutputWithContext(ctx context.Context) ApiSubscriptionKeyParameterNamesOutput
}

type ApiSubscriptionKeyParameterNamesArgs struct {
	Header pulumi.StringInput `pulumi:"header"`
	Query pulumi.StringInput `pulumi:"query"`
}

func (ApiSubscriptionKeyParameterNamesArgs) ElementType() reflect.Type {
	return apiSubscriptionKeyParameterNamesType
}

func (a ApiSubscriptionKeyParameterNamesArgs) ToApiSubscriptionKeyParameterNamesOutput() ApiSubscriptionKeyParameterNamesOutput {
	return pulumi.ToOutput(a).(ApiSubscriptionKeyParameterNamesOutput)
}

func (a ApiSubscriptionKeyParameterNamesArgs) ToApiSubscriptionKeyParameterNamesOutputWithContext(ctx context.Context) ApiSubscriptionKeyParameterNamesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ApiSubscriptionKeyParameterNamesOutput)
}

type ApiSubscriptionKeyParameterNamesOutput struct { *pulumi.OutputState }

func (o ApiSubscriptionKeyParameterNamesOutput) Header() pulumi.StringOutput {
	return o.Apply(func(v ApiSubscriptionKeyParameterNames) string {
		return v.Header
	}).(pulumi.StringOutput)
}

func (o ApiSubscriptionKeyParameterNamesOutput) Query() pulumi.StringOutput {
	return o.Apply(func(v ApiSubscriptionKeyParameterNames) string {
		return v.Query
	}).(pulumi.StringOutput)
}

func (ApiSubscriptionKeyParameterNamesOutput) ElementType() reflect.Type {
	return apiSubscriptionKeyParameterNamesType
}

func (o ApiSubscriptionKeyParameterNamesOutput) ToApiSubscriptionKeyParameterNamesOutput() ApiSubscriptionKeyParameterNamesOutput {
	return o
}

func (o ApiSubscriptionKeyParameterNamesOutput) ToApiSubscriptionKeyParameterNamesOutputWithContext(ctx context.Context) ApiSubscriptionKeyParameterNamesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ApiSubscriptionKeyParameterNamesOutput{}) }

