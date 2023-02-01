'use strict'

/*
#######################
### Aluno Functions ###
#######################
*/

/** @module AlunoFunctions */

/** @function
 *  @moduleOf AlunoFunctions
 *  @name alunoExists
 *  @param {Object} contract - hyperledger fabric contract
 *  @param {String} alunoCPF - cpf do aluno
 *  @returns {Object} - response from the chaincode 
 */
async function alunoExists(contract, alunoCPF) {
    return contract.evaluateTransaction('AlunoExists', alunoCPF)
}

/** @function
 * @moduleOf AlunoFunctions
 * @name createAluno
 * @param {Object} contract - hyperledger fabric contract
 * @param {json} alunoJSON - aluno json
 */
async function createAluno(contract, alunoJSON) {
    return contract.submitTransaction('CreateAluno', alunoJSON)
}

/** @function
 * @moduleOf AlunoFunctions
 * @name readAluno
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} alunoCPF - cpf do aluno
 * @returns {Object} response from the chaincode
 */
async function readAluno(contract, alunoCPF) {
    return contract.evaluateTransaction('ReadAluno', alunoCPF)
}

/** @function
* @name getAllAlunos
* @moduleOf AlunoFunctions
* @param {Object} contract - hyperledger fabric contract
* @returns {Object} response from the chaincode
*/
async function getAllAlunos(contract) {
    return contract.evaluateTransaction('GetAllAlunos')
}

/** @function
 * @moduleOf AlunoFunctions
 * @name updateAluno
 * @param {Object} contract - hyperledger fabric contract 
 * @param {Object} alunoJSON - JSON Object representing the aluno
 * @returns {Object} response from the chaincode
 */
async function updateAluno(contract, alunoJSON) {
    return contract.submitTransaction('UpdateAluno', alunoJSON)
}

/*
###########################
### Historico Functions ###
###########################
*/

/** @module HistoricoFunctions */

/** @function
 * @moduleOf module:HistoricoFunctions
 * @name createHistoricoEscolar
 * @param {Object} contract - hyperledger fabric contract 
 * @param {json} historicoJSON - JSON Object representing the historico
 * @returns {Object} response from the chaincode
 */
async function createHistoricoEscolar(contract, historicoJSON) {
    return contract.submitTransaction('CreateHistoricoEscolar', historicoJSON)
}

/** @function
 * @moduleOf module:HistoricoFunctions 
 * @name readHistorico 
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} historico - string representing the historico identificator 
 * @returns {Object} response from the chaincode
 */
async function readHistorico(contract, historico) {
    return contract.evaluateTransaction('ReadHistorico', historico)
}

/** @function
 * @moduleOf module:HistoricoFunctions
 * @name readHistoricosFromAluno
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} cpfAluno - string representing the aluno identificator
 * @returns {Object} response from the chaincode
 */
async function readHistoricosFromAluno(contract, cpfAluno) {
    return contract.evaluateTransaction('ReadHistoricosFromAluno', cpfAluno)
}

/** @function
 * @moduleOf module:HistoricoFunctions 
 * @name readHistoricoFromAluno
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} cpfAluno - string representing the aluno identificator
 * @returns {Object} response from the chaincode
 */
async function readLastHistoricoFromAluno(contract, cpfAluno) {
    return contract.evaluateTransaction('ReadLastHistoricoFromAluno', cpfAluno)
}

/** @function
 * @moduleOf module:HistoricoFunctions 
 * @name readHistoricosFromCurriculo
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} curriculo - string representing the curriculo identificator
 * @returns {Object} response from the chaincode
 */
async function readHistoricoFromCurriculo(contract, curriculo) {
    return contract.evaluateTransaction('ReadHistoricoFromCurriculo', curriculo)
}

/** @function
 * @moduleOf module:HistoricoFunctions 
 * @name readHistoricoFromIES
 * @param {Object} contract - hyperledger fabric contract 
 * @param {String} codigoIES - string representing the ies identificator
 * @returns {Object} response from the chaincode
 */
async function readHistoricoFromIes(contract, codigoIES) {
    return contract.evaluateTransaction('ReadHistoricoFromIES', codigoIES)
}

/** @function
 * @moduleOf module:HistoricoFunctions
 * @name readHistoricoFromCurso
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} codigoCurso - string representing the curso identificator
 * @returns {Object} response from the chaincode
 */
async function readHistoricoFromCurso(contract, codigoCurso) {
    return contract.evaluateTransaction('ReadHistoricoFromCurso', codigoCurso)
}

/*
###########################
### Student Functions   ###
###########################
*/

/** @module StudentFunctions */

/** @function
 * @moduleOf module:StudentFunctions
 * @name getCategoriasAtividades
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} CPF - CPF do aluno
 * @returns {Object} response from the chaincode
 **/
async function getCategoriasAtividades(contract, CPF) {
    return contract.evaluateTransaction('GetCategoriasAtividades', CPF)
}

/** @function
 * @moduleOf module:StudentFunctions
 * @name getCategoriasAtividades
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} CPF - CPF
 * @returns {Object} response from the chaincode
 **/
async function readAtividadesComplementaresAluno(contract, CPF) {
    return contract.evaluateTransaction('ReadAtividadesComplementaresFromAluno', CPF)
}

/** @function
 * @moduleOf module:StudentFunctions
 * @name getCategoriasAtividades
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} CPF - CPF
 * @returns {Object} response from the chaincode
 **/
async function readEstagiosFromAluno(contract, CPF) {
    return contract.evaluateTransaction('ReadEstagiosFromAluno', CPF)
}

/** @function
 * @moduleOf module:StudentFunctions
 * @name getCategoriasAtividades
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} CPF - CPF
 * @param {Strign} UUID Identificator of the activity
 * @returns {Object} response from the chaincode
 **/
async function approveAtividade(contract, CPF, UUID) {
    return contract.submitTransaction('ApproveAtividade', CPF, UUID)
}

/** @function
 * @moduleOf module:StudentFunctions
 * @name getCategoriasAtividades
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} CPF - CPF do aluno
 * @param {Strign} UUID Identificator of the internship
 * @returns {Object} response from the chaincode
 **/
async function approveEstagio(contract, CPF, UUID) {
    return contract.submitTransaction('ApproveEstagio', CPF, UUID)
}

/*
#######################
### Utils Functions ###
#######################
*/

/** @module UtilsFunctions */

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
    // Aluno Functions
    alunoExists,
    createAluno,
    readAluno,
    getAllAlunos,
    updateAluno,
    // Historico Functions
    createHistoricoEscolar,
    readHistorico,
    readHistoricoFromIes,
    readHistoricoFromCurso,
    readHistoricosFromAluno,
    readLastHistoricoFromAluno,
    readHistoricoFromCurriculo,
    // Student Functions
    getCategoriasAtividades,
    readAtividadesComplementaresAluno,
    readEstagiosFromAluno,
    approveAtividade,
    approveEstagio,
    // Utils Functions
    jsonPrettyOutput,
    bufferToStringResponse
}
