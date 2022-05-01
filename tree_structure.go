package treestructure

import (
	"github.com/caoshuyu/tree-structure/protofile/tsproto"
)

// TreeStructure 树结构
type treeStructure struct {
	tsHeader *tsproto.TreeHeader
}

// NewTreeStructure 新建结构体
func NewTreeStructure(header *tsproto.TreeHeader) *treeStructure {
	t := new(treeStructure)
	if header == nil {
		t.tsHeader = new(tsproto.TreeHeader)
	} else {
		t.tsHeader = header
	}
	return t
}

// AddTreeData 添加数据
func (t *treeStructure) AddTreeData(dataList []string) {
	for _, data := range dataList {
		t.tsHeader.C = t.addData(t.tsHeader.C, []rune(data))
	}
}

// DelTreeData 删除数据
func (t *treeStructure) DelTreeData(dataList []string) {
	haveKeyList := t.childData(t.tsHeader.C)
	haveKeyMap := make(map[string]bool)
	for k := range haveKeyList {
		haveKeyMap[haveKeyList[k]] = true
	}
	for _, data := range dataList {
		if _, h := haveKeyMap[data]; !h {
			return
		}
		delete(haveKeyMap, data)
	}
	bodyList := make([]*tsproto.TB, 0)
	for key := range haveKeyMap {
		bodyList = t.addData(bodyList, []rune(key))
	}
	t.tsHeader.C = bodyList
}

// GetTreeData 获取数据
func (t *treeStructure) GetTreeData() []string {
	return t.childData(t.tsHeader.C)
}

// GetProto 获取Proto文件
func (t *treeStructure) GetProto() *tsproto.TreeHeader {
	return t.tsHeader
}

func (t *treeStructure) addData(bodyList []*tsproto.TB, dataRune []rune) []*tsproto.TB {
	if len(dataRune) == 0 {
		return bodyList
	}
	isHave := false
	for k, body := range bodyList {
		bodyRune := []rune(body.B)
		if bodyRune[0] == dataRune[0] {
			isHave = true
			if len(dataRune) > 0 {
				bodyList[k].C = t.addData(bodyList[k].C, dataRune[1:])
			}
		}
	}
	if !isHave {
		bodyList = append(bodyList, t.buildDataTree(dataRune))
	}
	return bodyList
}

func (t *treeStructure) buildDataTree(dataRune []rune) *tsproto.TB {
	if len(dataRune) == 0 {
		return nil
	}
	body := new(tsproto.TB)
	body.B = string(dataRune[0])
	if len(dataRune) > 1 {
		body.C = append(body.C, t.buildDataTree(dataRune[1:]))
	}
	return body
}

func (t *treeStructure) childData(children []*tsproto.TB) []string {
	var dataList []string
	for _, one := range children {
		if len(one.C) > 0 {
			list := t.childData(one.C)
			for _, val := range list {
				dataList = append(dataList, one.B+val)
			}
		} else {
			dataList = append(dataList, one.B)
		}
	}
	return dataList
}
