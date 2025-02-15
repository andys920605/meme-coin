# meme-coin API Documentation

## 專案架構

```
.
├── cmd：啟動程式的進入點
├── config
├── internal
│   ├── domain
│   │   ├── model：定義領域模型
│   │   └── service：定義領域服務
|   |── mock：模擬實作，用於測試
│   ├── north
│   │   ├── local：應用服務（本地服務）
│   │   ├── message：應用服務的請求、回應和事件
│   │   └── remote：對外的契約，如 Restful API、gRPC、Subscriber 和 CronJob
│   └── south
│       ├── adapter：對外的實作，如 Repository、第三方 Client、File 和 Publisher
│       ├── message：對外的請求、回應
│       └── port：對外的介面定義
└── pkg：infra套件
```

---

## 架構介紹
![这是图片](https://image-static.segmentfault.com/274/745/2747459847-8eb326e46027217b_fix732 "Magic Gardens")

[菱形對稱架構介紹](https://segmentfault.com/a/1190000040533813/ "游標顯示")
菱形對稱架構是一種結合 六邊形架構 (Hexagonal Architecture)、Clean Architecture 與 領域驅動設計 (DDD) 的架構模式，旨在讓系統的內部結構保持清晰，並提高可維護性與可擴展性。

架構概念
這個架構的核心思想是：

1. 將業務邏輯獨立於基礎設施與框架，確保核心業務邏輯不被技術細節耦合。
2. 輸入與輸出對稱 (Input & Output Symmetry)，也就是說，無論是外部 API、資料庫、訊息隊列或是 UI 介面，它們與業務邏輯的交互方式都是一致的，保持架構的統一性與一致性。
3. 以 DDD 的方式組織代碼，使用 聚合 (Aggregate)、領域服務 (Domain Service)、應用層 (Application Layer) 來確保業務邏輯的可讀性與擴展性。

適用場景
高複雜度的業務邏輯，如金融、電商、廣告系統
需要長期維護的企業級應用
需要高度解耦與可擴展的架構設計
總結
菱形對稱架構的目標是 透過清晰的分層來降低耦合，讓系統更容易維護與擴展。
它保留了 六邊形架構的開放性、Clean Architecture 的清晰分層，並結合 DDD 的業務驅動，讓整體架構更符合現代後端開發的需求。


## 快速上手

### 使用 Docker Compose 啟動服務

```bash
docker-compose up -d
```

## API Interface
### Create meme coin API
- Endpoint：http://127.0.0.1:8080/srv/meme-coins
- HTTP Method：POST
- Request Body example：
```bash
{
    "name":"gogo",
    "description":"test description"
}
```
Response：
```
status code: 200
{
    "code": 0,
    "msg": "ok",
    "data": {
        "id": "1125823954557076224"
    }
}
```

### Get meme coin API
- Endpoint：http://127.0.0.1:8080/srv/meme-coins/{id}
- HTTP Method：GET

Response：
```
status code: 200
{
    "code": 0,
    "msg": "ok",
    "data": {
        "id": "1125823954557076224",
        "name": "gogo",
        "description": "test description",
        "popularity_score": 0,
        "created_at": "2025-02-15T16:05:37Z"
    }
}
```

### Update meme coin API
- Endpoint：http://127.0.0.1:8080/srv/meme-coins/{id}
- HTTP Method：PUT
- Request Body example：
```bash
{
    "description":"asdf"
}
```
Response：
```
status code: 200
{
    "code": 0,
    "msg": "ok"
}
```

### Delete meme coin API
- Endpoint：http://127.0.0.1:8080/srv/meme-coins/{id}
- HTTP Method：DELETE
Response：
```
status code: 200
{
    "code": 0,
    "msg": "ok"
}
```

### Poke meme coin API
- Endpoint：http://127.0.0.1:8080/srv/meme-coins/{id}/poke
- HTTP Method：POST

Response：
```
status code: 200
{
    "code": 0,
    "msg": "ok"
}
```
