# Commit Convention

This project follows [Conventional Commits](https://www.conventionalcommits.org/) specifications.

## Format

Each commit message consists of a **header**, an optional **body**, and an optional **footer**.

```
<type>(<optional scope>): <subject>

<optional body>

<optional footer>
```

The **header** is mandatory and must conform to the [Commit Message Header](#commit-message-header) format.

### Commit Message Header

```
<type>(<optional scope>): <subject>
```

The `type` and `subject` fields are mandatory, the `scope` field is optional.

#### Type

Must be one of the following:

* **feat**: A new feature
* **fix**: A bug fix
* **docs**: Documentation only changes
* **style**: Changes that do not affect the meaning of the code (white-space, formatting, etc)
* **refactor**: A code change that neither fixes a bug nor adds a feature
* **perf**: A code change that improves performance
* **test**: Adding missing tests or correcting existing tests
* **build**: Changes that affect the build system or external dependencies
* **ci**: Changes to our CI configuration files and scripts
* **chore**: Other changes that don't modify src or test files

#### Scope

The scope specifies what area of the codebase your commit touches, and is usually the package name. For example:

* **core**: Changes to the main hord package
* **cache**: Changes to the cache package
* **drivers/redis**: Changes to the Redis driver
* **drivers/bbolt**: Changes to the BBolt driver
* **drivers/cassandra**: Changes to the Cassandra driver
* **drivers/hashmap**: Changes to the HashMap driver
* **drivers/mock**: Changes to the Mock driver
* **drivers/nats**: Changes to the NATS driver

#### Subject

The subject is a succinct description of the change:

* Use the imperative, present tense: "change" not "changed" nor "changes"
* Don't capitalize the first letter
* No dot (.) at the end

### Commit Message Body

The body should include the motivation for the change and contrast this with previous behavior.

### Commit Message Footer

The footer should contain any information about **Breaking Changes** and is also the place to reference GitHub issues that this commit closes.

**Breaking Changes** should start with the word `BREAKING CHANGE:` with a space or two newlines. The rest of the commit message is then used for this.

## Examples

```
feat(drivers/redis): add cluster support

Add support for Redis clusters with configurable connection pooling

Closes #123
```

```
fix(drivers/cassandra): resolve connection leak 

When an error occurs during Cassandra query execution, connections were not
properly returned to the pool, causing connection leaks.

BREAKING CHANGE: The Cassandra driver now requires an explicit call to Close()
```

```
docs: update installation instructions
```

```
style(core): fix comment typos
```

```
refactor(cache): optimize key lookup algorithm
```