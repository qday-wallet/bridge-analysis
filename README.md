## bridge-analysis

QDayBridge 的组件之一

### Deploy 

- Download

  下载必要的配置文件和模版

     ```
     wget -O docker.zip https://github.com/qday-wallet/bridge-relay/releases/download/1.0.0/docker.zip && unzip docker.zip

     ```
- Update config

  根据自己的需要，更改配置文件（可以参考模版）

    ``````
    {
    
      "DB": { //trust relay 的db 
        "Addr": "190.92.213.101",
        "Port": 5432,
        "User": "postgres",
        "Password": "123456789",
        "DbName": "postgres"
      },
      "Log": { //日志文件
        "Path": "./log/analysis",
        "Delay": 2
      },
      "Port": 9092, //api 端口
      "Root": "./bridge" //api root path
    }
    
    ``````


- Start Analysis

  ```
  //start 
  docker compose -f docker-compose.yml up -d
  
  //down
  docker compose -f docker-compose.yml down -v
  ```



### API
        

``````
    
    //查询历史交易记录（包含 入账、出账的交易，按交易发起时间倒序排列，UTC时间）
    curl --location --request POST 'http://190.92.213.101:9092/bridge/query' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "address":"0x30ef9dF39C10C57a478f4c6733c3f210CE17C662",//当前账户
    "pageSize":5,
    "pageNumber":1
    }'
    
    //查询异常交易
    curl --location --request POST 'http://localhost:9092/bridge/monitor' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "startTime": "2024-10-10 01:00:00",
        "endTime": "2024-10-25 01:00:00",
        "status": [ //异常交易状态码
            9,
            100
        ]
    }'
    
    const (
        TxFail            = 101
        TxSuccess         = 100
        DepositSuccess    = 2
        Voting            = 3
        VoteFail          = 4
        VoteSuccess       = 5
        MintSuccess       = 6
        MintFail          = 7
        Locking           = 9
        WithdrawFail      = 10
    )
    
    //QDay 入账统计（按小时为单位）
    curl --location --request POST 'http://localhost:9092/bridge/analysis/income' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "startTime": "2024-10-10 01:00:00",
        "endTime": "2024-10-25 01:00:00"
    }'
    
    //QDay 出账统计（按小时为单位）
    curl --location --request POST 'http://localhost:9092/bridge/analysis/pay' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "startTime": "2024-10-10 01:00:00",
        "endTime": "2024-10-25 01:00:00"
    }'
    
``````