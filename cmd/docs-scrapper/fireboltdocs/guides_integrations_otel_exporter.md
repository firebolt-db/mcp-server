# [](#overview)Overview

[OpenTelemetry](https://opentelemetry.io/) is a [CNCF](https://www.cncf.io/) project that provides a collection of APIs, SDKs, and tools to instrument, generate, collect, and export telemetry data (metrics, logs, and traces). In the past few years, this project become the accepted standard for telemetry, with native support by all major vendors. As such, Firebolt provides an OpenTelemetry exporter which gives and compatibility with minimal effort.

Firebolt OpenTelemetry Exporter is provided as a docker image, which allows exporting engine metrics to any [OTLP](https://opentelemetry.io/docs/specs/otel/protocol/) compatible collector. This makes possible to integrate Firebolt runtime metrics into customer’s monitoring and alerting systems and be able to use homogeneous infrastructure for observability of the entire data stack.

# [](#enabling-firebolt-opentelemetry-exporter)Enabling Firebolt OpenTelemetry Exporter

For installation and usage instructions, see the [otel-exporter](https://github.com/firebolt-db/otel-exporter) repository on GitHub.