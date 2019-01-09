const etherlime = require('etherlime');
const RootValidator = require('../build/RootValidator.json');
const ethers = require('ethers');
const utils = ethers.utils;

// node scripts/verify.js 0xcontractAddress "Lime Trees Rock" "[\"0x28966527d02480b0c603c1712eecb885f9083dd3503bc21183a7cf287ecb6cfc\",\"0xed863c82cb4fa9c1b3bc8cd1d73bd6bc5d75628264c4618305078c6e50d0c7fb\",\"0x1f2046f5ede7895de3666059b52edcc36e3fa4f8812bfd9ff34553f5aea45ec1\"]" 4
const verify = async () => {

	if (process.argv.length < 6) {
		throw new Error('Invalid arguments');
	}

	const contractAddress = process.argv[2];
	const originalData = process.argv[3];
	const hashes = JSON.parse(process.argv[4]);
	const index = process.argv[5];

	const merkleContract = await etherlime.ContractAt(RootValidator, contractAddress)
	const isPart = await merkleContract.verifyDataInState(utils.toUtf8Bytes(originalData), hashes, index)
	console.log("This transaction is part of the Merkle tree: ", isPart);

};

verify()