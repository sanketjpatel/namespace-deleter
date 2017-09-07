FROM alpine:3.6
COPY namespace-deleter /namespace-deleter
CMD ["./namespace-deleter"]
