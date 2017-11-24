package main

import (
    "net/http"
    "context"
    "fmt"
)

const requestIDKey = "rid"

func requestIDFromContext(ctx context.Context) string {
    //should be val, ok = xxxx.(string)
    val, ok := ctx.Value(requestIDKey).(string)
    if ok {
        return val
    } else {
        return "" 
    }
}

func newContextWithRequestID(ctx, req *http.Request) conext.Context {
    reqID := req.Header.Get("X-Request-ID")
    if reqID == "" {
        reqID = ""
    }
    return context.WithValue(ctx, requestIDKey, reqID)
}

//实现了所有的请求都在上下文存了一个requestID
func middleWare(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        //Context returns the request's context. To change the context, use WithContext.
        //first->get the context and gen a new context => then set the new context to request
        ctx := newContextWithRequestID(req.Context(), req)

        //ServeHTTP calls f(w, r). => next(w, r)
        //WithContext returns a shallow copy of r with its context changed to ctx. The provided ctx must be non-nil.
        next.ServeHTTP(w, req.WithContext(ctx))
    })
}

func h(w http.ResponseWriter, req *http.Request) {
   reqID := requestIDFromContext(req.Context())
   fmt.Fprintln(w, "Request ID:", reqID)
   return
}

func main() {
    http.Handle("/", middleWare(http.HandlerFunc(h)))
    http.ListenAndServe(":9092", nil)
}
