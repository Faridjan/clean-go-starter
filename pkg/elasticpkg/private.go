package elasticpkg

import (
	"context"

	"github.com/olivere/elastic/v7"

	"git.centerhome.kz/bcc/backend/toolchain/common-libs/errors"
)

// Global default constants
const elasticLimit = 10000

// initConnect
// init elastic conn
func initConnect(ctx context.Context, nodesURLs []string, setSniff bool, username, password string) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(nodesURLs...),
		elastic.SetSniff(setSniff),
		elastic.SetBasicAuth(username, password),
	)
	if err != nil {
		return nil, errors.InternalServerError(ctx).
			SetDevMessage("initConnect.elastic.NewClient: " + err.Error())
	}

	return client, nil
}
