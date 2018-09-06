# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Probe(pulumi.CustomResource):
    """
    Manages a LoadBalancer Probe Resource.
    
    ~> **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached
    """
    def __init__(__self__, __name__, __opts__=None, interval_in_seconds=None, loadbalancer_id=None, location=None, name=None, number_of_probes=None, port=None, protocol=None, request_path=None, resource_group_name=None):
        """Create a Probe resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if interval_in_seconds and not isinstance(interval_in_seconds, int):
            raise TypeError('Expected property interval_in_seconds to be a int')
        __self__.interval_in_seconds = interval_in_seconds
        """
        The interval, in seconds between probes to the backend endpoint for health status. The default value is 15, the minimum value is 5.
        """
        __props__['intervalInSeconds'] = interval_in_seconds

        if not loadbalancer_id:
            raise TypeError('Missing required property loadbalancer_id')
        elif not isinstance(loadbalancer_id, basestring):
            raise TypeError('Expected property loadbalancer_id to be a basestring')
        __self__.loadbalancer_id = loadbalancer_id
        """
        The ID of the LoadBalancer in which to create the NAT Rule.
        """
        __props__['loadbalancerId'] = loadbalancer_id

        if location and not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        Specifies the name of the Probe.
        """
        __props__['name'] = name

        if number_of_probes and not isinstance(number_of_probes, int):
            raise TypeError('Expected property number_of_probes to be a int')
        __self__.number_of_probes = number_of_probes
        """
        The number of failed probe attempts after which the backend endpoint is removed from rotation. The default value is 2. NumberOfProbes multiplied by intervalInSeconds value must be greater or equal to 10.Endpoints are returned to rotation when at least one probe is successful.
        """
        __props__['numberOfProbes'] = number_of_probes

        if not port:
            raise TypeError('Missing required property port')
        elif not isinstance(port, int):
            raise TypeError('Expected property port to be a int')
        __self__.port = port
        """
        Port on which the Probe queries the backend endpoint. Possible values range from 1 to 65535, inclusive.
        """
        __props__['port'] = port

        if protocol and not isinstance(protocol, basestring):
            raise TypeError('Expected property protocol to be a basestring')
        __self__.protocol = protocol
        """
        Specifies the protocol of the end point. Possible values are `Http`, `Https` or `Tcp`. If Tcp is specified, a received ACK is required for the probe to be successful. If Http is specified, a 200 OK response from the specified URI is required for the probe to be successful.
        """
        __props__['protocol'] = protocol

        if request_path and not isinstance(request_path, basestring):
            raise TypeError('Expected property request_path to be a basestring')
        __self__.request_path = request_path
        """
        The URI used for requesting health status from the backend endpoint. Required if protocol is set to Http. Otherwise, it is not allowed.
        """
        __props__['requestPath'] = request_path

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to create the resource.
        """
        __props__['resourceGroupName'] = resource_group_name

        __self__.load_balancer_rules = pulumi.runtime.UNKNOWN

        super(Probe, __self__).__init__(
            'azure:lb/probe:Probe',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'intervalInSeconds' in outs:
            self.interval_in_seconds = outs['intervalInSeconds']
        if 'loadBalancerRules' in outs:
            self.load_balancer_rules = outs['loadBalancerRules']
        if 'loadbalancerId' in outs:
            self.loadbalancer_id = outs['loadbalancerId']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'numberOfProbes' in outs:
            self.number_of_probes = outs['numberOfProbes']
        if 'port' in outs:
            self.port = outs['port']
        if 'protocol' in outs:
            self.protocol = outs['protocol']
        if 'requestPath' in outs:
            self.request_path = outs['requestPath']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
