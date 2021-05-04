package thread

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/progfay/kataribe-playgound/client/request"
)

type thread struct {
	client *http.Client
}

func New() *thread {
	return &thread{
		client: new(http.Client),
	}
}

func (t *thread) do(ctx context.Context, requestLines []request.RequestLine) error {
	if len(requestLines) == 0 {
		return errors.New("requestLines must have at least one requestLine")
	}

	req, err := requestLines[rand.Intn(len(requestLines))].NewRequest()
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	res, err := t.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	defer io.Copy(ioutil.Discard, res.Body)
	return nil
}

func (t *thread) Run(ctx context.Context, requestLines []request.RequestLine) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		err := t.do(ctx, requestLines)
		if err != nil {
			cancel()
			return err
		}
	}
}
