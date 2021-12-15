#!/bin/bash

# reading os type from flags
CURRENT_OS=$1

if [ "${CURRENT_OS}" == "windows-latest" ];then
    extension=.exe
fi
go build -o functional-test$extension
echo 'Building functional-test binary'
go build -o functional-test$extension

echo 'Building Nuclei binary from current branch'
go build -o nuclei_dev$extension ../nuclei

echo 'Installing latest release of nuclei'
go build -o nuclei$extension -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei

echo 'Starting Nuclei functional test'
./functional-test$extension -main ./nuclei$extension -dev ./nuclei_dev$extension -testcases testcases.txt
