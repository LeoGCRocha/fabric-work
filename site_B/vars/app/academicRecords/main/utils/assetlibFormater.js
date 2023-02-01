'use strict'

/** @function
 *  @name formatAluno
 *  @param {Object} obj - Object with aluno data
 *  @returns {Object} - Formatted aluno data
 */

const crypto = require('crypto');

function formatAluno(aluno) {
    let formattedAluno = {
        "ID": aluno["ID"],
        "nome": aluno["Nome"],
        "sexo": aluno["Sexo"],
        "nacionalidade": aluno["Nacionalidade"],
        "CPF": aluno["CPF"],
        "dataNascimento": formatDateType(aluno["DataNascimento"]),
        "naturalidade": {},
        "nomeSocial": "",
        historicos: []
    }

    // Naturalidade
    if (aluno['Naturalidade']['NomeMunicipioEstrangeiro'] != undefined) {
        formattedAluno["naturalidade"] = {
            "nomeMunicipioEstrangeiro": aluno['Naturalidade']['NomeMunicipioEstrangeiro'],
            "ehEstrangeiro": true
        }
    } else {
        formattedAluno["naturalidade"] = {
            "codigoMunicipio": aluno["Naturalidade"]["CodigoMunicipio"],
            "nomeMunicipio": aluno["Naturalidade"]["NomeMunicipio"],
            "UF": aluno["Naturalidade"]["UF"],
            "ehEstrangeiro": false
        }
    }

    // Nome Social
    if (aluno["NomeSocial"] != undefined) {
        formattedAluno["nomeSocial"] = aluno["NomeSocial"]
    }

    // RG
    if (aluno['RG'] != undefined) {
        formattedAluno['RG'] = {
            "numero": aluno["RG"]["Numero"].toString(),
            "orgaoExpedidor": aluno["RG"]["OrgaoExpedidor"],
            "UF": aluno["RG"]["UF"].toString()
        }
        formattedAluno["outroDocumento"] = {}
        formattedAluno["documentoRG"] = true
    } else {
        formattedAluno["outroDocuento"] = {
            "tipoDocumento": aluno["OutroDocumentoIdentificacao"]["TipoDocumento"],
            "identificador": aluno["OutroDocumentoIdentificacao"]["Identificador"]
        }
        formattedAluno["RG"] = {}
        formattedAluno["documentoRG"] = false
    }
    return formattedAluno
} 

/** @function
 *  @name formatHistorico
 *  @param {Object} obj - Object with historico data
 *  @returns {Object} - Formatted historico data
 */
