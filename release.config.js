const repositoryUrl = process.env.GITHUB_REPOSITORY
  ? `https://github.com/${process.env.GITHUB_REPOSITORY}.git`
  : 'https://github.com/OffPeakEngineer/pulsed.git';

module.exports = {
  branches: ['main'],
  repositoryUrl,
  plugins: [
    '@semantic-release/commit-analyzer',
    '@semantic-release/release-notes-generator',
    [
      '@semantic-release/github',
      {
        assets: [
          { path: 'dist/pulsed-linux-amd64', label: 'Linux amd64 binary' },
          { path: 'dist/pulsed-linux-arm64', label: 'Linux arm64 binary' },
          { path: 'dist/pulsed-darwin-amd64', label: 'macOS Intel binary' },
          { path: 'dist/pulsed-darwin-arm64', label: 'macOS Apple Silicon binary' },
          { path: 'dist/pulsed-windows-amd64.exe', label: 'Windows amd64 binary' },
          { path: 'dist/pulsed-windows-arm64.exe', label: 'Windows arm64 binary' }
        ]
      }
    ]
  ]
};
