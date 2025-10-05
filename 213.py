import jwt
from jwt import InvalidTokenError, ExpiredSignatureError

# 你的 JWT 字串（從前端或 API header 傳過來的）
token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6bnVsbCwiZXhwIjoxNzUzNzgxMjA1LCJpYXQiOjE3NTM3ODExNDQsImlkIjoiMDE5NTQyYTEtNDY5OC03NWVhLWEwY2EtMzlhZTJiNzg4ZDhlIiwicGhvbmVfbnVtYmVyIjpudWxsLCJyb2xlX2lkcyI6bnVsbCwic3ViIjoiMDE5NTQyYTEtNDY5OC03NWVhLWEwY2EtMzlhZTJiNzg4ZDhlIiwidGVuYW50X2lkIjpudWxsLCJ0eXBlIjoxLCJ1c2VybmFtZSI6InJvb3QifQ.qtMYWR8146X0UzREchYesxopkYuO3wcKRACIZmtunfk"

# 產生 JWT 時用的密鑰
secret_key = "123456"

try:
    # 驗證並解析
    payload = jwt.decode(token, secret_key, algorithms=["HS256"])

    print("✅ Token 驗證成功")
    print("使用者名稱:", payload.get("username"))
    print("使用者 ID:", payload.get("id"))
    print("發行時間 (iat):", payload.get("iat"))
    print("過期時間 (exp):", payload.get("exp"))

except ExpiredSignatureError:
    print("❌ Token 已過期")
except InvalidTokenError as e:
    print("❌ Token 無效:", str(e))
