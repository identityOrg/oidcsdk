package oidcsdk

import (
	"github.com/google/wire"
	"github.com/identityOrg/oidcsdk/impl/factories"
	"github.com/identityOrg/oidcsdk/impl/manager"
	"github.com/identityOrg/oidcsdk/impl/processors"
	"github.com/identityOrg/oidcsdk/impl/strategies"
	"github.com/identityOrg/oidcsdk/impl/writers"
)

var DefaultProcessorSet = wire.NewSet(
	processors.NewProcessorSequence,
	processors.NewDefaultBearerUserAuthProcessor,
	processors.NewDefaultClientAuthenticationProcessor,
	processors.NewDefaultGrantTypeValidator,
	processors.NewDefaultResponseTypeValidator,
	processors.NewDefaultAccessCodeValidator,
	processors.NewDefaultRefreshTokenValidator,
	processors.NewDefaultStateValidator,
	processors.NewDefaultPKCEValidator,
	processors.NewDefaultRedirectURIValidator,
	processors.NewDefaultAudienceValidationProcessor,
	processors.NewDefaultScopeValidator,
	processors.NewDefaultUserValidator,
	processors.NewDefaultClaimProcessor,
	processors.NewDefaultTokenIntrospectionProcessor,
	processors.NewDefaultTokenRevocationProcessor,
	processors.NewDefaultAuthCodeIssuer,
	processors.NewDefaultAccessTokenIssuer,
	processors.NewDefaultIDTokenIssuer,
	processors.NewDefaultRefreshTokenIssuer,
	processors.NewDefaultTokenPersister,
)

var DefaultManagerSet = wire.NewSet(
	wire.Struct(new(manager.Options), "*"),
	writers.NewDefaultErrorWriter,
	writers.NewDefaultResponseWriter,
	strategies.NewDefaultStrategy,
	factories.NewDefaultRequestContextFactory,
	manager.NewDefaultManager,
	wire.Bind(new(IErrorWriter), new(*writers.DefaultErrorWriter)),
	wire.Bind(new(IResponseWriter), new(*writers.DefaultResponseWriter)),
	wire.Bind(new(IAccessTokenStrategy), new(*strategies.DefaultStrategy)),
	wire.Bind(new(IAuthorizationCodeStrategy), new(*strategies.DefaultStrategy)),
	wire.Bind(new(IRefreshTokenStrategy), new(*strategies.DefaultStrategy)),
	wire.Bind(new(IIDTokenStrategy), new(*strategies.DefaultStrategy)),
	wire.Bind(new(IRequestContextFactory), new(*factories.DefaultRequestContextFactory)),
	wire.Bind(new(IManager), new(*manager.DefaultManager)),
)
