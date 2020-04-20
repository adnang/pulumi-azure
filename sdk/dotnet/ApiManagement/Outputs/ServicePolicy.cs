// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Outputs
{

    [OutputType]
    public sealed class ServicePolicy
    {
        /// <summary>
        /// The XML Content for this Policy.
        /// </summary>
        public readonly string? XmlContent;
        /// <summary>
        /// A link to an API Management Policy XML Document, which must be publicly available.
        /// </summary>
        public readonly string? XmlLink;

        [OutputConstructor]
        private ServicePolicy(
            string? xmlContent,

            string? xmlLink)
        {
            XmlContent = xmlContent;
            XmlLink = xmlLink;
        }
    }
}