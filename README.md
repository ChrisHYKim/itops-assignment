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

![alt text](<Screenshot From 2025-06-15 18-46-02.png>)