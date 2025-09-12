package aigc

import (
	"log"
	"os"
	"testing"
)

func TestGetRandomAigcAgent(t *testing.T) {
	token, ok := os.LookupEnv("XFYUN_AI_X1_TOKEN")
	if !ok {
		log.Fatalln("请通过系统环境变量（XFYUN_AI_X1_TOKEN）配置讯飞云 AI 的 X1 模型的 Token")
	}
	var agent OpenAiClient
	for i := 0; i < 1; i++ {
		agent = XfyunAiClient{Model: "x1", Token: token}
		resp := agent.ApplyCompletion(`系统：提取事件信息。
		用户：Alice和Bob将在周五参加科学展览会。
		请以JSON格式提供以下信息，不要返回说明文字：

		{
			"name": "事件名称",
			"date": "事件日期",
			"participants": ["参与者列表"]
		}

		注：不要返回说明文字。`, &CompletionOptions{})
		if resp.IsValid() {
			t.Log(resp.GetResult())
		}
	}
}

/**
`虚构一个小说角色顾葳蕤，女性，需要补全角色的兴趣爱好，口头禅等信息，姓名符合个人特质，角色参考特质：活泼、敢爱敢恨、世家嫡女、不拘小节。
    请以JSON格式提供以下信息：
    {
        "name": "角色姓名",
        "gender": "角色性别",
		"appearance": ["外貌描述"],
		"clothing": ["衣着习惯"],
        "lip-services": ["口头禅"],
		"interests": ["兴趣"]
    }`
*/
