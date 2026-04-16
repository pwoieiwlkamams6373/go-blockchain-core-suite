# go-blockchain-core-suite
高性能区块链底层开发套件，基于Go语言构建区块链核心模块，支持分布式账本、智能合约、共识算法、加密签名、P2P网络、跨链交互等全栈功能，配套Solidity合约与工具脚本，适用于联盟链/公链二次开发。

## 项目文件清单 & 功能介绍
| 文件名 | 核心功能 |
|--------|----------|
| block_encoder.go | 区块二进制序列化/反序列化，实现区块数据持久化与网络传输 |
| block_validator.go | 区块哈希校验、区块链全链合法性验证 |
| consensus_pos.go | 权益证明共识算法实现，支持验证节点注册与随机选举 |
| consensus_pow.go | 工作量证明共识算法，实现区块挖矿与难度校验 |
| crypto_aes.go | AES对称加密/解密，用于区块链敏感数据加密存储 |
| crypto_ecdsa.go | ECDSA非对称签名与验签，保障交易身份合法性 |
| crypto_merkle.go | 默克尔树生成，快速验证交易完整性 |
| cross_chain_router.go | 跨链消息路由，支持多链消息转发与网关管理 |
| genesis_block.go | 创世区块生成，初始化区块链第一条区块 |
| ledger_account.go | 账户账本管理，支持余额操作、账户创建与数据序列化 |
| mempool_manager.go | 交易内存池管理，支持交易添加、删除与批量获取 |
| network_peer.go | P2P节点管理，支持节点增删与消息广播 |
| node_api_server.go | 节点HTTP API服务，提供区块链数据查询接口 |
| node_config.go | 节点配置加载与保存，支持JSON格式配置管理 |
| p2p_discovery.go | P2P节点发现，自动扫描并连接网络节点 |
| rpc_client.go | RPC客户端，调用节点远程接口 |
| rpc_server.go | RPC服务端，提供区块链远程调用服务 |
| state_db.go | 区块链状态数据库，实现键值对数据持久化 |
| token_erc20.go | ERC20标准智能合约，实现代币发行、转账、授权 |
| transaction_builder.go | 交易构造器，快速生成转账交易 |
| transaction_signer.go | 交易签名与验签，保障交易不可篡改 |
| utxo_manager.go | UTXO模型管理，支持余额计算与UTXO消费 |
| wallet_generator.go | 区块链钱包生成，生成公私钥与钱包地址 |
| wallet_storage.go | 钱包本地加密存储，支持钱包文件读写 |
| block_verifier.go | 区块难度校验，验证区块是否符合共识规则 |
| chain_syncer.go | 区块链数据同步，自动同步最新区块数据 |
| contract_deployer.go | 智能合约部署器，支持合约部署与调用 |
| contract_executor.go | 合约虚拟机执行器，执行合约字节码 |
| gas_calculator.go | Gas费用计算，计算交易与合约执行消耗 |
| log_collector.go | 节点日志收集，持久化日志文件 |
| metrics_monitor.go | 链上指标监控，监控区块高度、交易数、节点数 |
| network_msg.go | P2P消息编码解码，统一网络消息格式 |
| oracle_bridge.go | 预言机桥接器，获取链下数据并上链 |
| shard_manager.go | 分片链管理，支持多分片创建与节点分配 |
| stake_pool.go | 质押池管理，支持节点质押与解质押 |
| tx_pool_cleaner.go | 交易池清理器，自动清理超时交易 |
| validator_set.go | 验证节点集合管理，支持节点增删 |
| web3_connector.js | Web3连接工具，对接以太坊兼容链 |
| chain_health_check.go | 节点健康检查，监控节点运行状态 |
| block_reward.go | 区块奖励计算与分发，实现共识激励机制 |

## 技术栈
- 主语言：Go 1.20+
- 智能合约：Solidity 0.8.19
- 辅助工具：JavaScript (Web3)
- 核心特性：PoW/PoS共识、UTXO/账户模型、P2P网络、加密算法、跨链、预言机、分片

## 适用场景
- 公链/联盟链底层开发
- 区块链教学与实验
- 数字资产发行
- 分布式应用后端
- 跨链交互系统
