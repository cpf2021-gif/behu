package svc

import (
	"behu/application/article/api/internal/config"
	"behu/application/article/rpc/article"
	"behu/pkg/interceptors"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	defaultOssConnectionTimeout = 50
	defaultOssReadWriteTimeout  = 50
)

type ServiceContext struct {
	Config     config.Config
	OssClient  *oss.Client
	ArticleRPC article.Article
}

func NewServiceContext(c config.Config) *ServiceContext {
	if c.Oss.ConnectTimeout == 0 {
		c.Oss.ConnectTimeout = defaultOssConnectionTimeout
	}
	if c.Oss.ReadWriteTimeout == 0 {
		c.Oss.ReadWriteTimeout = defaultOssReadWriteTimeout
	}
	oc, err := oss.New(c.Oss.Endpoint, c.Oss.AccessKeyId, c.Oss.AccessKeySecret,
		oss.Timeout(c.Oss.ConnectTimeout, c.Oss.ReadWriteTimeout))
	if err != nil {
		panic(err)
	}
	// 自定义拦截器
	articleRPC := zrpc.MustNewClient(c.ArticleRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:     c,
		OssClient:  oc,
		ArticleRPC: article.NewArticle(articleRPC),
	}
}
