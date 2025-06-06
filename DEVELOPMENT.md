# Development

Thanks for contributing to Stytch's Management Go library! If you run into trouble, find us in [Slack].

## Setup

1. Clone this repo.
2. To test your changes locally replace the import in your test project's go.mod with `replace github.com/stytchauth/stytch-management-go => ../stytch-management-go` where `../stytch-management-go` is the path to your cloned copy of stytch-management-go.


## Local Testing

1. Create a workspace management key and secret from your Stytch workspace
1. Use the template in .env.example to create a file .env with these credentials
1. run `make test`

If using an IDE to test, you must add that environment to your test setup, otherwise the tests will be skipped. 

There are helper functions built in to our tests (see `DisposableProject()` in client_test.go) that will create temporary projects and then delete them in order to test all the endpoints. This will not affect any existing projects.

## Issues and Pull Requests

Please file issues in this repo. We don't have an issue template yet, but for now, say whatever you think is important!

If you have non-trivial changes you'd like us to incorporate, please open an issue first so we can discuss the changes before starting on a pull request. (It's fine to start with the PR for a typo or simple bug.) If we think the changes align with the direction of the project, we'll either ask you to open the PR or assign someone on the Stytch team to make the changes.

When you're ready for someone to look at your issue or PR, assign `@stytchauth/client-libraries` (GitHub should do this automatically). If we don't acknowledge it within one business day, please escalate it by tagging `@stytchauth/engineering` in a comment or letting us know in [Slack].

[Slack]: https://stytch.slack.com/join/shared_invite/zt-2f0fi1ruu-ub~HGouWRmPARM1MTwPESA
