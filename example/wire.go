//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/example/config"
	"github.com/identityOrg/oidcsdk/example/demosession"
	"github.com/identityOrg/oidcsdk/example/memdbstore"
	"github.com/identityOrg/oidcsdk/example/pages"
	"github.com/identityOrg/oidcsdk/example/secretkey"
	"github.com/identityOrg/oidcsdk/impl"
	"github.com/identityOrg/oidcsdk/impl/manager"
)

var DefaultStoreSet = wire.NewSet(
	memdbstore.NewInMemoryDB,
	demosession.NewManager,
	secretkey.NewDefaultMemorySecretStore,
	pages.NewPageRenderer,
	wire.Bind(new(oidcsdk.ISecretStore), new(*secretkey.DefaultMemorySecretStore)),
	wire.Bind(new(oidcsdk.ITokenStore), new(*memdbstore.InMemoryDB)),
	wire.Bind(new(oidcsdk.IUserStore), new(*memdbstore.InMemoryDB)),
	wire.Bind(new(oidcsdk.IClientStore), new(*memdbstore.InMemoryDB)),
	wire.Bind(new(oidcsdk.ISessionManager), new(*demosession.Manager)),
	wire.Bind(new(oidcsdk.IPageResponseHandler), new(*pages.PageRenderer)),
)

func ComposeNewManager(config *oidcsdk.Config, demo bool, demoConfig *config.DemoConfig) *manager.DefaultManager {
	wire.Build(impl.DefaultManagerSet, impl.DefaultProcessorSet, DefaultStoreSet)
	return nil
}

func ComposeSessionStore(demoConfig *config.DemoConfig) (manager *demosession.Manager) {
	wire.Build(DefaultStoreSet)
	return nil
}

func ComposeDemoStore(demoConfig *config.DemoConfig, demo bool) *memdbstore.InMemoryDB {
	wire.Build(DefaultStoreSet)
	return nil
}
