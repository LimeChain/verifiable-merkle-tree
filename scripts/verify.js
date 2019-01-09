const etherlime = require('etherlime');
const MerkleLime = require('../build/MerkleLime.json');
const ethers = require('ethers');
const utils = ethers.utils;

// node scripts/verify.js 0xa00f6A3a3D00979D7B7E23D7C7dF6CC7E255Ad88 "Gitcoin Livestream Rocks" "[\"0xac11234732f084af283c6f0abcd30bbab34de31fc1ae3040ae8b91cbe6a18794\",\"0xf6c6901f1cd8f45d193642065a7b88f9d3549006be25adfd53cef07d8c6c434b\"]" 3
const verify = async () => {

	if (process.argv.length < 6) {
		throw new Error('Invalid arguments');
	}

	const contractAddress = process.argv[2];
	const originalData = process.argv[3];
	const hashes = JSON.parse(process.argv[4]);
	const index = process.argv[5];

	const merkleContract = await etherlime.ContractAt(MerkleLime, contractAddress)
	const isPart = await merkleContract.verifyDataInState(utils.toUtf8Bytes(originalData), hashes, index)
	console.log("This transaction is part of the Merkle tree: ", isPart);

};

verify()