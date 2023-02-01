package assetlib

import (
	"time"
)

// Aluno will hold all informations about a Student.
// Struct definition:
// If DocumentRG has a true value the attribute RG need to be assigned.
// Else the attribute OutroDocumento need to be assigned.
// This struct will be primary struct on world state.
// The reference for this struct is CPF.
type Aluno struct {
	ID             string             `json:"ID"`
	CPF            string             `json:"CPF"`
	Nome           string             `json:"nome"`
	Sexo           string             `json:"sexo"`
	Nacionalidade  string             `json:"nacionalidade"`
	Naturalidade   Municipio          `json:"naturalidade"`
	RG             RG                 `json:"RG"`
	OutroDocumento OutroDocumento     `json:"outroDocumento"`
	DataNascimento time.Time          `json:"dataNascimento"`
	DocumentoRG    bool               `json:"documentoRG"`
	Historicos     []HistoricoEscolar `json:"historicos"`
	NomeSocial     string             `json:"nomeSocial"`
}

// Municipio will hold all informations about a municipio.
type Municipio struct {
	CodigoMunicipio          string `json:"codigoMunicipio"`
	NomeMunicipio            string `json:"nomeMunicipio"`
	UF                       string `json:"UF"`
	EhEstrangeiro            bool   `json:"ehEstrangeiro"`
	NomeMunicipioEstrangeiro string `json:"nomeMunicipioEstrangeiro"`
}

// RG will be the identification about a Student.
type RG struct {
	Numero         string `json:"numero"`
	OrgaoExpedidor string `json:"orgaoExpedidor"`
	UF             string `json:"UF"`
}

// OutroDocumento will be the identification about a Student if he doesnt have RG.
type OutroDocumento struct {
	TipoDocumento string `json:"tipoDocumento"`
	Identificador string `json:"identificador"`
}

// tipoDeCurso.
type tipoDeCurso string

// EEmTramite.
const (
	EEmTramite       tipoDeCurso = "EmTramite"
	EValidadoPeloMEC tipoDeCurso = "ValidadoPeloMEC"
)

// Curso will hold all informations about a Course.
// Struct definition:
// If tipoDeCurso is EmTramite, TramitacaoMEC need to be assigned.
// This struct will be a primary struct on world state.
// TramitacaoMEC.numeroDoProcesso or CodigoCursoEMEC.
// If the type of the course is EValidadoPeloMEC, the fields Autorizacao and Reconhecimento are required
// and RenovacaoReconhecimento are not required.
// If the type of the course is EEmTramite, the fields Autorizacao, Reconhecimento, RenovacaoReconhecimento
// are not required.
type Curso struct {
	Nome                    string         `json:"nome"`
	Tipo                    tipoDeCurso    `json:"tipo"`
	Habilitacoes            []Habilitacao  `json:"habilitacoes"`
	CodigoCursoEMEC         string         `json:"codigoCursoEMEC"`
	TramitacaoMEC           TramitacaoMEC  `json:"tramitacaoMEC"`
	Curriculos              []Curriculo    `json:"curriculos"`
	CodigoIES               string         `json:"codigoIES"`
	Autorizacao             AtoRegulatorio `json:"autorizacao"`
	Reconhecimento          AtoRegulatorio `json:"reconhecimento"`
	RenovacaoReconhecimento AtoRegulatorio `json:"renovacaoReconhecimento"`
}

type tipoAtoRegulatorio string

// EParecer
const (
	EParecer      tipoAtoRegulatorio = "Parecer"
	EResolucao    tipoAtoRegulatorio = "Resolução"
	EDecreto      tipoAtoRegulatorio = "Decreto"
	EPortaria     tipoAtoRegulatorio = "Portaria"
	EDeliberação  tipoAtoRegulatorio = "Deliberação"
	ELeiFederal   tipoAtoRegulatorio = "Lei Federal"
	ELeiEstadual  tipoAtoRegulatorio = "Lei Estadual"
	ELeiMunicipal tipoAtoRegulatorio = "Lei Municipal"
	EAtoPróprio   tipoAtoRegulatorio = "Ato Próprio"
)

