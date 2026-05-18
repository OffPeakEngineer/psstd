module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    'body-max-line-length': [0],
    'type-enum': [2, 'always', [
      'build',
      'chore',
      'ci',
      'docs',
      'wiki', 
      'feat',
      'fix',
      'perf',
      'refactor',
      'revert',
      'style',
      'test',
      'vibe'	    
      'plan'
      'vision',
    ]],
  },
};
