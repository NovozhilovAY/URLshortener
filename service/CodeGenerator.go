package service

import (
	"math/rand"
	"strings"
	"time"
)

const CODE_LENGTH = 32
const CODE_ALPHABET = "abcdefghijklmnopqrstuvwxyz0123456789"

type CodeGenerator struct {
	codeLength int
	symbols    string
}

func NewCodeGenerator() *CodeGenerator {
	rand.Seed(time.Now().UnixNano())
	return &CodeGenerator{codeLength: CODE_LENGTH, symbols: CODE_ALPHABET}
}

func (CodeGenerator) GenerateCode() string {
	var sb strings.Builder
	for i := 0; i < CODE_LENGTH; i++ {
		sb.WriteString(string(CODE_ALPHABET[rand.Intn(len(CODE_ALPHABET))]))
	}
	return sb.String()
}
