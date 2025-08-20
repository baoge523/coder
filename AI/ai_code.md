## AI

https://github.com/cloudwego/eino

https://github.com/cloudwego/eino/blob/main/README.zh_CN.md

github.com/cloudwego/eino@v0.3.51

https://www.cloudwego.io/zh/docs/eino/quick_start/



### AI基础 -- 总结
使用eino实现ai总结；基于LLM，将告警信息总结成prompt，然后告诉LLM,让其总结

开发步骤:
1、查看eino的官网，引入依赖、编写sys prompt 和 user prompt
2、创建LLM，我是用的deepseek(go1.24)
3、通过占位符的方式将告警信息传入到user prompt中，让LLM帮忙分析
```text
system prompt 限制ai行为，和预期希望ai做什么事情，比如作为一种角色，在看到用户的问题后，怎么理性的分析，并合理的根据用户提供的信息给出比较合理的建议，同时不能提供到自身系统的不足等问题

user prompt 用户的问题，这个可能是比较简单的问题，也可能是比较复杂且带有相关数据信息让ai总结的
```

常见场景有：
```text
1、简单对答场景，需要结合之前的对话
2、简单问题，拓展用户问题生成更加详细的user prompt 然后向LLM提问
3、。。。
```

### agent 平台
Agent = LLM+记忆+规划技能+工具使用
智能体开发平台的架构一般包含 插件配置、Agent 配置、Agent 执行模块、插件执行模块，发布模块。

agent的三大能力：
```text
规划：智能体能够根据给定目标，自主拆解任务步骤执行执行计划，例如OpenManus，OWL等通用智能体，通过规划能力，能够有条不紊地处理复杂任务，确保每一步都朝着目标迈进。
       针对特定领域的智能体，还可以通过预定义Agent的工作流，结合大模型的意图识别和流程控制，提升Agent在处理复杂任务过程中的稳定性，类似于Dify、Coze、元器等平台。

记忆：智能体具备长期记忆能力，能够存储和调用历史信息，从而在任务执行过程中保持上下文连贯性。记忆功能使 AI Agent 能够更好地理解用户需求，并在多轮交互中提供更精准的反馈。
       与单纯的大模型不同，AI Agent 的记忆能力使其能够在复杂任务中保持状态，避免信息丢失，从而更有效地处理多轮对话和长期任务。

工具使用：LLM虽然在信息处理方面表现出色，但它们缺乏直接感知和影响现实世界的能力，工具是LLM连接外部世界的关键，智能体能够通过使用工具，例如调用API、访问数据库等等，与外部世界进行交互。
        近期爆火的MCP协议，定义了工具的开发范式，通过标准化的接口规范，使得AI Agent能够更便捷地集成各种外部工具和服务，从而大大扩展了智能体的能力边界。
```
google对agent的定义
https://drive.google.com/file/d/1oEjiRCTbd54aSdB_eEe3UShxLBWK9xkt/view?pli=1


LLM与agent的能力对比:
```text
           LLM                                                                          Agent
知识范围	仅限于训练数据中包含的内容	                                                    可通过工具接入外部系统获取扩展知识
推理能力	仅能进行单次查询响应，除非特别设计，否则无法维护会话历史和上下文连续性	                能够维护完整会话历史，支持基于用户查询和编排层决策的多轮对话
工具使用	不具备内置工具调用能力	                                                        在架构层面直接支持工具集成
逻辑处理	无内置逻辑处理层，需要用户通过简单问询或利用CoT、ReAct等推理框架构建复杂提示来引导预测	    具备完整的认知架构，能够集成CoT、ReAct或LangChain等预置智能体框架
```


mcp出来之前
```text
1. 智能体开发平台需要单独的插件配置和插件执行模型，以屏蔽不同工具之间的协议差异，提供统一的接口给 Agent 使用；

2. 开发者如果要增加自定义的工具，需要按照平台规定的 http 协议实现工具。并且不同的平台之间的协议可能不同；

3. “M×N 问题”：每新增一个工具或模型，需重新开发全套接口，导致开发成本激增、系统脆弱；

4. 功能割裂：AI 模型无法跨工具协作（如同时操作 Excel 和数据库），用户需手动切换平台。

   没有标准，整个行业生态很难有大的发展，所以 MCP 作为一种标准的出现，是 AI 发展的必然需求。
```



coze 、dify 、 腾讯云智能体开发平台
https://cloud.dify.ai/signin
https://www.coze.cn/space/7540496974827634723/develop

### mcp 通信协议 （Model Context Protocol）
mcp的全球广场
https://mcpmarket.cn/
https://mcp.so/

百度: https://sai.baidu.com/zh/
阿里: https://www.modelscope.cn/mcp?category=research-and-data&page=1

MCP 协议旨在解决大型语言模型（LLM）与外部数据源、工具间的集成难题，被比喻为“AI应用的USB-C接口“。



### agent 工具


open-ai-agent sdk
https://openai.github.io/openai-agents-python/

agent 可以看成一个领域的"专家"，它可以通过提供的工具信息，在用户发出提问时合理的选择使用哪个工具

多agent互通

在mcp出来前，由于agent的工具配置，没有一个统一的标准，各种的工具提供者都按照自己喜欢的方式提供工具，导致在agent适配工具时成为了瓶颈；就相当于各说各的语言，谁都听不懂

在mcp统一通信协议后，工具按照mcp协议开发，那么所有的agent在使用工具的时候，就很好的适配了