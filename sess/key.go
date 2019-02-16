package sess

import (
	"context"
	"encoding/base64"
	"os"

	"google.golang.org/appengine/log"
)

func getSessionKey(ctx context.Context) []byte {
	b64sc := os.Getenv("SESSION_KEY")
	log.Debugf(ctx, "B64: %v", b64sc)
	sc, _ := base64.StdEncoding.DecodeString(b64sc)

	log.Debugf(ctx, "SESSION KEY [%v]", sc)

	return sc
}
