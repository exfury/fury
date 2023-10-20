import pytest

from .network import setup_fury, setup_geth


@pytest.fixture(scope="session")
def fury(tmp_path_factory):
    path = tmp_path_factory.mktemp("fury")
    yield from setup_fury(path, 26650)


@pytest.fixture(scope="session")
def geth(tmp_path_factory):
    path = tmp_path_factory.mktemp("geth")
    yield from setup_geth(path, 8545)


@pytest.fixture(scope="session", params=["fury", "fury-ws"])
def fury_rpc_ws(request, fury):
    """
    run on both fury and fury websocket
    """
    provider = request.param
    if provider == "fury":
        yield fury
    elif provider == "fury-ws":
        fury_ws = fury.copy()
        fury_ws.use_websocket()
        yield fury_ws
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["fury", "fury-ws", "geth"])
def cluster(request, fury, geth):
    """
    run on fury, fury websocket and geth
    """
    provider = request.param
    if provider == "fury":
        yield fury
    elif provider == "fury-ws":
        fury_ws = fury.copy()
        fury_ws.use_websocket()
        yield fury_ws
    elif provider == "geth":
        yield geth
    else:
        raise NotImplementedError
