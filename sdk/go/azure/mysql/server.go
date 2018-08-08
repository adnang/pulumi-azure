// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a MySQL Server.
type Server struct {
	s *pulumi.ResourceState
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOpt) (*Server, error) {
	if args == nil || args.AdministratorLogin == nil {
		return nil, errors.New("missing required argument 'AdministratorLogin'")
	}
	if args == nil || args.AdministratorLoginPassword == nil {
		return nil, errors.New("missing required argument 'AdministratorLoginPassword'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil || args.SslEnforcement == nil {
		return nil, errors.New("missing required argument 'SslEnforcement'")
	}
	if args == nil || args.StorageProfile == nil {
		return nil, errors.New("missing required argument 'StorageProfile'")
	}
	if args == nil || args.Version == nil {
		return nil, errors.New("missing required argument 'Version'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["administratorLogin"] = nil
		inputs["administratorLoginPassword"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["sku"] = nil
		inputs["sslEnforcement"] = nil
		inputs["storageProfile"] = nil
		inputs["tags"] = nil
		inputs["version"] = nil
	} else {
		inputs["administratorLogin"] = args.AdministratorLogin
		inputs["administratorLoginPassword"] = args.AdministratorLoginPassword
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["sku"] = args.Sku
		inputs["sslEnforcement"] = args.SslEnforcement
		inputs["storageProfile"] = args.StorageProfile
		inputs["tags"] = args.Tags
		inputs["version"] = args.Version
	}
	inputs["fqdn"] = nil
	s, err := ctx.RegisterResource("azure:mysql/server:Server", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Server{s: s}, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServerState, opts ...pulumi.ResourceOpt) (*Server, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["administratorLogin"] = state.AdministratorLogin
		inputs["administratorLoginPassword"] = state.AdministratorLoginPassword
		inputs["fqdn"] = state.Fqdn
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["sku"] = state.Sku
		inputs["sslEnforcement"] = state.SslEnforcement
		inputs["storageProfile"] = state.StorageProfile
		inputs["tags"] = state.Tags
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("azure:mysql/server:Server", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Server{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Server) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Server) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
func (r *Server) AdministratorLogin() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["administratorLogin"])
}

// The Password associated with the `administrator_login` for the MySQL Server.
func (r *Server) AdministratorLoginPassword() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["administratorLoginPassword"])
}

// The FQDN of the MySQL Server.
func (r *Server) Fqdn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fqdn"])
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *Server) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
func (r *Server) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
func (r *Server) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `sku` block as defined below.
func (r *Server) Sku() *pulumi.Output {
	return r.s.State["sku"]
}

// Specifies if SSL should be enforced on connections. Possible values are `Enforced` and `Disabled`.
func (r *Server) SslEnforcement() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sslEnforcement"])
}

// A `storage_profile` block as defined below.
func (r *Server) StorageProfile() *pulumi.Output {
	return r.s.State["storageProfile"]
}

// A mapping of tags to assign to the resource.
func (r *Server) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Specifies the version of MySQL to use. Valid values are `5.6` and `5.7`. Changing this forces a new resource to be created.
func (r *Server) Version() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering Server resources.
type ServerState struct {
	// The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
	AdministratorLogin interface{}
	// The Password associated with the `administrator_login` for the MySQL Server.
	AdministratorLoginPassword interface{}
	// The FQDN of the MySQL Server.
	Fqdn interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	Name interface{}
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `sku` block as defined below.
	Sku interface{}
	// Specifies if SSL should be enforced on connections. Possible values are `Enforced` and `Disabled`.
	SslEnforcement interface{}
	// A `storage_profile` block as defined below.
	StorageProfile interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the version of MySQL to use. Valid values are `5.6` and `5.7`. Changing this forces a new resource to be created.
	Version interface{}
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
	AdministratorLogin interface{}
	// The Password associated with the `administrator_login` for the MySQL Server.
	AdministratorLoginPassword interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	Name interface{}
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `sku` block as defined below.
	Sku interface{}
	// Specifies if SSL should be enforced on connections. Possible values are `Enforced` and `Disabled`.
	SslEnforcement interface{}
	// A `storage_profile` block as defined below.
	StorageProfile interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the version of MySQL to use. Valid values are `5.6` and `5.7`. Changing this forces a new resource to be created.
	Version interface{}
}
