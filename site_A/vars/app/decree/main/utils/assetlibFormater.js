'use strict'

/** @function
 *  @name formatIES
 *  @param {Object} ies - Unformatted IES object
 *  @returns {Object} - Standardized IES object for the chaincode
 */
function formatIES(ies) {
    let formattedIES = {
        'nome': ies['Nome'],
        'codigoMEC': ies['CodigoMEC'],
        'CNPJ': ies['CNPJ'],
        'mantenedora': {},
        'cursos': [],
        'endereco': {},
        'credenciamento': {},
        'recredenciamento': {},
        'renovacaoRecredenciamento': {}
    }

    // Endereco
    formattedIES.endereco = formatEndereco(ies['Endereco'])

    // Mantenedora (minoccurs 0)
    formattedIES['mantenedora'] = formatMantenedora(ies['Mantenedora'])

    // Credenciamento
    formattedIES['credenciamento'] = formatAtoRegulatorio(ies['Credenciamento'])

    // Recredenciamento
    if (ies['Recredenciamento'] !== undefined) {
        formattedIES['recredenciamento'] = formatAtoRegulatorio(ies['Recredenciamento'])
    }

    // RenovacaoRecredenciamento
    if (ies['RenovacaoDeRecredenciamento'] !== undefined) {
        formattedIES['renovacaoRecredenciamento'] = formatAtoRegulatorio(ies['RenovacaoDeRecredenciamento'])
    }
    
    return formattedIES
}


/** @function
 *  @name formatCurso
 *  @param {Object} curso - Unformatted curso object
 *  @returns {Object} - Standardized curso object for the chaincode
 */
function formatCurso(curso) {
    let infoCourse = curso['CurriculoEscolar']['infCurriculoEscolar']
    let tipoCurso = infoCourse['DadosCurso'] !== undefined ? 'DadosCurso' : 'DadosCursoNSF'
    let formattedCurso = {
        'nome': infoCourse[tipoCurso]['NomeCurso'],
        'habilitacoes': [],
        'codigoCursoEMEC': '',
        'autorizacao': {},
        'reconhecimento': {},
        'renovacaoReconhecimento': {},
        'tramitacaoMEC': {},
        'curriculos': [formatCurriculo(curso)],
        'tipo': '',
        'codigoIES': infoCourse['IesEmissora']['CodigoMEC']
    }

    if (infoCourse[tipoCurso]['CodigoCursoEMEC'] !== undefined) {
        formattedCurso['codigoCursoEMEC'] = infoCourse[tipoCurso]['CodigoCursoEMEC']
    } else {
        formattedCurso['tramitacaoMEC'] = formatTramitacaoMEC(infoCourse[tipoCurso]['SemCodigoCursoEMEC'])
    }

    if (infoCourse[tipoCurso]['CodigoCursoEMEC'] !== undefined) {
        formattedCurso.tipo = 'ValidadoPeloMEC'
    } else {
        formattedCurso.tipo = 'TramitacaoMEC'
    }
    let habilitacoesArray = infoCourse[tipoCurso]['Habilitacao']
    if (infoCourse[tipoCurso]['Habilitacao'] !== undefined) {
        if (Array.isArray(habilitacoesArray)) {
            for (let i = 0; i < habilitacoesArray.length; i++) {
                let currentHabilitacao = habilitacoesArray[i]
                formattedCurso.habilitacoes.push(formatHabilitacao(currentHabilitacao))
            }
        } else {
            formattedCurso.habilitacoes.push(formatHabilitacao(habilitacoesArray))
        }
    }
    if (infoCourse[tipoCurso]['CodigoCursoEMEC'] !== undefined) {
        formattedCurso['tipo'] = 'ValidadoPeloMEC'
    } else if (infoCourse[tipoCurso]['SemCodigoCursoEMEC'] != undefined) {
        formattedCurso['tipo'] = 'EmTramite'
        formattedCurso['tramitacaoMEC'] = {
            "numeroProcesso": infoCourse[tipoCurso]['SemCodigoCursoEMEC']['NumeroProcesso'],
            "tipoDeProcesso": infoCourse[tipoCurso]['SemCodigoCursoEMEC']['TipoProcesso'],
            "dataCadastro": formatDateType(infoCourse[tipoCurso]['SemCodigoCursoEMEC']['DataCadastro'].slice(0, 10)),
            "dataProtocolo": formatDateType(infoCourse[tipoCurso]['SemCodigoCursoEMEC']['DataProtocolo'].slice(0, 10))
        }
    }

    // Autorizacao
    if (tipoCurso == 'DadosCurso') {
        formattedCurso['autorizacao'] = formatAtoRegulatorio(infoCourse[tipoCurso]['Autorizacao'])
    } else {
        if (infoCourse[tipoCurso]['Autorizacao'] !== undefined) {
            formattedCurso['autorizacao'] = formatAtoRegulatorio(infoCourse[tipoCurso]['Autorizacao'])
        }
    }

    // Reconhecimento
    if (tipoCurso == 'DadosCurso') {
        formattedCurso['reconhecimento'] = formatAtoRegulatorio(infoCourse[tipoCurso]['Reconhecimento'])
    } else {
        if (infoCourse[tipoCurso]['Reconhecimento'] !== undefined) {
            formattedCurso['reconhecimento'] = formatAtoRegulatorio(infoCourse[tipoCurso]['Reconhecimento'])
        }
    }

    // RenovacaoReconhecimento
    if (infoCourse[tipoCurso]['RenovacaoReconhecimento'] !== undefined) {
        formattedCurso['renovacaoReconhecimento'] = formatAtoRegulatorio(infoCourse[tipoCurso]['RenovacaoReconhecimento'])
    }

    return formattedCurso
}


