#!/bin/bash
# Script to instantiate chaincode
cp $FABRIC_CFG_PATH/core.yaml /vars/core.yaml
cd /vars
export FABRIC_CFG_PATH=/vars

# Need to set to order admin to update channel stuff by default
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/ordererOrganizations/orgA-orderer.com/orderers/orderer1.orgA-orderer.com/tls/ca.crt
export CORE_PEER_LOCALMSPID=orgA-orderer-com
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/ordererOrganizations/orgA-orderer.com/users/Admin@orgA-orderer.com/msp
export ORDERER_TLS_CA=/vars/keyfiles/ordererOrganizations/orgA-orderer.com/orderers/orderer1.orgA-orderer.com/tls/ca.crt
export ORDERER_ADDRESS=192.168.0.26:7056

if [ -f "jornada_update_envelope.pb" ]; then
# Now finally submit the channel update tx
  peer channel update -f jornada_update_envelope.pb \
    -c jornada -o $ORDERER_ADDRESS --cafile $ORDERER_TLS_CA --tls
else
  echo "No channel configuration update envelop found, do channel sign first."
  exit 1
fi