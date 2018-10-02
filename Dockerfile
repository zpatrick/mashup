FROM alpine
RUN apk add --no-cache ca-certificates
ADD ./generator/matrix.json /generator/matrix.json
ADD ./mashup /
CMD ["/mashup"]
