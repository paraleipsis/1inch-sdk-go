package validate

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/paraleipsis/1inch-sdk-go/constants"
	"github.com/paraleipsis/1inch-sdk-go/internal/bigint"
	"github.com/paraleipsis/1inch-sdk-go/internal/slice_utils"
)

func CheckEthereumAddressRequired(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "EthereumAddress")
	}

	if value == "" {
		return NewParameterMissingError(variableName)
	}

	return CheckEthereumAddress(value, variableName)
}

func CheckEthereumAddress(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "EthereumAddress")
	}
	if value == "" {
		return nil
	}

	re := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	if !re.MatchString(value) {
		return NewParameterValidationError(variableName, "not a valid Ethereum address")
	}
	return nil
}

func CheckEthereumAddressListRequired(parameter interface{}, variableName string) error {
	addresses, ok := parameter.([]string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a list of strings", variableName, "EthereumAddressList")
	}
	if len(addresses) == 0 {
		return NewParameterMissingError(variableName)
	}

	for _, address := range addresses {
		if address == "" {
			return NewParameterMissingError(variableName)
		}
		re := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
		if !re.MatchString(address) {
			return NewParameterValidationError(variableName, "not a valid Ethereum address")
		}
	}

	return nil
}

var bigIntMax, _ = bigint.FromString("115792089237316195423570985008687907853269984665640564039457584007913129639935")
var bigIntZero = big.NewInt(0)

func CheckBigIntRequired(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "BigInt")
	}

	if value == "" {
		return NewParameterMissingError(variableName)
	}

	return CheckBigInt(value, variableName)
}
func CheckBigInt(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "BigInt")
	}

	if value == "" {
		return nil
	}

	parsedValue, err := bigint.FromString(value)
	if err != nil {
		return NewParameterValidationError(variableName, "not a valid value")
	}
	if parsedValue.Cmp(bigIntMax) > 0 {
		return NewParameterValidationError(variableName, "too big to fit in uint256")
	}
	if parsedValue.Cmp(bigIntZero) < 0 {
		return NewParameterValidationError(variableName, "must be a positive value")
	}
	return nil
}

func CheckChainIdIntRequired(parameter interface{}, variableName string) error {
	value, ok := parameter.(int)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be an int", variableName, "ChainIdInt")
	}

	if value == 0 {
		return NewParameterMissingError(variableName)
	}

	return CheckChainIdInt(value, variableName)
}

func CheckChainIdInt(parameter interface{}, variableName string) error {
	value, ok := parameter.(int)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be an int", variableName, "ChainIdInt")
	}
	if value == 0 {
		return nil
	}

	if !slice_utils.Contains(value, constants.ValidChainIds) {
		return NewParameterValidationError(variableName, fmt.Sprintf("is invalid, valid chain ids are: %v", constants.ValidChainIds))
	}
	return nil
}

func CheckChainIdFloat32Required(parameter interface{}, variableName string) error {
	value, ok := parameter.(float32)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a float32", variableName, "ChainIdFloat32")
	}

	if value == 0 {
		return NewParameterMissingError(variableName)
	}

	return CheckChainIdFloat32(value, variableName)
}

func CheckChainIdFloat32(parameter interface{}, variableName string) error {
	value, ok := parameter.(float32)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a float32", variableName, "ChainIdFloat32")
	}
	if value == 0 {
		return nil
	}

	if !slice_utils.Contains(int(value), constants.ValidChainIds) {
		return NewParameterValidationError(variableName, fmt.Sprintf("is invalid, valid chain ids are: %v", constants.ValidChainIds))
	}
	return nil
}

func CheckPrivateKeyRequired(parameter interface{}, variableName string) error {
	address, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "PrivateKey")
	}

	if address == "" {
		return NewParameterMissingError(variableName)
	}

	return CheckPrivateKey(address, variableName)
}

func CheckPrivateKey(parameter interface{}, variableName string) error {
	address, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "PrivateKey")
	}

	if address == "" {
		return nil
	}

	re := regexp.MustCompile(`^[a-fA-F0-9]{64}$`)
	if !re.MatchString(address) {
		return NewParameterValidationError(variableName, "not a valid private key")
	}
	return nil
}

