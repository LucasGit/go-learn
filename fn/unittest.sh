#!/bin/bash


test_report() {
	return

	echo "=========================="
	if [[ $? != 0 ]]; then
		echo "FAILED"
		exit 1
	fi
	echo "PASS"
}

echo -e "\n=> run all unit test"
file="fn"
go test -v ${file}_test.go ${file}.go
test_report

file="fn_recursive"
go test -v ${file}_test.go ${file}.go
test_report

file="fn_anonymous"
go test -v ${file}_test.go ${file}.go
test_report


exit 0
echo -e "\n=> single function run"
go test -v  -run TestHello ${file}_test.go ${file}.go
test_report

go test -v  -run TestMax_val ${file}_test.go ${file}.go
test_report




