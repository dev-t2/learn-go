const Lottery = artifacts.require('Lottery');

contract('Lottery', (accounts) => {
  let lottery;

  beforeEach(async () => {
    console.log('  Executed Callback Function: beforeEach');

    lottery = await Lottery.new();
  });

  it('Lottery Test', async () => {
    console.log('  Executed Callback Function: it');

    const owner = await lottery.owner();
    const value = await lottery.getSomeValue();

    console.log({ owner, value });

    assert.equal(value, 5);
  });
});
