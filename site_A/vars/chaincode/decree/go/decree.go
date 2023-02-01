package main

import (
	"assetlib"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	errors "errorMessages"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

/*
##############
### Decree ###
##############
*/

// Decree struct.
type Decree struct {
}

/*
###################################
### Chaincode Default Functions ###
###################################
*/

// Init function.
func (t *Decree) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke function.
func (t *Decree) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	errContext := "Não foi possível invocar a função."
	fn, args := stub.GetFunctionAndParameters()

	switch fn {

	case "CursoExists":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 3, len(args))
			return shim.Error(errorMsg)
		}
		result, err := cursoExists(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(strconv.FormatBool(result)))

	case "ReadCurso":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		curso, err := readCurso(stub, args[0])
		if err != nil {
			return shim.Success([]byte("{}"))
		}
		responseBytes, err := json.Marshal(*curso)
		if err != nil {
			return shim.Error(errors.Get(errContext, "cannotMarshal", "Curso"))
		}
		return shim.Success(responseBytes)

	case "ReadCursosFromNome":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		cursos, err := readCursosFromNome(stub, args[0])
		if err != nil {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(cursos)
		if cursos == nil {
			return shim.Success([]byte("[]"))
		}
		if err != nil {
			return shim.Error(errors.Get(errContext, "cannotMarshal", "Cursos"))
		}
		return shim.Success(result)

	case "ReadCursosFromIES":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		cursos, _ := readCursosFromIES(stub, args[0])
		if len(cursos) == 0 {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(cursos)
		if cursos == nil {
			return shim.Success([]byte("[]"))
		}
		if err != nil {
			return shim.Error(errors.Get(errContext, "cannotMarshal", "Cursos"))
		}
		return shim.Success(result)

	case "ReadAllCursos":
		if len(args) != 0 {
			errorMsg := errors.Get(errContext, "argsLen", 0, len(args))
			return shim.Error(errorMsg)
		}
		cursos, err := readAllCursos(stub)
		if err != nil || cursos == nil {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(cursos)
		if err != nil {
			return shim.Error(errors.Get(errContext, "cannotMarshal", "Cursos"))
		}
		return shim.Success(result)

	case "CurriculoExists":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		result, err := curriculoExists(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(strconv.FormatBool(result)))

	case "ReadCurriculo":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		curriculo, err := readCurriculo(stub, args[0])
		if err != nil {
			return shim.Success([]byte("{}"))
		}
		responseBytes, err := json.Marshal(curriculo)
		if err != nil {
			return shim.Error(errors.Get(errContext, "cannotMarshal", "Curriculo"))
		}
		return shim.Success(responseBytes)

	case "ReadCurriculosFromCurso":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		curriculos, err := readCurriculosFromCurso(stub, args[0])
		if err != nil {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(curriculos)
		if curriculos == nil {
			return shim.Success([]byte("[]"))
		}
		if err != nil {
			return shim.Error(errors.Get(errContext, "cannotMarshal", "Curriculos"))
		}
		return shim.Success(result)

	case "ReadCurriculosFromAmbiente":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		curriculos, err := readCurriculosFromAmbiente(stub, args[0])
		if err != nil {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(curriculos)
		if curriculos == nil {
			return shim.Success([]byte("[]"))
		}
		if err != nil {
			return shim.Error(errors.Get(errContext, "cannotMarshal", "Curriculos"))
		}
		return shim.Success(result)

	case "ReadCurriculosFromIES":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		curriculos, err := readCurriculosFromIES(stub, args[0])
		if err != nil {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(curriculos)
		if curriculos == nil {
			return shim.Success([]byte("[]"))
		}
		if err != nil {
			return shim.Error(errors.Get(errContext, "cannotMarshal", "Curriculos"))
		}
		return shim.Success(result)

	case "ReadAllCurriculos":
		if len(args) != 0 {
			errorMsg := errors.Get(errContext, "argsLen", 0, len(args))
			return shim.Error(errorMsg)
		}
		curriculos, err := readAllCurriculos(stub)
		if err != nil {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(curriculos)
		if curriculos == nil {
			return shim.Success([]byte("[]"))
		}
		if err != nil {
			return shim.Error(errors.Get(errContext, "cannotMarshal", "Curriculos"))
		}
		return shim.Success(result)

	case "IntegralizacaoCurricular":
		if len(args) != 2 {
			errorMsg := errors.Get(errContext, "argsLen", 2, len(args))
			return shim.Error(errorMsg)
		}
		result, err := integralizacaoCurricular(stub, args[0], args[1])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(strconv.FormatBool(result)))

	case "IESExists":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		result, err := IESExists(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(strconv.FormatBool(result)))

	case "CreateIES":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		err := createIES(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(nil))

	case "ReadIES":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		ies, err := readIES(stub, args[0])
		if err != nil {
			return shim.Success([]byte("{}"))
		}
		responseBytes, err := json.Marshal(*ies)
		if err != nil {
			return shim.Error(errors.Get(errContext, "cannotMarshal", "IES"))
		}
		return shim.Success(responseBytes)

	case "ReadAllIES":
		if len(args) != 0 {
			errorMsg := errors.Get(errContext, "argsLen", 0, len(args))
			return shim.Error(errorMsg)
		}
		ies, err := readAllIESToAPI(stub)
		if err != nil || ies == nil {
			return shim.Success([]byte("[]"))
		}
		result, err := json.Marshal(ies)
		if err != nil {
			return shim.Error(errors.Get(errContext, "cannotMarshal", "IES"))
		}
		return shim.Success(result)

	case "AddCurso":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		err := addCurso(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(nil))

	case "AddCurriculo":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "argsLen", 1, len(args))
			return shim.Error(errorMsg)
		}
		err := addCurriculo(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(nil))

	default:
		return shim.Error(errors.Get(errContext, "undefinedFunction"))
	}
}

func main() {
	if err := shim.Start(new(Decree)); err != nil {
		log.Panicf("Erro ao tentar iniciar a chaincode decree: %v", err)
	}
}

/*
#######################
### Curso Functions ###
#######################
*/

// readCurso receives a courseID as a parameter, which represents the identifier of a course.
// It returns a Curso if exists, and error if any problem occurred.
func readCurso(stub shim.ChaincodeStubInterface, idCurso string) (*assetlib.Curso, error) {

	errContext := "Nao foi possivel ler o curso."
	exists, err := cursoExists(stub, idCurso)
	if !exists {
		return nil, err
	}
	// Get value from ledger.
	assetsType := "Curso"
	cursoKey, _ := stub.CreateCompositeKey(assetsType, []string{idCurso})
	cursoResponse, _ := stub.GetState(cursoKey)
	var curso assetlib.Curso
	err = json.Unmarshal(cursoResponse, &curso)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Curso")
		return nil, fmt.Errorf(errorMsg)
	}
	return &curso, nil
}

// addCurso receives a JSON which represent a Curso struct and put it on the ledger.
// It returns a error if any problem occurred.
func addCurso(stub shim.ChaincodeStubInterface, cursoJSON string) error {
	errContext := "Nao foi possivel adicionar o curso."
	var curso assetlib.Curso
	parameters := []string{"nome", "habilitacoes", "tipo", "codigoCursoEMEC", "tramitacaoMEC", "codigoIES", "curriculos",
		"autorizacao", "reconhecimento", "renovacaoReconhecimento"}
	if !isValidJSONInput(parameters, cursoJSON) {
		return fmt.Errorf(errors.Get(errContext, "invalidJSONInput", parameters))
	}
	err := json.Unmarshal([]byte(cursoJSON), &curso)
	if err != nil {
		return fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "Curso"))
	}

	ies, err := readIES(stub, curso.CodigoIES)
	if err != nil {
		return err
	}

	cursoID, err := validateCursoAndReturnID(curso)
	if err != nil {
		return err
	}

	cursoExists, err := cursoExists(stub, cursoID)
	if err != nil {
		return err
	}
	if cursoExists {
		return fmt.Errorf(errors.Get(errContext, "cursoAlreadyExists"))
	}

	err = createCurso(stub, curso)
	if err != nil {
		return err
	}

	ies.Cursos = append(ies.Cursos, curso)
	iesJSON, err := json.Marshal(ies)
	if err != nil {
		return fmt.Errorf(errors.Get(errContext, "cannotMarshal", "IES"))
	}

	assetType := "IES"
	IESKey, _ := stub.CreateCompositeKey(assetType, []string{ies.CodigoMEC})

	return stub.PutState(IESKey, iesJSON)
}

