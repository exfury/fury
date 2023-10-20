import pytest

from .network import setup_fury
from .utils import CONTRACTS, deploy_contract, w3_wait_for_new_blocks


@pytest.fixture(scope="module")
def custom_fury(tmp_path_factory):
    path = tmp_path_factory.mktemp("storage-proof")
    yield from setup_fury(path, 26800)


@pytest.fixture(scope="module", params=["fury", "geth"])
def cluster(request, custom_fury, geth):
    """
    run on both fury and geth
    """
    provider = request.param
    if provider == "fury":
        yield custom_fury
    elif provider == "geth":
        yield geth
    else:
        raise NotImplementedError


def test_basic(cluster):
    # wait till height > 2 because
    # proof queries at height <= 2 are not supported
    if cluster.w3.eth.block_number <= 2:
        w3_wait_for_new_blocks(cluster.w3, 2)

    _, res = deploy_contract(
        cluster.w3,
        CONTRACTS["StateContract"],
    )
    method = "eth_getProof"
    storage_keys = ["0x0", "0x1"]
    proof = (
        cluster.w3.provider.make_request(
            method, [res["contractAddress"], storage_keys, hex(res["blockNumber"])]
        )
    )["result"]
    for proof in proof["storageProof"]:
        if proof["key"] == storage_keys[0]:
            assert proof["value"] != "0x0"
