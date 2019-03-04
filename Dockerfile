FROM golang:1.12.0-alpine


ENV ROOTPATH=$GOPATH/src/github.com/tuterdust/my-todo-list/src
RUN mkdir -p $GOPATH/src/github.com/tuterdust/my-todo-list
ADD . $GOPATH/src/github.com/tuterdust/my-todo-list
WORKDIR $ROOTPATH
RUN cd $ROOTPATH
RUN pwd
CMD go run *.go  

