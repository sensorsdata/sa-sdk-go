## 神策简介

[**神策数据**](https://www.sensorsdata.cn/)（Sensors Data），隶属于神策网络科技（北京）有限公司，是一家专业的大数据分析服务公司，大数据分析行业开拓者，为客户提供深度用户行为分析平台、以及专业的咨询服务和行业解决方案，致力于帮助客户实现数据驱动。神策数据立足大数据及用户行为分析的技术与实践前沿，业务现已覆盖以互联网、金融、零售快消、高科技、制造等为代表的十多个主要行业、并可支持企业多个职能部门。公司总部在北京，并在上海、深圳、合肥、武汉等地拥有本地化的服务团队，覆盖东区及南区市场；公司拥有专业的服务团队，为客户提供一对一的客户服务。公司在大数据领域积累的核心关键技术，包括在海量数据采集、存储、清洗、分析挖掘、可视化、智能应用、安全与隐私保护等领域。 [**More**](https://www.sensorsdata.cn/about/aboutus.html)

## 安装与更新

使用以下指令获取 `Sensors Analytics SDK `

```
go get github.com/sensorsdata/sa-sdk-go
```

使用以下指令更新 `Sensors Analytics SDK `

```
go get -u github.com/sensorsdata/sa-sdk-go
	
```

## 示例

```golang
import sdk "github.com/sensorsdata/sa-sdk-go"

// Gets the url of Sensors Analytics in the home page.
SA_SERVER_URL = 'YOUR_SERVER_URL'

// Initialized the Sensors Analytics SDK with Default Consumer
consumer = sdk.InitDefaultConsumer(SA_SERVER_URL)
sa = sdk.InitSensorsAnalytics(consumer)

properties := map[string]interface{}{
     "price": 12,
     "name": "apple",
     "somedata": []string{"a", "b"},
}

// Track the event 'ServerStart'
sa.track("ABCDEFG1234567", "ServerStart", properties, false)

sa.Close()
```

**更多示例**
([Examples](_examples)) 

## 更多帮助
可以查看官方帮助文档： [Golang SDK 使用说明](http://www.sensorsdata.cn/manual/golang_sdk.html)<br>
或者加入神策官方 SDK QQ 讨论群：<br><br>
![ QQ 讨论群](https://github.com/sensorsdata/sa-sdk-android/blob/master/docs/qrCode.jpeg)

## License

Copyright 2015－2020 Sensors Data Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

**禁止一切基于神策数据开源 SDK 的所有商业活动！**