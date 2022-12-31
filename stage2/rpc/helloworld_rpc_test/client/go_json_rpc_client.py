import json
import socket #不使用 request 包，因为 socket 不带 http header 是纯文本的，request 请求会带 header 等信息

request = {
    "id":0,
    "params":["winnie"],
    "method":"HelloServer.Hello"
}

client = socket.create_connection(("localhost",1234))
client.sendall(json.dumps(request).encode())

# 获取服务器返回的数据
rsp = client.recv(1024)
rsp = json.loads(rsp.decode())

print(rsp)
