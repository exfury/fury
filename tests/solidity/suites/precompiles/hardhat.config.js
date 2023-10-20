require('@nomicfoundation/hardhat-toolbox')

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: '0.8.18',
  networks: {
    fury: {
      url: 'http://127.0.0.1:8545',
      chainId: 710,
      accounts: [
        '0x2992593f994ce4456254d8f3b9238282a9cc7056a27f08ffe597607a683b133a',
        '0x2992593f994ce4456254d8f3b9238282a9cc7056a27f08ffe597607a683b133a'
      ]
    }
  }
}
