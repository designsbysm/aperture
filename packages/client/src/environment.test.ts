import environment from './environment';

describe('environment', () => {
  it('should be an object', () => {
    expect(typeof environment)
      .toBe('object');
  });

  it('should contain node', () => {
    expect(environment)
      .toHaveProperty('node');
  });

  it('should contain apiURL', () => {
    expect(environment)
      .toHaveProperty('apiURL');

    expect(environment.apiURL)
      .toBeDefined();
  });
});
