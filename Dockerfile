FROM alpine
COPY main /home/infocpu
ENTRYPOINT ["/home/infocpu"]