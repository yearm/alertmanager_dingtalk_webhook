## Alertmanager Dingtalk Webhook

webhook support send prometheus 2.0 alert message to dingtalk.

## Use

```
./cmd/webhook -dingtalkRobot="https://oapi.dingtalk.com/robot/send?access_token=xxx"
```

- alertmanager.yml configuration example

```
global:
  resolve_timeout: 5m

route:
  group_by: ['alertname']
  group_wait: 10s
  group_interval: 10s
  repeat_interval: 1h
  receiver: 'webhook'
receivers:
- name: 'webhook'
  webhook_configs:
  - url: 'http://localhost:8060/webhook'
    send_resolved: true
```

