## 📄 `README.md`

````markdown
# Example Monorepo

Welcome to the Example Monorepo — a **polyglot, Bazel-managed** repository containing Go, Python, Node/TypeScript, and Java code. This repo uses **Bazel 8.3.1** with **bzlmod** for dependency management, pre-configured devcontainer, and unified tooling across languages.

---

## Table of Contents
- [Repository Structure](#repository-structure)
- [Getting Started](#getting-started)
- [Building](#building)
- [Testing](#testing)
- [Linting & Formatting](#linting--formatting)
- [Contributing](#contributing)
- [Changelog](#changelog)
- [License](#license)

---

## Repository Structure

```text
.
├── apps
├── BUILD.bazel
├── CHANGELOG.md
├── CODE_OF_CONDUCT.md
├── config
├── CONTRIBUTING.md
├── docs
├── example.BUILD.bazel.go
├── example.BUILD.bazel.java
├── example.BUILD.bazel.py
├── example.BUILD.bazel.ts
├── examples
├── experimental
├── infra
├── internal
├── libs
├── LICENSE
├── MODULE.bazel
├── MODULE.bazel.lock
├── package.json
├── packages
├── platform
├── README.md
├── requirements_lock.txt
├── requirements.txt
├── scripts
├── SECURITY.md
├── services
├── tests
├── third_party
├── tools
└── WORKSPACE.bzlmod
````

---

## Getting Started

### Prerequisites

* Install [Bazelisk](https://github.com/bazelbuild/bazelisk) (respects `.bazelversion`)
* Node.js via [nvm](https://github.com/nvm-sh/nvm) (`nvm use`)
* Python via [pyenv](https://github.com/pyenv/pyenv) or asdf (`.python-version`)
* Java JDK 17 (`.java-version`)
* Recommended: VS Code + [devcontainer](.devcontainer/devcontainer.json)

### Devcontainer

```bash
# Open in VS Code
# Reopen in container for fully preconfigured environment
```

---

## Building

```bash
# Build all targets
bazel build //...
```

---

## Testing

```bash
# Run all tests
bazel test //:all_tests
```

---

## Linting & Formatting

```bash
# Run language-specific linting
bazel run //:lint_go
bazel run //:lint_python
bazel run //:lint_node
bazel run //:lint_java

# Format Bazel BUILD files
bazel run //:gazelle

# Pre-commit hooks
pre-commit install
pre-commit run --all-files
```

---

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines on:

* Branching & commit messages
* PRs & code reviews
* Running tests & linters
* Reporting issues

---

## Changelog

All notable changes are documented in [CHANGELOG.md](./CHANGELOG.md).

---

## License

This project is licensed under the [MIT License](./LICENSE).

---

## Security

See [SECURITY.md](./SECURITY.md) for reporting vulnerabilities.