const Lottery = artifacts.require('Lottery');

contract('Lottery', (accounts) => {
  let lottery;

  beforeEach(async () => {
    lottery = await Lottery.new();
  });

  it.only('getTotalPot', async () => {
    const totalPot = await lottery.getTotalPot();

    assert.equal(totalPot, 0);
  });
});
