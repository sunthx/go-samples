package main

import (
	"math/big"
	"crypto/rand"
	"crypto/x509/pkix"
	"crypto/x509"
	"time"
	"net"
	"crypto/rsa"
	"os"
	"encoding/pem"
)

func main()  {
	//创建128位的哈希值
	max := new (big.Int).Lsh(big.NewInt(1), 128)

	//证书序列号
	serialNumber,_ := rand.Int(rand.Reader, max)

	//证书的标题
	subject := pkix.Name{
		Organization: []string{"xxxxxxxxx Co."},
		OrganizationalUnit:[]string{"Games"},
		CommonName:"Create CA",
	}

	//
	template:=x509.Certificate{
		SerialNumber:serialNumber,
		Subject:subject,

		//证书的开始时间与过期时间
		NotBefore:time.Now(),
		NotAfter:time.Now().Add(365 * 24 * time.Hour),

		//证书的用途
		KeyUsage:x509.KeyUsageDataEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

		//证书限制
		IPAddresses:[]net.IP{net.ParseIP("127.0.0.2")},
	}

	pk,_ := rsa.GenerateKey(rand.Reader, 2048)

	derBytes,_ := x509.CreateCertificate(rand.Reader,&template,&template,&pk.PublicKey,pk)
	certOut,_ := os.Create("cert.pem")
	pem.Encode(certOut,&pem.Block{Type:"CERTIFICATE",Bytes:derBytes})
	certOut.Close()

	keyOut,_ := os.Create("key.pem")
	pem.Encode(keyOut,&pem.Block{Type:"RSA PRIVATE KEY",Bytes:x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