function formatHistorico(obj) {
    
    // Select the actual type of course
    let curso
    obj['DadosCurso'] !== undefined ? curso = obj['DadosCurso'] : curso = obj['DadosCursoNSF']

    let historicoEscolar = obj['HistoricoEscolar']
    let situacaoAtual = historicoEscolar['SituacaoAtualDiscente']
    let situacao = formatSituacao(situacaoAtual)
    let cargaHorariaCurso = formatCargaHoraria(historicoEscolar['CargaHorariaCurso'])
    let cargaHorariaIntegralizada = formatCargaHoraria(historicoEscolar['CargaHorariaCursoIntegralizada'])
    let elementoHistorico = formatElementoHistorico(historicoEscolar['ElementosHistorico'], formatCurso(curso))
    let historicoJsonObj = {
        "aluno": obj["Aluno"]["CPF"],
        "curso": formatCurso(curso),
        "iesEmissora": obj["IesEmissora"]["CodigoMEC"], 
        "situacaoAtualDiscente": {
            "tipoDeSituacao" : situacao
        },

        // AAAA-MM-DDs
        // Represents two different fields date and time
        "dataHoraEmissao": formatDateTypeHistorico(historicoEscolar['DataEmissaoHistorico'], historicoEscolar['HoraEmissaoHistorico']),
        
        // CargaHorariaCurso and CargaHorariaCursoIntegralizada need to be compared with curriculum
        "cargaHorariaCurso": cargaHorariaCurso,
        "cargaHorariaCursoIntegralizada": cargaHorariaIntegralizada,
        "ingressoCurso": {
            "formaDeAcesso": historicoEscolar['IngressoCurso']['FormaAcesso'],
            "data": formatDateType(historicoEscolar['IngressoCurso']['Data'])
        },
        "curriculo": "",
        "nomeParaAreas": "",
        "areas": [],
        "elementoHistorico": elementoHistorico,
        "ENADE": formatENADE(historicoEscolar["ENADE"]),
        "codigoValidacao": "",
        "digestValue": "",
        "informacoesAdicionais": ""
    }

    // Codigo de Validacao
    if (obj['SegurancaHistorico']['CodigoValidacao'] != undefined) {
        historicoJsonObj['codigoValidacao'] = obj['SegurancaHistorico']['CodigoValidacao']
    }

    // Periodo Letivo
    if (situacaoAtual["PeriodoLetivo"] != undefined) {
        historicoJsonObj['situacaoAtualDiscente']['periodoLetivo'] = situacaoAtual['PeriodoLetivo']
    }

    // Codigo Curriculo
    if (historicoEscolar['CodigoCurriculo'] !== undefined) {
        historicoJsonObj['curriculo'] = historicoEscolar['CodigoCurriculo']
    }

    // Informacoes Adicionais
    if (historicoEscolar['InformacoesAdicionais'] !== undefined) {
        historicoJsonObj['informacoesAdicionais'] = historicoEscolar['InformacoesAdicionais']
    }

    // Nome para Areas
    if (historicoEscolar['NomeParaAreas'] !== undefined) {
        historicoJsonObj['nomeParaAreas'] = historicoEscolar['NomeParaAreas']
    }
     
    // Areas
    if (historicoEscolar['Areas'] !== undefined) {
        let areas = historicoEscolar['Areas']['Area']
        if (Array.isArray(areas)) {
            areas.forEach((area) => {
                historicoJsonObj['areas'].push({
                    "codigo": area['Codigo'],
                    "nome": area['Nome']
                })
            })
        } else {
            historicoJsonObj['areas'].push({
                "codigo": areas['Codigo'],
                "nome": areas['Nome']
            })
        }
    }

    // Prepare struct SituacaoAtualDiscente
    historicoJsonObj['situacaoAtualDiscente']['Formado'] = {}
    historicoJsonObj['situacaoAtualDiscente']['Intercambio'] = {}
    let tipoDeSituacao = historicoJsonObj['situacaoAtualDiscente']['tipoDeSituacao']
    if (tipoDeSituacao === "Formado") {
        let situacaoFormado = situacaoAtual['Formado']
        historicoJsonObj['situacaoAtualDiscente']['Formado'] = formatSituacaoFormado(situacaoFormado)
    } else if (tipoDeSituacao == "IntercambioNacional" || tipoDeSituacao == "IntercambioInternacional") {
        let situacaoIntercambio = situacaoAtual[tipoDeSituacao]
        historicoJsonObj['situacaoAtualDiscente'][tipoDeSituacao] = "Intercambio"
        historicoJsonObj['situacaoAtualDiscente']['Intercambio'] = formatSituacaoIntercambio(situacaoIntercambio, tipoDeSituacao)
    }
    return historicoJsonObj
}

/** @function
 *  @name formatENADE
 *  @param {Object} obj - Object with ENADE data
 *  @returns {Object} - Formatted ENADE data
 */
