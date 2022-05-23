import hashlib
import base64
import hmac


def gen_signed_nonce(ssecret: str, nonce: str) -> str:
    """Nonce signed with ssecret."""
    m = hashlib.sha256()
    m.update(base64.b64decode(ssecret))
    m.update(base64.b64decode(nonce))
    return base64.b64encode(m.digest()).decode()


def gen_signature(url: str, signed_nonce: str, nonce: str, data: str) -> str:
    """Request signature based on url, signed_nonce, nonce and data."""
    sign = '&'.join([url, signed_nonce, nonce, 'data=' + data])
    signature = hmac.new(key=base64.b64decode(signed_nonce),
                         msg=sign.encode(),
                         digestmod=hashlib.sha256).digest()
    return base64.b64encode(signature).decode()


test = gen_signed_nonce("7AL5MZR1mL4R8UmaEn8QZA==", "mWX0lu1kqZENc6QB")

print(test)

print(gen_signature("123", test, "mWX0lu1kqZENc6QB", "123"))
