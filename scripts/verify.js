const etherlime = require('etherlime');
const RootValidator = require('../build/RootValidator.json');
const ethers = require('ethers');
const utils = ethers.utils;

/* node scripts/verify.js 0x9eD274314f0fB37837346C425D3cF28d89ca9599 "Etherlime Rulez4123" "[\"0xbde6939616df344b350a2a79e38775e42c5a3a38ceb599330c261d106170251d\",\"0x4d7290d7337e07588f3e708a9905ce5eea3e1a92983680d0f21149e25a5b26ef\",\"0x047d1da351a51b1d4e3dcfd38633f570ed2e852456b85a96ad83ebf67bf5d763\"]" 4
*/
const verify = async () => {

	if (process.argv.length < 6) {
		throw new Error('Invalid arguments');
	}

	const contractAddress = process.argv[2];
	const originalData = process.argv[3];
	const hashes = [
		"0x06d1078a22271274504f45c04e79a81631d3451a6b379619e38b0edef0d5101e",
		"0x1838d6076beca4c5a58f9797d56ab72d9d8dce30b184ad452554157bc2029317",
		"0x90f0cebca5479de2a926fc19c2855a8da412b25fce1d6b511130ad76685d6309",
		"0x697f3c3c44bfd8fbdf4809823e3a7627e683164256a495a9e0a76bc539de3264",
		"0x815ea8575558be074a01f4678b329c777aa5399e9ce8a84a2976e1ee28e1a509",
		"0xb843db0097ceacc3d4fcd8a267d8375f2f6e3b5c2350dd7afea0ac4d0500b21b",
		"0x72488ea4045c69e1727dfeacc715647a1c3835f0b7c792047af81461b23f49f2",
		"0x3446ca52f190f9d7e8f7bb4f12a54ca942899277dc655887d029bad616096e3c",
		"0x3b9b3ea08edad14116fef6c6f373a349a3434b7bce71aefef6180a3760e2192a",
		"0x7ca3dcbfa08ed3a3f1235802624eaed64e3a14e810c41811f6291fd63f92e196",
		"0x927fb8363fa64441cde888159f7c88ec6ee98529ea107557ec1e7bb3c703b10d",
		"0xdd46f56986aa7a215178c497f17f7f8a6648d553849e994e65973e49fde9f4b1",
		"0xadadf206abbb37075d90a43c4537da664a9df40ddb9331a6653c09aaff53c91f",
		"0xdfe95ca9b44f606a6ae5e1041c28bfb6c89de087ebf1510057363270ff34968b",
		"0x9d766c4085fc8ccaed78a6baaddcd44c31940629b3cfe7b142ec0713163d3437",
		"0x0ee8990ed0eb2bdd183fd6324e61d8fbe129c15fa89e827ebec9ff8b50c3bc61",
		"0xf892dda0f2df2adbf2670e76471fde02f1edfd458b1e55ac2b41ed0cbc11c810"
	];
	const index = process.argv[5];

	const merkleContract = await etherlime.ContractAt(RootValidator, contractAddress)
	const isPart = await merkleContract.verifyDataInState(utils.toUtf8Bytes(originalData), hashes, index)
	console.log("This transaction is part of the Merkle tree: ", isPart);

};

verify()