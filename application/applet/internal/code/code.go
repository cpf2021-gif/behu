package code

import "behu/pkg/xcode"

var (
	RegisterMobileEmpty   = xcode.New(10001, "手机号不能为空")
	VerificationCodeEmpty = xcode.New(10002, "验证码不能为空")
	MobileHasRegistered   = xcode.New(10003, "手机号已注册")
	LoginMobileEmpty      = xcode.New(10004, "手机号不能为空")
)