function formatENADE(obj) {
    let ENADE = {
        "habilitacoes": [],
        "naoHabilitacoes": [],
        "irregulares": []
    }
    if (obj != undefined) {

        if (obj['Habilitado'] != undefined) {
            let arrHabilitacoes = obj['Habilitado']
            if (Array.isArray(arrHabilitacoes)) {
                for (let i = 0; i < arrHabilitacoes.length; i++ ) {
                    let habilitacaoAtual = arrHabilitacoes[i]
                    let habilitacao = {
                        "edicao": Number(habilitacaoAtual['Edicao']),
                        "habilitado": true,
                        "condicao": habilitacaoAtual['Condicao']
                    }
                    ENADE['habilitacoes'].push(habilitacao)
                }
            } else {
                ENADE['habilitacoes'].push({
                    "edicao": Number(arrHabilitacoes['Edicao']),
                    "habilitado": true,
                    "condicao": arrHabilitacoes['Condicao']
                })
            }
        }
        if (obj['Irregular'] != undefined) {
            let arrIrregulares = obj['Irregular']
            if (Array.isArray(arrIrregulares)) {
                for (let i = 0; i < arrIrregulares.length; i++ ) {
                    let irregularAtual = arrIrregulares[i]
                    let irregular = {
                        "edicao": irregularAtual['Edicao'],
                        "habilitado": false,
                        "condicao": irregularAtual['Condicao']
                    }
                    ENADE['irregulares'].push(irregular)
                }
            } else {
                ENADE['irregulares'].push({
                    "edicao": Number(arrIrregulares['Edicao']),
                    "habilitado": false,
                    "condicao": arrIrregulares['Condicao']
                })
            }
        }
        if (obj['NaoHabilitado'] != undefined) {
            let arrNaoHabilitacoes = obj['NaoHabilitado']
            if (Array.isArray(arrNaoHabilitacoes)) {
                for (let i = 0; i < arrNaoHabilitacoes.length; i++ ) {
                    let naoHabilitacaoAtual = arrNaoHabilitacoes[i]
                    let naoHabilitacao = {
                        "edicao": Number(naoHabilitacaoAtual['Edicao']),
                        "habilitado": false,
                        "condicao": naoHabilitacaoAtual['Condicao'],
                        "motivo": "",
                        "outroMotivo": ""
                    }
                    if (naoHabilitacaoAtual['Motivo'] != undefined) {
                        naoHabilitacao['motivo'] = naoHabilitacaoAtual['Motivo']
                    } else if (naoHabilitacao['OutroMotivo'] != undefined) {
                        naoHabilitacao['outroMotivo'] = naoHabilitacaoAtual['OutroMotivo']
                    }
                    ENADE['naoHabilitacoes'].push(naoHabilitacao)
                }
            } else {
                let naoHabilitacao = {
                    "edicao": Number(arrNaoHabilitacoes['Edicao']),
                    "habilitado": false,
                    "condicao": arrNaoHabilitacoes['Condicao'],
                    "motivo": "",
                    "outroMotivo": ""
                }
                if (arrNaoHabilitacoes['Motivo'] != undefined) {
                    naoHabilitacao['motivo'] = arrNaoHabilitacoes['Motivo']
                } else if (arrNaoHabilitacoes['OutroMotivo'] != undefined) {
                    naoHabilitacao['outroMotivo'] = arrNaoHabilitacoes['OutroMotivo']
                }
                ENADE['naoHabilitacoes'].push(naoHabilitacao)
            }
        }
    }
    return ENADE
}

/** @function
 *  @name formatDateType
 *  @param {Object} date - Unformatted Date
 *  @returns {Object} - Formatted Date
 */
function formatDateType(date) {
    return `${date}T00:00:00.000Z`
}

/** @function
 *  @name formatDateType
 *  @param {Object} date - Unformatted Date
 *  @returns {Object} - Formatted Date
 */
function formatDateTypeHistorico(date, time) {
    return `${date}T${time}.000Z`
}

/** @function
 *  @name formatSituacaoFormado
 *  @param {Object} obj - Object with SituacaoFormado data
 *  @returns {Object} - Formatted SituacaoFormado data
 */
function formatSituacaoFormado(obj) {
    return {
        "dataConclusaoCurso": formatDateType(obj['DataConclusaoCurso']),
        "dataColacaoGrau": formatDateType(obj['DataColacaoGrau']),
        "dataExpedicaoDiploma": formatDateType(obj['DataExpedicaoDiploma'])
    }
}

/** @function
 *  @name formatSituacaoIntercambio
 *  @param {Object} obj - Object with SituacaoIntercambio data
 *  @returns {Object} - Formatted SituacaoIntercambio data
 */
function formatSituacaoIntercambio(obj, tipoDeSituacao) {
    return {
        "instituicao": obj['Instituicao'],
        "pais": obj['Pais'],
        "programaIntercambio": obj['NomeProgramaIntercambio'],
        tipoDeIntercambio: tipoDeSituacao.substring(11, tipoDeSituacao.length)
    }
}

/** @function
 *  @name formatCurso
 *  @param {Object} obj - Object with Curso data
 *  @returns {string} - idCurso
 */
function formatCurso(obj) {
    let idCurso
    if (obj['SemCodigoCursoEMEC'] != undefined) {
        idCurso = obj['SemCodigoCursoEMEC']['NumeroProcesso']
    } else {
        idCurso = obj['CodigoCursoEMEC']
    }
    return idCurso.toString()
}

