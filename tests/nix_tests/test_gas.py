from .utils import (
    ADDRS,
    CONTRACTS,
    KEYS,
    deploy_contract,
    send_transaction,
    w3_wait_for_new_blocks,
)


def test_gas_eth_tx(geth, fury):
    tx_value = 10

    # send a transaction with geth
    geth_gas_price = geth.w3.eth.gas_price
    tx = {"to": ADDRS["community"], "value": tx_value, "gasPrice": geth_gas_price}
    geth_receipt = send_transaction(geth.w3, tx, KEYS["validator"])

    # send an equivalent transaction with fury
    fury_gas_price = fury.w3.eth.gas_price
    tx = {"to": ADDRS["community"], "value": tx_value, "gasPrice": fury_gas_price}
    fury_receipt = send_transaction(fury.w3, tx, KEYS["validator"])

    # ensure that the gasUsed is equivalent
    assert geth_receipt.gasUsed == fury_receipt.gasUsed


def test_gas_deployment(geth, fury):
    # deploy an identical contract on geth and fury
    # ensure that the gasUsed is equivalent
    _, geth_contract_receipt = deploy_contract(geth.w3, CONTRACTS["TestERC20A"])
    _, fury_contract_receipt = deploy_contract(fury.w3, CONTRACTS["TestERC20A"])
    assert geth_contract_receipt.gasUsed == fury_contract_receipt.gasUsed


def test_gas_call(geth, fury):
    function_input = 10

    # deploy an identical contract on geth and fury
    # ensure that the contract has a function which consumes non-trivial gas
    geth_contract, _ = deploy_contract(geth.w3, CONTRACTS["BurnGas"])
    fury_contract, _ = deploy_contract(fury.w3, CONTRACTS["BurnGas"])

    # call the contract and get tx receipt for geth
    geth_gas_price = geth.w3.eth.gas_price
    geth_txhash = geth_contract.functions.burnGas(function_input).transact(
        {"from": ADDRS["validator"], "gasPrice": geth_gas_price}
    )
    geth_call_receipt = geth.w3.eth.wait_for_transaction_receipt(geth_txhash)

    # repeat the above for fury
    fury_gas_price = fury.w3.eth.gas_price
    fury_txhash = fury_contract.functions.burnGas(function_input).transact(
        {"from": ADDRS["validator"], "gasPrice": fury_gas_price}
    )
    fury_call_receipt = fury.w3.eth.wait_for_transaction_receipt(fury_txhash)

    # ensure that the gasUsed is equivalent
    assert geth_call_receipt.gasUsed == fury_call_receipt.gasUsed


def test_block_gas_limit(fury):
    tx_value = 10

    # get the block gas limit from the latest block
    w3_wait_for_new_blocks(fury.w3, 5)
    block = fury.w3.eth.get_block("latest")
    exceeded_gas_limit = block.gasLimit + 100

    # send a transaction exceeding the block gas limit
    fury_gas_price = fury.w3.eth.gas_price
    tx = {
        "to": ADDRS["community"],
        "value": tx_value,
        "gas": exceeded_gas_limit,
        "gasPrice": fury_gas_price,
    }

    # expect an error due to the block gas limit
    try:
        send_transaction(fury.w3, tx, KEYS["validator"])
    except Exception as error:
        assert "exceeds block gas limit" in error.args[0]["message"]

    # deploy a contract on fury
    fury_contract, _ = deploy_contract(fury.w3, CONTRACTS["BurnGas"])

    # expect an error on contract call due to block gas limit
    try:
        fury_txhash = fury_contract.functions.burnGas(exceeded_gas_limit).transact(
            {
                "from": ADDRS["validator"],
                "gas": exceeded_gas_limit,
                "gasPrice": fury_gas_price,
            }
        )
        (fury.w3.eth.wait_for_transaction_receipt(fury_txhash))
    except Exception as error:
        assert "exceeds block gas limit" in error.args[0]["message"]

    return
