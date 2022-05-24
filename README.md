# Your Application

```

➜  demo4 gogen3 init order
➜  demo4 gogen3 usecase CreateOrder
➜  demo4 gogen3 repository SaveOrder Order CreateOrder
➜  demo4 gogen3 gateway prod
➜  demo4 gogen3 controller restapi
➜  demo4 go mod tidy
go: finding module for package github.com/gin-gonic/gin
go: finding module for package github.com/gin-contrib/cors
go: finding module for package github.com/matoous/go-nanoid
go: found github.com/gin-gonic/gin in github.com/gin-gonic/gin v1.7.7
go: found github.com/gin-contrib/cors in github.com/gin-contrib/cors v1.3.1
go: found github.com/matoous/go-nanoid in github.com/matoous/go-nanoid v1.5.0
➜  demo4 gogen3 application demo4   
➜  demo4 go run main.go 
# golang.org/x/sys/unix
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/syscall_darwin.1_13.go:25:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.1_13.go:27:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.1_13.go:40:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:28:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:43:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:59:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:75:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:90:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:105:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:121:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:121:3: too many errors
➜  demo4 go get golang.org/x/sys/unix
go: upgraded golang.org/x/sys v0.0.0-20200116001909-b77594299b42 => v0.0.0-20220520151302-bc2c85ada10a
➜  demo4 go get golang.org/x/sys/unix
➜  demo4 go run main.go              
You may try 'go run main.go <app_name>' :
 - demo4
➜  demo4 go run main.go demo4
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> demo4/shared/infrastructure/server.NewGinHTTPHandler.func1 (3 handlers)
[GIN-debug] POST   /createorder              --> demo4/domain_order/controller/restapi.(*Controller).createOrderHandler.func1 (5 handlers)
{"appName":"demo4","appInstID":"H747","start":"2022-05-24 15:35:04","severity":"INFO","message":"0000000000000000 server is running at :8080","location":"server.(*GracefullyShutdown).RunWithGracefullyShutdown:40","time":"2022-05-24 15:35:04"}


```