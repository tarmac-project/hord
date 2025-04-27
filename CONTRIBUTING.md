# Contributing to Hord

Thank you for being so interested in helping develop Hord. The time, skills, and perspectives you contribute to this project are valued.

## Issues and Proposals

Bugs, Proposals, & Feature Requests are all welcome. To get started, please open an issue via GitHub. Please provide as much detail as possible.

## Contributing

Contributions are always appreciated. For additional drivers, please ensure that the contributed database driver adheres to the Documentation, Coding Standards, and Build practices followed by existing drivers.

### Commit Messages

This project follows [Conventional Commits](https://www.conventionalcommits.org/) specifications. All commit messages should be structured as follows:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Types include:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `build`: Changes that affect the build system or external dependencies
- `ci`: Changes to our CI configuration files and scripts
- `chore`: Other changes that don't modify src or test files

Examples:
```
feat(api): add ability to parse arrays
fix(driver): resolve deadlock in Redis driver
docs: update installation instructions
```

Commits following this convention will automatically be included in the changelog when a release is created.
