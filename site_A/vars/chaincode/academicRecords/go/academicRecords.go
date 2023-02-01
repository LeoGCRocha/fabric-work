package main

import (
	"assetlib"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	errors "errorMessages"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

/*
#######################
### AcademicRecords ###
#######################
*/

// AcademicRecords struct.
type AcademicRecords struct {
}

/*
###################################
### Chaincode Default Functions ###
###################################
*/

// Init function.
func (t *AcademicRecords) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke function.
func (t *AcademicRecords) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	errContext := "Nao foi possível invocar a funcao."
	fn, args := stub.GetFunctionAndParameters()

	switch fn {
	case "ReadAluno":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		aluno, err := readAluno(stub, args[0])
		if err != nil {
			return shim.Success([]byte("{}"))
		}
		alunoResponseInBytes, err := json.Marshal(aluno)
		if err != nil {
			errorMsg := errors.Get(errContext, "invalidStruct", "aluno")
			return shim.Error(errorMsg)
		}
		return shim.Success(alunoResponseInBytes)
	case "CreateAluno":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		err := createAluno(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(nil))
	case "AlunoExists":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		result, err := alunoExists(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(strconv.FormatBool(result)))
	case "GetAllAlunos":
		if len(args) != 0 {
			errorMsg := errors.Get(errContext, "argsLen", 0, len(args))
			return shim.Error(errorMsg)
		}
		alunos, err := getAllAlunos(stub)
		if err != nil {
			return shim.Success([]byte("[]"))
		}
		if len(alunos) == 0 {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(alunos)
		if err != nil {
			errorMsg := errors.Get(errContext, "invalidStruct", "aluno")
			return shim.Error(errorMsg)
		}
		return shim.Success(result)
	case "ReadAtividadesComplementaresFromAluno":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		atividades, err := readAtividadesComplementaresFromAluno(stub, args[0])
		if err != nil {
			return shim.Error("Nao foi possivel obter as atividades complementares.")
		}
		atividadesJSON, err := json.Marshal(atividades)
		if err != nil {
			errorMsg := errors.Get(errContext, "invalidStruct", "atividade")
			return shim.Error(errorMsg)
		}
		return shim.Success(atividadesJSON)
	case "ReadEstagiosFromAluno":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		estagios, err := readEstagiosFromAluno(stub, args[0])
		if err != nil {
			return shim.Error("Nao foi possivel obter os estagios.")
		}
		estagiosJSON, err := json.Marshal(estagios)
		if err != nil {
			errorMsg := errors.Get(errContext, "invalidStruct", "estagio")
			return shim.Error(errorMsg)
		}
		return shim.Success(estagiosJSON)
	case "GetCategoriasAtividades":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		categorias, err := getCategoriasAtividades(stub, args[0])
		if err != nil {
			return shim.Error("Nao foi possivel obter as categorias")
		}
		categoriasJSON, err := json.Marshal(categorias)
		if err != nil {
			errorMsg := errors.Get(errContext, "invalidStruct", "categoria")
			return shim.Error(errorMsg)
		}
		return shim.Success(categoriasJSON)
	case "UpdateAluno":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		err := updateAluno(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		var buffer bytes.Buffer
		buffer.WriteString("Aluno foi atualizado com sucesso")
		return shim.Success(buffer.Bytes())
	case "CreateHistoricoEscolar":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		err := createHistoricoEscolar(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(nil))
	case "ReadHistorico":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		historico, err := readHistoricoEscolar(stub, args[0])
		if err != nil {
			return shim.Success([]byte("{}"))
		}
		responseBytes, err := json.Marshal(*historico)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Historico")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	case "ReadHistoricosFromAluno":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		historicos, err := readHistoricosFromAluno(stub, args[0])
		if err != nil {
			return shim.Success([]byte("[]"))
		}
		if len(historicos) == 0 {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(historicos)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Historico")
			return shim.Error(errorMsg)
		}
		return shim.Success(result)
	case "ReadLastHistoricoFromAluno":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		historicos, err := readHistoricosFromAluno(stub, args[0])
		if err != nil {
			return shim.Success([]byte("{}"))
		}
		if len(historicos) == 0 {
			return shim.Success([]byte("{}"))
		}
		result, err := json.Marshal(historicos[len(historicos)-1])
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Historico")
			return shim.Error(errorMsg)
		}
		return shim.Success(result)
	case "ReadHistoricoFromCurriculo":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		historicos, err := readHistoricoEscolarFromCurriculo(stub, args[0])
		if err != nil {
			return shim.Success([]byte("[]"))
		}
		if len(historicos) == 0 {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(historicos)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Historico")
			return shim.Error(errorMsg)
		}
		return shim.Success(result)
	case "GetAtividadesPendente":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "1")
			return shim.Error(errorMsg)
		}
		atividades, err := getAtividadesPendente(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		if len(atividades) == 0 {
			return shim.Success([]byte("[]"))
		}
		responseBytes, err := json.Marshal(atividades)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Atividades")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	case "GetEstagiosPendente":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "1")
			return shim.Error(errorMsg)
		}
		estagios, err := getEstagiosPendente(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		if len(estagios) == 0 {
			return shim.Success([]byte("[]"))
		}
		responseBytes, err := json.Marshal(estagios)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Estagios")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	case "ReadHistoricoFromIES":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		historicos, err := readHistoricoFromIes(stub, args[0])
		if err != nil {
			return shim.Success([]byte("[]"))
		}
		if len(historicos) == 0 {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(historicos)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Historico")
			return shim.Error(errorMsg)
		}
		return shim.Success(result)
	case "ReadHistoricoFromCurso":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		historicos, err := readHistoricoEscolarFromCurso(stub, args[0])
		if err != nil {
			return shim.Success([]byte("[]"))
		}
		if len(historicos) == 0 {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(historicos)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Historico")
			return shim.Error(errorMsg)
		}
		return shim.Success(result)
	case "ApproveAtividade":
		if len(args) != 2 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "2")
			return shim.Error(errorMsg)
		}
		err := approveAtividade(stub, args[0], args[1])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(nil)
	case "ApproveEstagio":
		if len(args) != 2 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "2")
			return shim.Error(errorMsg)
		}
		err := approveEstagio(stub, args[0], args[1])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(nil)
	default:
		errorMsg := errors.Get(errContext, "undefinedFunction", fn)
		return shim.Error(errorMsg)
	}
}

func main() {
	if err := shim.Start(new(AcademicRecords)); err != nil {
		log.Panicf("Erro ao tentar iniciar a chaincode academicRecords: %v", err)
	}
}

/*
#######################
### Aluno Functions ###
#######################
*/

// readAluno Check if the aluno given as parameter exists.
// It returns nil if aluno doesn't exists, and a error if any problem occurred.
func readAluno(stub shim.ChaincodeStubInterface, alunoCPF string) (*assetlib.Aluno, error) {

	errContext := "Não foi possível ler o aluno"
	exists, _ := alunoExists(stub, alunoCPF)
	if !exists {
		errorMsg := errors.Get(errContext, "alunoDoesNotExist")
		return nil, fmt.Errorf(errorMsg)
	}

	assetType := "Aluno"
	alunoResposneKey, _ := stub.CreateCompositeKey(assetType, []string{alunoCPF})
	bytesAluno, err := stub.GetState(alunoResposneKey)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Aluno")
		return nil, fmt.Errorf(errorMsg)
	}

	var aluno assetlib.Aluno
	err = json.Unmarshal(bytesAluno, &aluno)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Aluno")
		return nil, fmt.Errorf(errorMsg)
	}

	return &aluno, nil
}

