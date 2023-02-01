'use strict'

/**@module EvaluateTransactions */
const academicRecords = require('../academicRecords')
const CustomError = require('./utils/CustomError')

/**
 * Process transactions of type read
 * @param {*} academicRecordsContract - Academic Records contract interface
 * @param {json} option - Define the type and parammeters about this executation
 */
async function evaluateTransaction(academicRecordsContract, option) {
    const errContext = "Não foi possível realizar a transação"

    let response
    let jsonReffer

    // Type check.
    switch (option['type']) {
        
    // Verify if aluno exists and return true or false as response.
    case 'alunoExists':
        response = await academicRecords.alunoExists(academicRecordsContract, option['CPF'])
        return  academicRecords.bufferToStringResponse(response)

    // Get all alunos and return them as response.
    case 'getAllAlunos':
        response = await academicRecords.getAllAlunos(academicRecordsContract)
        return academicRecords.jsonPrettyOutput(academicRecords.bufferToStringResponse(response))

    // Get aluno and return it as response, if aluno does not exist return error.
    case 'readAluno':
        response = await academicRecords.readAluno(academicRecordsContract, option['CPF'])
        return academicRecords.jsonPrettyOutput(academicRecords.bufferToStringResponse(response))

    // Get historicos from Aluno and return them as response.
    case 'readHistoricosFromAluno':
        response = await academicRecords.readHistoricosFromAluno(academicRecordsContract, option['CPF'])
        return academicRecords.jsonPrettyOutput(academicRecords.bufferToStringResponse(response))

    // Get last historico from Aluno and return it as response.
    case 'readLastHistoricoFromAluno':
        response = await academicRecords.readLastHistoricoFromAluno(academicRecordsContract, option['CPF'])
        return academicRecords.jsonPrettyOutput(academicRecords.bufferToStringResponse(response))

    // Get historicos from IES and return them as response.
    case 'readHistoricosFromIes':
        response = await academicRecords.readHistoricoFromIes(academicRecordsContract, option['IES'])
        return academicRecords.jsonPrettyOutput(academicRecords.bufferToStringResponse(response))

    // Get historicos from Curso and return them as response.
    case 'readHistoricosFromCurso':
        response = await academicRecords.readHistoricoFromCurso(academicRecordsContract, option['codigoCurso'])
        return academicRecords.jsonPrettyOutput(academicRecords.bufferToStringResponse(response))

    // Get historico from all parameters and return it as response.
    case 'readHistorico':
        jsonReffer = {
            iesEmissora: option['iesEmissora'],
            curso: option['curso'],
            curriculo: option['curriculo'],
            aluno: option['aluno'],
            digestValue: option['digestValue']
        }
        response = await academicRecords.readHistorico(academicRecordsContract,
            JSON.stringify(jsonReffer))
        return academicRecords.jsonPrettyOutput(academicRecords.bufferToStringResponse(response))
    case 'getCategoriasAtividades':
        response = await academicRecords.getCategoriasAtividades(academicRecordsContract, option['CPF'])
        return  academicRecords.bufferToStringResponse(response)
    case 'readAtividadesComplementaresAluno':
        response = await academicRecords.readAtividadesComplementaresAluno(academicRecordsContract, option['CPF'])
        return  academicRecords.bufferToStringResponse(response)
    case 'readEstagiosFromAluno':
        response = await academicRecords.readEstagiosFromAluno(academicRecordsContract, option['CPF'])
        return  academicRecords.bufferToStringResponse(response)
    default:
        throw new CustomError('invalidFunctionNameError', errContext)        
    }
}

module.exports = { evaluateTransaction }