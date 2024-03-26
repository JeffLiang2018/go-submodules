
# Multiple level dependencies
ipcMsg modules: v0.0.1, v0.0.2
events modules: v0.0.1 (ipcMsg@v0.0.2)
                v0.0.2 (ipcMsg@v0.0.1)

### Case 1
```
.
|-- ipcMsg   (ipcMsg@v0.0.2)
|   |-- go.mod
|   `-- convert.go
`-- events    (events@v0.0.1)
|   |-- go.mod  => (ipcMsg @v0.0.2)
|   `-- searchEngine.go
`-- pcm    (oro@v1.1.1)
|   |-- go.mod => event@v0.0.1, ipcMsg@v0.0.1
|   `-- main.go
```
pcm module would use ipcMsg@v0.0.2


### Case 2
```
.
|-- ipcMsg
|   |-- go.mod
|   `-- convert.go
`-- events
|   |-- go.mod  => (@v0.0.1)
|   `-- searchEngine.go
`-- pcm
|   |-- go.mod => event@v0.0.2, ipcMsg@v0.0.1
|   `-- main.go
```
pcm module would use ipcMsg@v0.0.1

### Case 3
```
.
|-- ipcMsg
|   |-- go.mod
|   `-- convert.go
`-- events
|   |-- go.mod  => (@v0.0.2)
|   `-- searchEngine.go
`-- pcm
|   |-- go.mod => event@v0.0.1, ipcMsg@v0.0.1
|   `-- main.go
```
Need to run 
```shell
go: updates to go.mod needed; to update it:
	go mod tidy
```
After running it would become ipcMsg will use the latest v0.0.2
```
.
|-- ipcMsg
|   |-- go.mod
|   `-- convert.go
`-- events
|   |-- go.mod  => (@v0.0.2)
|   `-- searchEngine.go
`-- pcm
|   |-- go.mod => event@v0.0.1, ipcMsg@v0.0.2
|   `-- main.go
```


### Case 4
```
.
|-- ipcMsg
|   |-- go.mod
|   `-- convert.go
`-- events
|   |-- go.mod  => (@v0.0.1)
|   `-- searchEngine.go
`-- pcm
|   |-- go.mod => event@v0.0.2, ipcMsg@v0.0.2
|   `-- main.go
```
Need to run 
```shell
main.go:5:2: missing go.sum entry for module providing package github.com/JeffLiang2018/go-submodules/events (imported by github.com/JeffLiang2018/go-submodules/pcm); to add:
	go get github.com/JeffLiang2018/go-submodules/pcm
```
pcm call events@v0.0.1 (with ipcMsg@v0.0.2)
(Optional) play with events@v0.0.3 and events@v0.0.4

### Conclusion
If multiple levels dependencies exists, go multiple modules would use the latest version for the leaf dependency.
`Unless, all dependency levels are mapping to the exact version.`
