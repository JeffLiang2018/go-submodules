# go-submodules

## go-submodules workflow
(Initialized with a empty repo)
Each step is finished by a pull request in github.
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
Create submodule pcm
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
Create submodule cbs
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
Document the steps updated until oro/v0.0.3 
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

#### Step 5:
Add daemons submodule and debugApi entry
```shell
# add debugApi daemons
git checkout -b <feat-gomodules-daemons>
mkdir debugApi
go mod init github.com/JeffLiang2018/go-submodules/daemons
go mod tidy
cd ..
mkdir debugApi
# create main.go (main package)
cd ..   # go back to daemons directory
go mod tidy
# check in changes and create pull request, then merge it.
# add daemons and entry for debugApi
git checkout main
git branch -d feat-gomodules-daemons    # optional
git pull --rebase origin
git checkout -b feat-gomodule-deamon-2
mkdir daemons
cd daemons
go mod init github.com/JeffLiang2018/go-submodules/daemons
mkdir debugApi
# create main.go (main package)
# check in the changes and create the pull request, then merge it
git checkout main
git branch -d feat-gomodule-daemons    # optional
git tag oro/v0.0.4
git push oro/v0.0.4
```

#### Step 6:
Update ipcMsg submodules and pcm submodules, add oamock
```shell
# git pull --rebase main # to get the latest version
git checkout -b feat-upgrade-ipcmsg
# update go files in ipcMsg submodules
# check in the changes and create the pull request, then merge it
git checkout main
git branch -d feat-upgrade-ipcmsg    # optional
git tag ipcMsg/v0.0.2
git push ipcMsg/v0.0.2
# now update pcm modules to refer to ipcMsg/v0.0.2
git checkout -b feat-upgrade-ipcmsg
# update pcm/go.mod file 
go mod tidy
# create oamock and the main (package main) under daemons 
# check in the changes and create the pull request, then merge it
git branch -d feat-upgrade-icpmsg
git tag oro/v0.0.5
```

#### Step 7:
Spiking on multiple levels dependency