/** @function
 *  @name formatCargaHoraria
 *  @param {Object} obj - Object with CargaHoraria data
 *  @returns {Object} - Formatted CargaHoraria data
 */
function formatCargaHoraria(obj) {
    if (obj['HoraAula'] != undefined) {
        return {
            "horaAula": true, 
            "cargaHoraria": Number(obj['HoraAula'])
        }
    } else if (obj['HoraRelogio'] != undefined) {
        return {
            "horaAula": false,
            "cargaHoraria": Number(obj['HoraRelogio'])
        }
    }
}

/** @function
 *  @name formatSituacaoElementoHistorico
 *  @param {Object} obj - Object with SituacaoElementoHistorico data
 *  @returns {Object} - Formatted SituacaoElementoHistorico data
 */
function formatSituacaoElementoHistorico(obj) {
    let situacaoAtual = obj
    let tipoDeSituacao = formatSituacao(situacaoAtual)
    let formattedSituacao = {
        "periodoLetivo": situacaoAtual['PeriodoLetivo'].toString(),
        "tipoDeSituacao": tipoDeSituacao,
    }
    formattedSituacao['Formado'] = {}
    formattedSituacao['Intercambio'] = {}
    if (tipoDeSituacao === "Formado") {
        let situacaoFormado = situacaoAtual['Formado']
        formattedSituacao['Formado'] = formatSituacaoFormado(situacaoFormado)
    } else if (tipoDeSituacao == "IntercambioNacional" || tipoDeSituacao == "IntercambioInternacional") {
        let situacaoIntercambio = situacaoAtual[tipoDeSituacao]
        formattedSituacao['tipoDeSituacao'] = "Intercambio"
        formattedSituacao['Intercambio'] = formatSituacaoIntercambio(situacaoIntercambio, tipoDeSituacao)
    }
    return formattedSituacao
}

/** @function
 *  @name formatDisciplinaElementoHistorico
 *  @param {Object} obj - Object with DisciplinaElementoHistorico data
 *  @param {int} codigoCursoEMEC - Curso ID
 *  @returns {Object} - Formatted DisciplinaElementoHistorico data
 */
function formatDisciplinaElementoHistorico(obj, codigoCursoEMEC) {
    let disciplinaAtual = obj
    let formattedDisciplina = {
        "nome": disciplinaAtual["NomeDisciplina"],
        "periodo": disciplinaAtual["PeriodoLetivo"],
        "codigo": disciplinaAtual["CodigoDisciplina"],
        "codigoCursoEMEC": codigoCursoEMEC,
        "cargaHoraria": formatCargaHoraria(disciplinaAtual['CargaHoraria']),
        "docentes": [], // There is not a docentes attr in .xml
        "curricular": false,
        "tipoDeNota": formatTipoNota(disciplinaAtual),
        "estadoDisciplina": formatEstadoDisciplina(disciplinaAtual)
    }
    formattedDisciplina['nota'] = disciplinaAtual[formattedDisciplina["tipoDeNota"]].toString()
    return formattedDisciplina
}

/** @function
 *  @name formatElementoHistorico
 *  @param {Object} elementoHistorico - Object with ElementoHistorico data
 *  @param {int} codigoCursoEMEC - Curso ID
 *  @returns {Object} - Formatted ElementoHistorico data
 */
