test:
		PATH=${GOPATH}/bin:${PATH}
		go install github.com/sirkon/ldetool
		go generate github.com/sirkon/ldetool/testing
		which ldetool
		go test -test.v github.com/sirkon/ldetool/testing

