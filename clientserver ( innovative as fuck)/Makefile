.PHONY: all client server clean

all: client server

client:
    cd client && go build -o ..\client.exe

server:
    cd server && go build -o ..\server.exe

clean:
    del /F /Q client.exe server.exe