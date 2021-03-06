package middleware

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	utillog "github.com/Azure/ARO-RP/pkg/util/log"
)

type logResponseWriter struct {
	http.ResponseWriter

	statusCode int
	bytes      int
}

func (w *logResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker := w.ResponseWriter.(http.Hijacker)
	return hijacker.Hijack()
}

func (w *logResponseWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.bytes += n
	return n, err
}

func (w *logResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

type logReadCloser struct {
	io.ReadCloser

	bytes int
}

func (rc *logReadCloser) Read(b []byte) (int, error) {
	n, err := rc.ReadCloser.Read(b)
	rc.bytes += n
	return n, err
}

func Log(baseLog *logrus.Entry) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()

			r.Body = &logReadCloser{ReadCloser: r.Body}
			w = &logResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			log := baseLog
			log = utillog.EnrichWithPath(log, r.URL.Path)

			username, _ := r.Context().Value(ContextKeyUsername).(string)

			log = log.WithFields(logrus.Fields{
				"request_method":      r.Method,
				"request_path":        r.URL.Path,
				"request_proto":       r.Proto,
				"request_remote_addr": r.RemoteAddr,
				"request_user_agent":  r.UserAgent(),
				"username":            username,
			})
			log.Print("read request")

			defer func() {
				log.WithFields(logrus.Fields{
					"body_read_bytes":      r.Body.(*logReadCloser).bytes,
					"body_written_bytes":   w.(*logResponseWriter).bytes,
					"duration":             time.Since(t).Seconds(),
					"response_status_code": w.(*logResponseWriter).statusCode,
				}).Print("sent response")
			}()

			h.ServeHTTP(w, r)
		})
	}
}
