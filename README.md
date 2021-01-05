# CQRSES_GROUP3

A repository to train on event sourcing, CQRS and DDD.

## Project setup

Generate RSA keys for the token authentication :

```sh
    # use the following password: private_key
    openssl genrsa -des3 -out private.pem 2048
    openssl rsa -in private.pem -outform PEM -pubout -out public.pem

```

Update your environment. Create a `.env` file thanks to the `.env.dist` example.

build app :

```sh
    docker-compose up --build
```

## Branch naming convention

You branch should have a name that reflects it's purpose.

It should use the same guidelines as [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md) (`feat`, `fix`, `build`, `perf`, `docs`), followed by an underscore (`_`) and a very quick summary of the subject in [kebab case][1].

Example: `feat_add-image-tag-database-relation`.

## Pull requests (PR)

Pull requests in this project follow two conventions, you will need to use the templates available in the [ISSUE_TEMPLATE](.github/ISSUE_TEMPLATE) folder :

- Adding a new feature should use the [FEATURE_REQUEST](.github/ISSUE_TEMPLATE/FEATURE_REQUEST.md) template.
- Reporting a bug should use the [BUG_REPORT](.github/ISSUE_TEMPLATE/bug_report.md) template.

If your pull request is still work in progress, please add "WIP: " (Work In Progress) in front of the title, therefor you inform the maintainers that your work is not done, and we can't merge it.

The naming of the PR should follow the same rules as the [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md)

## Linter

We use go linter [gofmt](https://blog.golang.org/gofmt) to automatically formats the source code.
