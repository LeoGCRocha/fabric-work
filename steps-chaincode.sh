##
# site A
##

minifab up -o orgA.com -e 7050  -n academicRecords -p ''

##
# site B
##
minifab netup -o orgB.com -e 7070

cp vars/JoinRequest_orgB-com.json ../site_A/vars/NewOrgJoinRequest.json

##
# site A
##
# 
minifab orgjoin

cp vars/profiles/endpoints.yaml ../site_B/vars/

##
# site B
##
minifab nodeimport,join -c jornada

### add academicRecords to site B

minifab install,approve 

##
# site A
##
minifab approve,discover,commit

###
#### add site C
####

##
# site C
##
minifab netup -o orgC.com -e 7080

cp vars/JoinRequest_orgC-com.json ../site_A/vars/NewOrgJoinRequest.json

##
# site A
##
# 
minifab channelquery,configmerge,channelsign

sudo cp vars/jornada_update_envelope.pb ../site_B/vars/

##
# site B
##
minifab channelsignenvelope

sudo cp vars/jornada_update_envelope.pb ../site_A/vars/


##
# site A
##
# 
minifab channelupdate

cp vars/profiles/endpoints.yaml ../site_C/vars/

##
# site C
##
minifab nodeimport,join -c jornada

##
# add academicRecords to site C
##
minifab install,approve

##
# site B
##
minifab approve

##
# site A
##
minifab approve,discover,commit

##
# site C and site B
##
## verify new configuration to chaincode
minifab discover

# todo: work until here

##
# site A
##
minifab install -n student -v 0.2

## 
# site B
##
minifab install -n student -v 0.2 
minifab approve

##
# site A
##
minifab discover

##
# site A
##
minifab install -n decree
minifab approve 

##
# site B
## 
minifab approve -n decree

##
# site C
## 
minifab approve -n decree

##
# site A
##
minifab commit,discover