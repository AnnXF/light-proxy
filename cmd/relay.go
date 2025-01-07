package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

func NewRelayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relay",
		Short: "Use relay CLI to forward requests to the real network.",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetString("port")
			relay(port)
		},
	}

	cmd.Flags().StringP("port", "p", "8080", "proxy port, default 8080")

	return cmd
}

func relay(port string) {
	proxy := http.NewServeMux()
	proxy.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		transport := &http.Transport{}
		log.Printf("[sys] get proxy request url:%s \n", r.URL.String())
		outReq, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
		if err != nil {
			log.Printf("[sys] create request error:%s \n", err.Error())
			http.Error(w, "[http] failed to create request", http.StatusInternalServerError)
			return
		}

		outReq.Header = r.Header
		resp, err := transport.RoundTrip(outReq)
		if err != nil {
			log.Printf("[sys] transport round trip error:%s \n", err.Error())
			http.Error(w, "[http] failed to forward request error:"+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		for k, v := range resp.Header {
			w.Header()[k] = v
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

	log.Printf("[sys] proxy server start success port:%s \n", port)
	log.Fatal(http.ListenAndServe(":"+port, proxy))
}
