# Commit Convention

`Tag`: `Message`

The `Tag` should be in the list above

The `Scope` should define the context of the affected changes.

The `Message` should not be confused with git commit message.

The `Tag` is one of the following:

- `build:`-> Changes that affect the build system or external dependencies (docker, npm, make…)

- `feat:`-> Added new functionality

- `fix:`-> Bug fix

- `refactor:`-> Modifications which brings neither new functionality nor performance improvement

- `docs:`-> Writing or updating documentation

- `test:`-> Adding or modifying tests

- The `scope` should define the context of the affected changes.

- The `message` summaries description of the change in one sentence.

Examples:

```
feat (frontend): Added /category route to access images by category.
fix (database): Added migration to correct category structure.
```
