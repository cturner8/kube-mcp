# Wiki Development Documentation Gaps Analysis

**Date:** 2026-01-17  
**Scope:** Development section of kube-mcp wiki  
**Wiki URL:** https://github.com/cturner8/kube-mcp/wiki

## Executive Summary

The kube-mcp wiki has a solid foundation with 6 development-related pages, but several critical gaps exist that would benefit contributors. This document identifies missing documentation and provides recommendations for improvement.

## Current Development Section Structure

From the sidebar (`_Sidebar.md`):

```
## Development
- Development Environment
- Create a dev cluster
- Working with the MCP API
- Helm Development
- Debugging
- Release Process
```

## Detailed Gap Analysis

### 1. Testing Documentation ⚠️ **CRITICAL GAP**

**Current State:** No documentation exists  
**What's Missing:**
- How to run tests: `go test -v ./...` (from test.yml workflow)
- Where test files are located (currently no `*_test.go` files exist)
- Testing practices and conventions
- How to write tests for new tools
- Integration testing approach
- Mock Kubernetes client usage

**Impact:** High - Contributors don't know if/how to test their changes

**Recommendation:** Create "Testing.md" page with:
```markdown
## Running Tests

```bash
cd api
go test -v ./...
```

## Writing Tests

[Guide on test structure, mocking, etc.]

## CI Testing

Tests run automatically via GitHub Actions (see `.github/workflows/test.yml`)
```

---

### 2. Building and Docker ⚠️ **CRITICAL GAP**

**Current State:** Partial - `Development-Environment.md` mentions `go build .` but no Docker/Compose docs  
**What's Missing:**
- Building Docker images locally
- Using `compose.yml` for local testing
- Multi-stage Dockerfile explanation (api/Dockerfile)
- Docker image architecture (Go 1.25 → Alpine 3.22)
- Building for multiple platforms (linux/arm64, linux/amd64)

**Impact:** High - Contributors can't test Docker deployments locally

**Recommendation:** Create "Building-and-Docker.md" page with:
```markdown
## Building Locally

```bash
cd api
go build .
```

## Building Docker Image

```bash
docker build -t kube-mcp:local api/
```

## Using Docker Compose

```bash
docker compose up --build
```

The `compose.yml` includes pre-configured OIDC settings for local testing.

## Multi-Architecture Builds

The project supports `linux/amd64` and `linux/arm64`. See `.github/workflows/build.yml` for CI build process.
```

---

### 3. Code Architecture/Structure 📊 **IMPORTANT GAP**

**Current State:** Excellent info exists in `.github/copilot-instructions.md` but NOT in wiki  
**What's Missing:**
- Architecture diagram (exists in copilot-instructions.md)
- Directory structure explanation
- Tool pattern documentation (list/get pattern)
- How tools are registered in `api/server/server.go`
- Configuration flow through `api/config/config.go`

**Impact:** Medium - New contributors lack understanding of codebase organization

**Recommendation:** Create "Code-Architecture.md" page, adapting content from copilot-instructions.md:
```markdown
## Architecture Overview

[Include the architecture diagram]

## Directory Structure

```
api/
├── main.go                 # Entry point
├── config/                 # Configuration parsing
├── kubernetes/            # K8s client wrapper
├── server/                # HTTP server & MCP protocol
└── tools/                 # Kubernetes resource tools
```

## Tool Pattern

Every Kubernetes resource follows a consistent pattern:
1. List Tool (list{Resource}s.go)
2. Get Tool (get{Resource}.go)

[More details...]
```

---

### 4. CI/CD Workflows 🔧 **IMPORTANT GAP**

**Current State:** No documentation  
**What's Missing:**
- Explanation of test.yml, build.yml, helm.yml workflows
- When workflows trigger (tags, manual, etc.)
- Image signing with cosign
- Attestation generation
- Release creation process

**Impact:** Medium - Contributors don't understand automation

**Recommendation:** Create "CI-CD-Workflows.md" page:
```markdown
## GitHub Actions Workflows

### test.yml
Manual workflow for running tests:
- Builds the Go application
- Runs all tests

### build.yml
Triggered on version tags (v*):
- Builds multi-arch Docker images
- Signs images with cosign
- Generates SBOM attestations
- Creates draft GitHub release

### helm.yml
Triggered on chart tags (chart-v*):
- Packages Helm chart
- Pushes to OCI registry
- Creates draft GitHub release
```

---

### 5. Release Process 🚨 **CRITICAL GAP**

**Current State:** Empty file (`Release-Process.md` has 0 bytes)  
**What's Missing:**
- Complete release workflow
- Version tagging conventions (v* for app, chart-v* for Helm)
- How to create releases
- Pre-release checklist

**Impact:** High - Maintainers have no documented release process

**Recommendation:** Populate "Release-Process.md":
```markdown
## Versioning

kube-mcp uses semantic versioning:
- Application: `v1.2.3`
- Helm Chart: `chart-v1.2.3`

## Release Workflow

### Application Release

1. Update version in relevant files
2. Create and push tag:
   ```bash
   git tag v1.2.3
   git push origin v1.2.3
   ```
3. GitHub Actions builds, signs, and creates draft release
4. Edit release notes and publish

### Helm Chart Release

1. Update Chart.yaml version
2. Create and push tag:
   ```bash
   git tag chart-v1.2.3
   git push origin chart-v1.2.3
   ```
3. GitHub Actions packages and publishes chart
```

---

### 6. Contributing Guidelines 📝 **IMPORTANT GAP**

