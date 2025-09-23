# Migrating from Stytch Management Go SDK v2 to v3

This guide helps you migrate from v2 to v3 of the Stytch Management Go SDK. v3 introduces significant architectural changes that provide a better developer experience and support Stytch custom environments.

For more details on how the management API changed, see our [PWA docs](https://stytch.com/docs/workspace-management/pwa/overview).

> [!WARNING]
> The v3 of this SDK is currently in development and only available as an alpha version, with unnanounced breaking changes. For the current major release, see our [v2 branch](https://github.com/stytchauth/stytch-management-go/tree/v2). 

## Key Changes

### 1. Project architecture: from Live/Test to environments

**V2** 
- Used a single `Project` resource to manage both live and test projects
- Used uuid-style identifiers for live and test project IDs

**V3**
- `Project` is now a resource that captures related environments
  - It has a name, a vertical, and an identifier (slug)
  - The slug is immutable and configurable at creation
  - The slug is unique for the workspace
- `Environment` holds the configuration for one instance of the project
  - An environment can be of type live or type test 
  - A maximum of one live environment is supported
  - It's possible to have many test environments
  - Each environment has an identifier, also called a slug
  - The slug is immutable and configurable at creation
  - The slug is unique for the project

- Most endpoints now require both the project slug and the environment slug instead of what in V2 used to be the project's live project ID or test project ID

### 2. Schema changes

Nesting is no longer used in most write request payloads. For example, in JWT Templates, this was the SetRequest object:

```
// v2 required nesting of the request payload inside a "jwt_template" key
// even though the endpoint only supports setting jwt templates.

type SetRequest struct {
	ProjectID string `json:"project_id"`
	JWTTemplate JWTTemplate `json:"jwt_template"`
}
```

In v3, the fields in the JWTTemplate struct are now the top-most level of the payload, with no need for further nesting.

```
type SetRequest struct {
	ProjectSlug string `json:"-"`
	EnvironmentSlug string `json:"-"`
	JWTTemplateType JWTTemplateType `json:"jwt_template_type"`
	TemplateContent string `json:"template_content"`
	CustomAudience string `json:"custom_audience"`
}
```

- Optional fields in write methods are pointers
- Read-only fields are no longer part of the write request struct
- There is validation to avoid empty resource IDs in read methods

### 3. New and deleted resources

v3 includes new resource clients:
- `Environments` - Manage environments within projects

v3 renames:
- `ProjectMetrics` - Replaced by `EnvironmentMetrics`

### 4. Email Templates

**v2**

In v2, email templates only accepted the live project ID, and it modified both the live and test email templates. 

**v3**

In v3, email templates are a project-scoped resource, meaning it does not take an environment. Any write/delete action on the `EmailTemplate` client will affect all email templates in all environments equally.

v3 also introduces support for default email templates.

*Note: If an email template is managed via PWA and then changed manually in a single environment, further PWA request will overwrite those changes and make all environments identical.*

### 5. Better error handling and validation

In v3 of the SDK as well as the management API we revamped all aspects of error handling and validation. 

## Migration Checklist

1. Update import paths to v3
2. Replace v2 project configuration with v3 environment configuration
3. Replace `ProjectID` parameters with `ProjectSlug` and `EnvironmentSlug` in most resources
4. Update any relevant schemas that no longer use nested payloads
5. Update metrics calls from `ProjectMetrics` to `EnvironmentMetrics`

## Questions 

### Can I still use stytch-management-go v2?

Yes. V2 will be considered deprecated once v3 GA is released, with no new features or functionality added. Only critical security patches and bugs will be fixed. 

*Note: v1 is no longer maintained.*

### How long will stytch-management-go v2 be supported?

See our documentation on our workspace management API for a deprecation calendar. 

### If I use custom environments and have more than one test environment, what is v2's test environment?

The test environment configuration in v2 is the oldest non-deleted test environment. There is no mechanism to configure other test environments in v2. The only way to manage multiple environments' configurations via the stytch-management-go SDK is by upgrading to v3.

## Support

If you encounter issues during migration, please:
- Check the [v3 documentation](README.md) for updated examples
- Check the [Stytch Management API docs](https://stytch.com/docs/workspace-management/pwa/overview) for further documentation and information
- File an issue on [GitHub](https://github.com/stytchauth/stytch-management-go/issues)
- Contact support via [Slack](https://stytch.com/docs/resources/support/overview) or email support@stytch.com