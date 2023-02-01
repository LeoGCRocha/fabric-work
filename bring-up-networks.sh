cd site_A/
minifab up -o orgA.com -e 7050  -n academicRecords -p ''
cd ../site_B/
minifab netup -o orgB.com -e 7070
cp vars/JoinRequest_orgB-com.json ../site_A/vars/NewOrgJoinRequest.json
cd ../site_A/
minifab orgjoin
cp vars/profiles/endpoints.yaml ../site_B/vars/
cd ../site_B/
minifab nodeimport,join -c jornada
minifab install,approve
cd ../site_A/
minifab approve,discover,commit
cd ../site_C/
minifab netup -o orgC.com -e 7080
cp vars/JoinRequest_orgC-com.json ../site_A/vars/NewOrgJoinRequest.json
cd ../site_A/
minifab channelquery,configmerge,channelsign
sudo cp vars/jornada_update_envelope.pb ../site_B/vars/
cd ../site_B/
minifab channelsignenvelope
sudo cp vars/jornada_update_envelope.pb ../site_A/vars/
cd ../site_A/
minifab channelupdate
cp vars/profiles/endpoints.yaml ../site_C/vars/
cd ../site_C/
minifab nodeimport,join -c jornada
minifab install,approve
cd ../site_B/
minifab approve
cd ../site_A/
minifab approve,discover,commit