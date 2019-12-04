// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Azure Container Service Instance
// 
// > **NOTE:** All arguments including the client secret will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
// 
// > **DEPRECATED:** [Azure Container Service (ACS) has been deprecated by Azure in favour of Azure (Managed) Kubernetes Service (AKS)](https://azure.microsoft.com/en-us/updates/azure-container-service-will-retire-on-january-31-2020/). Support for ACS will be removed in the next major version of the AzureRM Provider (2.0) - and we **strongly recommend** you consider using Azure Kubernetes Service (AKS) for new deployments.
// 
// ## Example Usage (DCOS)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/container_service.html.markdown.
type Service struct {
	pulumi.CustomResourceState

	// A Agent Pool Profile's block as documented below.
	AgentPoolProfile ServiceAgentPoolProfileOutput `pulumi:"agentPoolProfile"`

	// A VM Diagnostics Profile block as documented below.
	DiagnosticsProfile ServiceDiagnosticsProfileOutput `pulumi:"diagnosticsProfile"`

	// A Linux Profile block as documented below.
	LinuxProfile ServiceLinuxProfileOutput `pulumi:"linuxProfile"`

	// The location where the Container Service instance should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// A Master Profile block as documented below.
	MasterProfile ServiceMasterProfileOutput `pulumi:"masterProfile"`

	// Unique name of the agent pool profile in the context of the subscription and resource group.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the Container Orchestration Platform to use. Currently can be either `DCOS`, `Kubernetes` or `Swarm`. Changing this forces a new resource to be created.
	OrchestrationPlatform pulumi.StringOutput `pulumi:"orchestrationPlatform"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// A Service Principal block as documented below.
	ServicePrincipal ServiceServicePrincipalOutput `pulumi:"servicePrincipal"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil || args.AgentPoolProfile == nil {
		return nil, errors.New("missing required argument 'AgentPoolProfile'")
	}
	if args == nil || args.DiagnosticsProfile == nil {
		return nil, errors.New("missing required argument 'DiagnosticsProfile'")
	}
	if args == nil || args.LinuxProfile == nil {
		return nil, errors.New("missing required argument 'LinuxProfile'")
	}
	if args == nil || args.MasterProfile == nil {
		return nil, errors.New("missing required argument 'MasterProfile'")
	}
	if args == nil || args.OrchestrationPlatform == nil {
		return nil, errors.New("missing required argument 'OrchestrationPlatform'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AgentPoolProfile; i != nil { inputs["agentPoolProfile"] = i.ToServiceAgentPoolProfileOutput() }
		if i := args.DiagnosticsProfile; i != nil { inputs["diagnosticsProfile"] = i.ToServiceDiagnosticsProfileOutput() }
		if i := args.LinuxProfile; i != nil { inputs["linuxProfile"] = i.ToServiceLinuxProfileOutput() }
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.MasterProfile; i != nil { inputs["masterProfile"] = i.ToServiceMasterProfileOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.OrchestrationPlatform; i != nil { inputs["orchestrationPlatform"] = i.ToStringOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.ServicePrincipal; i != nil { inputs["servicePrincipal"] = i.ToServiceServicePrincipalOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource Service
	err := ctx.RegisterResource("azure:containerservice/service:Service", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AgentPoolProfile; i != nil { inputs["agentPoolProfile"] = i.ToServiceAgentPoolProfileOutput() }
		if i := state.DiagnosticsProfile; i != nil { inputs["diagnosticsProfile"] = i.ToServiceDiagnosticsProfileOutput() }
		if i := state.LinuxProfile; i != nil { inputs["linuxProfile"] = i.ToServiceLinuxProfileOutput() }
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.MasterProfile; i != nil { inputs["masterProfile"] = i.ToServiceMasterProfileOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.OrchestrationPlatform; i != nil { inputs["orchestrationPlatform"] = i.ToStringOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.ServicePrincipal; i != nil { inputs["servicePrincipal"] = i.ToServiceServicePrincipalOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource Service
	err := ctx.ReadResource("azure:containerservice/service:Service", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type ServiceState struct {
	// A Agent Pool Profile's block as documented below.
	AgentPoolProfile ServiceAgentPoolProfileInput `pulumi:"agentPoolProfile"`
	// A VM Diagnostics Profile block as documented below.
	DiagnosticsProfile ServiceDiagnosticsProfileInput `pulumi:"diagnosticsProfile"`
	// A Linux Profile block as documented below.
	LinuxProfile ServiceLinuxProfileInput `pulumi:"linuxProfile"`
	// The location where the Container Service instance should be created. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// A Master Profile block as documented below.
	MasterProfile ServiceMasterProfileInput `pulumi:"masterProfile"`
	// Unique name of the agent pool profile in the context of the subscription and resource group.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the Container Orchestration Platform to use. Currently can be either `DCOS`, `Kubernetes` or `Swarm`. Changing this forces a new resource to be created.
	OrchestrationPlatform pulumi.StringInput `pulumi:"orchestrationPlatform"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A Service Principal block as documented below.
	ServicePrincipal ServiceServicePrincipalInput `pulumi:"servicePrincipal"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// A Agent Pool Profile's block as documented below.
	AgentPoolProfile ServiceAgentPoolProfileInput `pulumi:"agentPoolProfile"`
	// A VM Diagnostics Profile block as documented below.
	DiagnosticsProfile ServiceDiagnosticsProfileInput `pulumi:"diagnosticsProfile"`
	// A Linux Profile block as documented below.
	LinuxProfile ServiceLinuxProfileInput `pulumi:"linuxProfile"`
	// The location where the Container Service instance should be created. Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// A Master Profile block as documented below.
	MasterProfile ServiceMasterProfileInput `pulumi:"masterProfile"`
	// Unique name of the agent pool profile in the context of the subscription and resource group.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the Container Orchestration Platform to use. Currently can be either `DCOS`, `Kubernetes` or `Swarm`. Changing this forces a new resource to be created.
	OrchestrationPlatform pulumi.StringInput `pulumi:"orchestrationPlatform"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A Service Principal block as documented below.
	ServicePrincipal ServiceServicePrincipalInput `pulumi:"servicePrincipal"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
type ServiceAgentPoolProfile struct {
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count *int `pulumi:"count"`
	// The DNS Prefix given to Agents in this Agent Pool.
	DnsPrefix string `pulumi:"dnsPrefix"`
	Fqdn *string `pulumi:"fqdn"`
	// Unique name of the agent pool profile in the context of the subscription and resource group.
	Name string `pulumi:"name"`
	// The VM Size of each of the Agent Pool VM's (e.g. Standard_F1 / Standard_D2v2).
	VmSize string `pulumi:"vmSize"`
}
var serviceAgentPoolProfileType = reflect.TypeOf((*ServiceAgentPoolProfile)(nil)).Elem()

type ServiceAgentPoolProfileInput interface {
	pulumi.Input

	ToServiceAgentPoolProfileOutput() ServiceAgentPoolProfileOutput
	ToServiceAgentPoolProfileOutputWithContext(ctx context.Context) ServiceAgentPoolProfileOutput
}

type ServiceAgentPoolProfileArgs struct {
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count pulumi.IntInput `pulumi:"count"`
	// The DNS Prefix given to Agents in this Agent Pool.
	DnsPrefix pulumi.StringInput `pulumi:"dnsPrefix"`
	Fqdn pulumi.StringInput `pulumi:"fqdn"`
	// Unique name of the agent pool profile in the context of the subscription and resource group.
	Name pulumi.StringInput `pulumi:"name"`
	// The VM Size of each of the Agent Pool VM's (e.g. Standard_F1 / Standard_D2v2).
	VmSize pulumi.StringInput `pulumi:"vmSize"`
}

func (ServiceAgentPoolProfileArgs) ElementType() reflect.Type {
	return serviceAgentPoolProfileType
}

func (a ServiceAgentPoolProfileArgs) ToServiceAgentPoolProfileOutput() ServiceAgentPoolProfileOutput {
	return pulumi.ToOutput(a).(ServiceAgentPoolProfileOutput)
}

func (a ServiceAgentPoolProfileArgs) ToServiceAgentPoolProfileOutputWithContext(ctx context.Context) ServiceAgentPoolProfileOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServiceAgentPoolProfileOutput)
}

type ServiceAgentPoolProfileOutput struct { *pulumi.OutputState }

// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
func (o ServiceAgentPoolProfileOutput) Count() pulumi.IntOutput {
	return o.Apply(func(v ServiceAgentPoolProfile) int {
		if v.Count == nil { return *new(int) } else { return *v.Count }
	}).(pulumi.IntOutput)
}

// The DNS Prefix given to Agents in this Agent Pool.
func (o ServiceAgentPoolProfileOutput) DnsPrefix() pulumi.StringOutput {
	return o.Apply(func(v ServiceAgentPoolProfile) string {
		return v.DnsPrefix
	}).(pulumi.StringOutput)
}

func (o ServiceAgentPoolProfileOutput) Fqdn() pulumi.StringOutput {
	return o.Apply(func(v ServiceAgentPoolProfile) string {
		if v.Fqdn == nil { return *new(string) } else { return *v.Fqdn }
	}).(pulumi.StringOutput)
}

// Unique name of the agent pool profile in the context of the subscription and resource group.
func (o ServiceAgentPoolProfileOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v ServiceAgentPoolProfile) string {
		return v.Name
	}).(pulumi.StringOutput)
}

// The VM Size of each of the Agent Pool VM's (e.g. Standard_F1 / Standard_D2v2).
func (o ServiceAgentPoolProfileOutput) VmSize() pulumi.StringOutput {
	return o.Apply(func(v ServiceAgentPoolProfile) string {
		return v.VmSize
	}).(pulumi.StringOutput)
}

func (ServiceAgentPoolProfileOutput) ElementType() reflect.Type {
	return serviceAgentPoolProfileType
}

func (o ServiceAgentPoolProfileOutput) ToServiceAgentPoolProfileOutput() ServiceAgentPoolProfileOutput {
	return o
}

func (o ServiceAgentPoolProfileOutput) ToServiceAgentPoolProfileOutputWithContext(ctx context.Context) ServiceAgentPoolProfileOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ServiceAgentPoolProfileOutput{}) }

type ServiceDiagnosticsProfile struct {
	// Should VM Diagnostics be enabled for the Container Service VM's
	Enabled bool `pulumi:"enabled"`
	StorageUri *string `pulumi:"storageUri"`
}
var serviceDiagnosticsProfileType = reflect.TypeOf((*ServiceDiagnosticsProfile)(nil)).Elem()

type ServiceDiagnosticsProfileInput interface {
	pulumi.Input

	ToServiceDiagnosticsProfileOutput() ServiceDiagnosticsProfileOutput
	ToServiceDiagnosticsProfileOutputWithContext(ctx context.Context) ServiceDiagnosticsProfileOutput
}

type ServiceDiagnosticsProfileArgs struct {
	// Should VM Diagnostics be enabled for the Container Service VM's
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	StorageUri pulumi.StringInput `pulumi:"storageUri"`
}

func (ServiceDiagnosticsProfileArgs) ElementType() reflect.Type {
	return serviceDiagnosticsProfileType
}

func (a ServiceDiagnosticsProfileArgs) ToServiceDiagnosticsProfileOutput() ServiceDiagnosticsProfileOutput {
	return pulumi.ToOutput(a).(ServiceDiagnosticsProfileOutput)
}

func (a ServiceDiagnosticsProfileArgs) ToServiceDiagnosticsProfileOutputWithContext(ctx context.Context) ServiceDiagnosticsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServiceDiagnosticsProfileOutput)
}

type ServiceDiagnosticsProfileOutput struct { *pulumi.OutputState }

// Should VM Diagnostics be enabled for the Container Service VM's
func (o ServiceDiagnosticsProfileOutput) Enabled() pulumi.BoolOutput {
	return o.Apply(func(v ServiceDiagnosticsProfile) bool {
		return v.Enabled
	}).(pulumi.BoolOutput)
}

func (o ServiceDiagnosticsProfileOutput) StorageUri() pulumi.StringOutput {
	return o.Apply(func(v ServiceDiagnosticsProfile) string {
		if v.StorageUri == nil { return *new(string) } else { return *v.StorageUri }
	}).(pulumi.StringOutput)
}

func (ServiceDiagnosticsProfileOutput) ElementType() reflect.Type {
	return serviceDiagnosticsProfileType
}

func (o ServiceDiagnosticsProfileOutput) ToServiceDiagnosticsProfileOutput() ServiceDiagnosticsProfileOutput {
	return o
}

func (o ServiceDiagnosticsProfileOutput) ToServiceDiagnosticsProfileOutputWithContext(ctx context.Context) ServiceDiagnosticsProfileOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ServiceDiagnosticsProfileOutput{}) }

type ServiceLinuxProfile struct {
	// The Admin Username for the Cluster.
	AdminUsername string `pulumi:"adminUsername"`
	// An SSH Key block as documented below.
	SshKey ServiceLinuxProfileSshKey `pulumi:"sshKey"`
}
var serviceLinuxProfileType = reflect.TypeOf((*ServiceLinuxProfile)(nil)).Elem()

type ServiceLinuxProfileInput interface {
	pulumi.Input

	ToServiceLinuxProfileOutput() ServiceLinuxProfileOutput
	ToServiceLinuxProfileOutputWithContext(ctx context.Context) ServiceLinuxProfileOutput
}

type ServiceLinuxProfileArgs struct {
	// The Admin Username for the Cluster.
	AdminUsername pulumi.StringInput `pulumi:"adminUsername"`
	// An SSH Key block as documented below.
	SshKey ServiceLinuxProfileSshKeyInput `pulumi:"sshKey"`
}

func (ServiceLinuxProfileArgs) ElementType() reflect.Type {
	return serviceLinuxProfileType
}

func (a ServiceLinuxProfileArgs) ToServiceLinuxProfileOutput() ServiceLinuxProfileOutput {
	return pulumi.ToOutput(a).(ServiceLinuxProfileOutput)
}

func (a ServiceLinuxProfileArgs) ToServiceLinuxProfileOutputWithContext(ctx context.Context) ServiceLinuxProfileOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServiceLinuxProfileOutput)
}