/** @function
 *  @name formatCurriculo
 *  @param {Object} curriculo - Object with Curriculum data
 *  @returns {Object} - Formatted Curriculum data
 */
function formatCurriculo(curriculo) {
    let tipoCurso = curriculo['CurriculoEscolar']['infCurriculoEscolar']['DadosCurso'] !== undefined ? 'DadosCurso' : 'DadosCursoNSF'

    let unidadeArray =
        curriculo['CurriculoEscolar']['infCurriculoEscolar']['infEstruturaCurricular']['UnidadeCurricular']
    
    let formattedCurriculo = {
        'unidadesCurriculares': [],
        'cargaHorariaCurso': {
            'cargaHoraria': 72,
            'horaAula': true
        },
        'codigoCurriculo': curriculo['CurriculoEscolar']['infCurriculoEscolar']['CodigoCurriculo'],
        'dadosDoCurso': {
            'nome': "",
            'idCurso': "",
            'codigoIES': ""
        },
        'iesEmissora': curriculo['CurriculoEscolar']['infCurriculoEscolar']['IesEmissora']['CodigoMEC'],
        'minutosHoraAula': Number(curriculo['CurriculoEscolar']['infCurriculoEscolar']['MinutosRelogioDaHoraAula']),
        'ambiente': 'Producao',
        'criteriosIntegralizacao': [],
        'categoriasAtividadesComplementares': [],
        'etiquetas': [],
        'areas': [],
        'dataCurriculo': '',
        'nomeParaAreas': '',
        'informacoesAdicionais': '',
        'segurancaCurriculo': curriculo['CurriculoEscolar']['infCurriculoEscolar']['SegurancaCurriculo']['CodigoValidacao']
    }

    // Etiquetas
    let infEtiquetas = curriculo['CurriculoEscolar']['infCurriculoEscolar']['infEtiquetas']['Etiqueta']
    if (infEtiquetas !== undefined) {
        if (Array.isArray(infEtiquetas)) {
            infEtiquetas.forEach(element => {
                formattedCurriculo['etiquetas'].push(formatEtiqueta(element))
            });
        } else {    
            formattedCurriculo.etiquetas.push(formatEtiqueta(infEtiquetas))
        }
    }

    // Nome Para Areas
    if (curriculo['CurriculoEscolar']['infCurriculoEscolar']['NomeParaAreas'] !== undefined) {
        let nomeParaAreas = curriculo['CurriculoEscolar']['infCurriculoEscolar']['NomeParaAreas']
        formattedCurriculo['nomeParaAreas'] = nomeParaAreas
    }

    // Data Curriculo
    if (curriculo['CurriculoEscolar']['infCurriculoEscolar']['DataCurriculo'] !== undefined) {
        let dataCurriculo = curriculo['CurriculoEscolar']['infCurriculoEscolar']['DataCurriculo']
        formattedCurriculo['dataCurriculo'] = formatDateType(dataCurriculo)
    }

    // Areas
    let infAreas = curriculo['CurriculoEscolar']['infCurriculoEscolar']['infAreas']['Area']
    if (infAreas !== undefined) {
        if (Array.isArray(infAreas)) {
            infAreas.forEach(element => {
                formattedCurriculo['areas'].push(formatArea(element))
            });
        } else {
            formattedCurriculo['areas'].push(formatArea(infAreas))
        }
    }

    // Course ID 
    let courseRef = curriculo['CurriculoEscolar']['infCurriculoEscolar'][tipoCurso]
    if (courseRef['CodigoCursoEMEC'] !== undefined) {
        formattedCurriculo['dadosDoCurso']['idCurso'] = courseRef['CodigoCursoEMEC']
    } else {
        formattedCurriculo['dadosDoCurso']['idCurso'] = courseRef['SemCodigoCursoEMEC']['NumeroProcesso']
    }

    // Course Name
    formattedCurriculo['dadosDoCurso']['nome'] = courseRef['NomeCurso']

    // Codigo IES
    formattedCurriculo['dadosDoCurso']['codigoIES'] = curriculo['CurriculoEscolar']['infCurriculoEscolar']['IesEmissora']['CodigoMEC']

    // Producao
    if (curriculo['CurriculoEscolar']['infCurriculoEscolar']['ambiente'] !== 'Producao') {
        formattedCurriculo['ambiente'] = curriculo['CurriculoEscolar']['infCurriculoEscolar']['ambiente']
    }

    // Versão
    if (curriculo['CurriculoEscolar']['infCurriculoEscolar']['versao'] !== undefined) {
        formattedCurriculo['versao'] = 'v'+curriculo['CurriculoEscolar']['infCurriculoEscolar']['versao']
    }

    // Unidades Curriculares
    if (unidadeArray !== undefined) {
        if (Array.isArray(unidadeArray)) {
            for (let i = 0; i < unidadeArray.length; i++) {
                let currentUnidade = unidadeArray[i];
                formattedCurriculo.unidadesCurriculares.push(formatUnidadeCurricular(currentUnidade))
            }
        } else {
     
            formattedCurriculo.unidadesCurriculares.push(formatUnidadeCurricular(unidadeArray))
        }
    }

    // Critérios Integralizacao
    let criteriosRef = curriculo['CurriculoEscolar']['infCurriculoEscolar']['infCriteriosIntegralizacao']
    let criteriosArray = criteriosRef['CriterioIntegralizacaoRotulos']
    if (criteriosArray !== undefined) {
        
        if (Array.isArray(criteriosArray)) {
            for (let i = 0; i < criteriosArray.length; i++) {
                let currentCriterio = criteriosArray[i]
                formattedCurriculo.criteriosIntegralizacao.push(formatCriterio(currentCriterio))
            }
        } else {

            // Only one criterio
            formattedCurriculo.criteriosIntegralizacao.push(formatCriterio(criteriosArray))
        }
    }

    criteriosArray = criteriosRef['CriterioIntegralizacaoExpressao']
    if (criteriosArray !== undefined) {
        
        if (Array.isArray(criteriosArray)) {
            for (let i = 0; i < criteriosArray.length; i++) {
                let currentCriterio = criteriosArray[i]
                formattedCurriculo.criteriosIntegralizacao.push(formatCriterio(currentCriterio))
            }
        } else {

            // Only one criterio
            formattedCurriculo.criteriosIntegralizacao.push(formatCriterio(criteriosArray))
        }
    }

    // Atividades Complementares
    let atividadesRef = curriculo['CurriculoEscolar']['infCurriculoEscolar']['infEstruturaAtividadesComplementares']
    if (atividadesRef !== undefined) {

        // Categorias
        if (atividadesRef['Categoria'] !== undefined) {
            let categoriasArray = atividadesRef['Categoria']
            if (Array.isArray(categoriasArray)) {
                for (let i = 0; i < categoriasArray.length; i++) {
                    let currentCategoria = categoriasArray[i]
                    let categoriaFormatted = formatCategoria(currentCategoria)
                    
                    // Atividades
                    let atividadesArray = currentCategoria['Atividades']['Atividade']
                    categoriaFormatted.atividadesComplementares = formatAtividadesComplementares(atividadesArray)
                    formattedCurriculo.categoriasAtividadesComplementares.push(categoriaFormatted)
                }
            } else {
                let categoriaFormatted = formatCategoria(categoriasArray)

                // Atividades
                let atividadesArray = categoriasArray['Atividades']['Atividade']
                categoriaFormatted.atividadesComplementares = formatAtividadesComplementares(atividadesArray)
                formattedCurriculo.categoriasAtividadesComplementares.push(categoriaFormatted)
            }
        }
    }

    // Informacoes Adicionais
    let informacoesAdicionais = curriculo['CurriculoEscolar']['infCurriculoEscolar']['InformacoesAdicionais']
    if (informacoesAdicionais !== undefined) {
        formattedCurriculo['informacoesAdicionais'] = informacoesAdicionais
    }

    return formattedCurriculo
}

