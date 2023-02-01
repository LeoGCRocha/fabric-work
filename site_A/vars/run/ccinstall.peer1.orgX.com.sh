#!/bin/bash
# Script to install chaincode onto a peer node
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ID=cli
export CORE_PEER_ADDRESS=192.168.0.26:7053
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgX.com/peers/peer1.orgX.com/tls/ca.crt
export CORE_PEER_LOCALMSPID=orgX-com
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgX.com/users/Admin@orgX.com/msp
cd /go/src/github.com/chaincode/academicRecords


if [ ! -f "academicRecords_go_0.1.tar.gz" ]; then
  cd go
  GO111MODULE=on
  go mod vendor
  cd -
  peer lifecycle chaincode package academicRecords_go_0.1.tar.gz \
    -p /go/src/github.com/chaincode/academicRecords/go/ \
    --lang golang --label academicRecords_0.1
fi

peer lifecycle chaincode install academicRecords_go_0.1.tar.gz
