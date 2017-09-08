#!/usr/bin/env python
# _*_ coding=utf-8 _*_

import requests
import json

url = "http://127.0.0.1:8090/users"

req = requests.get(url)
data = json.loads(req.text)

#print(data)
print(req.status_code)
print(data["code"])

for user in data["data"]:
    print(user["id"], user["name"], user["note"], user["isadmin"])

