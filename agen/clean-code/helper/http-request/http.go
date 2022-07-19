package httprequest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// HTTPApp ...
type HTTPApp struct {
	client      *http.Client
	headers     map[string]string
	baseURL     string
	debugEnable bool
}

// NewHTTPApp ...
func NewHTTPApp(baseURL string, headers map[string]string) *HTTPApp {
	transport := &http.Transport{DialContext: (&net.Dialer{
		// limits the time spent establishing a TCP connection
		Timeout: 30 * time.Second,
		//KeepAlive:     10 * time.Second,
	}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return &HTTPApp{
		client: &http.Client{
			Transport: transport,
		},
		headers: headers,
		baseURL: baseURL,
	}
}

// MakeRequest ...
func (o *HTTPApp) MakeRequest(ctx context.Context, method, path string, body io.Reader, headers map[string]string, q url.Values) (*http.Response, error) {

	url := o.baseURL
	if path != "" {
		url += "/" + path
	}
	// log.Printf("[MakeRequest]: method: %s, url: %s, headers: %#v, body: %#v, q: %#v \n", method, url, headers, body, q)

	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	if len(q) > 0 {
		request.URL.RawQuery = q.Encode()
	}
	if len(o.headers) > 0 {
		for k, v := range o.headers {
			request.Header.Set(k, v)
		}
	}
	if len(headers) > 0 {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}

	client := o.client // http.Client{}
	// log.Debug(httputil.DumpRequestOut(request, true))
	return client.Do(request)
}

// SetDebug ...
func (o *HTTPApp) SetDebug(b bool) {
	o.debugEnable = b
}

// Get ...
func (o HTTPApp) Get(ctx context.Context, path string, headers map[string]string, q url.Values) (*http.Response, error) {
	fURL, err := o.resolveURL(path)
	if err != nil {
		return nil, err
	}
	headersRequest := make(map[string]string)
	if len(o.headers) > 0 {
		for k, v := range o.headers {
			headersRequest[k] = v
		}
	}
	if len(headers) > 0 {
		for k, v := range headers {
			headersRequest[k] = v
		}
	}
	res, err := o.MakeRequest(ctx, http.MethodGet, fURL, nil, headersRequest, q)
	if err != nil {
		return res, err
	}
	if res.StatusCode != http.StatusOK {
		bb, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		res.Body = ioutil.NopCloser(bytes.NewBuffer(bb))
		//ll.Error("Get", l.String("res", string(bb)), l.Error(err))
		return res, fmt.Errorf("not succes, status: %v", res.StatusCode)
	}
	return res, err
}

func (o HTTPApp) resolveURL(urlPath string) (string, error) {
	if strings.HasPrefix(urlPath, "http") {
		return urlPath, nil
	}
	urlObj, err := url.Parse(o.baseURL)
	if err != nil {
		return "", err
	}
	urlObj.Path = path.Join(urlObj.Path, urlPath)
	finalURL := urlObj.String()
	return finalURL, nil
}
