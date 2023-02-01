package main

import (
	"assetlib"
	"encoding/json"
	errors "errorMessages"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

type Student struct {
}

// Init Function
func (s *Student) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *Student) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	errContext := "Nao foi possivel invocar a funcao."
	fn, args := stub.GetFunctionAndParameters()

	switch fn {
	case "VerifyStudentLogin":
		if len(args) != 2 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "2")
			return shim.Error(errorMsg)
		}
		studentLogin, err := verifyStudentLogin(stub, args[0], args[1])
		if err != nil {
			return shim.Error(err.Error())
		}
		responseBytes, err := json.Marshal(*studentLogin)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Historico")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	case "RegisterStudent":
		if len(args) != 4 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "4")
			return shim.Error(errorMsg)
		}
		studentLogin, err := registerStudent(stub, args[0], args[1], args[2], args[3])
		if err != nil {
			return shim.Error(err.Error())
		}
		responseBytes, err := json.Marshal(*studentLogin)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Historico")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	case "AllStudents":
		if len(args) != 0 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "0")
			return shim.Error(errorMsg)
		}
		students, err := allStudents(stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		responseBytes, err := json.Marshal(students)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Alunos")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	case "GetAtividades":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "1")
			return shim.Error(errorMsg)
		}
		atividades, err := getAtividades(stub, args[0])
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
	case "GetEstagios":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "1")
			return shim.Error(errorMsg)
		}
		estagios, err := getEstagios(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		responseBytes, err := json.Marshal(estagios)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Estagios")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	case "GetHistoricos":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "1")
			return shim.Error(errorMsg)
		}
		historicos, err := getHistoricos(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		responseBytes, err := json.Marshal(historicos)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Historicos")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	case "AddAtividade":
		if len(args) != 2 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "2")
			return shim.Error(errorMsg)
		}
		atividade, err := addAtividade(stub, args[0], args[1])
		if err != nil {
			return shim.Error(err.Error())
		}
		responseBytes, err := json.Marshal(atividade)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Atividade")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	case "AddEstagio":
		if len(args) != 2 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "2")
			return shim.Error(errorMsg)
		}
		estagio, err := addEstagio(stub, args[0], args[1])
		if err != nil {
			return shim.Error(err.Error())
		}
		responseBytes, err := json.Marshal(estagio)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Estagio")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	case "GetAtividadesPendente":
		if len(args) != 1 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "1")
			return shim.Error(errorMsg)
		}
		atividades, err := getAtividadesPendentes(stub, args[0])
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
	case "ApproveAndRemoveAtividade":
		if len(args) != 2 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "2")
			return shim.Error(errorMsg)
		}
		atividade, err := approveAndRemoveAtividade(stub, args[0], args[1])
		if err != nil {
			return shim.Error(err.Error())
		}
		responseBytes, err := json.Marshal(atividade)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Atividade")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	case "ApproveAndRemoveEstagio":
		if len(args) != 2 {
			errorMsg := errors.Get(errContext, "invalidNumberOfArguments", "2")
			return shim.Error(errorMsg)
		}
		estagio, err := approveAndRemoveEstagio(stub, args[0], args[1])
		if err != nil {
			return shim.Error(err.Error())
		}
		responseBytes, err := json.Marshal(estagio)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotMarshal", "Estagio")
			return shim.Error(errorMsg)
		}
		return shim.Success(responseBytes)
	default:
		errorMsg := errors.Get(errContext, "undefinedFunction", fn)
		return shim.Error(errorMsg)
	}
}

// Main function
func main() {
	if err := shim.Start(new(Student)); err != nil {
		log.Panicf("Erro ao tentar iniciar a chaincode academicRecords: %v", err)
	}
}

type StudentResponse struct {
	NomeDeUsuario            string                           `json:"nomeDeUsuario"` // required, login
	Email                    string                           `json:"email"`         // required
	CPF                      string                           `json:"CPF"`           // required
	HistoricosEscolares      []assetlib.HistoricoEscolar      `json:"historicosEscolares"`
	AtividadesComplementares []assetlib.AtividadeComplementar `json:"atividadesComplementares"`
	Estagios                 []assetlib.Estagio               `json:"estagios"`
}