// createCurso receives a JSON which represent a Curso struct and put it on the ledger.
// It return true if Curso exists, and a error if any problem occurred.
func createCurso(stub shim.ChaincodeStubInterface, curso assetlib.Curso) error {
	errContext := "Nao foi possivel criar o curso"
	cursoJSON, err := json.Marshal(curso)
	if err != nil {
		return fmt.Errorf(errors.Get(errContext, "cannotMarshal", "Curso"))
	}
	parameters := []string{"nome", "habilitacoes", "tipo", "codigoCursoEMEC", "tramitacaoMEC", "codigoIES", "curriculos",
		"autorizacao", "reconhecimento", "renovacaoReconhecimento"}
	if !isValidJSONInput(parameters, string(cursoJSON)) {
		return fmt.Errorf(errors.Get(errContext, "invalidJSONInput", parameters))
	}
	err = json.Unmarshal(cursoJSON, &curso)
	if err != nil {
		return fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "Curso"))
	}
	idCurso, err := validateCursoAndReturnID(curso)

	for _, curriculo := range curso.Curriculos {

		curriculoExists, err := curriculoExists(stub, curriculo.CodigoCurriculo)
		if err != nil {
			return err
		}
		if curriculoExists {
			return fmt.Errorf(errors.Get(errContext, "curriculoAlreadyExists", curriculo.CodigoCurriculo))
		}

		if curriculo.CodigoCurriculo == "" {
			return fmt.Errorf(errors.Get(errContext, "invalidCurso", "curriculo.CodigoCurriculo não pode estar vazio."))
		}

		if curriculo.DadosDoCurso.IDCurso != idCurso {
			return fmt.Errorf(errors.Get(errContext, "invalidCurso", "codigo apontado pelo curriculo diferente do codigo do curso."))
		}

		if curriculo.IesEmissora != curso.CodigoIES {
			return fmt.Errorf(errors.Get(errContext, "invalidCurso", "codigo IES apontado pelo curriculo diferente do codigo IES apontado pelo curso."))
		}

		if curriculo.DadosDoCurso.CodigoIES != curso.CodigoIES {
			return fmt.Errorf(errors.Get(errContext, "invalidCurso", "codigo IES apontado pelo curriculo diferente do codigo IES apontado pelo curso."))
		}

		if len(curriculo.UnidadesCurriculares) < 1 {
			return fmt.Errorf(errors.Get(errContext, "invalidCurso", "um curriculo deve ter pelo menos uma unidade curricular."))
		}

		valid, err := isValidCurriculo(curriculo)
		if !valid {
			return err
		}

		curriculoJSON, err := json.Marshal(curriculo)
		if err != nil {
			return fmt.Errorf(errors.Get(errContext, "cannotMarshal", "Curriculo"))
		}
		// Create curso.
		assetType := "Curriculo"
		if curriculo.Ambiente != "Producao" && curriculo.Ambiente != "Homologacao" {
			curriculoResponseKey, _ := stub.CreateCompositeKey(assetType, []string{"Producao", curriculo.CodigoCurriculo})
			stub.PutState(curriculoResponseKey, []byte(curriculoJSON))
		} else {
			curriculoResponseKey, _ := stub.CreateCompositeKey(assetType, []string{string(curriculo.Ambiente), curriculo.CodigoCurriculo})
			stub.PutState(curriculoResponseKey, []byte(curriculoJSON))
		}
	}
	if err != nil {
		return err
	}

	assetType := "Curso"
	cursoKey, _ := stub.CreateCompositeKey(assetType, []string{idCurso})
	return stub.PutState(cursoKey, cursoJSON)
}

