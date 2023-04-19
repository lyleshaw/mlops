namespace go service

// 公共结构体
struct BaseReq {
}

struct BaseResp {
  1: i32 code,
  2: string message
}

// Request 结构体
struct LoginReq {
  1: string email,
  2: string password
}

struct RegisterReq {
  1: string email,
  2: string username,
  3: string password,
  4: string code
}

struct AllUserReq {
}

struct DeleteUserReq {
  1: i32 user_id
}

struct CreateAppReq {
  1: string app_name,
  2: string app_image
}

struct AppInfoReq {
  1: i32 app_id
}

struct UpdateAppReq {
  1: i32 app_id,
  2: optional string app_name,
  3: optional string app_image,
  4: bool is_reload
}

struct DeleteAppReq {
  1: i32 app_id
}

struct SendCodeReq {
  1: string email
}

// Response 结构体
struct LoginResp {
  1: i32 code,
  2: string message
  3: LoginData data
}

struct UserListResp {
  1: i32 code,
  2: string message
  3: list<User> data
}

struct UserResp {
  1: i32 code,
  2: string message
  3: User data
}

struct AppListResp {
  1: i32 code,
  2: string message
  3: list<Application> data
}

struct ApplicationResp {
  1: i32 code,
  2: string message
  3: Application data
}

// 实体结构体
struct LoginData {
    1: string token,
    2: string expire
}

struct User {
  1: i32 user_id,
  2: string username,
  3: string email,
  4: string avatar,
  5: bool is_active,
  6: bool is_admin,
  7: string create_at,
  8: string update_at
}

struct Application {
  1: i32 app_id,
  2: string app_name,
  3: string app_status,
  4: i32 app_port,
  5: string app_image,
  6: i32 user_id,
  7: string create_at,
  8: string update_at
}

struct EmailCode {
  1:string email
}

service UserService {
  UserResp Login(1: LoginReq req) (api.post="/api/v1/user/login")

  UserResp Register(1: RegisterReq req) (api.post="/api/v1/user/register")

  UserResp GetUser() (api.get="/api/v1/user/me")

  UserListResp AllUser(1: AllUserReq req) (api.get="/api/v1/user/all")

  BaseResp DeleteUser(1: DeleteUserReq req) (api.delete="/api/v1/user")
}

service AppService {
  ApplicationResp CreateApp(1: CreateAppReq req) (api.post="/api/v1/app")

  AppListResp MyApp(1: AllUserReq req) (api.get="/api/v1/app/me")

  ApplicationResp AppInfo(1: AppInfoReq req) (api.get="/api/v1/app")

  ApplicationResp UpdateApp(1: UpdateAppReq req) (api.patch="/api/v1/app")

  BaseResp DeleteApp(1: DeleteAppReq req) (api.delete="/api/v1/app")
}

service EmailService {
  BaseResp SendCode(1: SendCodeReq req) (api.post="/api/v1/email/code")
}