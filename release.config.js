const repositoryUrl = process.env.GITHUB_REPOSITORY
  ? `https://github.com/${process.env.GITHUB_REPOSITORY}.git`
  : 'https://github.com/yourname/psstd.git';

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
          { path: 'psstd', label: 'psstd binary' }
        ]
      }
    ]
  ]
};
