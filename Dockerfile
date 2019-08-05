FROM alpine
ADD controller-api /controller-api
ENTRYPOINT [ "/controller-api" ]
