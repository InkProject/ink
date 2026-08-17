FROM golang:1.26
ADD . /code
WORKDIR /code
RUN go install
EXPOSE 8000
CMD ["ink", "preview", "template"]
