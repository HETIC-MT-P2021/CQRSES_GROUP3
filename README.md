# CQRSES_GROUP3

A repository to train on event sourcing, CQRS and DDD.

## Project setup

If this is the first time you run the project, please go to the `.env.dist` file and set the variables according to your wishes

then simply run `make init`

If you need more informations about the available make commands just run `make help`

## Tests

You can run tests using the `make test` command.

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

## API Documentation

### Account

Create new account

```http request
POST /register
Content-Type: application/json
{
    "name": "John Doe",
    "email": "johndoe@gmail.com",
    "password": "*********"
}
```

Log in

```http request
POST /login
Content-Type: application/json
{
    "email": "johndoe@gmail.com",
    "password": "*********",
}
```

> Please note that the token is set inside your cookies and that you need to place it inside the `Authorization` header of your request, prefixed by 'Bearer ' (the space is important).

### Articles

Get article

```http request
GET /api/v1/articles/:id
```

Create new article

```http request
POST /api/v1/articles
Content-Type: application/json
{
    "AuthorID": 1,
    "Title": "The best article you'll read today",
    "Content": "In fact it's the best article because it's only one line."
}
```

Update an existing article

```http request
POST /api/v1/articles/:id
Content-Type: application/json
{
    "AuthorID": 1,
    "Title": "The best article you'll read today",
    "Content": "In fact it's the best article because it's only one line."
}
```

Delete article

```http request
DELETE /api/v1/articles/:id
```

## Stack

| tool | port | note |
| - | - | - |
| postgres | 5432 | store the users |
| go | 8000 | the app |
| go | 8082 | the consumer |
| adminer | 8080 | administrate the db |
| elasticsearch | 9200 | store the articles |
| kibana | 5601 | manage elasticsearch |
| rabbitmq | 15672 - 5672 | queuing system |