// alunoExists Check if the aluno given as parameter exists.
// It returns false if aluno doesn't exists, and an error if any problem occurred.
func alunoExists(stub shim.ChaincodeStubInterface, alunoCPF string) (bool, error) {

	errContext := "Não foi possível verificar se o aluno existe"
	assetType := "Aluno"
	alunoResponseKey, _ := stub.CreateCompositeKey(assetType, []string{alunoCPF})
	bytesAluno, err := stub.GetState(alunoResponseKey)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotReadWorldState")
		return false, fmt.Errorf(errorMsg)
	}

	return bytesAluno != nil, nil
}

// getAllAlunos Get all alunos stored in the ledger.
// It returns an array of Aluno structs, and an error if any problem occurred.
func getAllAlunos(stub shim.ChaincodeStubInterface) ([]assetlib.Aluno, error) {

	errContext := "Não foi possível recuperar todos os alunos da ledger"
	assetType := "Aluno"
	studentsIterator, err := stub.GetStateByPartialCompositeKey(assetType, nil)
	if err != nil {
		return nil, err
	}

	var alunos []assetlib.Aluno
	defer studentsIterator.Close()
	for studentsIterator.HasNext() {
		queryResponse, err := studentsIterator.Next()
		if err != nil {
			return nil, err
		}
		var aluno assetlib.Aluno
		err = json.Unmarshal(queryResponse.Value, &aluno)
		if err != nil {
			return nil, fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "Aluno"))
		}
		alunos = append(alunos, aluno)
	}

	_, err = json.Marshal(alunos)
	if err != nil {
		return nil, fmt.Errorf(errors.Get(errContext, "cannotMarshal", "Alunos"))
	}

	return alunos, nil
}

// updateAluno Update Aluno on the ledger.
// It Returns an error if any problem occurred.
func updateAluno(stub shim.ChaincodeStubInterface, alunoJSON string) error {

	errContext := "Não foi possível atualizar o aluno"
	var aluno assetlib.Aluno
	parameters := []string{"ID", "CPF", "nome", "sexo", "nacionalidade", "naturalidade", "RG", "dataNascimento", "outroDocumento", "documentoRG", "historicos", "nomeSocial"}
	if !isValidJSONInput(parameters, alunoJSON) {
		errorMsg := errors.Get(errContext, "invalidJSONInput", parameters)
		return fmt.Errorf(errorMsg)
	}

	err := json.Unmarshal([]byte(alunoJSON), &aluno)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Aluno")
		return fmt.Errorf(errorMsg)
	}

	exists, _ := alunoExists(stub, aluno.CPF)
	if !exists {
		errorMsg := errors.Get(errContext, "alunoDoesNotExist")
		return fmt.Errorf(errorMsg)
	}

	var alunoOnLedger *assetlib.Aluno
	alunoOnLedger, err = readAluno(stub, aluno.CPF)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// Verify informations.
	if aluno.DocumentoRG {
		if aluno.RG.Numero == "" || aluno.RG.OrgaoExpedidor == "" || aluno.RG.UF == "" {
			errorMsg := errors.Get(errContext, "invalidJSONInputValues", "aluno", "a struct RG não foi preenchido de maneira adequada")
			return fmt.Errorf(errorMsg)
		}
		if aluno.OutroDocumento.Identificador != "" || aluno.OutroDocumento.TipoDocumento != "" {
			errorMsg := errors.Get(errContext, "invalidJSONInputValues", "aluno", "o campo outroDocumento não pode ser preenchido caso o valor de DocumentoRG seja true")
			return fmt.Errorf(errorMsg)
		}

	} else {
		if aluno.OutroDocumento.Identificador == "" || aluno.OutroDocumento.TipoDocumento == "" {
			errorMsg := errors.Get(errContext, "invalidJSONInputValues", "aluno", "a struct outroDocumento não foi preenchido corretamente")
			return fmt.Errorf(errorMsg)
		}
		if aluno.RG.Numero != "" || aluno.RG.OrgaoExpedidor != "" || aluno.RG.UF != "" {
			errorMsg := errors.Get(errContext, "invalidJSONInputValues", "aluno", "o campo RG  não deve ser preenchido caso o Documento.RG esteja em false")
			return fmt.Errorf(errorMsg)
		}
	}

	// Update informations.
	if aluno.Nome != alunoOnLedger.Nome {
		alunoOnLedger.Nome = aluno.Nome
	}
	if aluno.Sexo != alunoOnLedger.Sexo {
		alunoOnLedger.Sexo = aluno.Sexo
	}
	if aluno.Nacionalidade != alunoOnLedger.Nacionalidade {
		alunoOnLedger.Nacionalidade = aluno.Nacionalidade
	}
	if aluno.Naturalidade != alunoOnLedger.Naturalidade {
		alunoOnLedger.Naturalidade = aluno.Naturalidade
	}
	if aluno.DataNascimento != alunoOnLedger.DataNascimento {
		alunoOnLedger.DataNascimento = aluno.DataNascimento
	}
	if aluno.RG != alunoOnLedger.RG {
		alunoOnLedger.RG = aluno.RG
	}
	if aluno.OutroDocumento != alunoOnLedger.OutroDocumento {
		alunoOnLedger.OutroDocumento = aluno.OutroDocumento
	}
	if aluno.NomeSocial != alunoOnLedger.NomeSocial {
		alunoOnLedger.NomeSocial = aluno.NomeSocial
	}

	// Update on ledger.
	alunoUpdateJSON, err := json.Marshal(alunoOnLedger)
	if err != nil {
		errorMsg := errors.Get(errContext, "invalidStruct", "aluno")
		return fmt.Errorf(errorMsg)
	}

	assetType := "Aluno"
	alunoResponseKey, _ := stub.CreateCompositeKey(assetType, []string{aluno.CPF})

	return stub.PutState(alunoResponseKey, []byte(alunoUpdateJSON))
}