// cursoExists receives the cursoID which represent the Curso identificator.
// It return true if Curso exists, and a error if any problem occurred.
func cursoExists(stub shim.ChaincodeStubInterface, idCurso string) (bool, error) {
	errContext := "Nao foi possivel verificar se o curso existe"
	assetType := "Curso"
	cursoResponseKey, _ := stub.CreateCompositeKey(assetType, []string{idCurso})
	bytesCurso, err := stub.GetState(cursoResponseKey)
	if err != nil {
		return false, fmt.Errorf(errors.Get(errContext, "cannotReadWorldState"))
	}
	return bytesCurso != nil, nil
}

// readCursosFromNome reads all cursos with the given name in the ledger.
// It return a list of Curso struct with contain all same name Cursos, and an error if any problem ocurred.
func readCursosFromNome(stub shim.ChaincodeStubInterface, nomeCurso string) ([]assetlib.Curso, error) {
	errContext := "Nao foi possivel ler os cursos"
	assetType := "Curso"
	var cursos []assetlib.Curso

	cursosIterator, _ := stub.GetStateByPartialCompositeKey(assetType, nil)
	defer cursosIterator.Close()
	for cursosIterator.HasNext() {
		query, err := cursosIterator.Next()
		if err != nil {
			return nil, err
		}
		var curso assetlib.Curso
		err = json.Unmarshal(query.Value, &curso)
		if err != nil {
			return nil, fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "Curso"))
		}
		// Return 0 if curso.Nome equals to nomeCurso
		if strings.Compare(curso.Nome, nomeCurso) == 0 {
			cursos = append(cursos, curso)
		}
	}
	return cursos, nil
}

// readCursosFromIES reads all cursos from the given IES in the ledger.
// It returns a list of struct Curso, and an error if any problem ocurred.
func readCursosFromIES(stub shim.ChaincodeStubInterface, codigoIES string) ([]assetlib.Curso, error) {
	// errContext := "Nao foi possivel ler os cursos da IES"
	ies, err := readIES(stub, codigoIES)
	if err != nil {
		return nil, err
	}

	return ies.Cursos, nil
}

// readAllCursos reads all cursos from the ledger.
// It returns a list of struct Curso, and an error if any problem ocurred.
func readAllCursos(stub shim.ChaincodeStubInterface) ([]assetlib.Curso, error) {
	errContext := "Nao foi possivel ler todos os cursos"
	assetType := "Curso"
	cursoIterator, err := stub.GetStateByPartialCompositeKey(assetType, nil)
	if err != nil {
		return nil, err
	}
	var cursos []assetlib.Curso
	defer cursoIterator.Close()
	for cursoIterator.HasNext() {
		queryResponse, err := cursoIterator.Next()
		if err != nil {
			return nil, err
		}
		var curso assetlib.Curso
		err = json.Unmarshal(queryResponse.Value, &curso)
		if err != nil {
			return nil, fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "Curso"))
		}
		cursos = append(cursos, curso)
	}
	_, err = json.Marshal(cursos)
	if err != nil {
		return nil, fmt.Errorf(errors.Get(errContext, "cannotMarshal", "Curso"))
	}
	return cursos, nil
}

/*
###########################
### Curriculo Functions ###
###########################
*/

