Using :

1. go env -w GO111MODULE=on
2. go get github.com/wiliehidayat87/rmqp
3. go mod init [inside_your_project_name]
4. go mod tidy
5. go mod vendor

Run :

1. go build -o [your_project_name]
2. ./your_project_name consume
3. ./your_project_name publish "Something in your mind message"
