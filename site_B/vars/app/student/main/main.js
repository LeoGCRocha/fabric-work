'use strict'

const { Gateway, Wallets } = require('fabric-network')
const fs = require('fs')
const path = require('path')
const evaluateT = require("./evaluateTransactions")
const submitT = require("./submitTransactions")
const CustomError = require('./utils/CustomError')

/**
 * Runs appacademic applicaton
 * @param {*} reqParams - request parameters
 */
async function start(reqParams) {

    const errContext = 'Não foi possível executar a função.'
    // Load the network configuration.
    const ccpPath = path.resolve(__dirname, '..', '..', 'connection.json');
    const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));
    // Create a new file system based wallet for managing identities.
    const walletPath = path.join('/vars/profiles/vscode/wallets', process.env.ORG_NAME);
    const wallet = await Wallets.newFileSystemWallet(walletPath);        
    // Check to see if we've already enrolled the admin user.
    const identity = await wallet.get('Admin');
    if (!identity) { throw new CustomError('walletAdminNotFound', errContext) }
    // Network definitions.
    const gateway = new Gateway();
    await gateway.connect(ccp, { wallet, identity: 'Admin', discovery: { enabled: true, asLocalhost: false } });
    const network = await gateway.getNetwork('jornada');
    const studentContract = network.getContract('student')
    // Option
    let result
    if (!reqParams.hasOwnProperty('op')) { // eslint-disable-line
        throw new CustomError('invalidOPParameter', errContext)
    }    
    switch (reqParams['op']) {
    case "write":
        result = await submitT.submitTransaction(studentContract, reqParams)
        break
    case "read":
        result = await evaluateT.evaluateTransaction(studentContract, reqParams)
        break
    default:
        throw new CustomError('invalidFunctionNameError', errContext)
    }
    // Disconnect from the gateway.
    gateway.disconnect();
    return result
}

module.exports = { start }