// addCurriculo receives a JSON which represent a Currriculo struct and put it on the ledger.
// It returns a error if any problem occurred.
func addCurriculo(stub shim.ChaincodeStubInterface, curriculoJSON string) error {

	errContext := "Nao foi possivel adicionar o curriculo"
	var curriculo assetlib.Curriculo

	// Validate JSON structure
	parameters := []string{"unidadesCurriculares", "cargaHorariaCurso", "codigoCurriculo",
		"dadosDoCurso", "iesEmissora", "minutosHoraAula", "ambiente",
		"criteriosIntegralizacao", "categoriasAtividadesComplementares", "etiquetas", "areas",
		"dataCurriculo", "nomeParaAreas", "versao", "informacoesAdicionais", "segurancaCurriculo"}

	if !isValidJSONInput(parameters, curriculoJSON) {
		return fmt.Errorf(errors.Get(errContext, "invalidJSONInput", parameters))
	}

	err := json.Unmarshal([]byte(curriculoJSON), &curriculo)
	if err != nil {
		return fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "Curriculo"))
	}

	// Verify if there is a Curriculo with the same codigoCurriculo
	curriculoExists, err := curriculoExists(stub, curriculo.CodigoCurriculo)
	if err != nil {
		return err
	}
	if curriculoExists {
		return fmt.Errorf(errors.Get(errContext, "curriculoAlreadyExists", curriculo.CodigoCurriculo))
	}

	// Verify all curriculum information
	valid, err := isValidCurriculo(curriculo)
	if !valid {
		return err
	}

	curso, err := readCurso(stub, curriculo.DadosDoCurso.IDCurso)
	if err != nil {
		return err
	}

	if curriculo.IesEmissora != curso.CodigoIES {
		return fmt.Errorf(errors.Get(errContext, "invalidCurriculo", "curriculo aponta para uma IES diferente do curso"))
	}

	ies, err := readIES(stub, curso.CodigoIES)
	if err != nil {
		return fmt.Errorf(errors.Get(errContext, "IESDoesNotExist"))
	}
	curso.Curriculos = append(curso.Curriculos, curriculo)
	var cursoObj assetlib.Curso = assetlib.Curso{
		Nome:            curso.Nome,
		Habilitacoes:    curso.Habilitacoes,
		Tipo:            curso.Tipo,
		CodigoCursoEMEC: curso.CodigoCursoEMEC,
		Curriculos:      curso.Curriculos,
		CodigoIES:       curso.CodigoIES,
	}

	// Prepare to put curriculo on world state.
	assetType := "Curriculo"
	if curriculo.Ambiente != "Producao" && curriculo.Ambiente != "Homologacao" {
		curriculoResponseKey, _ := stub.CreateCompositeKey(assetType, []string{"Producao", curriculo.CodigoCurriculo})
		stub.PutState(curriculoResponseKey, []byte(curriculoJSON))
	} else {
		curriculoResponseKey, _ := stub.CreateCompositeKey(assetType, []string{string(curriculo.Ambiente), curriculo.CodigoCurriculo})
		stub.PutState(curriculoResponseKey, []byte(curriculoJSON))
	}
	// Add new curriculo to array of cursos.
	for index, curso := range ies.Cursos {
		if curso.CodigoCursoEMEC == cursoObj.CodigoCursoEMEC {
			ies.Cursos[index] = cursoObj
			break
		}
	}
	iesJSON, _ := json.Marshal(ies)
	cursoJSON, _ := json.Marshal(cursoObj)
	assetType = "IES"
	// Update IES on ledger.
	IESKey, _ := stub.CreateCompositeKey(assetType, []string{curriculo.IesEmissora})
	stub.PutState(IESKey, iesJSON)
	// Update Curso on ledger.
	assetType = "Curso"
	cursoKey, _ := stub.CreateCompositeKey(assetType, []string{curriculo.DadosDoCurso.IDCurso})
	return stub.PutState(cursoKey, cursoJSON)
}

// readCurriculosFromCourse reads all curriculos of a curso list of curriculo.
// It returns a list of Curriculo struct, and an error if any problem occurred.
func readCurriculosFromCurso(stub shim.ChaincodeStubInterface, cursoID string) ([]assetlib.Curriculo, error) {
	errContext := "Nao foi possivel ler os curriculos do curso"
	curso, err := readCurso(stub, cursoID)
	if err != nil {
		return nil, err
	}
	if curso == nil {
		return nil, fmt.Errorf(errors.Get(errContext, "cursoDoesNotExist"))
	}
	return curso.Curriculos, nil
}

// curriculoExists Check if the curriculo given as parameter exists.
// It returns false if curriculo doesn't exists, and an error if any problem occurred.
func curriculoExists(stub shim.ChaincodeStubInterface, curriculoID string) (bool, error) {
	errContext := "Nao foi possivel verificar se o curso existe"
	assetType := "Curriculo"
	curriculoResponseKey, _ := stub.CreateCompositeKey(assetType, []string{"Producao", curriculoID})
	bytesCurriculo, err := stub.GetState(curriculoResponseKey)
	if err != nil {
		return false, fmt.Errorf(errors.Get(errContext, "cannotReadWorldState"))
	}
	return bytesCurriculo != nil, nil
}

// readAllCurrriculos reads all curriculos from the ledger.
// It returns a list of struct Curriculo, and an error if any problem ocurred.
func readAllCurriculos(stub shim.ChaincodeStubInterface) ([]assetlib.Curriculo, error) {
	errContext := "Nao foi possivel ler todos os curriculos"
	assetType := "Curriculo"
	curriculoIterator, err := stub.GetStateByPartialCompositeKey(assetType, nil)
	if err != nil {
		return nil, err
	}
	var curriculos []assetlib.Curriculo
	defer curriculoIterator.Close()
	for curriculoIterator.HasNext() {
		queryResponse, err := curriculoIterator.Next()
		if err != nil {
			return nil, err
		}
		var curriculo assetlib.Curriculo
		err = json.Unmarshal(queryResponse.Value, &curriculo)
		if err != nil {
			return nil, fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "Curriculo"))
		}
		curriculos = append(curriculos, curriculo)
	}
	_, err = json.Marshal(curriculos)
	if err != nil {
		return nil, fmt.Errorf(errors.Get(errContext, "cannotMarshal", "Curriculos"))
	}
	return curriculos, nil
}

// readCurriculosFromAmbiente reads all curriculos in the ledger that have the given ambiente.
// It returns a list of struct Curriculo, and an error if any problem ocurred.
func readCurriculosFromAmbiente(stub shim.ChaincodeStubInterface, ambiente string) ([]assetlib.Curriculo, error) {
	errContext := "Nao foi possivel ler os curriculos do ambiente"
	assetType := "Curriculo"
	var curriculos []assetlib.Curriculo

	curriculoIterator, _ := stub.GetStateByPartialCompositeKey(assetType, []string{ambiente})
	defer curriculoIterator.Close()
	for curriculoIterator.HasNext() {
		query, err := curriculoIterator.Next()
		if err != nil {
			return nil, err
		}
		var curriculo assetlib.Curriculo
		err = json.Unmarshal(query.Value, &curriculo)
		if err != nil {
			return nil, fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "Curriculo"))
		}
		if string(curriculo.Ambiente) == ambiente {
			curriculos = append(curriculos, curriculo)
		}
	}
	return curriculos, nil
}

