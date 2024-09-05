# Stytch Management Go Library 

The Stytch Management Go library makes it easy to use Stytch's Programmatic Workspace Actions API via Go. 

This library is tested with go 1.22.

## Install

```
$ go get github.com/stytchauth/stytch-management-go/v1
```

## Pre-requisites

You need your Stytch Management API Credentials from the workspace management section of your [Stytch Dashboard](https://stytch.com/dashboard/settings/management-api).

**Note:** This key will allow you to perform read and write actions on your workspace,
potentially deleting important Stytch resources like projects or secrets that are in use.

## Usage

This library supports project-level actions on the following resources:

- [x] Projects
- [x] Email Templates
- [x] JWT Templates
- [x] Password strength configuration
- [x] Project Metrics
- [x] Project Secrets
- [x] Public Tokens
- [x] RBAC Policies
- [x] Redirect URLs
- [x] SDK Configuration

## Examples

Create a new API client:

```go

    // Set your Stytch Management API credentials as env variables
    keyID := os.Getenv("STYTCH_WORKSPACE_KEY_ID")
    keySecret := os.Getenv("STYTCH_WORKSPACE_KEY_SECRET")
    
    client := api.NewClient(keyID, keySecret)
    ctx := context.Background()

```

Create a new b2b project:

```go
    // Send the request
    res, err := client.Projects.Create(ctx, projects.CreateRequest{
        ProjectName: "My new project",
        Vertical: projects.VerticalB2B,
    })
    
    // Get the new project information
    // This is used in examples below
    newProject := res.Project
```

Create a new project secret in the test project:

```go
    res, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
        ProjectID: newProject.TestProjectID,
    })
```

Get all public tokens in the live project

```go
    resp, err := client.PublicTokens.GetAll(ctx, publictokens.GetAllRequest{
        ProjectID: newProject.LiveProjectID,
    })
```

Delete a redirect URL

```go
    res, err := client.RedirectURLs.Delete(ctx, redirecturls.DeleteRequest{
        ProjectID: newProject.TestProjectID,
        URL: "http://localhost:3000/authenticate",
    })
```

## Documentation

All request and response components are typed. There are docstrings for request and 
response attributes for all methods, as well as helper structs. 
These can all be found in the [models](./pkg/models) directory. 

Note that in some cases the methods might be specific to live or test projects. The 
documentation/docstrings will mention that where applicable.

## Support

If you've found a bug, [open an issue](https://github.com/stytchauth/stytch-management-go/issues/new)!

If you have questions or want help troubleshooting, join us in [Slack](https://stytch.com/docs/resources/support/overview) or email support@stytch.com.

If you've found a security vulnerability, please follow our [responsible disclosure instructions](https://stytch.com/docs/resources/security-and-trust/security#:~:text=Responsible%20disclosure%20program).

## Development

See [DEVELOPMENT.md](DEVELOPMENT.md)

## Code of Conduct

Everyone interacting in the Stytch project's codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](CODE_OF_CONDUCT.md).