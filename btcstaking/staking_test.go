package btcstaking_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/babylonchain/babylon/btcstaking"
	"github.com/babylonchain/babylon/testutil/datagen"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/mempool"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/stretchr/testify/require"
)

// StakingScriptData is a struct that holds data parsed from staking script
type StakingScriptData struct {
	StakerKey           *btcec.PublicKey
	FinalityProviderKey *btcec.PublicKey
	CovenantKey         *btcec.PublicKey
	StakingTime         uint16
}


func genRandomBTCAddress(r *rand.Rand) (*btcutil.AddressPubKeyHash, error) {
	return btcutil.NewAddressPubKeyHash(datagen.GenRandomByteArray(r, 20), &chaincfg.MainNetParams)
}

func taprootOutputWithValue(t *testing.T, r *rand.Rand, value btcutil.Amount) *wire.TxOut {
	bytes := datagen.GenRandomByteArray(r, 32)
	addrr, err := btcutil.NewAddressTaproot(bytes, &chaincfg.MainNetParams)
	require.NoError(t, err)
	return outputFromAddressAndValue(t, addrr, value)
}

func outputFromAddressAndValue(t *testing.T, addr btcutil.Address, value btcutil.Amount) *wire.TxOut {
	pkScript, err := txscript.PayToAddrScript(addr)
	require.NoError(t, err)
	return wire.NewTxOut(int64(value), pkScript)
}

func testSlashingTx(
	r *rand.Rand,
	t *testing.T,
	stakingTx *wire.MsgTx,
	stakingOutputIdx int,
	slashingRate sdkmath.LegacyDec,
	fee int64,
	stakerPk *btcec.PublicKey,
	slashingChangeLockTime uint16,
) {
	// Generate random slashing and change addresses
	slashingAddress, err := genRandomBTCAddress(r)
	require.NoError(t, err)

	// Construct slashing transaction using the provided parameters
	slashingTx, err := btcstaking.BuildSlashingTxFromStakingTxStrict(
		stakingTx,
		uint32(stakingOutputIdx),
		slashingAddress,
		stakerPk,
		slashingChangeLockTime,
		fee,
		slashingRate,
		&chaincfg.MainNetParams,
	)

	if btcstaking.IsRateValid(slashingRate) {
		// If the slashing rate is valid i.e., in the range (0,1) with at most 2 decimal places,
		// it is still possible that the slashing transaction is invalid. The following checks will confirm that
		// slashing tx is not constructed if
		// - the change output has insufficient funds.
		// - the change output is less than the dust threshold.
		// - The slashing output is less than the dust threshold.

		slashingRateFloat64, err2 := slashingRate.Float64()
		require.NoError(t, err2)

		stakingAmount := btcutil.Amount(stakingTx.TxOut[stakingOutputIdx].Value)
		slashingAmount := stakingAmount.MulF64(slashingRateFloat64)
		changeAmount := stakingAmount - slashingAmount - btcutil.Amount(fee)

		// check if the created outputs are not dust
		slashingOutput := outputFromAddressAndValue(t, slashingAddress, slashingAmount)
		changeOutput := taprootOutputWithValue(t, r, changeAmount)

		if changeAmount <= 0 {
			require.Error(t, err)
			require.ErrorIs(t, err, btcstaking.ErrInsufficientChangeAmount)
		} else if mempool.IsDust(slashingOutput, mempool.DefaultMinRelayTxFee) || mempool.IsDust(changeOutput, mempool.DefaultMinRelayTxFee) {
			require.Error(t, err)
			require.ErrorIs(t, err, btcstaking.ErrDustOutputFound)
		} else {
			require.NoError(t, err)
			err = btcstaking.CheckTransactions(
				slashingTx,
				stakingTx,
				uint32(stakingOutputIdx),
				fee,
				slashingRate,
				slashingAddress,
				stakerPk,
				slashingChangeLockTime,
				&chaincfg.MainNetParams,
			)
			require.NoError(t, err)
		}
	} else {
		require.Error(t, err)
		require.ErrorIs(t, err, btcstaking.ErrInvalidSlashingRate)
	}
}

func FuzzGeneratingSignatureValidation(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		pk, err := btcec.NewPrivateKey()
		require.NoError(t, err)
		inputHash, err := chainhash.NewHash(datagen.GenRandomByteArray(r, 32))
		require.NoError(t, err)

		tx := wire.NewMsgTx(2)
		foundingOutput := wire.NewTxOut(int64(r.Intn(1000)), datagen.GenRandomByteArray(r, 32))
		tx.AddTxIn(
			wire.NewTxIn(wire.NewOutPoint(inputHash, uint32(r.Intn(20))), nil, nil),
		)
		tx.AddTxOut(
			wire.NewTxOut(int64(r.Intn(1000)), datagen.GenRandomByteArray(r, 32)),
		)
		script := datagen.GenRandomByteArray(r, 150)

		sig, err := btcstaking.SignTxWithOneScriptSpendInputFromScript(
			tx,
			foundingOutput,
			pk,
			script,
		)

		require.NoError(t, err)

		err = btcstaking.VerifyTransactionSigWithOutput(
			tx,
			foundingOutput,
			script,
			pk.PubKey(),
			sig.Serialize(),
		)

		require.NoError(t, err)
	})
}
