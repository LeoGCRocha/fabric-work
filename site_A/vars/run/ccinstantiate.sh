#!/bin/bash
# Script to instantiate chaincode
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ID=cli
export CORE_PEER_ADDRESS=192.168.0.26:7055
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgA.com/peers/peer1.orgA.com/tls/ca.crt
export CORE_PEER_LOCALMSPID=orgA-com
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgA.com/users/Admin@orgA.com/msp
export ORDERER_ADDRESS=192.168.0.26:7056
export ORDERER_TLS_CA=/vars/keyfiles/ordererOrganizations/orgA-orderer.com/orderers/orderer1.orgA-orderer.com/tls/ca.crt

peer chaincode invoke -o $ORDERER_ADDRESS --isInit \
  --cafile $ORDERER_TLS_CA --tls -C jornada -n academicRecords \
  --peerAddresses 192.168.0.26:7054 \
  --tlsRootCertFiles /vars/keyfiles/peerOrganizations/orgA.com/peers/peer2.orgA.com/tls/ca.crt \
  --peerAddresses 192.168.0.26:7053 \
  --tlsRootCertFiles /vars/keyfiles/peerOrganizations/orgX.com/peers/peer1.orgX.com/tls/ca.crt \
  -c '{"Args":[  ]}' --waitForEvent
