package api

import (
    "golang.org/x/crypto/bcrypt"
    "crypto/sha1"
    "fmt"
    "io"
    "os"
)

type Response struct {
    Message   string       `json:"message,omitempty"`
    Type      string       `json:"type,omitempty"`
    Data      interface{}  `json:"data,omitempty"`
}

type Request struct {
    Token   string             `json:"auth,omitempty"`
    Data    map[string]string  `json:"data,omitempty"`
}

type List [] interface{}
type Dict map[interface{}] interface{}

func (self Dict) ToStrMap() map[string]interface{} {
    m := map[string]interface{}{}
    for v, k := range self {
        m[v.(string)] = k
    }
    return m
}

func ToHash(s string) string {
    h := sha1.New()
    io.WriteString(h, s)
    return fmt.Sprintf("%x", h.Sum(nil))
}
func ToPassHash(s string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
    return string(hash), err
}

func CheckPass(p []byte, s string) (error) {
    err := bcrypt.CompareHashAndPassword(p, []byte(s))
    return err
}

func GetEnv(key string, def string) string {
    val, ok := os.LookupEnv(key)
    if !ok {
        return def
    }
    return val
}

func ToLabel(ID uint, Type string) string {
    return fmt.Sprintf("<%s:%d>", Type, ID)
}
