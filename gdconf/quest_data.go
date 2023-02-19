package gdconf

import (
	"fmt"

	"hk4e/pkg/logger"

	"github.com/jszwec/csvutil"
)

// QuestData 任务配置表
type QuestData struct {
	QuestId       int32 `csv:"QuestId"`                 // ID
	ParentQuestId int32 `csv:"ParentQuestId,omitempty"` // 父任务ID
	Sequence      int32 `csv:"Sequence,omitempty"`      // 序列
	// 领取条件
	AcceptCondCompose     int32 `csv:"AcceptCondCompose,omitempty"`     // [领取条件]组合
	AcceptCondType1       int32 `csv:"AcceptCondType1,omitempty"`       // [领取条件]1类型
	AcceptCondType1Param1 int32 `csv:"AcceptCondType1Param1,omitempty"` // [领取条件]1参数1
	AcceptCondType1Param2 int32 `csv:"AcceptCondType1Param2,omitempty"` // [领取条件]1参数2
	AcceptCondType1Param3 int32 `csv:"AcceptCondType1Param3,omitempty"` // [领取条件]1参数3
	AcceptCondType2       int32 `csv:"AcceptCondType2,omitempty"`       // [领取条件]2类型
	AcceptCondType2Param1 int32 `csv:"AcceptCondType2Param1,omitempty"` // [领取条件]2参数1
	AcceptCondType2Param2 int32 `csv:"AcceptCondType2Param2,omitempty"` // [领取条件]2参数2
	AcceptCondType2Param3 int32 `csv:"AcceptCondType2Param3,omitempty"` // [领取条件]2参数3
	AcceptCondType3       int32 `csv:"AcceptCondType3,omitempty"`       // [领取条件]3类型
	AcceptCondType3Param1 int32 `csv:"AcceptCondType3Param1,omitempty"` // [领取条件]3参数1
	AcceptCondType3Param2 int32 `csv:"AcceptCondType3Param2,omitempty"` // [领取条件]3参数2
	AcceptCondType3Param3 int32 `csv:"AcceptCondType3Param3,omitempty"` // [领取条件]3参数3
	// 完成条件
	FinishCondCompose           int32  `csv:"FinishCondCompose,omitempty"`           // [完成条件]组合
	FinishCondType1             int32  `csv:"FinishCondType1,omitempty"`             // [完成条件]1类型
	FinishCondType1Param1       int32  `csv:"FinishCondType1Param1,omitempty"`       // [完成条件]1参数1
	FinishCondType1Param2       int32  `csv:"FinishCondType1Param2,omitempty"`       // [完成条件]1参数2
	FinishCondType1ComplexParam string `csv:"FinishCondType1ComplexParam,omitempty"` // [完成条件]1复杂参数
	FinishCondType1Count        int32  `csv:"FinishCondType1Count,omitempty"`        // [完成条件]1次数
	FinishCondType2             int32  `csv:"FinishCondType2,omitempty"`             // [完成条件]2类型
	FinishCondType2Param1       int32  `csv:"FinishCondType2Param1,omitempty"`       // [完成条件]2参数1
	FinishCondType2Param2       int32  `csv:"FinishCondType2Param2,omitempty"`       // [完成条件]2参数2
	FinishCondType2ComplexParam string `csv:"FinishCondType2ComplexParam,omitempty"` // [完成条件]2复杂参数
	FinishCondType2Count        int32  `csv:"FinishCondType2Count,omitempty"`        // [完成条件]2次数
	FinishCondType3             int32  `csv:"FinishCondType3,omitempty"`             // [完成条件]3类型
	FinishCondType3Param1       int32  `csv:"FinishCondType3Param1,omitempty"`       // [完成条件]3参数1
	FinishCondType3Param2       int32  `csv:"FinishCondType3Param2,omitempty"`       // [完成条件]3参数2
	FinishCondType3ComplexParam string `csv:"FinishCondType3ComplexParam,omitempty"` // [完成条件]3复杂参数
	FinishCondType3Count        int32  `csv:"FinishCondType3Count,omitempty"`        // [完成条件]3次数
	// 失败条件
	FailCondCompose           int32  `csv:"FailCondCompose,omitempty"`           // [失败条件]组合
	FailCondType1             int32  `csv:"FailCondType1,omitempty"`             // [失败条件]1类型
	FailCondType1Param1       int32  `csv:"FailCondType1Param1,omitempty"`       // [失败条件]1参数1
	FailCondType1Param2       int32  `csv:"FailCondType1Param2,omitempty"`       // [失败条件]1参数2
	FailCondType1ComplexParam string `csv:"FailCondType1ComplexParam,omitempty"` // [失败条件]1复杂参数
	FailCondType1Count        int32  `csv:"FailCondType1Count,omitempty"`        // [失败条件]1次数
	FailCondType2             int32  `csv:"FailCondType2,omitempty"`             // [失败条件]2类型
	FailCondType2Param1       int32  `csv:"FailCondType2Param1,omitempty"`       // [失败条件]2参数1
	FailCondType2Param2       int32  `csv:"FailCondType2Param2,omitempty"`       // [失败条件]2参数2
	FailCondType2ComplexParam string `csv:"FailCondType2ComplexParam,omitempty"` // [失败条件]2复杂参数
	FailCondType2Count        int32  `csv:"FailCondType2Count,omitempty"`        // [失败条件]2次数
	FailCondType3             int32  `csv:"FailCondType3,omitempty"`             // [失败条件]3类型
	FailCondType3Param1       int32  `csv:"FailCondType3Param1,omitempty"`       // [失败条件]3参数1
	FailCondType3Param2       int32  `csv:"FailCondType3Param2,omitempty"`       // [失败条件]3参数2
	FailCondType3ComplexParam string `csv:"FailCondType3ComplexParam,omitempty"` // [失败条件]3复杂参数
	FailCondType3Count        int32  `csv:"FailCondType3Count,omitempty"`        // [失败条件]3次数
}

func (g *GameDataConfig) loadQuestData() {
	g.QuestDataMap = make(map[int32]*QuestData)
	data := g.readCsvFileData("QuestData.csv")
	var questDataList []*QuestData
	err := csvutil.Unmarshal(data, &questDataList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, questData := range questDataList {
		// list -> map
		g.QuestDataMap[questData.QuestId] = questData
	}
	logger.Info("QuestData count: %v", len(g.QuestDataMap))
}

func GetQuestDataById(sceneId int32) *QuestData {
	return CONF.QuestDataMap[sceneId]
}

func GetQuestDataMap() map[int32]*QuestData {
	return CONF.QuestDataMap
}
