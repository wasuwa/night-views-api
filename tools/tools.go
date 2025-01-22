//go:build tools
// +build tools

package tools

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../api.yaml

// 空のインポートリストで問題ありません
