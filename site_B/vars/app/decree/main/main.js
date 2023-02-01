'use strict';

const { Gateway, Wallets } = require('fabric-network')
const fs = require('fs')
const path = require('path')
const submitT = require('./submitTransactions.js')
const evalT = require('./evaluateTransactions.js')
const CustomError = require('./utils/CustomError')

async function start(reqParams) {
    const errContext = 'Nao foi possivel executar a funcao'
    // Load the network configuration.
    const ccpPath = path.resolve(__dirname, '.', '../../connection.json');
    const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

    // Create a new file system based wallet for managing identities.
    const walletPath = path.join('/vars/profiles/vscode/wallets', process.env.ORG_NAME);
    const wallet = await Wallets.newFileSystemWallet(walletPath);

    // Check to see if we've already enrolled the admin user.
    const identity = await wallet.get('Admin')
    if (!identity) { throw new CustomError('walletAdminNotFound', errContext) }
    const gateway = new Gateway()
    await gateway.connect(ccp, { wallet, identity: 'Admin', discovery: { enabled: true, asLocalhost: false } });
    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork('jornada')

    // Get the contract from the network.
    const contractDecree = network.getContract('decree')

    if (!reqParams.hasOwnProperty('op')) { // eslint-disable-line
        throw new CustomError('invalidOPParameter', errContext)
    }
    // Handle transaction types.
    let result
    switch (reqParams.op) {
    case 'write':
        // Submit the specified transaction.
        await submitT.submitTransactions(contractDecree, reqParams.path)
        break
    case 'read':
        // Evaluate the specified transaction.
        result = await evalT.evaluateTransactions(contractDecree, reqParams)
        break
    default:
        throw new CustomError('invalidFunctionNameError', errContext)
    }
    // Disconnect from the gateway.
    gateway.disconnect()
    return result
}

module.exports = { start }