// createAluno Create a new Aluno on the ledger if the json is valid.
// It returns an error if any problem ocurred.
func createAluno(stub shim.ChaincodeStubInterface, alunoJSON string) error {
	errContext := "Nao foi possivel criar o aluno"
	var aluno assetlib.Aluno
	parameters := []string{"ID", "CPF", "nome", "sexo", "nacionalidade", "naturalidade", "RG", "dataNascimento", "outroDocumento", "documentoRG", "historicos", "nomeSocial"}
	if !isValidJSONInput(parameters, alunoJSON) {
		errorMsg := errors.Get(errContext, "invalidJSONInput", parameters)
		return fmt.Errorf(errorMsg)
	}
	err := json.Unmarshal([]byte(alunoJSON), &aluno)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotMarshal", "Aluno")
		return fmt.Errorf(errorMsg)
	}
	exists, err := alunoExists(stub, aluno.CPF)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if exists {
		errorMsg := errors.Get(errContext, "alunoAlreadyExists")
		return fmt.Errorf(errorMsg)
	}
	if aluno.DocumentoRG {
		if aluno.RG.Numero == "" || aluno.RG.OrgaoExpedidor == "" || aluno.RG.UF == "" {
			errorMsg := errors.Get(errContext, "invalidJSONInputValues", "aluno", "a struct RG não foi preenchido de maneira adequada")
			return fmt.Errorf(errorMsg)
		}
		if aluno.OutroDocumento.Identificador != "" || aluno.OutroDocumento.TipoDocumento != "" {
			errorMsg := errors.Get(errContext, "invalidJSONInputValues", "aluno", "o campo outroDocumento não pode ser preenchido caso o valor de DocumentoRG seja true")
			return fmt.Errorf(errorMsg)
		}
	} else {
		if aluno.OutroDocumento.Identificador == "" || aluno.OutroDocumento.TipoDocumento == "" {
			errorMsg := errors.Get(errContext, "invalidJSONInputValues", "aluno", "a struct outroDocumento não foi preenchido corretamente")
			return fmt.Errorf(errorMsg)
		}
		if aluno.RG.Numero != "" || aluno.RG.OrgaoExpedidor != "" || aluno.RG.UF != "" {
			errorMsg := errors.Get(errContext, "invalidJSONInputValues", "aluno", "o campo RG  não deve ser preenchido caso o Documento.RG esteja em false")
			return fmt.Errorf(errorMsg)
		}
	}
	// Validate municipio.
	isValidMunicipio, err := isValidMunicipio(aluno.Naturalidade)
	if !isValidMunicipio {
		return err
	}
	// Prepare to put on world state.
	assetType := "Aluno"
	alunoResponseKey, _ := stub.CreateCompositeKey(assetType, []string{aluno.CPF})
	return stub.PutState(alunoResponseKey, []byte(alunoJSON))
}

/*
##################################
### HistoricoEscolar Functions ###
##################################
*/

// readHistoricoEscolarFromCurriculo Get all historicos from a specific Curriculo.
// Return a list of historicos, or a empty list if there is no historicos.
// Return an error if any problem ocurred.
func readHistoricoEscolarFromCurriculo(stub shim.ChaincodeStubInterface, historicoParams string) ([]assetlib.HistoricoEscolar, error) {

	errContext := "Nao foi possivel ler o historico escolar do curriculo"
	parameters := []string{"curso", "curriculo"}
	if !isValidJSONInput(parameters, historicoParams) {
		errorMsg := errors.Get(errContext, "invalidJSONInput", parameters)
		return nil, fmt.Errorf(errorMsg)
	}

	// Parse to struct.
	type HistoricoParam struct {
		Curso     string
		Curriculo string
	}

	var historicoParam HistoricoParam
	json.Unmarshal([]byte(historicoParams), &historicoParam)

	curso, err := readCurso(stub, historicoParam.Curso)
	if err != nil {
		return nil, err
	}

	found := false
	for _, curriculo := range curso.Curriculos {
		if curriculo.CodigoCurriculo == historicoParam.Curriculo {
			found = true
			break
		}
	}
	if !found {
		errorMsg := errors.Get(errContext, "curriculoDoesNotExist", historicoParam.Curriculo)
		return nil, fmt.Errorf(errorMsg)
	}

	// Get all historico.
	assetType := "HistoricoEscolar"
	var historicos []assetlib.HistoricoEscolar
	historicosIterator, err := stub.GetStateByPartialCompositeKey(assetType, []string{curso.CodigoIES, historicoParam.Curso, historicoParam.Curriculo})
	if err != nil {
		errorMsg := errors.Get(errContext, "generic", err.Error())
		return nil, fmt.Errorf(errorMsg)
	}

	defer historicosIterator.Close()
	for historicosIterator.HasNext() {
		query, err := historicosIterator.Next()
		if err != nil {
			errorMsg := errors.Get(errContext, "generic", err.Error())
			return nil, fmt.Errorf(errorMsg)
		}
		var historico assetlib.HistoricoEscolar
		err = json.Unmarshal(query.Value, &historico)
		if err != nil {
			// return nil, err
			errorMsg := errors.Get(errContext, "cannotUnmarshal", "HistoricoEscolar")
			return nil, fmt.Errorf(errorMsg)
		}
		historicos = append(historicos, historico)
	}

	return historicos, nil
}

// readHistoricoFromIes Get all historicos from a specific IES.
// Return a list of historicos, if no historicos found, return an empty list.
// Return an error if any problem ocurred.
func readHistoricoFromIes(stub shim.ChaincodeStubInterface, iesRef string) ([]assetlib.HistoricoEscolar, error) {

	errContext := "Nao foi possivel ler o historico da IES"
	exists, _ := IESExists(stub, iesRef)
	if !exists {
		errorMsg := errors.Get(errContext, "IESDoesNotExist", iesRef)
		return nil, fmt.Errorf(errorMsg)
	}

	assetType := "HistoricoEscolar"
	var historicos []assetlib.HistoricoEscolar
	historicosIterator, _ := stub.GetStateByPartialCompositeKey(assetType, []string{iesRef})

	defer historicosIterator.Close()
	for historicosIterator.HasNext() {
		query, err := historicosIterator.Next()
		if err != nil {
			errorMsg := errors.Get(errContext, "generic", err.Error())
			return nil, fmt.Errorf(errorMsg)
		}
		var historico assetlib.HistoricoEscolar
		err = json.Unmarshal(query.Value, &historico)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotUnmarshal", "HistoricoEscolar")
			return nil, fmt.Errorf(errorMsg)
		}
		historicos = append(historicos, historico)
	}

	return historicos, nil
}

// readHistoricoEscolarFromCurso Get all historicos of a specific Curso.
// Return a list of historicos, if no historicos found, return an empty list.
// Return an error if any problem ocurred.
func readHistoricoEscolarFromCurso(stub shim.ChaincodeStubInterface, cursoRef string) ([]assetlib.HistoricoEscolar, error) {

	errContext := "Nao foi possivel ler todos os historicos do curso"
	exists, _ := cursoExists(stub, cursoRef)
	if !exists {
		msgError := errors.Get(errContext, "cursoDoesNotExist", cursoRef)
		return nil, fmt.Errorf(msgError)
	}

	curso, _ := readCurso(stub, cursoRef)
	var idCurso string
	if curso.CodigoCursoEMEC != "" {
		idCurso = curso.CodigoCursoEMEC
	} else {
		idCurso = curso.TramitacaoMEC.NumeroProcesso
	}

	assetType := "HistoricoEscolar"
	var historicos []assetlib.HistoricoEscolar
	historicosResultIterator, err := stub.GetStateByPartialCompositeKey(assetType, []string{curso.CodigoIES, idCurso})
	if err != nil {
		errorMsg := errors.Get(errContext, "generic", err.Error())
		return nil, fmt.Errorf(errorMsg)
	}

	defer historicosResultIterator.Close()
	for historicosResultIterator.HasNext() {
		query, err := historicosResultIterator.Next()
		if err != nil {
			errorMsg := errors.Get(errContext, "generic", err.Error())
			return nil, fmt.Errorf(errorMsg)
		}
		var historico assetlib.HistoricoEscolar
		err = json.Unmarshal(query.Value, &historico)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotUnmarshal", "HistoricoEscolar")
			return nil, fmt.Errorf(errorMsg)
		}
		historicos = append(historicos, historico)
	}

	return historicos, nil
}