// readCurriculosFromIES reads all curriculos in the ledger of the given IES.
// It returns a list of struct Curriculo, and an error if any problem ocurred.
func readCurriculosFromIES(stub shim.ChaincodeStubInterface, iesID string) ([]assetlib.Curriculo, error) {
	errContext := "Nao foi possivel ler os curriculos da IES"
	ies, err := readIES(stub, iesID)
	if err != nil {
		return nil, err
	}
	if ies == nil {
		return nil, fmt.Errorf(errors.Get(errContext, "IESDoesNotExist"))
	}
	// Read all cursos and all curriculos of each curso.
	curriculos := []assetlib.Curriculo{}
	for _, curso := range ies.Cursos {

		curriculos = append(curriculos, curso.Curriculos...)
	}
	return curriculos, nil
}

// readCurriculo reads a specific curriculo in the ledger.
// return a struct curriculo, and an error if any problem ocurred.
// Only curriculos with Ambiente equals to Producao.
func readCurriculo(stub shim.ChaincodeStubInterface, codigoCurriculo string) (*assetlib.Curriculo, error) {
	errContext := "Nao foi possivel ler o curriculo"
	exists, _ := curriculoExists(stub, codigoCurriculo)
	if !exists {
		return nil, fmt.Errorf(errors.Get(errContext, "curriculoDoesNotExist", codigoCurriculo))
	}
	assetType := "Curriculo"
	curriculoResponseKey, _ := stub.CreateCompositeKey(assetType, []string{"Producao", codigoCurriculo})
	bytesCurriculo, err := stub.GetState(curriculoResponseKey)
	if err != nil {
		return nil, fmt.Errorf(errors.Get(errContext, "cannotReadWorldState"))
	}
	var curriculo assetlib.Curriculo
	err = json.Unmarshal(bytesCurriculo, &curriculo)
	if err != nil {
		return nil, fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "Curriculo"))
	}
	return &curriculo, nil
}

// integralizacaoCurricular calculates based on historico escolar if a given aluno completes all the integralizacao criterios.
// It returns a string with information about the integralizacao, and an error if any problem ocurred.
func integralizacaoCurricular(stub shim.ChaincodeStubInterface, curriculoID string, alunoCPF string) (bool, error) {
	errContext := "Nao foi possivel verificar a integralizacao curricular"
	response := stub.InvokeChaincode("academicRecords", ToChaincodeArgs("ReadLastHistoricoFromAluno", alunoCPF), "jornada")

	if response.GetStatus() != shim.OK {
		return false, fmt.Errorf(errors.Get(errContext, "cannotInvokeChaincode", response.GetMessage()))
	}

	var historico assetlib.HistoricoEscolar
	payloadHistorico := response.GetPayload()
	err := json.Unmarshal(payloadHistorico, &historico)
	if err != nil {
		return false, fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "Historico"))
	}

	var sumHours int

	// Verify in the lists of disciplines the ones the giving student passed and add into sumHours.
	for _, disciplina := range historico.ElementoHistorico.Disciplinas {
		if disciplina.EstadoDisciplina.EstadoDisciplina == "Aprovado" {
			sumHours = sumHours + disciplina.CargaHoraria.CargaHoraria
		}
	}

	// Check if the sumHours is the same as the ones at CriterioIntegralizacao.
	curriculo, err := readCurriculo(stub, curriculoID)
	if err != nil {
		return false, err
	}

	// Get the correct criterioIntegralizacao.
	criterioIntegralizacao := curriculo.CriteriosIntegralizacao[len(curriculo.CriteriosIntegralizacao)-1]

	min := criterioIntegralizacao.CargasHorariasCriterio.CargaHorariaMinima.CargaHoraria
	max := criterioIntegralizacao.CargasHorariasCriterio.CargaHorariaMaxima.CargaHoraria

	if (sumHours >= min) && (sumHours <= max) {
		return true, nil
	}

	return false, nil
}

/*
#####################
### IES Functions ###
#####################
*/

// IESExists receives a codigoMEC which represent the IES identificator.
// It returns true if IES exists, and return a error if any problem occurred.
func IESExists(stub shim.ChaincodeStubInterface, codigoMEC string) (bool, error) {
	assetType := "IES"
	IESKey, _ := stub.CreateCompositeKey(assetType, []string{codigoMEC})
	iesResponse, _ := stub.GetState(IESKey)
	return iesResponse != nil, nil
}

// IESToAPI struct converts a IES struct to a API Pattern struct.
type IESToAPI struct {
	Nome          string               `json:"nome"`
	CodigoMEC     string               `json:"codigoMEC"`
	CNPJ          string               `json:"CNPJ"`
	Mantedenedora assetlib.Mantenedora `json:"mantenedora"`
}

func readAllIESToAPI(stub shim.ChaincodeStubInterface) ([]IESToAPI, error) {
	errContext := "Nao foi possivel ler todas as IES para a API"
	assetType := "IES"
	IESIterator, err := stub.GetStateByPartialCompositeKey(assetType, nil)
	if err != nil {
		return nil, err
	}

	var IESs []IESToAPI
	defer IESIterator.Close()
	for IESIterator.HasNext() {
		queryResponse, err := IESIterator.Next()
		if err != nil {
			return nil, err
		}
		var IES IESToAPI
		err = json.Unmarshal(queryResponse.Value, &IES)
		if err != nil {
			return nil, fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "IES"))
		}
		IESs = append(IESs, IES)
	}
	_, err = json.Marshal(IESs)
	if err != nil {
		return nil, fmt.Errorf(errors.Get(errContext, "cannotMarshal", "IESs"))
	}
	return IESs, nil
}

