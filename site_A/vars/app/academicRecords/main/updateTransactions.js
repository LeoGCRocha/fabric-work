'use strict'

/**@module EvaluateTransactions */
const academicRecords = require('../academicRecords')
const CustomError = require('./utils/CustomError')

/**
 * Process transactions of type read
 * @param {*} academicRecordsContract - Academic Records contract interface
 * @param {json} option - Define the type and parammeters about this executation
 */
async function updateTransaction(academicRecordsContract, option) {
    const errContext = "Não foi possível realizar a transação"
    let response
    let jsonReffer

    // Type check.
    switch (option['type']) {
    case 'approveAtividade':
        response = await academicRecords.approveAtividade(academicRecordsContract, option['CPF'], option['id'])
        return academicRecords.bufferToStringResponse(response)
    case 'approveEstagio':
        response = await academicRecords.approveEstagio(academicRecordsContract, option['CPF'], option['id'])
        return academicRecords.bufferToStringResponse(response)
    default:
        throw new CustomError('invalidFunctionNameError', errContext)        
    }
}

module.exports = { updateTransaction }