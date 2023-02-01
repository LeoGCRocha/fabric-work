'use strict'

/**@module SubmitTransactions 
 * @description Submits transactions to the ledgers */
const { XMLParser } = require('fast-xml-parser')
const interfaceDecree = require('../decree.js')
const diffUtil = require('./utils/diff.js')
const objConv = require('./utils/assetlibFormater.js')
const CustomError = require('./utils/CustomError')

/**
 * Process transactions of type write
 * @async
 * @param {*} decreeContract - Decree contract interface
 * @param {String} reqPath - Path to XML file
 * @returns {Promise<Boolean>} - A boolean value indicating whether the transaction was successful or not
 */
async function submitTransactions(decreeContract, reqPath) {
    const errContext = "Nao foi possivel submeter a transacao"

    if (reqPath === '') {
        throw new CustomError('invalidJsonParameter', errContext)
    }
    const alwaysArray = ['CurriculoEscolar.infCurriculoEscolar.infEstruturaCurricular.UnidadeCurricular']
    const options = {
        attrNodeName: "#attr",
        textNodeName: "#text",
        attributeNamePrefix: "",
        ignoreAttributes: false,
        ignoreDeclaration: true,
        parseAttributeValue: true,
        parseTagValue: false,
        format: false,
        suppressEmptyNode: false,
        isArray: (_, jpath) => {
            if (alwaysArray.indexOf(jpath) !== -1) return true
        }
    };

    // Parse XML to JSON.
    const parser = new XMLParser(options)

    // Reads file containing xml processed by API.
    let reqXml = reqPath
    let reqJSON = parser.parse(reqXml)

    // Reads XML stored in ledger.
    let ledgerXml
    let infoIES = reqJSON['CurriculoEscolar']['infCurriculoEscolar']['IesEmissora']
    
    let cursoRef = reqJSON['CurriculoEscolar']['infCurriculoEscolar']
    let typeCurso = cursoRef['DadosCurso'] !== undefined ? 'DadosCurso' : 'DadosCursoNSF'
    let cursoID 

    if (cursoRef[typeCurso]['CodigoCursoEMEC'] !== undefined) {
        cursoID = cursoRef[typeCurso]['CodigoCursoEMEC']
    } else {
        cursoID = cursoRef[typeCurso]['SemCodigoCursoEMEC']['NumeroProcesso']
    }

    // Check if IES exists.
    let iesExists = 'false'
    try {
        iesExists = await interfaceDecree.IESExists(decreeContract, infoIES['CodigoMEC'])
    } catch (err) {
        throw new CustomError('IESExistsError', errContext)
    }

    if (interfaceDecree.bufferToStringResponse(iesExists) === 'false') {
        let newIES = objConv.formatIES(infoIES)
        try {
            await interfaceDecree.createIES(decreeContract, JSON.stringify(newIES))
            try {
                await interfaceDecree.addCurso(decreeContract, JSON.stringify(objConv.formatCurso(reqJSON)))
                return true
            } catch (error) {
                throw new CustomError('courseCreationError', errContext)
            }
        } catch (error) {
            throw new CustomError('IESCreationError', errContext)
        }
    }

    // XML doesnt exists in ledger.
    let ledgerJSON
    let courseIsOffered = 'false'
    if (ledgerXml !== undefined) {
        try {
            courseIsOffered = await interfaceDecree.cursoExists(decreeContract, cursoID)
            if (interfaceDecree.bufferToStringResponse(courseIsOffered) === 'true') {
                ledgerJSON = parser.parse(ledgerXml)
                if (diffUtil.diffHist(ledgerJSON, reqJSON)) {
                    await interfaceDecree.addCurriculo(decreeContract, JSON.stringify(objConv.formatCurriculo(reqJSON)))
                }
            }
            return true
        } catch (error) {
            throw new CustomError('ledgerCreationError', errContext)
        }
    } else {
        try {
    
            // Check if the course exists.
            courseIsOffered = await interfaceDecree.cursoExists(decreeContract, cursoID)
            if (interfaceDecree.bufferToStringResponse(courseIsOffered) === 'false') {
                await interfaceDecree.addCurso(decreeContract, JSON.stringify(objConv.formatCurso(reqJSON)))
            }
        } catch (error) {
            throw new CustomError('courseCreationError', errContext)
        }
    }
}

module.exports = { submitTransactions }