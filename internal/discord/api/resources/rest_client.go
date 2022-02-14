package resources

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/holedaemon/hubris/internal/discord/types"
)

const (
	base    = "https://discord.com/api/v"
	version = "9"

	root = base + version
)

var defaultHeader = make(http.Header)

func init() {
	defaultHeader.Set("Accept", "application/json")
	defaultHeader.Set("Content-Type", "application/json")
}

type RestClient struct {
	token string
	cli   *http.Client
}

func NewRestClient(token string, cli *http.Client) *RestClient {
	return &RestClient{
		token: token,
		cli:   cli,
	}
}

type RequestOption func(*requestOptions)

type requestOptions struct {
	Body   io.Reader
	Query  url.Values
	Header http.Header
}

func (rc *RestClient) newRequest(ctx context.Context, uri, method string, opts ...RequestOption) (*http.Request, error) {
	u := root + uri

	ro := new(requestOptions)

	for _, o := range opts {
		o(ro)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, ro.Body)
	if err != nil {
		return nil, err
	}

	h := defaultHeader.Clone()
	h.Set("Authorization", rc.token)

	if ro.Header != nil {
		for k := range ro.Header {
			h.Set(k, ro.Header.Get(k))
		}
	}

	req.Header = h

	if ro.Query != nil {
		req.URL.RawQuery = ro.Query.Encode()
	}

	return req, nil
}

func (rc *RestClient) do(ctx context.Context, url, meth string, opts ...RequestOption) (*http.Response, error) {
	req, err := rc.newRequest(ctx, url, meth, opts...)
	if err != nil {
		return nil, err
	}

	res, err := rc.cli.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return nil, types.NewError(res)
	}

	return res, nil
}

func (rc *RestClient) Get(ctx context.Context, url string, v interface{}, opts ...RequestOption) error {
	res, err := rc.do(ctx, url, http.MethodGet, opts...)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

func (rc *RestClient) Post(ctx context.Context, url string, v interface{}, opts ...RequestOption) error {
	res, err := rc.do(ctx, url, http.MethodPost, opts...)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

func (rc *RestClient) Delete(ctx context.Context, url string, v interface{}, opts ...RequestOption) error {
	res, err := rc.do(ctx, url, http.MethodDelete, opts...)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

func (rc *RestClient) Patch(ctx context.Context, url string, v interface{}, opts ...RequestOption) error {
	res, err := rc.do(ctx, url, http.MethodPatch, opts...)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

func (rc *RestClient) Put(ctx context.Context, url string, v interface{}, opts ...RequestOption) error {
	res, err := rc.do(ctx, url, http.MethodPut, opts...)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