// AtoRegulatorio will hold all informations about a AtoRegulatorio.
// DOU need to be a positive number.
// The field Number is defined as string with a regex.
type AtoRegulatorio struct {
	TipoAtoRegulatorio tipoAtoRegulatorio `json:"tipoAtoRegulatorio"` // required field
	Numero             string             `json:"numero"`             // required field
	Data               time.Time          `json:"data"`               // required field
	VeiculoPublicacao  string             `json:"veiculoPublicacao"`  // min-occurs: 0
	DataPublicacao     time.Time          `json:"dataPublicacao"`     // min-occurs: 0
	SecaoPublicacao    string             `json:"secaoPublicacao"`    // min-occurs: 0
	PaginaPublicacao   string             `json:"paginaPublicacao"`   // min-occurs: 0
	DOU                string             `json:"dou"`                // min-occurs: 0
	TramitacaoMEC      TramitacaoMEC      `json:"tramitacaoMEC"`      // required field
}

// Habilitacao will contain all the information about the Habilitations of a course.
type Habilitacao struct {
	Nome string    `json:"nome"`
	Data time.Time `json:"data"`
}

// ambiente.
type tipoAmbiente string

// EProducao.
const (
	EProducao    tipoAmbiente = "Producao"
	EHomologacao tipoAmbiente = "Homologacao"
)

// TramitacaoMEC exists If a course has not yet been validated by the MEC, the validation process.
// code (NumeroDeProcesso) must be the reference for its identification.
type TramitacaoMEC struct {
	NumeroProcesso string    `json:"numeroProcesso"`
	TipoDeProcesso string    `json:"tipoDeProcesso"`
	DataCadastro   time.Time `json:"dataCadastro"`
	DataProtocolo  time.Time `json:"dataProtocolo"`
}

// Curriculo will hold all information about a Curriculo.
// Struct definition:
// DadosCurso will hold informations about the course.
// This struct will be a primary struct in world state.
// The reference for this struct in world state is a composite key of [Ambiente/CodigoCurriculo].
type Curriculo struct {
	CodigoCurriculo                    string                           `json:"codigoCurriculo"`
	DadosDoCurso                       DadosCurso                       `json:"dadosDoCurso"`
	IesEmissora                        string                           `json:"iesEmissora"`
	MinutosRelogioHoraAula             uint64                           `json:"minutosHoraAula"`
	UnidadesCurriculares               []UnidadeCurricular              `json:"unidadesCurriculares"`
	CriteriosIntegralizacao            []CriterioIntegralizacao         `json:"criteriosIntegralizacao"`
	CategoriasAtividadesComplementares []CategoriaAtividadeComplementar `json:"categoriasAtividadesComplementares"`
	Versao                             string                           `json:"versao"`
	Ambiente                           tipoAmbiente                     `json:"ambiente"`
	CargaHorariaCurso                  CargaHoraria                     `json:"cargaHorariaCurso"`
	NomeParaAreas                      string                           `json:"nomeParaAreas"` // required field
	Etiquetas                          []Etiqueta                       `json:"etiquetas"`     // min-occurs: 1
	Areas                              []Area                           `json:"areas"`         // min-occurs: 1
	DataCurriculo                      time.Time                        `json:"dataCurriculo"` // required field
	InformacoesAdicionais              string                           `json:"informacoesAdicionais"`
	SegurancaCurriculo                 string                           `json:"segurancaCurriculo"`
}

// DadosCurso will hold informations about the course.
type DadosCurso struct {
	Nome      string `json:"nome"`
	CodigoIES string `json:"codigoIES"`
	IDCurso   string `json:"idCurso"`
}

// UnidadeCurricular will hold informations about an UnidadeCurricular.
// Equivalencias is a list containing similar Unidades.
// Etiquetas will have a detail description that will be used in integralizacao.
// tipoUnidadeCurricular enum containing all the valid types.
type UnidadeCurricular struct {
	Codigo                    string                `json:"codigo"`
	Nome                      string                `json:"nome"`
	CargaHorariaEmHoraRelogio CargaHoraria          `json:"cargaHorariaEmHoraRelogio"`
	CargaHorariaEmHoraAula    CargaHoraria          `json:"cargaHorariaEmHoraAula"`
	Ementa                    []string              `json:"ementa"`
	Fase                      string                `json:"fase"`
	TipoUnidadeCurricular     TipoUnidadeCurricular `json:"tipoUnidadeCurricular"`
	Equivalencias             []string              `json:"equivalencias"`
	Etiquetas                 []Etiqueta            `json:"etiquetas"`
	Areas                     []Area                `json:"areas"`
	PreRequisitos             []string              `json:"preRequisitos"`
}

