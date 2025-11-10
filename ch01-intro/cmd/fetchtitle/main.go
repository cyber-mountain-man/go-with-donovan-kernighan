package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"time"
)

var titleRE = regexp.MustCompile(`(?is)<\s*title[^>]*>(.*?)</\s*title\s*>`)

func main() {
	timeoutStr := flag.String("timeout", "12s", "total request timeout (e.g. 5s, 10s)")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: fetchtitle [-timeout=12s] URL")
		os.Exit(2)
	}
	url := flag.Arg(0)

	// Parse timeout
	tout, err := time.ParseDuration(*timeoutStr)
	if err != nil || tout <= 0 {
		fmt.Fprintf(os.Stderr, "invalid timeout: %q\n", *timeoutStr)
		os.Exit(2)
	}

	// Transport with sane timeouts
	dialer := &net.Dialer{
		Timeout:   7 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		DialContext:         dialer.DialContext,
		TLSHandshakeTimeout: 7 * time.Second,
		ForceAttemptHTTP2:   true,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   tout,
	}

	ctx, cancel := context.WithTimeout(context.Background(), tout)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "fetchtitle/1.0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		fmt.Fprintf(os.Stderr, "fetch: HTTP %s\n", resp.Status)
		os.Exit(1)
	}

	body, _ := io.ReadAll(resp.Body)
	m := titleRE.FindSubmatch(body)
	if len(m) == 2 {
		fmt.Println(string(m[1]))
	} else {
		fmt.Println("(no <title> found)")
	}
}
