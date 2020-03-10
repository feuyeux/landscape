## A simple redis client demo with golang

### How to config
#### config file
put redis config into `/opt/landscape/conf` as below three lines
```
host
port
password
```
#### build and deploy
```bash
cd landscape
go build
mv landscape /usr/local/bin/
```
### How to use
```bash
$landscape -h
NAME:
   landscape - A simple redis client cli

USAGE:
   landscape [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

AUTHOR:
   Lu Han/LiuWeng <feuyeux@gmail.com>

COMMANDS:
   save, w   write kv to redis
   read, r   read kv from redis
   queue, q  queue commands
   map, q    map commands
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value  redis config file (default: "...")
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)
```

#### kv
```bash
▶ landscape save today 20200309
OK

▶ landscape read today         
20200309
```

#### queue
```bash
▶ landscape queue push x 1     
1

▶ landscape queue push x 2
2

▶ landscape queue push x 3
3

▶ landscape queue pop x   
1
  
▶ landscape queue all x
[2 3]

▶ landscape queue last x
3
```

#### map
```bash                        
▶ landscape map save f1 a1 1
true

▶ landscape map save f1 a2 2
true

▶ landscape map read f1 a1  
1

▶ landscape map all f1    
map[a1:1 a2:2]
```

### How to deploy
```bash
RUN wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz && \
  tar -xzf go1.14.linux-amd64.tar.gz && mv go /usr/local/ && \
  echo "export GOROOT=/usr/local/go" >> /etc/profile && \
  echo "export GOPATH=/home/admin/go_path" >> /etc/profile && \
  echo "export PATH=/usr/local/go/bin:/home/admin/go_path/bin:$PATH" >> /etc/profile && \
  echo "export GOPROXY=http://gomodule-repository.aone.alibaba-inc.com" >> /etc/profile && \
  echo "export GO111MODULE=on" >> /etc/profile && \
  mkdir -p /home/admin/go_path/src && \
  cd /home/admin/go_path/src && \
  git clone https://github.com/feuyeux/landscape.git && \
  cd landscape && \
  source /etc/profile && \
  go build && mv landscape /usr/local/bin/
```