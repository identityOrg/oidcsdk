// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/example/config"
	"github.com/identityOrg/oidcsdk/example/demosession"
	"github.com/identityOrg/oidcsdk/example/memdbstore"
	"github.com/identityOrg/oidcsdk/example/pages"
	"github.com/identityOrg/oidcsdk/example/secretkey"
	"github.com/identityOrg/oidcsdk/impl/factories"
	"github.com/identityOrg/oidcsdk/impl/manager"
	"github.com/identityOrg/oidcsdk/impl/processors"
	"github.com/identityOrg/oidcsdk/impl/strategies"
	"github.com/identityOrg/oidcsdk/impl/writers"
)

// Injectors from wire.go:

func ComposeNewManager(config2 *oidcsdk.Config, demo bool, demoConfig *config.DemoConfig) *manager.DefaultManager {
	pageRenderer := pages.NewPageRenderer()
	defaultRequestContextFactory := factories.NewDefaultRequestContextFactory()
	defaultErrorWriter := writers.NewDefaultErrorWriter()
	defaultResponseWriter := writers.NewDefaultResponseWriter()
	demosessionManager := demosession.NewManager(demoConfig)
	defaultMemorySecretStore := secretkey.NewDefaultMemorySecretStore()
	inMemoryDB := memdbstore.NewInMemoryDB(demo)
	defaultStrategy := strategies.NewDefaultStrategy(defaultMemorySecretStore, config2)
	defaultBearerUserAuthProcessor := processors.NewDefaultBearerUserAuthProcessor(inMemoryDB, inMemoryDB, defaultStrategy)
	defaultClientAuthenticationProcessor := processors.NewDefaultClientAuthenticationProcessor(inMemoryDB)
	defaultGrantTypeValidator := processors.NewDefaultGrantTypeValidator()
	defaultRPILogoutIDTokenValidator := processors.NewDefaultRPILogoutIDTokenValidator(defaultStrategy, inMemoryDB)
	defaultResponseTypeValidator := processors.NewDefaultResponseTypeValidator()
	defaultAccessCodeValidator := processors.NewDefaultAccessCodeValidator(inMemoryDB, defaultStrategy)
	defaultRefreshTokenValidator := processors.NewDefaultRefreshTokenValidator(defaultStrategy, inMemoryDB)
	defaultStateValidator := processors.NewDefaultStateValidator(config2)
	defaultPKCEValidator := processors.NewDefaultPKCEValidator(config2)
	defaultRedirectURIValidator := processors.NewDefaultRedirectURIValidator()
	defaultAudienceValidationProcessor := processors.NewDefaultAudienceValidationProcessor()
	defaultScopeValidator := processors.NewDefaultScopeValidator()
	defaultUserValidator := processors.NewDefaultUserValidator(inMemoryDB, inMemoryDB, config2)
	defaultClaimProcessor := processors.NewDefaultClaimProcessor(inMemoryDB)
	defaultTokenIntrospectionProcessor := processors.NewDefaultTokenIntrospectionProcessor(inMemoryDB, defaultStrategy, defaultStrategy)
	defaultTokenRevocationProcessor := processors.NewDefaultTokenRevocationProcessor(inMemoryDB, defaultStrategy, defaultStrategy)
	defaultAuthCodeIssuer := processors.NewDefaultAuthCodeIssuer(defaultStrategy, config2)
	defaultAccessTokenIssuer := processors.NewDefaultAccessTokenIssuer(defaultStrategy, config2)
	defaultIDTokenIssuer := processors.NewDefaultIDTokenIssuer(defaultStrategy, config2)
	defaultRefreshTokenIssuer := processors.NewDefaultRefreshTokenIssuer(defaultStrategy, config2)
	defaultTokenPersister := processors.NewDefaultTokenPersister(inMemoryDB, inMemoryDB, config2)
	v := processors.NewProcessorSequence(defaultBearerUserAuthProcessor, defaultClientAuthenticationProcessor, defaultGrantTypeValidator, defaultRPILogoutIDTokenValidator, defaultResponseTypeValidator, defaultAccessCodeValidator, defaultRefreshTokenValidator, defaultStateValidator, defaultPKCEValidator, defaultRedirectURIValidator, defaultAudienceValidationProcessor, defaultScopeValidator, defaultUserValidator, defaultClaimProcessor, defaultTokenIntrospectionProcessor, defaultTokenRevocationProcessor, defaultAuthCodeIssuer, defaultAccessTokenIssuer, defaultIDTokenIssuer, defaultRefreshTokenIssuer, defaultTokenPersister)
	options := &manager.Options{
		PageResponseHandler:   pageRenderer,
		RequestContextFactory: defaultRequestContextFactory,
		ErrorWriter:           defaultErrorWriter,
		ResponseWriter:        defaultResponseWriter,
		UserSessionManager:    demosessionManager,
		SecretStore:           defaultMemorySecretStore,
		Sequence:              v,
	}
	defaultManager := manager.NewDefaultManager(config2, options)
	return defaultManager
}

func ComposeSessionStore(demoConfig *config.DemoConfig) *demosession.Manager {
	demosessionManager := demosession.NewManager(demoConfig)
	return demosessionManager
}

func ComposeDemoStore(demoConfig *config.DemoConfig, demo bool) *memdbstore.InMemoryDB {
	inMemoryDB := memdbstore.NewInMemoryDB(demo)
	return inMemoryDB
}

// wire.go:

var DefaultStoreSet = wire.NewSet(memdbstore.NewInMemoryDB, demosession.NewManager, secretkey.NewDefaultMemorySecretStore, pages.NewPageRenderer, wire.Bind(new(oidcsdk.ISecretStore), new(*secretkey.DefaultMemorySecretStore)), wire.Bind(new(oidcsdk.ITokenStore), new(*memdbstore.InMemoryDB)), wire.Bind(new(oidcsdk.IUserStore), new(*memdbstore.InMemoryDB)), wire.Bind(new(oidcsdk.IClientStore), new(*memdbstore.InMemoryDB)), wire.Bind(new(oidcsdk.ISessionManager), new(*demosession.Manager)), wire.Bind(new(oidcsdk.IPageResponseHandler), new(*pages.PageRenderer)))
