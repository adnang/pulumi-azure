# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class RunBook(pulumi.CustomResource):
    """
    Manages a Automation Runbook.
    """
    def __init__(__self__, __name__, __opts__=None, account_name=None, description=None, location=None, log_progress=None, log_verbose=None, name=None, publish_content_link=None, resource_group_name=None, runbook_type=None, tags=None):
        """Create a RunBook resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not account_name:
            raise TypeError('Missing required property account_name')
        elif not isinstance(account_name, basestring):
            raise TypeError('Expected property account_name to be a basestring')
        __self__.account_name = account_name
        """
        The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
        """
        __props__['accountName'] = account_name

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        A description for this credential.
        """
        __props__['description'] = description

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        __props__['location'] = location

        if not log_progress:
            raise TypeError('Missing required property log_progress')
        elif not isinstance(log_progress, bool):
            raise TypeError('Expected property log_progress to be a bool')
        __self__.log_progress = log_progress
        """
        Progress log option.
        """
        __props__['logProgress'] = log_progress

        if not log_verbose:
            raise TypeError('Missing required property log_verbose')
        elif not isinstance(log_verbose, bool):
            raise TypeError('Expected property log_verbose to be a bool')
        __self__.log_verbose = log_verbose
        """
        Verbose log option.
        """
        __props__['logVerbose'] = log_verbose

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        Specifies the name of the Runbook. Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if not publish_content_link:
            raise TypeError('Missing required property publish_content_link')
        elif not isinstance(publish_content_link, dict):
            raise TypeError('Expected property publish_content_link to be a dict')
        __self__.publish_content_link = publish_content_link
        """
        The published runbook content link.
        """
        __props__['publishContentLink'] = publish_content_link

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if not runbook_type:
            raise TypeError('Missing required property runbook_type')
        elif not isinstance(runbook_type, basestring):
            raise TypeError('Expected property runbook_type to be a basestring')
        __self__.runbook_type = runbook_type
        """
        The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
        """
        __props__['runbookType'] = runbook_type

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        __props__['tags'] = tags

        super(RunBook, __self__).__init__(
            'azure:automation/runBook:RunBook',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'accountName' in outs:
            self.account_name = outs['accountName']
        if 'description' in outs:
            self.description = outs['description']
        if 'location' in outs:
            self.location = outs['location']
        if 'logProgress' in outs:
            self.log_progress = outs['logProgress']
        if 'logVerbose' in outs:
            self.log_verbose = outs['logVerbose']
        if 'name' in outs:
            self.name = outs['name']
        if 'publishContentLink' in outs:
            self.publish_content_link = outs['publishContentLink']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'runbookType' in outs:
            self.runbook_type = outs['runbookType']
        if 'tags' in outs:
            self.tags = outs['tags']
