build:
	go build -o bin/kvmux cmd/main.go

install:
	cp bin/kvmux ${HOME}/.local/bin

clean:
	rm bin/kvmux
	rm ${HOME}/.local/bin/kvmux
