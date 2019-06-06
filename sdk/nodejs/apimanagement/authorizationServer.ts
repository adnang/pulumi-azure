// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Authorization Server within an API Management Service.
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleAuthorizationServer = new azure.apimanagement.AuthorizationServer("example", {
 *     apiManagementName: azurerm_api_management_example.name,
 *     authorizationEndpoint: "https://example.mydomain.com/client/authorize",
 *     clientId: "42424242-4242-4242-4242-424242424242",
 *     clientRegistrationEndpoint: "https://example.mydomain.com/client/register",
 *     displayName: "Test Server",
 *     grantTypes: ["authorizationCode"],
 *     name: "test-server",
 *     resourceGroupName: azurerm_api_management_example.resourceGroupName,
 * });
 * const exampleApi = pulumi.output(azure.apimanagement.getApi({
 *     apiManagementName: "search-api-management",
 *     name: "search-api",
 *     resourceGroupName: "search-service",
 *     revision: "2",
 * }));
 * ```
 */
export class AuthorizationServer extends pulumi.CustomResource {
    /**
     * Get an existing AuthorizationServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthorizationServerState, opts?: pulumi.CustomResourceOptions): AuthorizationServer {
        return new AuthorizationServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/authorizationServer:AuthorizationServer';

    /**
     * Returns true if the given object is an instance of AuthorizationServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthorizationServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthorizationServer.__pulumiType;
    }

    /**
     * The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * The OAUTH Authorization Endpoint.
     */
    public readonly authorizationEndpoint!: pulumi.Output<string>;
    /**
     * The HTTP Verbs supported by the Authorization Endpoint. Possible values are `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`.
     */
    public readonly authorizationMethods!: pulumi.Output<string[]>;
    /**
     * The mechanism by which Access Tokens are passed to the API. Possible values are `authorizationHeader` and `query`.
     */
    public readonly bearerTokenSendingMethods!: pulumi.Output<string[] | undefined>;
    /**
     * The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are `Basic` and `Body`.
     */
    public readonly clientAuthenticationMethods!: pulumi.Output<string[] | undefined>;
    /**
     * The Client/App ID registered with this Authorization Server.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The URI of page where Client/App Registration is performed for this Authorization Server.
     */
    public readonly clientRegistrationEndpoint!: pulumi.Output<string>;
    /**
     * The Client/App Secret registered with this Authorization Server.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
     */
    public readonly defaultScope!: pulumi.Output<string | undefined>;
    /**
     * A description of the Authorization Server, which may contain HTML formatting tags.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The user-friendly name of this Authorization Server.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Form of Authorization Grants required when requesting an Access Token. Possible values are `authorizationCode`, `clientCredentials`, `implicit` and `resourceOwnerPassword`.
     */
    public readonly grantTypes!: pulumi.Output<string[]>;
    /**
     * The name of this Authorization Server. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The password associated with the Resource Owner.
     */
    public readonly resourceOwnerPassword!: pulumi.Output<string | undefined>;
    /**
     * The username associated with the Resource Owner.
     */
    public readonly resourceOwnerUsername!: pulumi.Output<string | undefined>;
    /**
     * Does this Authorization Server support State? If this is set to `true` the client may use the state parameter to raise protocol security.
     */
    public readonly supportState!: pulumi.Output<boolean | undefined>;
    public readonly tokenBodyParameters!: pulumi.Output<{ name: string, value: string }[] | undefined>;
    /**
     * The OAUTH Token Endpoint.
     */
    public readonly tokenEndpoint!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthorizationServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthorizationServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthorizationServerArgs | AuthorizationServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthorizationServerState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["authorizationEndpoint"] = state ? state.authorizationEndpoint : undefined;
            inputs["authorizationMethods"] = state ? state.authorizationMethods : undefined;
            inputs["bearerTokenSendingMethods"] = state ? state.bearerTokenSendingMethods : undefined;
            inputs["clientAuthenticationMethods"] = state ? state.clientAuthenticationMethods : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientRegistrationEndpoint"] = state ? state.clientRegistrationEndpoint : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["defaultScope"] = state ? state.defaultScope : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["grantTypes"] = state ? state.grantTypes : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["resourceOwnerPassword"] = state ? state.resourceOwnerPassword : undefined;
            inputs["resourceOwnerUsername"] = state ? state.resourceOwnerUsername : undefined;
            inputs["supportState"] = state ? state.supportState : undefined;
            inputs["tokenBodyParameters"] = state ? state.tokenBodyParameters : undefined;
            inputs["tokenEndpoint"] = state ? state.tokenEndpoint : undefined;
        } else {
            const args = argsOrState as AuthorizationServerArgs | undefined;
            if (!args || args.apiManagementName === undefined) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if (!args || args.authorizationEndpoint === undefined) {
                throw new Error("Missing required property 'authorizationEndpoint'");
            }
            if (!args || args.authorizationMethods === undefined) {
                throw new Error("Missing required property 'authorizationMethods'");
            }
            if (!args || args.clientId === undefined) {
                throw new Error("Missing required property 'clientId'");
            }
            if (!args || args.clientRegistrationEndpoint === undefined) {
                throw new Error("Missing required property 'clientRegistrationEndpoint'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.grantTypes === undefined) {
                throw new Error("Missing required property 'grantTypes'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["authorizationEndpoint"] = args ? args.authorizationEndpoint : undefined;
            inputs["authorizationMethods"] = args ? args.authorizationMethods : undefined;
            inputs["bearerTokenSendingMethods"] = args ? args.bearerTokenSendingMethods : undefined;
            inputs["clientAuthenticationMethods"] = args ? args.clientAuthenticationMethods : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientRegistrationEndpoint"] = args ? args.clientRegistrationEndpoint : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["defaultScope"] = args ? args.defaultScope : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["grantTypes"] = args ? args.grantTypes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceOwnerPassword"] = args ? args.resourceOwnerPassword : undefined;
            inputs["resourceOwnerUsername"] = args ? args.resourceOwnerUsername : undefined;
            inputs["supportState"] = args ? args.supportState : undefined;
            inputs["tokenBodyParameters"] = args ? args.tokenBodyParameters : undefined;
            inputs["tokenEndpoint"] = args ? args.tokenEndpoint : undefined;
        }
        super(AuthorizationServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthorizationServer resources.
 */
export interface AuthorizationServerState {
    /**
     * The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * The OAUTH Authorization Endpoint.
     */
    readonly authorizationEndpoint?: pulumi.Input<string>;
    /**
     * The HTTP Verbs supported by the Authorization Endpoint. Possible values are `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`.
     */
    readonly authorizationMethods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The mechanism by which Access Tokens are passed to the API. Possible values are `authorizationHeader` and `query`.
     */
    readonly bearerTokenSendingMethods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are `Basic` and `Body`.
     */
    readonly clientAuthenticationMethods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Client/App ID registered with this Authorization Server.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The URI of page where Client/App Registration is performed for this Authorization Server.
     */
    readonly clientRegistrationEndpoint?: pulumi.Input<string>;
    /**
     * The Client/App Secret registered with this Authorization Server.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
     */
    readonly defaultScope?: pulumi.Input<string>;
    /**
     * A description of the Authorization Server, which may contain HTML formatting tags.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The user-friendly name of this Authorization Server.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Form of Authorization Grants required when requesting an Access Token. Possible values are `authorizationCode`, `clientCredentials`, `implicit` and `resourceOwnerPassword`.
     */
    readonly grantTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of this Authorization Server. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The password associated with the Resource Owner.
     */
    readonly resourceOwnerPassword?: pulumi.Input<string>;
    /**
     * The username associated with the Resource Owner.
     */
    readonly resourceOwnerUsername?: pulumi.Input<string>;
    /**
     * Does this Authorization Server support State? If this is set to `true` the client may use the state parameter to raise protocol security.
     */
    readonly supportState?: pulumi.Input<boolean>;
    readonly tokenBodyParameters?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, value: pulumi.Input<string> }>[]>;
    /**
     * The OAUTH Token Endpoint.
     */
    readonly tokenEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthorizationServer resource.
 */
export interface AuthorizationServerArgs {
    /**
     * The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * The OAUTH Authorization Endpoint.
     */
    readonly authorizationEndpoint: pulumi.Input<string>;
    /**
     * The HTTP Verbs supported by the Authorization Endpoint. Possible values are `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`.
     */
    readonly authorizationMethods: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The mechanism by which Access Tokens are passed to the API. Possible values are `authorizationHeader` and `query`.
     */
    readonly bearerTokenSendingMethods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are `Basic` and `Body`.
     */
    readonly clientAuthenticationMethods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Client/App ID registered with this Authorization Server.
     */
    readonly clientId: pulumi.Input<string>;
    /**
     * The URI of page where Client/App Registration is performed for this Authorization Server.
     */
    readonly clientRegistrationEndpoint: pulumi.Input<string>;
    /**
     * The Client/App Secret registered with this Authorization Server.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
     */
    readonly defaultScope?: pulumi.Input<string>;
    /**
     * A description of the Authorization Server, which may contain HTML formatting tags.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The user-friendly name of this Authorization Server.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Form of Authorization Grants required when requesting an Access Token. Possible values are `authorizationCode`, `clientCredentials`, `implicit` and `resourceOwnerPassword`.
     */
    readonly grantTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of this Authorization Server. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The password associated with the Resource Owner.
     */
    readonly resourceOwnerPassword?: pulumi.Input<string>;
    /**
     * The username associated with the Resource Owner.
     */
    readonly resourceOwnerUsername?: pulumi.Input<string>;
    /**
     * Does this Authorization Server support State? If this is set to `true` the client may use the state parameter to raise protocol security.
     */
    readonly supportState?: pulumi.Input<boolean>;
    readonly tokenBodyParameters?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, value: pulumi.Input<string> }>[]>;
    /**
     * The OAUTH Token Endpoint.
     */
    readonly tokenEndpoint?: pulumi.Input<string>;
}
