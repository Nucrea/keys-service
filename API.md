# HTTP API Сервиса Генерации Ключей

## GET /health
Проверка доступности сервиса (в том числе для K8S)

**params** : нет

**body** : нет

### 200 OK
Запрос успешно обработан

**body** : нет

## GET /key
Получение ключа шифрования RSA

**params** : нет

**body** : нет

### 200 OK
Успешный запрос

**body** : `text/plain`
```
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC3vNxy8+7oxtRl
fsWEhz8mXWOO9vws0bDAlXxUsHC7sktoteVAhvYXknn9c6YUYmDGAGsXyOWmJRuF
nPSQBrgvDPl9jBb2C9FH4sz5jR/KKC4Xdr4Pj7iNFjApmsynYABSRHD+0KpPEiAU
IuI7cK+2q2j2ffQFOfHFpJ2A97l8SFKUpK5EnoJMCq/0JSU43NS4BZ8XW3etbArz
WL8LWZye/bdyX4ZSszkISEhq/pDOzuCymtjkIl3DixEYqD+3Ejex+iFGYBphMOrF
nq45nyFcDry3Vwi4tFNCVaTGDQgbf9wDDHreY1kWE9NaSZFO5WD0imuzK+imO2OV
LdBTscH1AgMBAAECggEATCaQT2xeRkex2+bwwf6z45itHKGs3n6B/S93ejouXvyE
WH+a5mWJNTfWK391A9nKNgYSXKc81uFmSLhGfDLsv0NnW+tl0NARCvQq8/rThBun
Dr5+A0ETwhXCHFw2GeCmAXKwSkv9agwfE3lpviz0BoeAVzjYnBnp5HpXdQ426BT1
iSjXJvmwRgrl1ly553rlkis3kZSa4mNS/r0A0jEvzzT4eif0DOc1v6B24zoUePEW
0pVo614T23is9Y0nouMtKdQCwjjjh2BbmC6KjTkOGofgiEzZsNsg5knurrdSLMv3
OqRh+zj+AyQxNjaaJ0MhN8sPewRaFDVPR/CFQ2SH+QKBgQDQVDnFRtsx866m8Dzi
Isb/5vRSMPETv5spHp81WeUalltes+RIuH+Z5khyr9CCPHHm7wLZp3IYttbTHyDQ
ZlTVJtNhtm0FhR0iYW1PhZ3ocQ4uRW/WbKTWNqtz1djD88YcvTp3bqAlPOwiTI1i
Ln6dkKt+Gsa5QaF18hY2w+NLPwKBgQDhyBhrkOWlgulVvb4mIYUEShhaS3fTb3ad
sCfyF2BR2Yoz9RXdNwVIUkDEj9W/hk2r+UalTAqzJ+sSgTpnMNG+CAVIxK0JqQ1A
Z0OcFJ0xRNJruMjPgft0D0axQxZEh/Pq54l5KF5wOR4mdzOrf86tkncNqH6fQvqY
XgwIlpgpywKBgQCGNXxyJs+XQpFFYocWd6kTusmzGWx1eH6Q4vVV/W+mzS5XuDRc
1N7/WmdZ2wHMpPwL9fY0GbdbTI7gu7D8ELCeEMEktc1OPQ8j0vgEvuOXlx23mWwP
Cza1+cpCeYWH10fNw+oiftYUp0bIYeDDW4ieIVEZkE5tkmZeAXNmHJQVKQKBgDlA
LcEIytJ/MX+GT3MHyNzflPFAda/tcZxmkJp4hvn6OWsXWGXxj6tZAAdXmZGpEoTq
/pjngUcQdjEJB7Am1uhizEQ5as8qSKvuA1zOdVWK5/hcsL69bO9u+DP2mOzjtFv6
Pge0zs2SDi0eyMFR9SxaGUojUYg8yaJdJpv+47KlAoGAJTGO+42Ee09sI7YmkSqi
bnwtytl4R+rJ6DTM3oVy8DupHKBm6Npug4Pqi6AKTdQ0vIckR1yUrvCNlxye0YLX
eqw+EG2TVsZQWPy06PLhPmSUc20WNtmvFmTXIGToMyqFzqZYJh44YbX8Rrn+r+p4
48nLaAerXnKLfTSegBqRpV0=
-----END PRIVATE KEY-----
```

### 429 TOO MANY REQUESTS
Сервис не успевает генерировать ключи в нужном объеме. Часть соединений, долго висящая в ожидании ключа, отваливается с данной ошибкой

**body** : нет

---