/** @function
 *  @name formatEndereco
 *  @param {Object} endereco - Unformatted Endereco data
 *  @returns {Object} - Standardized Endereco for the chaincode
 */
function formatEndereco(endereco) {
    let formattedEndereco = {
        'logradouro': endereco['Logradouro'] !== undefined ? endereco['Logradouro'] : '',
        'numero': endereco['Numero'] !== undefined ? endereco['Numero'] : '',
        'complemento': endereco['Complemento'] !== undefined ? endereco['Complemento'] : '',
        'bairro': endereco['Bairro'] !== undefined ? endereco['Bairro'] : '',
        'CEP': endereco['CEP'] !== undefined ? endereco['CEP'] : '',
        'municipio': {}
    }

    if (endereco['NomeMunicipioEstrangeiro'] !== undefined) {
        formattedEndereco['municipio'] = {
            'codigoMunicipio': '',
            'nomeMunicipio': '',
            'UF': '',
            'ehEstrangeiro': true,
            'nomeMunicipioEstrangeiro': endereco['NomeMunicipioEstrangeiro']
        }
    } else {
        formattedEndereco['municipio'] = {
            'codigoMunicipio': endereco['CodigoMunicipio'] !== undefined ? endereco['CodigoMunicipio'] : '',
            'nomeMunicipio': endereco['NomeMunicipio'] !== undefined ? endereco['NomeMunicipio'] : '',
            'UF': endereco['UF'] !== undefined ? endereco['UF'] : '',
            'ehEstrangeiro': false,
            'nomeMunicipioEstrangeiro': ''
        }
    }
    return formattedEndereco
}

