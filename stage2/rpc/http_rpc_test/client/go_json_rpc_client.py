import requests

request= {
    "id":0,
    "params":["jiazong"],
    "method":"HelloService.Hello"
}

rsp = requests.post("http://localhost:1234/jsonrpc",json=request)
print(rsp.text)