// CriterioIntegralizacao will hold informations about the CriterioIntegralizacao.
// Each CriterioIntegralizacao will have a Equacao to calculate it.
// CargasHorariasCriterio and CargasHorariasExtensao will hold the minimum and maximum hours.
type CriterioIntegralizacao struct {
	Codigo                 string            `json:"codigo"`
	Equacao                string            `json:"equacao"`
	UnidadesCurriculares   []string          `json:"unidadesCurriculares"`
	Etiquetas              []string          `json:"etiquetas"`
	Tipo                   string            `json:"tipo"`
	CargasHorariasCriterio CriterioConclusao `json:"cargasHorariasCriterio"`
	Expressao              []string          `json:"expressao"`
}

// CriterioConclusao will hold informations about the CargaHorariaMinima/CargaHorariaMaxima.
// of a CriterioConclusao.
type CriterioConclusao struct {
	CargaHorariaMinima CargaHoraria `json:"cargaHorariaMinima"`
	CargaHorariaMaxima CargaHoraria `json:"cargaHorariaMaxima"`
	CargaHorariaTotal  CargaHoraria `json:"cargaHorariaTotal"`
}

// TipoUnidadeCurricular Enum.
type TipoUnidadeCurricular string

// EDisciplina.
const (
	EDisciplina             TipoUnidadeCurricular = "Disciplina"
	EModulo                 TipoUnidadeCurricular = "Módulo"
	EAtividade              TipoUnidadeCurricular = "Atividade"
	EEstagio                TipoUnidadeCurricular = "Estágio"
	ETrabalhoConclusaoCurso TipoUnidadeCurricular = "Trabalho de Conclusão de Curso"
	EMonografia             TipoUnidadeCurricular = "Monografia"
	EArtigo                 TipoUnidadeCurricular = "Artigo"
	EProjeto                TipoUnidadeCurricular = "Projeto"
	EProduto                TipoUnidadeCurricular = "Produto"
	EAtividadeComplementar  TipoUnidadeCurricular = "Atividade Complementar"
	EAtividadeExtensao      TipoUnidadeCurricular = "Atividade de Extensão"
)

// Etiqueta will have a detail description about an UnidadeCurricular.
type Etiqueta struct {
	Nome                    string `json:"nome"`
	Codigo                  string `json:"codigo"`
	AplicadaAutomaticamente bool   `json:"aplicadaAutomaticamente"`
}

// tipoDeNota.
type tipoDeNota string

// ENota.
const (
	ENota                      tipoDeNota = "Nota"
	ENotaAteCem                tipoDeNota = "NotaAteCem"
	EConceito                  tipoDeNota = "Conceito"
	EConceitoRM                tipoDeNota = "ConceitoRM"
	EConceitoEspecificoDoCurso tipoDeNota = "ConceitoEspecificoDoCurso"
)

// Disciplina will hold all information about a single discipline.
// The attribute TipoDeNota can be five diferents types : Nota, NotaAteCem, Conceito, ConceitoRM, ConceitoEspecificoDoCurso.
// Nota : 0 up to 10tipoDenota.
// ConceitoNota : A+, A, A-, ..., F+, F, F-.
// ConceitoEspecificoDoCurso : A non-default value.
// In the attribute "Nota" all possibility of values is saved as String.
type Disciplina struct {
	Nome             string           `json:"nome"`
	Periodo          string           `json:"periodo"`
	Codigo           string           `json:"codigo"`
	EstadoDisciplina EstadoDisciplina `json:"estadoDisciplina"`
	Docentes         []Docente        `json:"docentes"`
	CargaHoraria     CargaHoraria     `json:"cargaHoraria"`
	Nota             string           `json:"nota"`
	TipoDeNota       tipoDeNota       `json:"tipoDeNota"`
	Curricular       bool             `json:"curricular"`
}

