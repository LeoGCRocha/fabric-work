#!/bin/bash
cd $FABRIC_CFG_PATH
# cryptogen generate --config crypto-config.yaml --output keyfiles
configtxgen -profile OrdererGenesis -outputBlock genesis.block -channelID systemchannel

configtxgen -printOrg orgA-com > JoinRequest_orgA-com.json
configtxgen -printOrg orgX-com > JoinRequest_orgX-com.json
