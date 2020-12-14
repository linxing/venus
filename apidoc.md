<a name="top"></a>
#  v0.0.0



 - [User](#User)
   - [更新手机号码/昵称](#更新手机号码/昵称)
   - [获取用户个人信息](#获取用户个人信息)
   - [获取refresh token](#获取refresh-token)
   - [手机号码登录](#手机号码登录)
   - [用户更新个人密码](#用户更新个人密码)
   - [用户名登录](#用户名登录)
   - [用户注册](#用户注册)

___


# <a name='User'></a> User

## <a name='更新手机号码/昵称'></a> 更新手机号码/昵称
[Back to top](#top)

```
POST /venus/user/info
```

### Headers - `Header`

| Name    | Type      | Description                          |
|---------|-----------|--------------------------------------|
| token | `String` | <p>Token</p> |

### Parameters - `Parameter`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| phone_number | `String` | <p>新手机号码</p> |
| nickname | `String` | <p>用户昵称</p> |

### Success response

#### Success response - `Success 200`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| success | `String` | <p>状态码</p> |

### Success response example

#### Success response example - `Success-Response:`

```json
HTTP/1.2 200 OK
{
 "success": "ok"
}
```

## <a name='获取用户个人信息'></a> 获取用户个人信息
[Back to top](#top)

```
GET /venus/user/info
```

### Parameters - `Parameter`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| nickname | `String` | <p>昵称</p> |
| avatar | `String` | <p>头像地址</p> |
| position | `String` | <p>职位名称</p> |
| department_name | `String` | <p>部门名称</p> |
| last_login_at_sec | `Number` | <p>上一次登录时间</p> |

### Success response example

#### Success response example - `Success-Response:`

```json
HTTP/1.2 200 OK
{
 "nickname": "test",
 "avatar": "test",
 "position": "test",
 "department_name": "gov",
 "last_login_at_sec": 111112222331,
 "internal_phone": "10020"
 "role_id": 1,
 "user_id": 1
}
```

## <a name='获取refresh-token'></a> 获取refresh token
[Back to top](#top)

```
POST /venus/token/refresh
```

### Headers - `Header`

| Name    | Type      | Description                          |
|---------|-----------|--------------------------------------|
| token | `String` | <p>Token HTTP/1.2 200 OK { &quot;token&quot;: &quot;token&quot; }</p> |

## <a name='手机号码登录'></a> 手机号码登录
[Back to top](#top)

```
POST /venus/auth/user/signin_with_phone
```

### Parameters - `Parameter`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| phone_number | `String` | <p>手机号码</p> |
| password | `String` | <p>密码</p> |

### Success response

#### Success response - `Success 200`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| token | `String` | <p>Token</p> |

### Error response example

#### Error response example - `Error-Response:`

```json
HTTP/1.1 403 Bad Request
{
    "errorCode": 11003,
    "errorMessage": "invalid_oauth_code"
}
```

## <a name='用户更新个人密码'></a> 用户更新个人密码
[Back to top](#top)

```
POST /venus/user/password
```

### Headers - `Header`

| Name    | Type      | Description                          |
|---------|-----------|--------------------------------------|
| token | `String` | <p>Token</p> |

### Parameters - `Parameter`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| password | `String` | <p>用户密码</p> |
| confirm_password | `String` | <p>用户确认密码</p> |

### Success response

#### Success response - `Success 200`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| success | `String` | <p>状态码</p> |

### Success response example

#### Success response example - `Success-Response:`

```json
HTTP/1.2 200 OK
{
 "success": "ok"
}
```

## <a name='用户名登录'></a> 用户名登录
[Back to top](#top)

```
POST /venus/auth/user/signin_with_name
```

### Parameters - `Parameter`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| user_name | `String` | <p>用户名</p> |
| password | `String` | <p>密码</p> |

### Success response

#### Success response - `Success 200`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| token | `String` | <p>Token.</p> |

### Error response

#### Error response - `Error 4xx`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| UserNotFound |  | <p>User not found</p> |
| InvalidOAuthCode |  | <p>Invalid Token</p> |

### Error response example

#### Error response example - `Error-Response:`

```json
HTTP/1.1 404 Not Found
{
    "errorCode": 10010,
    "errorMessage": "user_not_found"
}
```

#### Error response example - `Error-Response:`

```json
HTTP/1.1 404 Bad Request
{
    "errorCode": 11003,
    "errorMessage": "invalid_oauth_code"
}
```

## <a name='用户注册'></a> 用户注册
[Back to top](#top)

```
POST /venus/auth/user/create
```

### Parameters - `Parameter`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| user_name | `String` | <p>用户名</p> |
| password | `String` | <p>密码</p> |
| nickname | `String` | <p>昵称</p> |
| phone_number | `String` | <p>手机号码</p> |
| internal_number | `String` | <p>内部电话号码</p> |

### Error response

#### Error response - `Error 4xx`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| UserExists |  | <p>The user account already exists</p> |

### Error response example

#### Error response example - `Error-Response:`

```json
HTTP/1.1 400 Bad Request
{
   "errorCode": 10011,
   "errorMessage": "user_exists"
}
```
