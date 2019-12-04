// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an HTTP Action within a Logic App Workflow
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/logic_app_action_http.html.markdown.
type ActionHttp struct {
	pulumi.CustomResourceState

	// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
	Body pulumi.StringOutput `pulumi:"body"`

	// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
	Headers pulumi.StringMapOutput `pulumi:"headers"`

	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringOutput `pulumi:"logicAppId"`

	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
	Method pulumi.StringOutput `pulumi:"method"`

	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the URI which will be called when this HTTP Action is triggered.
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewActionHttp registers a new resource with the given unique name, arguments, and options.
func NewActionHttp(ctx *pulumi.Context,
	name string, args *ActionHttpArgs, opts ...pulumi.ResourceOption) (*ActionHttp, error) {
	if args == nil || args.LogicAppId == nil {
		return nil, errors.New("missing required argument 'LogicAppId'")
	}
	if args == nil || args.Method == nil {
		return nil, errors.New("missing required argument 'Method'")
	}
	if args == nil || args.Uri == nil {
		return nil, errors.New("missing required argument 'Uri'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Body; i != nil { inputs["body"] = i.ToStringOutput() }
		if i := args.Headers; i != nil { inputs["headers"] = i.ToStringMapOutput() }
		if i := args.LogicAppId; i != nil { inputs["logicAppId"] = i.ToStringOutput() }
		if i := args.Method; i != nil { inputs["method"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Uri; i != nil { inputs["uri"] = i.ToStringOutput() }
	}
	var resource ActionHttp
	err := ctx.RegisterResource("azure:logicapps/actionHttp:ActionHttp", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionHttp gets an existing ActionHttp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionHttp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionHttpState, opts ...pulumi.ResourceOption) (*ActionHttp, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Body; i != nil { inputs["body"] = i.ToStringOutput() }
		if i := state.Headers; i != nil { inputs["headers"] = i.ToStringMapOutput() }
		if i := state.LogicAppId; i != nil { inputs["logicAppId"] = i.ToStringOutput() }
		if i := state.Method; i != nil { inputs["method"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Uri; i != nil { inputs["uri"] = i.ToStringOutput() }
	}
	var resource ActionHttp
	err := ctx.ReadResource("azure:logicapps/actionHttp:ActionHttp", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionHttp resources.
type ActionHttpState struct {
	// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
	Body pulumi.StringInput `pulumi:"body"`
	// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
	Headers pulumi.StringMapInput `pulumi:"headers"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringInput `pulumi:"logicAppId"`
	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
	Method pulumi.StringInput `pulumi:"method"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the URI which will be called when this HTTP Action is triggered.
	Uri pulumi.StringInput `pulumi:"uri"`
}

// The set of arguments for constructing a ActionHttp resource.
type ActionHttpArgs struct {
	// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
	Body pulumi.StringInput `pulumi:"body"`
	// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
	Headers pulumi.StringMapInput `pulumi:"headers"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringInput `pulumi:"logicAppId"`
	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
	Method pulumi.StringInput `pulumi:"method"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the URI which will be called when this HTTP Action is triggered.
	Uri pulumi.StringInput `pulumi:"uri"`
}
