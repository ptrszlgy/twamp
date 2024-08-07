# twamp

TWAMP server and client for go

Client implements https://github.com/tcaine/twamp

Server implements https://github.com/vvidic/go-twamp

Codes are improved to be able to ping TWAMP-Light targets. (EXPERIMENTAL)

# Pinging TWAMP-Light responder

## Run the server and bind to the proper interface

`./twampd -listen 10.10.0.1:862`

## Run the client and specify the server and the TWAMP-Light responder separately

`./twampc -server-ip 10.10.0.1 -responder-ip 192.168.0.20 -responder-port 862 -count 100 10.10.0.1`

# Command line arguments

## Server

```
[me@devel twamp]$ ./twampd -h
Usage of ./twampd:
  -listen string
        listen address (default "localhost:862")
```

## Client

```
[me@devel twamp]$ ./twampc -h
Usage of ./twampc:
  -count uint
        Number of requests to send (0..18446744073709551615 packets, 0 being continuous) (default 5)
  -interval float
        Delay between TWAMP-test requests (seconds). For sub-second intervals, use floating points (default 1)
  -mode string
        Mode of operation (ping, json) (default "ping")
  -port int
        UDP port to send request packets (default 1234)
  -rapid
        Send requests as rapidly as possible (default count of 5, ignores interval and sends next packet as soon as we have a response/timeout)
  -responder-ip string
        Remote reflector IP
  -responder-port int
        UDP port to send request packets
  -server-ip string
        Remote server IP (default "127.0.0.1")
  -server-port int
        Remote host port (default 862)
  -size int
        Size of request packets (0..65468 bytes) (default 42)
  -timeout int
        Maximum wait time for a response for the last packet (seconds). If rapid is set, this is a timeout for every packet (default 1)
  -tos int
        IP type-of-service value (0..255)
```