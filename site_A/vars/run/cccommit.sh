#!/bin/bash
# Script to instantiate chaincode
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ID=cli
export CORE_PEER_ADDRESS=192.168.0.26:7054
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgA.com/peers/peer2.orgA.com/tls/ca.crt
export CORE_PEER_LOCALMSPID=orgA-com
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgA.com/users/Admin@orgA.com/msp
export ORDERER_ADDRESS=192.168.0.26:7056
export ORDERER_TLS_CA=/vars/keyfiles/ordererOrganizations/orgA-orderer.com/orderers/orderer1.orgA-orderer.com/tls/ca.crt
SID=$(peer lifecycle chaincode querycommitted -C jornada -O json \
  | jq -r '.chaincode_definitions|.[]|select(.name=="academicRecords")|.sequence' || true)

if [[ -z $SID ]]; then
  SEQUENCE=1
else
  SEQUENCE=$((1+$SID))
fi

peer lifecycle chaincode commit -o $ORDERER_ADDRESS --channelID jornada \
  --name academicRecords --version 0.1 --sequence $SEQUENCE \
  --peerAddresses 192.168.0.26:7052 \
  --tlsRootCertFiles /vars/discover/jornada/orgX-com/tlscert \
  --init-required \
  --cafile $ORDERER_TLS_CA --tls
