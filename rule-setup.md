# 规则配置方法

### 配置读取流程

Spacer 的配置读取流程如下：

- 通过 `conf` 或者 `-c` 参数显式的指定要使用配置的文件。
- 如果没有显式的指定配置文件，则尝试读取 `${HOME}/.speacer.yaml` 文件的配置。
- 如果 `${HOME}/.speacer.yaml` 文件不存在，则不加载任何额外的配置，仅使用默认的配置。
 
### 启用规则

如果你想要启用一些默认没有启用的规则，可以在配置文件的 `enableds` 列表中添加它们：

```yaml
enables:
    - c
    - d
    - e
```

配置文件中的 `enableds` 规则会追加至`默认列表`中，而不是`仅启用`配置文件中的规则。

### 禁用规则

如果你想要明确地禁用一些规则，可以在配置文件的 `disableds` 列表中添加它们：

```yaml
disables:
    - a
    - c
    - e
```

配置文件中的 `disables` 列表具有`最高优先级`。这意味着，即使该则规则是默认启用的，或者你在 `enables` 列表中显式的启用了它，只要在该列表中出现的规则，总是会被禁用。