type ServiceLinuxProfileOutput struct { *pulumi.OutputState }

// The Admin Username for the Cluster.
func (o ServiceLinuxProfileOutput) AdminUsername() pulumi.StringOutput {
	return o.Apply(func(v ServiceLinuxProfile) string {
		return v.AdminUsername
	}).(pulumi.StringOutput)
}

// An SSH Key block as documented below.
func (o ServiceLinuxProfileOutput) SshKey() ServiceLinuxProfileSshKeyOutput {
	return o.Apply(func(v ServiceLinuxProfile) ServiceLinuxProfileSshKey {
		return v.SshKey
	}).(ServiceLinuxProfileSshKeyOutput)
}

func (ServiceLinuxProfileOutput) ElementType() reflect.Type {
	return serviceLinuxProfileType
}

func (o ServiceLinuxProfileOutput) ToServiceLinuxProfileOutput() ServiceLinuxProfileOutput {
	return o
}

func (o ServiceLinuxProfileOutput) ToServiceLinuxProfileOutputWithContext(ctx context.Context) ServiceLinuxProfileOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ServiceLinuxProfileOutput{}) }

type ServiceLinuxProfileSshKey struct {
	// The Public SSH Key used to access the cluster.
	KeyData string `pulumi:"keyData"`
}
var serviceLinuxProfileSshKeyType = reflect.TypeOf((*ServiceLinuxProfileSshKey)(nil)).Elem()

