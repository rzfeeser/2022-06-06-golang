/* Alta3 Research | RZFeeser
   Prometheus - exposing metrics in a go application requires about
     HTTP endpoint. */
   
package main

import (
        "net/http"

        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":2112", nil)
}

