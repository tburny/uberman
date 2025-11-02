Verify that the latest GitHub release includes all binary artifacts.

Check that the following files are present:
- uberman-linux-amd64.tar.gz
- uberman-linux-arm64.tar.gz
- uberman-darwin-amd64.tar.gz
- uberman-darwin-arm64.tar.gz
- uberman-windows-amd64.exe.zip
- uberman-freebsd-amd64.tar.gz
- All corresponding .sha256 checksum files

If binaries are missing, investigate the GitHub Actions workflow logs.
