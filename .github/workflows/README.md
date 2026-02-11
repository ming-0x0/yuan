# GitHub Actions Workflows

This directory contains GitHub Actions workflows for CI/CD automation.

## Available Workflows

### 1. CI (`ci.yaml`)
**Trigger**: Push/PR to `main` or `develop` branches

**Jobs**:
- **Lint**: Run golangci-lint for code quality
- **Test**: Run unit tests with MySQL service, generate coverage
- **Build**: Build application and upload binary artifact

**Features**:
- MySQL 9 service container for integration tests
- Code coverage upload to Codecov
- Parallel job execution with dependencies

---

### 2. Deploy to Dev (`deploy-dev.yaml`)
**Trigger**: Push to `develop` branch or manual workflow dispatch

**Jobs**:
- Build and push Docker image to Docker Hub
- Deploy to dev server via SSH
- Run health checks
- Send Slack notifications

**Required Secrets**:
- `DOCKER_USERNAME`: Docker Hub username
- `DOCKER_PASSWORD`: Docker Hub password
- `DEV_SERVER_HOST`: Dev server hostname/IP
- `DEV_SERVER_USER`: SSH username
- `DEV_SERVER_SSH_KEY`: SSH private key
- `SLACK_WEBHOOK`: Slack webhook URL (optional)

---

### 3. Docker Build (`docker-build.yaml`)
**Trigger**: Push to `main`/`develop`, tags `v*`, or PR to `main`

**Jobs**:
- Build Docker image with multi-tag support
- Run Trivy security scanner
- Upload security results to GitHub Security

**Features**:
- Smart tagging based on branch/tag/PR
- Docker layer caching for faster builds
- Vulnerability scanning

---

### 4. PR Checks (`pr-checks.yaml`)
**Trigger**: Pull request opened/updated

**Jobs**:
- **Validate**: Check PR title format (conventional commits)
- **Code Quality**: Check formatting, run go vet, verify go.mod
- **Dependency Review**: Scan for vulnerable dependencies

**Features**:
- Enforces semantic PR titles
- Detects merge conflicts
- Checks code formatting

---

### 5. Release (`release.yaml`)
**Trigger**: Push tags matching `v*.*.*`

**Jobs**:
- Build binaries for multiple platforms (Linux, macOS, Windows)
- Generate changelog from commits
- Create GitHub release with binaries
- Build and push Docker image with version tags

**Platforms**:
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

---

## Setup Instructions

### 1. Configure Secrets

Go to **Settings → Secrets and variables → Actions** and add:

```
DOCKER_USERNAME=your-dockerhub-username
DOCKER_PASSWORD=your-dockerhub-token
DEV_SERVER_HOST=dev.example.com
DEV_SERVER_USER=deploy
DEV_SERVER_SSH_KEY=<your-ssh-private-key>
SLACK_WEBHOOK=https://hooks.slack.com/services/xxx
```

### 2. Enable GitHub Actions

Ensure Actions are enabled in **Settings → Actions → General**

### 3. Configure Environments

Create a `dev` environment in **Settings → Environments** for deployment protection rules.

---

## Workflow Triggers

| Workflow | Push `main` | Push `develop` | PR | Tag `v*` | Manual |
|----------|-------------|----------------|-----|----------|--------|
| CI | ✅ | ✅ | ✅ | ❌ | ❌ |
| Deploy Dev | ❌ | ✅ | ❌ | ❌ | ✅ |
| Docker Build | ✅ | ✅ | ✅ | ✅ | ❌ |
| PR Checks | ❌ | ❌ | ✅ | ❌ | ❌ |
| Release | ❌ | ❌ | ❌ | ✅ | ❌ |

---

## Branching Strategy

- **`main`**: Production-ready code
- **`develop`**: Development branch, auto-deploys to dev environment
- **Feature branches**: Create PR to `develop`
- **Tags**: `v1.2.3` format triggers release workflow

---

## PR Title Format

Use conventional commit format:

```
feat: add new feature
fix: resolve bug
docs: update documentation
style: format code
refactor: restructure code
perf: improve performance
test: add tests
build: update build config
ci: update CI/CD
chore: maintenance tasks
```

---

## Example Usage

### Deploy to Dev
```bash
# Push to develop branch
git checkout develop
git push origin develop

# Or trigger manually from GitHub UI
# Actions → Deploy to Dev → Run workflow
```

### Create Release
```bash
# Create and push tag
git tag v1.0.0
git push origin v1.0.0

# Release workflow will:
# - Build binaries for all platforms
# - Generate changelog
# - Create GitHub release
# - Push Docker image
```

### Run Tests Locally
```bash
# Same as CI workflow
make test
make lint
make build
```

---

## Troubleshooting

### Deployment Fails
1. Check secrets are configured correctly
2. Verify SSH key has access to dev server
3. Ensure dev server has Docker installed
4. Check server logs: `docker compose logs`

### Docker Build Fails
1. Verify Dockerfile syntax
2. Check Docker Hub credentials
3. Review build logs in Actions tab

### Tests Fail
1. Ensure MySQL service is healthy
2. Check test database configuration
3. Run tests locally: `make test`

---

## Maintenance

- Update Go version in all workflows when upgrading
- Rotate secrets regularly
- Review and update dependencies
- Monitor workflow execution times