// CargaHoraria will define the type of workload in a discipline.
// If HoraAula is false the workload will be reference to hours.
type CargaHoraria struct {
	CargaHoraria int  `json:"cargaHoraria"`
	HoraAula     bool `json:"horaAula"`
}

// Docente will hold informations about a docente.
type Docente struct {
	Nome      string `json:"nome"`
	Titulacao string `json:"titulacao"`
	Lattes    string `json:"lattes"`
	CPF       string `json:"CPF"`
}

// tipoDeEstadoDisciplina.
type tipoDeEstadoDisciplina string

// EAprovado.
const (
	EAprovado  tipoDeEstadoDisciplina = "Aprovado"
	EReprovado tipoDeEstadoDisciplina = "Reprovado"
	EPendente  tipoDeEstadoDisciplina = "Pendente"
)

// EstadoDisciplina defines a single state of discipline state.
// If the attribute Discipline Status was set to EAproved the attribute Integralizacao needs to be assigned.
type EstadoDisciplina struct {
	Integralizacao   Integralizacao         `json:"integralizacao"`
	EstadoDisciplina tipoDeEstadoDisciplina `json:"tipoDeEstadoDisciplina"`
}

// tipoDeFormaIntegralizacao.
type tipoDeFormaIntegralizacao string

// All possibilities to FormaIntegralização in Intregalizacao Struct.
const (
	EAprovadoIntegralizacao   tipoDeFormaIntegralizacao = "Aprovado"
	ECursadoIntegralizacao    tipoDeFormaIntegralizacao = "Cursado"
	EValidadoIntegralizacao   tipoDeFormaIntegralizacao = "Validado"
	EOutraFormaIntegralizacao tipoDeFormaIntegralizacao = "OutraForma"
)

// Integralizacao about a student.
type Integralizacao struct {
	OutraForma          string                    `json:"outraFormaIntegralizacao"`
	FormaIntegralizacao tipoDeFormaIntegralizacao `json:"formaIntegralizacao"`
}

// IES will hold all information about a IES (Instituicao De Ensino Superior).
// This struct is a primary struct in world state.
// The reference for this struct in world state is CodigoMEC.
type IES struct {
	Nome                      string         `json:"nome"`        // required field
	CodigoMEC                 string         `json:"codigoMEC"`   // required field
	CNPJ                      string         `json:"CNPJ"`        // required field
	Mantedenedora             Mantenedora    `json:"mantenedora"` // min-occurs: 0
	Cursos                    []Curso        `json:"cursos"`
	Endereco                  Endereco       `json:"endereco"`                  // required field
	Credenciamento            AtoRegulatorio `json:"credenciamento"`            // require field
	Recredenciamento          AtoRegulatorio `json:"recredenciamento"`          // min-occurs: 0
	RenovacaoRecredenciamento AtoRegulatorio `json:"renovacaoRecredenciamento"` // min-occurs: 0
}

// Endereco will define the address struct.
type Endereco struct {
	Logradouro  string    `json:"logradouro"`
	Numero      string    `json:"numero"`
	Complemento string    `json:"complemento"`
	Bairro      string    `json:"bairro"`
	Municipio   Municipio `json:"municipio"`
	CEP         string    `json:"CEP"`
}

// Mantenedora will hold informations about a mantenedora.
type Mantenedora struct {
	RazaoSocial string   `json:"razaoSocial"`
	CNPJ        string   `json:"CNPJ"`
	Endereco    Endereco `json:"endereco"`
}

