// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Inputs
{

    public sealed class ServiceHostnameConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("developerPortals")]
        private InputList<Inputs.ServiceHostnameConfigurationDeveloperPortalArgs>? _developerPortals;

        /// <summary>
        /// One or more `developer_portal` blocks as documented below.
        /// </summary>
        public InputList<Inputs.ServiceHostnameConfigurationDeveloperPortalArgs> DeveloperPortals
        {
            get => _developerPortals ?? (_developerPortals = new InputList<Inputs.ServiceHostnameConfigurationDeveloperPortalArgs>());
            set => _developerPortals = value;
        }

        [Input("managements")]
        private InputList<Inputs.ServiceHostnameConfigurationManagementArgs>? _managements;

        /// <summary>
        /// One or more `management` blocks as documented below.
        /// </summary>
        public InputList<Inputs.ServiceHostnameConfigurationManagementArgs> Managements
        {
            get => _managements ?? (_managements = new InputList<Inputs.ServiceHostnameConfigurationManagementArgs>());
            set => _managements = value;
        }

        [Input("portals")]
        private InputList<Inputs.ServiceHostnameConfigurationPortalArgs>? _portals;

        /// <summary>
        /// One or more `portal` blocks as documented below.
        /// </summary>
        public InputList<Inputs.ServiceHostnameConfigurationPortalArgs> Portals
        {
            get => _portals ?? (_portals = new InputList<Inputs.ServiceHostnameConfigurationPortalArgs>());
            set => _portals = value;
        }

        [Input("proxies")]
        private InputList<Inputs.ServiceHostnameConfigurationProxyArgs>? _proxies;

        /// <summary>
        /// One or more `proxy` blocks as documented below.
        /// </summary>
        public InputList<Inputs.ServiceHostnameConfigurationProxyArgs> Proxies
        {
            get => _proxies ?? (_proxies = new InputList<Inputs.ServiceHostnameConfigurationProxyArgs>());
            set => _proxies = value;
        }

        [Input("scms")]
        private InputList<Inputs.ServiceHostnameConfigurationScmArgs>? _scms;

        /// <summary>
        /// One or more `scm` blocks as documented below.
        /// </summary>
        public InputList<Inputs.ServiceHostnameConfigurationScmArgs> Scms
        {
            get => _scms ?? (_scms = new InputList<Inputs.ServiceHostnameConfigurationScmArgs>());
            set => _scms = value;
        }

        public ServiceHostnameConfigurationArgs()
        {
        }
    }
}
