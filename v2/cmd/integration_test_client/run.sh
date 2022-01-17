#!/bin/bash

cleanup(){
  killall nuclei-server
}
cd ../nuclei-server
go build
mv nuclei-server ../integration_test_client/nuclei-server
cd ../nuclei-client
go build
mv nuclei-client ../integration_test_client/nuclei-client
cd ../integration_test_client
go build
cleanup
#running nuclei api server
./nuclei-server &
#delay the test cases
sleep 2
# #running tests
./integration_test_client -url "http://localhost:8822/api/v1"

if [ $? -eq 0 ]
then
  cleanup
  exit 0
else
  cleanup
  exit 1
fi