type ServiceLinuxProfileSshKeyInput interface {
	pulumi.Input

	ToServiceLinuxProfileSshKeyOutput() ServiceLinuxProfileSshKeyOutput
	ToServiceLinuxProfileSshKeyOutputWithContext(ctx context.Context) ServiceLinuxProfileSshKeyOutput
}

type ServiceLinuxProfileSshKeyArgs struct {
	// The Public SSH Key used to access the cluster.
	KeyData pulumi.StringInput `pulumi:"keyData"`
}

func (ServiceLinuxProfileSshKeyArgs) ElementType() reflect.Type {
	return serviceLinuxProfileSshKeyType
}

func (a ServiceLinuxProfileSshKeyArgs) ToServiceLinuxProfileSshKeyOutput() ServiceLinuxProfileSshKeyOutput {
	return pulumi.ToOutput(a).(ServiceLinuxProfileSshKeyOutput)
}

func (a ServiceLinuxProfileSshKeyArgs) ToServiceLinuxProfileSshKeyOutputWithContext(ctx context.Context) ServiceLinuxProfileSshKeyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServiceLinuxProfileSshKeyOutput)
}

type ServiceLinuxProfileSshKeyOutput struct { *pulumi.OutputState }

// The Public SSH Key used to access the cluster.
func (o ServiceLinuxProfileSshKeyOutput) KeyData() pulumi.StringOutput {
	return o.Apply(func(v ServiceLinuxProfileSshKey) string {
		return v.KeyData
	}).(pulumi.StringOutput)
}

