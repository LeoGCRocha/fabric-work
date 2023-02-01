'use strict'
/**@module EvaluateTransactions */
const interfaceDecree = require('../decree.js')
const CustomError = require('./utils/CustomError')

/**
 * Process transactions of type read
 * @param {*} contract - Decree contract interface
 * @param {Array} params - Formated struct with parameters
 * @returns {Promise<String>} - Fabric contract result message
 */

async function evaluateTransactions(contract, params) {

    const errContext = "Não foi possível realizar a transação"

    // Checks which read transaction needs to be made
    let response
    switch (params.type) {
    // Receive a ies code and return the ies information if exists
    case 'readIES':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.readIES(contract, params.codigoIES))
        return interfaceDecree.jsonPrettyOutput(response)
    // Return all ies information
    case 'readAllIES':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.readAllIES(contract))
        return interfaceDecree.jsonPrettyOutput(response)
    // Receive a ies code and return true if exists
    case 'iesExists':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.IESExists(contract, params.codigoIES))
        return interfaceDecree.jsonPrettyOutput(JSON.stringify({
            'codigoIES': params.codigoIES,
            'exists': response
        }))
    // Receive a curso code and return true if exists
    case 'cursoExists':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.cursoExists(contract, params.codigoCurso))
        return interfaceDecree.jsonPrettyOutput(JSON.stringify({
            'codigoCurso': params.codigoCurso,
            'exists': response
        }))
    // Receive a curso code and return the curso information if exists
    case 'readCurso':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.readCurso(contract, params.codigoCurso))
        return interfaceDecree.jsonPrettyOutput(response)
    // Return a list of all cursos
    case 'readAllCursos':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.readAllCursos(contract))
        return interfaceDecree.jsonPrettyOutput(response)
    // Receive a curse name and return a list of cursos with that name
    case 'readCursosFromNome':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.readCursosFromNome(contract, params.nome))
        return interfaceDecree.jsonPrettyOutput(response)
    // Receive a IES code and return a list of cursos with that IES
    case 'readCursosFromIES':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.readCursosFromIES(contract, params.codigoIES))
        return interfaceDecree.jsonPrettyOutput(response)
    // Receive a curso code and return all curriculo from that curso
    case 'readCurriculosFromCurso':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.readCurriculosFromCurso(contract, params.codigoCurso))
        return interfaceDecree.jsonPrettyOutput(response)
    // Receive a IES code and return all curriculo from that IES
    case 'readCurriculosFromIES':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.readCurriculosFromIES(contract, params.codigoIES))
        return interfaceDecree.jsonPrettyOutput(response)
    // Receive a curriculo code and return the curriculo information if exists
    case 'readCurriculo':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.readCurriculo(contract, params.codigoCurriculo))
        return interfaceDecree.jsonPrettyOutput(response)    
    // Receive a curriculo code and return true if exists
    case 'curriculoExists':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.curriculoExists(contract, params.codigoCurriculo))
        return interfaceDecree.jsonPrettyOutput(JSON.stringify({
            'codigoCurriculo': params.codigoCurriculo,
            'exists': response
        }))
    // Return a list of all curriculos
    case 'readAllCurriculos':
        response = interfaceDecree.bufferToStringResponse(await interfaceDecree.readAllCurriculos(contract))
        return interfaceDecree.jsonPrettyOutput(response)
    default:
        throw new CustomError('invalidFunctionNameError', errContext)
    }
}

module.exports = {evaluateTransactions}