// readHistoricoEscolar Reads a specific historico from the ledger.
// return a JSON with all the historico informations.
func readHistoricoEscolar(stub shim.ChaincodeStubInterface, historicoParams string) (*assetlib.HistoricoEscolar, error) {

	errContext := "Nao foi possivel ler o historico escolar"
	parameters := []string{"iesEmissora", "curso", "curriculo", "aluno", "digestValue"}
	if !isValidJSONInput(parameters, historicoParams) {
		errorMsg := errors.Get(errContext, "invalidJSONInput", parameters)
		return nil, fmt.Errorf(errorMsg)
	}

	// Parse to struct.
	type HistoricoParam struct {
		IesEmissora string
		Curso       string
		Curriculo   string
		Aluno       string
		DigestValue string
	}

	var historicoParam HistoricoParam
	json.Unmarshal([]byte(historicoParams), &historicoParam)
	exists, _ := historicoExists(stub, historicoParam.IesEmissora, historicoParam.Curso, historicoParam.Curriculo, historicoParam.Aluno,
		historicoParam.DigestValue)
	if !exists {
		errorMsg := errors.Get(errContext, "historicoDoesNotExist", "requisitado")
		return nil, fmt.Errorf(errorMsg)
	}

	// Get value from ledger.
	assetsType := "HistoricoEscolar"
	historicoKey, _ := stub.CreateCompositeKey(assetsType, []string{historicoParam.IesEmissora, historicoParam.Curso,
		historicoParam.Curriculo, historicoParam.Aluno, historicoParam.DigestValue})

	historicoResponse, _ := stub.GetState(historicoKey)
	var historico assetlib.HistoricoEscolar
	err := json.Unmarshal(historicoResponse, &historico)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Historico")
		return nil, fmt.Errorf(errorMsg)
	}

	return &historico, nil
}

// readHistoricosFromAluno Read all historicos from a specific Aluno on the ledger.
// Return an array of historicos, or error if any problem ocurred.
func readHistoricosFromAluno(stub shim.ChaincodeStubInterface, cpf string) ([]assetlib.HistoricoEscolar, error) {
	aluno, err := readAluno(stub, cpf)
	if err != nil {
		return nil, err
	}

	return aluno.Historicos, nil
}

// historicoExists, receives a historicoKey as paramets and verify on the ledger if some historic with this key exists.
// It return true if historico exists, and error if any problem occurred.
func historicoExists(stub shim.ChaincodeStubInterface, iesEmissora string, curso string, curriculo string, aluno string, digestValue string) (bool, error) {
	errContext := "Nao foi possivel verificar se o historico existe"
	assetsType := "HistoricoEscolar"
	historicoKey, _ := stub.CreateCompositeKey(assetsType, []string{iesEmissora, curso, curriculo, aluno, digestValue})
	historicoResponse, err := stub.GetState(historicoKey)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotReadWorldState")
		return false, fmt.Errorf(errorMsg)
	}

	return historicoResponse != nil, nil
}

