#!/bin/bash
# Script to discover endorsers and channel config
cd /vars

export PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgA.com/users/Admin@orgA.com/tls/ca.crt
export ADMINPRIVATEKEY=/vars/keyfiles/peerOrganizations/orgA.com/users/Admin@orgA.com/msp/keystore/priv_sk
export ADMINCERT=/vars/keyfiles/peerOrganizations/orgA.com/users/Admin@orgA.com/msp/signcerts/Admin@orgA.com-cert.pem

discover endorsers --peerTLSCA $PEER_TLS_ROOTCERT_FILE \
  --userKey $ADMINPRIVATEKEY \
  --userCert $ADMINCERT \
  --MSP orgA-com --channel jornada \
  --server 192.168.0.26:7054 \
  --chaincode academicRecords | jq '.[0]' | \
  jq 'del(.. | .Identity?)' | jq 'del(.. | .LedgerHeight?)' \
  > /vars/discover/jornada_academicRecords_endorsers.json

discover config --peerTLSCA $PEER_TLS_ROOTCERT_FILE \
  --userKey $ADMINPRIVATEKEY \
  --userCert $ADMINCERT \
  --MSP orgA-com --channel jornada \
  --server 192.168.0.26:7054 > /vars/discover/jornada_config.json
