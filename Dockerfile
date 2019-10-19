FROM golang:1.13
ENV APP_DIR $GOPATH/src/basic_blog_go

RUN go get github.com/astaxie/beego && go get github.com/beego/bee
RUN mkdir -p $APP_DIR


ADD . $APP_DIR
WORKDIR $APP_DIR

RUN go get ./