func CheckApprovalType(parameter interface{}, variableName string) error {
	value, ok := parameter.(int)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be an int", variableName, "ApprovalType")
	}

	if value == 0 {
		return nil
	}

	if !slice_utils.Contains(value, []int{0, 1, 2}) {
		return NewParameterValidationError(variableName, "invalid approval type")
	}
	return nil
}

func CheckSlippageRequired(parameter interface{}, variableName string) error {
	value, ok := parameter.(float32)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a float32", variableName, "Slippage")
	}
	if value == 0 {
		return NewParameterMissingError(variableName)
	}
	return CheckSlippage(value, variableName)
}

func CheckSlippage(parameter interface{}, variableName string) error {
	value, ok := parameter.(float32)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a float32", variableName, "Slippage")
	}
	if value == 0 {
		return nil
	}
	if value < 0.01 || value > 50 {
		return NewParameterValidationError(variableName, fmt.Sprintf("invalid slippage value (%v) - only values 0.01-50 are allowed", value))
	}
	return nil
}

func CheckPage(parameter interface{}, variableName string) error {
	value, ok := parameter.(float32)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a float32", variableName, "Page")
	}
	if value == 0 {
		return nil
	}

	if value < 1 {
		return NewParameterValidationError(variableName, "must be greater than 0")
	}
	return nil
}

func CheckLimit(parameter interface{}, variableName string) error {
	value, ok := parameter.(float32)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a float32", variableName, "Limit")
	}
	if value == 0 {
		return nil
	}

	if value < 1 {
		return NewParameterValidationError(variableName, "must be greater than 0")
	}
	return nil
}

func CheckStatusesInts(parameter interface{}, variableName string) error {
	value, ok := parameter.([]float32)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a []float32", variableName, "StatusesInts")
	}

	if value == nil {
		return nil
	}

	if HasDuplicates(value) {
		return NewParameterValidationError(variableName, "must not contain duplicates")
	}
	validStatuses := []float32{1, 2, 3}
	if !IsSubset(value, validStatuses) {
		return NewParameterValidationError(variableName, fmt.Sprintf("can only contain %v", validStatuses))
	}
	return nil
}

func CheckStatusesStrings(parameter interface{}, variableName string) error {
	value, ok := parameter.([]string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a []string", variableName, "StatusesStrings")
	}

	if HasDuplicates(value) {
		return NewParameterValidationError(variableName, "must not contain duplicates")
	}
	validStatuses := []string{"1", "2", "3"}
	if !IsSubset(value, validStatuses) {
		return NewParameterValidationError(variableName, fmt.Sprintf("can only contain %v", validStatuses))
	}
	return nil
}

func CheckSortBy(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "SortBy")
	}

	if value == "" {
		return nil
	}

	validSortBy := []string{"createDateTime", "takerRate", "makerRate", "makerAmount", "takerAmount"}
	if !slice_utils.Contains(value, validSortBy) {
		return NewParameterValidationError(variableName, fmt.Sprintf("can only contain %v", validSortBy))
	}
	return nil
}

func CheckOrderHashRequired(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "OrderHash")
	}

	if value == "" {
		return NewParameterMissingError(variableName)
	}
	return CheckOrderHash(value, variableName)
}

func CheckOrderHash(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "OrderHash")
	}

	if value == "" {
		return nil
	}
	// TODO add criteria that captures valid order hash strings here
	return nil
}

func CheckProtocols(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "Protocols")
	}

	if value == "" {
		return nil
	}

	pattern := `^[a-zA-Z0-9_]+(,[a-zA-Z0-9_]+)*$`
	re := regexp.MustCompile(pattern)

	ok = re.MatchString(value)
	if !ok {
		return NewParameterValidationError(variableName, "must be formatted as a single-string list exactly in the format 'Protocol1,Protocol2,Protocol3' without any "+
			"spaces between each protocol name. These names must match the exact protocol id used by the 1inch APIs "+
			"(use the Swap service's GetLiquiditySources function to see this list). Additionally, there cannot be a trailing comma at the end of the list.")
	}

	if ok {
		addresses := strings.Split(value, ",")
		addressesMap := make(map[string]bool)
		for _, address := range addresses {
			if _, exists := addressesMap[address]; exists {
				return NewParameterValidationError(variableName, "Duplicate protocol found in list")
			}
			addressesMap[address] = true
		}
	}

	return nil
}

