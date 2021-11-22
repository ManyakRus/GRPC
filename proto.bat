set DST_DIR=C:\Users\nikitin.a\GolandProjects\GRPC
set SRC_DIR=%DST_DIR%\proto

echo %DST_DIR%

C:\Users\nikitin.a\GolandProjects\protoc-3.19.1-win64\bin\protoc.exe -I=%SRC_DIR% --go_out=%DST_DIR% %SRC_DIR%\greeter.proto --go-grpc_out=. 

rem C:\Users\nikitin.a\GolandProjects\protoc-3.19.1-win64\bin\protoc.exe -I=%SRC_DIR% --go_out=%DST_DIR% %SRC_DIR%\greeter.proto --go-grpc_out=.  --go_opt=. 