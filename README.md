## A simple redis client demo with golang

#### prepare
put redis config into `/opt/landscape/conf`
```
host
port
password
```

#### kv
```bash
▶ ./landscape save today 20200309
OK

▶ ./landscape read today         
20200309
```

#### queue
```bash
▶ ./landscape queue push x 1     
1

▶ ./landscape queue push x 2
2

▶ ./landscape queue push x 3
3

▶ ./landscape queue pop x   
1
  
▶ ./landscape queue all x
[2 3]

▶ ./landscape queue last x
3
```

#### map
```bash                        
▶ ./landscape map save f1 a1 1
true

▶ ./landscape map save f1 a2 2
true

▶ ./landscape map read f1 a1  
1

▶ ./landscape map all f1    
map[a1:1 a2:2]
```