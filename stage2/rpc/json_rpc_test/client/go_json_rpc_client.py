import json
import socket

# 使用 socket 而不使用 request 包，因为socket 发送纯文本，不会带 http header 那些元素

request = {
    "id":0,
    "params":["winnie"],
    "method":"HelloService.Hello"
}

client = socket.create_connection(("localhost",1234))
client.sendall(json.dumps(request).encode())

# 获取服务器返回的数据
rsp = client.recv(1024)
rsp = json.loads(rsp.decode())

print(rsp["result"])