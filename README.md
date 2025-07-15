# 이슈 관리 API

백엔드 개발자 채용 과제 - 이슈 관리 시스템 REST API

## 프로젝트 개요

Go 언어와 Gin 프레임워크를 사용하여 구현한 이슈 관리 REST API입니다.
이슈의 생성, 조회, 수정 기능을 제공하며, 담당자 할당과 상태 관리 기능을 포함합니다.

## 기술 스택

- **언어**: Go 1.21+
- **프레임워크**: Gin
- **데이터 저장**: 메모리 (In-memory)
- **포트**: 8080

## 실행 방법

### 1. 필수 요구사항

- Go 1.21 이상 설치
- Git 설치

### 2. 프로젝트 클론 및 의존성 설치

```bash
git clone https://github.com/kimauto/issue-api.git
cd issue-api
go mod tidy
go get github.com/gin-gonic/gin
```

### 3. 서버 실행

```bash
go run *.go
```

### 4. 서버 실행 확인

서버가 성공적으로 시작되면 다음 메시지가 출력됩니다:

```
[GIN-debug] Listening and serving HTTP on :8080
```

## API 명세

### 기본 정보

- **Base URL**: `http://localhost:8080`
- **Content-Type**: `application/json`

### 사용자 정보

시스템에 미리 등록된 사용자:

- ID: 1, 이름: "김개발"
- ID: 2, 이름: "이디자인"
- ID: 3, 이름: "박기획"

### 이슈 상태

- `PENDING`: 대기 중
- `IN_PROGRESS`: 진행 중
- `COMPLETED`: 완료
- `CANCELLED`: 취소

### API 엔드포인트

#### 1. 이슈 생성

```http
POST /issue
```

**요청 예시:**

```json
{
  "title": "버그 수정 필요",
  "description": "로그인 페이지에서 오류 발생",
  "userId": 1
}
```

**응답 예시 (201 Created):**

```json
{
  "id": 1,
  "title": "버그 수정 필요",
  "description": "로그인 페이지에서 오류 발생",
  "status": "IN_PROGRESS",
  "user": {
    "id": 1,
    "name": "김개발"
  },
  "createdAt": "2025-07-15T10:00:00Z",
  "updatedAt": "2025-07-15T10:00:00Z"
}
```

#### 2. 이슈 목록 조회

```http
GET /issues
GET /issues?status=PENDING
```

**응답 예시 (200 OK):**

```json
{
  "issues": [
    {
      "id": 1,
      "title": "버그 수정 필요",
      "description": "로그인 페이지에서 오류 발생",
      "status": "PENDING",
      "createdAt": "2025-07-15T10:00:00Z",
      "updatedAt": "2025-07-15T10:05:00Z"
    }
  ]
}
```

#### 3. 이슈 상세 조회

```http
GET /issue/{id}
```

**응답 예시 (200 OK):**

```json
{
  "id": 1,
  "title": "버그 수정 필요",
  "description": "로그인 페이지에서 오류 발생",
  "status": "PENDING",
  "createdAt": "2025-07-15T10:00:00Z",
  "updatedAt": "2025-07-15T10:05:00Z"
}
```

#### 4. 이슈 수정

```http
PATCH /issue/{id}
```

**요청 예시:**

```json
{
  "title": "로그인 버그 수정",
  "status": "IN_PROGRESS",
  "userId": 2
}
```

**응답 예시 (200 OK):**

```json
{
  "id": 1,
  "title": "로그인 버그 수정",
  "description": "로그인 페이지에서 오류 발생",
  "status": "IN_PROGRESS",
  "user": {
    "id": 2,
    "name": "이디자인"
  },
  "createdAt": "2025-07-15T10:00:00Z",
  "updatedAt": "2025-07-15T10:10:00Z"
}
```

## 비즈니스 규칙

### 이슈 생성 시

- 담당자가 있는 경우: 상태를 `IN_PROGRESS`로 설정
- 담당자가 없는 경우: 상태를 `PENDING`으로 설정
- 존재하지 않는 사용자를 담당자로 지정할 수 없음

### 이슈 수정 시

- `COMPLETED` 또는 `CANCELLED` 상태의 이슈는 수정 불가
- 담당자 없이 `IN_PROGRESS` 또는 `COMPLETED` 상태로 변경 불가
- `PENDING` 상태에서 담당자 할당 시 자동으로 `IN_PROGRESS`로 변경
- 담당자 제거 시 상태를 `PENDING`으로 변경

## API 테스트 방법

### 1. curl 사용

```bash
# 이슈 생성
curl -X POST http://localhost:8080/issue \
  -H "Content-Type: application/json" \
  -d '{"title":"테스트 이슈","description":"테스트 설명","userId":1}'

# 이슈 목록 조회
curl -X GET http://localhost:8080/issues

# 상태별 필터링
curl -X GET "http://localhost:8080/issues?status=PENDING"

# 이슈 상세 조회
curl -X GET http://localhost:8080/issue/1

# 이슈 수정
curl -X PATCH http://localhost:8080/issue/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"COMPLETED"}'
```

### 2. Postman 사용

1. Postman 실행
2. 새 요청 생성
3. 위의 엔드포인트와 JSON 데이터를 사용하여 테스트

## 에러 응답

모든 에러는 다음 형식으로 응답됩니다:

```json
{
  "error": "에러 메시지",
  "code": 400
}
```

### 주요 에러 코드

- `400 Bad Request`: 잘못된 요청 데이터
- `404 Not Found`: 존재하지 않는 리소스
- `500 Internal Server Error`: 서버 내부 오류

## 파일 구조

```
.
├── main.go           # 메인 서버 실행 파일
├── model.go         # 데이터 모델 정의
├── handler.go       # API 핸들러 함수들
├── data.go        # 데이터 저장 및 관리
├── utils.go          # 유틸리티 함수들
├── go.mod           # Go 모듈 파일
├── go.sum           # 의존성 체크섬
└── README.md        # 프로젝트 설명서
```

## 개발자 정보

이 프로젝트는 백엔드 개발자 채용 과제로 개발되었습니다.
