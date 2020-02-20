
# demo-camp
A code demo gather.


<!-- vscode-markdown-toc -->
* 1. [elastic](#elastic)
* 2. [envoy](#envoy)
* 3. [Kubernetes](#Kubernetes)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

##  1. <a name='elastic'></a>elastic
> elastic客户端的使用

目录结构

```shell
.
├── go.mod
├── go.sum
├── main.go
└── vendor
    ├── github.com
    └── modules.txt
```

##  2. <a name='envoy'></a>envoy
> envoy Tcp/Auth/Dynamic 配置文件

目录结构

```
.
├── auth
│   └── server
├── bootstarp-tcp.yaml
├── bootstrap-grpc-tio.yml
├── bootstrap-grpc.yml
├── bootstrap-rds.yml
├── bootstrap-tio.yml
├── bootstrap.yml
├── data-plane-api
│   ├── API_OVERVIEW.md
│   ├── API_VERSIONING.md
│   ├── BUILD
│   ├── CONTRIBUTING.md
│   ├── LICENSE
│   ├── README.md
│   ├── STYLE.md
│   ├── bazel
│   ├── ci
│   ├── diagrams
│   ├── docs
│   ├── envoy
│   ├── examples
│   ├── test
│   ├── tools
│   └── xds_protocol.rst
├── docker-compose.yml
├── eds
│   ├── eds
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── vendor
└── rds
    ├── go.mod
    ├── go.sum
    ├── main
    ├── main.go
    ├── rds
    └── vendor
```

##  3. <a name='Kubernetes'></a>Kubernetes
> cli简单用法， go.mod正确