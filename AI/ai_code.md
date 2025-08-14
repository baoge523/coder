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

### mcp 通信协议