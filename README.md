### 기획
* 해외에서도 메뉴를 쉽게 이해하고 즐길 수 있도록 돕는 <여행자 맞춤 음식 안내 서비스>

### 각종 명령어
* [실행] 프로젝트 루트에서 `go run ./cmd/api-server`
* [의존성 꼬임 해결] 프로젝트 루트에서 `go mod tidy`
* [의존성 추가] 프로젝트 루트에서 `go get google.golang.org/api/option`, `go get github.com/joho/godotenv`, ...

### 프로젝트 구조 컨벤션
* https://github.com/golang-standards/project-layout

### 자잘 팁
* 프로젝트 루트에 `.env` 파일을 만들고, `GEMINI_API_KEY="실제 발급받은 api key 값"` 이렇게 적어줘야 동작함
