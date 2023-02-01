'use strict'

function reqParamsToAtividade(req) {
    // return {
    //     'atividade': req.params.atividade,
    //     'CPF': req.params.CPF,
    //     'type': 'ApproveAtividade'
    // }
    return {
        'codigo': req.params.codigo,
    }

    // type AtividadeComplementar struct {
	//     Id                                string            `json:"id"`
	//     Codigo                            string            `json:"codigo"`
	//     DataInicio                        time.Time         `json:"dataInicio"`
	//     DataFim                           time.Time         `json:"dataFim"`
	//     DataRegistro                      time.Time         `json:"dataRegistro"`
	//     TipoAtividadeComplementar         string            `json:"tipoAtividadeComplementar"`
	//     Descricao                         string            `json:"descricao"`
	//     CargaHorariaEmHoraRelogio         CargaHoraria      `json:"cargaHorariaEmHoraRelogio"`
	//     DocentesResponsaveisPelaValidacao []Docente         `json:"docentesResponsaveisPelaValidacao"`
	//     Certificado                       string            `json:"certificado"`
	//     Situacao                          SituacaoAtividade `json:"situacao"`
    // }
}

function reqParamsToEstagio(req) {
    return {}
}

module.exports = {
    reqParamsToAtividade,
    reqParamsToEstagio
}