// verifyStudentLogin verifies if the student's login is valid
func verifyStudentLogin(stub shim.ChaincodeStubInterface, studentLogin string, password string) (*StudentResponse, error) {
	// Get the student's data from the ledger
	errContext := "Nao foi possivel validar o login do estudante."
	assetType := "student"
	studentResponseKey, _ := stub.CreateCompositeKey(assetType, []string{studentLogin})
	studentResponseBytes, err := stub.GetState(studentResponseKey)
	if err != nil {
		errorMsg := errors.Get(errContext, "InvalidLogin", studentLogin)
		return nil, fmt.Errorf(errorMsg)
	}
	// Verify if the password is correct
	var student assetlib.PerfilDoEstudante
	err = json.Unmarshal(studentResponseBytes, &student)
	if err != nil {
		return nil, fmt.Errorf(errContext)
	}
	if student.Senha != password {
		errorMsg := errors.Get(errContext, "InvalidPassword")
		return nil, fmt.Errorf(errorMsg)
	}
	studentResponse, err := getStudentData(stub, student)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotGetStudentData")
		return nil, fmt.Errorf(errorMsg)
	}
	return studentResponse, nil
}

// Register a new student to the ledger and make the student login
// The password is encrypted
func registerStudent(stub shim.ChaincodeStubInterface, nomeDoUsuario string, CPF string, email string, senha string) (*StudentResponse, error) {
	errContext := "Nao foi possivel registrar o estudante."
	students, err := allStudents(stub)
	if err != nil {
		return nil, err
	}
	// Verify if the student already exists
	for _, student := range students {
		if student.CPF == CPF {
			return nil, fmt.Errorf("cpf já cadastrado")
		}
		if student.Email == email {
			return nil, fmt.Errorf("email já cadastrado")
		}
		if student.NomeDeUsuario == nomeDoUsuario {
			return nil, fmt.Errorf("nome de usuário já cadastrado")
		}
	}
	// Verify if the student's data is valid to be registered
	// A student profile is valid if the CPF was been registered in academicRecords chaincode.
	// The CPF is the key to the student's profile in academicRecords chaincode.
	response := stub.InvokeChaincode("academicRecords", ToChaincodeArgs("AlunoExists", CPF), "jornada")
	if response.Status != shim.OK {
		return nil, fmt.Errorf(response.Message)
	}
	// The password is encrypted
	// User can be registered
	student := assetlib.PerfilDoEstudante{
		NomeDeUsuario: nomeDoUsuario,
		CPF:           CPF,
		Email:         email,
		Senha:         senha,
	}
	// Save the student's data in the ledger
	assetType := "student"
	studentResponseKey, err := stub.CreateCompositeKey(assetType, []string{student.NomeDeUsuario})
	if err != nil {
		return nil, fmt.Errorf("nao foi possivel registrar o estudante")
	}
	studentProfileJson, err := json.Marshal(student)
	if err != nil {
		return nil, fmt.Errorf("nao foi possivel registrar o estudante")
	}
	err = stub.PutState(studentResponseKey, []byte(studentProfileJson))
	if err != nil {
		// return nil, shim.Error("Nao foi possivel registrar o estudante.")
		errorMsg := errors.Get(errContext, "cannotRegisterStudent")
		return nil, fmt.Errorf(errorMsg)
	}
	// Return the student's data
	studentResponse, err := getStudentData(stub, student)
	if err != nil {
		return nil, err
	}

	return studentResponse, nil
}

func getStudentData(stub shim.ChaincodeStubInterface, student assetlib.PerfilDoEstudante) (*StudentResponse, error) {
	errContext := "Falha para obter dados do estudante."
	// Return the student's data
	response := stub.InvokeChaincode("academicRecords", ToChaincodeArgs("ReadHistoricosFromAluno", student.CPF), "jornada")
	var historicos []assetlib.HistoricoEscolar
	err := json.Unmarshal(response.Payload, &historicos)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Historicos")
		return nil, fmt.Errorf(errorMsg)
	}

	response = stub.InvokeChaincode("academicRecords", ToChaincodeArgs("ReadAtividadesComplementaresFromAluno", student.CPF), "jornada")
	var atividades []assetlib.AtividadeComplementar
	err = json.Unmarshal(response.Payload, &atividades)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Atividades")
		return nil, fmt.Errorf(errorMsg)
	}

	response = stub.InvokeChaincode("academicRecords", ToChaincodeArgs("ReadEstagiosFromAluno", student.CPF), "jornada")
	var estagios []assetlib.Estagio
	err = json.Unmarshal(response.Payload, &estagios)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Estagios")
		return nil, fmt.Errorf(errorMsg)
	}

	var studentResponse StudentResponse = StudentResponse{
		NomeDeUsuario:            student.NomeDeUsuario,
		Email:                    student.Email,
		CPF:                      student.CPF,
		HistoricosEscolares:      historicos,
		AtividadesComplementares: atividades,
		Estagios:                 estagios,
	}

	return &studentResponse, nil
}