// createIES receives a JSON which represent a IES struct and put it on the ledger.
// It returns a error if any problem occurred.

func createIES(stub shim.ChaincodeStubInterface, IESJSON string) error {
	errContext := "Nao foi possivel criar a IES"
	var ies assetlib.IES
	parameters := []string{"nome", "codigoMEC", "CNPJ", "mantenedora", "cursos", "endereco", "credenciamento", "recredenciamento", "renovacaoRecredenciamento"}
	if !isValidJSONInput(parameters, IESJSON) {
		return fmt.Errorf(errors.Get(errContext, "invalidJSONInput", parameters))
	}
	err := json.Unmarshal([]byte(IESJSON), &ies)
	if err != nil {
		return fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "IES"))
	}

	assetType := "IES"
	IESKey, _ := stub.CreateCompositeKey(assetType, []string{ies.CodigoMEC})
	exists, err := IESExists(stub, ies.CodigoMEC)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf(errors.Get(errContext, "IESAlreadyExists"))
	}

	if !isValidEndereco(ies.Endereco) {
		return fmt.Errorf(errors.Get(errContext, "invalidEndereco"))
	}

	if !isValidAtoRegulatorio(ies.Credenciamento) {
		return fmt.Errorf(errors.Get(errContext, "invalidAtoRegulatorio", "Credenciamento"))
	}

	if (ies.Recredenciamento != assetlib.AtoRegulatorio{}) {
		if !isValidAtoRegulatorio(ies.Recredenciamento) {
			return fmt.Errorf(errors.Get(errContext, "invalidAtoRegulatorio", "Recredenciamento"))
		}
	}

	if (ies.RenovacaoRecredenciamento != assetlib.AtoRegulatorio{}) {
		if !isValidAtoRegulatorio(ies.RenovacaoRecredenciamento) {
			return fmt.Errorf(errors.Get(errContext, "invalidAtoRegulatorio", "RenovacaoRecredenciamento"))
		}
	}

	if !isValidEndereco(ies.Endereco) {
		return fmt.Errorf(errors.Get(errContext, "invalidEndereco"))
	}

	if !isValidEndereco(ies.Mantedenedora.Endereco) {
		return fmt.Errorf(errors.Get(errContext, "invalidEndereco"))
	}

	for _, curso := range ies.Cursos {
		cursoID, err := validateCursoAndReturnID(curso)
		if err != nil {
			return err
		}
		if curso.CodigoIES != ies.CodigoMEC {
			return fmt.Errorf(errors.Get(errContext, "invalidIES", "curso.CodigoIES deve ser igual a ies.CodigoMEC"))
		}
		cursoExists, err := cursoExists(stub, cursoID)
		if err != nil {
			return err
		}
		if cursoExists {
			return fmt.Errorf(errors.Get(errContext, "cursoAlreadyExists"))
		}
		err = createCurso(stub, curso)
		if err != nil {
			return err
		}
	}
	return stub.PutState(IESKey, []byte(IESJSON))
}

// readIES receives the ID identificator to a IES.
// It returns a IES if exists, and return a error if any problem occurred.
func readIES(stub shim.ChaincodeStubInterface, iesCodigo string) (*assetlib.IES, error) {
	errContext := "Nao foi possivel ler a IES"
	exists, _ := IESExists(stub, iesCodigo)
	if !exists {
		return nil, fmt.Errorf(errors.Get(errContext, "IESDoesNotExist"))
	}

	assetType := "IES"
	IESKey, _ := stub.CreateCompositeKey(assetType, []string{iesCodigo})

	iesJSON, _ := stub.GetState(IESKey)
	var ies assetlib.IES
	err := json.Unmarshal(iesJSON, &ies)
	if err != nil {
		return nil, fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "IES"))
	}
	return &ies, nil
}

/*
############################
### Validation Functions ###
############################
*/

// ToChaincodeArgs receives dynamic number of strings as parameters.
// It returns array byte of chaincode args.
func ToChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

// isValidJSONInput Checks if the JSON given as parameter is valid.
// It returns true if the JSON is valid and a error if any problem occurred.
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

