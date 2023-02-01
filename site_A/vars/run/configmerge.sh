#! /bin/bash

jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {(.[1].values.MSP.value.config.name): .[1]}}}}}' \
  /vars/jornada_config.json /vars/NewOrgJoinRequest.json | \
  jq -s '.[0] * {"channel_group":{"groups":{"Application":{"version": (.[0].channel_group.groups.Application.version|tonumber + 1)|tostring }}}}' \
  > /vars/jornada_update_config.json

mv /vars/jornada_update_config.json /vars/jornada_config.json
