# syntax=docker/dockerfile:1

FROM gcr.io/distroless/base:nonroot
COPY ./firebolt-mcp-server /usr/local/bin/firebolt-mcp-server
ENTRYPOINT ["/usr/local/bin/firebolt-mcp-server"]
