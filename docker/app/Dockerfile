FROM golang:1.10.3
ADD ./src /go/src/DXFimporter
ADD ./resources/DXF /resources
ADD ./output /output
ADD ./bin /go/bin
RUN go get github.com/olekukonko/tablewriter
WORKDIR /go/src/DXFimporter
RUN go build -a -o /go/bin/DXFimporter
CMD [ "DXFimporter", "--mount=/resources", "--output=/output/output.wol" ]