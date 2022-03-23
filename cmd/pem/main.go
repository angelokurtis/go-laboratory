package main

import (
	"crypto/x509"
	"encoding/pem"
)

func main() {
	const certPEM = `
-----BEGIN CERTIFICATE-----
MIIC5zCCAc+gAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTIxMTAxNDE0NDQyOVoXDTMxMTAxMjE0NDQyOVowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAPMU
KQ1ddesokO6sBdqONBrR5NS8lCuhKVGyr8Equ4tXKMwEcwPEebEKKeVx+icfaCMV
UKExxOWj66Gony0kLtPwhqNTRseNIggu0ebxYWIbcCgZM8N6Lk97T3r1Wl0//+6m
oQmAkkJAgdz3/N+LJKMpCWhFzWajwoMV/5pMtq8N+eguQcU8LC8JmAgz91GWYSwK
a6ha422wGIL4ALcnVTGzBQS7/JOi05+cfJD1UjfR18h0gDvFnx9O8LqZqLaVNhTl
KBXi8GrDYsM6jbEeGtPcgK8aamafATt5ItbvCub1+nxTwLSRbQv8yMTY79U1hqpT
CM7XTPawjjXIvAR3ljkCAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wHQYDVR0OBBYEFH4IMaaKY72EjFdVJ73QWLjFsz5aMA0GCSqGSIb3
DQEBCwUAA4IBAQBwiugQK5jp9CT+STNDxFkpcW/NxcEmJGWNBN6PvAUqXxpd2xH9
m9jsk9PL6L9ZNeGOnv6HT1n7KdpNXDvHe0QLjFEUnEvxPU2SeRlaG0Xhg2xqcqfq
EkMcRwxrm4IjybG98FpMF7ax/z0Bo2Tu+OIWEqImM2bt+/sN8J3W0k7pl7RpE2kv
xh6lxjiqPczX0ef2/AHUjfUMeBei6cIuo5ZqoKZ223FxmJv1YzYJ/rCJKzNGj2/V
bEJKcNvo2DmUrzYZee4kCBM9iyVrIfGQG/Es1hLIUNQKfTk3vcg4tXoOMVOBvBhH
TJT6xOrljrj54FvVFHeI5BJtumFIkhkHf8nQ
-----END CERTIFICATE-----`
	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		panic("failed to parse cert")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse cert: " + err.Error())
	}
	if _, err := cert.Verify(x509.VerifyOptions{}); err != nil {
		panic("failed to verify cert: " + err.Error())
	}
}
