# go-submodules

## go-submodules workflow
(Initialized with a empty repo)
#### Step 1:
Create submodule ipcMsg
```shell
mkdir icpMsg
cd ipcMsg
go mod init github.com/JeffLiang2018/go-submodules/ipcMsg
# add go files such as convert.go
# check in the codes
git tag ipcMsg/v0.0.1    
git push origin ipcMsg/v0.0.1
# tag should be handled by JenkinsCD file  
```

#### Step 2:
```shell
# go back to go-submodules directory
mkdir pcm
go mod init github.com/JeffLiang2018/go-submodules/pcm
# add go files such as main.go (package main)
go mod tidy
# check in the codes
git tag oro/v0.0.1
git push origin oro/v0.0.1
# tag should be handled by JenkinsCD file
```

#### Step 3:
```shell
# go back to go-submodules directory
git checkout -b <feat-branch>
mkdir cbs
go mod init github.com/JeffLiang2018/go-submodules/cbs
# add go files such as main.go (package main)
go mod tidy
# check in the codes
# Create pull request and merge it: https://github.com/JeffLiang2018/go-submodules/pull/1
git checkout main
git pull --rebase origin
git tag oro/v0.0.2
git push origin oro/v0.0.2
# tag should be handled by JenkinsCD file
```

#### Step 4:
```shell
git checkout -b <doc-feat-branch>
# edit README.md
# check in README.md
# Create pull request and merge it: https://github.com/JeffLiang2018/go-submodules/pull/1
git checkout main
git pull --rebase origin
git tag oro/v0.0.3
git push origin oro/v0.0.3
# tag should be handled by JenkinsCD file
```