// validateCursoAndReturnID Checks if the Curso given as parameter is valid.
// It returns a empty string if the Curso is invalid or the courseId if valid and returns a error if any problem occurred.
func validateCursoAndReturnID(curso assetlib.Curso) (string, error) {
	errContext := "Nao foi possivel validar o curso"
	var idCurso string
	for _, habilitacao := range curso.Habilitacoes {
		if habilitacao.Nome == "" {
			return "", fmt.Errorf(errors.Get(errContext, "invalidCurso", "a struct Habilitacao possui o campo nome como obrigatório."))
		}
	}

	if curso.CodigoCursoEMEC != "" {
		if (curso.Autorizacao == assetlib.AtoRegulatorio{}) {
			return "", fmt.Errorf(errors.Get(errContext, "invalidAutorizacao"))
		}

		if !(isValidAtoRegulatorio(curso.Autorizacao)) {
			return "", fmt.Errorf(errors.Get(errContext, "invalidAtoRegulatorio", "Autorizacao"))
		}

		if (curso.Reconhecimento != assetlib.AtoRegulatorio{}) {
			if !(isValidAtoRegulatorio(curso.Reconhecimento)) {
				return "", fmt.Errorf(errors.Get(errContext, "invalidAtoRegulatorio", "Reconhecimento"))
			}
		}
	} else {

		if curso.TramitacaoMEC.NumeroProcesso == "" {
			return "", fmt.Errorf(errors.Get(errContext, "invalidTramitacaoMEC"))
		}

		if (curso.Autorizacao != assetlib.AtoRegulatorio{}) {
			if !(isValidAtoRegulatorio(curso.Autorizacao)) {
				return "", fmt.Errorf(errors.Get(errContext, "invalidAtoRegulatorio", "Autorizacao"))
			}
		}

		if (curso.Reconhecimento != assetlib.AtoRegulatorio{}) {
			if !(isValidAtoRegulatorio(curso.Reconhecimento)) {
				return "", fmt.Errorf(errors.Get(errContext, "invalidAtoRegulatorio", "Reconhecimento"))
			}
		}

		if (curso.RenovacaoReconhecimento != assetlib.AtoRegulatorio{}) {
			if !(isValidAtoRegulatorio(curso.RenovacaoReconhecimento)) {
				return "", fmt.Errorf(errors.Get(errContext, "invalidAtoRegulatorio", "RenovacaoReconhecimento"))
			}
		}
	}

	if len(curso.Curriculos) < 1 {
		return "", fmt.Errorf(errors.Get(errContext, "invalidCurso", "o Curso precisa ter pelo menos um curriculo definido."))
	}
	if curso.Nome == "" {
		return "", fmt.Errorf(errors.Get(errContext, "invalidCurso", "o atributo nome é obrigatório na struct Curso."))
	}
	if curso.Tipo == "EmTramite" {
		if curso.CodigoCursoEMEC != "" {
			return "", fmt.Errorf(errors.Get(errContext, "invalidCurso", "caso o tipo do curso seja EmTramite o atributo CodigoCursoEMEC deve ser vazio."))
		}
		if curso.TramitacaoMEC.NumeroProcesso == "" || curso.TramitacaoMEC.TipoDeProcesso == "" {
			return "", fmt.Errorf(errors.Get(errContext, "invalidCurso", "NumeroProcesso e TipoDeProcesso devem estar preenchidos na struct tramitacaoMEC."))
		}
		idCurso = curso.TramitacaoMEC.NumeroProcesso
		return idCurso, nil
	} else if curso.Tipo == "ValidadoPeloMEC" {
		if curso.CodigoCursoEMEC == "" {
			return "", fmt.Errorf(errors.Get(errContext, "invalidCurso", "o CodigoCursoEMEC não pode estar vazio."))
		}
		if (curso.TramitacaoMEC != assetlib.TramitacaoMEC{}) {
			return "", fmt.Errorf(errors.Get(errContext, "invalidCurso", "caso o tipo do curso não seja TramitacaoMEC a estrutura TramitacaoMEC deve estar vazia."))
		}
		idCurso = curso.CodigoCursoEMEC
		return idCurso, nil
	} else {
		return "", fmt.Errorf(errors.Get(errContext, "invalidCurso", "o TipoDeCurso passado como parametro não é definido na struct."))
	}
}

// isValidCurriculo checks if the given curriculo is valid with validations.
// return a bool value if it is valid, and an error if any problem ocurred.
func isValidCurriculo(curriculo assetlib.Curriculo) (bool, error) {

	errContext := "Nao foi possivel validar o curriculo"
	if curriculo.CodigoCurriculo == "" {
		return false, fmt.Errorf(errors.Get(errContext, "invalidCurriculo", "o campo CodigoCurriculo é obrigatorio."))
	}

	// Seguranca cannot be empty
	if curriculo.SegurancaCurriculo == "" {
		return false, fmt.Errorf(errors.Get(errContext, "invalidCurriculo", "o campo SegurancaCurriculo é obrigatorio."))
	}

	// min 1 max unbouded
	if len(curriculo.Etiquetas) < 1 {
		return false, fmt.Errorf(errors.Get(errContext, "invalidEtiquetas", "o curriculo precisa ter pelo menos uma etiqueta."))
	}

	// minx 1 max unbouded
	if len(curriculo.UnidadesCurriculares) < 1 {
		return false, fmt.Errorf(errors.Get(errContext, "invalidCurriculo", "o campo UnidadesCurriculares é obrigatorio"))
	}

	// min 1 max unbouded
	if len(curriculo.CriteriosIntegralizacao) < 1 {
		return false, fmt.Errorf(errors.Get(errContext, "invalidCurriculo", "o campo CriteriosIntegralizacao é obrigatorio"))
	}

	for _, criterio := range curriculo.CriteriosIntegralizacao {
		validCriterio := isValidCriterio(criterio)

		if !validCriterio {

			json, _ := json.Marshal(criterio)
			return false, fmt.Errorf(string(json))

			return false, fmt.Errorf(errors.Get(errContext, "invalidCurriculo", "Não foi possível validar algum critério de integralizacao"))
		}
	}

	for _, unidade := range curriculo.UnidadesCurriculares {
		validUnidade := isValidUnidade(unidade)
		if !validUnidade {
			return false, fmt.Errorf(errors.Get(errContext, "invalidCurriculo", "Não foi possível validar as unidades curriculares"))
		}

		if unidade.TipoUnidadeCurricular == "Atividade Complementar" {
			isValidAtividade := false

			for _, categoria := range curriculo.CategoriasAtividadesComplementares {
				for _, atividade := range categoria.AtividadesComplementares {
					if atividade.Codigo == unidade.Codigo {
						isValidAtividade = true
						break
					}
				}
			}

			if !isValidAtividade {
				return false, fmt.Errorf(unidade.Codigo + " não é uma atividade complementar válida")
			}
		}

		// Validare Areas in UnidadeCurricular
		validAreas := true
		for _, area := range unidade.Areas {
			validAreas := false

			for _, areaCurriculo := range curriculo.Areas {
				if area.Codigo == areaCurriculo.Codigo {
					validAreas = true
					break
				}
			}

			if !validAreas {
				validAreas = false
			}
		}

		if !validAreas {
			return false, fmt.Errorf(errors.Get(errContext, "invalidCurriculo", "Não foi possível validar as areas das unidades curriculares"))
		}

		// Validate Etiquetas in UnidadeCurricular
		var isValidEtiqueta bool = true

		for _, etiqueta := range unidade.Etiquetas {

			// All etiquetas need to be validate
			var isValidEtiquetaCurrilar bool = false
			for _, etiquetaCurricular := range curriculo.Etiquetas {
				if etiqueta.Codigo == etiquetaCurricular.Codigo {
					isValidEtiquetaCurrilar = true
					break
				}
			}

			if !isValidEtiquetaCurrilar {
				isValidEtiqueta = false
				break
			}
		}

		if !isValidEtiqueta {
			return false, fmt.Errorf(errors.Get(errContext, "invalidEtiquetaCurriculo"))
		}
	}

	// isValidCategoria
	for _, categoria := range curriculo.CategoriasAtividadesComplementares {
		validCategoria := isValidCategoria(categoria)
		if !validCategoria {
			return false, fmt.Errorf(errors.Get(errContext, "invalidCurriculo", "Não foi possível validar as categorias de atividades complementares"))
		}
	}

	// isValidAtividade
	for _, categoria := range curriculo.CategoriasAtividadesComplementares {
		for _, atividade := range categoria.AtividadesComplementares {
			validAtividade := isValidAtividadeComplementar(atividade)
			if !validAtividade {
				return false, fmt.Errorf(errors.Get(errContext, "invalidCurriculo", "Não foi possível validar as atividades complementares"))
			}
		}
	}

	return true, nil
}

