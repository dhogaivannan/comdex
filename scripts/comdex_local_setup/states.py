from constants import *

APPS = [
    # [name, shortName, minGovDeposit, govTimeInSeconds]
    ["harbor", "hbr", 1000000, 5],  # ID - 1
    ["cswap", "cswap", 0, 0],  # ID - 2
    ["commodo", "comdo", 0, 0],  # ID - 3
]

ASSETS = [
    # [name, denom, isOnChain, assetOraclePriceRequired]
    ["ATOM", "uatom", 0, 1],  # ID - 1
    ["CMDX", "ucmdx", 0, 1],  # ID - 2
    ["CMST", "ucmst", 0, 0],  # ID - 3
    ["OSMO", "uosmo", 0, 1],  # ID - 4
    ["CATOM", "ucatom", 0, 0],  # ID - 5
    ["CCMDX", "uccmdx", 0, 0],  # ID - 6
    ["CCMST", "uccmst", 0, 0],  # ID - 7
    ["COSMO", "ucosmo", 0, 0],  # ID - 8
    ["HARBOR", "uharbor", 1, 0],  # ID - 9
]

PAIRS = [
    # [assetID1, assetID2]
    [1, 3],  # ID - 1
    [2, 3],  # ID - 2
    [4, 3],  # ID - 3
]

LIQUIDITY_PAIRS = [
    # [appID, baseCoinDenom, quoteCoinDenom]
    [2, ASSETS[1][1], ASSETS[8][1]],  # ID - 1
    [2, ASSETS[1][1], ASSETS[2][1]],  # ID - 2
]

LIQUIDITY_POOLS = [
    # [appID, pairID, depositCoins]
    [2, 1, f"1000000000000{ASSETS[1][1]},2000000000000{ASSETS[8][1]}"],  # ID - 1
]

