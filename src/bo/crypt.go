package bo

import (
    "crypto/md5"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "hash"
    "log"
    "io/ioutil"
)

func Crypt(text []byte, crypt string) ( []byte, bool) {

    var pem_file_path string
    var block *pem.Block
    var private_key *rsa.PrivateKey
    var public_key *rsa.PublicKey
    var pem_data, label []byte
    var err error

    // A PEM file can contain a Private key among others (Public certificate, Intermidiate Certificate, Root certificate, ...)
    pem_file_path = "/Users/cfrancisco/Documents/workspace/go/certs/self-ssl-no-password.key"
    if pem_data, err = ioutil.ReadFile(pem_file_path); err != nil {
        return []byte("Error reading pem file: %s"), false
    }

    //Package pem implements the PEM data encoding, most commonly used in TLS keys and certificates.
    //Decode will find the next PEM formatted block (certificate, private key etc) in the input.
    //Expected Block type "RSA PRIVATE KEY"
    //http://golang.org/pkg/encoding/pem/
    if block, _ = pem.Decode(pem_data); block == nil || block.Type != "RSA PRIVATE KEY" {
        return []byte("No valid PEM data found"), false
    }

    //x509 parses X.509-encoded keys and certificates.
    //ParsePKCS1PrivateKey returns an RSA private key from its ASN.1 PKCS#1 DER encoded form.
    if private_key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
        return []byte("Private key can't be decoded: %s"), false
    }

    public_key = &private_key.PublicKey

	if crypt == "encrypt" {
    	data, check := Encrypt_oaep(public_key, text, label); if check {
    		return data, true
    	} else {
    		return data, false
    	}
	} else {
		data, check := Decrypt_oaep(private_key, text, label); if check {
			return data, true
		} else {
			return data, false
		}
	}
}

//OAEP Encrypt
func Encrypt_oaep(public_key *rsa.PublicKey, plain_text, label []byte) ( []byte, bool) {
    var encrypted []byte
    var err error
    var md5_hash hash.Hash

    md5_hash = md5.New()

    if encrypted, err = rsa.EncryptOAEP(md5_hash, rand.Reader, public_key, plain_text, label); err != nil {
        log.Fatal(err)
        return encrypted, false
    }
    return encrypted, true
}

func Decrypt_oaep(private_key *rsa.PrivateKey, encrypted, label []byte) ( []byte, bool) {
    var decrypted []byte
    var err error
    var md5_hash hash.Hash

    md5_hash = md5.New()
    if decrypted, err = rsa.DecryptOAEP(md5_hash, rand.Reader, private_key, encrypted, label); err != nil {
        log.Fatal(err)
        return decrypted, false
    }
    return decrypted, true
}