// HistoricoEscolar will hold all information about a single historic.
// This strict is a primary struct in world state.
// There reference for this struct in world state is a composite key of [Aluno/Curso/IesEmissora/DataHoraEmissao/Curriculo].
type HistoricoEscolar struct {
	Aluno                          string            `json:"aluno"`
	Curso                          string            `json:"curso"`
	IesEmissora                    string            `json:"iesEmissora"`
	DataHoraEmissao                time.Time         `json:"dataHoraEmissao"`
	SituacaoAtualDiscente          Situacao          `json:"situacaoAtualDiscente"`
	ENADE                          ENADE             `json:"ENADE"`
	Curriculo                      string            `json:"curriculo"`
	IngressoCurso                  IngressoCurso     `json:"ingressoCurso"`
	ElementoHistorico              ElementoHistorico `json:"elementoHistorico"`
	CodigoValidacao                string            `json:"codigoValidacao"`
	DigestValue                    string            `json:"digestValue"`
	CargaHorariaCurso              CargaHoraria      `json:"cargaHorariaCurso"`
	CargaHorariaCursoIntegralizada CargaHoraria      `json:"cargaHorariaCursoIntegralizada"`
	NomeParaAreas                  string            `json:"nomeParaAreas"`         // mini-occurs: 0
	Areas                          []Area            `json:"areas"`                 // mini-occurs: 0
	InformacoesAdicionais          string            `json:"informacoesAdicionais"` // mini-occurs: 0
}

// Area will hold all information about a single area.
type Area struct {
	Codigo string `json:"codigo"`
	Nome   string `json:"nome"`
}

// ENADE will hold all ENADE habilitations for a single history.
type ENADE struct {
	Habilitacoes    []HabilitacaoEnade `json:"habilitacoes"`
	NaoHabilitacoes []HabilitacaoEnade `json:"naoHabilitacoes"`
	Irregulares     []HabilitacaoEnade `json:"irregulares"`
}

// tipoMotivoENADE.
type tipoMotivoENADE string

// All possibilities to Motivo in HabilitacaoENADE Struct.
const (
	ECicloAvaliativo   tipoMotivoENADE = "Estudante não habilitado ao Enade em razão do calendário do ciclo avaliativo"
	EProjetoPedagogico tipoMotivoENADE = "Estudante não habilitado ao Enade em razão da natureza do projeto pedagógico do curso"
)

// tipoCondicaoENADE.
type tipoCondicaoENADE string

// All possibilities to Condicao in HabilitacaoENADE Struct.
const (
	EIngressante tipoCondicaoENADE = "Ingressante"
	EConcluinte  tipoCondicaoENADE = "Concluinte"
)

// HabilitacaoEnade will hold all information about a habiltation on ENADE.
// If Habilitado value is set to false, the attribute Motivo or OutroMotivo need to be assigned.
// The attribute Edicao represent a year, for example: 2022.
type HabilitacaoEnade struct {
	Condicao    tipoCondicaoENADE `json:"condicao"`
	Edicao      int               `json:"edicao"`
	Motivo      tipoMotivoENADE   `json:"motivo"`
	OutroMotivo string            `json:"outroMotivo"`
	Habilitado  bool              `json:"habilitado"`
}

// TipoDeSituacao Enum.
type TipoDeSituacao string

// All possibilities to TipoDeSituacao in Situacao.
const (
	EDesistencia           TipoDeSituacao = "Desistencia"
	ELicenca               TipoDeSituacao = "Licenca"
	EMatriculaEmDisciplina TipoDeSituacao = "MatriculadoEmDisciplina"
	EOutraSituacao         TipoDeSituacao = "OutraSituacao"
	ETrancamento           TipoDeSituacao = "Trancamento"
	EFormado               TipoDeSituacao = "Formado"
	EIntercambio           TipoDeSituacao = "Intercambio"
	EJubilado              TipoDeSituacao = "Jubilado"
)

// Situacao Will hold all information about a single situation.
// If TipoDeSituacao is set to "Formado", the attribute Formado need to be assigned.
// Else if TipoDeSituacao is set to "Intercambio", the attribute Intercambio need to be assigned.
type Situacao struct {
	PeriodoLetivo  string         `json:"periodoLetivo"`
	TipoDeSituacao TipoDeSituacao `json:"tipoDeSituacao"`
	Formado        Formado        `json:"formado"`
	Intercambio    Intercambio    `json:"intercambio"`
}

// Formado will hold informations about a student after graduate.
type Formado struct {
	DataConclusaoCurso   time.Time `json:"dataConclusaoCurso"`
	DataColacaoGrau      time.Time `json:"dataColacaoGrau"`
	DataExpedicaoDiploma time.Time `json:"dataExpedicaoDiploma"`
}

// tipoDeIntercambio.
type tipoDeIntercambio string