ADD_ASSET_RATES = [
    # [ assetName, jsonObject]
    [
        "CMST",
        {
            "asset_id": "3",
            "u_optimal": "0.8",
            "base": "0.002",
            "slope_1": "0.06",
            "slope_2": "0.6",
            "enable_stable_borrow": "1",
            "stable_base": "0.04",
            "stable_slope_1": "0.04",
            "stable_slope_2": "0.6",
            "ltv": "0.8",
            "liquidation_threshold": "0.85",
            "liquidation_penalty": "0.025",
            "liquidation_bonus": "0.025",
            "reserve_factor": "0.1",
            "c_asset_id": "7",
            "title": "Add Asset Rates Stats CMST",
            "description": "adding asset rates stats for CMST",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "ATOM",
        {
            "asset_id": "1",
            "u_optimal": "0.75",
            "base": "0.002",
            "slope_1": "0.07",
            "slope_2": "1.25",
            "enable_stable_borrow": "0",
            "stable_base": "0.0",
            "stable_slope_1": "0.0",
            "stable_slope_2": "0.0",
            "ltv": "0.7",
            "liquidation_threshold": "0.75",
            "liquidation_penalty": "0.05",
            "liquidation_bonus": "0.05",
            "reserve_factor": "0.2",
            "c_asset_id": "5",
            "title": "Add Asset Rates Stats ATOM",
            "description": "adding asset rates stats ATOM",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "OSMO",
        {
            "asset_id": "4",
            "u_optimal": "0.65",
            "base": "0.002",
            "slope_1": "0.08",
            "slope_2": "1.5",
            "enable_stable_borrow": "0",
            "stable_base": "0.0",
            "stable_slope_1": "0.0",
            "stable_slope_2": "0.0",
            "ltv": "0.6",
            "liquidation_threshold": "0.65",
            "liquidation_penalty": "0.05",
            "liquidation_bonus": "0.05",
            "reserve_factor": "0.2",
            "c_asset_id": "8",
            "title": "Add Asset Rates Stats OSMO",
            "description": "adding asset rates stats OSMO",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "CMDX",
        {
            "asset_id": "2",
            "u_optimal": "0.5",
            "base": "0.002",
            "slope_1": "0.08",
            "slope_2": "2.0",
            "enable_stable_borrow": "0",
            "stable_base": "0.0",
            "stable_slope_1": "0.0",
            "stable_slope_2": "0.0",
            "ltv": "0.5",
            "liquidation_threshold": "0.55",
            "liquidation_penalty": "0.05",
            "liquidation_bonus": "0.05",
            "reserve_factor": "0.2",
            "c_asset_id": "6",
            "title": "Add Asset Rates Stats CMDX",
            "description": "adding asset rates stats CMDX",
            "deposit": "10000000ucmdx",
        },
    ],
]

ADD_LEND_POOL = [
    {
        "module_name": "cmdx",
        "main_asset_id": "2",
        "first_bridged_asset_id": "3",
        "second_bridged_asset_id": "1",
        "asset_id": "1,2,3",
        "is_bridged_asset": "1,0,1",
        "c_pool_name": "CMDX-ATOM-CMST",
        "reserve_funds": "100000000",
        "title": "Add pool",
        "description": "adding pool",
        "deposit": "10000000ucmdx",
    },
    {
        "module_name": "osmo",
        "main_asset_id": "4",
        "first_bridged_asset_id": "3",
        "second_bridged_asset_id": "1",
        "asset_id": "1,4,3",
        "is_bridged_asset": "1,0,1",
        "c_pool_name": "OSMO-ATOM-CMST",
        "reserve_funds": "100000000",
        "title": "Add pool",
        "description": "adding pool",
        "deposit": "10000000ucmdx",
    },
]

ADD_LEND_PAIR = [
    [
        "CMDX-CMST",
        {
            "asset_in": "2",
            "asset_out": "3",
            "is_inter_pool": "0",
            "asset_out_pool_id": "1",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair CMDX-CMST",
            "description": "adding extended pairs for CMDX-CMST same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "CMDX-ATOM",
        {
            "asset_in": "2",
            "asset_out": "1",
            "is_inter_pool": "0",
            "asset_out_pool_id": "1",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair CMDX-ATOM",
            "description": "adding extended pairs CMDX-ATOM same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "ATOM-CMDX",
        {
            "asset_in": "1",
            "asset_out": "2",
            "is_inter_pool": "0",
            "asset_out_pool_id": "1",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair ATOM-CMDX",
            "description": "adding extended pairs ATOM-CMDX same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "ATOM-CMST",
        {
            "asset_in": "1",
            "asset_out": "3",
            "is_inter_pool": "0",
            "asset_out_pool_id": "1",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair ATOM-CMST",
            "description": "adding extended pairs ATOM-CMST same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "CMST-CMDX",
        {
            "asset_in": "3",
            "asset_out": "2",
            "is_inter_pool": "0",
            "asset_out_pool_id": "1",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair CMST-CMDX",
            "description": "adding extended pairs CMST-CMDX same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "CMST-ATOM",
        {
            "asset_in": "3",
            "asset_out": "1",
            "is_inter_pool": "0",
            "asset_out_pool_id": "1",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair CMST-ATOM",
            "description": "adding extended pairs CMST-ATOM same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "OSMO-CMST",
        {
            "asset_in": "4",
            "asset_out": "3",
            "is_inter_pool": "0",
            "asset_out_pool_id": "2",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair OSMO-CMST",
            "description": "adding extended pairs OSMO-CMST same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "OSMO-ATOM",
        {
            "asset_in": "4",
            "asset_out": "1",
            "is_inter_pool": "0",
            "asset_out_pool_id": "2",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair OSMO-ATOM",
            "description": "adding extended pairs OSMO-ATOM same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "ATOM-OSMO",
        {
            "asset_in": "1",
            "asset_out": "4",
            "is_inter_pool": "0",
            "asset_out_pool_id": "2",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair ATOM-OSMO",
            "description": "adding extended pairs ATOM-OSMO same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "ATOM-CMST",
        {
            "asset_in": "1",
            "asset_out": "3",
            "is_inter_pool": "0",
            "asset_out_pool_id": "2",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair ATOM-CMST",
            "description": "adding extended pairs ATOM-CMST same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "CMST-OSMO",
        {
            "asset_in": "3",
            "asset_out": "4",
            "is_inter_pool": "0",
            "asset_out_pool_id": "2",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair CMST-OSMO",
            "description": "adding extended pairs CMST-OSMO same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "CMST-ATOM",
        {
            "asset_in": "3",
            "asset_out": "1",
            "is_inter_pool": "0",
            "asset_out_pool_id": "2",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair CMST-ATOM",
            "description": "adding extended pairs CMST-ATOM same pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "CMDX-OSMO",
        {
            "asset_in": "2",
            "asset_out": "4",
            "is_inter_pool": "1",
            "asset_out_pool_id": "2",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair CMDX-OSMO",
            "description": "adding extended pairs CMDX-OSMO cross pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "CMST-OSMO",
        {
            "asset_in": "3",
            "asset_out": "4",
            "is_inter_pool": "1",
            "asset_out_pool_id": "2",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair CMST-OSMO",
            "description": "adding extended pairs CMST-OSMO cross pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "ATOM-OSMO",
        {
            "asset_in": "1",
            "asset_out": "4",
            "is_inter_pool": "1",
            "asset_out_pool_id": "2",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair ATOM-OSMO",
            "description": "adding extended pairs ATOM-OSMO cross pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "OSMO-CMDX",
        {
            "asset_in": "4",
            "asset_out": "2",
            "is_inter_pool": "1",
            "asset_out_pool_id": "1",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair OSMO-CMDX",
            "description": "adding extended pairs OSMO-CMDX cross pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "CMST-CMDX",
        {
            "asset_in": "3",
            "asset_out": "2",
            "is_inter_pool": "1",
            "asset_out_pool_id": "1",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair CMST-CMDX",
            "description": "adding extended pairs CMST-CMDX cross pool",
            "deposit": "10000000ucmdx",
        },
    ],
    [
        "ATOM-CMDX",
        {
            "asset_in": "1",
            "asset_out": "2",
            "is_inter_pool": "1",
            "asset_out_pool_id": "1",
            "min_usd_value_left": "100000000000",
            "title": "Add Extended pair ATOM-CMDX",
            "description": "adding extended pairs ATOM-CMDX cross pool",
            "deposit": "10000000ucmdx",
        },
    ],
]

LEND_ASSET_PAIR_MAPPING = [
    # [assetID, poolID, pairIDs]
    [1, 1, [3, 4, 15]],
    [2, 1, [1, 2, 13]],
    [3, 1, [5, 6, 14]],
    [4, 2, [7, 8, 16]],
    [1, 2, [9, 10, 18]],
    [3, 2, [11, 12, 17]],
]

WASM_CONTRACTS = [
    {
        "name": "Vesting Contract",
        "contractAddressKey": "vesting_contract",
        "contractLink": "https://github.com/comdex-official/test-wasm-artifacts/raw/main/token_vesting.wasm",
        "contractPath": f"{COMDEX_DIR_PATH}/scripts/comdex_local_setup/token_vesting.wasm",
        "initator": {},
        "formatKeys": []
    },
    {
        "name": "Locking Contract",
        "contractAddressKey": "locking_contract",
        "contractLink": "https://github.com/comdex-official/test-wasm-artifacts/raw/main/locking_contract.wasm",
        "contractPath": f"{COMDEX_DIR_PATH}/scripts/comdex_local_setup/locking_contract.wasm",
        "initator": {
            "t1": {"period": 500, "weight": "0.25"},
            "t2": {"period": 1000, "weight": "0.50"},
            "t3": {"period": 3000, "weight": "0.75"},
            "t4": {"period": 5000, "weight": "1.0"},
            "voting_period": 22500,
            "vesting_contract": "",
            "foundation_addr": ["comdex1rljg3wwgv6qezu3p05vxny9pwk3mdwl0ja407z"],
            "foundation_percentage": "0.2",
            "surplus_asset_id": 3,
            "emission": {
                "app_id": 1,
                "total_rewards": "10000000000000",
                "rewards_pending": "10000000000000",
                "emmission_rate": "0.01",
                "distributed_rewards": "0",
            },
        },
        "formatKeys": ['vesting_contract']
    },
    {
        "name": "Governance Contract",
        "contractAddressKey": "governance_contract",
        "contractLink": "https://github.com/comdex-official/test-wasm-artifacts/raw/main/governance.wasm",
        "contractPath": f"{COMDEX_DIR_PATH}/scripts/comdex_local_setup/governance.wasm",
        "initator": {
            "threshold": {"threshold_quorum": {"threshold": "0.50", "quorum": "0.33"}},
            "target": "0.0.0.0:9090",
            "locking_contract": "",
        },
        "formatKeys": ['locking_contract']
    },
]

WASM_PROPOSALS = [
    {
        "proposalID": 0,
        "isProposal": False,
        "contractAddressKey": "locking_contract",
        "content": {"lock": {"app_id": 1, "locking_period": "t3"}},
    },
    {
        "proposalID": 1,
        "isProposal": True,
        "contractAddressKey": "governance_contract",
        "content": {
            "propose": {
                "propose": {
                    "title": "New proposal for add vault pair for CMDX C - CMST",
                    "description": "This is a base execution proposal to add CMDX C - CMST vault pair with given Vault properties a. Liquidation ratio : 140 % b. Stability Fee : 1%  c. Liquidation Penalty : 12% d. DrawDown Fee : 1% e. Debt Cieling : 100000000 CMST f. Debt Floor : 100 CMST ",
                    "msgs": [
                        {
                            "msg_add_extended_pairs_vault": {
                                "app_id": 1,
                                "pair_id": 1,
                                "stability_fee": "0.025",
                                "closing_fee": "0.00",
                                "liquidation_penalty": "0.12",
                                "draw_down_fee": "0.001",
                                "is_vault_active": True,
                                "debt_ceiling": 1000000000000,
                                "debt_floor": 100000000,
                                "is_stable_mint_vault": False,
                                "min_cr": "1.7",
                                "pair_name": "ATOM-A",
                                "asset_out_oracle_price": False,
                                "asset_out_price": 1000000,
                                "min_usd_value_left": 100000,
                            }
                        }
                    ],
                    "app_id_param": 1,
                }
            }
        },
    },
    {
        "proposalID": 2,
        "isProposal": True,
        "contractAddressKey": "governance_contract",
        "content": {
            "propose": {
                "propose": {
                    "title": "New proposal to initialise collector param for stability fee and auction (surplus and debt)threshold data",
                    "description": "This is an base  execution proposal to initialise CMST and HARBOR pair for Surplus and dutch auction with Debt Threshold being  1000 CMST and Surplus Threshold as 100000000 CMST  ",
                    "msgs": [
                        {
                            "msg_set_collector_lookup_table": {
                                "app_id": 1,
                                "collector_asset_id": 3,
                                "secondary_asset_id": 9,
                                "surplus_threshold": 10000000000,
                                "debt_threshold": 10000000,
                                "locker_saving_rate": "0.06",
                                "lot_size": 20000,
                                "bid_factor": "0.01",
                                "debt_lot_size": 200,
                            }
                        }
                    ],
                    "app_id_param": 1,
                }
            }
        },
    },
    {
        "proposalID": 3,
        "isProposal": True,
        "contractAddressKey": "governance_contract",
        "content": {
            "propose": {
                "propose": {
                    "title": "New proposal for setting auction params for collateral auctions.",
                    "description": "This is an base proposal to set collateral auction params with auction duration being 600 seconds.",
                    "msgs": [
                        {
                            "msg_add_auction_params": {
                                "app_id": 1,
                                "auction_duration_seconds": 20,
                                "buffer": "1.2",
                                "cusp": "0.4",
                                "step": 360,
                                "price_function_type": 1,
                                "surplus_id": 1,
                                "debt_id": 2,
                                "dutch_id": 3,
                                "bid_duration_seconds": 10,
                            }
                        }
                    ],
                    "app_id_param": 1,
                }
            }
        },
    },
    {
        "proposalID": 4,
        "isProposal": True,
        "contractAddressKey": "governance_contract",
        "content": {
            "propose": {
                "propose": {
                    "title": "New proposal to whitelist CMST for locker",
                    "description": "This is an base  execution proposal to add use CMST as locker deposit asset.",
                    "msgs": [
                        {
                            "msg_set_auction_mapping_for_app": {
                                "app_id": 1,
                                "asset_id": 3,
                                "is_surplus_auction": False,
                                "is_debt_auction": False,
                                "is_distributor": True,
                                "asset_out_oracle_price": False,
                                "asset_out_price": 1000000,
                            }
                        }
                    ],
                    "app_id_param": 1,
                }
            }
        },
    },
    {
        "proposalID": 5,
        "isProposal": True,
        "contractAddressKey": "governance_contract",
        "content": {
            "propose": {
                "propose": {
                    "title": "New proposal to whitelist CMST for locker",
                    "description": "This is an base  execution proposal to add use CMST as locker deposit asset.",
                    "msgs": [
                        {"msg_white_list_asset_locker": {"app_id": 1, "asset_id": 3}}
                    ],
                    "app_id_param": 1,
                }
            }
        },
    },
    {
        "proposalID": 6,
        "isProposal": True,
        "contractAddressKey": "governance_contract",
        "content": {
            "propose": {
                "propose": {
                    "title": "New proposal to whitelist CMST for locker",
                    "description": "This is an base  execution proposal to add use CMST as locker deposit asset.",
                    "msgs": [
                        {
                            "msg_whitelist_app_id_locker_rewards": {
                                "app_id": 1,
                                "asset_id": 3,
                            }
                        }
                    ],
                    "app_id_param": 1,
                }
            }
        },
    },
    {
        "proposalID": 7,
        "isProposal": True,
        "contractAddressKey": "governance_contract",
        "content": {
            "propose": {
                "propose": {
                    "title": "New proposal to whitelist CMST for locker",
                    "description": "This is an base  execution proposal to add use CMST as locker deposit asset.",
                    "msgs": [{"msg_whitelist_app_id_vault_interest": {"app_id": 1}}],
                    "app_id_param": 1,
                }
            }
        },
    },
    {
        "proposalID": 8,
        "isProposal": True,
        "contractAddressKey": "governance_contract",
        "content": {
            "propose": {
                "propose": {
                    "title": "New proposal for add pair for CMDX",
                    "description": "This is an base proposal execution proposal to add CMDX-CMST n.",
                    "msgs": [{"msg_whitelist_app_id_liquidation": {"app_id": 1}}],
                    "app_id_param": 1,
                }
            }
        },
    },
    {
        "proposalID": 9,
        "isProposal": True,
        "contractAddressKey": "governance_contract",
        "content": {
            "propose": {
                "propose": {
                    "title": "New proposal for setting auction params for collateral auctions.",
                    "description": "This is an base proposal to set collateral auction params with auction duration being 600 seconds.",
                    "msgs": [
                        {
                            "msg_add_e_s_m_trigger_params": {
                                "app_id": 1,
                                "target_value": {"amount": "200", "denom": "uharbor"},
                                "cool_off_period": 60,
                                "asset_id": [3],
                                "rates": [1000000],
                            }
                        }
                    ],
                    "app_id_param": 1,
                }
            }
        },
    },
    {
        "proposalID": 10,
        "isProposal": True,
        "contractAddressKey": "governance_contract",
        "content": {
            "propose": {
                "propose": {
                    "title": "New proposal for add vault pair for CMDX C - CMST",
                    "description": "This is a base execution proposal to add CMDX C - CMST vault pair with given Vault properties a. Liquidation ratio : 140 % b. Stability Fee : 1%  c. Liquidation Penalty : 12% d. DrawDown Fee : 1% e. Debt Cieling : 100000000 CMST f. Debt Floor : 100 CMST ",
                    "msgs": [
                        {
                            "msg_add_extended_pairs_vault": {
                                "app_id": 1,
                                "pair_id": 2,
                                "stability_fee": "0.025",
                                "closing_fee": "0.00",
                                "liquidation_penalty": "0.12",
                                "draw_down_fee": "0.001",
                                "is_vault_active": True,
                                "debt_ceiling": 1000000000000,
                                "debt_floor": 100000000,
                                "is_stable_mint_vault": False,
                                "min_cr": "1.7",
                                "pair_name": "CMDX-A",
                                "asset_out_oracle_price": False,
                                "asset_out_price": 1000000,
                                "min_usd_value_left": 100000,
                            }
                        }
                    ],
                    "app_id_param": 1,
                }
            }
        },
    },
]
