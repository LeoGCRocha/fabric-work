'use strict'

/**@module EvaluateTransactions */

const student = require('../student')
const CustomError = require('./utils/CustomError')

async function evaluateTransaction(studentContract, option) {
    const errContext = "Não foi possível realizar a transação"

    let response
    switch (option['type']) {
        case 'getHistoricos':
            response = await student.getHistoricos(studentContract, option['CPF'])
            return student.jsonPrettyOutput(student.bufferToStringResponse(response))
        case 'getAtividades':
            response = await student.getAtividades(studentContract, option['CPF'])
            return student.jsonPrettyOutput(student.bufferToStringResponse(response))
        case 'getEstagios':
            response = await student.getEstagios(studentContract, option['CPF'])
            return student.jsonPrettyOutput(student.bufferToStringResponse(response))
        case 'validStudentLogin':
            response = await student.verifyStudentLogin(studentContract, option['nomeDeUsuario'], option['senha'])
            return student.jsonPrettyOutput(student.bufferToStringResponse(response))
        default:
            throw new CustomError('invalidTransactionTypeError', errContext)
    }
}

module.exports = { evaluateTransaction }