**Current State:** No contributing documentation  
**What's Missing:**
- How to contribute (fork, branch, PR)
- Code style and conventions
- PR review process
- Commit message format
- Issue reporting guidelines

**Impact:** Medium - New contributors lack guidance on contribution process

**Recommendation:** Create "Contributing.md" page:
```markdown
## How to Contribute

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Make changes and test
4. Submit a pull request

## Code Conventions

- Use `gofmt` for code formatting
- Follow Go best practices
- Tool names use snake_case (e.g., `get_pod`, `list_pods`)
- Add tests for new features
- Update wiki documentation for significant changes

## Pull Request Process

1. Ensure tests pass
2. Update documentation
3. Request review
4. Address feedback
```

---

### 7. Debugging (Expansion Needed) 🔍 **MODERATE GAP**

**Current State:** Only covers MCP Inspector (`Debugging.md` is 455 bytes)  
**What's Missing:**
- Go debugging with delve
- Logging levels and configuration (`KUBE_MCP_LOG_LEVEL`)
- Common issues and solutions
- How to debug tool handlers
- Kubernetes client debugging

**Impact:** Medium - Developers have limited debugging options documented

**Recommendation:** Expand "Debugging.md":
```markdown
## MCP Inspector

[Current content]

## Go Debugging

### Using Delve

```bash
cd api
dlv debug . -- --out-of-cluster
```

### Logging Levels

Set log level via environment variable:
```bash
export KUBE_MCP_LOG_LEVEL=debug  # debug|info|warn|error
```

## Common Issues

### "Cannot connect to Kubernetes"
- Verify kubeconfig: `kubectl cluster-info`
- Check `--out-of-cluster` flag
- Review RBAC permissions

### "JWT validation failed"
- Verify OIDC configuration
- Check token expiration
- Validate signing method matches provider
```

---

### 8. Adding New Tools Guide 🛠️ **IMPORTANT GAP**

**Current State:** Pattern documented in copilot-instructions.md but not in wiki  
**What's Missing:**
- Step-by-step guide for adding Kubernetes resource tools
- Template/example code
- Registration process
- Testing new tools

**Impact:** High - Contributors can't easily extend functionality

**Recommendation:** Create "Adding-New-Tools.md" page:
```markdown
## Tool Pattern

Every Kubernetes resource follows this pattern:

1. **List Tool** (`list{Resource}s.go`)
2. **Get Tool** (`get{Resource}.go`)

## Step-by-Step Guide

### 1. Create List Tool

File: `api/tools/listSecrets.go`

```go
package tools

import (
    "encoding/json"
    // imports...
)

var ListSecretsTool = mcp.NewTool(/* ... */)
```

### 2. Create Get Tool

File: `api/tools/getSecret.go`

[Similar pattern]

### 3. Register Tools

In `api/server/server.go`:

```go
if tools.IsToolAllowed("list_secrets") {
    mcp.AddTool(tools.ListSecretsTool)
}
if tools.IsToolAllowed("get_secret") {
    mcp.AddTool(tools.GetSecretTool)
}
```

### 4. Test

```bash
cd api
go build .
go run . --out-of-cluster
# Use MCP Inspector to test new tools
```
```

---

## Priority Recommendations

### High Priority (Do First)
1. ✅ **Complete Release-Process.md** - Currently empty, blocks releases
2. ✅ **Create Testing.md** - Critical for code quality
3. ✅ **Create Building-and-Docker.md** - Needed for local testing

### Medium Priority (Do Soon)
4. ✅ **Create Code-Architecture.md** - Helps new contributors
5. ✅ **Create Adding-New-Tools.md** - Enables feature contributions
6. ✅ **Expand Debugging.md** - Improves developer experience
7. ✅ **Create Contributing.md** - Standardizes contribution process

### Lower Priority (Nice to Have)
8. ✅ **Create CI-CD-Workflows.md** - Documents automation

## Sidebar Updates Needed

After creating new pages, update `_Sidebar.md`:

```markdown
## Development

- [Development Environment](...)
- [Create a dev cluster](...)
- [Working with the MCP API](...)
- [Code Architecture](...)                    # NEW
- [Adding New Tools](...)                     # NEW
- [Building and Docker](...)                  # NEW
- [Testing](...)                              # NEW
- [Debugging](...)
- [Helm Development](...)
- [Contributing](...)                         # NEW
- [CI/CD Workflows](...)                      # NEW
- [Release Process](...)
```

## Cross-Reference Opportunities

Several pages should reference each other:
- Development-Environment.md → link to Testing.md, Building-and-Docker.md
- Working-with-the-MCP-API.md → link to Debugging.md
- Adding-New-Tools.md → link to Code-Architecture.md, Testing.md
- Contributing.md → link to all development pages

## Metrics

| Category | Current Pages | Needed Pages | Completion |
|----------|---------------|--------------|------------|
| Setup | 2 | 2 | 100% ✅ |
| Development | 4 | 12 | 33% ⚠️ |
| **Total** | **6** | **14** | **43%** |

## Conclusion

The kube-mcp wiki development section has a solid foundation but needs expansion in 8 key areas. Addressing the high-priority gaps (Release Process, Testing, Building) should be done first, followed by architecture and tooling guides. This will provide a comprehensive developer experience and enable easier contributions to the project.

## Next Steps

1. Review this analysis with project maintainers
2. Prioritize which gaps to address first
3. Create new wiki pages for identified gaps
4. Update existing pages with cross-references
5. Update sidebar navigation
6. Announce documentation improvements to contributors
