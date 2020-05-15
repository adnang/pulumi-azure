// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql.Inputs
{

    public sealed class ServerAzureadAdministratorGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Required)  The login username of the Azure AD Administrator of this SQL Server.
        /// </summary>
        [Input("loginUsername", required: true)]
        public Input<string> LoginUsername { get; set; } = null!;

        /// <summary>
        /// (Required) The object id of the Azure AD Administrator of this SQL Server.
        /// </summary>
        [Input("objectId", required: true)]
        public Input<string> ObjectId { get; set; } = null!;

        /// <summary>
        /// (Optional) The tenant id of the Azure AD Administrator of this SQL Server.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public ServerAzureadAdministratorGetArgs()
        {
        }
    }
}
