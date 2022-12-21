# Local Logging Server

## Log 저장 위치

`HOME/logs`

## 통신 프로토콜
```
TEST_LOG // 테스트용
SYNC_LOG // Batch SyncAll 성능 로그
API_LOG // API Request 로그

PROGRAM_EXIT // 프로그램 종료
```

## 데이터 전달 포맷

```json
{
    "protocol": "" // ... protocol,
    "data": {
        // ... some object
    }
}
```

