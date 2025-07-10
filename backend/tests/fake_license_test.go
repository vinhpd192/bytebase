// Copyright (c) 2025 by The A.C.E Studio Teams
//

package tests

import "testing"
import "time"
import "github.com/golang-jwt/jwt/v5"

func TestGenerateLicense(t *testing.T) {
	print("Test license")

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"kid":           "v1",
		"iss":           "bytebase",
		"aud":           "bb.license",
		"mod":           "prod",
		"sub":           "00001000.",
		"instanceCount": 999999,
		"trialing":      false,
		"plan":          "ENTERPRISE",
		"orgName":       "bb",
		"workspaceId":   "",
		"exp":           time.Now().Add(time.Hour * 24 * 365 * 10).Unix(),
		"iat":           time.Now().Unix(),
	})

	key := `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDARPQtLB6iIdy2
pfzGxseeyLPTgY+YLGuZrYlOjwRUItjGUYxsV7Jjgj6X2FMRT3xtzv2l/ksDgzrk
QV0OEIANM6zZm1Qss1Sh1+Okm4rVN0ov/lYc85XTZRWHB04rGSCxewB+qRjIQxNW
61BKvJPUYtxO2XmOnCeSEzYsGljpZQ4JDu6YTuMe9aCdIU1AsKOeQ6BbVUsg4xeU
CQvdCOEPyZwvjwaheNnxn7hL1N3k9tAIPBxJJMxrokEQqYIomriWAMd4ibVExYKy
G1kPNuFnyMwlAMGGq8Rhxs8MK6/Hz0l/NqQkav/IAugvoTJ3Zdhraqtja44yepKF
ZgNtRmO1AgMBAAECggEAGUgX2Wa/Qm90b7nkwph5wCnA55NWN9rDbUoxz1gj6BJK
8czgS5C+frIuG9ZQmeqbQG8i0DRfVb9oO4JSw5os1E+Sr2xqxZPxEXTrVIY1W7Lw
+N2XIihvb4QPsBkoUKd1spCrcHw4q4CSZu/7N1CTma5ELMQQ5U0oMN+A1IzEWy/5
o0pdCV+ZyfGyVmAlu4mIyafXviKfWWQp80htWyNdRO90PUxqal8CilmEiqEWtSfJ
K/Lb+/kvVT1is/7tOTALWhlal0uhM4UovRtlQ/7wmKats+xc9hGVToQ9QiqPG8oP
svppMb/HMUFtQ8p1H8Gtuq10ri4J9P9MgLlb93B6YQKBgQDnCf4pNwo85y7gNrZZ
0auEEnl4QTqTzVDT6ipr951HERpxc2LOva1Ano1pWJFNxCvflITmRDynlgHBV/up
MWnia/2QnAlQUJtzIuK5+2S5yKeo0qKYuEpvafm3MBaqjm3MmSmWpLczLmi2eT7R
SigdmshIgyGXWNAKMhkxX1TS5QKBgQDVCq5P9L6Tqh53y0YBzm5SJXH5dnAMxzhj
+LNllFQbJuaf6+/cubPlNxX3MQGLR1KpKcJaZn0/wFIPLhPH0j3+34+LY11esiyj
OZoFAV3kywNro/D1vevUYNf8zKedwlrN3SaFdzyWOGmQMGYvQzrsQG5b0RKKs2z2
+7aOkPAwkQKBgQC4nTOKm4CCyfibER0FCJWVI24MuZKlKBq+Ow9K1sHGV9RZc2ow
6T/Xstoso6j4CxlDvtmzZjNRdnxQko111RZW5xPmychndQHCLs4c7tS936iZLjbe
8eQIwxXXW4XBI655HqosxdHeA+eIZ7naYvdLMLN5vc7JsR4MB5w554NQQQKBgHqJ
6jmTTQsijuPYyCvB0TLu1t/mgAzZilX2dzm2V6pGdeEGlLEhH5h4VuoYBgtWuvUS
T/gve1twA8M3fakyBPbVIHvMa6hMU+CPAUhfwOIY457GbZLr6vn9uj5uePsFD7XZ
vo6GRxMg35dQP3Zv9TRu2wt09nVRPxvRLPBT7dQBAoGAZfYts2v1VuVN9FMXoa+h
hpuG0tJkPbyT4ZY5S7noSV0wueLYnaWB/HQWL6u/aBz/BnS2NchMzGkFRUgusYLs
6hAGrUncIRgNQS+K1XPPEv1GWBJFl7fz9Y7WM8pJraxU0E/cS0q+QimfYEDXb2U+
hGQj2B+rqXZ76ZkexuMgCwc=
-----END PRIVATE KEY-----`

	pkey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
	if err != nil {
		t.Errorf("failed to parse private key: %v", err)
	}

	tokenStr, err := token.SignedString(pkey)
	if err != nil {
		t.Errorf("failed to sign token: %v", err)
	}

	t.Logf("token: %s", tokenStr)
}
