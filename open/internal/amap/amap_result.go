package amap

import "github.com/MurphyL/lego-kits/open/internal/third_party"

type ParsedResult map[string]any

func (r ParsedResult) Success() bool {
	return r["status"] == "1" || r["status"] == 1
}

func (r ParsedResult) Code() uint {
	return r["infocode"].(uint)
}

func (r ParsedResult) Get(key string) any {
	return r[key]
}

func ErrorOf(code string) *third_party.EndpointStatus {
	switch code {
	case "10000":
		return third_party.NewStatus("10000", "OK", "请求正常")
	case "10001":
		return third_party.NewStatus("10001", "INVALID_USER_KEY", "开发者发起请求时，传入的key不正确或者过期")
	case "10002":
		return third_party.NewStatus("10002", "SERVICE_NOT_AVAILABLE", "1.开发者没有权限使用相应的服务，2.开发者请求接口的路径拼写错误。")
	case "10003":
		return third_party.NewStatus("10003", "DAILY_QUERY_OVER_LIMIT", "开发者的日访问量超限，被系统自动封停，第二天0:00会自动解封。")
	case "10004":
		return third_party.NewStatus("10004", "ACCESS_TOO_FREQUENT", "开发者的单位时间内（1分钟）访问量超限，被系统自动封停，下一分钟自动解封。")
	case "10005":
		return third_party.NewStatus("10005", "INVALID_USER_IP", "开发者在LBS官网控制台设置的IP白名单不正确。白名单中未添加对应服务器的出口IP。可到“控制台>配置”中设定IP白名单。")
	case "10006":
		return third_party.NewStatus("10006", "INVALID_USER_DOMAIN", "开发者绑定的域名无效，需要在官网控制台重新设置")
	case "10007":
		return third_party.NewStatus("10007", "INVALID_USER_SIGNATURE", "开发者签名未通过开发者在key控制台中，开启了“数字签名”功能，但没有按照指定算法生成“数字签名”。")
	case "10008":
		return third_party.NewStatus("10008", "INVALID_USER_SCODE", "需要开发者判定key绑定的SHA1,package是否与sdk包里的一致")
	case "10009":
		return third_party.NewStatus("10009", "USERKEY_PLAT_NOMATCH", "请求中使用的key与绑定平台不符，例如：开发者申请的是js api的key，却用来调web服务接口")
	case "10010":
		return third_party.NewStatus("10010", "IP_QUERY_OVER_LIMIT", "未设定IP白名单的开发者使用key发起请求，从单个IP向服务器发送的请求次数超出限制，被系统自动封停。封停后无法自动恢复，需要提交工单联系我们。")
	case "10011":
		return third_party.NewStatus("10011", "NOT_SUPPORT_HTTPS", "服务不支持https请求，如果需要申请支持，请提交工单联系我们")
	case "10012":
		return third_party.NewStatus("10012", "INSUFFICIENT_PRIVILEGES", "由于不具备请求该服务的权限，所以服务被拒绝。")
	case "10013":
		return third_party.NewStatus("10013", "USER_KEY_RECYCLED", "开发者删除了key，key被删除后无法正常使用")
	case "10014":
		return third_party.NewStatus("10014", "QPS_HAS_EXCEEDED_THE_LIMIT", "QPS超出限制，超出部分的请求被拒绝。限流阈值内的请求依旧会正常返回")
	case "10015":
		return third_party.NewStatus("10015", "GATEWAY_TIMEOUT", "受单机QPS限流限制时出现该问题，建议降低请求的QPS或在控制台提工单联系我们")
	case "10016":
		return third_party.NewStatus("10016", "SERVER_IS_BUSY", "服务器负载过高，请稍后再试")
	case "10017":
		return third_party.NewStatus("10017", "RESOURCE_UNAVAILABLE", "所请求的资源不可用")
	case "10019":
		return third_party.NewStatus("10019", "CQPS_HAS_EXCEEDED_THE_LIMIT", "QPS超出限制，超出部分的请求被拒绝。限流阈值内的请求依旧会正常返回")
	case "10020":
		return third_party.NewStatus("10020", "CKQPS_HAS_EXCEEDED_THE_LIMIT", "QPS超出限制，超出部分的请求被拒绝。限流阈值内的请求依旧会正常返回")
	case "10021":
		return third_party.NewStatus("10021", "CUQPS_HAS_EXCEEDED_THE_LIMIT", "QPS超出限制，超出部分的请求被拒绝。限流阈值内的请求依旧会正常返回")
	case "10026":
		return third_party.NewStatus("10026", "INVALID_REQUEST", "由于违规行为账号被封禁不可用，如有异议请登录控制台提交工单进行申诉")
	case "10029":
		return third_party.NewStatus("10029", "ABROAD_DAILY_QUERY_OVER_LIMIT", "QPS超出限制，超出部分的请求被拒绝。限流阈值内的请求依旧会正常返回")
	case "10041":
		return third_party.NewStatus("10041", "NO_EFFECTIVE_INTERFACE", "开发者发起请求时，请求的接口权限过期。请提交工单联系我们")
	case "10044":
		return third_party.NewStatus("10044", "USER_DAILY_QUERY_OVER_LIMIT", "账号维度日调用量超出限制，超出部分的请求被拒绝。限流阈值内的请求依旧会正常返回")
	case "10045":
		return third_party.NewStatus("10045", "USER_ABROAD_DAILY_QUERY_OVER_LIMIT", "账号维度海外服务接口日调用量超出限制，超出部分的请求被拒绝。限流阈值内的请求依旧会正常返回")
	case "20000":
		return third_party.NewStatus("20000", "INVALID_PARAMS", "请求参数的值没有按照规范要求填写。例如，某参数值域范围为[1,3],开发者误填了’4’")
	case "20001":
		return third_party.NewStatus("20001", "MISSING_REQUIRED_PARAMS", "缺少接口中要求的必填参数")
	case "20002":
		return third_party.NewStatus("20002", "ILLEGAL_REQUEST", "请求协议非法比如某接口仅支持get请求，结果用了POST方式")
	case "20003":
		return third_party.NewStatus("20003", "UNKNOWN_ERROR", "其他未知错误")
	case "20011":
		return third_party.NewStatus("20011", "INSUFFICIENT_ABROAD_PRIVILEGES", "使用逆地理编码接口、输入提示接口、周边搜索接口、路径规划接口时可能出现该问题，规划点（包括起点、终点、途经点）不在中国陆地范围内")
	case "20012":
		return third_party.NewStatus("20012", "ILLEGAL_CONTENT", "使用搜索接口时可能出现该问题，通常是由于查询内容非法导致")
	case "20800":
		return third_party.NewStatus("20800", "OUT_OF_SERVICE", "使用路径规划服务接口时可能出现该问题，规划点（包括起点、终点、途经点）不在中国大陆陆地范围内")
	case "20801":
		return third_party.NewStatus("20801", "NO_ROADS_NEARBY", "使用路径规划服务接口时可能出现该问题，划点（起点、终点、途经点）附近搜不到路")
	case "20802":
		return third_party.NewStatus("20802", "ROUTE_FAIL", "使用路径规划服务接口时可能出现该问题，路线计算失败，通常是由于道路连通关系导致")
	case "20803":
		return third_party.NewStatus("20803", "OVER_DIRECTION_RANGE", "使用路径规划服务接口时可能出现该问题，路线计算失败，通常是由于道路起点和终点距离过长导致。")
	case "300**":
		return third_party.NewStatus("300**", "ENGINE_RESPONSE_DATA_ERROR", "出现3开头的错误码，建议先检查传入参数是否正确，若无法解决，请详细描述错误复现信息，提工单给我们。（大数据接口请直接跟负责商务反馈）如，30001、30002、30003、32000、32001、32002、32003、32200、32201、32202、32203。")
	case "40000":
		return third_party.NewStatus("40000", "QUOTA_PLAN_RUN_OUT", "所购买服务的余额耗尽，无法继续使用服务")
	case "40001":
		return third_party.NewStatus("40001", "GEOFENCE_MAX_COUNT_REACHED", "Key可创建的地理围栏的数量，已达上限。")
	case "40002":
		return third_party.NewStatus("40002", "SERVICE_EXPIRED", "所购买的服务期限已到，无法继续使用")
	case "40003":
		return third_party.NewStatus("40003", "ABROAD_QUOTA_PLAN_RUN_OUT", "所购买服务的海外余额耗尽，无法继续使用服务")
	default:
		return nil
	}
}
