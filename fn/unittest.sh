#!/bin/bash

file=fn

test_report() {
	echo "=========================="
	if [[ $? != 0 ]]; then
		echo "FAILED"
		exit 1
	fi
	echo "PASS"
}

echo -e "\n=> run all unit test"
go test -v ${file}_test.go ${file}.go
test_report


# 
echo -e "\n=> single function run"
go test -v  -run TestHello ${file}_test.go ${file}.go
test_report

go test -v  -run TestMax_val ${file}_test.go ${file}.go
test_report




