## 📄 `CONTRIBUTING.md`

````markdown
# Contributing to Example Monorepo

Welcome! We’re excited that you want to contribute. To make collaboration smooth, please follow these guidelines.

---

## 1. Code of Conduct
All contributors are expected to follow our [Code of Conduct](./CODE_OF_CONDUCT.md). Harassment or unprofessional behavior will not be tolerated.

---

## 2. Getting Started

### Prerequisites
- **Bazel**: pinned via `.bazelversion` (`8.3.1`)  
- **Bazelisk**: installed to respect `.bazelversion` and `.bazeliskrc`  
- **Node**: `nvm use` will pick `.nvmrc` (`20.11.1`)  
- **Python**: use pyenv/asdf to select `.python-version` (`3.12`)  
- **Java**: JDK 17, per `.java-version`  
- Recommended: VS Code with extensions installed in `.devcontainer/devcontainer.json`

### Devcontainer
For an isolated environment:
```bash
# Open in VS Code and choose "Reopen in Container"
````

---

## 3. Repository Layout

```text
/.
├── apps/
├── BUILD.bazel
├── CHANGELOG.md
├── CODE_OF_CONDUCT.md
├── config/
├── CONTRIBUTING.md
├── docs/
├── example.BUILD.bazel.go
├── example.BUILD.bazel.java
├── example.BUILD.bazel.py
├── example.BUILD.bazel.ts
├── examples/
├── experimental/
├── infra/
├── internal/
├── libs/
├── LICENSE
├── MODULE.bazel
├── MODULE.bazel.lock
├── package.json
├── packages/
├── platform/
├── README.md
├── requirements_lock.txt
├── requirements.txt
├── scripts/
├── SECURITY.md
├── services/
├── tests/
├── third_party/
├── tools/
└── WORKSPACE.bzlmod
```

---

## 4. Development Workflow

### Building

```bash
bazel build //...
```

### Running Tests

```bash
bazel test //:all_tests
```

### Formatting / Linting

```bash
# Go
bazel run //:lint_go

# Python
bazel run //:lint_python

# Node
bazel run //:lint_node

# Java
bazel run //:lint_java

# Bazel BUILD files
bazel run //:gazelle
```

---

## 5. Pre-Commit Hooks

We use [pre-commit](https://pre-commit.com/) to enforce linting and formatting:

```bash
pre-commit install
pre-commit run --all-files
```

---

## 6. Pull Request Guidelines

* Branch from `main` and create a descriptive name.
* Include tests for new features or bug fixes.
* Run all linters and formatters before opening a PR.
* Follow [Conventional Commits](https://www.conventionalcommits.org/) for commit messages.
* Link the PR to any relevant issue or changelog entry.

---

## 7. Reporting Issues

* Open GitHub issues for bugs, feature requests, or questions.
* Include enough context to reproduce the issue (logs, versions, steps).

---

## 8. Security

Report security vulnerabilities privately via [SECURITY.md](./SECURITY.md).

---

## 9. License

By contributing, you agree that your contributions will be licensed under the [MIT License](./LICENSE).