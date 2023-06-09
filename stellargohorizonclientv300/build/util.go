package build

import (
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/xdr"
)

func setAccountId(addressOrSeed string, aid *xdr.AccountId) error {
	kp, err := keypair.Parse(addressOrSeed)
	if err != nil {
		return err
	}

	if aid == nil {
		return errors.New("aid is nil in setAccountId")
	}

	return aid.SetAddress(kp.Address())
}

func setMuxedAccount(addressOrSeed string, m *xdr.MuxedAccount) error {
	kp, err := keypair.Parse(addressOrSeed)
	if err != nil {
		return err
	}

	if m == nil {
		return errors.New("m is nil in setMuxedAccount")
	}

	return m.SetAddress(kp.Address())
}

func createAlphaNumAsset(code, issuerAccountId string) (xdr.Asset, error) {
	var issuer xdr.AccountId
	err := setAccountId(issuerAccountId, &issuer)
	if err != nil {
		return xdr.Asset{}, err
	}

	length := len(code)
	switch {
	case length >= 1 && length <= 4:
		var codeArray xdr.AssetCode4
		byteArray := []byte(code)
		copy(codeArray[:], byteArray[0:length])
		asset := xdr.AlphaNum4{AssetCode: codeArray, Issuer: issuer}
		return xdr.NewAsset(xdr.AssetTypeAssetTypeCreditAlphanum4, asset)
	case length >= 5 && length <= 12:
		var codeArray xdr.AssetCode12
		byteArray := []byte(code)
		copy(codeArray[:], byteArray[0:length])
		asset := xdr.AlphaNum12{AssetCode: codeArray, Issuer: issuer}
		return xdr.NewAsset(xdr.AssetTypeAssetTypeCreditAlphanum12, asset)
	default:
		return xdr.Asset{}, errors.New("Asset code length is invalid")
	}
}
