const etherlime = require('etherlime');
const RootValidator = require('../build/RootValidator.json');
const ethers = require('ethers');
const utils = ethers.utils;

const verify = async () => {



	const contractAddress = "0xD813E6D0a509a615c968088f47358009c5Db9569";
	const originalData = "Forgae 124!4";
	const index = 100002;
	const hashes = [
		"0x7c209f20f29fccf0d7e1fcc0ef8c9722dcfc9a08a3c8b537f46d85a3f81cc93f",
		"0xec64688b7d881b6e27a05683cbbf3f585b1d2565a08a47daaa1844ea0b9c0c9d",
		"0xaf1346fa856a19dff7a51ad3c0ddecc529371c729ce5a26ccd4e5e995311c77e",
		"0x0b6625f901306fe9e6f11ab7cf7a1364e33a1c7d3b96e1c871ce995d7766c991",
		"0xf5e918f248c26b25e3e5886a0e200923ae8a962b6f0067e777ce4801c46b7640",
		"0xb843db0097ceacc3d4fcd8a267d8375f2f6e3b5c2350dd7afea0ac4d0500b21b",
		"0x0beb567bfce8088b8e940536c1c6a7d686a930079b8d6352e867d38019c6770d",
		"0x3446ca52f190f9d7e8f7bb4f12a54ca942899277dc655887d029bad616096e3c",
		"0xa6b2087403c4254b44ca0e889eab6f5160e0c792a928f50a6049d2d1c057ef05",
		"0x7ca3dcbfa08ed3a3f1235802624eaed64e3a14e810c41811f6291fd63f92e196",
		"0x927fb8363fa64441cde888159f7c88ec6ee98529ea107557ec1e7bb3c703b10d",
		"0x1844bbc3d8b61d18555cd1e90c00522d48cf0a5134df2e3d1c30de81a56e1f23",
		"0xcae7e899640d223233c298879b6d280421d79998839ec77aadf1181fa4271ae2",
		"0x17b80efdcea01c28a2ee84781d22b9af73fce9011cf96f0195a7602801c0592f",
		"0x02ec98d8c762185391e3f68a5e547f615ef7f3ca0f3a38a97c58a343b3c458c0",
		"0x0ee8990ed0eb2bdd183fd6324e61d8fbe129c15fa89e827ebec9ff8b50c3bc61",
		"0xf892dda0f2df2adbf2670e76471fde02f1edfd458b1e55ac2b41ed0cbc11c810"
	];

	const provider = new ethers.providers.InfuraProvider('rinkeby', 'H4UAAWyThMPs2WB9LsHD')
	const wallet = new ethers.Wallet('d723d3cdf932464de15845c0719ca13ce15e64c83625d86ddbfc217bd2ac5f5a', provider)
	const merkleContract = await etherlime.ContractAt(RootValidator, contractAddress, wallet, provider)
	const isPart = await merkleContract.verifyDataInState(utils.toUtf8Bytes(originalData), hashes, index)
	console.log("This transaction is part of the Merkle tree: ", isPart);

};

verify()