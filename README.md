# udp-mirror
Listens for UDP packets and echos them back

Original: https://github.com/czerwonk/udp-mirror

# Install
```
go get -u github.com/rob-deutsch/udp-mirror
```
# Application
This tool is helpful as a UDP endpoint

# Use
In this example we want to listen for packets on port 4560, and each will be echoed back.
```
udp-mirror -listen-address ":4560"
```
## Docker
```
docker run -it -p 4560:9999 rob-deutsch/udp-mirror
```