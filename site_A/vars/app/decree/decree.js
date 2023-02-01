'use strict'

/*
#######################
### Curso Functions ###
#######################
*/

/** @module CursoFunctions */

/** @function
 * @moduleOf module:CursoFunctions
 * @name addCurso 
 * @param {Object} contract - hyperledger fabric contract
 * @param {json} cursoJSON - JSON object representing the curso
 * @returns {Object} response from the chaincode
 */
async function addCurso(contract, cursoJSON) {
    return await contract.submitTransaction('AddCurso', cursoJSON)
}

/** @function
 * @moduleOf module:CursoFunctions
 * @name cursoExists
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} cursoID - string representing the curso identifier
 * @returns {Object} response from the chaincode
 */
async function cursoExists(contract, cursoID) {
    return await contract.evaluateTransaction('CursoExists', cursoID)
    
}

/** @function
 * @moduleOf module:CursoFunctions
 * @name readCurso
 * @param {Object} contract - hyperledger fabric contract 
 * @param {String} cursoID - string representing the curso identifier
 * @returns {Object} response from the chaincode
 */
async function readCurso(contract, cursoID) {
    return await contract.evaluateTransaction('ReadCurso', cursoID)
}

/** @function
 * @moduleOf module:CursoFunctions
 * @name readAllCursos 
 * @param {Object} contract - hyperledger fabric contract
 * @returns {Object} response from the chaincode
 */
async function  readAllCursos(contract) {
    return await contract.evaluateTransaction('ReadAllCursos')
}

/** @function
 * @moduleOf module:CursoFunctions
 * @name readCursosFromNome
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} courseName - string representing the curso name
 * @returns {Object} response from the chaincode
 */
async function readCursosFromNome(contract, courseName) {
    return await contract.evaluateTransaction('ReadCursosFromNome', courseName)
}
 
/** @function
 * @moduleOf module:CursoFunctions
 * @name readCursosFromIES
 * @param {Object} contract - hyperledger fabric contract
 * @param {string} codigoIES - string representing the IES identifier
 * @returns {Object} response from the chaincode
 */
async function readCursosFromIES(contract, codigoIES) {
    return await contract.evaluateTransaction('ReadCursosFromIES', codigoIES)
}

/*
#######################
### IES Functions   ###
#######################
*/

/** @module IesFunctions */

/** @function
 * @moduleOf module:IesFunctions
 * @name createIES
 * @param {Object} contract - hyperledger fabric contract
 * @param {json} IESJSON - JSON object representing the IES 
 * @returns {Object} response from the chaincode
 */
async function createIES(contract, IESJSON) {
    return await contract.submitTransaction('CreateIES', IESJSON)
}

/** @function
 * @moduleOf module:IesFunctions
 * @name IESExists
 * @param {Object} contract - hyperledger fabric contract
 * @param {string} codigoIES - string representing the IES identifier
 * @returns {Object} response from the chaincode
 */
async function IESExists(contract, codigoIES) {
    return await contract.evaluateTransaction('IESExists', codigoIES)
}

/** @function
 * @moduleOf module:IesFunctions
 * @name readAllIES
 * @param {Object} contract - hyperledger fabric contract 
 * @returns {Object} response from the chaincode
 */
async function readAllIES(contract) {
    return await contract.evaluateTransaction('ReadAllIES')
}

/** @function
 * @moduleOf module:IesFunctions
 * @name readIES
 * @param {Object} contract - hyperledger fabric contract
 * @param {string} codigoIES - string representing the IES identifier 
 * @returns {Object} response from the chaincode
 */
async function readIES(contract, codigoIES) {
    return await contract.evaluateTransaction('ReadIES', codigoIES)
}

/*
###########################
### Curriculo Functions ###
###########################
*/

/** @function
 * @moduleOf module:IesFunctions
 * @name addCurriculo
 * @param {Object} contract - hyperledger fabric contract 
 * @param {json} curriculoJSON - JSON object representing the curriculo
 * @returns {Object} response from the chaincode
 */
async function addCurriculo(contract, curriculoJSON) {
    return await contract.submitTransaction('AddCurriculo', curriculoJSON)
}

/** @function
 * @moduleOf module:IesFunctions
 * @name curriculoExists
 * @param {Object} contract - hyperledger fabric contract
 * @param {String} curriculoID - string representing the curriculo identifier 
 * @returns {Object} response from the chaincode
 */
async function curriculoExists(contract, curriculoID) {
    return await contract.evaluateTransaction('CurriculoExists', curriculoID)
}

/** @function
 * @moduleOf module:IesFunctions
 * @name readCurriculosFromCurso
 * @param {Object} contract - hyperledger fabric contract 
 * @param {String} cursoID - string representing the curso identifier
 * @returns {Object} response from the chaincode
*/
async function readCurriculosFromCurso(contract, cursoID) {
    return await contract.evaluateTransaction('ReadCurriculosFromCurso', cursoID)
}

/** @function
 * @moduleOf module:IesFunctions
 * @name readCurriculosFromIES
 * @param {Object} contract - hyperledger fabric contract 
 * @param {String} iesID - string representing the IES identifier 
 * @returns {Object} response from the chaincode 
*/
async function readCurriculosFromIES(contract, iesID) {
    return await contract.evaluateTransaction('ReadCurriculosFromIES', iesID)
}

/** @function
 * @moduleOf module:IesFunctions
 * @name readCurriculo
 * @param {Object} contract - hyperledger fabric contract 
 * @param {String} curriculoID - string representing the curriculo identifier 
 * @returns {Object} response from the chaincode
 */
async function readCurriculo(contract, curriculoID) {
    return await contract.evaluateTransaction('ReadCurriculo', curriculoID)
}

/** @function
 * @moduleOf module:IesFunctions
 * @name readAllCurriculos
 * @param {Object} contract - hyperledger fabric contract
 * @returns {Object} response from the chaincode
 */
async function readAllCurriculos(contract) {
    return await contract.evaluateTransaction('ReadAllCurriculos')
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

/** @function
 * @moduleOf module:UtilsFunctions
 * @name jsonPrettyOutput 
 * @param {json} jsonResponse - json string
 * @returns {String} - parse a json to a formatted json as string
 */
function jsonPrettyOutput(jsonResponse) {
    return JSON.stringify(JSON.parse(jsonResponse), undefined, 4)
}

// Exports functions.
module.exports = {
    // Utils functions.
    bufferToStringResponse,
    jsonPrettyOutput,
    // IES functions.
    IESExists,
    createIES,
    readIES,
    readAllIES,
    // Curso functions.
    cursoExists,
    addCurso,
    readCurso,
    readAllCursos,
    readCursosFromNome,
    readCursosFromIES,
    // Curriculo functions.
    addCurriculo,
    curriculoExists,
    readCurriculo,
    readCurriculosFromCurso,
    readCurriculosFromIES,
    readAllCurriculos
}