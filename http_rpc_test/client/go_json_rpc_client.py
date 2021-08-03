import requests
import socket

request = {
    "id":　0,
    "params": ["test"],
    "method": "HelloService.Hello"
}

rsp = requests.post("http://localhost:1234/jsonrpc", json=request)
print(rsp.text)