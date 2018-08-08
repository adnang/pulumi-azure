# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class VirtualMachine(pulumi.CustomResource):
    """
    Manages a Virtual Machine.
    
    ~> **NOTE:** Data Disks can be attached either directly on the `azurerm_virtual_machine` resource, or using the `azurerm_virtual_machine_data_disk_attachment` resource - but the two cannot be used together. If both are used against the same Virtual Machine, spurious changes will occur.
    """
    def __init__(__self__, __name__, __opts__=None, availability_set_id=None, boot_diagnostics=None, delete_data_disks_on_termination=None, delete_os_disk_on_termination=None, identity=None, license_type=None, location=None, name=None, network_interface_ids=None, os_profile=None, os_profile_linux_config=None, os_profile_secrets=None, os_profile_windows_config=None, plan=None, primary_network_interface_id=None, resource_group_name=None, storage_data_disks=None, storage_image_reference=None, storage_os_disk=None, tags=None, vm_size=None, zones=None):
        """Create a VirtualMachine resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if availability_set_id and not isinstance(availability_set_id, basestring):
            raise TypeError('Expected property availability_set_id to be a basestring')
        __self__.availability_set_id = availability_set_id
        """
        The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        """
        __props__['availabilitySetId'] = availability_set_id

        if boot_diagnostics and not isinstance(boot_diagnostics, dict):
            raise TypeError('Expected property boot_diagnostics to be a dict')
        __self__.boot_diagnostics = boot_diagnostics
        """
        A `boot_diagnostics` block.
        """
        __props__['bootDiagnostics'] = boot_diagnostics

        if delete_data_disks_on_termination and not isinstance(delete_data_disks_on_termination, bool):
            raise TypeError('Expected property delete_data_disks_on_termination to be a bool')
        __self__.delete_data_disks_on_termination = delete_data_disks_on_termination
        """
        Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        """
        __props__['deleteDataDisksOnTermination'] = delete_data_disks_on_termination

        if delete_os_disk_on_termination and not isinstance(delete_os_disk_on_termination, bool):
            raise TypeError('Expected property delete_os_disk_on_termination to be a bool')
        __self__.delete_os_disk_on_termination = delete_os_disk_on_termination
        """
        Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        """
        __props__['deleteOsDiskOnTermination'] = delete_os_disk_on_termination

        if identity and not isinstance(identity, dict):
            raise TypeError('Expected property identity to be a dict')
        __self__.identity = identity
        """
        A `identity` block.
        """
        __props__['identity'] = identity

        if license_type and not isinstance(license_type, basestring):
            raise TypeError('Expected property license_type to be a basestring')
        __self__.license_type = license_type
        """
        Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
        """
        __props__['licenseType'] = license_type

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
        """
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        Specifies the name of the OS Disk.
        """
        __props__['name'] = name

        if not network_interface_ids:
            raise TypeError('Missing required property network_interface_ids')
        elif not isinstance(network_interface_ids, list):
            raise TypeError('Expected property network_interface_ids to be a list')
        __self__.network_interface_ids = network_interface_ids
        """
        A list of Network Interface ID's which should be associated with the Virtual Machine.
        """
        __props__['networkInterfaceIds'] = network_interface_ids

        if os_profile and not isinstance(os_profile, dict):
            raise TypeError('Expected property os_profile to be a dict')
        __self__.os_profile = os_profile
        """
        An `os_profile` block. Required when `create_option` in the `storage_os_disk` block is set to `FromImage`.
        """
        __props__['osProfile'] = os_profile

        if os_profile_linux_config and not isinstance(os_profile_linux_config, dict):
            raise TypeError('Expected property os_profile_linux_config to be a dict')
        __self__.os_profile_linux_config = os_profile_linux_config
        """
        A `os_profile_linux_config` block.
        """
        __props__['osProfileLinuxConfig'] = os_profile_linux_config

        if os_profile_secrets and not isinstance(os_profile_secrets, list):
            raise TypeError('Expected property os_profile_secrets to be a list')
        __self__.os_profile_secrets = os_profile_secrets
        """
        One or more `os_profile_secrets` blocks.
        """
        __props__['osProfileSecrets'] = os_profile_secrets

        if os_profile_windows_config and not isinstance(os_profile_windows_config, dict):
            raise TypeError('Expected property os_profile_windows_config to be a dict')
        __self__.os_profile_windows_config = os_profile_windows_config
        """
        A `os_profile_windows_config` block.
        """
        __props__['osProfileWindowsConfig'] = os_profile_windows_config

        if plan and not isinstance(plan, dict):
            raise TypeError('Expected property plan to be a dict')
        __self__.plan = plan
        """
        A `plan` block.
        """
        __props__['plan'] = plan

        if primary_network_interface_id and not isinstance(primary_network_interface_id, basestring):
            raise TypeError('Expected property primary_network_interface_id to be a basestring')
        __self__.primary_network_interface_id = primary_network_interface_id
        """
        The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
        """
        __props__['primaryNetworkInterfaceId'] = primary_network_interface_id

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if storage_data_disks and not isinstance(storage_data_disks, list):
            raise TypeError('Expected property storage_data_disks to be a list')
        __self__.storage_data_disks = storage_data_disks
        """
        One or more `storage_data_disk` blocks.
        """
        __props__['storageDataDisks'] = storage_data_disks

        if storage_image_reference and not isinstance(storage_image_reference, dict):
            raise TypeError('Expected property storage_image_reference to be a dict')
        __self__.storage_image_reference = storage_image_reference
        """
        A `storage_image_reference` block.
        """
        __props__['storageImageReference'] = storage_image_reference

        if not storage_os_disk:
            raise TypeError('Missing required property storage_os_disk')
        elif not isinstance(storage_os_disk, dict):
            raise TypeError('Expected property storage_os_disk to be a dict')
        __self__.storage_os_disk = storage_os_disk
        """
        A `storage_os_disk` block.
        """
        __props__['storageOsDisk'] = storage_os_disk

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the Virtual Machine.
        """
        __props__['tags'] = tags

        if not vm_size:
            raise TypeError('Missing required property vm_size')
        elif not isinstance(vm_size, basestring):
            raise TypeError('Expected property vm_size to be a basestring')
        __self__.vm_size = vm_size
        """
        Specifies the [size of the Virtual Machine](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-size-specs/).
        """
        __props__['vmSize'] = vm_size

        if zones and not isinstance(zones, basestring):
            raise TypeError('Expected property zones to be a basestring')
        __self__.zones = zones
        """
        A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
        """
        __props__['zones'] = zones

        super(VirtualMachine, __self__).__init__(
            'azure:compute/virtualMachine:VirtualMachine',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'availabilitySetId' in outs:
            self.availability_set_id = outs['availabilitySetId']
        if 'bootDiagnostics' in outs:
            self.boot_diagnostics = outs['bootDiagnostics']
        if 'deleteDataDisksOnTermination' in outs:
            self.delete_data_disks_on_termination = outs['deleteDataDisksOnTermination']
        if 'deleteOsDiskOnTermination' in outs:
            self.delete_os_disk_on_termination = outs['deleteOsDiskOnTermination']
        if 'identity' in outs:
            self.identity = outs['identity']
        if 'licenseType' in outs:
            self.license_type = outs['licenseType']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'networkInterfaceIds' in outs:
            self.network_interface_ids = outs['networkInterfaceIds']
        if 'osProfile' in outs:
            self.os_profile = outs['osProfile']
        if 'osProfileLinuxConfig' in outs:
            self.os_profile_linux_config = outs['osProfileLinuxConfig']
        if 'osProfileSecrets' in outs:
            self.os_profile_secrets = outs['osProfileSecrets']
        if 'osProfileWindowsConfig' in outs:
            self.os_profile_windows_config = outs['osProfileWindowsConfig']
        if 'plan' in outs:
            self.plan = outs['plan']
        if 'primaryNetworkInterfaceId' in outs:
            self.primary_network_interface_id = outs['primaryNetworkInterfaceId']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'storageDataDisks' in outs:
            self.storage_data_disks = outs['storageDataDisks']
        if 'storageImageReference' in outs:
            self.storage_image_reference = outs['storageImageReference']
        if 'storageOsDisk' in outs:
            self.storage_os_disk = outs['storageOsDisk']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'vmSize' in outs:
            self.vm_size = outs['vmSize']
        if 'zones' in outs:
            self.zones = outs['zones']
