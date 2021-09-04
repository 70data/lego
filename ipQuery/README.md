## 使用说明

启动时加载至内存

```
ipQuery.LoadIPdata("IP.txt")

func main() {
	ipQuery.LoadIPdata("IP.txt")
}
```

使用时直接解析

```
province := ipQuery.ProvinceIsp(ipAddr)[0]
isp := ipQuery.ProvinceIsp(ipAddr)[1]
```
