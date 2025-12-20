package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestConvertTransferLog(t *testing.T) {
	type args struct {
		log string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.Log
		wantErr bool
	}{
		{
			name: "TestConvertTransferLog",
			args: args{
				log: "{\"address\":\"0x1c7d4b196cb0c7b01d743fbc6116a902379c7238\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x0000000000000000000000003c3380cdfb94dfeeaa41cad9f58254ae380d752d\",\"0x000000000000000000000000157b434ed20aea9d647b5a7c308f0eb0b626f32e\"],\"data\":\"0x00000000000000000000000000000000000000000000000000000000000f4240\",\"blockNumber\":\"0x960611\",\"transactionHash\":\"0x208123e31df6bf73032f03d2176320a6601f9f9bf58645753a86f531b4e81c50\",\"transactionIndex\":\"0x4\",\"blockHash\":\"0x3e2e72f1c56476833b85dbe399f1eadb6c0021c752d1d298d288e81a5a2d8626\",\"blockTimestamp\":\"0x693d71d8\",\"logIndex\":\"0x3\",\"removed\":false}",
			},
			want: &types.Log{
				Address: common.HexToAddress("0x1c7d4b196cb0c7b01d743fbc6116a902379c7238"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertTransferLog(tt.args.log)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertTransferLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Address.Hex() != tt.want.Address.Hex() {
				t.Errorf("ConvertTransferLog() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseTransferLog(t *testing.T) {
	logInfo := "{\"address\":\"0x1c7d4b196cb0c7b01d743fbc6116a902379c7238\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x0000000000000000000000003c3380cdfb94dfeeaa41cad9f58254ae380d752d\",\"0x000000000000000000000000157b434ed20aea9d647b5a7c308f0eb0b626f32e\"],\"data\":\"0x00000000000000000000000000000000000000000000000000000000000f4240\",\"blockNumber\":\"0x960611\",\"transactionHash\":\"0x208123e31df6bf73032f03d2176320a6601f9f9bf58645753a86f531b4e81c50\",\"transactionIndex\":\"0x4\",\"blockHash\":\"0x3e2e72f1c56476833b85dbe399f1eadb6c0021c752d1d298d288e81a5a2d8626\",\"blockTimestamp\":\"0x693d71d8\",\"logIndex\":\"0x3\",\"removed\":false}"
	got, err := ConvertTransferLog(logInfo)
	if err != nil {
		t.Errorf("ParseTransferLog() error = %v", err)
		return
	}
	ParseTransferLog(got)
	t.Log("ok")
}
