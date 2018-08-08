# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Cluster(pulumi.CustomResource):
    """
    Manage a Service Fabric Cluster.
    """
    def __init__(__self__, __name__, __opts__=None, add_on_features=None, certificate=None, client_certificate_thumbprint=None, diagnostics_config=None, fabric_settings=None, location=None, management_endpoint=None, name=None, node_types=None, reliability_level=None, resource_group_name=None, tags=None, upgrade_mode=None, vm_image=None):
        """Create a Cluster resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if add_on_features and not isinstance(add_on_features, list):
            raise TypeError('Expected property add_on_features to be a list')
        __self__.add_on_features = add_on_features
        """
        A List of one or more features which should be enabled, such as `DnsService`.
        """
        __props__['addOnFeatures'] = add_on_features

        if certificate and not isinstance(certificate, dict):
            raise TypeError('Expected property certificate to be a dict')
        __self__.certificate = certificate
        """
        A `certificate` block as defined below.
        """
        __props__['certificate'] = certificate

        if client_certificate_thumbprint and not isinstance(client_certificate_thumbprint, dict):
            raise TypeError('Expected property client_certificate_thumbprint to be a dict')
        __self__.client_certificate_thumbprint = client_certificate_thumbprint
        """
        A `client_certificate_thumbprint` block as defined below.
        """
        __props__['clientCertificateThumbprint'] = client_certificate_thumbprint

        if diagnostics_config and not isinstance(diagnostics_config, dict):
            raise TypeError('Expected property diagnostics_config to be a dict')
        __self__.diagnostics_config = diagnostics_config
        """
        A `diagnostics_config` block as defined below. Changing this forces a new resource to be created.
        """
        __props__['diagnosticsConfig'] = diagnostics_config

        if fabric_settings and not isinstance(fabric_settings, list):
            raise TypeError('Expected property fabric_settings to be a list')
        __self__.fabric_settings = fabric_settings
        """
        One or more `fabric_settings` blocks as defined below.
        """
        __props__['fabricSettings'] = fabric_settings

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
        """
        __props__['location'] = location

        if not management_endpoint:
            raise TypeError('Missing required property management_endpoint')
        elif not isinstance(management_endpoint, basestring):
            raise TypeError('Expected property management_endpoint to be a basestring')
        __self__.management_endpoint = management_endpoint
        """
        Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
        """
        __props__['managementEndpoint'] = management_endpoint

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the Node Type. Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if not node_types:
            raise TypeError('Missing required property node_types')
        elif not isinstance(node_types, list):
            raise TypeError('Expected property node_types to be a list')
        __self__.node_types = node_types
        """
        One or more `node_type` blocks as defined below.
        """
        __props__['nodeTypes'] = node_types

        if not reliability_level:
            raise TypeError('Missing required property reliability_level')
        elif not isinstance(reliability_level, basestring):
            raise TypeError('Expected property reliability_level to be a basestring')
        __self__.reliability_level = reliability_level
        """
        Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
        """
        __props__['reliabilityLevel'] = reliability_level

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        __props__['tags'] = tags

        if not upgrade_mode:
            raise TypeError('Missing required property upgrade_mode')
        elif not isinstance(upgrade_mode, basestring):
            raise TypeError('Expected property upgrade_mode to be a basestring')
        __self__.upgrade_mode = upgrade_mode
        """
        Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
        """
        __props__['upgradeMode'] = upgrade_mode

        if not vm_image:
            raise TypeError('Missing required property vm_image')
        elif not isinstance(vm_image, basestring):
            raise TypeError('Expected property vm_image to be a basestring')
        __self__.vm_image = vm_image
        """
        Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
        """
        __props__['vmImage'] = vm_image

        __self__.cluster_endpoint = pulumi.runtime.UNKNOWN
        """
        The Cluster Endpoint for this Service Fabric Cluster.
        """

        super(Cluster, __self__).__init__(
            'azure:servicefabric/cluster:Cluster',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'addOnFeatures' in outs:
            self.add_on_features = outs['addOnFeatures']
        if 'certificate' in outs:
            self.certificate = outs['certificate']
        if 'clientCertificateThumbprint' in outs:
            self.client_certificate_thumbprint = outs['clientCertificateThumbprint']
        if 'clusterEndpoint' in outs:
            self.cluster_endpoint = outs['clusterEndpoint']
        if 'diagnosticsConfig' in outs:
            self.diagnostics_config = outs['diagnosticsConfig']
        if 'fabricSettings' in outs:
            self.fabric_settings = outs['fabricSettings']
        if 'location' in outs:
            self.location = outs['location']
        if 'managementEndpoint' in outs:
            self.management_endpoint = outs['managementEndpoint']
        if 'name' in outs:
            self.name = outs['name']
        if 'nodeTypes' in outs:
            self.node_types = outs['nodeTypes']
        if 'reliabilityLevel' in outs:
            self.reliability_level = outs['reliabilityLevel']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'upgradeMode' in outs:
            self.upgrade_mode = outs['upgradeMode']
        if 'vmImage' in outs:
            self.vm_image = outs['vmImage']