func CheckFee(parameter interface{}, variableName string) error {
	value, ok := parameter.(float32)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "Fee")
	}

	if value < 0 {
		return NewParameterValidationError(variableName, "must be a positive value")
	}

	if value > 3 {
		return NewParameterValidationError(variableName, "must be a value between 0 and 3")
	}

	return nil
}

//  TODO The enforced naming pattern for the variable name string literal doesn't work for generic types like "Float32NonNegativeWhole"

func CheckFloat32NonNegativeWhole(parameter interface{}, variableName string) error {
	value, ok := parameter.(float32)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "Float32NonNegativeWhole")
	}

	if value < 0 {
		return NewParameterValidationError(variableName, "must be 0 or greater")
	}

	// Cast it to an int to truncate it, then cast back to float32 for the comparison
	if float32(int(value)) != value {
		return NewParameterValidationError(variableName, "must be an integer")
	}

	return nil
}

func CheckConnectorTokens(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "ConnectorTokens")
	}

	if value == "" {
		return nil
	}

	pattern := `^0x[a-fA-F0-9]{40}(,0x[a-fA-F0-9]{40})*$`
	re := regexp.MustCompile(pattern)

	ok = re.MatchString(value)
	if !ok {
		return NewParameterValidationError(variableName, "must be formatted as a single-string list exactly in the format '0x123,0x456,0x789' "+
			"without any spaces between each protocol name. Additionally, there cannot be a trailing comma at the end of the list.")
	}

	if ok {
		// Split the string by commas to get individual addresses
		addresses := strings.Split(value, ",")

		// Use a map to check for duplicates
		addressesMap := make(map[string]bool)

		for _, address := range addresses {
			if _, exists := addressesMap[address]; exists {
				return NewParameterValidationError(variableName, "Duplicate address found in list")
			}
			addressesMap[address] = true
		}
	}

	return nil
}

func CheckPermitHash(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "PermitHash")
	}
	if value == "" {
		return nil
	}

	re := regexp.MustCompile(`^0x[a-fA-F0-9]*$`)
	if !re.MatchString(value) {
		return NewParameterValidationError(variableName, "not a valid permit hash")
	}
	return nil
}

func CheckExpireAfterUnix(parameter interface{}, variableName string) error {
	value, ok := parameter.(int64)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "ExpireAfterUnix")
	}
	if value == 0 {
		return nil
	}

	if value < time.Now().Unix() {
		return NewParameterValidationError(variableName, "must be a future timestamp")
	}
	return nil
}

func CheckFiatCurrency(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "FiatCurrency")
	}

	if len(value) != 3 {
		return NewParameterValidationError(variableName, "must have len = 3 (like USD, EUR, etc)")
	}

	return nil
}

func CheckTimerange(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "Timerange")
	}

	validTimerangeValues := []string{"1day", "1week", "1month", "1year", "3years"}
	if !slice_utils.Contains(value, validTimerangeValues) {
		return NewParameterValidationError(variableName, fmt.Sprintf("is invalid, valid chain ids are: %v", validTimerangeValues))
	}
	return nil
}

func CheckJsonRpcVersionRequired(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "JsonRpcVersion")
	}

	if value == "" {
		return NewParameterMissingError(variableName)
	}

	return CheckJsonRpcVersion(value, variableName)
}

func CheckJsonRpcVersion(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "JsonRpcVersion")
	}

	validJsonRpcValues := []string{"1.0", "2.0"}
	if !slice_utils.Contains(value, validJsonRpcValues) {
		return NewParameterValidationError(variableName, fmt.Sprintf("is invalid, valid rpc version are: %v", validJsonRpcValues))
	}
	return nil
}

func CheckNodeType(parameter interface{}, variableName string) error {
	value, ok := parameter.(string)
	if !ok {
		return fmt.Errorf("for parameter '%v' to be validated as '%v', it must be a string", variableName, "NodeType")
	}

	validJsonRpcValues := []string{"1.0", "2.0"}
	if !slice_utils.Contains(value, validJsonRpcValues) {
		return NewParameterValidationError(variableName, fmt.Sprintf("is invalid, valid rpc version are: %v", validJsonRpcValues))
	}
	return nil
}
