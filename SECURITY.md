# Security Policy

## Supported Versions
We actively maintain and monitor the following versions of our software in this repository:

- Bazel: 8.3.1 (via Bazelisk)
- Go: 1.23.2
- Python: 3.12
- Node: 20.11.1
- Java: 17

## Reporting a Vulnerability
If you discover a security vulnerability, please report it **privately** by emailing:

security@appliedinnovationcorp.com


Please do **not** create a public issue for security vulnerabilities. Include:

- A detailed description of the issue
- Steps to reproduce
- Impact assessment (if known)
- Your contact information

We will respond within **3 business days** and work with you to resolve the issue responsibly.

## Disclosure Policy
Once a vulnerability has been fixed, we may publicly disclose the issue with credit to the reporter, unless the reporter requests anonymity.  

## Security Best Practices
- Keep dependencies up to date (use `bazel fetch //...` regularly)
- Run all tests and linters before committing
- Use the pinned language/tool versions specified in `.bazelversion`, `.nvmrc`, `.python-version`, and `.java-version`
