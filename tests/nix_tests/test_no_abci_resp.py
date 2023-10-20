from pathlib import Path

import pytest

from .network import setup_custom_fury
from .utils import w3_wait_for_new_blocks


@pytest.fixture(scope="module")
def fury(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp")
    yield from setup_custom_fury(
        path,
        26260,
        Path(__file__).parent / "configs/discard-abci-resp.jsonnet",
    )


def test_gas_eth_tx(fury):
    """
    When node does not persist ABCI responses
    eth_gasPrice should return an error instead of crashing
    """
    w3_wait_for_new_blocks(fury.w3, 3)
    try:
        fury.w3.eth.gas_price
        raise Exception("This query should have failed")
    except Exception as error:
        assert "node is not persisting abci responses" in error.args[0]["message"]