type StudentProfile struct {
	NomeDeUsuario string `json:"nomeDeUsuario"`
	CPF           string `json:"CPF"`
	Email         string `json:"email"`
}

func allStudents(stub shim.ChaincodeStubInterface) ([]StudentProfile, error) {

	errContext := "Não foi possível recuperar todos os alunos da ledger"
	assetType := "student"
	studentsIterator, err := stub.GetStateByPartialCompositeKey(assetType, nil)
	if err != nil {
		return nil, err
	}

	var students []StudentProfile
	defer studentsIterator.Close()
	for studentsIterator.HasNext() {
		queryResponse, err := studentsIterator.Next()
		if err != nil {
			return nil, err
		}
		var aluno StudentProfile
		err = json.Unmarshal(queryResponse.Value, &aluno)
		if err != nil {
			return nil, fmt.Errorf(errors.Get(errContext, "cannotUnmarshal", "Student"))
		}
		students = append(students, aluno)
	}

	_, err = json.Marshal(students)
	if err != nil {
		return nil, fmt.Errorf(errors.Get(errContext, "cannotMarshal", "Student"))
	}

	return students, nil
}

func getAtividades(stub shim.ChaincodeStubInterface, CPF string) ([]assetlib.AtividadeComplementar, error) {
	errContext := "Nao foi possivel obter todas as atividades."
	response := stub.InvokeChaincode("academicRecords", ToChaincodeArgs("ReadAtividadesComplementaresFromAluno", CPF), "jornada")
	if response.GetStatus() != shim.OK {
		return nil, fmt.Errorf(errContext)
	}
	// Get atividades from LAST HISTORICO (only approved ones)
	var atividades []assetlib.AtividadeComplementar
	payloadAtividades := response.GetPayload()
	err := json.Unmarshal(payloadAtividades, &atividades)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Atividades")
		return nil, fmt.Errorf(errorMsg)
	}
	// Get other atividades including pending and rejected ones
	resultsIterator, err := stub.GetStateByPartialCompositeKey("activity", []string{CPF})
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotGetStateByPartialCompositeKey", "Atividades")
		return nil, fmt.Errorf(errorMsg)
	}
	defer resultsIterator.Close()
	for i := 0; resultsIterator.HasNext(); i++ {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf(errContext)
		}
		var atividade assetlib.AtividadeComplementar
		err = json.Unmarshal(response.Value, &atividade)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotUnmarshal", "Atividade")
			return nil, fmt.Errorf(errorMsg)
		}
		atividades = append(atividades, atividade)
	}
	return atividades, nil
}

func getAtividadesPendentes(stub shim.ChaincodeStubInterface, CPF string) ([]assetlib.AtividadeComplementar, error) {
	errContext := "Nao foi possivel obter todas as atividades."
	resultsIterator, err := stub.GetStateByPartialCompositeKey("activity", []string{CPF})
	var atividades []assetlib.AtividadeComplementar
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotGetStateByPartialCompositeKey", "Atividades")
		return nil, fmt.Errorf(errorMsg)
	}
	defer resultsIterator.Close()
	for i := 0; resultsIterator.HasNext(); i++ {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf(errContext)
		}
		var atividade assetlib.AtividadeComplementar
		err = json.Unmarshal(response.Value, &atividade)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotUnmarshal", "Atividade")
			return nil, fmt.Errorf(errorMsg)
		}
		atividades = append(atividades, atividade)
	}
	return atividades, nil
}

