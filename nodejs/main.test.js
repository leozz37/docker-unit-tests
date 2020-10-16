const helloWorld = require('./main');

test('Hello World', () => {
  expect(helloWorld()).toBe("Hello from Nodejs!");
});