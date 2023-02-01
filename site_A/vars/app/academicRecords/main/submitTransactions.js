'use strict'

/**@module SubmitTransactions 
 * @description Submits transactions to the ledgers */

const { XMLParser } = require('fast-xml-parser')
const academicRecords = require('../academicRecords.js')
const objConv = require('./utils/assetlibFormater.js')
const CustomError = require('./utils/CustomError')

/**
 * Process transactions of type write
 * @async
 * @param {*} academicRecordsContract - academicRecords contract interface
 * @param {String} reqXml - XML file
 * @returns {Promise<Boolean>} - A boolean value indicating whether the transaction was successful or not
 */

async function submitTransaction(academicRecordsContract, reqXml) {
    const errContext = "Não foi possível submeter a transação."
    if (reqXml === '') {
        throw new CustomError('invalidXMLError', errContext)
    }
    let options = {
        attrNodeName: "#attr",
        textNodeName: "#text",
        attributeNamePrefix: "",
        arrayMode: "false",
        ignoreAttributes: false,
        parseAttributeValue: false,
        ignoreDeclaration: true,
        explicitArray: false,
        parseTagValue: false
    };
    const parser = new XMLParser(options);
    var newXML = reqXml
    let reqJSON = parser.parse(newXML.toString());
    let jsonInAssetlibPattern = objConv.formatHistorico(reqJSON['DocumentoHistoricoEscolarFinal']['infHistoricoEscolar'])
    jsonInAssetlibPattern.digestValue = reqJSON['DocumentoHistoricoEscolarFinal']['Signature']['SignedInfo']['DigestValue']
    let jsonAlunoAssetPattern = objConv.formatAluno(reqJSON['DocumentoHistoricoEscolarFinal']['infHistoricoEscolar']['Aluno'])

    const digest = reqJSON['DocumentoHistoricoEscolarFinal']['Signature']['SignedInfo']['Reference']['DigestValue']
    jsonInAssetlibPattern['digestValue'] = digest

    let exists = "false"
    exists = await academicRecords.alunoExists(academicRecordsContract, jsonAlunoAssetPattern['CPF'])
    exists = academicRecords.bufferToStringResponse(exists) == "true"

    if (exists) {
        await academicRecords.updateAluno(academicRecordsContract, JSON.stringify(jsonAlunoAssetPattern))
        await academicRecords.createHistoricoEscolar(academicRecordsContract, JSON.stringify(jsonInAssetlibPattern))
    } else {
        await academicRecords.createAluno(academicRecordsContract, JSON.stringify(jsonAlunoAssetPattern))
        await academicRecords.createHistoricoEscolar(academicRecordsContract, JSON.stringify(jsonInAssetlibPattern))
    }
    return true
}
module.exports = { submitTransaction }
