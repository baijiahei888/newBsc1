package vm

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	"github.com/baijiahei888/newBsc1/core/vm/lightclient"
	"github.com/stretchr/testify/require"
)

const (
	testHeight uint64 = 66848226
)

func TestTmHeaderValidateAndMerkleProofValidate(t *testing.T) {
	consensusStateBytes, err := hex.DecodeString("42696e616e63652d436861696e2d4e696c650000000000000000000000000000000000000" +
		"3fc05e2b7029751d2a6581efc2f79712ec44d8b4981850325a7feadaa58ef4ddaa18a9380d9ab0fc10d18ca0e0832d5f4c063c5489ec1443dfb738" +
		"252d038a82131b27ae17cbe9c20cdcfdf876b3b12978d3264a007fcaaa71c4cdb701d9ebc0323f44f000000174876e800184e7b103d34c41003f9b" +
		"864d5f8c1adda9bd0436b253bb3c844bc739c1e77c9000000174876e8004d420aea843e92a0cfe69d89696dff6827769f9cb52a249af537ce89bf2" +
		"a4b74000000174876e800bd03de9f8ab29e2800094e153fac6f696cfa512536c9c2f804dcb2c2c4e4aed6000000174876e8008f4a74a07351895dd" +
		"f373057b98fae6dfaf2cd21f37a063e19601078fe470d53000000174876e8004a5d4753eb79f92e80efe22df7aca4f666a4f44bf81c536c4a09d4b" +
		"9c5b654b5000000174876e800c80e9abef7ff439c10c68fe8f1303deddfc527718c3b37d8ba6807446e3c827a000000174876e8009142afcc691b7" +
		"cc05d26c7b0be0c8b46418294171730e079f384fde2fa50bafc000000174876e80049b288e4ebbb3a281c2d546fc30253d5baf08993b6e5d295fb7" +
		"87a5b314a298e000000174876e80004224339688f012e649de48e241880092eaa8f6aa0f4f14bfcf9e0c76917c0b6000000174876e8004034b37ce" +
		"da8a0bf13b1abaeee7a8f9383542099a554d219b93d0ce69e3970e8000000174876e800")
	require.NoError(t, err)

	cs, err := lightclient.DecodeConsensusState(consensusStateBytes)
	require.NoError(t, err)

	headerBytes, err := hex.DecodeString("e3210a92130abb020a02080a121242696e616e63652d436861696e2d4e696c6518e38bf01f220c08e191aef20510f5f4e4c70230dae0c7173a480a20102b54820dd8fb5bc2c4e875ee573fa294d9b7b7ceb362aa8fd21b33dee41b1c12240801122082f341511f3e6b89d6177fd31f8a106013ba09d6e12ef40a7dec885d81b687634220b1b77e6977e0cd0177e3102a78833c9e152aa646ed4fb5a77e8af58c9867eec0522080d9ab0fc10d18ca0e0832d5f4c063c5489ec1443dfb738252d038a82131b27a5a2080d9ab0fc10d18ca0e0832d5f4c063c5489ec1443dfb738252d038a82131b27a6220294d8fbd0b94b767a7eba9840f299a3586da7fe6b5dead3b7eecba193c400f936a20a3e248bc209955054d880e4d89ff3c0419c0cd77681f4b4c6649ead5545054b982011462633d9db7ed78e951f79913fdc8231aa77ec12b12d1100a480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be212b601080210e38bf01f22480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be22a0b08e291aef20510cebfe23e321406fd60078eb4c2356137dd50036597db267cf61642409276f20ad4b152f91c344bd63ac691bad66e04e228a8b58dca293ff0bd10f8aef6dfbcecae49e32b09d89e10b771a6c01628628596a95e126b04763560c66c0f12b801080210e38bf01f22480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be22a0b08e291aef20510a4caa532321418e69cc672973992bb5f76d049a5b2c5ddf77436380142409ed2b74fa835296d552e68c439dd4ee3fa94fb197282edcc1cc815c863ca42a2c9a73475ff6be9064371a61655a3c31d2f0acc89c3a4489ad4c2671aef52360512b801080210e38bf01f22480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be22a0b08e291aef20510a69eca2f3214344c39bb8f4512d6cab1f6aafac1811ef9d8afdf38024240de2768ead90011bcbb1914abc1572749ab7b81382eb81cff3b41c56edc12470a7b8a4d61f8b4ca7b2cb7e24706edd219455796b4db74cd36965859f91dc8910312b801080210e38bf01f22480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be22a0b08e291aef20510dcdd833b321437ef19af29679b368d2b9e9de3f8769b357866763803424072ddfe0aeb13616b3f17eb60b19a923ec51fcc726625094aa069255c829c8cdd9e242080a1e559b0030fe9a0db19fd34e392bd78df12a9caff9f2b811bc1ac0a12b801080210e38bf01f22480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be22a0b08e291aef20510e9f2f859321462633d9db7ed78e951f79913fdc8231aa77ec12b38044240f5f61c640ab2402b44936de0d24e7b439df78bc3ef15467ecb29b92ece4aa0550790d5ce80761f2ac4b0e3283969725c42343749d9b44b179b2d4fced66c5d0412b801080210e38bf01f22480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be22a0b08e291aef20510ff90f55532147b343e041ca130000a8bc00c35152bd7e774003738054240df6e298b3efd42eb536e68a0210bc921e8b5dc145fe965f63f4d3490064f239f2a54a6db16c96086e4ae52280c04ad8b32b44f5ff3d41f0c364949ccb628c50312b801080210e38bf01f22480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be22a0b08e291aef20510cad7c931321491844d296bd8e591448efc65fd6ad51a888d58fa3806424030298627da1afd28229aac150f553724b594989e59136d6a175d84e45a4dee344ff9e0eeb69fdf29abb6d833adc3e1ccdc87b2a65019ef5fb627c44d9d132c0012b801080210e38bf01f22480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be22a0b08e291aef20510c8c296323214b3727172ce6473bc780298a2d66c12f1a14f5b2a38074240918491100730b4523f0c85409f6d1cca9ebc4b8ca6df8d55fe3d85158fa43286608693c50332953e1d3b93e3e78b24e158d6a2275ce8c6c7c07a7a646a19200312b801080210e38bf01f22480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be22a0b08e291aef2051086f1a2403214b6f20c7faa2b2f6f24518fa02b71cb5f4a09fba338084240ca59c9fc7f6ab660e9970fc03e5ed588ccb8be43fe5a3e8450287b726f29d039e53fe888438f178ac63c3d2ca969cd8c2fbc8606f067634339b6a94a7382960212b801080210e38bf01f22480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be22a0b08e291aef2051080efbb543214e0dd72609cc106210d1aa13936cb67b93a0aee2138094240e787a21f5cb7052624160759a9d379dd9db144f2b498bca026375c9ce8ecdc2a0936af1c309b3a0f686c92bf5578b595a4ca99036a19c9fc50d3718fd454b30012b801080210e38bf01f22480a207eaabf7df1081377e06e08efe7ad17974049380bdd65a9b053c099ef80ff6e6f122408011220d153cc308d9cb96ca43ffeceaae1ee85794c83d17408ff76cfee92f5e91d0be22a0b08e291aef20510ddf8d85a3214fc3108dc3814888f4187452182bc1baf83b71bc9380a4240d51ea31f6449eed71de22339722af1edbb0b21401037d85882b32a2ed8ae9127f2df4d1da2092729e582812856227ed6cdf98a3f60203d1ff80bd635fb03bb0912a4070a4f0a1406fd60078eb4c2356137dd50036597db267cf61612251624de6420e17cbe9c20cdcfdf876b3b12978d3264a007fcaaa71c4cdb701d9ebc0323f44f1880d0dbc3f4022080e0ebdaf2e2ffffff010a4b0a1418e69cc672973992bb5f76d049a5b2c5ddf7743612251624de6420184e7b103d34c41003f9b864d5f8c1adda9bd0436b253bb3c844bc739c1e77c91880d0dbc3f4022080d0dbc3f4020a4b0a14344c39bb8f4512d6cab1f6aafac1811ef9d8afdf12251624de64204d420aea843e92a0cfe69d89696dff6827769f9cb52a249af537ce89bf2a4b741880d0dbc3f4022080d0dbc3f4020a4b0a1437ef19af29679b368d2b9e9de3f8769b3578667612251624de6420bd03de9f8ab29e2800094e153fac6f696cfa512536c9c2f804dcb2c2c4e4aed61880d0dbc3f4022080d0dbc3f4020a4b0a1462633d9db7ed78e951f79913fdc8231aa77ec12b12251624de64208f4a74a07351895ddf373057b98fae6dfaf2cd21f37a063e19601078fe470d531880d0dbc3f4022080d0dbc3f4020a4b0a147b343e041ca130000a8bc00c35152bd7e774003712251624de64204a5d4753eb79f92e80efe22df7aca4f666a4f44bf81c536c4a09d4b9c5b654b51880d0dbc3f4022080d0dbc3f4020a4b0a1491844d296bd8e591448efc65fd6ad51a888d58fa12251624de6420c80e9abef7ff439c10c68fe8f1303deddfc527718c3b37d8ba6807446e3c827a1880d0dbc3f4022080d0dbc3f4020a4b0a14b3727172ce6473bc780298a2d66c12f1a14f5b2a12251624de64209142afcc691b7cc05d26c7b0be0c8b46418294171730e079f384fde2fa50bafc1880d0dbc3f4022080d0dbc3f4020a4b0a14b6f20c7faa2b2f6f24518fa02b71cb5f4a09fba312251624de642049b288e4ebbb3a281c2d546fc30253d5baf08993b6e5d295fb787a5b314a298e1880d0dbc3f4022080d0dbc3f4020a4b0a14e0dd72609cc106210d1aa13936cb67b93a0aee2112251624de642004224339688f012e649de48e241880092eaa8f6aa0f4f14bfcf9e0c76917c0b61880d0dbc3f4022080d0dbc3f4020a4b0a14fc3108dc3814888f4187452182bc1baf83b71bc912251624de64204034b37ceda8a0bf13b1abaeee7a8f9383542099a554d219b93d0ce69e3970e81880d0dbc3f4022080d0dbc3f402124f0a1406fd60078eb4c2356137dd50036597db267cf61612251624de6420e17cbe9c20cdcfdf876b3b12978d3264a007fcaaa71c4cdb701d9ebc0323f44f1880d0dbc3f4022080e0ebdaf2e2ffffff011aa4070a4f0a1406fd60078eb4c2356137dd50036597db267cf61612251624de6420e17cbe9c20cdcfdf876b3b12978d3264a007fcaaa71c4cdb701d9ebc0323f44f1880d0dbc3f4022080e0ebdaf2e2ffffff010a4b0a1418e69cc672973992bb5f76d049a5b2c5ddf7743612251624de6420184e7b103d34c41003f9b864d5f8c1adda9bd0436b253bb3c844bc739c1e77c91880d0dbc3f4022080d0dbc3f4020a4b0a14344c39bb8f4512d6cab1f6aafac1811ef9d8afdf12251624de64204d420aea843e92a0cfe69d89696dff6827769f9cb52a249af537ce89bf2a4b741880d0dbc3f4022080d0dbc3f4020a4b0a1437ef19af29679b368d2b9e9de3f8769b3578667612251624de6420bd03de9f8ab29e2800094e153fac6f696cfa512536c9c2f804dcb2c2c4e4aed61880d0dbc3f4022080d0dbc3f4020a4b0a1462633d9db7ed78e951f79913fdc8231aa77ec12b12251624de64208f4a74a07351895ddf373057b98fae6dfaf2cd21f37a063e19601078fe470d531880d0dbc3f4022080d0dbc3f4020a4b0a147b343e041ca130000a8bc00c35152bd7e774003712251624de64204a5d4753eb79f92e80efe22df7aca4f666a4f44bf81c536c4a09d4b9c5b654b51880d0dbc3f4022080d0dbc3f4020a4b0a1491844d296bd8e591448efc65fd6ad51a888d58fa12251624de6420c80e9abef7ff439c10c68fe8f1303deddfc527718c3b37d8ba6807446e3c827a1880d0dbc3f4022080d0dbc3f4020a4b0a14b3727172ce6473bc780298a2d66c12f1a14f5b2a12251624de64209142afcc691b7cc05d26c7b0be0c8b46418294171730e079f384fde2fa50bafc1880d0dbc3f4022080d0dbc3f4020a4b0a14b6f20c7faa2b2f6f24518fa02b71cb5f4a09fba312251624de642049b288e4ebbb3a281c2d546fc30253d5baf08993b6e5d295fb787a5b314a298e1880d0dbc3f4022080d0dbc3f4020a4b0a14e0dd72609cc106210d1aa13936cb67b93a0aee2112251624de642004224339688f012e649de48e241880092eaa8f6aa0f4f14bfcf9e0c76917c0b61880d0dbc3f4022080d0dbc3f4020a4b0a14fc3108dc3814888f4187452182bc1baf83b71bc912251624de64204034b37ceda8a0bf13b1abaeee7a8f9383542099a554d219b93d0ce69e3970e81880d0dbc3f4022080d0dbc3f402124f0a1406fd60078eb4c2356137dd50036597db267cf61612251624de6420e17cbe9c20cdcfdf876b3b12978d3264a007fcaaa71c4cdb701d9ebc0323f44f1880d0dbc3f4022080e0ebdaf2e2ffffff01")
	require.NoError(t, err)

	parameterInput := make([]byte, 32+len(consensusStateBytes)+len(headerBytes))
	binary.BigEndian.PutUint64(parameterInput[24:32], uint64(len(consensusStateBytes)))
	copy(parameterInput[32:32+len(consensusStateBytes)], consensusStateBytes)
	copy(parameterInput[32+len(consensusStateBytes):], headerBytes)

	totalLengthPrefix := make([]byte, 32)
	binary.BigEndian.PutUint64(totalLengthPrefix[0:8], 0)
	binary.BigEndian.PutUint64(totalLengthPrefix[8:16], 0)
	binary.BigEndian.PutUint64(totalLengthPrefix[16:24], 0)
	binary.BigEndian.PutUint64(totalLengthPrefix[24:], uint64(len(parameterInput)))
	input := append(totalLengthPrefix, parameterInput...)

	var tmHeaderValidateContract tmHeaderValidate
	syncedConsensusStateBytes, err := tmHeaderValidateContract.Run(input)
	require.NoError(t, err)
	syncedConsensusState, err := lightclient.DecodeConsensusState(syncedConsensusStateBytes[32:])
	require.NoError(t, err)
	require.Equal(t, testHeight+1, syncedConsensusState.Height)
	require.Equal(t, cs.ChainID, syncedConsensusState.ChainID)

	headerBytes, err = hex.DecodeString("c11d0af00e0aba020a02080a121242696e616e63652d436861696e2d4e696c65188596a822220b08bd85c2f30510e8f3d65730a986f7173a480a209e4b742cb3d59618c5af9d9b9d8b6a8408802131a968c6d1e3a516c9548cedc01224080112201dd8f3e89b01db52d79f89a87c8abfe9836e8b828ed33d78e202664dfead9c0e422018c8c46d7665a5db717331ff3deebe8f9986352d308c5950cf1a09286439dd9d522080d9ab0fc10d18ca0e0832d5f4c063c5489ec1443dfb738252d038a82131b27a5a2080d9ab0fc10d18ca0e0832d5f4c063c5489ec1443dfb738252d038a82131b27a6220294d8fbd0b94b767a7eba9840f299a3586da7fe6b5dead3b7eecba193c400f936a202c69c314b4de5c8035253c8bc0771d9ca17b1b23a57c0c6d068b57579791cae2820114b3727172ce6473bc780298a2d66c12f1a14f5b2a12b00c0a480a20c24c5e913fe4b8eec95d473d671c2dfc2cf651118b37e83661a24a1b8ef842251224080112206310a09cbcf41fd71c55e524c7a91f71414bb31ecf6e7b74161e8ec04a3c7b2b120012b9010802108596a82222480a20c24c5e913fe4b8eec95d473d671c2dfc2cf651118b37e83661a24a1b8ef842251224080112206310a09cbcf41fd71c55e524c7a91f71414bb31ecf6e7b74161e8ec04a3c7b2b2a0c08bd85c2f30510cff5c5a102321418e69cc672973992bb5f76d049a5b2c5ddf7743638014240129b1ac9bcf178022f358161f90a0c83d8344be7ad7e1f55c98b4a7f366a95816e80a7e2beac669c73fa5bd6e17693f8a70347ab785c7883a8ef2c9d7464b50a12b9010802108596a82222480a20c24c5e913fe4b8eec95d473d671c2dfc2cf651118b37e83661a24a1b8ef842251224080112206310a09cbcf41fd71c55e524c7a91f71414bb31ecf6e7b74161e8ec04a3c7b2b2a0c08bd85c2f30510e3f7919d023214344c39bb8f4512d6cab1f6aafac1811ef9d8afdf38024240ad6d56ffd8062b70f6c2bf7345a26fcb0a5de2bc477d21e71d780c7f029098aba7efaa2aae1b945602c22d0d1ccdec1e45e7d14d177677a2ba23feeffa88c103120012b9010802108596a82222480a20c24c5e913fe4b8eec95d473d671c2dfc2cf651118b37e83661a24a1b8ef842251224080112206310a09cbcf41fd71c55e524c7a91f71414bb31ecf6e7b74161e8ec04a3c7b2b2a0c08bd85c2f305109c88b58202321462633d9db7ed78e951f79913fdc8231aa77ec12b3804424017c205e1eade76097e659a2c3a20bd7fc61e383b5b631499a8bf26ac81c537d7369ddc5be0c4f8a62c6cfb55d5d6cd4ce75aa57c03888aef365d0702032e6d0c12b9010802108596a82222480a20c24c5e913fe4b8eec95d473d671c2dfc2cf651118b37e83661a24a1b8ef842251224080112206310a09cbcf41fd71c55e524c7a91f71414bb31ecf6e7b74161e8ec04a3c7b2b2a0c08bd85c2f30510bcbda5ff0132147b343e041ca130000a8bc00c35152bd7e7740037380542404c8536a8a2af12e88176699c93ee6c9032240b58b6f92897a7724bd41447f81ae8533e2e3676a44d398296006c1d51d14c9eb2ebca5ac02a5925c549de571b0012b9010802108596a82222480a20c24c5e913fe4b8eec95d473d671c2dfc2cf651118b37e83661a24a1b8ef842251224080112206310a09cbcf41fd71c55e524c7a91f71414bb31ecf6e7b74161e8ec04a3c7b2b2a0c08bd85c2f30510968bff9c02321491844d296bd8e591448efc65fd6ad51a888d58fa38064240f1b55166334e16b2cf1e6bf35521a17b6fc8699807df66495c14508ed5c75b27bf0ba19f7ad4302166d673f027a37148f6ecf4b87fe60b6f196bfc944df3c90412b9010802108596a82222480a20c24c5e913fe4b8eec95d473d671c2dfc2cf651118b37e83661a24a1b8ef842251224080112206310a09cbcf41fd71c55e524c7a91f71414bb31ecf6e7b74161e8ec04a3c7b2b2a0c08bd85c2f30510899ab59f023214b3727172ce6473bc780298a2d66c12f1a14f5b2a38074240b90deb286343c378597d36c381b11f720b9994cc4aefb28ebe840240bd68b4e8855b8508f072133e36a92f4160e5172dd4e086cde6799dc957e3d971784a130e120012b9010802108596a82222480a20c24c5e913fe4b8eec95d473d671c2dfc2cf651118b37e83661a24a1b8ef842251224080112206310a09cbcf41fd71c55e524c7a91f71414bb31ecf6e7b74161e8ec04a3c7b2b2a0c08bd85c2f30510e7c1dffe013214e0dd72609cc106210d1aa13936cb67b93a0aee213809424085424cd5908cc12b9ce35b4891f20af1791f5a7a8f7dfef225a790890a1639badef168a57005ad64d167faf86c97c4bdfeba1d5307182abeb4ef768b4056c00d12b9010802108596a82222480a20c24c5e913fe4b8eec95d473d671c2dfc2cf651118b37e83661a24a1b8ef842251224080112206310a09cbcf41fd71c55e524c7a91f71414bb31ecf6e7b74161e8ec04a3c7b2b2a0c08bd85c2f305108bd8c5ff013214fc3108dc3814888f4187452182bc1baf83b71bc9380a424013382f84fa5d21f6a274159300b5fb20e21705c80b45beff6ca48d84ffa7ab700275e8017bcf69ac70f0c76bd497977c7373e3a9cfdfc09355cf5427272e4d0512a4070a4f0a1406fd60078eb4c2356137dd50036597db267cf61612251624de6420e17cbe9c20cdcfdf876b3b12978d3264a007fcaaa71c4cdb701d9ebc0323f44f1880d0dbc3f4022080e0ebdaf2e2ffffff010a4b0a1418e69cc672973992bb5f76d049a5b2c5ddf7743612251624de6420184e7b103d34c41003f9b864d5f8c1adda9bd0436b253bb3c844bc739c1e77c91880d0dbc3f4022080d0dbc3f4020a4b0a14344c39bb8f4512d6cab1f6aafac1811ef9d8afdf12251624de64204d420aea843e92a0cfe69d89696dff6827769f9cb52a249af537ce89bf2a4b741880d0dbc3f4022080d0dbc3f4020a4b0a1437ef19af29679b368d2b9e9de3f8769b3578667612251624de6420bd03de9f8ab29e2800094e153fac6f696cfa512536c9c2f804dcb2c2c4e4aed61880d0dbc3f4022080d0dbc3f4020a4b0a1462633d9db7ed78e951f79913fdc8231aa77ec12b12251624de64208f4a74a07351895ddf373057b98fae6dfaf2cd21f37a063e19601078fe470d531880d0dbc3f4022080d0dbc3f4020a4b0a147b343e041ca130000a8bc00c35152bd7e774003712251624de64204a5d4753eb79f92e80efe22df7aca4f666a4f44bf81c536c4a09d4b9c5b654b51880d0dbc3f4022080d0dbc3f4020a4b0a1491844d296bd8e591448efc65fd6ad51a888d58fa12251624de6420c80e9abef7ff439c10c68fe8f1303deddfc527718c3b37d8ba6807446e3c827a1880d0dbc3f4022080d0dbc3f4020a4b0a14b3727172ce6473bc780298a2d66c12f1a14f5b2a12251624de64209142afcc691b7cc05d26c7b0be0c8b46418294171730e079f384fde2fa50bafc1880d0dbc3f4022080d0dbc3f4020a4b0a14b6f20c7faa2b2f6f24518fa02b71cb5f4a09fba312251624de642049b288e4ebbb3a281c2d546fc30253d5baf08993b6e5d295fb787a5b314a298e1880d0dbc3f4022080d0dbc3f4020a4b0a14e0dd72609cc106210d1aa13936cb67b93a0aee2112251624de642004224339688f012e649de48e241880092eaa8f6aa0f4f14bfcf9e0c76917c0b61880d0dbc3f4022080d0dbc3f4020a4b0a14fc3108dc3814888f4187452182bc1baf83b71bc912251624de64204034b37ceda8a0bf13b1abaeee7a8f9383542099a554d219b93d0ce69e3970e81880d0dbc3f4022080d0dbc3f402124f0a1406fd60078eb4c2356137dd50036597db267cf61612251624de6420e17cbe9c20cdcfdf876b3b12978d3264a007fcaaa71c4cdb701d9ebc0323f44f1880d0dbc3f4022080e0ebdaf2e2ffffff011aa4070a4f0a1406fd60078eb4c2356137dd50036597db267cf61612251624de6420e17cbe9c20cdcfdf876b3b12978d3264a007fcaaa71c4cdb701d9ebc0323f44f1880d0dbc3f4022080e0ebdaf2e2ffffff010a4b0a1418e69cc672973992bb5f76d049a5b2c5ddf7743612251624de6420184e7b103d34c41003f9b864d5f8c1adda9bd0436b253bb3c844bc739c1e77c91880d0dbc3f4022080d0dbc3f4020a4b0a14344c39bb8f4512d6cab1f6aafac1811ef9d8afdf12251624de64204d420aea843e92a0cfe69d89696dff6827769f9cb52a249af537ce89bf2a4b741880d0dbc3f4022080d0dbc3f4020a4b0a1437ef19af29679b368d2b9e9de3f8769b3578667612251624de6420bd03de9f8ab29e2800094e153fac6f696cfa512536c9c2f804dcb2c2c4e4aed61880d0dbc3f4022080d0dbc3f4020a4b0a1462633d9db7ed78e951f79913fdc8231aa77ec12b12251624de64208f4a74a07351895ddf373057b98fae6dfaf2cd21f37a063e19601078fe470d531880d0dbc3f4022080d0dbc3f4020a4b0a147b343e041ca130000a8bc00c35152bd7e774003712251624de64204a5d4753eb79f92e80efe22df7aca4f666a4f44bf81c536c4a09d4b9c5b654b51880d0dbc3f4022080d0dbc3f4020a4b0a1491844d296bd8e591448efc65fd6ad51a888d58fa12251624de6420c80e9abef7ff439c10c68fe8f1303deddfc527718c3b37d8ba6807446e3c827a1880d0dbc3f4022080d0dbc3f4020a4b0a14b3727172ce6473bc780298a2d66c12f1a14f5b2a12251624de64209142afcc691b7cc05d26c7b0be0c8b46418294171730e079f384fde2fa50bafc1880d0dbc3f4022080d0dbc3f4020a4b0a14b6f20c7faa2b2f6f24518fa02b71cb5f4a09fba312251624de642049b288e4ebbb3a281c2d546fc30253d5baf08993b6e5d295fb787a5b314a298e1880d0dbc3f4022080d0dbc3f4020a4b0a14e0dd72609cc106210d1aa13936cb67b93a0aee2112251624de642004224339688f012e649de48e241880092eaa8f6aa0f4f14bfcf9e0c76917c0b61880d0dbc3f4022080d0dbc3f4020a4b0a14fc3108dc3814888f4187452182bc1baf83b71bc912251624de64204034b37ceda8a0bf13b1abaeee7a8f9383542099a554d219b93d0ce69e3970e81880d0dbc3f4022080d0dbc3f402124f0a1406fd60078eb4c2356137dd50036597db267cf61612251624de6420e17cbe9c20cdcfdf876b3b12978d3264a007fcaaa71c4cdb701d9ebc0323f44f1880d0dbc3f4022080e0ebdaf2e2ffffff01")
	require.NoError(t, err)

	syncedConsensusStateBytes = syncedConsensusStateBytes[32:]
	parameterInput = make([]byte, 32+len(syncedConsensusStateBytes)+len(headerBytes))
	binary.BigEndian.PutUint64(parameterInput[24:32], uint64(len(syncedConsensusStateBytes)))
	copy(parameterInput[32:32+len(syncedConsensusStateBytes)], syncedConsensusStateBytes)
	copy(parameterInput[32+len(syncedConsensusStateBytes):], headerBytes)

	totalLengthPrefix = make([]byte, 32)
	binary.BigEndian.PutUint64(totalLengthPrefix[24:32], uint64(len(parameterInput)))
	input = append(totalLengthPrefix, parameterInput...)

	syncedAgainConsensusStateBytes, err := tmHeaderValidateContract.Run(input)
	require.NoError(t, err)
	syncedAgainConsensusState, err := lightclient.DecodeConsensusState(syncedAgainConsensusStateBytes[32:])
	require.NoError(t, err)

	newAppHash := syncedAgainConsensusState.AppHash
	key, err := hex.DecodeString("6163636f756e743a8a4e2eb018bdf98a8f53ec755740ffc728637a1d")
	require.NoError(t, err)
	value, err := hex.DecodeString("4bdc4c270a750a148a4e2eb018bdf98a8f53ec755740ffc728637a1d12110a0941544348412d3733301080f69bf321120b0a03424e4210e8baeb8d44120f0a075050432d303041108094ebdc031a26eb5ae98721031c199c92e5b0080967da99be27cf2da53317441b4a663e6d9c6caf02be1fdbdc20d7962b2815")
	require.NoError(t, err)
	proofBytes, err := hex.DecodeString("0add070a066961766c3a76121c6163636f756e743a8a4e2eb018bdf98a8f53ec755740ffc728637a1d1ab407b2070aaf070a2d081810cdfd2b188096a82222209f223f804e2d94ac51c4321b0687397012e6d95eb9783b03bc790da631004c7c0a2d081710adb31a18f395a8222a20d2a38865de82383ccce0140513b65cec1bf2ae6cd7dfeb22eb6faadb4e26b26f0a2d081510b2990b18f395a82222208a02bbd5a695dfc772627ac8744aa9cf30ae26575bdce8c96a9a0d0999175b430a2d081410e6ff0418f395a8222a20d39619c779be909e67f23499fb74eb2c19afd7f21523401d4ccf7e917db5cd600a2d081210e3fe0118f395a8222a20a10cc73843f889d9e03a463eb135e928bb980e19734344cba0fbf4e8a4c5258b0a2c081010dd6518f395a8222a2007fd15843a2fd3f58d021b0e072a6c70742d7a3d993a922445e3491e1c14ee8e0a2c080f10cc2a18eda6a7222a20088942d7b30abd021d8e9505cc41313fad87c8c10a799f3b51018b7b2cfe4ad90a2c080d10b70d18eda6a7222a2091a37bc44d0c61e3752ddc59eb390355ab65e8a9fb453be4f0acec537f1ca14f0a2c080c10890818eda6a72222201cfc317855a06667c45812fe36efe33af05671dfe0d9b56b02662011af2e79e30a2c080b10ac0318c4b0ee212220aeb454a4b3243b6269a2fd8841dca9a951c53b30f1e27da91063dae7224402c70a2c080910e40118c4b0ee212a20441340a4de6498f861b97b3f3ad9603af055e5af51a0d96fff2ae28e3c5c6c9a0a2c0808108d0118c4b0ee212220ae32ea4b9ab7b53571da320e2815fd8b2c278124961cca4a1849a799842424450a2b0807104d18c4b0ee212220e2804c9b7f045ec0b4ab20920a937b82fda8b7a9ddd12b21637335b915cfda550a2b0806102418a5f4c7192a20ec85f22addedfc82c771af5b4c77544b7c1d7c5bbac33f2712dfba1045ebdbd00a2b0805101118a5f4c7192a2071ade34dcc447a0ba8adc603080633d15c06f3525830c86ebce35eca0a4921fc0a2b0804100c18a5f4c7192a205190bce93993e65b266a3417ed511df8897a812cb4b62569e5afcfbec10b69cd0a2b0803100618a5f4c7192220b76c6884f1d412ac10bfb3987fb7d26f0330b2a85539509ebc5c6bdec2f95d520a2b0802100418a5f4c71922206a285b4a4f9d1c687bbafa1f3649b6a6e32b1a85dd0402421210683e846cf0020a2b0801100218a5f4c7192220033b3f7c6dcb258b6e55545e7a4f51539447cd595eb8a2e373ba0015502da1051a450a1c6163636f756e743a8a4e2eb018bdf98a8f53ec755740ffc728637a1d12201a272295e94cf1d8090bdb019dde48e9dab026ad2c3e43aaa7e61cc954a9245d18a5f4c7190ab6040a0a6d756c746973746f726512036163631aa204a0040a9d040a300a0364657812290a27088496a822122038fc49f49648fec62acc434151a51eaa378c1b20a730a749548e36f1529422500a300a03676f7612290a27088496a8221220a78ce489bdf08b9ee869c184876e1623dc38b3e64a5cf1a0005f97976c64deac0a380a0b61746f6d69635f7377617012290a27088496a8221220544c2fa38f61e10a39ec00b3e724d5834761268bb455cdbf5843bcf1531f8fbc0a300a0376616c12290a27088496a82212201f71082c9f6f45fb456b2c00b41e50d2f662f2dfec3cb6965f19d214bf02f3980a0f0a046d61696e12070a05088496a8220a320a057374616b6512290a27088496a82212200dd467343c718f240e50b4feac42970fc8c1c69a018be955f9c27913ac1f8b3c0a300a0361636312290a27088496a8221220270c19ccc9c40c5176b3dfbd8af734c97a307e0dbd8df9e286dcd5d709f973ed0a330a06746f6b656e7312290a27088496a8221220c4f96eedf50c83964de9df013afec2e545012d92528b643a5166c828774187b60a320a05706169727312290a27088496a8221220351c55cfda84596ecd22ebc77013662aba97f81f19d9ef3d150213bb07c823060a360a0974696d655f6c6f636b12290a27088496a8221220e7adf5bd30ce022decf0e9341bf05c464ed70cdbc97423bd2bab8f3571e5179b0a330a06706172616d7312290a27088496a822122042a9dfc356ca435db131eb41fb1975c8482f2434537918665e530b0b4633b5f9")
	require.NoError(t, err)

	merkleProofInput := make([]byte, 32+32+len(key)+32+len(value)+32+len(proofBytes))
	copy(merkleProofInput[:32], "acc")
	binary.BigEndian.PutUint64(merkleProofInput[32+24:32+32], uint64(len(key)))
	copy(merkleProofInput[32+32:32+32+len(key)], key)

	binary.BigEndian.PutUint64(merkleProofInput[32+32+len(key)+24:32+32+len(key)+32], uint64(len(value)))
	copy(merkleProofInput[32+32+len(key)+32:32+32+len(key)+32+len(value)], value)

	copy(merkleProofInput[32+32+len(key)+32+len(value):32+32+len(key)+32+len(value)+32], newAppHash)
	copy(merkleProofInput[32+32+len(key)+32+len(value)+32:], proofBytes)

	totalLengthPrefix = make([]byte, 32)
	binary.BigEndian.PutUint64(totalLengthPrefix[0:8], 0)
	binary.BigEndian.PutUint64(totalLengthPrefix[8:16], 0)
	binary.BigEndian.PutUint64(totalLengthPrefix[16:24], 0)
	binary.BigEndian.PutUint64(totalLengthPrefix[24:], uint64(len(merkleProofInput)))
	input = append(totalLengthPrefix, merkleProofInput...)

	iavlMerkleProofValidateContract := iavlMerkleProofValidate{}
	success, err := iavlMerkleProofValidateContract.Run(input)
	require.NoError(t, err)
	expectedResult := make([]byte, 32)
	binary.BigEndian.PutUint64(expectedResult[24:], 0x01)
	require.Equal(t, expectedResult, success)
}
