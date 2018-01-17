FROM scratch
ADD main /
EXPOSE 8000
ENTRYPOINT ["./main"]
