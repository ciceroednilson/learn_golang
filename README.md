
# Install Golang.


https://go.dev/learn/


# Linux - execute the command after the download go1.20.3.linux-amd64.tar.gz.


rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz



# Permission.


sudo chmod -R 777 /usr/local/go/


# Show version.


go version


# Create/Start a module.

 
go mod init name-module


# Build.

go build

# Example of download of dependency.


go get github.com/badoux/checkmail


# Remove unnecessary dependency.


go mod tidy


# Execute an application.


go run file_main.go



# Execute in the root folder to execute all unit test.


go test ./...


# Execute the unit test as detail.


go test -v ./...


# Execute the unit test as coverage.


go test --cover  ./...

go tool conver --html=cobertura.txt
