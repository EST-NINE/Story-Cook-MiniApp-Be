package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ncuhome/story-cook/config"
	"github.com/ncuhome/story-cook/pkg/util"
	"io"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func extendStory(title string, background string, keywords string) string {
	apiKey := config.ApiKey
	charaSetting := `## 指令\n你是一个作家，你现在要根据用户给出的故事标题、背景和关键词和用户共同完成一篇故事\n## 规则\n- 用户会
提供一个故事背景，标题，以及续写关键词，之后根据关键词对故事续写500字左右，然后给出3个branch(故事情节的不同发展)，每个branch字数控
制在300字左右\n - 当用户给出\"!end\"指令时，你下次续写的内容应该为故事的结局，并给出结束的分支\n- 我要你只回复续写相关内容，不要写
任何解释，你输出的每个续写必须控制在500字左右，每个分支必须控制在300字左右\n## 输出格式\n要求输出结构化 JSON 对象 (**注意例子中'}'
反大括号的位置** ，json 对象的 KEY 要严格与例子一致， **不要把branch_1，翻译成分支1或分支_1**)，符合下面 TypeScript：interface 
{ext?: string; branch_1?: string; branch_2?: string; branch_3?: string; }  这是例子{"ext": "在这一日，贾宝玉偶然来到园林，
看见了那片独特的垂柳林和林中若有所思的黛玉。他悄然走近，见黛玉正专注地望着一株百年垂柳，其根部因年久而隆起，似乎与周边环境格格不入。
贾宝玉心生好奇，询问黛玉为何对此柳如此关注。黛玉轻叹，提及此柳曾是她与已故母亲共同栽种，如今欲以自己的力量将其移植至母亲墓旁，作为永久
的陪伴。然而，柳树根基深厚，黛玉力量单薄，遂有倒拔垂柳之举，以此抒发对母亲深深的怀念之情。","branch_1":"贾宝玉深受感动，决定帮助黛玉
完成心愿。两人合力，经过一番艰苦卓绝的努力，最终成功将垂柳连根拔起，并一同护送至黛玉母亲的墓园，此举不仅加深了二人的情感纽带，也为他们
的人生故事留下了浓墨重彩的一笔。","branch_2":"贾宝玉虽然心疼黛玉，却深知强行移栽恐损及垂柳生机。他提议制作柳枝标本，并在母亲墓前立碑
纪念，黛玉听后略显犹豫，但终究接受了这个更为温柔且环保的方案。于是，他们在柳树下剪下一枝，制成永恒的纪念物，这份深情厚意同样感动了天地
。","branch_3":"面对黛玉的忧郁，贾宝玉突发奇想，命人用精巧的手工技艺，在垂柳附近建起一座雅致的小亭，亭内摆放石桌石凳，四周环绕着诗词
碑刻，将此处打造成了一个缅怀之地，使得黛玉可以随时在此处追忆母爱，不再需要以倒拔柳树的方式来寄托哀思。"}`
	prompt := fmt.Sprintf("标题：%s 故事背景：%s 关键词：%s", title, background, keywords)

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
		return fmt.Sprintf("Error marshalling JSON: %v", err)
	}

	// 创建新的HTTP请求
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-genera"+
		"tion/generation", bytes.NewBuffer(jsonData))
	if err != nil {
		util.LogrusObj.Infoln(err)
		return fmt.Sprintf("Error creating request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return fmt.Sprintf("Error sending request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return fmt.Sprintf("Error reading response body: %v", err)
	}

	// 解析JSON响应
	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		return fmt.Sprintf("Error unmarshalling JSON response: %v", err)
	}

	// 提取output字段
	output, ok := responseMap["output"].(map[string]interface{})
	if !ok {
		return fmt.Sprintf("Output field not found or not a map")
	}

	// 提取text字段
	text, ok := output["text"].(string)
	if !ok {
		return fmt.Sprintf("Text field not found or not a string")
	}
	return text
}