// All possibilities to attribute TipoDeIntercambio in Intercambio.
const (
	EInternacional tipoDeIntercambio = "Internacional"
	ENacional      tipoDeIntercambio = "Nacional"
)

// Intercambio will hold informations about a possible student exchange.
type Intercambio struct {
	Instituicao         string            `json:"instituicao"`
	Pais                string            `json:"pais"`
	ProgramaIntercambio string            `json:"programaIntercambio"`
	TipoDeIntercambio   tipoDeIntercambio `json:"tipoDeIntercambio"`
}

// TipoFormaDeAcesso Enum.
type TipoFormaDeAcesso string

// All possibilities to attribute FormaDeAcesso in IngressoCurso.
const (
	EPrograma         TipoFormaDeAcesso = "Programas de avaliação seriada ou continuada"
	EConvenios        TipoFormaDeAcesso = "Convenio"
	EHistoricoEScolar TipoFormaDeAcesso = "Historico Escolar"
	ESisu             TipoFormaDeAcesso = "Sisu"
	EVestibular       TipoFormaDeAcesso = "Vestibular"
	EEntrevista       TipoFormaDeAcesso = "Entrevista"
	ETransferencia    TipoFormaDeAcesso = "Transferencia"
	EOutros           TipoFormaDeAcesso = "Outros"
)

// IngressoCurso will hold informations about when and how the student entered the IES.
type IngressoCurso struct {
	Data          time.Time         `json:"data"`
	FormaDeAcesso TipoFormaDeAcesso `json:"formaDeAcesso"`
}

// ElementoHistorico Will hold all situations, disciplines, activites and internships from a specfic historic.
type ElementoHistorico struct {
	Situacoes                []Situacao              `json:"situacoes"`
	Disciplinas              []Disciplina            `json:"disciplinas"`
	AtividadesComplementares []AtividadeComplementar `json:"atividadeComplementares"`
	Estagios                 []Estagio               `json:"estagios"`
}

// CategoriaAtividadeComplementar represent a list of Atividades Complementares from the same category.
type CategoriaAtividadeComplementar struct {
	Codigo                          string                            `json:"codigo"`
	Nome                            string                            `json:"nome"`
	LimiteCargaHorariaEmHoraRelogio CargaHoraria                      `json:"limiteCargaHorariaEmHoraRelogio"`
	AtividadesComplementares        []AtividadeComplementarCurricular `json:"atividadesComplementares"`
}

// AtividadeComplementar will hold all informations about a atividade complementar.
type AtividadeComplementar struct {
	Codigo                            string       `json:"codigo"`
	DataInicio                        time.Time    `json:"dataInicio"`
	DataFim                           time.Time    `json:"dataFim"`
	DataRegistro                      time.Time    `json:"dataRegistro"`
	TipoAtividadeComplementar         string       `json:"tipoAtividadeComplementar"`
	Descricao                         string       `json:"descricao"`
	CargaHorariaEmHoraRelogio         CargaHoraria `json:"cargaHorariaEmHoraRelogio"`
	DocentesResponsaveisPelaValidacao []Docente    `json:"docentesResponsaveisPelaValidacao"`
}

// AtividadeComplementarCurricular will hold informations about an atividade complementar.
type AtividadeComplementarCurricular struct {
	Codigo                          string       `json:"codigo"`
	Nome                            string       `json:"nome"`
	Descricao                       string       `json:"descricao"`
	LimiteCargaHorariaEmHoraRelogio CargaHoraria `json:"limiteCargaHorariaEmHoraRelogio"`
}

// Estagio will hold informations about a estagio.
type Estagio struct {
	DataInicio                 time.Time    `json:"dataInicio"`
	DataFim                    time.Time    `json:"dataFim"`
	Descricao                  string       `json:"descricao"`
	Concedente                 Concedente   `json:"concedente"`
	CargaHorariaEmHorasRelogio CargaHoraria `json:"cargaHorariaEmHorasRelogio"`
	DocentesOrientadores       []Docente    `json:"docentesOrientadores"`
}

// Concedente will hold informations about concedente.
type Concedente struct {
	RazaoSocial  string `json:"razaoSocial"`
	NomeFantasia string `json:"nomeFantasia"`
	CNPJ         string `json:"CNPJ"`
}

