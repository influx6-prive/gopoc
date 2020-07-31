setup:
	mkdir ~/go
	wget https://golang.org/dl/go1.14.6.linux-amd64.tar.gz -O ~/go/go1.14.6.linux-amd64.tar.gz
	cd ~/go && mkdir 1.14.6 && tar -xvf go1.14.6.linux-amd64.tar.gz -C 1.14.6
	export PATH="$HOME/go/1.14.6/bin:$PATH"
	cd -

test:
	go test ./pkg/...

deps:
	go mod download

creditsuisse_cli:
	cd cmd/creditsuisse_mapper/ && go run main.go && cd -
