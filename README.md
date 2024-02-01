go-udp
======

A simple client/server UDP implemetation using Golang stdlib. This is not meant to be used for anything serious in production, just as an example of the protocol.

This was mainly written to help triage networking issues while setting up various docker containers.

## Example run

Here is an example run. You can see that the server startes

### server  

```
.\server.exe
[udp-server] 2024/02/01 10:50:05 listening on addr=[::]:59033 with block size=1024
[udp-server] 2024/02/01 10:51:15 <127.0.0.1:62886> asdfasd
[udp-server] 2024/02/01 10:51:15 <127.0.0.1:62886> fas
[udp-server] 2024/02/01 10:51:15 <127.0.0.1:62886> dfa
[udp-server] 2024/02/01 10:51:16 <127.0.0.1:62886> sdf
[udp-server] 2024/02/01 10:51:16 <127.0.0.1:62886> asdf
[udp-server] 2024/02/01 10:51:19 <127.0.0.1:62886>
[udp-server] 2024/02/01 10:51:19 <127.0.0.1:62886>
[udp-server] 2024/02/01 10:51:20 <127.0.0.1:62886> a
[udp-server] 2024/02/01 10:51:20 <127.0.0.1:62886> sdf
[udp-server] 2024/02/01 10:51:20 <127.0.0.1:62886> asd
[udp-server] 2024/02/01 10:51:20 <127.0.0.1:62886> fa
[udp-server] 2024/02/01 10:51:21 <127.0.0.1:62886> sdfasdfkjhsdjkfhalskdjhflaskdfa
```

### client

```
PS C:\Users\partkyle\Code\go-udp> .\client.exe -port 59033
[udp-server] 2024/02/01 10:51:13 sending from 0.0.0.0:0 to 0.0.0.0:59033
[udp-server] 2024/02/01 10:51:13 using stdin for payload
asdfasd
fas
dfa
sdf
asdf


a
sdf
asd
fa
sdfasdfkjhsdjkfhalskdjhflaskdfa
```
