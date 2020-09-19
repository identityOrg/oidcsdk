package impl

import (
	"github.com/google/wire"
	"github.com/identityOrg/oidcsdk"
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
	wire.Bind(new(oidcsdk.IErrorWriter), new(*writers.DefaultErrorWriter)),
	wire.Bind(new(oidcsdk.IResponseWriter), new(*writers.DefaultResponseWriter)),
	wire.Bind(new(oidcsdk.IAccessTokenStrategy), new(*strategies.DefaultStrategy)),
	wire.Bind(new(oidcsdk.IAuthorizationCodeStrategy), new(*strategies.DefaultStrategy)),
	wire.Bind(new(oidcsdk.IRefreshTokenStrategy), new(*strategies.DefaultStrategy)),
	wire.Bind(new(oidcsdk.IIDTokenStrategy), new(*strategies.DefaultStrategy)),
	wire.Bind(new(oidcsdk.IRequestContextFactory), new(*factories.DefaultRequestContextFactory)),
	manager.NewDefaultManager,
)
