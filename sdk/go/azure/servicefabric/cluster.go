// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manage a Service Fabric Cluster.
type Cluster struct {
	s *pulumi.ResourceState
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ManagementEndpoint == nil {
		return nil, errors.New("missing required argument 'ManagementEndpoint'")
	}
	if args == nil || args.NodeTypes == nil {
		return nil, errors.New("missing required argument 'NodeTypes'")
	}
	if args == nil || args.ReliabilityLevel == nil {
		return nil, errors.New("missing required argument 'ReliabilityLevel'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.UpgradeMode == nil {
		return nil, errors.New("missing required argument 'UpgradeMode'")
	}
	if args == nil || args.VmImage == nil {
		return nil, errors.New("missing required argument 'VmImage'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["addOnFeatures"] = nil
		inputs["azureActiveDirectory"] = nil
		inputs["certificate"] = nil
		inputs["clientCertificateThumbprints"] = nil
		inputs["clusterCodeVersion"] = nil
		inputs["diagnosticsConfig"] = nil
		inputs["fabricSettings"] = nil
		inputs["location"] = nil
		inputs["managementEndpoint"] = nil
		inputs["name"] = nil
		inputs["nodeTypes"] = nil
		inputs["reliabilityLevel"] = nil
		inputs["resourceGroupName"] = nil
		inputs["reverseProxyCertificate"] = nil
		inputs["tags"] = nil
		inputs["upgradeMode"] = nil
		inputs["vmImage"] = nil
	} else {
		inputs["addOnFeatures"] = args.AddOnFeatures
		inputs["azureActiveDirectory"] = args.AzureActiveDirectory
		inputs["certificate"] = args.Certificate
		inputs["clientCertificateThumbprints"] = args.ClientCertificateThumbprints
		inputs["clusterCodeVersion"] = args.ClusterCodeVersion
		inputs["diagnosticsConfig"] = args.DiagnosticsConfig
		inputs["fabricSettings"] = args.FabricSettings
		inputs["location"] = args.Location
		inputs["managementEndpoint"] = args.ManagementEndpoint
		inputs["name"] = args.Name
		inputs["nodeTypes"] = args.NodeTypes
		inputs["reliabilityLevel"] = args.ReliabilityLevel
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["reverseProxyCertificate"] = args.ReverseProxyCertificate
		inputs["tags"] = args.Tags
		inputs["upgradeMode"] = args.UpgradeMode
		inputs["vmImage"] = args.VmImage
	}
	inputs["clusterEndpoint"] = nil
	s, err := ctx.RegisterResource("azure:servicefabric/cluster:Cluster", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterState, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addOnFeatures"] = state.AddOnFeatures
		inputs["azureActiveDirectory"] = state.AzureActiveDirectory
		inputs["certificate"] = state.Certificate
		inputs["clientCertificateThumbprints"] = state.ClientCertificateThumbprints
		inputs["clusterCodeVersion"] = state.ClusterCodeVersion
		inputs["clusterEndpoint"] = state.ClusterEndpoint
		inputs["diagnosticsConfig"] = state.DiagnosticsConfig
		inputs["fabricSettings"] = state.FabricSettings
		inputs["location"] = state.Location
		inputs["managementEndpoint"] = state.ManagementEndpoint
		inputs["name"] = state.Name
		inputs["nodeTypes"] = state.NodeTypes
		inputs["reliabilityLevel"] = state.ReliabilityLevel
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["reverseProxyCertificate"] = state.ReverseProxyCertificate
		inputs["tags"] = state.Tags
		inputs["upgradeMode"] = state.UpgradeMode
		inputs["vmImage"] = state.VmImage
	}
	s, err := ctx.ReadResource("azure:servicefabric/cluster:Cluster", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Cluster) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Cluster) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A List of one or more features which should be enabled, such as `DnsService`.
func (r *Cluster) AddOnFeatures() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["addOnFeatures"])
}

