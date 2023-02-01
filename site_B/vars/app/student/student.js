'use strict'

async function registerStudent(contract, nomeDeUsuario, CPF, email, senha) {
    return contract.submitTransaction('RegisterStudent', nomeDeUsuario, CPF, email, senha)
}

async function verifyStudentLogin(contract, nomeDeUsuario, senha) {
    return contract.evaluateTransaction('VerifyStudentLogin', nomeDeUsuario, senha)
}

async function allStudents(contract) {
    return contract.evaluateTransaction('AllStudents')
}

async function getHistoricos(contract, CPF) {
    return contract.evaluateTransaction('GetHistoricos', CPF)
}

async function getAtividades(contract, CPF) {
    return contract.evaluateTransaction('GetAtividades', CPF)
}

async function getEstagios(contract, CPF) {
    return contract.evaluateTransaction('GetEstagios', CPF)
}

async function addAtividade(contract, CPF, atividade) {
    return contract.submitTransaction('AddAtividade', CPF, atividade)
}

async function addEstagio(contract, CPF, estagio) {
    return contract.submitTransaction('AddEstagio', CPF, estagio)
}

/** @function
 * @moduleOf module:UtilsFunctions
 * @name bufferToStringResponse
 * @param {string} bufferResponse - buffer response from the chaincode
 * @returns {String} Buffer response converted to string
 */
function bufferToStringResponse(bufferResponse) {
    return Buffer.from(bufferResponse, 'utf-8').toString()
}

/**
 * @name jsonPrettyOutput 
 * @moduleOf module:UtilsFunctions
 * @param {json} jsonResponse - json string
 * @returns {String} - parse a json to a formatted json as string
 */
function jsonPrettyOutput(jsonResponse) {
    return JSON.stringify(JSON.parse(jsonResponse), undefined, 4)
}


module.exports = {
    registerStudent,
    verifyStudentLogin,
    allStudents,
    getHistoricos,
    getAtividades,
    getEstagios,
    addAtividade,
    addEstagio,
    bufferToStringResponse,
    jsonPrettyOutput
}