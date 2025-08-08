import { createWalletClient, http } from "viem";
import { privateKeyToAccount } from "viem/accounts";
import { base } from "viem/chains";
import { wrapFetchWithPayment } from "x402-fetch";

const dinGatewayURL = "https://ethereum-mainnet.gateway.din.dev/rpc";
const rpcReq = {
  jsonrpc: "2.0",
  method: "eth_getBlockByNumber",
  params: ["latest", false],
  id: 1,
};

const { X402_BUYER_PRIVATE_KEY } = process.env;

if (!X402_BUYER_PRIVATE_KEY) {
  console.error("X402_BUYER_PRIVATE_KEY not set");
  process.exit(1);
}

const account = privateKeyToAccount("0x" +X402_BUYER_PRIVATE_KEY as `0x${string}`);
const client = createWalletClient({
  account,
  transport: http(),
  chain: base,
});

const fetchWithPay = wrapFetchWithPayment(fetch, client);

const response = await fetchWithPay(dinGatewayURL, {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
  },
  body: JSON.stringify(rpcReq),
});

const data = await response.json();
console.log(data);
