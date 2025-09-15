# Release Guide

## Creating a New Release

### 1. Update Version
Update the version in `version.go` if needed (usually handled by CI/CD).

### 2. Create and Push Tag
```bash
# Create a new tag
git tag -a v1.0.0 -m "Release version 1.0.0"

# Push the tag to trigger the release
git push origin v1.0.0
```

### 3. Automatic Release Process
The GitHub Actions workflow will automatically:
- Build binaries for all supported platforms
- Create archives (zip for Windows, tar.gz for Unix)
- Generate release notes
- Create a GitHub release with all artifacts

## Supported Platforms

The release process builds binaries for:

### Windows
- `file-uploader-windows-amd64.exe` (64-bit)
- `file-uploader-windows-386.exe` (32-bit)

### macOS
- `file-uploader-darwin-amd64.tar.gz` (Intel)
- `file-uploader-darwin-arm64.tar.gz` (Apple Silicon)

### Linux
- `file-uploader-linux-amd64.tar.gz` (64-bit)
- `file-uploader-linux-386.tar.gz` (32-bit)
- `file-uploader-linux-arm64.tar.gz` (ARM 64-bit)
- `file-uploader-linux-arm.tar.gz` (ARM 32-bit)

## Manual Release (if needed)

If you need to create a release manually:

```bash
# Build for your platform
go build -ldflags="-s -w" -o file-uploader main.go

# Test the binary
./file-uploader --version
./file-uploader --port 8080
```

## Version Format

Use semantic versioning: `vMAJOR.MINOR.PATCH`

Examples:
- `v1.0.0` - First stable release
- `v1.0.1` - Bug fix
- `v1.1.0` - New features
- `v2.0.0` - Breaking changes

## Release Notes Template

The automatic release includes:
- Download links for all platforms
- Installation instructions
- Feature list
- System requirements
- Usage instructions

## Testing Before Release

1. Run tests: `go test -v ./...`
2. Build and test locally: `go build -o file-uploader main.go`
3. Test the web interface
4. Verify file upload functionality
5. Test on different browsers/devices