/** @function
 *  @name formatTramitacaoMEC
 *  @param {Object} tramitacaoMEC - Unformatted TramitacaoMEC data
 *  @returns {Object} - Standardized TramitacaoMEC for the chaincode
 */
function formatTramitacaoMEC(tramitacaoMEC) {
    let formattedTramitacaoMEC = {
        "numeroProcesso": tramitacaoMEC['NumeroProcesso'] !== undefined ? tramitacaoMEC['NumeroProcesso'] : '',
        "tipoDeProcesso": tramitacaoMEC['TipoProcesso'] !== undefined ? tramitacaoMEC['TipoProcesso'] : '',
        "dataCadastro": tramitacaoMEC['DataCadastro'] !== undefined ? formatDateType(tramitacaoMEC['DataCadastro']) : '',
        "dataProtocolo": tramitacaoMEC['DataProtocolo'] !== undefined ? formatDateType(tramitacaoMEC['DataProtocolo']) : '',
    }
    return formattedTramitacaoMEC
}

/** @function
 *  @name formatAtoRegulatorio
 *  @param {Object} atoRegulatorio - Unformatted TramitacaoMEC data
 *  @returns {Object} - Standardized TramitacaoMEC for the chaincode
 */
function formatAtoRegulatorio(atoRegulatorio) {
    let formattedAtoRegulatorio = {
        'tramitacaoMEC': {},
        'numero': ''
    }

    if (atoRegulatorio['InformacoesTramitacaoEMEC'] !== undefined) {
        formattedAtoRegulatorio['tramitacaoMEC'] = formatTramitacaoMEC(atoRegulatorio['InformacoesTramitacaoEMEC'])
    } else {
        formattedAtoRegulatorio['tipoAtoRegulatorio'] = atoRegulatorio['Tipo']
        formattedAtoRegulatorio['numero'] = atoRegulatorio['Numero']
        formattedAtoRegulatorio['data'] = formatDateType(atoRegulatorio['Data'])
        formattedAtoRegulatorio['veiculoPublicado'] = atoRegulatorio['VeiculoPublicado'] !== undefined ? atoRegulatorio['VeiculoPublicado'] : ''
        formattedAtoRegulatorio['dataPublicacao'] = atoRegulatorio['DataPublicacao'] !== undefined ? formatDateType(atoRegulatorio['DataPublicacao']) : ''
        formattedAtoRegulatorio['secaoPublicacao'] = atoRegulatorio['SecaoPublicacao'] !== undefined ? atoRegulatorio['SecaoPublicacao'] : ''
        formattedAtoRegulatorio['paginaPublicacao'] = atoRegulatorio['PaginaPublicacao'] !== undefined ? atoRegulatorio['PaginaPublicacao'] : ''
        formattedAtoRegulatorio['dou'] = atoRegulatorio['NumeroDOU'] !== undefined ? atoRegulatorio['NumeroDOU'] : ''
    }

    return formattedAtoRegulatorio
}

