rm bits_*.go
mv example_test.go example_test
go generate
mv example_test example_test.go
go test
