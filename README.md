# CQRSES_GROUP3

A repository to train on event sourcing, CQRS and DDD.
[Google Drive for CQRSES_GROUP3](https://drive.google.com/drive/folders/1YXihm3yICXDTGh6KhaKwunQL6gave4u4?usp=sharing)

## Project setup

If this is the first time you run the project, please go to the `.env.dist` file and set the variables according to your wishes

then simply run `make init`

If you need more informations about the available make commands just run `make help`

### Docker Setup

Refer to the [DOCKER.md](DOCKER.md) for setup


## Useful resources

Google Drive: [CQRSES_GROUP3](https://drive.google.com/drive/folders/1YXihm3yICXDTGh6KhaKwunQL6gave4u4?usp=sharing)
> All the technical documentations and schemes.

[Specifications board](https://github.com/HETIC-MT-P2021/CQRSES_GROUP3/projects/1)
> All the Epic and User stories

[Application board](https://github.com/HETIC-MT-P2021/CQRSES_GROUP3/projects/2)
> All the technical tasks

## Specification process

1. The process follow a simple logic: first we create an Epic issue in the [specifications board](https://github.com/HETIC-MT-P2021/CQRSES_GROUP3/projects/1) to reference all the correspondent User/Dev-Stories. Some of them are focused for developers, others for users.
2. Those Epic issues will be described with their initiatives and goals and captures a large amount of work granularized into User/Dev-stories.
3. Then the User/Dev-stories will be created with the desired behaviour for a specific feature.
4. Finally, the User/Dev-stories will be breakdown into technical tasks referenced as feature request in the issues list.

> All of those issues follow the stages `Todo`, `In Progress`, `Done`

## Development process

1. Most of the time a ticket on the [application board](https://github.com/HETIC-MT-P2021/CQRSES_GROUP3/projects/2) in the `Todo` column is assign to a team member.
2. When a member starts working on the ticket, he should move the concerned ticket to `In Progress`.
3. Then he creates a branch specifically for this ticket with a name that follows the [conventions specified below](#branch-naming-convention).
4. Commit regularly at each significant step with unambiguous commit messages (see [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md) file).
5. Create a merge request that follows the [conventions specified below](#pull-requests-pr) to the develop branch.
6. On the [application board](https://github.com/HETIC-MT-P2021/CQRSES_GROUP3/projects/2), move the ticket to the status `In Review`
7. Assign the merge request to a maintainer
8. It may take some back and forth before your pull request will be validated
9. Your pull request will be then merged into develop branch and the concerned ticket will be moved to `Done`

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

## Contributors

<table>
  <tr>
    <td align="center">
    <a href="https://github.com/jasongauvin">
      <img src="https://avatars1.githubusercontent.com/u/41618366?s=400&u=b970ed03cbb921ce1312ef86b39093e4fa0be7e3&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Jason Gauvin</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/JackMaarek/">
      <img src="https://avatars3.githubusercontent.com/u/28316928?s=400&u=3cdfb5b0683245ad333a39cfca3a5251f3829824&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Jacques Maarek</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/Para0234">
      <img src="https://avatars.githubusercontent.com/u/31258019?s=460&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Benoît Covet</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/edwinvautier">
      <img src="https://avatars3.githubusercontent.com/u/35581502?s=460&u=d9096f90151f35552d9adcd57bacaee366f0aaef&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Edwin Vautier</b></sub>
    </a>
    </td>
  </tr>
</table>
