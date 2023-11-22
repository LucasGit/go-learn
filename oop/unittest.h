
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
file="geometry"
go test -v ${file}_test.go ${file}.go
test_report
