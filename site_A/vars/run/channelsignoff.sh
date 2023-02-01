#!/bin/bash
# Script to instantiate chaincode
cp $FABRIC_CFG_PATH/core.yaml /vars/core.yaml
cd /vars
export FABRIC_CFG_PATH=/vars

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ID=cli
export CORE_PEER_ADDRESS=192.168.0.26:7054
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgA.com/peers/peer2.orgA.com/tls/ca.crt
export CORE_PEER_LOCALMSPID=orgA-com
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgA.com/users/Admin@orgA.com/msp
export ORDERER_ADDRESS=192.168.0.26:7056
export ORDERER_TLS_CA=/vars/keyfiles/ordererOrganizations/orgA-orderer.com/orderers/orderer1.orgA-orderer.com/tls/ca.crt

# 1. Fetch the channel configuration
peer channel fetch config config_block.pb -o $ORDERER_ADDRESS \
  --cafile $ORDERER_TLS_CA --tls -c jornada

# 2. Translate the configuration into json format
configtxlator proto_decode --input config_block.pb --type common.Block \
  | jq .data.data[0].payload.data.config > jornada_current_config.json

# 3. Translate the current config in json format to protobuf format
configtxlator proto_encode --input jornada_current_config.json \
  --type common.Config --output config.pb

# 4. Translate the desired config in json format to protobuf format
configtxlator proto_encode --input jornada_config.json \
  --type common.Config --output modified_config.pb

# 5. Calculate the delta of the current config and desired config
configtxlator compute_update --channel_id jornada \
  --original config.pb --updated modified_config.pb \
  --output jornada_update.pb

# 6. Decode the delta of the config to json format
configtxlator proto_decode --input jornada_update.pb \
  --type common.ConfigUpdate | jq . > jornada_update.json

# 7. Now wrap of the delta config to fabric envelop block
echo '{"payload":{"header":{"channel_header":{"channel_id":"jornada", "type":2}},"data":{"config_update":'$(cat jornada_update.json)'}}}' | jq . > jornada_update_envelope.json

# 8. Encode the json format into protobuf format
configtxlator proto_encode --input jornada_update_envelope.json \
  --type common.Envelope --output jornada_update_envelope.pb

# 9. Need to sign channel update envelop by each org admin
export CORE_PEER_LOCALMSPID=orgA-com
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgA.com/peers/peer2.orgA.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgA.com/users/Admin@orgA.com/msp
export CORE_PEER_ADDRESS=192.168.0.26:7054

peer channel signconfigtx -f jornada_update_envelope.pb

export CORE_PEER_LOCALMSPID=orgX-com
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgX.com/peers/peer1.orgX.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgX.com/users/Admin@orgX.com/msp
export CORE_PEER_ADDRESS=192.168.0.26:7053

peer channel signconfigtx -f jornada_update_envelope.pb

