FROM golang:1.18.4 AS build
WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-extldflags '-static'" -o /bin/vid-dl-go

FROM scratch
WORKDIR /
COPY --from=build /bin/vid-dl-go .
ENV GIN_MODE=release
ENV PORT=80
EXPOSE 80
CMD ["/vid-dl-go"]