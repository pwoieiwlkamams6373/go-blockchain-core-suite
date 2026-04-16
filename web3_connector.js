const Web3 = require('web3');

class Web3Connector {
    constructor(rpcUrl) {
        this.web3 = new Web3(rpcUrl);
    }

    async getBlockNumber() {
        return await this.web3.eth.getBlockNumber();
    }

    async getBalance(address) {
        return await this.web3.eth.getBalance(address);
    }

    async sendTransaction(tx) {
        return await this.web3.eth.sendTransaction(tx);
    }
}

module.exports = Web3Connector;
