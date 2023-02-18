package azure

import "github.com/spctr-cc/backend-bugtrack/model"

type Store interface {
	Create(*model.Azure) error
}