func isValidCategoria(categoria assetlib.CategoriaAtividadeComplementar) bool {

	if categoria.Codigo == "" {
		return false
	}

	if categoria.Nome == "" {
		return false
	}

	if len(categoria.AtividadesComplementares) < 1 {
		return false
	}

	return true
}

// isValidCriterio Checks if criterio de integralizacao given as parameter is valid.
// It returns true if criterio is valid and false otherwise.
func isValidCriterio(criterio assetlib.CriterioIntegralizacao) bool {

	if criterio.Tipo == "CriterioExpressao" {

		if criterio.Codigo == "" {
			return false
		}

		if len(criterio.Expressao) < 1 {
			return false
		}

	} else if criterio.Tipo == "CriterioRotulo" {

		if criterio.Codigo == "" {
			return false
		}

	} else {
		return false
	}

	return true
}

// isValidUnidade Checks if unidade curricular given as parameter is valid.
// It returns true if unidade is valid and false otherwise.
func isValidUnidade(unidade assetlib.UnidadeCurricular) bool {

	if unidade.Codigo == "" || unidade.Nome == "" || unidade.TipoUnidadeCurricular == "" {
		return false
	}

	tiposDeUnidade := []assetlib.TipoUnidadeCurricular{assetlib.EDisciplina, assetlib.EModulo, assetlib.EAtividade, assetlib.EEstagio, assetlib.ETrabalhoConclusaoCurso, assetlib.EMonografia, assetlib.EArtigo, assetlib.EProjeto, assetlib.EProduto, assetlib.EAtividadeComplementar, assetlib.EAtividadeExtensao}

	for _, unidadeAtual := range tiposDeUnidade {
		if unidadeAtual == unidade.TipoUnidadeCurricular {
			return true
		}
	}
	return false
}

// isValidAtividadeComplementar Checks if atividade complementar given as parameter is valid.
// It returns true if atividade is valid and false otherwise.
func isValidAtividadeComplementar(atividade assetlib.AtividadeComplementarCurricular) bool {

	if atividade.Codigo == "" {
		return false
	}

	if atividade.Nome == "" {
		return false
	}

	if atividade.LimiteCargaHorariaEmHoraRelogio == (assetlib.CargaHoraria{}) {
		return false
	}

	if !(isValidCargaHoraria(atividade.LimiteCargaHorariaEmHoraRelogio)) {
		return false
	}

	return true
}

// isValidEndereco Checks if endereco given as parameter is valid.

func isValidEndereco(endereco assetlib.Endereco) bool {
	if endereco.Logradouro == "" || endereco.Numero == "" || endereco.Bairro == "" || endereco.CEP == "" || (endereco.Municipio == assetlib.Municipio{}) {
		return false
	}

	found, err := regexp.MatchString("[0-9]{8}", endereco.CEP)

	if err != nil || !found {
		return false
	}

	return true
}

// isValidAtoRegulatorio
func isValidAtoRegulatorio(atoRegulatorio assetlib.AtoRegulatorio) bool {
	return (atoRegulatorio.Numero != "" && atoRegulatorio.TramitacaoMEC.NumeroProcesso == "") || (atoRegulatorio.Numero == "" && atoRegulatorio.TramitacaoMEC.NumeroProcesso != "")
}

// isValidCargaHoraria
func isValidCargaHoraria(cargaHoraria assetlib.CargaHoraria) bool {
	return cargaHoraria.CargaHoraria >= 0
}