/** @function
 *  @name formatUnidadeCurricular
 *  @param {Object} unidadeCurricular - Unformatted UnidadeCurricular data
 *  @returns {Object} - Standardized UnidadeCurricular for the chaincode
 */
function formatUnidadeCurricular(unidadeCurricular) {
    let formattedUnidadeCurricular = {
        'nome': unidadeCurricular['Nome'],
        'codigo': unidadeCurricular['Codigo'],
        'cargaHorariaEmHoraRelogio': {
            'cargaHoraria': 0,
            'HoraAula': false
        },
        'cargaHorariaEmHoraAula': {
            'cargaHoraria': 0,
            'HoraAula': true
        },
        'ementa': '',
        'fase': '',
        'tipoUnidadeCurricular': unidadeCurricular['Tipo'],
        'areas': [],
        'equivalencias': [],
        'etiquetas': [],
        'preRequisitos': [],
    }
    
    if (unidadeCurricular['CargaHorariaEmHoraRelogio'] != undefined) {
        formattedUnidadeCurricular['cargaHorariaEmHoraRelogio']['cargaHoraria'] = Number(unidadeCurricular['CargaHorariaEmHoraRelogio'])
    }

    if (unidadeCurricular['CargaHorariaEmHoraAula'] != undefined) {
        formattedUnidadeCurricular['cargaHorariaEmHoraAula']['cargaHoraria'] = Number(unidadeCurricular['CargaHorariaEmHoraAula'])
    }
    
    // Pre-requisitos
    if (unidadeCurricular['PreRequisitos'] !== undefined) {
        let requisitos = unidadeCurricular['PreRequisitos']['CodigoDependencia']

        if (Array.isArray(requisitos)) {
            requisitos.forEach(requisito => {
                formattedUnidadeCurricular['preRequisitos'].push(requisito)
            })
        } else {
            formattedUnidadeCurricular['preRequisitos'].push(requisitos)
        }
    }

    // Ementa (minoccurs 0)
    if (unidadeCurricular['Ementa'] !== undefined) {
        let arrayOfEmentas = unidadeCurricular['Ementa']['ItemEmenta']
        if (Array.isArray(arrayOfEmentas)) {
            formattedUnidadeCurricular['ementa'] = arrayOfEmentas
        } else {
            formattedUnidadeCurricular['ementa'] = [arrayOfEmentas]
        }
    }
    
    // Fase (minoccurs 0)
    if (unidadeCurricular['Fase'] !== undefined) {
        formattedUnidadeCurricular['fase'] = unidadeCurricular['Fase']
    }
    
    // Pattern Areas 
    if (unidadeCurricular['Areas'] !== undefined) {
        let areas = unidadeCurricular['Areas']['Area']
        if (areas !== undefined) {
            if (Array.isArray(areas)) {
                areas.forEach(element => {
                    formattedUnidadeCurricular['areas'].push(formatArea(element))
                });
            } else {
                formattedUnidadeCurricular['areas'].push(formatArea(areas))
            }
        }
    }

    // Pattern Equivalencias
    if (unidadeCurricular['Equivalencias'] !== undefined) {
        let equi = unidadeCurricular['Equivalencias']['UnidadesCurricularesEquivalente']
        if (equi !== undefined) {   
            let arrayOfEquiva = equi['CodigoUnidadeEquivalente']
            if (Array.isArray(arrayOfEquiva)) {
                formattedUnidadeCurricular['equivalencias'] = arrayOfEquiva
            } else {
                formattedUnidadeCurricular['equivalencias'] = [arrayOfEquiva]
            }
        }
    }

    // Pattern etiquetas
    if (unidadeCurricular['Etiquetas'] !== undefined) {
        let etiquetasObj = unidadeCurricular['Etiquetas']['Etiqueta']
        if (Array.isArray(etiquetasObj)) {
            for (let i = 0; i < etiquetasObj.length; i++) {
                let currentEtiqueta = etiquetasObj[i]
                formattedUnidadeCurricular['etiquetas'].push(formatEtiqueta(currentEtiqueta))
            }
        } else {
            formattedUnidadeCurricular['etiquetas'].push(formatEtiqueta(etiquetasObj))
        }
    }

    return formattedUnidadeCurricular
}