func getEstagios(stub shim.ChaincodeStubInterface, CPF string) ([]assetlib.Estagio, error) {
	errContext := "Nao foi possivel obter todos os estagios."
	response := stub.InvokeChaincode("academicRecords", ToChaincodeArgs("ReadEstagiosFromAluno", CPF), "jornada")
	if response.GetStatus() != shim.OK {
		return nil, fmt.Errorf(errContext)
	}
	// Get estagios from LAST HISTORICO (only approved ones)
	var estagios []assetlib.Estagio
	payloadAtividades := response.GetPayload()
	err := json.Unmarshal(payloadAtividades, &estagios)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Estagios")
		return nil, fmt.Errorf(errorMsg)
	}
	// Get other estagios including pending and rejected ones
	resultsIterator, err := stub.GetStateByPartialCompositeKey("internship", []string{CPF})
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotGetStateByPartialCompositeKey", "Estagios")
		return nil, fmt.Errorf(errorMsg)
	}
	defer resultsIterator.Close()
	for i := 0; resultsIterator.HasNext(); i++ {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf(errContext)
		}
		var estagio assetlib.Estagio
		err = json.Unmarshal(response.Value, &estagio)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotUnmarshal", "Estagio")
			return nil, fmt.Errorf(errorMsg)
		}
		estagios = append(estagios, estagio)
	}
	return estagios, nil
}

func getEstagiosPendente(stub shim.ChaincodeStubInterface, CPF string) ([]assetlib.Estagio, error) {
	errContext := "Nao foi possivel obter todas as estagios."
	resultsIterator, err := stub.GetStateByPartialCompositeKey("internship", []string{CPF})
	var estagios []assetlib.Estagio
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotGetStateByPartialCompositeKey", "Estagios")
		return nil, fmt.Errorf(errorMsg)
	}
	defer resultsIterator.Close()
	for i := 0; resultsIterator.HasNext(); i++ {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf(errContext)
		}
		var estagio assetlib.Estagio
		err = json.Unmarshal(response.Value, &estagio)
		if err != nil {
			errorMsg := errors.Get(errContext, "cannotUnmarshal", "Estagio")
			return nil, fmt.Errorf(errorMsg)
		}
		estagios = append(estagios, estagio)
	}
	return estagios, nil
}

func getHistoricos(stub shim.ChaincodeStubInterface, CPF string) ([]assetlib.HistoricoEscolar, error) {
	errContext := "Nao foi possivel obter todos os historicos."
	response := stub.InvokeChaincode("academicRecords", ToChaincodeArgs("ReadHistoricosFromAluno", CPF), "jornada")
	if response.GetStatus() != shim.OK {
		errorMsg := errors.Get(errContext, "cannotInvokeChaincode", response.GetMessage())
		return nil, fmt.Errorf(errorMsg)
	}
	var historicos []assetlib.HistoricoEscolar
	payloadHistoricos := response.GetPayload()
	err := json.Unmarshal(payloadHistoricos, &historicos)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Historicos")
		return nil, fmt.Errorf(errorMsg)
	}
	return historicos, nil
}

func addAtividade(stub shim.ChaincodeStubInterface, CPF string, atividadeJSON string) (*assetlib.AtividadeComplementar, error) {
	errContext := "Nao foi possivel criar uma nova atividade."
	var atividadeComplementar assetlib.AtividadeComplementar
	err := json.Unmarshal([]byte(atividadeJSON), &atividadeComplementar)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "AtividadeComplementar")
		return nil, fmt.Errorf(errorMsg)
	}
	// Create unique ID to each activity
	id := uuid.New()
	atividadeComplementar.Id = id.String()
	// As the activity is created, it is automatically set to "Pendente"
	atividadeComplementar.Situacao = "Pendente"
	// Update dataRegistro to current date
	atividadeComplementar.DataRegistro = time.Now()
	// Save activity on the ledge temporarily, before it is approved by the coordinator
	assetType := "activity"
	activityKey, err := stub.CreateCompositeKey(assetType, []string{CPF, atividadeComplementar.Id})
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotCreateCompositeKey", "AtividadeComplementar")
		return nil, fmt.Errorf(errorMsg)
	}
	activityJSON, err := json.Marshal(atividadeComplementar)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotMarshal", "AtividadeComplementar")
		return nil, fmt.Errorf(errorMsg)
	}
	err = stub.PutState(activityKey, activityJSON)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotPutState", "AtividadeComplementar")
		return nil, fmt.Errorf(errorMsg)
	}
	return &atividadeComplementar, nil
}

