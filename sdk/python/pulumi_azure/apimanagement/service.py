# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Service(pulumi.CustomResource):
    additional_locations: pulumi.Output[list]
    """
    One or more `additional_location` blocks as defined below.

      * `gateway_regional_url` (`str`) - The URL of the Regional Gateway for the API Management Service in the specified region.
      * `location` (`str`) - The name of the Azure Region in which the API Management Service should be expanded to.
      * `public_ip_addresses` (`list`) - Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
    """
    certificates: pulumi.Output[list]
    """
    One or more (up to 10) `certificate` blocks as defined below.

      * `certificatePassword` (`str`) - The password for the certificate.
      * `encodedCertificate` (`str`) - The Base64 Encoded PFX Certificate.
      * `storeName` (`str`) - The name of the Certificate Store where this certificate should be stored. Possible values are `CertificateAuthority` and `Root`.
    """
    gateway_regional_url: pulumi.Output[str]
    """
    The URL of the Regional Gateway for the API Management Service in the specified region.
    """
    gateway_url: pulumi.Output[str]
    """
    The URL of the Gateway for the API Management Service.
    """
    hostname_configuration: pulumi.Output[dict]
    """
    A `hostname_configuration` block as defined below.

      * `developerPortals` (`list`) - One or more `developer_portal` blocks as documented below.
        * `certificate` (`str`) - One or more (up to 10) `certificate` blocks as defined below.
        * `certificatePassword` (`str`) - The password for the certificate.
        * `host_name` (`str`) - The Hostname to use for the Management API.
        * `key_vault_id` (`str`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
        * `negotiateClientCertificate` (`bool`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

      * `managements` (`list`) - One or more `management` blocks as documented below.
        * `certificate` (`str`) - The Base64 Encoded Certificate.
        * `certificatePassword` (`str`) - The password associated with the certificate provided above.
        * `host_name` (`str`) - The Hostname to use for the Management API.
        * `key_vault_id` (`str`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
        * `negotiateClientCertificate` (`bool`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

      * `portals` (`list`) - One or more `portal` blocks as documented below.
        * `certificate` (`str`) - One or more (up to 10) `certificate` blocks as defined below.
        * `certificatePassword` (`str`) - The password for the certificate.
        * `host_name` (`str`) - The Hostname to use for the Management API.
        * `key_vault_id` (`str`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
        * `negotiateClientCertificate` (`bool`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

      * `proxies` (`list`) - One or more `proxy` blocks as documented below.
        * `certificate` (`str`) - The Base64 Encoded Certificate.
        * `certificatePassword` (`str`) - The password associated with the certificate provided above.
        * `defaultSslBinding` (`bool`) - Is the certificate associated with this Hostname the Default SSL Certificate? This is used when an SNI header isn't specified by a client. Defaults to `false`.
        * `host_name` (`str`) - The Hostname to use for the Management API.
        * `key_vault_id` (`str`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
        * `negotiateClientCertificate` (`bool`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

      * `scms` (`list`) - One or more `scm` blocks as documented below.
        * `certificate` (`str`) - One or more (up to 10) `certificate` blocks as defined below.
        * `certificatePassword` (`str`) - The password for the certificate.
        * `host_name` (`str`) - The Hostname to use for the Management API.
        * `key_vault_id` (`str`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
        * `negotiateClientCertificate` (`bool`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.
    """
    identity: pulumi.Output[dict]
    """
    An `identity` block is documented below.

      * `identityIds` (`list`) - A list of IDs for User Assigned Managed Identity resources to be assigned.
      * `principal_id` (`str`) - The Principal ID associated with this Managed Service Identity.
      * `tenant_id` (`str`) - The Tenant ID associated with this Managed Service Identity.
      * `type` (`str`) - Specifies the type of Managed Service Identity that should be configured on this API Management Service. Possible values are `SystemAssigned`, `UserAssigned` or `SystemAssigned, UserAssigned` (to enable both).
    """
    location: pulumi.Output[str]
    """
    The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
    """
    management_api_url: pulumi.Output[str]
    """
    The URL for the Management API associated with this API Management service.
    """
    name: pulumi.Output[str]
    """
    The name of the API Management Service. Changing this forces a new resource to be created.
    """
    notification_sender_email: pulumi.Output[str]
    """
    Email address from which the notification will be sent.
    """
    policy: pulumi.Output[dict]
    """
    A `policy` block as defined below.

      * `xml_content` (`str`) - The XML Content for this Policy.
      * `xml_link` (`str`) - A link to an API Management Policy XML Document, which must be publicly available.
    """
    portal_url: pulumi.Output[str]
    """
    The URL for the Publisher Portal associated with this API Management service.
    """
    private_ip_addresses: pulumi.Output[list]
    protocols: pulumi.Output[dict]
    """
    A `protocols` block as defined below.

      * `enable_http2` (`bool`) - Should HTTP/2 be supported by the API Management Service? Defaults to `false`.
    """
    public_ip_addresses: pulumi.Output[list]
    """
    Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
    """
    publisher_email: pulumi.Output[str]
    """
    The email of publisher/company.
    """
    publisher_name: pulumi.Output[str]
    """
    The name of publisher/company.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
    """
    scm_url: pulumi.Output[str]
    """
    The URL for the SCM (Source Code Management) Endpoint associated with this API Management service.
    """
    security: pulumi.Output[dict]
    """
    A `security` block as defined below.

      * `enableBackendSsl30` (`bool`) - Should SSL 3.0 be enabled on the backend of the gateway? Defaults to `false`.
      * `enableBackendTls10` (`bool`) - Should TLS 1.0 be enabled on the backend of the gateway? Defaults to `false`.
      * `enableBackendTls11` (`bool`) - Should TLS 1.1 be enabled on the backend of the gateway? Defaults to `false`.
      * `enableFrontendSsl30` (`bool`) - Should SSL 3.0 be enabled on the frontend of the gateway? Defaults to `false`.
      * `enableFrontendTls10` (`bool`) - Should TLS 1.0 be enabled on the frontend of the gateway? Defaults to `false`.
      * `enableFrontendTls11` (`bool`) - Should TLS 1.1 be enabled on the frontend of the gateway? Defaults to `false`.
      * `enableTripleDesCiphers` (`bool`) - Should the `TLS_RSA_WITH_3DES_EDE_CBC_SHA` cipher be enabled for alL TLS versions (1.0, 1.1 and 1.2)? Defaults to `false`.
    """
    sign_in: pulumi.Output[dict]
    """
    A `sign_in` block as defined below.

      * `enabled` (`bool`) - Should anonymous users be redirected to the sign in page?
    """
    sign_up: pulumi.Output[dict]
    """
    A `sign_up` block as defined below.

      * `enabled` (`bool`) - Can users sign up on the development portal?
      * `termsOfService` (`dict`) - A `terms_of_service` block as defined below.
        * `consentRequired` (`bool`) - Should the user be asked for consent during sign up?
        * `enabled` (`bool`) - Should Terms of Service be displayed during sign up?.
        * `text` (`str`) - The Terms of Service which users are required to agree to in order to sign up.
    """
    sku_name: pulumi.Output[str]
    """
    `sku_name` is a string consisting of two parts separated by an underscore(\_). The fist part is the `name`, valid values include: `Consumption`, `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags assigned to the resource.
    """
    virtual_network_configuration: pulumi.Output[dict]
    """
    A `virtual_network_configuration` block as defined below. Required when `virtual_network_type` is `External` or `Internal`.

      * `subnet_id` (`str`) - The id of the subnet that will be used for the API Management.
    """
    virtual_network_type: pulumi.Output[str]
    """
    The type of virtual network you want to use, valid values include: `None`, `External`, `Internal`.
    """
    def __init__(__self__, resource_name, opts=None, additional_locations=None, certificates=None, hostname_configuration=None, identity=None, location=None, name=None, notification_sender_email=None, policy=None, protocols=None, publisher_email=None, publisher_name=None, resource_group_name=None, security=None, sign_in=None, sign_up=None, sku_name=None, tags=None, virtual_network_configuration=None, virtual_network_type=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an API Management Service.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.apimanagement.Service("exampleService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            publisher_name="My Company",
            publisher_email="company@exmaple.com",
            sku_name="Developer_1",
            policy={
                "xmlContent": \"\"\"    <policies>
              <inbound />
              <backend />
              <outbound />
              <on-error />
            </policies>
        \"\"\",
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] additional_locations: One or more `additional_location` blocks as defined below.
        :param pulumi.Input[list] certificates: One or more (up to 10) `certificate` blocks as defined below.
        :param pulumi.Input[dict] hostname_configuration: A `hostname_configuration` block as defined below.
        :param pulumi.Input[dict] identity: An `identity` block is documented below.
        :param pulumi.Input[str] location: The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] notification_sender_email: Email address from which the notification will be sent.
        :param pulumi.Input[dict] policy: A `policy` block as defined below.
        :param pulumi.Input[dict] protocols: A `protocols` block as defined below.
        :param pulumi.Input[str] publisher_email: The email of publisher/company.
        :param pulumi.Input[str] publisher_name: The name of publisher/company.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] security: A `security` block as defined below.
        :param pulumi.Input[dict] sign_in: A `sign_in` block as defined below.
        :param pulumi.Input[dict] sign_up: A `sign_up` block as defined below.
        :param pulumi.Input[str] sku_name: `sku_name` is a string consisting of two parts separated by an underscore(\_). The fist part is the `name`, valid values include: `Consumption`, `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
        :param pulumi.Input[dict] tags: A mapping of tags assigned to the resource.
        :param pulumi.Input[dict] virtual_network_configuration: A `virtual_network_configuration` block as defined below. Required when `virtual_network_type` is `External` or `Internal`.
        :param pulumi.Input[str] virtual_network_type: The type of virtual network you want to use, valid values include: `None`, `External`, `Internal`.

        The **additional_locations** object supports the following:

          * `gateway_regional_url` (`pulumi.Input[str]`) - The URL of the Regional Gateway for the API Management Service in the specified region.
          * `location` (`pulumi.Input[str]`) - The name of the Azure Region in which the API Management Service should be expanded to.
          * `public_ip_addresses` (`pulumi.Input[list]`) - Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.

        The **certificates** object supports the following:

          * `certificatePassword` (`pulumi.Input[str]`) - The password for the certificate.
          * `encodedCertificate` (`pulumi.Input[str]`) - The Base64 Encoded PFX Certificate.
          * `storeName` (`pulumi.Input[str]`) - The name of the Certificate Store where this certificate should be stored. Possible values are `CertificateAuthority` and `Root`.

        The **hostname_configuration** object supports the following:

          * `developerPortals` (`pulumi.Input[list]`) - One or more `developer_portal` blocks as documented below.
            * `certificate` (`pulumi.Input[str]`) - One or more (up to 10) `certificate` blocks as defined below.
            * `certificatePassword` (`pulumi.Input[str]`) - The password for the certificate.
            * `host_name` (`pulumi.Input[str]`) - The Hostname to use for the Management API.
            * `key_vault_id` (`pulumi.Input[str]`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
            * `negotiateClientCertificate` (`pulumi.Input[bool]`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

          * `managements` (`pulumi.Input[list]`) - One or more `management` blocks as documented below.
            * `certificate` (`pulumi.Input[str]`) - The Base64 Encoded Certificate.
            * `certificatePassword` (`pulumi.Input[str]`) - The password associated with the certificate provided above.
            * `host_name` (`pulumi.Input[str]`) - The Hostname to use for the Management API.
            * `key_vault_id` (`pulumi.Input[str]`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
            * `negotiateClientCertificate` (`pulumi.Input[bool]`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

          * `portals` (`pulumi.Input[list]`) - One or more `portal` blocks as documented below.
            * `certificate` (`pulumi.Input[str]`) - One or more (up to 10) `certificate` blocks as defined below.
            * `certificatePassword` (`pulumi.Input[str]`) - The password for the certificate.
            * `host_name` (`pulumi.Input[str]`) - The Hostname to use for the Management API.
            * `key_vault_id` (`pulumi.Input[str]`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
            * `negotiateClientCertificate` (`pulumi.Input[bool]`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

          * `proxies` (`pulumi.Input[list]`) - One or more `proxy` blocks as documented below.
            * `certificate` (`pulumi.Input[str]`) - The Base64 Encoded Certificate.
            * `certificatePassword` (`pulumi.Input[str]`) - The password associated with the certificate provided above.
            * `defaultSslBinding` (`pulumi.Input[bool]`) - Is the certificate associated with this Hostname the Default SSL Certificate? This is used when an SNI header isn't specified by a client. Defaults to `false`.
            * `host_name` (`pulumi.Input[str]`) - The Hostname to use for the Management API.
            * `key_vault_id` (`pulumi.Input[str]`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
            * `negotiateClientCertificate` (`pulumi.Input[bool]`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

          * `scms` (`pulumi.Input[list]`) - One or more `scm` blocks as documented below.
            * `certificate` (`pulumi.Input[str]`) - One or more (up to 10) `certificate` blocks as defined below.
            * `certificatePassword` (`pulumi.Input[str]`) - The password for the certificate.
            * `host_name` (`pulumi.Input[str]`) - The Hostname to use for the Management API.
            * `key_vault_id` (`pulumi.Input[str]`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
            * `negotiateClientCertificate` (`pulumi.Input[bool]`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

        The **identity** object supports the following:

          * `identityIds` (`pulumi.Input[list]`) - A list of IDs for User Assigned Managed Identity resources to be assigned.
          * `principal_id` (`pulumi.Input[str]`) - The Principal ID associated with this Managed Service Identity.
          * `tenant_id` (`pulumi.Input[str]`) - The Tenant ID associated with this Managed Service Identity.
          * `type` (`pulumi.Input[str]`) - Specifies the type of Managed Service Identity that should be configured on this API Management Service. Possible values are `SystemAssigned`, `UserAssigned` or `SystemAssigned, UserAssigned` (to enable both).

        The **policy** object supports the following:

          * `xml_content` (`pulumi.Input[str]`) - The XML Content for this Policy.
          * `xml_link` (`pulumi.Input[str]`) - A link to an API Management Policy XML Document, which must be publicly available.

        The **protocols** object supports the following:

          * `enable_http2` (`pulumi.Input[bool]`) - Should HTTP/2 be supported by the API Management Service? Defaults to `false`.

        The **security** object supports the following:

          * `enableBackendSsl30` (`pulumi.Input[bool]`) - Should SSL 3.0 be enabled on the backend of the gateway? Defaults to `false`.
          * `enableBackendTls10` (`pulumi.Input[bool]`) - Should TLS 1.0 be enabled on the backend of the gateway? Defaults to `false`.
          * `enableBackendTls11` (`pulumi.Input[bool]`) - Should TLS 1.1 be enabled on the backend of the gateway? Defaults to `false`.
          * `enableFrontendSsl30` (`pulumi.Input[bool]`) - Should SSL 3.0 be enabled on the frontend of the gateway? Defaults to `false`.
          * `enableFrontendTls10` (`pulumi.Input[bool]`) - Should TLS 1.0 be enabled on the frontend of the gateway? Defaults to `false`.
          * `enableFrontendTls11` (`pulumi.Input[bool]`) - Should TLS 1.1 be enabled on the frontend of the gateway? Defaults to `false`.
          * `enableTripleDesCiphers` (`pulumi.Input[bool]`) - Should the `TLS_RSA_WITH_3DES_EDE_CBC_SHA` cipher be enabled for alL TLS versions (1.0, 1.1 and 1.2)? Defaults to `false`.

        The **sign_in** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Should anonymous users be redirected to the sign in page?

        The **sign_up** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Can users sign up on the development portal?
          * `termsOfService` (`pulumi.Input[dict]`) - A `terms_of_service` block as defined below.
            * `consentRequired` (`pulumi.Input[bool]`) - Should the user be asked for consent during sign up?
            * `enabled` (`pulumi.Input[bool]`) - Should Terms of Service be displayed during sign up?.
            * `text` (`pulumi.Input[str]`) - The Terms of Service which users are required to agree to in order to sign up.

        The **virtual_network_configuration** object supports the following:

          * `subnet_id` (`pulumi.Input[str]`) - The id of the subnet that will be used for the API Management.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['additional_locations'] = additional_locations
            __props__['certificates'] = certificates
            __props__['hostname_configuration'] = hostname_configuration
            __props__['identity'] = identity
            __props__['location'] = location
            __props__['name'] = name
            __props__['notification_sender_email'] = notification_sender_email
            __props__['policy'] = policy
            __props__['protocols'] = protocols
            if publisher_email is None:
                raise TypeError("Missing required property 'publisher_email'")
            __props__['publisher_email'] = publisher_email
            if publisher_name is None:
                raise TypeError("Missing required property 'publisher_name'")
            __props__['publisher_name'] = publisher_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['security'] = security
            __props__['sign_in'] = sign_in
            __props__['sign_up'] = sign_up
            if sku_name is None:
                raise TypeError("Missing required property 'sku_name'")
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            __props__['virtual_network_configuration'] = virtual_network_configuration
            __props__['virtual_network_type'] = virtual_network_type
            __props__['gateway_regional_url'] = None
            __props__['gateway_url'] = None
            __props__['management_api_url'] = None
            __props__['portal_url'] = None
            __props__['private_ip_addresses'] = None
            __props__['public_ip_addresses'] = None
            __props__['scm_url'] = None
        super(Service, __self__).__init__(
            'azure:apimanagement/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, additional_locations=None, certificates=None, gateway_regional_url=None, gateway_url=None, hostname_configuration=None, identity=None, location=None, management_api_url=None, name=None, notification_sender_email=None, policy=None, portal_url=None, private_ip_addresses=None, protocols=None, public_ip_addresses=None, publisher_email=None, publisher_name=None, resource_group_name=None, scm_url=None, security=None, sign_in=None, sign_up=None, sku_name=None, tags=None, virtual_network_configuration=None, virtual_network_type=None):
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] additional_locations: One or more `additional_location` blocks as defined below.
        :param pulumi.Input[list] certificates: One or more (up to 10) `certificate` blocks as defined below.
        :param pulumi.Input[str] gateway_regional_url: The URL of the Regional Gateway for the API Management Service in the specified region.
        :param pulumi.Input[str] gateway_url: The URL of the Gateway for the API Management Service.
        :param pulumi.Input[dict] hostname_configuration: A `hostname_configuration` block as defined below.
        :param pulumi.Input[dict] identity: An `identity` block is documented below.
        :param pulumi.Input[str] location: The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] management_api_url: The URL for the Management API associated with this API Management service.
        :param pulumi.Input[str] name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] notification_sender_email: Email address from which the notification will be sent.
        :param pulumi.Input[dict] policy: A `policy` block as defined below.
        :param pulumi.Input[str] portal_url: The URL for the Publisher Portal associated with this API Management service.
        :param pulumi.Input[dict] protocols: A `protocols` block as defined below.
        :param pulumi.Input[list] public_ip_addresses: Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
        :param pulumi.Input[str] publisher_email: The email of publisher/company.
        :param pulumi.Input[str] publisher_name: The name of publisher/company.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scm_url: The URL for the SCM (Source Code Management) Endpoint associated with this API Management service.
        :param pulumi.Input[dict] security: A `security` block as defined below.
        :param pulumi.Input[dict] sign_in: A `sign_in` block as defined below.
        :param pulumi.Input[dict] sign_up: A `sign_up` block as defined below.
        :param pulumi.Input[str] sku_name: `sku_name` is a string consisting of two parts separated by an underscore(\_). The fist part is the `name`, valid values include: `Consumption`, `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
        :param pulumi.Input[dict] tags: A mapping of tags assigned to the resource.
        :param pulumi.Input[dict] virtual_network_configuration: A `virtual_network_configuration` block as defined below. Required when `virtual_network_type` is `External` or `Internal`.
        :param pulumi.Input[str] virtual_network_type: The type of virtual network you want to use, valid values include: `None`, `External`, `Internal`.

        The **additional_locations** object supports the following:

          * `gateway_regional_url` (`pulumi.Input[str]`) - The URL of the Regional Gateway for the API Management Service in the specified region.
          * `location` (`pulumi.Input[str]`) - The name of the Azure Region in which the API Management Service should be expanded to.
          * `public_ip_addresses` (`pulumi.Input[list]`) - Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.

        The **certificates** object supports the following:

          * `certificatePassword` (`pulumi.Input[str]`) - The password for the certificate.
          * `encodedCertificate` (`pulumi.Input[str]`) - The Base64 Encoded PFX Certificate.
          * `storeName` (`pulumi.Input[str]`) - The name of the Certificate Store where this certificate should be stored. Possible values are `CertificateAuthority` and `Root`.

        The **hostname_configuration** object supports the following:

          * `developerPortals` (`pulumi.Input[list]`) - One or more `developer_portal` blocks as documented below.
            * `certificate` (`pulumi.Input[str]`) - One or more (up to 10) `certificate` blocks as defined below.
            * `certificatePassword` (`pulumi.Input[str]`) - The password for the certificate.
            * `host_name` (`pulumi.Input[str]`) - The Hostname to use for the Management API.
            * `key_vault_id` (`pulumi.Input[str]`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
            * `negotiateClientCertificate` (`pulumi.Input[bool]`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

          * `managements` (`pulumi.Input[list]`) - One or more `management` blocks as documented below.
            * `certificate` (`pulumi.Input[str]`) - The Base64 Encoded Certificate.
            * `certificatePassword` (`pulumi.Input[str]`) - The password associated with the certificate provided above.
            * `host_name` (`pulumi.Input[str]`) - The Hostname to use for the Management API.
            * `key_vault_id` (`pulumi.Input[str]`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
            * `negotiateClientCertificate` (`pulumi.Input[bool]`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

          * `portals` (`pulumi.Input[list]`) - One or more `portal` blocks as documented below.
            * `certificate` (`pulumi.Input[str]`) - One or more (up to 10) `certificate` blocks as defined below.
            * `certificatePassword` (`pulumi.Input[str]`) - The password for the certificate.
            * `host_name` (`pulumi.Input[str]`) - The Hostname to use for the Management API.
            * `key_vault_id` (`pulumi.Input[str]`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
            * `negotiateClientCertificate` (`pulumi.Input[bool]`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

          * `proxies` (`pulumi.Input[list]`) - One or more `proxy` blocks as documented below.
            * `certificate` (`pulumi.Input[str]`) - The Base64 Encoded Certificate.
            * `certificatePassword` (`pulumi.Input[str]`) - The password associated with the certificate provided above.
            * `defaultSslBinding` (`pulumi.Input[bool]`) - Is the certificate associated with this Hostname the Default SSL Certificate? This is used when an SNI header isn't specified by a client. Defaults to `false`.
            * `host_name` (`pulumi.Input[str]`) - The Hostname to use for the Management API.
            * `key_vault_id` (`pulumi.Input[str]`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
            * `negotiateClientCertificate` (`pulumi.Input[bool]`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

          * `scms` (`pulumi.Input[list]`) - One or more `scm` blocks as documented below.
            * `certificate` (`pulumi.Input[str]`) - One or more (up to 10) `certificate` blocks as defined below.
            * `certificatePassword` (`pulumi.Input[str]`) - The password for the certificate.
            * `host_name` (`pulumi.Input[str]`) - The Hostname to use for the Management API.
            * `key_vault_id` (`pulumi.Input[str]`) - The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
            * `negotiateClientCertificate` (`pulumi.Input[bool]`) - Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.

        The **identity** object supports the following:

          * `identityIds` (`pulumi.Input[list]`) - A list of IDs for User Assigned Managed Identity resources to be assigned.
          * `principal_id` (`pulumi.Input[str]`) - The Principal ID associated with this Managed Service Identity.
          * `tenant_id` (`pulumi.Input[str]`) - The Tenant ID associated with this Managed Service Identity.
          * `type` (`pulumi.Input[str]`) - Specifies the type of Managed Service Identity that should be configured on this API Management Service. Possible values are `SystemAssigned`, `UserAssigned` or `SystemAssigned, UserAssigned` (to enable both).

        The **policy** object supports the following:

          * `xml_content` (`pulumi.Input[str]`) - The XML Content for this Policy.
          * `xml_link` (`pulumi.Input[str]`) - A link to an API Management Policy XML Document, which must be publicly available.

        The **protocols** object supports the following:

          * `enable_http2` (`pulumi.Input[bool]`) - Should HTTP/2 be supported by the API Management Service? Defaults to `false`.

        The **security** object supports the following:

          * `enableBackendSsl30` (`pulumi.Input[bool]`) - Should SSL 3.0 be enabled on the backend of the gateway? Defaults to `false`.
          * `enableBackendTls10` (`pulumi.Input[bool]`) - Should TLS 1.0 be enabled on the backend of the gateway? Defaults to `false`.
          * `enableBackendTls11` (`pulumi.Input[bool]`) - Should TLS 1.1 be enabled on the backend of the gateway? Defaults to `false`.
          * `enableFrontendSsl30` (`pulumi.Input[bool]`) - Should SSL 3.0 be enabled on the frontend of the gateway? Defaults to `false`.
          * `enableFrontendTls10` (`pulumi.Input[bool]`) - Should TLS 1.0 be enabled on the frontend of the gateway? Defaults to `false`.
          * `enableFrontendTls11` (`pulumi.Input[bool]`) - Should TLS 1.1 be enabled on the frontend of the gateway? Defaults to `false`.
          * `enableTripleDesCiphers` (`pulumi.Input[bool]`) - Should the `TLS_RSA_WITH_3DES_EDE_CBC_SHA` cipher be enabled for alL TLS versions (1.0, 1.1 and 1.2)? Defaults to `false`.

        The **sign_in** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Should anonymous users be redirected to the sign in page?

        The **sign_up** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Can users sign up on the development portal?
          * `termsOfService` (`pulumi.Input[dict]`) - A `terms_of_service` block as defined below.
            * `consentRequired` (`pulumi.Input[bool]`) - Should the user be asked for consent during sign up?
            * `enabled` (`pulumi.Input[bool]`) - Should Terms of Service be displayed during sign up?.
            * `text` (`pulumi.Input[str]`) - The Terms of Service which users are required to agree to in order to sign up.

        The **virtual_network_configuration** object supports the following:

          * `subnet_id` (`pulumi.Input[str]`) - The id of the subnet that will be used for the API Management.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_locations"] = additional_locations
        __props__["certificates"] = certificates
        __props__["gateway_regional_url"] = gateway_regional_url
        __props__["gateway_url"] = gateway_url
        __props__["hostname_configuration"] = hostname_configuration
        __props__["identity"] = identity
        __props__["location"] = location
        __props__["management_api_url"] = management_api_url
        __props__["name"] = name
        __props__["notification_sender_email"] = notification_sender_email
        __props__["policy"] = policy
        __props__["portal_url"] = portal_url
        __props__["private_ip_addresses"] = private_ip_addresses
        __props__["protocols"] = protocols
        __props__["public_ip_addresses"] = public_ip_addresses
        __props__["publisher_email"] = publisher_email
        __props__["publisher_name"] = publisher_name
        __props__["resource_group_name"] = resource_group_name
        __props__["scm_url"] = scm_url
        __props__["security"] = security
        __props__["sign_in"] = sign_in
        __props__["sign_up"] = sign_up
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        __props__["virtual_network_configuration"] = virtual_network_configuration
        __props__["virtual_network_type"] = virtual_network_type
        return Service(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