// createHistoricoEscolar receives a JSON to represent the HistoricoEscolar struct and put it on the ledger.
// It return a error if any problem occurred.
func createHistoricoEscolar(stub shim.ChaincodeStubInterface, historicoJSON string) error {
	errContext := "Nao foi possivel criar um historico escolar"
	parameters := []string{"aluno", "curso", "iesEmissora", "dataHoraEmissao", "situacaoAtualDiscente", "ENADE",
		"curriculo", "ingressoCurso", "elementoHistorico", "cargaHorariaCursoIntegralizada", "cargaHorariaCurso",
		"codigoValidacao", "digestValue", "nomeParaAreas", "areas", "informacoesAdicionais"}

	if !isValidJSONInput(parameters, historicoJSON) {
		errorMsg := errors.Get(errContext, "invalidJSONInput", parameters)
		return fmt.Errorf(errorMsg)
	}

	var historicoEscolar assetlib.HistoricoEscolar
	err := json.Unmarshal([]byte(historicoJSON), &historicoEscolar)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "HistoricoEscolar")
		return fmt.Errorf(errorMsg)
	}

	// Historico Curriculo.
	// In 1.05 version of the schema, the field "codigoCurriculo" is required.
	if historicoEscolar.Curriculo == "" {
		errorMsg := errors.Get(errContext, "emptyCurriculo")
		return fmt.Errorf(errorMsg)
	}

	// The value of the variable represents a composite key representing the structure in the ledger.
	assetType := "HistoricoEscolar"
	historicoEscolarKey, _ := stub.CreateCompositeKey(assetType, []string{historicoEscolar.IesEmissora, historicoEscolar.Curso,
		historicoEscolar.Curriculo, historicoEscolar.Aluno, historicoEscolar.DigestValue})
	historicoExists, _ := historicoExists(stub, historicoEscolar.IesEmissora, historicoEscolar.Curso, historicoEscolar.Curriculo, historicoEscolar.Aluno, historicoEscolar.DigestValue)
	if historicoExists {
		errorMsg := errors.Get(errContext, "historicoAlreadyExists")
		return fmt.Errorf(errorMsg)
	}

	aluno, err := readAluno(stub, historicoEscolar.Aluno)
	if err != nil {
		return err
	}

	curso, err := readCurso(stub, historicoEscolar.Curso)
	if err != nil {
		return err
	}

	_, err = readIES(stub, historicoEscolar.IesEmissora)
	if err != nil {
		return err
	}

	validSituacao, err := isValidSituacao(historicoEscolar.SituacaoAtualDiscente)
	if !validSituacao {
		return err
	}

	for _, habilitacao := range historicoEscolar.ENADE.Habilitacoes {
		validHabilitacao, err := isValidHabilitacao(habilitacao)
		if !validHabilitacao {
			return err
		}
	}

	for _, habilitacao := range historicoEscolar.ENADE.NaoHabilitacoes {
		validHabilitacao, err := isValidHabilitacao(habilitacao)
		if !validHabilitacao {
			return err
		}
	}

	for _, habilitacao := range historicoEscolar.ENADE.Irregulares {
		validHabilitacao, err := isValidHabilitacao(habilitacao)
		if !validHabilitacao {
			return err
		}
	}

	formaDeAcessoValida := false
	formasDeAcesso := []assetlib.TipoFormaDeAcesso{assetlib.EPrograma, assetlib.EConvenios, assetlib.EHistoricoEScolar, assetlib.ESisu, assetlib.EVestibular, assetlib.EEntrevista, assetlib.ETransferencia, assetlib.EOutros}
	for _, formaDeAcesso := range formasDeAcesso {
		if formaDeAcesso == historicoEscolar.IngressoCurso.FormaDeAcesso {
			formaDeAcessoValida = true
			break
		}
	}

	if !formaDeAcessoValida {
		errorMsg := errors.Get(errContext, "invalidFormaDeAcesso")
		return fmt.Errorf(errorMsg)
	}

	validCurriculo := false
	var foundCurriculo assetlib.Curriculo
	for _, curriculo := range curso.Curriculos {
		if curriculo.CodigoCurriculo == historicoEscolar.Curriculo {
			foundCurriculo = curriculo
			validCurriculo = true
			break
		}
	}
	if !validCurriculo {
		errorMsg := errors.Get(errContext, "invalidCurriculo", "O curriculo nao faz parte dos curriculos do curso")
		return fmt.Errorf(errorMsg)
	}

	for _, situacao := range historicoEscolar.ElementoHistorico.Situacoes {
		validSituacao, err := isValidSituacao(situacao)
		if !validSituacao {
			return err
		}
	}

	// Validate Atividades Complementares
	isValidAtividadesComplesmentares := true
	for _, atividade := range historicoEscolar.ElementoHistorico.AtividadesComplementares {
		isValidAtividade := false

		for _, categoriasAtividades := range foundCurriculo.CategoriasAtividadesComplementares {
			for _, atividadeCurricular := range categoriasAtividades.AtividadesComplementares {
				if atividadeCurricular.Codigo == atividade.Codigo {
					isValidAtividade = true
					break
				}
			}
		}

		if !isValidAtividade {
			isValidAtividadesComplesmentares = false
			break
		}
	}

	if !isValidAtividadesComplesmentares {
		errorMsg := errors.Get(errContext, "invalidAtividadeComplementar", "A atividade complementar nao faz parte das atividades do curriculo")
		return fmt.Errorf(errorMsg)
	}

	// Historico Escolar Carga Horaria need to me equal to curriculo
	if (historicoEscolar.CargaHorariaCurso.CargaHoraria != foundCurriculo.CargaHorariaCurso.CargaHoraria) &&
		(historicoEscolar.CargaHorariaCurso.HoraAula != foundCurriculo.CargaHorariaCurso.HoraAula) {
		errorMsg := errors.Get(errContext, "invalidCargaHoraria", "Carga horaria do historico escolar nao e igual a do curriculo")
		return fmt.Errorf(errorMsg)
	}

	mapDisciplina := make(map[string]bool)
	for _, disciplina := range historicoEscolar.ElementoHistorico.Disciplinas {
		if disciplina.Curricular {
			errorMsg := errors.Get(errContext, "disciplinaIsCurricular", disciplina.Nome)
			return fmt.Errorf(errorMsg)
		}

		isValid, err := isValidDiscipline(disciplina)
		if isValid {
			str := disciplina.Codigo
			_, presentDisciplina := mapDisciplina[str]
			if presentDisciplina {
				errorMsg := errors.Get(errContext, "repeatedDisciplinaID")
				return fmt.Errorf(errorMsg)
			}

			disciplinaInCurriculo := false
			for _, disciplinaCurriculo := range foundCurriculo.UnidadesCurriculares {
				if (disciplina.Codigo == disciplinaCurriculo.Codigo) && (disciplina.Nome == disciplinaCurriculo.Nome) {
					disciplinaInCurriculo = true

					// valid carga horaria
					if disciplina.CargaHoraria.HoraAula {
						if disciplina.CargaHoraria.CargaHoraria != disciplinaCurriculo.CargaHorariaEmHoraAula.CargaHoraria {
							errorMsg := errors.Get(errContext, "invalidCargaHoraria", disciplina.Nome)
							return fmt.Errorf(errorMsg)
						}
					} else {
						if disciplina.CargaHoraria.CargaHoraria != disciplinaCurriculo.CargaHorariaEmHoraRelogio.CargaHoraria {
							errorMsg := errors.Get(errContext, "invalidCargaHoraria", disciplina.Nome)
							return fmt.Errorf(errorMsg)
						}
					}
					break
				}
			}

			if !disciplinaInCurriculo {
				errorMsg := errors.Get(errContext, "disciplinaNotInCurriculo", disciplina.Nome)
				return fmt.Errorf(errorMsg)
			}

		} else {
			return err
		}
	}

	for _, atividadeComplementar := range historicoEscolar.ElementoHistorico.AtividadesComplementares {
		validAtividadeComplementar, err := isValidAtividadeComplementar(atividadeComplementar)
		if !validAtividadeComplementar {
			return err
		}
	}

	for _, estagio := range historicoEscolar.ElementoHistorico.Estagios {
		validEstagio, err := isValidEstagio(estagio)
		if !validEstagio {
			return err
		}

	}

	// Codigo Validacao.
	if historicoEscolar.CodigoValidacao == "" {
		errorMsg := errors.Get(errContext, "emptyCodigoValidacao")
		return fmt.Errorf(errorMsg)
	}

	// Validate Areas
	for _, area := range historicoEscolar.Areas {
		if area.Nome == "" || area.Codigo == "" {
			erroMsg := errors.Get(errContext, "emptyArea")
			return fmt.Errorf(erroMsg)
		}
	}

	// Put historic on students array.
	aluno.Historicos = append(aluno.Historicos, historicoEscolar)
	assetType2 := "Aluno"
	alunoResponseKey, _ := stub.CreateCompositeKey(assetType2, []string{aluno.CPF})

	alunoJSON, err := json.Marshal(aluno)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotMarshal", "Aluno")
		return fmt.Errorf(errorMsg)
	}

	stub.PutState(alunoResponseKey, []byte(alunoJSON))

	return stub.PutState(historicoEscolarKey, []byte(historicoJSON))
}

/*
##################################
### Student Functions          ###
##################################
*/

// readAtividadesComplementaresFromAluno Read all atividades complementares from a specific Aluno on the ledger.
func readAtividadesComplementaresFromAluno(stub shim.ChaincodeStubInterface, cpf string) ([]assetlib.AtividadeComplementar, error) {
	var atividadesComplemenatares []assetlib.AtividadeComplementar
	aluno, err := readAluno(stub, cpf)
	if err != nil {
		return nil, err
	}
	if len(aluno.Historicos) > 0 {
		lastHistorico := aluno.Historicos[len(aluno.Historicos)-1]
		atividadesComplemenatares = lastHistorico.ElementoHistorico.AtividadesComplementares
		return atividadesComplemenatares, nil
	}
	return []assetlib.AtividadeComplementar{}, nil
}

// readEstagiosFromAluno Read all estagios from a specific Aluno on the ledger.
func readEstagiosFromAluno(stub shim.ChaincodeStubInterface, cpf string) ([]assetlib.Estagio, error) {
	var estagios []assetlib.Estagio
	aluno, err := readAluno(stub, cpf)
	if err != nil {
		return nil, err
	}
	if len(aluno.Historicos) > 0 {
		lastHistorico := aluno.Historicos[len(aluno.Historicos)-1]
		estagios = lastHistorico.ElementoHistorico.Estagios
		return estagios, nil
	}
	return []assetlib.Estagio{}, nil
}

func getAtividadesPendente(stub shim.ChaincodeStubInterface, CPF string) ([]assetlib.AtividadeComplementar, error) {
	response := stub.InvokeChaincode("student", ToChaincodeArgs("GetAtividadesPendente", CPF), "jornada")
	var atividades []assetlib.AtividadeComplementar
	err := json.Unmarshal(response.Payload, &atividades)
	if err != nil {
		return nil, fmt.Errorf("erro ao deserializar atividades")
	}
	return atividades, nil
}

func getEstagiosPendente(stub shim.ChaincodeStubInterface, CPF string) ([]assetlib.Estagio, error) {
	response := stub.InvokeChaincode("student", ToChaincodeArgs("GetEstagiosPendente", CPF), "jornada")
	var estagios []assetlib.Estagio
	err := json.Unmarshal(response.Payload, &estagios)
	if err != nil {
		return nil, fmt.Errorf("erro ao deserializar estagios")
	}
	return estagios, nil
}