function formatElementoHistorico(elementoHistorico, codigoCursoEMEC) {
    let formattedElementoHistorico = {
        "situacoes": [],
        "disciplinas": [],
        "atividadeComplementares": [],
        "estagios": []
    }
    
    // All situacoes
    let arrSituacao = elementoHistorico['SituacaoDiscente']
    if (Array.isArray(arrSituacao)) {
        for (let i = 0; i < arrSituacao.length; i++ ) {
            let situacao = formatSituacaoElementoHistorico(arrSituacao[i])
            formattedElementoHistorico['situacoes'].push(situacao)
        }
    } else if (arrSituacao != undefined) {
        formattedElementoHistorico['situacoes'].push(formatSituacaoElementoHistorico(arrSituacao))
    }
    
    // All disciplinas
    let arrDisciplina = elementoHistorico['Disciplina']
    if (Array.isArray(arrDisciplina)) {
        for (let i = 0; i < arrDisciplina.length; i++) {
            let disciplina = formatDisciplinaElementoHistorico(arrDisciplina[i], codigoCursoEMEC)
            formattedElementoHistorico['disciplinas'].push(disciplina)
        }
    } else if (arrDisciplina != undefined) {
        formattedElementoHistorico['disciplinas'].push(formatDisciplinaElementoHistorico(arrDisciplina, codigoCursoEMEC))
    }

    // All atividades complementares
    let arrAtividadesComplementares = elementoHistorico['AtividadeComplementar']
    if (Array.isArray(arrAtividadesComplementares)) {
        for (let i = 0; i < arrAtividadesComplementares.length; i++) {
            let atividadeComplementar = formatAtividadeComplementar(arrAtividadesComplementares[i])
            formattedElementoHistorico['atividadeComplementares'].push(atividadeComplementar)
        }
    } else if (arrAtividadesComplementares != undefined){
        formattedElementoHistorico['atividadeComplementares'].push(formatAtividadeComplementar(arrAtividadesComplementares))
    }
    
    // All estagios
    let arrEstagios = elementoHistorico['Estagio']
    if (Array.isArray(arrEstagios)) {
        for (let i = 0; i < arrEstagios.length; i++) {
            let estagio = formatEstagio(arrEstagios[i])
            formattedElementoHistorico['estagios'].push(estagio)
        }
    } else if (arrEstagios != undefined){
        formattedElementoHistorico['estagios'].push(formatEstagio(arrEstagios))
    }
    return formattedElementoHistorico
}

/** @function
 *  @name formatEstagio
 *  @param {Object} estagio - Object with Estagio data
 *  @returns {Object} - Formatted Estagio data
 */
function formatEstagio(estagio) {
    let formattedEstagio = {
        "id": crypto.randomUUID(),
        "situacao": "Aprovado",
        "dataInicio": formatDateType(estagio['DataInicio']),
        "dataFim": formatDateType(estagio['DataFim']),
        "cargaHorariaEmHorasRelogio": {
            "horaAula": false,
            "cargaHoraria": Number(estagio['CargaHorariaEmHorasRelogio'])
        },
        "docentesOrientadores": [],
        "concedente": {},
        "descricao": ""
    }
    if (estagio['Descricao'] != undefined) {
        formattedEstagio['descricao'] = estagio['Descricao']
    } 
    if (estagio['Concedente'] != undefined) {
        let concedente = {
            "razaoSocial": estagio['Concedente']['RazaoSocial'],
            "CNPJ": estagio['Concedente']['CNPJ'].toString()
        }
        if (estagio['Concedente']['NomeFantasia'] != undefined) {
            concedente['nomeFantasia'] = estagio['Concedente']['NomeFantasia']
        }
        formattedEstagio['concedente'] = concedente
    }
    formattedEstagio["docentesOrientadores"] = formatDocentesResponsaveisPelaValidacao(estagio['DocentesOrientadores'])
    return formattedEstagio
}

/** @function
 *  @name formatAtividadeComplementar
 *  @param {Object} atividadeComplementar - Object with AtividadeComplementar data
 *  @returns {Object} - Formatted AtividadeComplementar data
 */
function formatAtividadeComplementar(atividadeComplementar) {
    let formattedAtividadeComplementar = {
        "id": crypto.randomUUID(),
        "codigo": "",
        "situacao": "Aprovado",
        "dataInicio": formatDateType(atividadeComplementar['DataInicio']),
        "dataFim": formatDateType(atividadeComplementar['DataFim']),
        "dataRegistro": formatDateType(atividadeComplementar['DataRegistro']),
        "cargaHorariaEmHoraRelogio": {
            "horaAula": false,
            "cargaHoraria": Number(atividadeComplementar['CargaHorariaEmHoraRelogio'])
        },
        "tipoAtividadeComplementar": atividadeComplementar['TipoAtividadeComplementar'],
        "descricao": "",
        "docentesResponsaveisPelaValidacao": []
    }
    if (atividadeComplementar['CodigoAtividadeComplementar'] != undefined) {
        formattedAtividadeComplementar['codigo'] = atividadeComplementar['CodigoAtividadeComplementar']
    }
    if (atividadeComplementar['Descricao'] != undefined) {
        formattedAtividadeComplementar['descricao'] = atividadeComplementar['Descricao']
    }
    formattedAtividadeComplementar["docentesResponsaveisPelaValidacao"] = formatDocentesResponsaveisPelaValidacao(atividadeComplementar['DocentesResponsaveisPelaValidacao'])
    return formattedAtividadeComplementar
}

