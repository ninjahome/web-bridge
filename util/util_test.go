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

func TestTxt2Img(t *testing.T) {
	var img, err = ConvertLongTweetToImg(testContent)
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
