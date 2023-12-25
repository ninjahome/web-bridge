package util

import (
	"image/jpeg"
	"os"
	"testing"
)

var testContent = `罗马帝国是历史上的一个文明帝国，承接着先前的罗马共和国。中国史书称为“大秦”或“海西国”。

前44年，罗马共和国将领凯撒成为终身独裁官，象征着共和制的结束；至前27年，屋大维成为奥古斯都，象征着罗马帝国的开端。其首都罗马城在公元前100年至公元400年这段时期是欧洲最大的城市，直至公元500年君士坦丁堡取代罗马城成为欧洲最大城市[5][6]，帝国人口亦增长到五千至九千万，大约是当时世界总人口的约20%[7][8]

罗马帝国可分为前期（前27年—200年）、中期（200年—395年）和后期〔395年—1204年[注 2]/1453年[注 3]〕三个阶段。

罗马共和国末年，政局由于一连串的内战和政治角力变得非常不稳定。公元前44年，共和国将领凯撒被元老院封为终身独裁官后不久，便遭到刺杀身亡。直至公元前31年，屋大维在亚克兴角战役击败对手马克·安东尼和女王克娄巴特拉七世，吞并埃及托勒密王国后，共和国的政局仍然不明朗。至公元前27年，元老院放弃共和制，赐君权及奥古斯都头衔予屋大维，这象征着罗马共和国的终结。这时元老院仍然存在，但大权已掌握在屋大维手中[9]。最初几位皇帝都以“第一公民”自居。

屋大维征战的胜利扩张了帝国的领土。立国之初的两个世纪，帝国的政局有着前所未见的稳定，这段时期被称为“罗马治世”。直到公元41年，帝国第三位皇帝卡里古拉被刺身亡后，元老院曾经考虑恢复共和政制，但罗马禁卫军架空元老院，遂立克劳狄一世为帝。克劳狄一世在位期间，帝国继屋大维后首次征战不列颠尼亚。公元68年，克劳狄一世的继位者尼禄在兵变中自杀身亡，帝国遭遇一连串短暂的内战，同时犹太地区更爆发第一次起义，这段时期曾经有四位军团将领称帝。维斯帕先在公元69年战胜其他将领，建立弗拉维王朝。其继位者提图斯，在公元79年维苏威火山爆发后开放斗兽场。提图斯只短暂在位两年，便由其兄弟图密善继位为帝国第11位皇帝。图密善最后亦遭到刺杀身亡。元老院后来封涅尔瓦为皇帝，这亦是罗马五贤帝之首，开辟一段长达八十多年的政局稳定时期。罗马第13位皇帝图拉真是罗马五贤帝之中的第二位，他在位时见证帝国的最大版图。

康茂德在位时，帝国开始出现衰退之兆。公元192年康茂德被刺杀身亡，触发五帝之年，有五人同时称帝，分别是佩蒂纳克斯、尤利安努斯、奈哲尔、阿尔拜努斯和塞普蒂米乌斯·塞维鲁，乱象最后由塞普蒂米乌斯·塞维鲁取得胜利而告终。公元235年皇帝亚历山大·塞维鲁被刺杀身亡，导致三世纪危机，这段时期短短50年之内有26人被元老院封为皇帝。直至戴克里先在位时创立四帝共治，帝国才全面恢复稳定，这段时期一共有四位皇帝共同统治罗马帝国。这种制度并不能维持下去，很快便招致内战。内战最终由君士坦丁一世胜出，成为帝国的唯一统治者。后来君士坦丁一世迁都至拜占庭，并命名为新罗马，但史家更喜欢以其名字―君士坦丁称之为君士坦丁堡。君士坦丁堡自此之后一直是帝国的首都，一直到其终结。君士坦丁一世亦于313年将基督公教（中文译为天主教会）合法化，并由狄奥西亚一世将基督教定为国教，基督教从而成为西方世界的主要宗教。

这时的罗马帝国仍然是世界上的强权，并与其东面安息帝国的继承者波斯第二帝国互相抗衡，持续着一个世纪的纷争[10][11]。狄奥多西一世是最后一位统治一个完整的罗马帝国的皇帝，随后帝国的领土因滥权、内战、野蛮人入侵、军事改革和经济衰退等负面因素被日益蚕食，这时的罗马帝国实际上已完全分裂成东西两部分，自此之后再没有被统一过。公元410年及公元455年，西面的罗马城相继被西哥德人和汪达尔人等日耳曼部族入侵。公元476年，西罗马帝国皇帝罗慕路斯·奥古斯都被奥多亚塞废黜，这象征着西罗马帝国的终结。但由于罗慕路斯·奥古斯都从未被东罗马帝国所承认，所以严格上来说，西罗马帝国上一位皇帝尼波斯在公元480年去世后才算是罗马帝国在西欧的统治划上句号。而东罗马帝国则一直存在至1453年，君士坦丁堡被土耳其人攻陷、皇帝君士坦丁十一世战死为止。史学家通称东罗马帝国为拜占庭帝国。

罗马帝国是世界历史上一个伟大的帝国，无论经济、文化、政治和军事上的成就都达到很高的水准，并和在与公元前一世纪兴起于亚洲的汉帝国西、东遥相并立。后世多将两者并列为当时世界上最先进及文明的强大帝国[12]。整个罗马帝国(包括东西罗马帝国)存在将近一千五百年，帝国的疆域在图拉真在位末年（117年）达到全盛，控制着大约五百万平方公里的土地[2]，统治着七千万的人口，这相当于当时世界总人口的百分之二十一。罗马帝国幅原辽阔，而且国祚长久，使拉丁希腊的语言、文化、宗教、发明、建筑、哲学、法律及政府模式对后世的影响相当深远。欧洲在整个中世纪时期，有数次对罗马帝国的复辟，这包括神圣罗马帝国。文艺复兴而后的欧洲帝国主义的兴起，更使希腊、罗马、犹太和基督教的文化向全世界传播开去，对现代社会文明的发展有着重要影响。`

