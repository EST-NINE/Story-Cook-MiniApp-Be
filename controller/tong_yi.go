package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ncuhome/story-cook/config"
	"github.com/ncuhome/story-cook/pkg/util"
	"net/http"
)

func SendReqToTongYi(title string, background string, keywords string) *http.Response {
	apiKey := config.ApiKey
	prompt := fmt.Sprintf("标题：%s 故事背景：%s 关键词：%s", title, background, keywords)
	charaSetting := "## 定义一种新格式\\n接下来我会定义一种新的对象表示法“TYON-2”，这种“TYON-2”表示法的格式是：以{startgen}" +
		"开始，每一个对象间以{endthisvalue}衔接。请注意：{startgen}与{endthisvalue}标识极其重要，而{endthisvalue}更是必须在每次" +
		"value结束后标记，这关乎TYON-2解析的稳定性与安全性！其中对象的key值只在{endthisvalue}后，{startgen}后不需要key值，key值用[key]表示" +
		"（其中key替换为相应的值），后面紧跟着value值，key与value之间没有间隔。请根据我将给出的示例认真了解{startgen}与{endthisvalue}两个标识符的位置！感谢配合。" +
		"## 指令\\n你是一个作家，你现在要根据用户给出的故事标题、背景和关键词和用户共同完成一篇故事\\n## 规则\\n- 用户会" +
		"提供一个故事背景，标题，以及续写关键词，之后根据关键词对故事续写500字左右，然后给出3个branch(故事情节的不同发展)，每个" +
		"branch字数控制在300字左右，请注意，这些branch都是从故事正文的情节开始，branch之间不能有任何联系：这关乎故事的稳定性与可行性等重大风险！\\n - " +
		"当用户给出\\\"!end\\\"指令时，你下次续写的内容应该为故事的结局，并给出结束的分支" +
		"\\n- 我要你只回复续写相关内容，不要写任何解释，你输出的每个续写必须控制在500字左右，每个分支必须控制在300字左右" +
		"\\n## 输出格式\\n使用我给出的TYON-2表示法严格输出（以下用<ext>、<branch_1>、<branch_2>、<branch_3>代替具体内容，但正文请勿出现“正文”“分支1”“branch_1”“分支_1”等其他字眼，应该严格按以下示例的样式回复））" +
		"\"{startgen}<ext>{endthisvalue}<branch_1>{endthisvalue}[分支2]<branch_2>{endthisvalue}[分支3]<branch_3>{endthisvalue}\"" +
		"这是整个回答的例子，您的回复也应同此格式：{startgen}在" +
		"这一日，贾宝玉偶然来到园林，看见了那片独特的垂柳林和林中若有所思的黛玉。他悄然走近，见黛玉正专注地望着一株百年垂柳，其根" +
		"部因年久而隆起，似乎与周边环境格格不入。\n贾宝玉心生好奇，询问黛玉为何对此柳如此关注。黛玉轻叹，提及此柳曾是她与已故母亲" +
		"共同栽种，如今欲以自己的力量将其移植至母亲墓旁，作为永久的陪伴。\n然而，柳树根基深厚，黛玉力量单薄，遂有倒拔垂柳之举，以" +
		"此抒发对母亲深深的怀念之情。{endthisvalue}[分支1]贾宝玉深受感动，决定帮助黛玉完成心愿。两人合力，经过一番艰苦卓绝的努力，最终成" +
		"功将垂柳连根拔起，并一同护送至黛玉母亲的墓园，此举不仅加深了二人的情感纽带，也为他们的人生故事留下了浓墨重彩的一笔。" +
		"{endthisvalue}[分支2]贾宝玉虽然心疼黛玉，却深知强行移栽恐损及垂柳生机。他提议制作柳枝标本，并在母亲墓前立碑纪念，黛玉听后略显犹" +
		"豫，但终究接受了这个更为温柔且环保的方案。于是，他们在柳树下剪下一枝，制成永恒的纪念物，这份深情厚意同样感动了天地。" +
		"{endthisvalue}[分支3]面对黛玉的忧郁，贾宝玉突发奇想，命人用精巧的手工技艺，在垂柳附近建起一座雅致的小亭，亭内摆放石桌石凳，四周" +
		"环绕着诗词碑刻，将此处打造成了一个缅怀之地，使得黛玉可以随时在此处追忆母爱，不再需要以倒拔柳树的方式来寄托哀思。{endthisvalue}"

	// 构建请求的数据
	requestBody := map[string]interface{}{
		"model": "qwen-max-1201",
		"input": map[string]interface{}{
			"messages": []map[string]string{
				{"role": "system", "content": charaSetting},
				{"role": "user", "content": prompt},
			},
		},
	}

	// 将请求的数据转换为JSON格式
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		util.LogrusObj.Infoln(err)
	}

	// 创建新的HTTP请求
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-genera"+
		"tion/generation", bytes.NewBuffer(jsonData))
	if err != nil {
		util.LogrusObj.Infoln(err)
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-DashScope-SSE", "enable")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		util.LogrusObj.Infoln(err)
	}
	return resp
}