/** @function
 *  @name formatEtiqueta
 *  @param {Object} etiqueta - Unformatted Etiqueta data
 *  @returns {Object} - Standardized Etiqueta for the chaincode
 */
function formatEtiqueta(etiqueta) {
    let formattedEtiqueta = {
        "codigo": etiqueta['Codigo'] !== undefined ? etiqueta["Codigo"] : '',
        "nome": etiqueta["Nome"] !== undefined ? etiqueta["Nome"] : '',
        aplicadaAutomaticamente: false
    }
    if (etiqueta['AplicadoAutomaticamenteUnidadesNaoPertencentesAoCurriculo'] !== undefined) {
        if (etiqueta['AplicadoAutomaticamenteUnidadesNaoPertencentesAoCurriculo'] === 'Sim') {
            formattedEtiqueta['aplicadaAutomaticamente'] = true
        } else {
            formattedEtiqueta['aplicadaAutomaticamente'] = false
        }
    }
    return formattedEtiqueta
}

/** @function
 *  @name formatAtvidadesComplementares
 *  @param {[Object]} atividadesComplementares - List of Atividades Complementares
 *  @returns {[Object]} - Array with formatted Atividades Complementares objects
 */
function formatAtividadesComplementares(atividadesComplementares) {
    let arrayAtividades = []

    if (Array.isArray(atividadesComplementares)) {
        for (let i = 0; i < atividadesComplementares.length; i++) {
            let currentAtividade = atividadesComplementares[i]
            arrayAtividades.push(formatAtividadeComplementar(currentAtividade))
        }
    } else {

        // Only one atividade
        arrayAtividades.push(formatAtividadeComplementar(atividadesComplementares))
    }

    return arrayAtividades
}


