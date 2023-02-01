# Minifabric Acadblock Fork

Este repositório utiliza como base o projeto desenvolvimento pela Hyperledger Labs [Minifabric](https://github.com/hyperledger-labs/minifabric).

## Pré-requisitos para do Minifabric

[Docker](https://www.docker.com/) (18.03 or newer) environment.


## Pré-requisitos para execução das Chaincodes:

| Dependência | Versão |
| :-----------|:-------|
| `golang`    |1.18.3  |
| `docker`    |20.10.17|

## Pré-requisitos para execução das Applications:

| Dependência | Versão |
| :-----------|:-------|
| `node`      |16.10.0 |
| `java`      |11.0.16 |
| `javac`     |11.0.16 |
| `npm`       |7.24.0  |
| `docker`    |20.10.17|

## Modificações Realizadas

### Modificações Gerais:

:pushpin: Atualização do script **minifab**, para referenciar a imagem Docker do projeto (labsec/minifab-acadblock);

:pushpin: Atualização do arquivo de configuração **envsettings**, para que utilize as definições utilizadas no projeto Jornada;

:pushpin: Atualização do arquivo **script/manfuncs.sh**, para adicionar as novas operações (appacademic, appdecree, appxmlog, apptest, channelsignenvelope).

### Modificações nos Playbooks:

Os playbooks, representam as operações que são realizadas por cada um dos comandos. Dessa forma, para adicionar as novas operações foram necessárias adição de novas pastas para configuração e execução: 

:pushpin: **minifabric/playbooks/appacademic**

:pushpin: **minifabric/playbooks/appdecree**

:pushpin: **minifabric/playbooks/appxmlog**

:pushpin: **minifabric/playbooks/appregister**

### Adicição da Operação Channel Sign Enveloped:

Para que fosse possível a geodistribuição dos pares, em múltiplos participantes da rede, foi necessário a adição da operação desenvolvida pelo grupo CT-Blockchain, adicionando o playbook [channelsignenveloped]: (https://github.com/ct-blockchain/minifabric).

### Scripts:

:pushpin: **pulljornada** prepara os arquivos das chaincodes e application do projeto Jornada;

:pushpin: **networkInit** inicializa toda a rede projeto para desenvolvimento, das chaincodes até a interface web;

## Documentação:

Para saber mais sobre Minifabric, veja [docs](https://github.com/hyperledger-labs/minifabric/blob/main/README.md).
