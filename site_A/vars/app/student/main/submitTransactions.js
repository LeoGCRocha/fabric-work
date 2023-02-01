'use strict'

const student = require('../student.js')
const CustomError = require('./utils/CustomError')

async function submitTransaction(studentContract, reqParams) {
    const errContext = "Nao foi possivel subemeter a transacao."
    let response
    switch(reqParams.type) {
        case 'addAtividade':
            if (reqParams.CPF === undefined || reqParams.atividade === undefined) {
                throw new CustomError('invalidTransactionTypeError', errContext)
            }
            response = await student.addAtividade(studentContract, reqParams.CPF,JSON.stringify(reqParams.atividade))
            return student.bufferToStringResponse(response)
        case 'addEstagio':
            if (reqParams.CPF === undefined || reqParams.estagio === undefined) {
                throw new CustomError('invalidTransactionTypeError', errContext)
            }
            response = await student.addEstagio(studentContract, reqParams.CPF, JSON.stringify(reqParams.estagio))
            return student.bufferToStringResponse(response)
        case 'registerStudent':
            response = await student.registerStudent(studentContract, 
                                            reqParams.nomeDeUsuario, 
                                            reqParams.CPF, 
                                            reqParams.email, 
                                            reqParams.senha)
            return student.bufferToStringResponse(response)
        default:
            throw new CustomError('invalidTransactionTypeError', errContext)
    }
}

module.exports = { submitTransaction }