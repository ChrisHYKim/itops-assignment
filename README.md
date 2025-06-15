# IT Operations Assignment - Issue Management Backend API & SQL Queries
## 🚀 주요 기능

### 백엔드 API (Go)
* 이슈 생성
* 이슈 목록 조회 (상태 필터링 포함)
* 특정 이슈 상세 조회
* 이슈 정보 수정 (제목, 설명, 상태, 담당자)
* **데이터 영속성**: 
Oracle 데이터베이스를 사용하여 이슈 데이터를 저장하고 관리합니다.
## 💻 사용 기술 스택

### 백엔드
* **Go 언어**: 애플리케이션 로직 구현
* **Oracle DB**: 데이터베이스
* **`gorilla/mux`**: HTTP 라우팅
* **`database/sql`**: 데이터베이스 연동
* **`github.com/go-godbc/godbc`**: Oracle DB 드라이버 (ODBC 기반)
* **`github.com/rs/cors`**: CORS 처리 미들웨어

### SQL 쿼리 문제 해결
* 제공된 `FMS_HBL_MST` 및 `FMS_HBL_CNTR` 테이블 데이터를 기반으로 한 SQL 쿼리 문제 해결.
## ✨ 프로젝트 구조 (구현된 프로젝트)
itops-assignment/
├── backend/
│   ├── main.go                     # 백엔드 애플리케이션 진입점
│   ├── go.mod                      # Go 모듈 파일
│   ├── go.sum
│   └── internal/
│       ├── api/                    # HTTP 핸들러 (API 엔드포인트 로직)
│       │   └── handlers.go
│       ├── database/               # 데이터베이스 연결 및 초기화 (Oracle 스키마 포함)
│       │   └── database.go
│       ├── model/                  # 데이터 모델 (Issue, User 등)
│       │   └── models.go
│       ├── repository/             # 데이터 저장소 인터페이스 및 구현 (Oracle DB 연동)
│       │   └── issue_repository.go
│       └── service/                # 비즈니스 로직 (리포지토리와 핸들러 연결)
├── sql/                            # SQL 쿼리 문제 풀이 파일
│   ├── problem1.sql
│   └── problem2.sql
└── README.md

