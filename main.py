import os
import asyncio
from eth_account import Account
from x402.clients.httpx import x402HttpxClient

async def main():
    din_gateway_url = "https://ethereum-mainnet.gateway.din.dev"
    rpc_req = '{"jsonrpc": "2.0","method": "eth_getBlockByNumber","params": ["latest", false],"id": 1}'

    private_key = os.environ.get("X402_BUYER_PRIVATE_KEY")
    if not private_key:
        print("X402_BUYER_PRIVATE_KEY not set")
        exit(1)

    account = Account.from_key(private_key)

    async with x402HttpxClient(account=account, base_url=din_gateway_url) as client:
        resp = await client.post(url="/rpc", data=rpc_req)
        print(await resp.aread())

if __name__ == "__main__":
    asyncio.run(main())
