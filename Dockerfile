FROM golang:latest
RUN go install github.com/stephenafamo/bob/gen/bobgen-psql@v0.26.1
ENTRYPOINT []
#ENTRYPOINT ["go", "run", "github.com/stephenafamo/bob/gen/bobgen-psql@v0.26.1"]