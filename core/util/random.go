package util

import (
	"strings"
	"math/rand"
	"time"
)

var names="刀白凤 丁春秋 马夫人 马五德 小翠 于光豪 巴天石 不平道人 邓百川 风波恶 甘宝宝 公冶乾 木婉清 包不同 天狼子 太皇太后 王语嫣 乌老大 无崖子 云岛主 云中鹤 止清 白世镜 天山童姥 本参 本观 本相 本因 出尘子 冯阿三 古笃诚 少林老僧 过彦之 兰剑 平婆婆 石清露 石嫂 司空玄 司马林 玄慈 玄寂 玄苦 玄难 玄生 玄痛 叶二娘 左子穆 耶律莫哥 李春来 李傀儡 李秋水 刘竹庄 祁六三 乔峰 全冠清 朴者和尚 阮星竹 许卓诚 朱丹臣 竹剑 阿碧 阿洪 阿胜 西夏宫女 阿朱 阿紫 波罗星 陈孤雁 何望海 鸠摩智 来福儿 耶律洪基 努儿海 宋长老 苏星河 苏辙 完颜阿古打 吴长风 枯荣长老 辛双清 严妈妈 余婆婆 岳老三 张全祥 单伯山 单季山 单叔山 单小山 单正 段延庆 段誉 段正淳 段正明 范禹 宗赞王子 范骅 苟读 和里布 孟师叔 华赫艮 耶律涅鲁古 耶律重元 郁光标 卓不凡 范百龄 哈大霸 姜师叔 吴光胜 梦姑 神山上人 神音 狮鼻子 室里 项长老 姚伯当 幽草 赵钱孙 黄眉和尚 哲罗星 钟灵 钟万仇 高升泰 龚光杰 贾老者 康广陵 秦红棉 容子矩 桑土公 唐光雄 奚长老 徐长老 诸保昆 崔百泉 崔绿华 符敏仪 赵洵 菊剑 梅剑 萧远山 虚竹 游骥 聋哑婆婆 游驹 游坦之 程青霜 傅思归 葛光佩 缘根 鲍千灵 智光大师 褚万里 瑞婆婆 端木元 黎夫人 慕容博 慕容复 谭公 赫连铁树 谭婆 谭青 摘星子 慧方 慧观 慧净 慧真 穆贵妃 薛慕华 吴领军 易大彪 卜沉 丁坚 丁勉 上官云 万大平 于人豪 于嫂 不戒和尚 长青子 仇松年 丹青生 邓八公 方人智 方生 方证 天门道人 计无施 木高峰 风清扬 丛不弃 王伯奋 王诚 王二叔 天乙道人 王夫人 王家驹 王家骏 王元霸 王仲强 白二 白熊 天松道人 乐厚 令狐冲 宁中则 平夫人 平一指 申人俊 史镖头 东方不败 史登达 司马大 田伯光 仪和 仪琳 仪清 玉玑子 玉灵道人 玉磬子 玉音子 玉钟子 左冷禅 成不忧 齐堂主 吉人通 冲虚道长 老不死 老头子 刘菁 刘芹 刘正风 米为义 农妇 成高道人 曲非烟 曲洋 任我行 英颚 西宝 向大年 向问天 陈七 陈歪嘴 迟百诚 狄镖头 狄修 杜长老 何三七 季镖头 定静师太 劳德诺 陆伯 陆大有 任盈盈 沙天江 秃笔翁 吴柏英 吴天德 辛国梁 严三星 杨莲亭 余沧海 余人彦 岳灵珊 张夫人 张金鏊 定逸 建除 林平之 林远图 林震南 罗人杰 易国梓 易师爷 易堂主 英白罗 英长老 岳不群 郑镖头 郑萼 周孤桐 费彬 封不平 洪人雄 侯人英 觉月 施戴子 施令威 闻先生 哑婆婆 钟镇 祝镖头 祖千秋 高克新 高明根 贾布 贾人达 夏老拳师 秦娟 秦伟帮 桑三娘 桃干仙 桃根仙 桃花仙 桃实仙 桃叶仙 桃枝仙 陶钧 莫大 崔镖头 黄伯流 黄国柏 黄钟公 梁发 绿竹翁 游迅 葛长老 震山子 黑白子 黑熊 麻衣汉子 鲁连荣 舒奇 童百熊 鲍大楚 解风 蓝凤凰 谭迪人 清虚道人 九难 卫周祚 马喇 马佑 马宝 马博仁 于八 马超兴 小桂子 小玄子 马齐 心溪 韦小宝 韦春花 毛文珠 巴泰 方怡 风际中 邓炳春 云素梅 王潭 无根道人 五符 元义方 巴郎星 王武通 王进宝 王琪 双儿 史松 冯难敌 邝天雄 平威 白寒松 白寒枫 卢一峰 归辛树 玄真道人 司徒鹤 对喀纳 冯锡范 孙思克 归钟 归二娘 玉林 司徒伯雷 汤若望 李自成 老吴 守备 米思翰 江百胜 齐元凯 华伯斯基 西奥图三世 刘一舟 沐剑声 庄夫人 许雪亭 多隆 齐洛诺夫 祁清彪 关安基 吕留良 陈珂 李西华 吕葆中 吕毅中 行颠 庄廷龙 庄允城 陆高轩 杜立德 吴之荣 苏菲亚 陈圆圆 罕贴摩 吴大鹏 沐剑屏 吴三桂 阿济赤 阿尔尼 张淡月 苏荃 苏冈 吴六奇 李式开 李力世 陈近南 吴应熊 杨溢之 佟国纲 吴立身 张康年 张勇 张妈 吴宝宇 何惕守 劳太监 明珠 费要多罗 柳燕 图海道 杰书 郎师傅 净清 净济 林兴珠 图尔布青 林永超 柳大洪 呼巴音 昌齐 郑克爽 赵齐贤 茅十八 建宁公主 洪朝 姚春 施琅 皇甫阁 胡逸之 南怀仁 钟志灵 神照上人 洪安通 胡德第 姚必达 赵良栋 查继左 胖头陀 郝太监 徐天川 陶红英 索额图 教士 陶师傅 高里津 敖彪 高颜超 钱老本 海大富 殷锦 贾老六 笔贴式 顾炎武 夏国相 桑结 晦聪禅师 章老三 黄甫 黄金魁 崔瞎子 黄宗羲 菊芳 彭参将 葛尔丹 程维藩 温有方 温有道 舒化龙 曾柔 富春 葛通 路副将 雷一啸 瘦头陀 蕊初 瑞栋 蔡德忠 察尔珠 潘先生 澄光 澄通 澄观 澄心 澄识 樊纲 慕天颜 鳌拜 巴颜法师 行痴 子聪 丁大全 人厨子 九死生 马钰 小棒头 大头鬼 一灯大师 马光佐 小龙女 尹志平 丘处机 王处一 王十三 公孙止 小王将军 王志坦 王惟忠 无常鬼 尹克西 天竺僧 少妇 孙婆婆 公孙绿萼 孙不二 皮清云 申志凡 冯默风 讨债鬼 史伯威 史仲猛 史叔刚 史季强 史孟龙 尼摩星 李莫愁 达尔巴 刘处玄 朱子柳 圣因师太 曲傻姑 吕文德 祁志诚 李志常 刘瑛姑 吊死鬼 百草仙 陆鼎立 陆二娘 阿根 张志光 完颜萍 陆冠英 宋德方 陈大方 觉远大师 沙通天 张君宝 张一氓 陈老丐 张二叔 陆无双 杨过 灵智上人 武三通 武敦儒 武修文 武三娘 林朝英 耶律晋 耶律燕 耶律楚材 忽必烈 丧门鬼 狗头陀 青灵子 欧阳峰 耶律齐 周伯通 金轮法王 洪凌波 柔儿 郭破虏 侯通海 宋五 俏鬼 柯镇恶 点苍渔隐 赵志敬 洪七公 郭靖 郭芙 郭襄 姬清玄 笑脸鬼 鹿清笃 崔志方 鄂尔多 萨多 黄药师 黄蓉 程遥迦 鲁有脚 彭连虎 韩无垢 童大海 韩老丐 彭长老 瘦丐 程瑛 雷猛 裘千尺 蒙哥 煞神鬼 催命鬼 蓝天和 裘千仞 赫大通 潇湘子 霍都 樊一翁 藏边大丑 藏边二丑 藏边三丑 藏边四丑 藏边五丑"

func RandomPickName()string  {
	array:=strings.Split(names," ")
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	i:=r.Intn(len(array)-1)
	return array[i]
}
