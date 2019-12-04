// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a custom virtual machine image that can be used to create virtual machines.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/image.html.markdown.
type Image struct {
	pulumi.CustomResourceState

	// One or more `dataDisk` elements as defined below.
	DataDisks ImageDataDisksArrayOutput `pulumi:"dataDisks"`

	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`

	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`

	// One or more `osDisk` elements as defined below.
	OsDisk ImageOsDiskOutput `pulumi:"osDisk"`

	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`

	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId pulumi.StringOutput `pulumi:"sourceVirtualMachineId"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient pulumi.BoolOutput `pulumi:"zoneResilient"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.DataDisks; i != nil { inputs["dataDisks"] = i.ToImageDataDisksArrayOutput() }
		if i := args.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.OsDisk; i != nil { inputs["osDisk"] = i.ToImageOsDiskOutput() }
		if i := args.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := args.SourceVirtualMachineId; i != nil { inputs["sourceVirtualMachineId"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.ZoneResilient; i != nil { inputs["zoneResilient"] = i.ToBoolOutput() }
	}
	var resource Image
	err := ctx.RegisterResource("azure:compute/image:Image", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.DataDisks; i != nil { inputs["dataDisks"] = i.ToImageDataDisksArrayOutput() }
		if i := state.Location; i != nil { inputs["location"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.OsDisk; i != nil { inputs["osDisk"] = i.ToImageOsDiskOutput() }
		if i := state.ResourceGroupName; i != nil { inputs["resourceGroupName"] = i.ToStringOutput() }
		if i := state.SourceVirtualMachineId; i != nil { inputs["sourceVirtualMachineId"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.ZoneResilient; i != nil { inputs["zoneResilient"] = i.ToBoolOutput() }
	}
	var resource Image
	err := ctx.ReadResource("azure:compute/image:Image", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type ImageState struct {
	// One or more `dataDisk` elements as defined below.
	DataDisks ImageDataDisksArrayInput `pulumi:"dataDisks"`
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// One or more `osDisk` elements as defined below.
	OsDisk ImageOsDiskInput `pulumi:"osDisk"`
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId pulumi.StringInput `pulumi:"sourceVirtualMachineId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient pulumi.BoolInput `pulumi:"zoneResilient"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// One or more `dataDisk` elements as defined below.
	DataDisks ImageDataDisksArrayInput `pulumi:"dataDisks"`
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// One or more `osDisk` elements as defined below.
	OsDisk ImageOsDiskInput `pulumi:"osDisk"`
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId pulumi.StringInput `pulumi:"sourceVirtualMachineId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient pulumi.BoolInput `pulumi:"zoneResilient"`
}
type ImageDataDisks struct {
	// Specifies the URI in Azure storage of the blob that you want to use to create the image.
	BlobUri *string `pulumi:"blobUri"`
	// Specifies the caching mode as `ReadWrite`, `ReadOnly`, or `None`. The default is `None`.
	Caching *string `pulumi:"caching"`
	// Specifies the logical unit number of the data disk.
	Lun *int `pulumi:"lun"`
	// Specifies the ID of the managed disk resource that you want to use to create the image.
	ManagedDiskId *string `pulumi:"managedDiskId"`
	// Specifies the size of the image to be created. The target size can't be smaller than the source size.
	SizeGb *int `pulumi:"sizeGb"`
}
var imageDataDisksType = reflect.TypeOf((*ImageDataDisks)(nil)).Elem()

type ImageDataDisksInput interface {
	pulumi.Input

	ToImageDataDisksOutput() ImageDataDisksOutput
	ToImageDataDisksOutputWithContext(ctx context.Context) ImageDataDisksOutput
}

type ImageDataDisksArgs struct {
	// Specifies the URI in Azure storage of the blob that you want to use to create the image.
	BlobUri pulumi.StringInput `pulumi:"blobUri"`
	// Specifies the caching mode as `ReadWrite`, `ReadOnly`, or `None`. The default is `None`.
	Caching pulumi.StringInput `pulumi:"caching"`
	// Specifies the logical unit number of the data disk.
	Lun pulumi.IntInput `pulumi:"lun"`
	// Specifies the ID of the managed disk resource that you want to use to create the image.
	ManagedDiskId pulumi.StringInput `pulumi:"managedDiskId"`
	// Specifies the size of the image to be created. The target size can't be smaller than the source size.
	SizeGb pulumi.IntInput `pulumi:"sizeGb"`
}

func (ImageDataDisksArgs) ElementType() reflect.Type {
	return imageDataDisksType
}

func (a ImageDataDisksArgs) ToImageDataDisksOutput() ImageDataDisksOutput {
	return pulumi.ToOutput(a).(ImageDataDisksOutput)
}

func (a ImageDataDisksArgs) ToImageDataDisksOutputWithContext(ctx context.Context) ImageDataDisksOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ImageDataDisksOutput)
}

type ImageDataDisksOutput struct { *pulumi.OutputState }

// Specifies the URI in Azure storage of the blob that you want to use to create the image.
func (o ImageDataDisksOutput) BlobUri() pulumi.StringOutput {
	return o.Apply(func(v ImageDataDisks) string {
		if v.BlobUri == nil { return *new(string) } else { return *v.BlobUri }
	}).(pulumi.StringOutput)
}

// Specifies the caching mode as `ReadWrite`, `ReadOnly`, or `None`. The default is `None`.
func (o ImageDataDisksOutput) Caching() pulumi.StringOutput {
	return o.Apply(func(v ImageDataDisks) string {
		if v.Caching == nil { return *new(string) } else { return *v.Caching }
	}).(pulumi.StringOutput)
}

// Specifies the logical unit number of the data disk.
func (o ImageDataDisksOutput) Lun() pulumi.IntOutput {
	return o.Apply(func(v ImageDataDisks) int {
		if v.Lun == nil { return *new(int) } else { return *v.Lun }
	}).(pulumi.IntOutput)
}

// Specifies the ID of the managed disk resource that you want to use to create the image.
func (o ImageDataDisksOutput) ManagedDiskId() pulumi.StringOutput {
	return o.Apply(func(v ImageDataDisks) string {
		if v.ManagedDiskId == nil { return *new(string) } else { return *v.ManagedDiskId }
	}).(pulumi.StringOutput)
}

// Specifies the size of the image to be created. The target size can't be smaller than the source size.
func (o ImageDataDisksOutput) SizeGb() pulumi.IntOutput {
	return o.Apply(func(v ImageDataDisks) int {
		if v.SizeGb == nil { return *new(int) } else { return *v.SizeGb }
	}).(pulumi.IntOutput)
}

func (ImageDataDisksOutput) ElementType() reflect.Type {
	return imageDataDisksType
}

func (o ImageDataDisksOutput) ToImageDataDisksOutput() ImageDataDisksOutput {
	return o
}

func (o ImageDataDisksOutput) ToImageDataDisksOutputWithContext(ctx context.Context) ImageDataDisksOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ImageDataDisksOutput{}) }

var imageDataDisksArrayType = reflect.TypeOf((*[]ImageDataDisks)(nil)).Elem()

type ImageDataDisksArrayInput interface {
	pulumi.Input

	ToImageDataDisksArrayOutput() ImageDataDisksArrayOutput
	ToImageDataDisksArrayOutputWithContext(ctx context.Context) ImageDataDisksArrayOutput
}

type ImageDataDisksArrayArgs []ImageDataDisksInput

func (ImageDataDisksArrayArgs) ElementType() reflect.Type {
	return imageDataDisksArrayType
}

func (a ImageDataDisksArrayArgs) ToImageDataDisksArrayOutput() ImageDataDisksArrayOutput {
	return pulumi.ToOutput(a).(ImageDataDisksArrayOutput)
}

func (a ImageDataDisksArrayArgs) ToImageDataDisksArrayOutputWithContext(ctx context.Context) ImageDataDisksArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ImageDataDisksArrayOutput)
}

type ImageDataDisksArrayOutput struct { *pulumi.OutputState }

func (o ImageDataDisksArrayOutput) Index(i pulumi.IntInput) ImageDataDisksOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) ImageDataDisks {
		return vs[0].([]ImageDataDisks)[vs[1].(int)]
	}).(ImageDataDisksOutput)
}

func (ImageDataDisksArrayOutput) ElementType() reflect.Type {
	return imageDataDisksArrayType
}

func (o ImageDataDisksArrayOutput) ToImageDataDisksArrayOutput() ImageDataDisksArrayOutput {
	return o
}

func (o ImageDataDisksArrayOutput) ToImageDataDisksArrayOutputWithContext(ctx context.Context) ImageDataDisksArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ImageDataDisksArrayOutput{}) }

type ImageOsDisk struct {
	// Specifies the URI in Azure storage of the blob that you want to use to create the image.
	BlobUri *string `pulumi:"blobUri"`
	// Specifies the caching mode as `ReadWrite`, `ReadOnly`, or `None`. The default is `None`.
	Caching *string `pulumi:"caching"`
	// Specifies the ID of the managed disk resource that you want to use to create the image.
	ManagedDiskId *string `pulumi:"managedDiskId"`
	// Specifies the state of the operating system contained in the blob. Currently, the only value is Generalized.
	OsState *string `pulumi:"osState"`
	// Specifies the type of operating system contained in the the virtual machine image. Possible values are: Windows or Linux.
	OsType *string `pulumi:"osType"`
	// Specifies the size of the image to be created. The target size can't be smaller than the source size.
	SizeGb *int `pulumi:"sizeGb"`
}
var imageOsDiskType = reflect.TypeOf((*ImageOsDisk)(nil)).Elem()

type ImageOsDiskInput interface {
	pulumi.Input

	ToImageOsDiskOutput() ImageOsDiskOutput
	ToImageOsDiskOutputWithContext(ctx context.Context) ImageOsDiskOutput
}

type ImageOsDiskArgs struct {
	// Specifies the URI in Azure storage of the blob that you want to use to create the image.
	BlobUri pulumi.StringInput `pulumi:"blobUri"`
	// Specifies the caching mode as `ReadWrite`, `ReadOnly`, or `None`. The default is `None`.
	Caching pulumi.StringInput `pulumi:"caching"`
	// Specifies the ID of the managed disk resource that you want to use to create the image.
	ManagedDiskId pulumi.StringInput `pulumi:"managedDiskId"`
	// Specifies the state of the operating system contained in the blob. Currently, the only value is Generalized.
	OsState pulumi.StringInput `pulumi:"osState"`
	// Specifies the type of operating system contained in the the virtual machine image. Possible values are: Windows or Linux.
	OsType pulumi.StringInput `pulumi:"osType"`
	// Specifies the size of the image to be created. The target size can't be smaller than the source size.
	SizeGb pulumi.IntInput `pulumi:"sizeGb"`
}

func (ImageOsDiskArgs) ElementType() reflect.Type {
	return imageOsDiskType
}

func (a ImageOsDiskArgs) ToImageOsDiskOutput() ImageOsDiskOutput {
	return pulumi.ToOutput(a).(ImageOsDiskOutput)
}

func (a ImageOsDiskArgs) ToImageOsDiskOutputWithContext(ctx context.Context) ImageOsDiskOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ImageOsDiskOutput)
}

type ImageOsDiskOutput struct { *pulumi.OutputState }

// Specifies the URI in Azure storage of the blob that you want to use to create the image.
func (o ImageOsDiskOutput) BlobUri() pulumi.StringOutput {
	return o.Apply(func(v ImageOsDisk) string {
		if v.BlobUri == nil { return *new(string) } else { return *v.BlobUri }
	}).(pulumi.StringOutput)
}

// Specifies the caching mode as `ReadWrite`, `ReadOnly`, or `None`. The default is `None`.
func (o ImageOsDiskOutput) Caching() pulumi.StringOutput {
	return o.Apply(func(v ImageOsDisk) string {
		if v.Caching == nil { return *new(string) } else { return *v.Caching }
	}).(pulumi.StringOutput)
}

// Specifies the ID of the managed disk resource that you want to use to create the image.
func (o ImageOsDiskOutput) ManagedDiskId() pulumi.StringOutput {
	return o.Apply(func(v ImageOsDisk) string {
		if v.ManagedDiskId == nil { return *new(string) } else { return *v.ManagedDiskId }
	}).(pulumi.StringOutput)
}

// Specifies the state of the operating system contained in the blob. Currently, the only value is Generalized.
func (o ImageOsDiskOutput) OsState() pulumi.StringOutput {
	return o.Apply(func(v ImageOsDisk) string {
		if v.OsState == nil { return *new(string) } else { return *v.OsState }
	}).(pulumi.StringOutput)
}

// Specifies the type of operating system contained in the the virtual machine image. Possible values are: Windows or Linux.
func (o ImageOsDiskOutput) OsType() pulumi.StringOutput {
	return o.Apply(func(v ImageOsDisk) string {
		if v.OsType == nil { return *new(string) } else { return *v.OsType }
	}).(pulumi.StringOutput)
}

// Specifies the size of the image to be created. The target size can't be smaller than the source size.
func (o ImageOsDiskOutput) SizeGb() pulumi.IntOutput {
	return o.Apply(func(v ImageOsDisk) int {
		if v.SizeGb == nil { return *new(int) } else { return *v.SizeGb }
	}).(pulumi.IntOutput)
}

func (ImageOsDiskOutput) ElementType() reflect.Type {
	return imageOsDiskType
}

func (o ImageOsDiskOutput) ToImageOsDiskOutput() ImageOsDiskOutput {
	return o
}

func (o ImageOsDiskOutput) ToImageOsDiskOutputWithContext(ctx context.Context) ImageOsDiskOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ImageOsDiskOutput{}) }

