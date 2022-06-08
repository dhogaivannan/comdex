package bindings

import sdk "github.com/cosmos/cosmos-sdk/types"

type ComdexQuery struct {
	AppData                                *AppData                                `json:"get_app,omitempty"`
	AssetData                              *AssetData                              `json:"get_asset_data,omitempty"`
	State                                  *State                                  `json:"state,omitempty"`
	MintedToken                            *MintedToken                            `json:"total_supply,omitempty"`
	RemoveWhiteListAssetLocker             *RemoveWhiteListAssetLocker             `json:"remove_white_list_asset,omitempty"`
	WhitelistAppIdVaultInterest            *WhitelistAppIdVaultInterest            `json:"whitelist_app_id_vault_interest,omitempty"`
	WhitelistAppIdLockerRewards            *WhitelistAppIdLockerRewards            `json:"whitelist_app_id_locker_rewards,omitempty"`
	ExternalLockerRewards                  *ExternalLockerRewards                  `json:"external_locker_rewards,omitempty"`
	ExternalVaultRewards                   *ExternalVaultRewards                   `json:"external_vault_rewards,omitempty"`
	CollectorLookupTableQuery              *CollectorLookupTableQuery              `json:"collector_lookup_table_query,omitempty"`
	ExtendedPairsVaultRecordsQuery         *ExtendedPairsVaultRecordsQuery         `json:"extended_pairs_vault_records_query,omitempty"`
	AuctionMappingForAppQuery              *AuctionMappingForAppQuery              `json:"auction_mapping_for_app_query,omitempty"`
	WhiteListedAssetQuery                  *WhiteListedAssetQuery                  `json:"white_listed_asset_query,omitempty"`
	UpdateLsrInPairsVaultQuery             *UpdateLsrInPairsVaultQuery             `json:"update_lsr_in_pairs_vault_query,omitempty"`
	UpdateLsrInCollectorLookupTableQuery   *UpdateLsrInCollectorLookupTableQuery   `json:"update_lsr_in_collector_lookup_table_query,omitempty"`
	RemoveWhitelistAssetLockerQuery        *RemoveWhitelistAssetLockerQuery        `json:"remove_whitelist_asset_locker_query,omitempty"`
	RemoveWhitelistAppIdVaultInterestQuery *RemoveWhitelistAppIdVaultInterestQuery `json:"remove_whitelist_app_id_vault_interest_query,omitempty"`
	WhitelistAppIdLiquidationQuery         *WhitelistAppIdLiquidationQuery         `json:"whitelist_app_id_liquidation_query,omitempty"`
	RemoveWhitelistAppIdLiquidationQuery   *RemoveWhitelistAppIdLiquidationQuery   `json:"remove_whitelist_app_id_liquidation_query,omitempty"`
}

type AppData struct {
	App_Id uint64 `json:"app_mapping_id"`
}

type AppDataResponse struct {
	MinGovDeposit    int64  `json:"min_gov_deposit"`
	GovTimeInSeconds int64  `json:"gov_time_in_seconds"`
	AssetId          uint64 `json:"gov_token_id"` //only when isGovToken true
}

type AssetData struct {
	Asset_Id uint64 `json:"asset_id"`
}

type AssetDataResponse struct {
	Denom string `json:"denom"`
}

type State struct {
	Address string `json:"address"`
	Denom   string `json:"denom"`
	Height  string `json:"height"`
	Target  string `json:"target"`
}

type StateResponse struct {
	Amount sdk.Coin `json:"amount"`
}

type MintedToken struct {
	App_Id   uint64 `json:"app_mapping_id"`
	Asset_Id uint64 `json:"asset_id"`
}

type MintedTokenResponse struct {
	MintedTokens int64 `json:"current_supply"`
}

type RemoveWhiteListAssetLocker struct {
	App_Id   uint64   `json:"app_mapping_id"`
	Asset_Id []uint64 `json:"asset_id"`
}

type RemoveWhiteListAssetResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type WhitelistAppIdVaultInterest struct {
	App_Id uint64 `json:"app_mapping_id"`
}

type WhitelistAppIdVaultInterestResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type WhitelistAppIdLockerRewards struct {
	App_Id   uint64   `json:"app_mapping_id"`
	Asset_Id []uint64 `json:"asset_id"`
}

type WhitelistAppIdLockerRewardsResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type ExternalLockerRewards struct {
	App_Id   uint64 `json:"app_mapping_id"`
	Asset_Id uint64 `json:"asset_id"`
}

type ExternalLockerRewardsResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type ExternalVaultRewards struct {
	App_Id   uint64 `json:"app_mapping_id"`
	Asset_Id uint64 `json:"asset_id"`
}

type ExternalVaultRewardsResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type CollectorLookupTableQuery struct {
	AppMappingId     uint64 `json:"app_mapping_id"`
	CollectorAssetId uint64 `json:"collector_asset_id"`
	SecondaryAssetId uint64 `json:"secondary_asset_id"`
}

type CollectorLookupTableQueryResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type ExtendedPairsVaultRecordsQuery struct {
	AppMappingId uint64  `json:"app_mapping_id"`
	PairId       uint64  `json:"pair_id"`
	StabilityFee sdk.Dec `json:"stability_fee"`
	ClosingFee   sdk.Dec `json:"closing_fee"`
	DrawDownFee  sdk.Dec `json:"draw_down_fee"`
	DebtCeiling  uint64  `json:"debt_ceiling"`
	DebtFloor    uint64  `json:"debt_floor"`
	PairName     string  `json:"pair_name"`
}

type ExtendedPairsVaultRecordsQueryResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type AuctionMappingForAppQuery struct {
	AppMappingId uint64 `json:"app_mapping_id"`
}

type AuctionMappingForAppQueryResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type WhiteListedAssetQuery struct {
	AppMappingId uint64 `json:"app_mapping_id"`
	AssetId      uint64 `json:"asset_id"`
}

type WhiteListedAssetQueryResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type UpdateLsrInPairsVaultQuery struct {
	AppMappingId uint64 `json:"app_mapping_id"`
	ExtPairId    uint64 `json:"ext_pair_id"`
}

type UpdateLsrInPairsVaultQueryResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type UpdateLsrInCollectorLookupTableQuery struct {
	AppMappingId uint64 `json:"app_mapping_id"`
	AssetId      uint64 `json:"asset_id"`
}

type UpdateLsrInCollectorLookupTableQueryResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type RemoveWhitelistAssetLockerQuery struct {
	AppMappingId uint64 `json:"app_mapping_id"`
	AssetId      uint64 `json:"asset_id"`
}

type RemoveWhitelistAssetLockerQueryResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type RemoveWhitelistAppIdVaultInterestQuery struct {
	AppMappingId uint64 `json:"app_mapping_id"`
}

type RemoveWhitelistAppIdVaultInterestQueryResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type WhitelistAppIdLiquidationQuery struct {
	AppMappingId uint64 `json:"app_mapping_id"`
}

type WhitelistAppIdLiquidationQueryResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}

type RemoveWhitelistAppIdLiquidationQuery struct {
	AppMappingId uint64 `json:"app_mapping_id"`
}

type RemoveWhitelistAppIdLiquidationQueryResponse struct {
	Found bool   `json:"found"`
	Err   string `json:"err"`
}
