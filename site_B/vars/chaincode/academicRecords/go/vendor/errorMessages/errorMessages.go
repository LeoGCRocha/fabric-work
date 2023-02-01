package errorMessages

import (
	"fmt"
)

var errors = map[string]string{
	"cpfNullError":                 "O campo CPF nao pode ser nulo nesta funcao.",
	"argsLen":                      "A funcao espera %d argumento(s), mas recebeu %d.",
	"invalidStruct":                "Estrutura %s invalida.",
	"undefinedFunction":            "Funcao passada como argumento nao esta definida.",
	"cannotInvokeChaincode":        "Nao foi possivel invocar a chaincode. Causa: %s",
	"alunoDoesNotExist":            "O aluno nao existe.",
	"alunoAlreadyExists":           "O aluno ja existe.",
	"historicoDoesNotExist":        "O historico %s nao existe.",
	"historicoAlreadyExists":       "Um historico com essas informacoes ja existe.",
	"requiredField":                "O campo %s eh obrigatorio.",
	"IESDoesNotExist":              "A IES passada como parametro nao existe.",
	"IESAlreadyExists":             "A IES especificada ja existe.",
	"cursoHasNoCurriculos":         "O curso %s nao possui curriculos.",
	"cursoDoesNotExist":            "O curso %s nao existe.",
	"cursoAlreadyExists":           "O curso ja existe.",
	"curriculoAlreadyExists":       "O curriculo %s ja existe.",
	"cannotMarshal":                "Nao foi possivel transformar os bytes de %s em struct.",
	"cannotUnmarshal":              "Nao foi possivel converter o JSON de %s para struct.",
	"invalidJSONInput":             "O JSON passado como parametro nao e valido. Causa: os seguintes parametros eram esperados: %v",
	"invalidJSONInputValues":       "O JSON de %s não representa a struct corretamente. Causa: %s",
	"generic":                      "Causa: %s",
	"registerDoesNotExist":         "O estudante nao possui um/o registro.",
	"indexOutOfRange":              "O index especificado e maior do que a quantidade de itens.",
	"curriculoDoesNotExist":        "O curriculo %s nao existe.",
	"cannotReadWorldState":         "Nao foi possivel ler o WorldState.",
	"cannotUpdateWorldState":       "Nao foi possivel atualizar o WorldState.",
	"invalidFormaDeAcesso":         "Forma de acesso invalida.",
	"invalidCodigoCursoEMEC":       "CodigoCursoEMEC invalido. Causa: %s",
	"disciplinaIsCurricular":       "Uma disciplina em um elementoHistorico deve ser não curricular, mas a disciplina %s e.",
	"repeatedDisciplinaID":         "Ha duas diciplinas com a mesma identificação.",
	"disciplinaNotInCurriculo":     "A disciplina %s não corresponde a disciplina apontada pelo curriculo.",
	"emptyCodigoValidacao":         "CodigoValidacao não pode ser vazio.",
	"emptyCurriculo":               "Curriculo não pode ser vazio.",
	"emptyArea":                    "Area precisa ter Nome e Codigo",
	"invalidEtiquetaCurriculo":     "Uma o mais etiquetas definidas na unidade não estão predefinidas no curriculo",
	"cannotEncode":                 "Nao foi possivel codificar a estrutura %s.",
	"studentDoesNotExist":          "O estudante nao existe.",
	"studentAlreadyExists":         "O estudante ja existe.",
	"documentNotFound":             "O documento %s nao foi encontrado no estudante.",
	"documentAlreadyExists":        "O documento %s ja existe.",
	"invalidCurriculo":             "O curriculo nao eh valido. Causa: %s",
	"invalidHabilitacao":           "A habilitacao nao eh valida. Causa: %s",
	"invalidMunicipio":             "O municipio nao eh valido. Causa: %s",
	"invalidSituacao":              "A situacao nao eh valida. Causa: %s",
	"invalidAtividadeComplementar": "A atividade complementar nao e valida. Causa: %s",
	"invalidCargaHoraria":          "A carga horaria nao eh valida. Causa: %s",
	"invalidDisciplina":            "A disciplina nao eh valida. Causa: %s",
	"invalidIES":                   "A IES nao eh valida. Causa: %s",
	"invalidEndereco":              "O endereco nao eh valido.",
	"invalidAtoRegulatorio":        "O ato regulatorio (%s) nao eh valido.",
	"invalidAreas":                 "O numero de areas deve ser no minimo 1 e no maximo 1",
	"invalidEtiquetas":             "O numero de etiquetas deve ser no minimo 1 e no maximo 1",
	"invalidAutorizacao":           "A autorização pode ser vazia",
	"invalidCurso":                 "O curso nao e valido. Causa: %s",
	"invalidIESReguladora":         "IESReguladora nao e valido. causa: %s",
	"IESReguladoraAlreadyExists":   "A IESReguladora %s ja existe.",
	"invalidCursoReguladora":       "CursoReguladora nao e valido. causa: %s",
	"CursoReguladoraAlreadyExists": "A Reguladora do Curso %s ja existe.",
	// Student chaincode errors
	"InvalidLogin":          "Nao foi possivel validar o estudante %s",
	"InvalidPassword":       "Senha invalida",
	"cannotGetStudentData":  "Nao foi possivel obter os dados do estudante.",
	"cannotRegisterStudent": "Nao foi possivel registrar o estudante.",
	// Go error
	"cannotCreateCompositeKey": "Nao foi possivel criar a chave composta.",
}

func Get(context string, errorID string, args ...interface{}) string {

	var message string

	message, found := errors[errorID]
	if !found {
		message = context + " | " + fmt.Sprintf("ErrorID: %s | ErrorID não encontrado", errorID)
		return message
	}

	message = context + " | " + fmt.Sprintf(message, args...)

	return message
}