// XMLog will hold a student Id and one array containing all his XML registers.
type XMLog struct {
	ID       string
	XMLArray []string
}

/*
#########################
### Dados Reguladores ###
#########################
*/


// DadosRegulatoriosIES will hold all informations about regulation data from an IES.
// Struct definition:
// InstituicaoEnsino will have informations about an IES regulation data.
// HistoricoAtos will hold the regulation acts historic about that IES.
type DadosRegulatoriosIES struct {
	InstituicaoEnsino InformacoesRegulatoriasIES `json:"insituicaoEnsino"`
	HistoricoAtos     HistoricoAtosIES           `json:"historicoAtos"`
}

// InformacoesRegulatoriasIES will hold all informations about an IES regulation data.
type InformacoesRegulatoriasIES struct {
	Nome                                string                 `json:"nome"`
	Sigla                               string                 `json:"sigla"`
	CodigoMEC                           string                 `json:"codigoMEC"`
	CNPJ                                string                 `json:"CNPJ"`
	Endereco                            Endereco               `json:"endereco"`
	MantenedoraRegulatoria              MantenedoraRegulatoria `json:"mantenedora"` //min-occur =0
	CategoriaOrganizacaoAcademica       CategoriaOrganizacao   `json:"categoriaOrganizacaoAcademica"`
	Ativa                               bool                   `json:"ativa"`
	PossuiPrerrogativaDeRegistroInterno bool                   `json:"possuiPrerrogativaDeRegistroInterno"`
	PossuiPrerrogativaDeRegistroExterno bool                   `json:"possuiPrerrogativaDeRegistroExterno"`
	CredenciadaParaCursosPresenciais    bool                   `json:"credenciadaParaCursosPresenciais"`
	CredenciadaParaCursosEAD            bool                   `json:"credenciadaParaCursosEAD"`
}

// CategoriaOrganizacao will hold which categoria that IES is described.
type CategoriaOrganizacao string

// Possibilities to categoriaOrganizacao
const (
	EFaculdade           CategoriaOrganizacao = "Faculdade"
	ECentroUniversitario CategoriaOrganizacao = "Centro Universitário"
	EUniversidade        CategoriaOrganizacao = "Universidade"
)

// MantenedoraRegulatoria will hold informations about mantenedora regulatoria.
type MantenedoraRegulatoria struct {
	RazaoSocial string   `json:"razaoSocial"`
	CodigoMEC   string   `json:"codigoMEC"`
	CNPJ        string   `json:"CNPJ"`
	Endereco    Endereco `json:"endereco"`
}

// HistoricoAtosIES will hold the list of regulation acts historic about that IES.
type HistoricoAtosIES struct {
	Credenciamento      []InformacoesAtoRegulatorio `json:"credenciamento"`      //min-occur = 0
	CredenciamentoEAD   []InformacoesAtoRegulatorio `json:"credenciamentoEAD"`   //min-occur = 0
	Recredenciamento    []InformacoesAtoRegulatorio `json:"recredenciamento"`    //min-occur = 0
	CriacaoPoloEAD      []InformacoesAtoRegulatorio `json:"criacaoPoloEAD"`      //min-occur = 0
	RecredenciamentoEAD []InformacoesAtoRegulatorio `json:"recredenciamentoEAD"` //min-occur = 0
	Descredenciamento   []InformacoesAtoRegulatorio `json:"descredenciamento"`   //min-occur = 0
}

// InformacoesAtoRegulatorio will hold informations about one regulation act.
type InformacoesAtoRegulatorio struct {
	Vigente                   bool           `json:"vigente"`       //min-occur = 0
	PrazoValidade             time.Time      `json:"prazoValidade"` //min-occur = 0
	InformacoesAtoRegulatorio AtoRegulatorio `json:"informacoesAtoRegulatorio"`
}

// DadosRegulatoriosCurso will hold all informations about regulation data from a course.
// Struct definition:
// InstituicaoEnsino will have minimum data about the IES that have the course.
// CursoGraduacao will hold all regulation informations about the course.
// HistoricoAtos will hold the regulation acts historic about that course.
type DadosRegulatoriosCurso struct {
	InstituicaoEnsino DadosMinimosIES    `json:"instituicaoEnsino"`
	CursoGraduacao    CursoGraduacao     `json:"cursoGraduacao"`
	HistoricoAtos     HistoricoAtosCurso `json:"historicoAtos"`
}

