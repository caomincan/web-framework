# web-framwork

## 前置要求

### 安装测试工具

在运行单元测试脚本之前，需要在 Mac 上安装以下工具：

```bash
# 一次性设置安装目录
sudo GOBIN=/usr/local/bin go install github.com/wadey/gocovmerge@latest
sudo GOBIN=/usr/local/bin go install github.com/t-yuki/gocover-cobertura@latest
```

验证安装：
```bash
command -v gocovmerge
command -v gocover-cobertura
```