/** @function
 *  @name formatDocente
 *  @param {Object} docente - Object with Docente data
 *  @returns {Object} - Formatted Docente data
 */
function formatDocente(docente) {
    let formattedDocente = {
        "nome": docente['Nome'],
        "lattes": "",
        "CPF": "",
        "titulacao": docente['Titulacao']
    }
    if (docente['Lattes'] != undefined) {
        formattedDocente["lattes"] = docente['Lattes']
    }
    if (docente['CPF'] != undefined) {
        formattedDocente["CPF"] = docente['CPF']
    }
    return formattedDocente
}

/** @function
 *  @name formatDocentesResponsaveisPelaValidacao
 *  @param {Object} docentesResponsaveis - Object with DocentesResponsaveisPelaValidacao data
 *  @returns {Object} - Formatted DocentesResponsaveisPelaValidacao data
 */
function formatDocentesResponsaveisPelaValidacao(docentesResponsaveis) {
    if (Array.isArray(docentesResponsaveis)) {
        let docentes = []
        for (let i = 0; i < docentesResponsaveis.length; i++) {
            let docente = formatDocente(docentesResponsaveis[i]['Docente'])
            docentes.push(docente)
        }
        return docentes
    } else {
        return [formatDocente(docentesResponsaveis['Docente'])]
    }
}

/** @function
 *  @name formatEstadoDisciplina
 *  @param {Object} estado - Object with EstadoDisciplina data
 *  @returns {Object} - Formatted EstadoDisciplina data
 */
function formatEstadoDisciplina(estado) {
    let typeOfEstados = ["Aprovado", "Pendente", "Reprovado"]
    for (let i = 0; i < typeOfEstados.length; i++) {
        if (estado[typeOfEstados[i]] != undefined) {
            let formattedEstado = {
                "tipoDeEstadoDisciplina": typeOfEstados[i]
            }
            if (typeOfEstados[i] == "Aprovado") {
                formattedEstado["integralizacao"] = {
                    "formaIntegralizacao" : "",
                    "outraFormaIntegralizacao": ""
                }
                if (estado[typeOfEstados[i]]['FormaIntegralizacao'] != undefined) {
                    formattedEstado["integralizacao"]["formaIntegralizacao"] = estado[typeOfEstados[i]]['FormaIntegralizacao']
                } else if (estado[typeOfEstados[i]]['OutraFormaIntegralizacao'] != undefined) {
                    formattedEstado["integralizacao"]["outraFormaIntegralizacao"]
                }
            }
            return formattedEstado
        }
    }
}

/** @function
 *  @name formatTipoNota
 *  @param {Object} typeGrade - Object with TipoNota data
 *  @returns {Object} - Formatted TipoNota data
 */
function formatTipoNota(typeGrade) {
    let typesOfGrade = ["Nota", "NotaAteCem", "Conceito", "ConceitoRM", "ConceitoEspecificoDoCurso"]
    for (let i = 0; i < typesOfGrade.length; i++) {
        if (typeGrade[typesOfGrade[i]] != undefined) {
            return typesOfGrade[i]
        } 
    }
}

/** @function
 *  @name formatSituacao
 *  @param {Object} situacao - Object with Situacao data
 *  @returns {Object} - Formatted Situacao data
 */
function formatSituacao(situacao) {
    let typeOfSituacoes = ["Trancamento", "MatriculadoEmDisciplina", "Licenca", "Desistencia", "Abandono", "Jubilado", "OutraSituacao",
        "Formado", "IntercambioInternacional", "IntercambioNacional"]
    for (let i = 0; i < typeOfSituacoes.length; i++) {
        if ((situacao[typeOfSituacoes[i]] != undefined) || (situacao["TipoDeSituacao"] == typeOfSituacoes[i])) {
            return typeOfSituacoes[i]
        }
    }
}

module.exports = {formatHistorico, formatAluno}