package kiwivm

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/google/go-querystring/query"
	"resty.dev/v3"
)

type Client struct {
	auth *Auth
	hc   *resty.Client
}

func NewClient(veID, apiKey string, optFunc ...ClientOptFunc) *Client {
	hc := resty.New().SetResponseMiddlewares(
		resty.AutoParseResponseMiddleware,
		customErrResponseMiddleware,
	).AddContentTypeDecoder(plainTextType, DecodeJson)
	for _, f := range optFunc {
		f(hc)
	}
	return &Client{
		auth: &Auth{
			VeID:   veID,
			APIKey: apiKey,
		},
		hc: hc,
	}
}

func GetDefaultTestClient() *Client {
	return NewClient(os.Getenv(EnvVeID), os.Getenv(EnvApiKey), WithDebug(true))
}

func Get[T any](ctx context.Context, c *Client, path string) (*T, error) {
	return GetWithQueryParams[T](ctx, c, path, nil)
}

func GetWithQueryParams[T any](ctx context.Context, c *Client, path string, req any) (*T, error) {
	values, err := query.Values(c.auth)
	if err != nil {
		return nil, err
	}

	if req != nil {
		reqQueryValues, err := query.Values(req)
		if err != nil {
			return nil, err
		}
		for k, vs := range reqQueryValues {
			for _, v := range vs {
				values.Add(k, v)
			}
		}
	}

	r, err := c.hc.R().
		SetContext(ctx).
		SetQueryParamsFromValues(values).
		SetExpectResponseContentType(plainTextType).
		SetError(new(ErrorRsp)).
		SetResult(new(T)).
		Get(Host + path)
	if err != nil {
		return nil, err
	}
	rResult, rError := r.Result(), r.Error()
	if rError != nil {
		if rError1, ok := rError.(*ErrorRsp); ok {
			return nil, &Error{
				Code:    rError1.Error,
				Message: rError1.Message,
			}
		}
	} else {
		if rResult1, ok := rResult.(*T); ok {
			return rResult1, nil
		}
	}
	return nil, UnknownError
}

func customErrResponseMiddleware(c *resty.Client, res *resty.Response) error {
	if res.Error() == nil {
		res.Request.SetError(new(ErrorRsp))
	}

	defer func() {
		if cl, ok := res.Body.(io.Closer); ok {
			_ = cl.Close()
		}
	}()

	if err := DecodeJson(res.Body, res.Request.Error); err != nil {
		return err
	}
	res.IsRead = true

	errorRsp, ok := res.Request.Error.(*ErrorRsp)
	if !ok {
		return errors.New("unknown error type")
	}

	if errorRsp.Error == 0 {
		res.Request.Error = nil
	}

	return nil
}
