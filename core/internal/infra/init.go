package infra

import (
	"github.com/jazzsewera/notipie/core/internal/impl"
	"github.com/jazzsewera/notipie/core/internal/impl/grid"
	"github.com/jazzsewera/notipie/core/pkg/lib/log"
)

type AppContext struct {
	gr grid.Grid
	ap grid.AppProxy
	up grid.UserProxy
	ep *impl.Endpoint
}

func (c *AppContext) Init(config Config) {
	c.initLogger(config.prod)
	c.initGrid()
	c.initEndpoint()
}

func (c *AppContext) initLogger(prod bool) {
	log.Init(prod)
}

func (c *AppContext) initGrid() {
	c.gr = grid.NewGrid()
	c.ap = grid.NewAppProxy(c.gr)
	c.up = grid.NewUserProxy()
	return
}

func (c *AppContext) initEndpoint() {
	c.ep = impl.NewEndpoint(c.gr, c.ap, c.up)
}

func (c *AppContext) Start() {

}