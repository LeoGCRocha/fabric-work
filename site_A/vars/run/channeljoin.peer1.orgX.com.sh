#!/bin/bash
# Script to join a peer to a channel
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ID=cli
export CORE_PEER_ADDRESS=192.168.0.26:7053
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgX.com/peers/peer1.orgX.com/tls/ca.crt
export CORE_PEER_LOCALMSPID=orgX-com
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgX.com/users/Admin@orgX.com/msp
export ORDERER_ADDRESS=192.168.0.26:7056
export ORDERER_TLS_CA=/vars/keyfiles/ordererOrganizations/orgA-orderer.com/orderers/orderer1.orgA-orderer.com/tls/ca.crt
if [ ! -f "jornada.genesis.block" ]; then
  peer channel fetch oldest -o $ORDERER_ADDRESS --cafile $ORDERER_TLS_CA \
  --tls -c jornada /vars/jornada.genesis.block
fi

peer channel join -b /vars/jornada.genesis.block \
  -o $ORDERER_ADDRESS --cafile $ORDERER_TLS_CA --tls