/** @function
 *  @name formatAtividadeComplementar
 *  @param {Object} atividade - Unformatted Atividade Complementar
 *  @returns {Object} - Formatted Atividade Complementar for the chaincode  
 */

function formatAtividadeComplementar(atividade) {
    let formattedAtividadeComplementar = {
        'codigo': atividade['Codigo'] !== undefined ? atividade['Codigo'] : '',
        'nome': atividade['Nome'] !== undefined ? atividade['Nome'] : '',
        'descricao': atividade['Descricao'] !== undefined ? atividade['Descricao'] : '',
        'limiteCargaHorariaEmHoraRelogio': {}
    }

    // Limite Carga Horaria
    if (atividade['LimiteCargaHorariaEmHoraRelogio'] !== undefined) {
        formattedAtividadeComplementar['limiteCargaHorariaEmHoraRelogio'] = {
            'horaAula': false,
            'cargaHoraria': parseInt(atividade['LimiteCargaHorariaEmHoraRelogio'])
        }
    }

    return formattedAtividadeComplementar
}

/** @function
 *  @name formatCategoria
 *  @param {Object} categoria - Unformatted Categoria
 *  @returns {Object} - Formatted Categoria for the chaincode
 */

function formatCategoria(categoria) {
    let formattedCategoria = {
        'codigo': categoria['Codigo'] !== undefined ? categoria['Codigo'] : '',
        'nome': categoria['Nome'] !== undefined ? categoria['Nome'] : '',
        'atividadesComplementares': [],
    }

    // Limite Carga Horaria
    if (categoria['LimiteCargaHorariaEmHoraRelogio'] !== undefined) {
        formattedCategoria['limiteCargaHoraria'] = {
            'horaAula': false,
            'cargaHoraria': parseInt(categoria['LimiteCargaHorariaEmHoraRelogio'])
        }
    }

    return formattedCategoria
}

/** @function
 *  @name formatArea
 *  @param {Object} area - Unformatted Area
 *  @returns {Object} - Formatted Area for the chaincode
 */

function formatArea (area) {
    let formattedArea = {
        codigo: area['Codigo'] !== undefined ? area['Codigo'] : '',
        nome: area['Nome'] !== undefined ? area['Nome'] : ''
    }
    return formattedArea
}

/** @function
 *  @name formatCategoria
 *  @param {Object} categoria - Unformatted Categoria
 *  @returns {Object} - Formatted Categoria for the chaincode
 */

