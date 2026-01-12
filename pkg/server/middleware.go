package server

import (
	"context"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func logging(logger *slog.Logger) mcp.Middleware {
	return func(handler mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (result mcp.Result, err error) {
			l := logger.With(
				slog.String("id", req.GetSession().ID()),
				slog.String("method", method),
			)

			if method == "tools/call" {
				if params, ok := req.GetParams().(*mcp.CallToolParamsRaw); ok && params != nil {
					l = l.With(
						slog.Any("tool", params.Name),
						slog.Any("arguments", params.Arguments),
					)
				}
			}

			l.DebugContext(ctx, "received  request")

			result, err = handler(ctx, method, req)

			if err != nil {
				l.ErrorContext(ctx, "error occurred",
					slog.String("error", err.Error()),
				)
			} else {
				l.InfoContext(ctx, "request finished")
			}

			return result, err
		}
	}
}
