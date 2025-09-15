# GitHub Actions Workflows

This repository includes several GitHub Actions workflows for automated CI/CD:

## Workflows

### 1. CI (`.github/workflows/ci.yml`)
**Triggers:** Push to main/develop branches, Pull Requests

**Features:**
- Runs tests and linting
- Builds on multiple platforms (Ubuntu, Windows, macOS)
- Security scanning with Gosec
- Code coverage reporting
- Cache optimization for faster builds

### 2. Release (`.github/workflows/release.yml`)
**Triggers:** Push tags, Manual dispatch

**Features:**
- Builds for all supported platforms
- Creates release archives
- Generates changelog
- Creates GitHub releases with artifacts

### 3. Auto Release (`.github/workflows/auto-release.yml`)
**Triggers:** Push tags matching `v*.*.*`

**Features:**
- Automatic release creation
- Multi-platform binary builds
- Comprehensive release notes
- Archive creation (zip/tar.gz)

## Supported Platforms

The workflows build binaries for:

- **Windows**: amd64, 386
- **macOS**: amd64 (Intel), arm64 (Apple Silicon)
- **Linux**: amd64, 386, arm64, arm

## Usage

### Creating a Release

1. **Update version** (if needed):
   ```bash
   # Edit version.go if needed
   ```

2. **Create and push tag**:
   ```bash
   git tag -a v1.0.0 -m "Release version 1.0.0"
   git push origin v1.0.0
   ```

3. **Automatic release**: GitHub Actions will automatically:
   - Build all platform binaries
   - Create archives
   - Generate release notes
   - Create GitHub release

### Manual Release

You can also trigger releases manually:
1. Go to Actions tab
2. Select "Build and Release" workflow
3. Click "Run workflow"
4. Choose branch and run

## Configuration

### Required Secrets
- `GITHUB_TOKEN` (automatically provided)

### Environment Variables
- `GO_VERSION`: Go version to use (default: 1.21)

## Build Flags

The build process uses these flags:
- `-ldflags="-s -w"`: Strip debug info and symbol table
- `-X main.version=${{ github.ref_name }}"`: Inject version info
- `CGO_ENABLED=0`: Disable CGO for static binaries

## Artifacts

Each release includes:
- Platform-specific binaries
- Compressed archives (zip for Windows, tar.gz for Unix)
- Release notes with installation instructions
- Checksums for verification

## Monitoring

Monitor workflow runs in the Actions tab:
- Green checkmark: Success
- Red X: Failure
- Yellow circle: In progress

## Troubleshooting

### Common Issues

1. **Build failures**: Check Go version compatibility
2. **Upload failures**: Verify GitHub token permissions
3. **Test failures**: Review test output in logs

### Debugging

1. Check workflow logs in Actions tab
2. Verify all required files are present
3. Ensure proper file permissions
4. Check for syntax errors in workflow files
