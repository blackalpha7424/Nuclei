#!/bin/bash

cleanup(){
  killall nuclei-server
}
cd ../nuclei-server
go build
mv nuclei-server ../integration-test-client/nuclei-server
cd ../nuclei-client
go build
mv nuclei-client ../integration-test-client/nuclei-client
cd ../integration-test-client
go build
cleanup
#running nuclei api server
./nuclei-server &
#delay the test cases
sleep 2
# #running tests
./integration-test-client -url "http://localhost:8822/api/v1"

if [ $? -eq 0 ]
then
  cleanup
  exit 0
else
  cleanup
  exit 1
fi