// Get list of avaible categories from the last historico escolar.
func getCategoriasAtividades(stub shim.ChaincodeStubInterface, CPF string) ([]assetlib.CategoriaAtividadeComplementar, error) {
	aluno, err := readAluno(stub, CPF)
	if err != nil {
		return nil, err
	}
	lastHistorico := aluno.Historicos[len(aluno.Historicos)-1]
	codigoCurriculo := lastHistorico.Curriculo
	errContext := "Nao foi possivel invocar a chaincode decree"
	response := stub.InvokeChaincode("decree", ToChaincodeArgs("ReadCurriculo", codigoCurriculo), "jornada")
	if response.GetStatus() != shim.OK {
		errorMsg := errors.Get(errContext, "cannotInvokeChaincode", response.GetMessage())
		return nil, fmt.Errorf(errorMsg)
	}
	var curriculo assetlib.Curriculo
	payloadCurriculo := response.GetPayload()
	err = json.Unmarshal(payloadCurriculo, &curriculo)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Curriculo")
		return nil, fmt.Errorf(errorMsg)
	}
	categorias := curriculo.CategoriasAtividadesComplementares
	return categorias, nil
}

func approveAtividade(stub shim.ChaincodeStubInterface, CPF string, UUID string) error {
	aluno, err := readAluno(stub, CPF)
	if err != nil {
		return err
	}
	if len(aluno.Historicos) == 0 {
		return fmt.Errorf("nao ha historicos para o aluno")
	}
	errContext := "Nao foi possivel aprovar a atividade."
	response := stub.InvokeChaincode("student", ToChaincodeArgs("ApproveAndRemoveAtividade", CPF, UUID), "jornada")
	if response.GetStatus() != shim.OK {
		errorMsg := errors.Get(errContext, "cannotInvokeChaincode", response.GetMessage())
		return fmt.Errorf(errorMsg)
	}
	// Add atividade in last historico escolar.
	// Get last historico escolar.
	lastHistorico := aluno.Historicos[len(aluno.Historicos)-1]
	// Get atividade from student chaincode.
	var atividade assetlib.AtividadeComplementar
	err = json.Unmarshal(response.GetPayload(), &atividade)
	if err != nil {
		// errorMsg := errors.Get(errContext, "cannotUnmarshal", "AtividadeComplementar")
		return fmt.Errorf(string(response.GetPayload()))
	}
	// Add atividade in historico escolar.
	lastHistorico.ElementoHistorico.AtividadesComplementares = append(lastHistorico.ElementoHistorico.AtividadesComplementares, atividade)
	// Create new historico escolar
	// Generate digest value for historico escolar
	hasher := sha256.New()
	horaAtual := time.Now()
	lastHistorico.DataHoraEmissao = horaAtual
	hasher.Write([]byte(horaAtual.String()))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	lastHistorico.DigestValue = sha
	// Add historico escolar in aluno
	historicoJSON, _ := json.Marshal(lastHistorico)
	createHistoricoEscolar(stub, string(historicoJSON))
	return nil
}

func approveEstagio(stub shim.ChaincodeStubInterface, CPF string, UUID string) error {
	aluno, err := readAluno(stub, CPF)
	if err != nil {
		return err
	}
	if len(aluno.Historicos) == 0 {
		return fmt.Errorf("nao ha historicos para o aluno")
	}
	errContext := "Nao foi possivel aprovar o estagio."
	response := stub.InvokeChaincode("student", ToChaincodeArgs("ApproveAndRemoveEstagio", CPF, UUID), "jornada")
	if response.GetStatus() != shim.OK {
		errorMsg := errors.Get(errContext, "cannotInvokeChaincode", response.GetMessage())
		return fmt.Errorf(errorMsg)
	}
	// Add estagio in last historico escolar.
	// Get last historico escolar.
	lastHistorico := aluno.Historicos[len(aluno.Historicos)-1]
	// Get estagio from student chaincode.
	var estagio assetlib.Estagio
	err = json.Unmarshal(response.GetPayload(), &estagio)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Estagio")
		return fmt.Errorf(errorMsg)
	}
	// Add estagio in historico escolar.
	lastHistorico.ElementoHistorico.Estagios = append(lastHistorico.ElementoHistorico.Estagios, estagio)
	hasher := sha256.New()
	horaAtual := time.Now()
	lastHistorico.DataHoraEmissao = horaAtual
	hasher.Write([]byte(horaAtual.String()))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	lastHistorico.DigestValue = sha
	// Add historico escolar in aluno
	historicoJSON, _ := json.Marshal(lastHistorico)
	createHistoricoEscolar(stub, string(historicoJSON))
	return nil
}

// Listar as atividades de um estudante
// Listar os estagios de um estudante

/*
#####################
### IES Functions ###
#####################
*/

// IESExists receives a codigoMEC which represent the IES identificator.
// It return true if IES exists, and return a error if any problem occurred.
func IESExists(stub shim.ChaincodeStubInterface, codigoMEC string) (bool, error) {
	errContext := "Nao foi possivel verificar se a IES existe"
	response := stub.InvokeChaincode("decree", ToChaincodeArgs("IESExists", codigoMEC), "jornada")
	if response.GetStatus() != shim.OK {
		errorMsg := errors.Get(errContext, "cannotInvokeChaincode", response.GetMessage())
		return false, fmt.Errorf(errorMsg)
	}
	result, err := strconv.ParseBool(string(response.GetPayload()))
	if err != nil {
		errorMsg := errors.Get(errContext, "parseBoolFailure")
		return false, fmt.Errorf(errorMsg)
	}

	return result, nil
}

// readIES receives the ID identificator to a IES.
// It return a IES if exists, and return a error if any problem occurred.
func readIES(stub shim.ChaincodeStubInterface, iesID string) (*assetlib.IES, error) {
	errContext := "Nao foi possivel ler a IES"
	exists, _ := IESExists(stub, iesID)
	if !exists {
		errorMsg := errors.Get(errContext, "IESDoesNotExist", iesID)
		return nil, fmt.Errorf(errorMsg)
	}
	response := stub.InvokeChaincode("decree", ToChaincodeArgs("ReadIES", iesID), "jornada")
	if response.GetStatus() != shim.OK {
		errorMsg := errors.Get(errContext, "cannotInvokeChaincode", response.GetMessage())
		return nil, fmt.Errorf(errorMsg)
	}
	var ies assetlib.IES
	payloadIes := response.GetPayload()
	err := json.Unmarshal(payloadIes, &ies)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "IES")
		return nil, fmt.Errorf(errorMsg)
	}
	return &ies, nil
}

/*
#######################
### Curso Functions ###
#######################
*/