func addEstagio(stub shim.ChaincodeStubInterface, CPF string, estagioJSON string) (*assetlib.Estagio, error) {
	errContext := "Nao foi possivel criar um novo estagio."
	var estagio assetlib.Estagio
	err := json.Unmarshal([]byte(estagioJSON), &estagio)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Estagio")
		return nil, fmt.Errorf(errorMsg)
	}
	// Create unique ID to each internship
	id := uuid.New()
	estagio.Id = id.String()
	// As the internship is created, it is automatically set to "Pendente"
	estagio.Situacao = "Pendente"
	// Save internship on the ledge temporarily, before it is approved by the coordinator
	assetType := "internship"
	internshipKey, err := stub.CreateCompositeKey(assetType, []string{CPF, estagio.Id})
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotCreateCompositeKey", "Estagio")
		return nil, fmt.Errorf(errorMsg)
	}
	internshipJSON, err := json.Marshal(estagio)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotMarshal", "Estagio")
		return nil, fmt.Errorf(errorMsg)
	}
	err = stub.PutState(internshipKey, internshipJSON)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotPutState", "Estagio")
		return nil, fmt.Errorf(errorMsg)
	}
	return &estagio, nil
}

// Returns assetlib.Atividade as JSON string from the ledger
func approveAndRemoveAtividade(stub shim.ChaincodeStubInterface, CPF string, UUID string) (*assetlib.AtividadeComplementar, error) {
	errContext := "Nao foi possivel aprovar a atividade."
	assetType := "activity"
	activityKey, err := stub.CreateCompositeKey(assetType, []string{CPF, UUID})
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotCreateCompositeKey", "AtividadeComplementar")
		return nil, fmt.Errorf(errorMsg)
	}
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotGetState", "AtividadeComplementar")
		return nil, fmt.Errorf(errorMsg)
	}
	atividadesPendentes, _ := getAtividadesPendentes(stub, CPF)
	var atividadeComplementar assetlib.AtividadeComplementar
	for _, atividade := range atividadesPendentes {
		if atividade.Id == UUID {
			atividadeComplementar = atividade
			break
		}
	}
	// As the activity is approved, it is automatically set to "Aprovado"
	atividadeComplementar.Situacao = "Aprovado"
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotMarshal", "AtividadeComplementar")
		return nil, fmt.Errorf(errorMsg)
	}
	// Remove activity from the ledge
	err = stub.DelState(activityKey)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotDelState", "AtividadeComplementar")
		return nil, fmt.Errorf(errorMsg)
	}
	return &atividadeComplementar, nil
}

func approveAndRemoveEstagio(stub shim.ChaincodeStubInterface, CPF string, UUID string) (*assetlib.Estagio, error) {
	errContext := "Nao foi possivel aprovar o estagio."
	assetType := "internship"
	internshipKey, err := stub.CreateCompositeKey(assetType, []string{CPF, UUID})
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotCreateCompositeKey", "Estagio")
		return nil, fmt.Errorf(errorMsg)
	}
	internshipJSON, err := stub.GetState(internshipKey)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotGetState", "Estagio")
		return nil, fmt.Errorf(errorMsg)
	}
	var estagio assetlib.Estagio
	err = json.Unmarshal(internshipJSON, &estagio)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotUnmarshal", "Estagio")
		return nil, fmt.Errorf(errorMsg)
	}
	// As the internship is approved, it is automatically set to "Aprovado"
	estagio.Situacao = "Aprovado"
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotMarshal", "Estagio")
		return nil, fmt.Errorf(errorMsg)
	}
	err = stub.PutState(internshipKey, internshipJSON)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotPutState", "Estagio")
		return nil, fmt.Errorf(errorMsg)
	}
	// Remove internship from the ledge
	err = stub.DelState(internshipKey)
	if err != nil {
		errorMsg := errors.Get(errContext, "cannotDelState", "Estagio")
		return nil, fmt.Errorf(errorMsg)
	}
	return &estagio, nil
}

// / ToChaincodeArgs receives dynamic number of strings as parameters.
// It returns array byte of chaincode args.
func ToChaincodeArgs(args ...string) [][]byte {

	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}

	return bargs
}
