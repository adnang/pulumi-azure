// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kusto

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) Database Principal
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/kusto_database_principal.html.markdown.
type DatabasePrincipal struct {
	pulumi.CustomResourceState

	// The app id, if not empty, of the principal.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created. 
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The email, if not empty, of the principal.
	Email pulumi.StringOutput `pulumi:"email"`
	// The fully qualified name of the principal.
	FullyQualifiedName pulumi.StringOutput `pulumi:"fullyQualifiedName"`
	// The name of the Kusto Database Principal.
	Name pulumi.StringOutput `pulumi:"name"`
	// An Object ID of a User, Group, or App. Changing this forces a new resource to be created. 
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
	Role pulumi.StringOutput `pulumi:"role"`
	// Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabasePrincipal registers a new resource with the given unique name, arguments, and options.
func NewDatabasePrincipal(ctx *pulumi.Context,
	name string, args *DatabasePrincipalArgs, opts ...pulumi.ResourceOption) (*DatabasePrincipal, error) {
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.ObjectId == nil {
		return nil, errors.New("missing required argument 'ObjectId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &DatabasePrincipalArgs{}
	}
	var resource DatabasePrincipal
	err := ctx.RegisterResource("azure:kusto/databasePrincipal:DatabasePrincipal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabasePrincipal gets an existing DatabasePrincipal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabasePrincipal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabasePrincipalState, opts ...pulumi.ResourceOption) (*DatabasePrincipal, error) {
	var resource DatabasePrincipal
	err := ctx.ReadResource("azure:kusto/databasePrincipal:DatabasePrincipal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabasePrincipal resources.
type databasePrincipalState struct {
	// The app id, if not empty, of the principal.
	AppId *string `pulumi:"appId"`
	// The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created. 
	ClientId *string `pulumi:"clientId"`
	// Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
	ClusterName *string `pulumi:"clusterName"`
	// Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// The email, if not empty, of the principal.
	Email *string `pulumi:"email"`
	// The fully qualified name of the principal.
	FullyQualifiedName *string `pulumi:"fullyQualifiedName"`
	// The name of the Kusto Database Principal.
	Name *string `pulumi:"name"`
	// An Object ID of a User, Group, or App. Changing this forces a new resource to be created. 
	ObjectId *string `pulumi:"objectId"`
	// Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
	Role *string `pulumi:"role"`
	// Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	Type *string `pulumi:"type"`
}

type DatabasePrincipalState struct {
	// The app id, if not empty, of the principal.
	AppId pulumi.StringPtrInput
	// The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created. 
	ClientId pulumi.StringPtrInput
	// Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringPtrInput
	// Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// The email, if not empty, of the principal.
	Email pulumi.StringPtrInput
	// The fully qualified name of the principal.
	FullyQualifiedName pulumi.StringPtrInput
	// The name of the Kusto Database Principal.
	Name pulumi.StringPtrInput
	// An Object ID of a User, Group, or App. Changing this forces a new resource to be created. 
	ObjectId pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
	Role pulumi.StringPtrInput
	// Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	Type pulumi.StringPtrInput
}

func (DatabasePrincipalState) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePrincipalState)(nil)).Elem()
}

type databasePrincipalArgs struct {
	// The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created. 
	ClientId string `pulumi:"clientId"`
	// Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
	ClusterName string `pulumi:"clusterName"`
	// Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// An Object ID of a User, Group, or App. Changing this forces a new resource to be created. 
	ObjectId string `pulumi:"objectId"`
	// Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
	Role string `pulumi:"role"`
	// Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DatabasePrincipal resource.
type DatabasePrincipalArgs struct {
	// The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created. 
	ClientId pulumi.StringInput
	// Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput
	// Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// An Object ID of a User, Group, or App. Changing this forces a new resource to be created. 
	ObjectId pulumi.StringInput
	// Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
	Role pulumi.StringInput
	// Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	Type pulumi.StringInput
}

func (DatabasePrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePrincipalArgs)(nil)).Elem()
}

