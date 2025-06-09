# easytier-ffi go with ebitengine/purego

## running

* build easytier-ffi library & move to lib dir

* running test

app.yaml 

```code
instance_name = "easytier"
instance_id = "2d97d178-8f67-4003-b286-6d47fa9dbed3"
dhcp = true
listeners = [
    "tcp://0.0.0.0:11010",
    "udp://0.0.0.0:11010",
    "wg://0.0.0.0:11011",
]
rpc_portal = "0.0.0.0:0"

[network_identity]
network_name = "xxxx"
network_secret = "xxxxx"

[[peer]]
uri = "tcp://xxxxx:11010"

[flags]
```
do demo

```code
sudo go run main.go
```

