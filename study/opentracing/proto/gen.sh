protoc -I ./ --go_out=plugins=micro,service_name=go.micro.srv.greeter:. ./greeter.proto

ls *.pb.go | xargs -n1 -IX bash -c "sed -e 's/,omitempty//' X > X.tmp && mv X{.tmp,}"