func (ServiceLinuxProfileSshKeyOutput) ElementType() reflect.Type {
	return serviceLinuxProfileSshKeyType
}

func (o ServiceLinuxProfileSshKeyOutput) ToServiceLinuxProfileSshKeyOutput() ServiceLinuxProfileSshKeyOutput {
	return o
}

func (o ServiceLinuxProfileSshKeyOutput) ToServiceLinuxProfileSshKeyOutputWithContext(ctx context.Context) ServiceLinuxProfileSshKeyOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ServiceLinuxProfileSshKeyOutput{}) }

type ServiceMasterProfile struct {
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count *int `pulumi:"count"`
	// The DNS Prefix given to Agents in this Agent Pool.
	DnsPrefix string `pulumi:"dnsPrefix"`
	Fqdn *string `pulumi:"fqdn"`
}
var serviceMasterProfileType = reflect.TypeOf((*ServiceMasterProfile)(nil)).Elem()

type ServiceMasterProfileInput interface {
	pulumi.Input

	ToServiceMasterProfileOutput() ServiceMasterProfileOutput
	ToServiceMasterProfileOutputWithContext(ctx context.Context) ServiceMasterProfileOutput
}

type ServiceMasterProfileArgs struct {
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count pulumi.IntInput `pulumi:"count"`
	// The DNS Prefix given to Agents in this Agent Pool.
	DnsPrefix pulumi.StringInput `pulumi:"dnsPrefix"`
	Fqdn pulumi.StringInput `pulumi:"fqdn"`
}