var testContent2 = `Visitors sunbathe at La Marbella beach ( Image: Bloomberg via Getty Images)
He thinks this means employees will be more open to collaboration, and more motivated generally, which is essential for the success of the business. Each office is assigned a weekly team happiness tool, to check how staff are feeling about work.

Their workplace environments are designed to make sure staff feel creatively inspired and fulfilled at work. As a mid-tier financial advisory firm, the company is quickly challenging the "big four" - Deloitte, Ernst & Young, KPMG and PwC - having secured plenty of new clients.

It's representing organisations such as Rugby League superpowers Leeds Rhinos and fast growing retailers such as the beauty and cosmetic brand P Louise, and the rapidly growing sportswear brand, Montirex. Paul said: “It’s been another great year for Sedulo, which simply could not have been possible without the hard work and dedication shown by our team year after year.

"We started the Sedulo Christmas trip over 12 years ago now and it’s something we’ll continue to do, despite the rapid pace at which the business is growing. These trips are about giving something back and showing our people that their hard work is appreciated, particularly after completion of our 2023 record breaking Christmas Toy Appeal, while also allowing us to come together as a team and celebrate the continued success of our business.`
var testContent3 = `最近，Solana带领Avalanche、NEAR等layer1公链大杀四方，市场又传出杀死以太坊的声音。 确实，ETH Killer是上一轮牛市大部分公链手握的王牌叙事。
不过，在EVM一统江湖，layer2尚未爆发的当下，单靠MEME、DePin等叙事不足以撼动ETH的市场地位。这只是坎昆升级前，Alt-Layer1的短暂狂欢，Why？
1）各大公链争相杀死以太坊的叙事，已然被一轮牛熊周期充分验证“失败了”。表面上看这源于以太坊强大的市场共识，和开发者群体创新力量，以及DeFi、NFT金融应用无限组合的魔力。
实际上因为，Crypto市场还受限于技术、市场、合规等因素并未实现Mass Adoption大规模普及，这些新公链带来的技术“跃迁”并没有成为激发新叙事，扩大新市场增量的抓手，还仅仅在吃以太坊的溢出红利。
Solana、Avalanche、Aptos等公链想从开发语言、代码复杂度、运行机制等底层框架彻底提升技术水平，为应用市场提供更好的infra。比如，Solana的高并发处理性能和用户体验UX优势，单论技术确实更适用于未来增量Crypto市场。
只不过，眼下市场基本运作逻辑还没真正蜕变。
偏好风险的游资，渴求财富密码的市场受众，不断刷新的多样玩法，始终存在的信息差，偶尔溢出圈的暴富故事等等，这些完美构成了周期性的牛市基本要素。
这让技术“先天局限”的以太坊，愣是靠各类EIP、ERC标准协议等缝缝补补，也足以衍生出一个庞大的应用市场，还能让其他竞品公链仅靠溢出效应就能频频冒头。
但大家都在吃以太坊DeFi市场红利，还没到Alt-layer1 可以轻松取代并超越以太坊的时候。
2）以太坊“先天缺陷”已经探索出了一整套成熟解决方案，比如：扩容问题，演变出了Rollup、Plasma、Validium等各类方案；又比如：EOA地址局限，靠ERC4337 Account Abstraction也得到了升级，甚至演变出了一个账户抽象赛道；
此外，layer2也同样成了一个叙事赛道，OP-Rollup和ZK-Rollup展开持续拉锯战等；后续还有坎昆升级后的Blob空间以及更遥远的Sharding分片、底层SNARK化等来提供后续发展支撑；
即便是区块容量上限潜在的DA能力局限问题，也延伸出了Eigenlayer等基于Restaking方案来优化DA，再通过模块化组合Celestia这类三方DA方案，以及对VM执行层的可选替代方案等等。
整个以太坊的开发、扩容、外延环境已经足够成熟和繁荣。它背后的开发者力量才是大以太屹立不倒的基石。
虽然以太坊过去几年，持续叠乐高构建生态的落地结果确实不及预期，但能在频发的黑客攻击下，完成POW到POS的关键升级，又能把开发者资源聚拢在以以太坊EVM为中心的主线上，且能演变出一个更加宏大的layer2叙事板块，以太坊的后续潜在可能性不容低估。
相信以太坊，是对以太坊多年稳固共识的敬畏，也是对其背后庞大开发者群体Builder的Respect。
依稀记得18年岁末，EOS号称全新范式公链，且掀起了一轮菠菜游戏狂潮，但事实结果大家都看到了，短暂繁华散尽后，最终还是慢却稳健的以太坊笑到了最后。
真正的价值发现，一定得慢慢Capture捕捉。
3）Layer2在熊市的Build速度确实慢了，尤其是没有一轮layer2 Summer的市场大馈赠，让参与layer2生态构建的每个人都有点心有不甘。
不过， layer2 Build速度过慢和以太坊DeFi叙事溢出到各大新锐公链另立山头道理类似。 以太坊layer2的后半程要靠一些高频交易和应用驱动，单靠以太坊金融玩法的溢出效应和路径依赖，和Alt-layer1硬碰硬对抗并没有优势。
一方面，Arbitrum、Optimism等OP-Rollup有了layer2先发生态优势，并在Stack战略下，扩大了市场疆域，但这些战略扩张到底属于B端布局，OP-Rollup需要解决被诟病的中心化难题且带动C端市场增长。
另一方面，zkSync、Starknet等ZK-Rollup有更Advanced技术优势，只不过ZK也是着眼于未来的技术，已有用户量级无法充分显现ZK的强大之处。只有用户量级扩大，Gas才能低到忽略不计，体验也越丝滑，这才是ZK layer2的最终形态。
此外，layer2市场腰尾部力量正在举事，比如，Metis，尝试用Hyper（OP+ZK）Rollup技术，做POS去中心化Sequencer，变革Token的激励方式（治理—>实用）等等。还有，Espresso、Astria等共享Sequencer方案，也在以Rollup as a Service的方式也不断外延layer2市场的潜能。
别以为OP+ZK就已经把layer2的故事讲完了，在我看来，Layer2 War才真正开始，真正内卷的layer2市场坎昆升级后可能才拉开局面。当坎昆升级时间敲定后，layer2的抱团逆势上涨不正是对眼下layer2憋屈处境的一次情绪释放？
当未来应用链叙事场景打开，Mass Adoption的局面被打开后，layer2赛道能沉淀的资金、用户和DApp应用，一定比其他Alt-layer1更稳固。
4）当然，此刻为以太坊生态发声，并非否定Solana的市场潜力。不可否认，Solana的技术创新起点确实高于已有区块链架构，无论是它的存储和计算分离特性以及高并发交易处理特性，都使得它对用户很友好，更容易打造生态。
以DePin为例，物理基础设施+Token代币激励，这是过去一直在失败漩涡中反复的叙事，比如Filecoin Arweave等。在Solana上会不会真正成功，我不知道，但DePin发生在Solana上，我对DePin的信心会多一些。毕竟，高并发的技术起点和web2天然接轨，这和靠模块组合并入的生态逻辑还不一样。
Solana眼下的崛起，一方面是以太坊layer2短暂沉寂带来的空窗机会，另一方面则是Solana上本来就活跃着一批开发力量的结果。但需要纠正的是，Solana的目标并非杀死以太坊，它其实在找以太坊的“空白”点伺机突破，如果有所建树且有于以太坊匹配的生态当量，顶多也就是比肩而已，何来取代一说。
以太坊不可避免的会被一些全新技术起点的Alt-layer1链冲击，但他们都不是以太坊“杀手”，我更倾向于称之为，web3破局创新者。
以太坊在DeFi金融应用以及庞大的组合性生态已经获得了成功，而layer2、layer3未来的新征途还正在路上。
如果“开放、包容、可信、组合”，这样的以太坊最终都成就不了区块链价值落地，我很难相信一个新链就会。
Note：若认同我的思考，喜欢我文章的话，帮忙一键三连下，谢谢大家支持。`

func TestTxt2Img(t *testing.T) {
	var img, err = ConvertLongTweetToImg(testContent3)
	if err != nil {
		t.Fatal(err)
	}
	// 保存图像
	outFile, err := os.Create("output.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer outFile.Close()
	err = jpeg.Encode(outFile, img, nil)
	if err != nil {
		t.Fatal(err)
	}
}
