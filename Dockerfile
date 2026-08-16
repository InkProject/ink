FROM golang:1.25
ADD . /code
WORKDIR /code
RUN go install
RUN chmod +x entrypoint.sh
EXPOSE 8000
ENTRYPOINT ["/code/entrypoint.sh"]
CMD ["ink", "preview", "builds"]
