# Stytch Management Go Library (v3)

The Stytch Management Go library makes it easy to use Stytch's Programmatic Workspace Actions API via Go. 

This library is tested with go 1.22.

## Install

```
$ go get github.com/stytchauth/stytch-management-go/v3
```

## Pre-requisites

You need your Stytch Management API Credentials from the workspace management section of your 
[Stytch Dashboard](https://stytch.com/dashboard/settings/management-api).

**Note:** This key will allow you to perform read and write actions on your workspace,
potentially deleting important Stytch resources like projects or secrets that are in use.

## Usage

This library supports project- and environment-level actions on the following resources:

- [x] Projects
- [x] Environments
- [x] Country Code Allowlists
- [x] Email Templates
- [x] Environment Metrics
- [x] Event Log Streaming
- [x] JWT Templates
- [x] Password Strength Configuration
- [x] Public Tokens
- [x] RBAC Policies
- [x] Redirect URLs
- [x] SDK Configuration
- [x] (Environment) Secrets
- [x] Trusted Token Profiles / PEM Files

## Examples

Create a new API client:

```go
    // Set your Stytch Management API credentials as env variables.
    keyID := os.Getenv("STYTCH_WORKSPACE_KEY_ID")
    keySecret := os.Getenv("STYTCH_WORKSPACE_KEY_SECRET")
    
    client := api.NewClient(keyID, keySecret)
    ctx := context.Background()
```

Create a new B2B project:

```go
    // Send the request.
    resp, err := client.Projects.Create(ctx, projects.CreateRequest{
        Name: "My new project",
        Vertical: projects.VerticalB2B,
    })
    
    // Get the new project information. This is used in the examples below.
    newProject := resp.Project
```

Get the live environment in the new project:

```go
    resp, err := client.Environments.GetAll(ctx, environments.GetAllRequest{
        ProjectID: newProject.Project,
    })

    var liveEnv string
    for _, env := range resp.Environments
        if env.EnvironmentType == environments.EnvironmentTypeLive {
            liveEnv = env.Environment
			break
        }
    }
```

Alternatively, create a new custom environment in the new project:

```go
    resp, err := client.Environments.Create(ctx, environments.CreateRequest{
        Project: newProject.Project,
        Name: "My custom environment",
		Type: EnvironmentTypeTest,
    })

    customEnv := resp.Environment
```

Create a new secret in the live environment:

```go
    resp, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
        Project: newProject.Project,
		Environment: liveEnv,
    })
```

Get all public tokens in the custom test environment:

```go
    resp, err := client.PublicTokens.GetAll(ctx, publictokens.GetAllRequest{
        Project: newProject.Project,
        Environment: customEnv,
    })
```

Delete a redirect URL in the custom test environment:

```go
    resp, err := client.RedirectURLs.Delete(ctx, redirecturls.DeleteRequest{
        Project: newProject.Project,
		Environment: customEnv,
        URL: "http://localhost:3000/authenticate",
    })
```

## Documentation

All request and response components are typed. There are docstrings for request and response
attributes for all methods, as well as helper structs. These can all be found in the [models](./pkg/models)
directory.

## Support

If you've found a bug, [open an issue](https://github.com/stytchauth/stytch-management-go/issues/new)!

If you have questions or want help troubleshooting, join us in [Slack](https://stytch.com/docs/resources/support/overview) or email support@stytch.com.

If you've found a security vulnerability, please follow our [responsible disclosure instructions](https://stytch.com/docs/resources/security-and-trust/security#:~:text=Responsible%20disclosure%20program).

## Development

See [DEVELOPMENT.md](DEVELOPMENT.md)

## Code of Conduct

Everyone interacting in the Stytch project's codebases, issue trackers, chat rooms and mailing lists
is expected to follow the [code of conduct](CODE_OF_CONDUCT.md).
