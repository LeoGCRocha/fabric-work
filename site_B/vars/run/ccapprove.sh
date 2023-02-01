#!/bin/bash
# Script to approve chaincode
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ID=cli
export CORE_PEER_ADDRESS=192.168.0.26:7072
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgB.com/peers/peer1.orgB.com/tls/ca.crt
export CORE_PEER_LOCALMSPID=orgB-com
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgB.com/users/Admin@orgB.com/msp
export ORDERER_ADDRESS=192.168.0.26:7056
export ORDERER_TLS_CA=/vars/keyfiles/ordererOrganizations/orgA-orderer.com/tlsca/tlsca.crt

peer lifecycle chaincode queryinstalled -O json | jq -r '.installed_chaincodes | .[] | select(.package_id|startswith("academicRecords_0.1:"))' > ccstatus.json

PKID=$(jq '.package_id' ccstatus.json | xargs)
REF=$(jq '.references.jornada' ccstatus.json)

SID=$(peer lifecycle chaincode querycommitted -C jornada -O json \
  | jq -r '.chaincode_definitions|.[]|select(.name=="academicRecords")|.sequence' || true)
if [[ -z $SID ]]; then
  SEQUENCE=1
elif [[ -z $REF ]]; then
  SEQUENCE=$SID
else
  SEQUENCE=$((1+$SID))
fi


export CORE_PEER_LOCALMSPID=orgB-com
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgB.com/peers/peer2.orgB.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgB.com/users/Admin@orgB.com/msp
export CORE_PEER_ADDRESS=192.168.0.26:7071

# approved=$(peer lifecycle chaincode checkcommitreadiness --channelID jornada \
#   --name academicRecords --version 0.1 --init-required --sequence $SEQUENCE --tls \
#   --cafile $ORDERER_TLS_CA --output json | jq -r '.approvals.orgB-com')

# if [[ "$approved" == "false" ]]; then
  peer lifecycle chaincode approveformyorg --channelID jornada --name academicRecords \
    --version 0.1 --package-id $PKID \
  --init-required \
    --sequence $SEQUENCE -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA
# fi
