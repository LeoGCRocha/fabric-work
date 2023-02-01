#!/bin/bash
# Script to approve chaincode
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ID=cli
export CORE_PEER_ADDRESS=192.168.0.26:7054
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgA.com/peers/peer2.orgA.com/tls/ca.crt
export CORE_PEER_LOCALMSPID=orgA-com
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgA.com/users/Admin@orgA.com/msp
export ORDERER_ADDRESS=192.168.0.26:7056
export ORDERER_TLS_CA=/vars/keyfiles/ordererOrganizations/orgA-orderer.com/orderers/orderer1.orgA-orderer.com/tls/ca.crt

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


export CORE_PEER_LOCALMSPID=orgA-com
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgA.com/peers/peer2.orgA.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgA.com/users/Admin@orgA.com/msp
export CORE_PEER_ADDRESS=192.168.0.26:7054

# approved=$(peer lifecycle chaincode checkcommitreadiness --channelID jornada \
#   --name academicRecords --version 0.1 --init-required --sequence $SEQUENCE --tls \
#   --cafile $ORDERER_TLS_CA --output json | jq -r '.approvals.orgA-com')

# if [[ "$approved" == "false" ]]; then
  peer lifecycle chaincode approveformyorg --channelID jornada --name academicRecords \
    --version 0.1 --package-id $PKID \
  --init-required \
    --sequence $SEQUENCE -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA
# fi

export CORE_PEER_LOCALMSPID=orgX-com
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgX.com/peers/peer2.orgX.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgX.com/users/Admin@orgX.com/msp
export CORE_PEER_ADDRESS=192.168.0.26:7052

# approved=$(peer lifecycle chaincode checkcommitreadiness --channelID jornada \
#   --name academicRecords --version 0.1 --init-required --sequence $SEQUENCE --tls \
#   --cafile $ORDERER_TLS_CA --output json | jq -r '.approvals.orgX-com')

# if [[ "$approved" == "false" ]]; then
  peer lifecycle chaincode approveformyorg --channelID jornada --name academicRecords \
    --version 0.1 --package-id $PKID \
  --init-required \
    --sequence $SEQUENCE -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA
# fi
