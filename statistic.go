package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"texas/statistic"
)

type winCountData struct {
	winCount   float64
	totalCount float64
}

type Statistic struct {
	data map[int32]map[int32]*winCountData
}

func NewStatistic() *Statistic {
	data := make(map[int32]map[int32]*winCountData)
	for i := int32(2); i <= 14; i++ {
		m := make(map[int32]*winCountData)
		for j := int32(2); j <= 14; j++ {
			m[j] = &winCountData{0, 0}
		}
		data[i] = m
	}
	return &Statistic{data}
}

func (s *Statistic) Add(card []*Card, isWin bool, winnerCount int) {
	if len(card) != 2 {
		panic(card)
	}
	var a, b int32
	if card[0].Num < card[1].Num {
		a = card[0].Num
		b = card[1].Num
	} else {
		a = card[1].Num
		b = card[0].Num
	}
	if card[0].Color != card[1].Color {
		a, b = b, a
	}
	v := s.data[a][b]
	v.totalCount++
	if isWin {
		v.winCount += 1.0 / float64(winnerCount)
	}
}

func (s *Statistic) Get(card []*Card) float64 {
	if len(card) != 2 {
		panic(card)
	}
	var a, b int32
	if card[0].Num < card[1].Num {
		a = card[0].Num
		b = card[1].Num
	} else {
		a = card[1].Num
		b = card[0].Num
	}
	if card[0].Color != card[1].Color {
		a, b = b, a
	}
	v := s.data[a][b]
	return v.winCount / v.totalCount
}

//func (s *Statistic) Save() {
//	stat := &statistic.StatisticData{}
//	for i := int32(14); i >= 2; i-- {
//		m := s.data[i]
//		for j := int32(14); j >= 2; j-- {
//			stat.Data = append(stat.Data, &statistic.CountData{
//				A:          i,
//				B:          j,
//				WinCount:   m[j].winCount,
//				TotalCount: m[j].totalCount,
//			})
//		}
//	}
//	buf, err := proto.Marshal(stat)
//	if err != nil {
//		panic(err)
//	}
//	err = ioutil.WriteFile("statistic.dat", buf, 0644)
//	if err != nil {
//		panic(err)
//	}
//}

