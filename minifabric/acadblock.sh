print_options() {
  printf "Opções disponíveis:
-b : Realiza o build da imagem labsec/minifab-acadblock"
}

sudo echo "Inicializando instalação"

# Get optional parameters
while getopts 'b' flag; do
  case "${flag}" in
    b) build=true ;;
    *) print_options
       exit 1 ;;
  esac
done

# Set path variables, if not already set
if [ -z "$MINIFAB_PATH" ]; then
    MINIFAB_PATH=$(pwd)
    echo "Utilizando $MINIFAB_PATH como caminho para o minifabric"
else 
    echo "Utilizando variável de ambiente MINIFAB_PATH configurada para $MINIFAB_PATH"
fi

cd ..
JORNADA_PATH=$(pwd)

# Checks if there is already a acadblock installation
if [ -d "$MINIFAB_PATH/vars" ]; then
	if [ "$(ls -A $MINIFAB_PATH/vars)" ]; then
        echo "Já existe uma instalação do acadblock em $MINIFAB_PATH. Você deseja removê-la e prosseguir com a instalação? (y/n)"
        read answer
        if [ "$answer" = "n" ]; then
            exit
        else
            cd $MINIFAB_PATH
            ./minifab cleanup
        fi
	fi
fi

sudo apt-get install docker-compose-plugin

# Remove old files
cd $JORNADA_PATH
sudo rm -rf minifabric/chaincode
sudo rm -rf minifabric/app

# Copy chaincodes and applications to minifabric
cp -R chaincode minifabric/chaincode
cp -R application minifabric/app

# Copy errorHandling files to applications
cp -R utils/errorHandling/js/* minifabric/app/academicRecords/main/utils
cp -R utils/errorHandling/js/* minifabric/app/decree/main/utils
if [ ! -d "./minifabric/app/xmlog/main/utils" ]; then
    mkdir minifabric/app/xmlog/main/utils
fi
cp -R utils/errorHandling/js/* minifabric/app/xmlog/main/utils

# Initialize go modules and prepare chaincodes for installation
cp minifabric/modules.sh minifabric/chaincode
cd minifabric/chaincode
docker run --rm -v "$(pwd)":/go/chaincode golang:1.19.3-alpine /bin/sh ./chaincode/modules.sh

cd $JORNADA_PATH

# Copy errorMessages.go to chaincodes
sudo cp utils/errorHandling/go/errorMessages.go minifabric/chaincode/academicRecords/go/vendor/errorMessages
sudo cp utils/errorHandling/go/errorMessages.go minifabric/chaincode/decree/go/vendor/errorMessages
sudo cp utils/errorHandling/go/errorMessages.go minifabric/chaincode/registerBook/go/vendor/errorMessages
sudo cp utils/errorHandling/go/errorMessages.go minifabric/chaincode/XMLog/go/vendor/errorMessages

cd minifabric

if [ "$build" = true ]; then
    echo "Construindo imagem labsec/minifab-acadblock"
    docker build -t labsec/minifab-acadblock .
fi

# Initialize base network
sudo echo Inicializando a rede

cd $MINIFAB_PATH
cp $JORNADA_PATH/minifabric/minifab $MINIFAB_PATH
mkdir $MINIFAB_PATH/vars
cp $JORNADA_PATH/config/spec.yaml $MINIFAB_PATH/vars
cp -R $JORNADA_PATH/minifabric/chaincode $MINIFAB_PATH/vars
cp -R $JORNADA_PATH/minifabric/app $MINIFAB_PATH/vars

cd $MINIFAB_PATH/vars/chaincode
sudo rm -rf Dockerfile
sudo rm -rf modules.sh

cd $MINIFAB_PATH
./minifab up -e true
sudo chmod 777 -R $MINIFAB_PATH

# Instalando chaincodes
echo Instalando chaincodes

cd $MINIFAB_PATH
./minifab ccup -l go -n XMLog -p ''
./minifab ccup -l go -n decree -p ''
# ./minifab ccup -l go -n registerBook -p ''

# Inicializando Applications
echo Inicializando Applications

cd $MINIFAB_PATH
./minifab appacademic
./minifab appdecree

# Instalando APIs
echo Instalando APIs

cd $JORNADA_PATH
cp -R api/utils api/academicRecords/api
cp -R api/utils api/decree/api

cp -R utils/errorHandling/js/* api/academicRecords/api/utils
cp -R utils/errorHandling/js/* api/decree/api/utils

cp docker/api/Dockerfile api/academicRecords
cp docker/api/Dockerfile api/decree
cp docker/web/Dockerfile web

docker compose up --build -d
