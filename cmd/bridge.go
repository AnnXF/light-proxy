package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"net/url"
)

func NewBridgeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge",
		Short: "Use bridge CLI to forward requests to the relay PC.",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetString("port")
			relayAddress, _ := cmd.Flags().GetString("relay")
			bridge(port, relayAddress)
		},
	}

	cmd.Flags().StringP("port", "p", "8080", "proxy port, default 8080")
	cmd.Flags().StringP("relay", "r", "", "relay address (required)")
	cmd.MarkFlagRequired("relay")

	return cmd
}

func bridge(port, relayAddress string) {
	relayURL, err := url.Parse(relayAddress)
	if err != nil {
		log.Printf("[sys] parse relay address error:%s \n", err.Error())
		return
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(relayURL),
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[sys] get proxy request url:%s \n", r.URL.String())
		outReq, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
		if err != nil {
			http.Error(w, "[http] failed to create request: "+err.Error(), http.StatusInternalServerError)
			log.Printf("[sys] failed to create request: %s \n", err.Error())
			return
		}
		outReq.Header = r.Header

		resp, err := client.Do(outReq)
		if err != nil {
			http.Error(w, "[http] send request error:"+err.Error(), http.StatusInternalServerError)
			log.Printf("[sys] send request error:%s \n", err.Error())
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})
	log.Printf("[sys] proxy server start success port:%s \n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
