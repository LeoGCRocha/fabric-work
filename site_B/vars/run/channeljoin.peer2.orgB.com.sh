#!/bin/bash
# Script to join a peer to a channel
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ID=cli
export CORE_PEER_ADDRESS=192.168.0.26:7071
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/orgB.com/peers/peer2.orgB.com/tls/ca.crt
export CORE_PEER_LOCALMSPID=orgB-com
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/orgB.com/users/Admin@orgB.com/msp
export ORDERER_ADDRESS=192.168.0.26:7056
export ORDERER_TLS_CA=/vars/keyfiles/ordererOrganizations/orgA-orderer.com/tlsca/tlsca.crt
if [ ! -f "jornada.genesis.block" ]; then
  peer channel fetch oldest -o $ORDERER_ADDRESS --cafile $ORDERER_TLS_CA \
  --tls -c jornada /vars/jornada.genesis.block
fi

peer channel join -b /vars/jornada.genesis.block \
  -o $ORDERER_ADDRESS --cafile $ORDERER_TLS_CA --tls
