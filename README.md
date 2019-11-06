# bus-demo

尝试使用 grafana 提供的 Bus Mode，体会如何解决 web 框架经常遇到的组件耦合问题

grafana 中系统启动或者运行时， 通过 bus.AddHandler 函数注册 HandlerFunc 处理函数到 globalBus 对象， 然后在 api 层调用 bus.Dispatch 的时候，我们会提供一个 strcut 对象作为 msg，Dispatch 会调用注册的 HandlerFunc 函数处理 strcut 对象，并且将处理结果写入到 strcut 对象中。dispatch 机制中，每个 msg 由一个 HandlerFunc 处理。

除了 dispatch 模式外， bus 还提供了 Publish 机制， 每个 msg 可以由多个 handlerFunc 处理。实现机制和 dispatch 类似，只是可以注册多个 handlerFunc 处理对应的 event。（publish 机制中，把 msg 成为 event）。我们以 signup 为例解释 publish 机制。在用户 signup 成功时， pkg/api/signup.go 会调用 bus.Publish 方法产生 SignUpCompleted event。在用户注册完成后，我们往往会发送注册成功邮件，在 grafana 中，发送注册成功邮件是由 NotificationService 完成的。NotificationService 会监听 SignUpCompleted event, 如果有 SignUpCompleted event 产生，NotificationService 就会执行发送邮件的逻辑


output:

```
=========== dispatch ===========
[biz3] call it
[biz2] call it
[kafka-msg]: content hello
[biz1] dispatch msg hello with err: <nil>
=========== dispatch ===========
=========== publish ===========
[biz4-msg]: content hello
[kafka-msg]: content hello
[biz4] publish msg hello with err: <nil>
=========== publish ===========
=========== fuzz ===========
[kafka-msg]: content hello
[biz6] dispatch msg: hello with err: <nil>
[biz4-msg]: content hello
[kafka-msg]: content hello
[biz6] publish msg hello with err: <nil>
=========== fuzz ===========
```