// readCurso receives a courseID as a parameter, which represents the identifier of a course.
// It return a Curso if exists, and error if any problem occurred.
func readCurso(stub shim.ChaincodeStubInterface, cursoID string) (*assetlib.Curso, error) {

	errContext := "Nao foi possivel ler o curso"
	exists, err := cursoExists(stub, cursoID)
	if err != nil {
		return nil, err
	}
	if !exists {
		errorMsg := errors.Get(errContext, "cursoDoesNotExist", cursoID)
		return nil, fmt.Errorf(errorMsg)
	}

	response := stub.InvokeChaincode("decree", ToChaincodeArgs("ReadCurso", cursoID), "jornada")
	if response.GetStatus() != shim.OK {
		errorMsg := errors.Get(errContext, "cannotInvokeChaincode", response.GetMessage())
		return nil, fmt.Errorf(errorMsg)
	}

	var curso assetlib.Curso
	payloadCurso := response.GetPayload()
	err = json.Unmarshal(payloadCurso, &curso)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Curso")
		return nil, fmt.Errorf(errorMsg)
	}

	return &curso, nil
}

// cursoExists receives the cursoID which represent the Curso identificator.
// It return true if Curso exists, and a error if any problem occurred.
func cursoExists(stub shim.ChaincodeStubInterface, idCurso string) (bool, error) {
	errContext := "Nao foi possivel verificar se o curso existe"
	response := stub.InvokeChaincode("decree", ToChaincodeArgs("CursoExists", idCurso), "jornada")
	if response.GetStatus() != shim.OK {
		errorMsg := errors.Get(errContext, "cannotInvokeChaincode", response.GetMessage())
		return false, fmt.Errorf(errorMsg)
	}
	result, err := strconv.ParseBool(string(response.GetPayload()))
	if err != nil {
		errorMsg := errors.Get(errContext, "parseBoolFailure")
		return false, fmt.Errorf(errorMsg)
	}
	return result, nil
}

/*
############################
### Validation Functions ###
############################
*/

// isValidJSONInput Checks if the JSON given as parameter is valid.
// It return true if the JSON is valid, and an error if any problem ocurred.
func isValidJSONInput(parameters []string, args string) bool {

	var input map[string]interface{}
	err := json.Unmarshal([]byte(args), &input)
	if err != nil {
		return false
	}

	i := 0
	keys := make([]string, len(input))
	for k := range input {
		keys[i] = k
		i++
	}
	if len(keys) != len(parameters) {
		return false
	}
	present := make([]bool, len(parameters))
	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(parameters); j++ {
			if parameters[j] == keys[i] {
				if present[j] {
					return false
				}
				present[j] = true
				break
			}
		}
	}
	for i := 0; i < len(present); i++ {
		if !present[i] {
			return false
		}

	}

	return true
}

// isValidEstagio Checks if the Estagio given as parameter is valid.
// It return true if the Estagio is valid and a error if any problem occurred.
func isValidEstagio(estagio assetlib.Estagio) (bool, error) {

	errContext := "Nao foi possivel validar o estagio"
	if len(estagio.DocentesOrientadores) == 0 {
		errMessage := errors.Get(errContext, "noDefinedDocente")
		return false, fmt.Errorf(errMessage)
	}

	for _, docente := range estagio.DocentesOrientadores {
		if docente.Nome == "" || docente.Titulacao == "" {
			errMessage := errors.Get(errContext, "invalidDocente")
			return false, fmt.Errorf(errMessage)
		}
	}

	if estagio.CargaHorariaEmHorasRelogio.HoraAula {
		errMessage := errors.Get(errContext, "cannotBeHoraAula")
		return false, fmt.Errorf(errMessage)
	}

	isValidCargaHoraria, err := isValidCargaHoraria(estagio.CargaHorariaEmHorasRelogio)
	if err != nil || !isValidCargaHoraria {
		return false, err
	}

	return true, nil
}

// isValidHabilitacao Checks if the Habilitacao given as parameter is valid.
// It return true if the Habilitacao is valid and a error if any problem occurred.
func isValidHabilitacao(habilitacao assetlib.HabilitacaoEnade) (bool, error) {
	errContext := "Nao foi possivel validar a Habilitacao."
	if habilitacao.Condicao != assetlib.EIngressante && habilitacao.Condicao != assetlib.EConcluinte {
		cause := "Condicao deve ser \"Ingressante\" ou \"Concluinte\""
		errMessage := errors.Get(errContext, "invalidHabilitacao", cause)

		return false,
			fmt.Errorf(errMessage)
	}

	if habilitacao.Edicao <= 0 {
		cause := "Edicao deve ser maior que 0."
		errMessage := errors.Get(errContext, "invalidHabilitacao", cause)
		return false,
			fmt.Errorf(errMessage)
	}

	if !habilitacao.Habilitado {
		if habilitacao.Motivo != assetlib.ECicloAvaliativo && habilitacao.Motivo != assetlib.EProjetoPedagogico && habilitacao.OutroMotivo == "" {
			cause := "Pelo menos um motivo deve ser especificado caso, a Habilitacao seja irregular ou não habilitada."
			errMessage := errors.Get(errContext, "invalidHabilitacao", cause)
			return false,
				fmt.Errorf(errMessage)
		}
		if habilitacao.OutroMotivo != "" && habilitacao.Motivo != "" {
			cause := "Somente um dos campos de motivo deve estar preenchido."
			errMessage := errors.Get(errContext, "invalidHabilitacao", cause)
			return false,
				fmt.Errorf(errMessage)
		}
	}

	return true, nil
}

// isValidMunicipio checks if a Municipio is valid.
// Return true if it is Valid.
func isValidMunicipio(municipio assetlib.Municipio) (bool, error) {

	errContext := "Nao foi possivel validar o Municipio."
	if municipio.EhEstrangeiro {
		// Municipio estrangeiro.
		if municipio.NomeMunicipio != "" || municipio.UF != "" || municipio.CodigoMunicipio != "" {
			cause := "NomeMunicipio, UF e CodigoMunicipio não devem ser preenchidos caso o Municipio seja estrangeiro."
			errMessage := errors.Get(errContext, "invalidMunicipio", cause)
			return false,
				fmt.Errorf(errMessage)
		}

	} else {
		// Municipio nacional.
		if municipio.NomeMunicipio == "" || municipio.UF == "" || municipio.CodigoMunicipio == "" {
			cause := "NomeMunicipio, UF e CodigoMunicipio devem ser preenchidos caso o Municipio seja nacional."
			errMessage := errors.Get(errContext, "invalidMunicipio", cause)
			return false, fmt.Errorf(errMessage)
		}
		if municipio.NomeMunicipioEstrangeiro != "" {
			cause := "NomeMunicipioEstrangeiro não deve ser preenchido caso o Municipio seja nacional."
			errMessage := errors.Get(errContext, "invalidMunicipio", cause)
			return false, fmt.Errorf(errMessage)
		}
	}

	return true, nil
}

// isValidSituacao Checks if the Situacao given as parameter is valid.
// It return true if the Situacao is valid and a error if any problem occurred.
func isValidSituacao(situacao assetlib.Situacao) (bool, error) {

	errContext := "Falha ao validar a situacao"
	tiposDeSituacao := []assetlib.TipoDeSituacao{assetlib.EDesistencia, assetlib.ELicenca, assetlib.EMatriculaEmDisciplina, assetlib.EOutraSituacao, assetlib.ETrancamento, assetlib.EFormado, assetlib.EIntercambio, assetlib.EJubilado}
	for _, situacaoAtual := range tiposDeSituacao {
		if situacaoAtual == situacao.TipoDeSituacao {
			if situacaoAtual == assetlib.EIntercambio {
				if situacao.Intercambio.Instituicao == "" || situacao.Intercambio.Pais == "" || situacao.Intercambio.ProgramaIntercambio == "" {
					cause := "Os campos Instituicao, Pais, ProgramaIntercambio não podem estar vazios."
					errMessage := errors.Get(errContext, "invalidSituacao", cause)
					return false, fmt.Errorf(errMessage)
				}
				return true, nil
			}
			return true, nil
		}
	}

	cause := ""
	errMessage := errors.Get(errContext, "invalidSituacao", cause)
	return false, fmt.Errorf(errMessage)
}