// An `azure_active_directory` block as defined below. Changing this forces a new resource to be created.
func (r *Cluster) AzureActiveDirectory() *pulumi.Output {
	return r.s.State["azureActiveDirectory"]
}

// A `certificate` block as defined below.
func (r *Cluster) Certificate() *pulumi.Output {
	return r.s.State["certificate"]
}

// One or two `client_certificate_thumbprint` blocks as defined below.
func (r *Cluster) ClientCertificateThumbprints() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["clientCertificateThumbprints"])
}

// Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
func (r *Cluster) ClusterCodeVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterCodeVersion"])
}

// The Cluster Endpoint for this Service Fabric Cluster.
func (r *Cluster) ClusterEndpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterEndpoint"])
}

// A `diagnostics_config` block as defined below. Changing this forces a new resource to be created.
func (r *Cluster) DiagnosticsConfig() *pulumi.Output {
	return r.s.State["diagnosticsConfig"]
}

// One or more `fabric_settings` blocks as defined below.
func (r *Cluster) FabricSettings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["fabricSettings"])
}

// Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
func (r *Cluster) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
func (r *Cluster) ManagementEndpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["managementEndpoint"])
}

// The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
func (r *Cluster) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// One or more `node_type` blocks as defined below.
func (r *Cluster) NodeTypes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["nodeTypes"])
}

// Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
func (r *Cluster) ReliabilityLevel() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["reliabilityLevel"])
}

// The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
func (r *Cluster) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `reverse_proxy_certificate` block as defined below.
func (r *Cluster) ReverseProxyCertificate() *pulumi.Output {
	return r.s.State["reverseProxyCertificate"]
}

// A mapping of tags to assign to the resource.
func (r *Cluster) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
func (r *Cluster) UpgradeMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["upgradeMode"])
}

// Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
func (r *Cluster) VmImage() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vmImage"])
}

// Input properties used for looking up and filtering Cluster resources.
type ClusterState struct {
	// A List of one or more features which should be enabled, such as `DnsService`.
	AddOnFeatures interface{}
	// An `azure_active_directory` block as defined below. Changing this forces a new resource to be created.
	AzureActiveDirectory interface{}
	// A `certificate` block as defined below.
	Certificate interface{}
	// One or two `client_certificate_thumbprint` blocks as defined below.
	ClientCertificateThumbprints interface{}
	// Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
	ClusterCodeVersion interface{}
	// The Cluster Endpoint for this Service Fabric Cluster.
	ClusterEndpoint interface{}
	// A `diagnostics_config` block as defined below. Changing this forces a new resource to be created.
	DiagnosticsConfig interface{}
	// One or more `fabric_settings` blocks as defined below.
	FabricSettings interface{}
	// Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
	ManagementEndpoint interface{}
	// The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
	Name interface{}
	// One or more `node_type` blocks as defined below.
	NodeTypes interface{}
	// Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
	ReliabilityLevel interface{}
	// The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `reverse_proxy_certificate` block as defined below.
	ReverseProxyCertificate interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
	UpgradeMode interface{}
	// Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
	VmImage interface{}
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// A List of one or more features which should be enabled, such as `DnsService`.
	AddOnFeatures interface{}
	// An `azure_active_directory` block as defined below. Changing this forces a new resource to be created.
	AzureActiveDirectory interface{}
	// A `certificate` block as defined below.
	Certificate interface{}
	// One or two `client_certificate_thumbprint` blocks as defined below.
	ClientCertificateThumbprints interface{}
	// Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
	ClusterCodeVersion interface{}
	// A `diagnostics_config` block as defined below. Changing this forces a new resource to be created.
	DiagnosticsConfig interface{}
	// One or more `fabric_settings` blocks as defined below.
	FabricSettings interface{}
	// Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
	ManagementEndpoint interface{}
	// The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
	Name interface{}
	// One or more `node_type` blocks as defined below.
	NodeTypes interface{}
	// Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
	ReliabilityLevel interface{}
	// The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `reverse_proxy_certificate` block as defined below.
	ReverseProxyCertificate interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
	UpgradeMode interface{}
	// Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
	VmImage interface{}
}