func (ServiceMasterProfileArgs) ElementType() reflect.Type {
	return serviceMasterProfileType
}

func (a ServiceMasterProfileArgs) ToServiceMasterProfileOutput() ServiceMasterProfileOutput {
	return pulumi.ToOutput(a).(ServiceMasterProfileOutput)
}

func (a ServiceMasterProfileArgs) ToServiceMasterProfileOutputWithContext(ctx context.Context) ServiceMasterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServiceMasterProfileOutput)
}

type ServiceMasterProfileOutput struct { *pulumi.OutputState }

// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
func (o ServiceMasterProfileOutput) Count() pulumi.IntOutput {
	return o.Apply(func(v ServiceMasterProfile) int {
		if v.Count == nil { return *new(int) } else { return *v.Count }
	}).(pulumi.IntOutput)
}

// The DNS Prefix given to Agents in this Agent Pool.
func (o ServiceMasterProfileOutput) DnsPrefix() pulumi.StringOutput {
	return o.Apply(func(v ServiceMasterProfile) string {
		return v.DnsPrefix
	}).(pulumi.StringOutput)
}

func (o ServiceMasterProfileOutput) Fqdn() pulumi.StringOutput {
	return o.Apply(func(v ServiceMasterProfile) string {
		if v.Fqdn == nil { return *new(string) } else { return *v.Fqdn }
	}).(pulumi.StringOutput)
}