// isValidAtividadeComplementar Checks if the AtividadeComplementar given as parameter is valid.
// It return true if the AtividadeComplementar is valid and a error if any problem occurred.
func isValidAtividadeComplementar(atividadeComplementar assetlib.AtividadeComplementar) (bool, error) {

	errContext := "Falha ao validar a AtividadeComplementar"
	if len(atividadeComplementar.DocentesResponsaveisPelaValidacao) == 0 {
		return false, fmt.Errorf(errors.Get(errContext, "invalidAtividadeComplementar", "AtividadeComplementar deve ter pelo menos um Docente responsavel."))
	}

	if len(atividadeComplementar.DocentesResponsaveisPelaValidacao) == 0 {
		return false, fmt.Errorf(errors.Get(errContext, "invalidAtividadeComplementar", "AtividadeComplementar deve ter pelo menos um Docente responsavel."))
	}

	for _, docente := range atividadeComplementar.DocentesResponsaveisPelaValidacao {
		if docente.Nome == "" || docente.Titulacao == "" {
			return false, fmt.Errorf(errors.Get(errContext, "invalidAtividadeComplementar", "Docente deve ter Nome e Titulacao definido."))
		}
	}

	if atividadeComplementar.CargaHorariaEmHoraRelogio.HoraAula {
		return false, fmt.Errorf(errors.Get(errContext, "invalidAtividadeComplementar", "AtividadeComplementar não pode ser do tipo HoraAula."))
	}

	isValidCargaHoraria, err := isValidCargaHoraria(atividadeComplementar.CargaHorariaEmHoraRelogio)
	if err != nil || !isValidCargaHoraria {
		return false, err
	}

	return true, nil
}

// isValidCargaHoraria Checks if the cargaHoraria given as parameter is valid.
// It returns true if cargaHoraria is valid and  error an error if any problem occurred.
func isValidCargaHoraria(cargaHoraria assetlib.CargaHoraria) (bool, error) {

	errContext := "Falha ao validar a CargaHoraria"
	if cargaHoraria.CargaHoraria <= 0 {
		return false, fmt.Errorf(errors.Get(errContext, "invalidCargaHoraria", "CargaHoraria deve ser maior que 0"))
	}

	return true, nil
}

// isValidDiscipline check if the Disciplina  ed as a parameter is valid.
// It returns true if the Disciplina is valid, and return an error if some problem occurred.
func isValidDiscipline(disciplina assetlib.Disciplina) (bool, error) {
	errContext := "Falha ao validar a Disciplina"
	if disciplina.Nome == "" || (disciplina.CargaHoraria == assetlib.CargaHoraria{}) ||
		disciplina.TipoDeNota == "" || disciplina.Codigo == "" {
		return false, fmt.Errorf(errors.Get(errContext, "invalidDisciplina", "os atributos Nome,CargaHoraria,TipoDenota,Codigo,CodigoCUrsoEMEC são obrigatórios para uma disciplina."))
	}

	if disciplina.TipoDeNota != "Nota" && disciplina.TipoDeNota != "Conceito" &&
		disciplina.TipoDeNota != "ConceitoRM" && disciplina.TipoDeNota != "NotaAteCem" &&
		disciplina.TipoDeNota != "ConceitoEspecificoDoCurso" {
		return false, fmt.Errorf(errors.Get(errContext, "invalidDisciplina", "a struct Disciplina passada como parâmetro não esta definida corretamente."))
	}

	isValidCargaHoraria, err := isValidCargaHoraria(disciplina.CargaHoraria)
	if err != nil || !isValidCargaHoraria {
		return false, fmt.Errorf("falha ao validar a Disciplina: %s", &disciplina.Nome)
		return false, err
	}

	if disciplina.Curricular {
		if disciplina.Periodo != "" || len(disciplina.Docentes) > 0 ||
			(disciplina.EstadoDisciplina != assetlib.EstadoDisciplina{}) || disciplina.TipoDeNota == "" {
			return false, fmt.Errorf(errors.Get(errContext, "invalidDisciplina", "algum dos atributos [Periodo, Docentes, EstadoDisciplina, TipoDeNota] não foi definido corretamente."))
		}

	} else {
		if disciplina.Periodo == "" {
			if (disciplina.EstadoDisciplina == assetlib.EstadoDisciplina{}) || (disciplina.CargaHoraria == assetlib.CargaHoraria{}) ||
				disciplina.TipoDeNota == "" || disciplina.Nota == "" {
				if disciplina.TipoDeNota == "Nota" {
					_, err := strconv.ParseFloat(disciplina.Nota, 64)
					if err != nil {
						return false, fmt.Errorf(errors.Get(errContext, "invalidDisciplina", "a Nota definida para disciplina não é valida para seu tipo."))
					}

				} else if disciplina.TipoDeNota == "NotaAteCem" {
					notaInteger, err := strconv.Atoi(disciplina.Nota)
					if err != nil || notaInteger < 0 || notaInteger > 100 {
						return false, fmt.Errorf(errors.Get(errContext, "invalidDisciplina", "a Nota definida para a disciplina deve ser um inteiro >= 0 && <= 100."))
					}

				} else if disciplina.TipoDeNota == "Conceito" {
					grades := []string{"A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-", "D+", "D", "D-", "E+", "E", "E-", "F+", "F", "F-"}
					found := false
					for _, grade := range grades {
						if grade == disciplina.Nota {
							found = true
							break
						}
					}

					if !found {
						return false, fmt.Errorf(errors.Get(errContext, "invalidDisciplina", "O valor definido na Nota não esta definido para Conceito."))
					}

				} else if disciplina.TipoDeNota == "ConceitoRM" {
					grades := []string{"A", "B", "C", "APD", "APP", "APR"}
					found := false
					for _, grade := range grades {
						if grade == disciplina.Nota {
							found = true
							break
						}
					}

					if !found {
						return false, fmt.Errorf(errors.Get(errContext, "invalidDisciplina", "O valor definido na Nota não esta definido para ConceitoRM."))
					}

				} else if disciplina.TipoDeNota == "ConceitoEspecificoDoCurso" {
				} else {
					return false, fmt.Errorf(errors.Get(errContext, "invalidDisciplina", "o valor atribuido para TipoDeNota não esta definido."))
				}
				return false, fmt.Errorf(errors.Get(errContext, "invalidDisciplina", "disciplina do Tipo curricular invalida"))
			}
			return false, fmt.Errorf(errors.Get(errContext, "invalidDisciplina", "disciplina do Tipo não curricular invalida"))
		}
	}

	return true, nil
}

// ToChaincodeArgs receives dynamic number of strings as parameters.
// It returns array byte of chaincode args.
func ToChaincodeArgs(args ...string) [][]byte {

	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}

	return bargs
}
