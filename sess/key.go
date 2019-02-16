package sess

import (
	"context"
	"encoding/base64"
	"os"

	"google.golang.org/appengine/log"
)

func getSessionKey(ctx context.Context) ([]byte, error) {
	sc, err := base64.StdEncoding.DecodeString(os.Getenv("SESSION_KEY"))
	if err != nil {
		log.Debugf(ctx, "Could not get session key from environment. [%v]", err)

		return nil, err
	}

	return sc, nil
}

func getEncKey(ctx context.Context) ([]byte, error) {
	sc, err := base64.StdEncoding.DecodeString(os.Getenv("SESSION_ENC"))
	if err != nil {
		log.Debugf(ctx, "Could not get session key from environment. [%v]", err)

		return nil, err
	}

	return sc, nil
}