// DadosMinimosIES have minimum data about the IES that have the course.
type DadosMinimosIES struct {
	Nome        string                 `json:"nome"`
	CodigoMEC   string                 `json:"codigoMEC"`
	CNPJ        string                 `json:"CNPJ"`
	Mantenedora MantenedoraRegulatoria `json:"mantenedora"` //min-occurr =0
}

// CursoGraduacao will hold all regulation informations about the course.
type CursoGraduacao struct {
	Nome                         string                      `json:"nome"`
	CodigoCursoEMEC              string                      `json:"codigoCursoEMEC"`
	TramitacaoEMEC               TramitacaoMEC               `json:"tramitacaoMEC"`
	EnderecoCurso                Endereco                    `json:"enderecoCurso"`
	Grau                         GrauConferido               `json:"grau"`
	Modalidade                   ModalidadeCurso             `json:"modalidade"`
	Polo                         PoloInformacoesRegulatorias `json:"polo"` //min-ocurr =0
	DataFuncionamento            time.Time                   `json:"dataFuncionamento"`
	CargaHoraria                 int                         `json:"cargaHoraria"`
	Periodicidade                PeriodicidadeCurso          `json:"periodicidade"`
	Integralizacao               int                         `json:"integralizacao"`
	EmAtividade                  bool                        `json:"emAtividade"`
	NumeroVagasAnuaisAutorizadas int                         `json:"numeroVagasAnuaisAutorizadas"` //positive number
}

// PoloInformacoesRegulatorias will hold informations about the polo.
type PoloInformacoesRegulatorias struct {
	Nome                         string        `json:"nome"`
	Endereco                     Endereco      `json:"endereco"`
	CodigoEMEC                   string        `json:"codigoEMEC"`
	SemCodigoEMEC                TramitacaoMEC `json:"semCodigoEMEC"`
	NumeroVagasAnuaisAutorizadas int           `json:"numeroVagasAnuaisAutorizadas"` //positive number
}

// GrauConferido will hold which grau the student will have after graduate.
type GrauConferido string

// Possibilities to Grau
const (
	ETecnologo       GrauConferido = "Tecnólogo"
	EBacharelado     GrauConferido = "Bacharelado"
	ELicenciatura    GrauConferido = "Licenciatura"
	ECursoSequencial GrauConferido = "Curso sequencial"
)

// ModalidadeCurso will hold which modalidade the course have.
type ModalidadeCurso string

// Possibilities to Modalidade
const (
	EPresencial ModalidadeCurso = "Presencial"
	EEAD        ModalidadeCurso = "EAD"
)

// PeriodicidadeCurso will hold which periodicidade the course have.
type PeriodicidadeCurso string

// Possibilities to Periodicidade
const (
	EIntegral      PeriodicidadeCurso = "Integral"
	ESemestral     PeriodicidadeCurso = "Semestral"
	EAnual         PeriodicidadeCurso = "Anual"
	EModular       PeriodicidadeCurso = "Modular"
	ETrimestral    PeriodicidadeCurso = "Trimestral"
	EQuadrimestral PeriodicidadeCurso = "Quadrimestral"
)

// HistoricoAtosCurso will hold the list of regulation acts historic about that course.
type HistoricoAtosCurso struct {
	Autorizacao                []InformacoesAtoRegulatorio `json: "autorizacao"`
	AutorizacaoEAD             []InformacoesAtoRegulatorio `json: "autorizacaoEAD"`
	Reconhecimento             []InformacoesAtoRegulatorio `json: "reconhecimento"`
	ReconhecimentoEAD          []InformacoesAtoRegulatorio `json: "reconhecimentoEAD"`
	RenovacaoReconhecimento    []InformacoesAtoRegulatorio `json: "renovacaoReconhecimento"`
	RenovacaoReconhecimentoEAD []InformacoesAtoRegulatorio `json: "renovacaoReconhecimentoEAD"`
}