func (s *Statistic) Load() {
	// 这是一个预先算好的概率表
	js := []byte("{\"data\":[{\"a\":14,\"b\":14,\"win_count\":1155963.6702378977,\"total_count\":3237772},{\"a\":14,\"b\":13,\"win_count\":1356926.254761887,\"total_count\":6480851},{\"a\":14,\"b\":12,\"win_count\":1230294.9190474427,\"total_count\":6478334},{\"a\":14,\"b\":11,\"win_count\":1137266.6595235793,\"total_count\":6474635},{\"a\":14,\"b\":10,\"win_count\":1066474.2738092903,\"total_count\":6476278},{\"a\":14,\"b\":9,\"win_count\":917942.508333109,\"total_count\":6480892},{\"a\":14,\"b\":8,\"win_count\":875077.1285712493,\"total_count\":6477869},{\"a\":14,\"b\":7,\"win_count\":835169.5857141144,\"total_count\":6479375},{\"a\":14,\"b\":6,\"win_count\":798735.5488093703,\"total_count\":6477146},{\"a\":14,\"b\":5,\"win_count\":836673.7773807705,\"total_count\":6477933},{\"a\":14,\"b\":4,\"win_count\":811618.6511902668,\"total_count\":6476994},{\"a\":14,\"b\":3,\"win_count\":786997.5440473979,\"total_count\":6477538},{\"a\":14,\"b\":2,\"win_count\":757549.6380950137,\"total_count\":6477739},{\"a\":13,\"b\":14,\"win_count\":896848.2119045414,\"total_count\":3238527},{\"a\":13,\"b\":13,\"win_count\":942697.945237863,\"total_count\":3239155},{\"a\":13,\"b\":12,\"win_count\":1170693.8988093196,\"total_count\":6481771},{\"a\":13,\"b\":11,\"win_count\":1078813.5226188067,\"total_count\":6477564},{\"a\":13,\"b\":10,\"win_count\":1010853.4761902686,\"total_count\":6483722},{\"a\":13,\"b\":9,\"win_count\":856927.9547617204,\"total_count\":6477597},{\"a\":13,\"b\":8,\"win_count\":759979.1369046267,\"total_count\":6481358},{\"a\":13,\"b\":7,\"win_count\":731405.9416665612,\"total_count\":6478874},{\"a\":13,\"b\":6,\"win_count\":703432.4249998918,\"total_count\":6481943},{\"a\":13,\"b\":5,\"win_count\":678085.959523697,\"total_count\":6477441},{\"a\":13,\"b\":4,\"win_count\":653923.2797618064,\"total_count\":6477213},{\"a\":13,\"b\":3,\"win_count\":636652.6226189513,\"total_count\":6475258},{\"a\":13,\"b\":2,\"win_count\":619188.2202380039,\"total_count\":6476745},{\"a\":12,\"b\":14,\"win_count\":844825.6166665013,\"total_count\":3239636},{\"a\":12,\"b\":13,\"win_count\":783922.2428570034,\"total_count\":3237692},{\"a\":12,\"b\":12,\"win_count\":783214.7738093869,\"total_count\":3241383},{\"a\":12,\"b\":11,\"win_count\":1034795.8738092977,\"total_count\":6478196},{\"a\":12,\"b\":10,\"win_count\":966649.8309522142,\"total_count\":6475003},{\"a\":12,\"b\":9,\"win_count\":813416.5785712709,\"total_count\":6477520},{\"a\":12,\"b\":8,\"win_count\":713940.2428570306,\"total_count\":6477859},{\"a\":12,\"b\":7,\"win_count\":633205.4119046764,\"total_count\":6475045},{\"a\":12,\"b\":6,\"win_count\":612687.9726189643,\"total_count\":6475097},{\"a\":12,\"b\":5,\"win_count\":590787.4404761097,\"total_count\":6476445},{\"a\":12,\"b\":4,\"win_count\":570009.5059523052,\"total_count\":6478785},{\"a\":12,\"b\":3,\"win_count\":551119.3023808999,\"total_count\":6478951},{\"a\":12,\"b\":2,\"win_count\":535045.8678571263,\"total_count\":6477348},{\"a\":11,\"b\":14,\"win_count\":800708.2166665458,\"total_count\":3239865},{\"a\":11,\"b\":13,\"win_count\":748583.6702379794,\"total_count\":3235194},{\"a\":11,\"b\":12,\"win_count\":709246.9059522862,\"total_count\":3235534},{\"a\":11,\"b\":11,\"win_count\":655836.6309523217,\"total_count\":3241449},{\"a\":11,\"b\":10,\"win_count\":949373.5321427224,\"total_count\":6475365},{\"a\":11,\"b\":9,\"win_count\":799284.8499998905,\"total_count\":6475163},{\"a\":11,\"b\":8,\"win_count\":698200.3976189581,\"total_count\":6476307},{\"a\":11,\"b\":7,\"win_count\":612366.5404761166,\"total_count\":6480461},{\"a\":11,\"b\":6,\"win_count\":539318.5166665874,\"total_count\":6477625},{\"a\":11,\"b\":5,\"win_count\":527315.3797618103,\"total_count\":6478915},{\"a\":11,\"b\":4,\"win_count\":507320.7321427837,\"total_count\":6477678},{\"a\":11,\"b\":3,\"win_count\":488256.1773809112,\"total_count\":6477208},{\"a\":11,\"b\":2,\"win_count\":473324.72499998106,\"total_count\":6475338},{\"a\":10,\"b\":14,\"win_count\":769984.4488094205,\"total_count\":3239021},{\"a\":10,\"b\":13,\"win_count\":723573.5047618088,\"total_count\":3239825},{\"a\":10,\"b\":12,\"win_count\":688025.3642856353,\"total_count\":3239372},{\"a\":10,\"b\":11,\"win_count\":667940.6999999341,\"total_count\":3240593},{\"a\":10,\"b\":10,\"win_count\":559106.151190486,\"total_count\":3239535},{\"a\":10,\"b\":9,\"win_count\":803144.4726189527,\"total_count\":6475756},{\"a\":10,\"b\":8,\"win_count\":705163.8297618277,\"total_count\":6481900},{\"a\":10,\"b\":7,\"win_count\":617229.4392856415,\"total_count\":6478514},{\"a\":10,\"b\":6,\"win_count\":539686.3738094374,\"total_count\":6478101},{\"a\":10,\"b\":5,\"win_count\":474159.69880944997,\"total_count\":6474322},{\"a\":10,\"b\":4,\"win_count\":461566.1488094592,\"total_count\":6473033},{\"a\":10,\"b\":3,\"win_count\":444744.8416666279,\"total_count\":6475102},{\"a\":10,\"b\":2,\"win_count\":430025.0011904619,\"total_count\":6478707},{\"a\":9,\"b\":14,\"win_count\":705306.0666665802,\"total_count\":3241320},{\"a\":9,\"b\":13,\"win_count\":657846.8440475531,\"total_count\":3238353},{\"a\":9,\"b\":12,\"win_count\":626133.0690475749,\"total_count\":3238540},{\"a\":9,\"b\":11,\"win_count\":607915.0607142472,\"total_count\":3236386},{\"a\":9,\"b\":10,\"win_count\":602214.1285714007,\"total_count\":3241409},{\"a\":9,\"b\":9,\"win_count\":474374.8095238414,\"total_count\":3239397},{\"a\":9,\"b\":8,\"win_count\":686287.4321427825,\"total_count\":6477010},{\"a\":9,\"b\":7,\"win_count\":614456.9976189809,\"total_count\":6473287},{\"a\":9,\"b\":6,\"win_count\":537309.1023808892,\"total_count\":6477042},{\"a\":9,\"b\":5,\"win_count\":466844.6726189966,\"total_count\":6484748},{\"a\":9,\"b\":4,\"win_count\":402311.34761903563,\"total_count\":6478445},{\"a\":9,\"b\":3,\"win_count\":393318.89880951867,\"total_count\":6472673},{\"a\":9,\"b\":2,\"win_count\":378651.7666666831,\"total_count\":6475485},{\"a\":8,\"b\":14,\"win_count\":685657.0785713559,\"total_count\":3240091},{\"a\":8,\"b\":13,\"win_count\":618435.3559523404,\"total_count\":3239691},{\"a\":8,\"b\":12,\"win_count\":585933.9654761726,\"total_count\":3235660},{\"a\":8,\"b\":11,\"win_count\":570386.9535714167,\"total_count\":3241005},{\"a\":8,\"b\":10,\"win_count\":564389.2440476012,\"total_count\":3239182},{\"a\":8,\"b\":9,\"win_count\":550007.0202381006,\"total_count\":3239266},{\"a\":8,\"b\":8,\"win_count\":416492.5404762219,\"total_count\":3235590},{\"a\":8,\"b\":7,\"win_count\":625750.2773808904,\"total_count\":6483448},{\"a\":8,\"b\":6,\"win_count\":560319.5690475579,\"total_count\":6469007},{\"a\":8,\"b\":5,\"win_count\":488278.2988094753,\"total_count\":6476571},{\"a\":8,\"b\":4,\"win_count\":415602.7547618976,\"total_count\":6475229},{\"a\":8,\"b\":3,\"win_count\":357194.371428598,\"total_count\":6479821},{\"a\":8,\"b\":2,\"win_count\":349942.4166666967,\"total_count\":6478843},{\"a\":7,\"b\":14,\"win_count\":665886.7380951772,\"total_count\":3240111},{\"a\":7,\"b\":13,\"win_count\":608115.8083333012,\"total_count\":3239343},{\"a\":7,\"b\":12,\"win_count\":552580.3964285691,\"total_count\":3239752},{\"a\":7,\"b\":11,\"win_count\":533484.8619047652,\"total_count\":3235179},{\"a\":7,\"b\":10,\"win_count\":528728.123809532,\"total_count\":3236910},{\"a\":7,\"b\":9,\"win_count\":520793.6511904939,\"total_count\":3239177},{\"a\":7,\"b\":8,\"win_count\":519972.55714288045,\"total_count\":3237992},{\"a\":7,\"b\":7,\"win_count\":371376.01190479286,\"total_count\":3236554},{\"a\":7,\"b\":6,\"win_count\":577397.0499999378,\"total_count\":6475679},{\"a\":7,\"b\":5,\"win_count\":518010.23214280483,\"total_count\":6480907},{\"a\":7,\"b\":4,\"win_count\":446720.0309523686,\"total_count\":6478504},{\"a\":7,\"b\":3,\"win_count\":377807.6559524009,\"total_count\":6476395},{\"a\":7,\"b\":2,\"win_count\":321393.65833338536,\"total_count\":6476953},{\"a\":6,\"b\":14,\"win_count\":651446.0630951772,\"total_count\":3237953},{\"a\":6,\"b\":13,\"win_count\":594024.8511904465,\"total_count\":3237315},{\"a\":6,\"b\":12,\"win_count\":546518.5666666594,\"total_count\":3240255},{\"a\":6,\"b\":11,\"win_count\":502502.53928572225,\"total_count\":3234374},{\"a\":6,\"b\":10,\"win_count\":496245.488095244,\"total_count\":3238829},{\"a\":6,\"b\":9,\"win_count\":490828.36190477916,\"total_count\":3240410},{\"a\":6,\"b\":8,\"win_count\":493592.252380973,\"total_count\":3237235},{\"a\":6,\"b\":7,\"win_count\":498280.56785715895,\"total_count\":3237104},{\"a\":6,\"b\":6,\"win_count\":339507.3500000278,\"total_count\":3238757},{\"a\":6,\"b\":5,\"win_count\":549356.402380906,\"total_count\":6475293},{\"a\":6,\"b\":4,\"win_count\":489422.7940476099,\"total_count\":6474898},{\"a\":6,\"b\":3,\"win_count\":421207.2226190662,\"total_count\":6480900},{\"a\":6,\"b\":2,\"win_count\":352333.74047623033,\"total_count\":6480517},{\"a\":5,\"b\":14,\"win_count\":665370.1238094575,\"total_count\":3236857},{\"a\":5,\"b\":13,\"win_count\":584124.9976190217,\"total_count\":3236198},{\"a\":5,\"b\":12,\"win_count\":535383.9154761914,\"total_count\":3235502},{\"a\":5,\"b\":11,\"win_count\":500389.2000000021,\"total_count\":3240519},{\"a\":5,\"b\":10,\"win_count\":467966.580952389,\"total_count\":3238015},{\"a\":5,\"b\":9,\"win_count\":458830.43333335174,\"total_count\":3239523},{\"a\":5,\"b\":8,\"win_count\":463080.16666668945,\"total_count\":3239738},{\"a\":5,\"b\":7,\"win_count\":473045.9404762106,\"total_count\":3240440},{\"a\":5,\"b\":6,\"win_count\":483606.7750000254,\"total_count\":3235618},{\"a\":5,\"b\":5,\"win_count\":312124.2345238356,\"total_count\":3238421},{\"a\":5,\"b\":4,\"win_count\":525593.0571428204,\"total_count\":6478636},{\"a\":5,\"b\":3,\"win_count\":466006.2523809457,\"total_count\":6483942},{\"a\":5,\"b\":2,\"win_count\":398267.4178571703,\"total_count\":6477356},{\"a\":4,\"b\":14,\"win_count\":654491.3785713644,\"total_count\":3237862},{\"a\":4,\"b\":13,\"win_count\":575940.2809523673,\"total_count\":3241771},{\"a\":4,\"b\":12,\"win_count\":529635.5630952485,\"total_count\":3243388},{\"a\":4,\"b\":11,\"win_count\":492739.76785715215,\"total_count\":3239790},{\"a\":4,\"b\":10,\"win_count\":464150.73333333986,\"total_count\":3239522},{\"a\":4,\"b\":9,\"win_count\":431487.62023811927,\"total_count\":3236310},{\"a\":4,\"b\":8,\"win_count\":432433.0952381177,\"total_count\":3236039},{\"a\":4,\"b\":7,\"win_count\":441992.4845238338,\"total_count\":3241199},{\"a\":4,\"b\":6,\"win_count\":457909.2059524203,\"total_count\":3240543},{\"a\":4,\"b\":5,\"win_count\":470727.0059524096,\"total_count\":3236526},{\"a\":4,\"b\":4,\"win_count\":292079.8476190768,\"total_count\":3243112},{\"a\":4,\"b\":3,\"win_count\":434035.8321428702,\"total_count\":6481251},{\"a\":4,\"b\":2,\"win_count\":378929.945238136,\"total_count\":6476407},{\"a\":3,\"b\":14,\"win_count\":643799.5309523267,\"total_count\":3238380},{\"a\":3,\"b\":13,\"win_count\":568153.3309523712,\"total_count\":3240071},{\"a\":3,\"b\":12,\"win_count\":519625.3952381198,\"total_count\":3237785},{\"a\":3,\"b\":11,\"win_count\":483469.20476192643,\"total_count\":3238336},{\"a\":3,\"b\":10,\"win_count\":457455.75952382636,\"total_count\":3239426},{\"a\":3,\"b\":9,\"win_count\":427395.69285717065,\"total_count\":3239588},{\"a\":3,\"b\":8,\"win_count\":405943.78809526993,\"total_count\":3240789},{\"a\":3,\"b\":7,\"win_count\":411590.357142891,\"total_count\":3240432},{\"a\":3,\"b\":6,\"win_count\":427989.42023813626,\"total_count\":3237637},{\"a\":3,\"b\":5,\"win_count\":446137.41547622223,\"total_count\":3237572},{\"a\":3,\"b\":4,\"win_count\":432070.16428575356,\"total_count\":3240733},{\"a\":3,\"b\":3,\"win_count\":274778.80000002944,\"total_count\":3237664},{\"a\":3,\"b\":2,\"win_count\":350255.89642861957,\"total_count\":6480953},{\"a\":2,\"b\":14,\"win_count\":630407.0297618513,\"total_count\":3239012},{\"a\":2,\"b\":13,\"win_count\":559999.5773809584,\"total_count\":3238164},{\"a\":2,\"b\":12,\"win_count\":513750.79642860364,\"total_count\":3242169},{\"a\":2,\"b\":11,\"win_count\":477829.7821428849,\"total_count\":3238402},{\"a\":2,\"b\":10,\"win_count\":450618.70119049936,\"total_count\":3239642},{\"a\":2,\"b\":9,\"win_count\":421856.4023809875,\"total_count\":3241678},{\"a\":2,\"b\":8,\"win_count\":402520.81309527665,\"total_count\":3236165},{\"a\":2,\"b\":7,\"win_count\":384740.79642861325,\"total_count\":3237405},{\"a\":2,\"b\":6,\"win_count\":397082.5011905212,\"total_count\":3240297},{\"a\":2,\"b\":5,\"win_count\":414603.6511905198,\"total_count\":3238780},{\"a\":2,\"b\":4,\"win_count\":405199.1154762367,\"total_count\":3240898},{\"a\":2,\"b\":3,\"win_count\":395054.25714290637,\"total_count\":3240455},{\"a\":2,\"b\":2,\"win_count\":265137.2797619359,\"total_count\":3238920}]}")
	stat := &statistic.StatisticData{}
	err := json.Unmarshal(js, stat)
	if err != nil {
		return
	}
	for _, d := range stat.Data {
		v := s.data[d.A][d.B]
		v.winCount = d.WinCount
		v.totalCount = d.TotalCount
	}
}

func (s *Statistic) String() string {
	var sb strings.Builder
	_, _ = sb.WriteString("     Ao     Ko     Qo     Jo     To     9o     8o     7o     6o     5o     4o     3o     2o")
	for i := int32(14); i >= 2; i-- {
		m := s.data[i]
		switch i {
		case 14:
			_, _ = sb.WriteString("\nAs ")
		case 13:
			_, _ = sb.WriteString("\nKs ")
		case 12:
			_, _ = sb.WriteString("\nQs ")
		case 11:
			_, _ = sb.WriteString("\nJs ")
		case 10:
			_, _ = sb.WriteString("\nTs ")
		default:
			_, _ = sb.Write([]byte{'\n', '0' + byte(i), 's', ' '})
		}
		for j := int32(14); j >= 2; j-- {
			v := m[j]
			if v.totalCount == 0 {
				_, _ = sb.WriteString(fmt.Sprintf("%5.2f%% ", 50.0))
			} else {
				_, _ = sb.WriteString(fmt.Sprintf("%5.2f%% ", float64(v.winCount)/float64(v.totalCount)*100.0))
			}
		}
	}
	return sb.String()
}