function formatCriterio(criterio) {
    let formattedCriterio = {
        'codigo': criterio['Codigo'] !== undefined ? criterio['Codigo'] : '',
        'unidadesCurriculares': [],
        'expressao': []
    }

    // Verify type
    if (criterio['Expressao'] !== undefined) {

        // Criterio Expressao
        formattedCriterio['tipo'] = 'CriterioExpressao'
        formattedCriterio['expressao'] = []

        let codSomas = criterio['Expressao']['Soma']['Codigo']
        if (codSomas !== undefined) {
            if (Array.isArray(codSomas)) {
                formattedCriterio['expressao'] = codSomas
            } else {
                formattedCriterio['expressao'].push(codSomas['Codigo'])
            }
        }

        let cargaHoraria = criterio['CargasHorariasCriterio']
        formattedCriterio['cargasHorariasCriterio'] = {
            'cargaHorariaMinima': cargaHoraria['CargaHorariaMinima'] !== undefined ? {
                'horaAula': false,
                'cargaHoraria': parseInt(cargaHoraria['CargaHorariaMinima'])
            } : {},
            'cargaHorariaMaxima': cargaHoraria['CargaHorariaMaxima'] !== undefined ? {
                'horaAula': false,
                'cargaHoraria': parseInt(cargaHoraria['CargaHorariaMaxima']) 
            }: {},
            'cargaHorariaTotal': cargaHoraria['CargaHorariaParaTotal'] !== undefined ? {
                'horaAula': false, 
                'cargaHoraria': parseInt(cargaHoraria['CargaHorariaParaTotal']) 
            }: {}
        }

    } else {

        // Critério Rotulo
        formattedCriterio['tipo'] = 'CriterioRotulo'
    
        let unidadesCurriculares = criterio['UnidadeCurricular'] !== undefined ? criterio['UnidadeCurricular'] : ''
        if (Array.isArray(unidadesCurriculares)) {
            formattedCriterio['unidadesCurriculares'] = unidadesCurriculares
        } else {
            formattedCriterio['unidadesCurriculares'].push(unidadesCurriculares)
        }

        // Etiquetas as String
        if (criterio['Etiqueta'] !== undefined) {
            let infEtiquetas = criterio['Etiqueta']
            formattedCriterio['etiquetas'] = []
            
            if (Array.isArray(infEtiquetas)) {
                
                // Multiple Etiquetas
                infEtiquetas.forEach(element => {
                    formattedCriterio['etiquetas'].push(element)
                });
            } else {

                // Only one etiqueta
                formattedCriterio['etiquetas'].push(infEtiquetas)
            }
        }

        // CargasHorariasCriterio
        let cargaHoraria = criterio['CargasHorariasCriterio']
        formattedCriterio['cargasHorariasCriterio'] = {
            'cargaHorariaMinima': cargaHoraria['CargaHorariaMinima'] !== undefined ? {
                'cargaHoraria': parseInt(cargaHoraria['CargaHorariaMinima']),
                'horaAula': false
            }: {},
            'cargaHorariaMaxima': cargaHoraria['CargaHorariaMaxima'] !== undefined ?  {
                'cargaHoraria': parseInt(cargaHoraria['CargaHorariaMaxima']),
                'horaAula': false
            } : {},
            'cargaHorariaTotal': cargaHoraria['CargaHorariaParaTotal'] !== undefined ? {
                'cargaHoraria': parseInt(cargaHoraria['CargaHorariaParaTotal']),
                'horaAula': false
            }: {}
        }
    }

    return formattedCriterio
}

/** @function
 *  @name formatHabilitacao
 *  @param {Object} habilitacao - Unformatted Habilitacao
 *  @returns {Object} - Formatted Habilitacao for the chaincode
 */
function formatHabilitacao(habilitacao) {
    if (habilitacao === undefined) {
        return {
            'nome': 'NaoFornecido',
            'data': formatDateType('0000-00-00')
        }
    } else {
        return {
            'nome': habilitacao['NomeHabilitacao'],
            'data': formatDateType(habilitacao['DataHabilitacao'].slice(0, 10))
        }
    }
}

/** @function
 *  @name formatMantenedora
 *  @param {Object} mantenedora - Unformatted Mantenedora
 *  @returns {Object} - Formatted Mantenedora for the chaincode
 */
function formatMantenedora(mantenedora) {
    let formattedMantenedora = {
        'razaoSocial': mantenedora['RazaoSocial'] !== undefined? mantenedora['RazaoSocial'] : '',
        'CNPJ': mantenedora['CNPJ'] !== undefined? mantenedora['CNPJ'] : '',
        'endereco': formatEndereco(mantenedora['Endereco']),
    }
    return formattedMantenedora
}

/** @function
 *  @name formatDateType
 *  @param {Object} obj - Unformatted Date
 *  @returns {Object} - Formatted Date
 */
function formatDateType(obj) {
    return `${obj}T00:00:00.000Z`
}

module.exports = {
    formatIES,
    formatCurso,
    formatCurriculo
}