func (ServiceMasterProfileOutput) ElementType() reflect.Type {
	return serviceMasterProfileType
}

func (o ServiceMasterProfileOutput) ToServiceMasterProfileOutput() ServiceMasterProfileOutput {
	return o
}

func (o ServiceMasterProfileOutput) ToServiceMasterProfileOutputWithContext(ctx context.Context) ServiceMasterProfileOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ServiceMasterProfileOutput{}) }

type ServiceServicePrincipal struct {
	// The ID for the Service Principal.
	ClientId string `pulumi:"clientId"`
	// The secret password associated with the service principal.
	ClientSecret string `pulumi:"clientSecret"`
}
var serviceServicePrincipalType = reflect.TypeOf((*ServiceServicePrincipal)(nil)).Elem()

type ServiceServicePrincipalInput interface {
	pulumi.Input

	ToServiceServicePrincipalOutput() ServiceServicePrincipalOutput
	ToServiceServicePrincipalOutputWithContext(ctx context.Context) ServiceServicePrincipalOutput
}

type ServiceServicePrincipalArgs struct {
	// The ID for the Service Principal.
	ClientId pulumi.StringInput `pulumi:"clientId"`
	// The secret password associated with the service principal.
	ClientSecret pulumi.StringInput `pulumi:"clientSecret"`
}

func (ServiceServicePrincipalArgs) ElementType() reflect.Type {
	return serviceServicePrincipalType
}

func (a ServiceServicePrincipalArgs) ToServiceServicePrincipalOutput() ServiceServicePrincipalOutput {
	return pulumi.ToOutput(a).(ServiceServicePrincipalOutput)
}

func (a ServiceServicePrincipalArgs) ToServiceServicePrincipalOutputWithContext(ctx context.Context) ServiceServicePrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServiceServicePrincipalOutput)
}

type ServiceServicePrincipalOutput struct { *pulumi.OutputState }

// The ID for the Service Principal.
func (o ServiceServicePrincipalOutput) ClientId() pulumi.StringOutput {
	return o.Apply(func(v ServiceServicePrincipal) string {
		return v.ClientId
	}).(pulumi.StringOutput)
}

// The secret password associated with the service principal.
func (o ServiceServicePrincipalOutput) ClientSecret() pulumi.StringOutput {
	return o.Apply(func(v ServiceServicePrincipal) string {
		return v.ClientSecret
	}).(pulumi.StringOutput)
}

func (ServiceServicePrincipalOutput) ElementType() reflect.Type {
	return serviceServicePrincipalType
}

func (o ServiceServicePrincipalOutput) ToServiceServicePrincipalOutput() ServiceServicePrincipalOutput {
	return o
}

func (o ServiceServicePrincipalOutput) ToServiceServicePrincipalOutputWithContext(ctx context.Context) ServiceServicePrincipalOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ServiceServicePrincipalOutput{}) }

